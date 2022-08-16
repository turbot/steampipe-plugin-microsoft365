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

func driveFileColumns() []*plugin.Column {
	return []*plugin.Column{
		{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the item (filename and extension).", Transform: transform.FromMethod("GetName")},
		{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier of the item within the Drive.", Transform: transform.FromMethod("GetId")},
		{Name: "drive_id", Type: proto.ColumnType_STRING, Description: "The unique id of the drive.", Transform: transform.FromField("DriveID")},
		{Name: "web_url", Type: proto.ColumnType_STRING, Description: "URL that displays the resource in the browser.", Transform: transform.FromMethod("GetWebUrl")},
		{Name: "description", Type: proto.ColumnType_STRING, Description: "Provides a user-visible description of the item.", Transform: transform.FromMethod("GetDescription")},
		{Name: "created_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time of item creation.", Transform: transform.FromMethod("GetCreatedDateTime")},
		{Name: "last_modified_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time the item was last modified.", Transform: transform.FromMethod("GetLastModifiedDateTime")},
		{Name: "etag", Type: proto.ColumnType_STRING, Description: "ETag for the entire item (metadata + content).", Transform: transform.FromMethod("GetETag")},
		{Name: "size", Type: proto.ColumnType_INT, Description: "Size of the item in bytes.", Transform: transform.FromMethod("GetSize")},
		{Name: "web_dav_url", Type: proto.ColumnType_STRING, Description: "WebDAV compatible URL for the item.", Transform: transform.FromMethod("GetWebDavUrl")},
		{Name: "ctag", Type: proto.ColumnType_STRING, Description: "An eTag for the content of the item. This eTag is not changed if only the metadata is changed. This property is not returned if the item is a folder.", Transform: transform.FromMethod("GetCTag")},
		{Name: "created_by", Type: proto.ColumnType_JSON, Description: "Identity of the user, device, and application which created the item.", Transform: transform.FromMethod("DriveItemCreatedBy")},
		{Name: "last_modified_by", Type: proto.ColumnType_JSON, Description: "Identity of the user, device, and application which last modified the item.", Transform: transform.FromMethod("DriveItemLastModifiedBy")},
		{Name: "parent_Reference", Type: proto.ColumnType_JSON, Description: "Parent information, if the item has a parent.", Transform: transform.FromMethod("DriveItemParentReference")},
		{Name: "file", Type: proto.ColumnType_JSON, Description: "File metadata, if the item is a file.", Transform: transform.FromMethod("DriveItemFile")},

		// Standard columns
		{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromMethod("GetName")},
		{Name: "user_identifier", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromQual("user_identifier")},
		{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Hydrate: plugin.HydrateFunc(getTenant).WithCache(), Transform: transform.FromValue()},
	}
}

//// TABLE DEFINITION

func tableOffice365DriveFile(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "office365_drive_file",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate:       listOffice365DriveFiles,
			ParentHydrate: listOffice365Drives,
			KeyColumns:    plugin.SingleColumn("user_identifier"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Columns: driveFileColumns(),
	}
}

//// LIST FUNCTION

func listOffice365DriveFiles(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	driveData := h.Item.(*Office365DriveInfo)

	var driveID string
	if driveData != nil {
		driveID = *driveData.GetId()
	}

	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("error creating client: %v", err)
	}

	userIdentifier := d.KeyColumnQuals["user_identifier"].GetStringValue()
	result, err := client.UsersById(userIdentifier).DrivesById(driveID).Root().Children().Get()
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateDriveItemCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listOffice365DriveFiles", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(func(pageItem interface{}) bool {
		item := pageItem.(models.DriveItemable)

		d.StreamLeafListItem(ctx, &Office365DriveItemInfo{item, driveID})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.QueryStatus.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listOffice365DriveFiles", "paging_error", err)
		return nil, err
	}

	return nil, nil
}
