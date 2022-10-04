package microsoft365

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"

	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

func driveFileColumns() []*plugin.Column {
	return []*plugin.Column{
		{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the item (filename and extension).", Transform: transform.FromMethod("GetName")},
		{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier of the item within the Drive.", Transform: transform.FromMethod("GetId")},
		{Name: "path", Type: proto.ColumnType_STRING, Description: "URL that displays the resource in the browser.", Transform: transform.FromMethod("DriveItemFilePath")},
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
		{Name: "folder", Type: proto.ColumnType_JSON, Description: "Folder metadata, if the item is a folder.", Transform: transform.FromMethod("DriveItemFolder")},

		// Standard columns
		{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromMethod("GetName")},
		{Name: "user_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionUserID},
		{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Hydrate: plugin.HydrateFunc(getTenant).WithCache(), Transform: transform.FromValue()},
	}
}

//// TABLE DEFINITION

func tableMicrosoft365DriveFile(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_drive_file",
		Description: "Retrieves file's metadata or content owned by an user.",
		List: &plugin.ListConfig{
			Hydrate:       listMicrosoft365DriveFiles,
			ParentHydrate: listMicrosoft365Drives,
			KeyColumns:    plugin.SingleColumn("user_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getMicrosoft365DriveFile,
			KeyColumns: plugin.AllColumns([]string{"id", "drive_id", "user_id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Columns: driveFileColumns(),
	}
}

//// LIST FUNCTION

func listMicrosoft365DriveFiles(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	driveData := h.Item.(*Microsoft365DriveInfo)

	var driveID string
	if driveData != nil {
		driveID = *driveData.GetId()
	}
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_drive_file.listMicrosoft365DriveFiles", "connection_error", err)
		return nil, err
	}

	userID := d.KeyColumnQuals["user_id"].GetStringValue()
	result, err := client.UsersById(userID).DrivesById(driveID).Root().Children().Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateDriveItemCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listMicrosoft365DriveFiles", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
		var resultFiles []Microsoft365DriveItemInfo

		item := pageItem.(models.DriveItemable)

		resultFiles = append(resultFiles, Microsoft365DriveItemInfo{item, driveID, userID})
		if item.GetFolder() != nil && item.GetFolder().GetChildCount() != nil && *item.GetFolder().GetChildCount() != 0 {
			childData, err := expandDriveFolders(ctx, client, adapter, item, userID, driveID)
			if err != nil {
				return false
			}
			resultFiles = append(resultFiles, childData...)
		}

		for _, i := range resultFiles {
			d.StreamListItem(ctx, i)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				break
			}
		}

		return true
	})
	if err != nil {
		logger.Error("listMicrosoft365DriveFiles", "paging_error", err)
		return nil, err
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getMicrosoft365DriveFile(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	userID := d.KeyColumnQualString("user_id")
	driveID := d.KeyColumnQualString("drive_id")
	id := d.KeyColumnQualString("id")
	if userID == "" || driveID == "" || id == "" {
		return nil, nil
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_drive_file.listMicrosoft365DriveFiles", "connection_error", err)
		return nil, err
	}

	result, err := client.UsersById(userID).DrivesById(driveID).ItemsById(id).Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	return Microsoft365DriveItemInfo{result, driveID, userID}, nil
}

func expandDriveFolders(ctx context.Context, c *msgraphsdkgo.GraphServiceClient, a *msgraphsdkgo.GraphRequestAdapter, data models.DriveItemable, userID string, driveID string) ([]Microsoft365DriveItemInfo, error) {
	var items []Microsoft365DriveItemInfo
	result, err := c.UsersById(userID).DrivesById(driveID).ItemsById(*data.GetId()).Children().Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, a, models.CreateDriveItemCollectionResponseFromDiscriminatorValue)
	if err != nil {
		plugin.Logger(ctx).Error("expandDriveFolders", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
		var data []Microsoft365DriveItemInfo

		item := pageItem.(models.DriveItemable)

		data = append(data, Microsoft365DriveItemInfo{item, driveID, userID})
		if item.GetFolder() != nil && item.GetFolder().GetChildCount() != nil && *item.GetFolder().GetChildCount() != 0 {
			childData, err := expandDriveFolders(ctx, c, a, item, userID, driveID)
			if err != nil {
				return false
			}
			data = append(data, childData...)
		}
		items = append(items, data...)

		return true
	})
	if err != nil {
		plugin.Logger(ctx).Error("expandDriveFolders", "paging_error", err)
		return nil, err
	}

	return items, nil
}
