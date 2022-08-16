package office365

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func driveColumns() []*plugin.Column {
	return []*plugin.Column{
		{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier of the drive.", Transform: transform.FromMethod("GetId")},
		{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the item.", Transform: transform.FromMethod("GetName")},
		{Name: "user_identifier", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromQual("user_identifier")},
		{Name: "drive_type", Type: proto.ColumnType_STRING, Description: "Describes the type of drive represented by this resource. OneDrive personal drives will return personal. OneDrive for Business will return business. SharePoint document libraries will return documentLibrary.", Transform: transform.FromMethod("GetDriveType")},
		{Name: "web_url", Type: proto.ColumnType_STRING, Description: "URL that displays the resource in the browser.", Transform: transform.FromMethod("GetWebUrl")},
		{Name: "description", Type: proto.ColumnType_STRING, Description: "Provide a user-visible description of the drive.", Transform: transform.FromMethod("GetDescription")},
		{Name: "etag", Type: proto.ColumnType_STRING, Description: "Specifies the eTag.", Transform: transform.FromMethod("GetETag")},

		// Standard columns
		{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromMethod("GetName")},
		{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Hydrate: plugin.HydrateFunc(getTenant).WithCache(), Transform: transform.FromValue()},
	}
}

//// TABLE DEFINITION

func tableOffice365Drive(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "office365_drive",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate:    listOffice365Drives,
			KeyColumns: plugin.SingleColumn("user_identifier"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Columns: driveColumns(),
	}
}

//// LIST FUNCTION

func listOffice365Drives(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("error creating client: %v", err)
	}

	userIdentifier := d.KeyColumnQuals["user_identifier"].GetStringValue()
	result, err := client.UsersById(userIdentifier).Drives().Get()
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateDriveCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listOffice365Drives", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(func(pageItem interface{}) bool {
		drive := pageItem.(models.Driveable)

		d.StreamListItem(ctx, &Office365DriveInfo{drive})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.QueryStatus.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listOffice365Drives", "paging_error", err)
		return nil, err
	}

	return nil, nil
}
