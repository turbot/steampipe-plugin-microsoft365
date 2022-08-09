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

//// TABLE DEFINITION

func tableOffice365DriveFiles() *plugin.Table {
	return &plugin.Table{
		Name:        "office365_drive_file",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listOffice365DriveFiles,
			// ParentHydrate: listOffice365Drives,
			KeyColumns: plugin.SingleColumn("user_identifier"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},

		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetName")},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetId")},
			{Name: "drive_id", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromField("DriveID")},
			{Name: "web_url", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetWebUrl")},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetDescription")},
			{Name: "created_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "", Transform: transform.FromMethod("GetCreatedDateTime")},
			{Name: "last_modified_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "", Transform: transform.FromMethod("GetLastModifiedDateTime")},
			{Name: "etag", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetETag")},
			{Name: "size", Type: proto.ColumnType_INT, Description: "", Transform: transform.FromMethod("GetSize")},
			{Name: "web_dav_url", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetWebDavUrl")},
			{Name: "ctag", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetCTag")},

			{Name: "created_by", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("DriveItemCreatedBy")},
			{Name: "last_modified_by", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("DriveItemLastModifiedBy")},
			{Name: "parent_Reference", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("DriveItemParentReference")},
			{Name: "file", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("DriveItemFile")},

			// Standard columns
			{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromMethod("GetName")},
			{Name: "user_identifier", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromQual("user_identifier")},
			{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Hydrate: plugin.HydrateFunc(getTenant).WithCache(), Transform: transform.FromValue()},
		},
	}
}

//// LIST FUNCTION

func listOffice365DriveFiles(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// driveData := h.Item.(*Office365DriveInfo)

	// var driveID string
	// if driveData != nil {
	// 	driveID = *driveData.GetId()
	// }

	logger := plugin.Logger(ctx)
	driveID := "b!9Gsmy9eA9UueQZNqu5gVXeJOR1WJQBhNnG6bef2-YkjFaNz_gf8ERI4H7WKA3uv4"

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

		d.StreamListItem(ctx, &Office365DriveItemInfo{item, driveID})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.QueryStatus.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listOffice365DriveFiles", "paging_error", err)
		return nil, err
	}

	return nil, nil
}
