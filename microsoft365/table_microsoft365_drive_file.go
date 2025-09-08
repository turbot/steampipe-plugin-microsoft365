package microsoft365

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

func driveFileColumns() []*plugin.Column {
	return commonColumns([]*plugin.Column{
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
	})
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
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"itemNotFound"}),
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getMicrosoft365DriveFile,
			KeyColumns: plugin.AllColumns([]string{"id", "drive_id", "user_id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"itemNotFound"}),
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

	userID := d.EqualsQuals["user_id"].GetStringValue()

	// Get all items with a filter that includes everything
	input := &drives.ItemItemsRequestBuilderGetQueryParameters{
		Filter: StringPtr("folder ne null or file ne null"),
	}
	options := &drives.ItemItemsRequestBuilderGetRequestConfiguration{
		QueryParameters: input,
	}
	result, err := client.Drives().ByDriveId(driveID).Items().Get(ctx, options)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator[models.DriveItemable](result, adapter, models.CreateDriveItemCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listMicrosoft365DriveFiles", "create_iterator_instance_error", err)
		return nil, err
	}

	// Track processed items to avoid duplicates
	processedItems := make(map[string]bool)

	err = pageIterator.Iterate(ctx, func(pageItem models.DriveItemable) bool {
		item := pageItem
		itemID := *item.GetId()

		// Skip if we've already processed this item
		if processedItems[itemID] {
			return true
		}

		// Mark as processed
		processedItems[itemID] = true

		// Stream the current item
		d.StreamListItem(ctx, Microsoft365DriveItemInfo{item, driveID, userID})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return false
		}

		// If this is a folder with children, expand it
		if item.GetFolder() != nil && item.GetFolder().GetChildCount() != nil && *item.GetFolder().GetChildCount() != 0 {
			childData, err := expandDriveFolders(ctx, client, adapter, item, userID, driveID)
			if err != nil {
				return false
			}

			// Stream all child items
			for _, childItem := range childData {
				childID := *childItem.DriveItemable.GetId()

				// Skip if we've already processed this child item
				if processedItems[childID] {
					continue
				}

				// Mark as processed
				processedItems[childID] = true

				d.StreamListItem(ctx, childItem)

				// Context can be cancelled due to manual cancellation or the limit has been hit
				if d.RowsRemaining(ctx) == 0 {
					return false
				}
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

	userID := d.EqualsQualString("user_id")
	driveID := d.EqualsQualString("drive_id")
	id := d.EqualsQualString("id")
	if userID == "" || driveID == "" || id == "" {
		return nil, nil
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_drive_file.listMicrosoft365DriveFiles", "connection_error", err)
		return nil, err
	}

	result, err := client.Drives().ByDriveId(driveID).Items().ByDriveItemId(id).Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	return Microsoft365DriveItemInfo{result, driveID, userID}, nil
}

func expandDriveFolders(ctx context.Context, c *msgraphsdkgo.GraphServiceClient, a *msgraphsdkgo.GraphRequestAdapter, data models.DriveItemable, userID string, driveID string) ([]Microsoft365DriveItemInfo, error) {
	var items []Microsoft365DriveItemInfo
	result, err := c.Drives().ByDriveId(driveID).Items().ByDriveItemId(*data.GetId()).Children().Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator[models.DriveItemable](result, a, models.CreateDriveItemCollectionResponseFromDiscriminatorValue)
	if err != nil {
		plugin.Logger(ctx).Error("expandDriveFolders", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem models.DriveItemable) bool {
		var data []Microsoft365DriveItemInfo

		item := pageItem

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
