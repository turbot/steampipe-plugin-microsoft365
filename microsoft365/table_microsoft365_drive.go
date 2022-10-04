package microsoft365

import (
	"context"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/drives"
)

func driveColumns() []*plugin.Column {
	return []*plugin.Column{
		{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the item.", Transform: transform.FromMethod("GetName")},
		{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier of the drive.", Transform: transform.FromMethod("GetId")},
		{Name: "drive_type", Type: proto.ColumnType_STRING, Description: "Describes the type of drive represented by this resource. OneDrive personal drives will return personal. OneDrive for Business will return business. SharePoint document libraries will return documentLibrary.", Transform: transform.FromMethod("GetDriveType")},
		{Name: "web_url", Type: proto.ColumnType_STRING, Description: "URL that displays the resource in the browser.", Transform: transform.FromMethod("GetWebUrl")},
		{Name: "description", Type: proto.ColumnType_STRING, Description: "Provide a user-visible description of the drive.", Transform: transform.FromMethod("GetDescription")},
		{Name: "created_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time of item creation.", Transform: transform.FromMethod("GetCreatedDateTime")},

		// Other Fields
		{Name: "etag", Type: proto.ColumnType_STRING, Description: "Specifies the eTag.", Transform: transform.FromMethod("GetETag")},
		{Name: "last_modified_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time the item was last modified.", Transform: transform.FromMethod("GetLastModifiedDateTime")},

		// JSON Fields
		{Name: "created_by", Type: proto.ColumnType_JSON, Description: "Identity of the user, device, or application which created the item.", Transform: transform.FromMethod("DriveCreatedBy")},
		{Name: "last_modified_by", Type: proto.ColumnType_JSON, Description: "Identity of the user, device, and application which last modified the item.", Transform: transform.FromMethod("DriveLastModifiedBy")},
		{Name: "parent_reference", Type: proto.ColumnType_JSON, Description: "Parent information, if the drive has a parent.", Transform: transform.FromMethod("DriveParentReference")},

		// Standard columns
		{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromMethod("GetName")},
		{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Hydrate: plugin.HydrateFunc(getTenant).WithCache(), Transform: transform.FromValue()},
		{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Odata query to search for resources."},
		{Name: "user_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionUserID},
	}
}

//// TABLE DEFINITION

func tableMicrosoft365Drive(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_drive",
		Description: "Drives defined user's shared drives in the Google Drive.",
		List: &plugin.ListConfig{
			Hydrate: listMicrosoft365Drives,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "user_id", Require: plugin.Required},
				{Name: "filter", Require: plugin.Optional},
				{Name: "id", Require: plugin.Optional},
				{Name: "name", Require: plugin.Optional},
				{Name: "drive_type", Require: plugin.Optional},
			},
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getMicrosoft365Drive,
			KeyColumns: plugin.AllColumns([]string{"id", "user_id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Columns: driveColumns(),
	}
}

//// LIST FUNCTION

func listMicrosoft365Drives(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_drive.listMicrosoft365Drives", "connection_error", err)
		return nil, err
	}

	input := &drives.DrivesRequestBuilderGetQueryParameters{}

	var queryFilter string
	equalQuals := d.KeyColumnQuals
	filter := buildDriveQueryFilter(equalQuals)

	if equalQuals["filter"] != nil {
		queryFilter = equalQuals["filter"].GetStringValue()
	}

	if queryFilter != "" {
		filter = append(filter, queryFilter)
	}

	if len(filter) > 0 {
		joinStr := strings.Join(filter, " and ")
		input.Filter = &joinStr
	}

	options := &drives.DrivesRequestBuilderGetRequestConfiguration{
		QueryParameters: input,
	}

	userID := d.KeyColumnQuals["user_id"].GetStringValue()
	result, err := client.UsersById(userID).Drives().Get(ctx, options)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateDriveCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listMicrosoft365Drives", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
		drive := pageItem.(models.Driveable)

		d.StreamListItem(ctx, &Microsoft365DriveInfo{drive, userID})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.QueryStatus.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listMicrosoft365Drives", "paging_error", err)
		return nil, err
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getMicrosoft365Drive(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	driveID := d.KeyColumnQualString("id")
	userID := d.KeyColumnQualString("user_id")
	if driveID == "" || userID == "" {
		return nil, nil
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_my_drive.listMicrosoft365MyDrives", "connection_error", err)
		return nil, err
	}

	result, err := client.UsersById(userID).DrivesById(driveID).Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	return &Microsoft365DriveInfo{result, userID}, nil
}

func buildDriveQueryFilter(equalQuals plugin.KeyColumnEqualsQualMap) []string {
	filters := []string{}

	filterQuals := map[string]string{
		"name":       "string",
		"id":         "string",
		"drive_type": "string",
	}

	for qual, qualType := range filterQuals {
		switch qualType {
		case "string":
			if equalQuals[qual] != nil {
				filters = append(filters, fmt.Sprintf("%s eq '%s'", strcase.ToCamel(qual), equalQuals[qual].GetStringValue()))
			}
		}
	}
	return filters
}
