package microsoft365

import (
	"context"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableMicrosoft365MyDriveFile(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_my_drive_file",
		Description: "Retrieves file's metadata or content owned by an user.",
		List: &plugin.ListConfig{
			Hydrate:       listMicrosoft365MyDriveFiles,
			ParentHydrate: listMicrosoft365MyDrives,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"itemNotFound"}),
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getMicrosoft365MyDriveFile,
			KeyColumns: plugin.AllColumns([]string{"drive_id", "id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"itemNotFound"}),
			},
		},
		Columns: driveFileColumns(),
	}
}

//// LIST FUNCTION

func listMicrosoft365MyDriveFiles(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	driveData := h.Item.(*Microsoft365DriveInfo)

	var driveID string
	if driveData != nil {
		driveID = *driveData.GetId()
	}
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_my_drive_file.listMicrosoft365MyDriveFiles", "connection_error", err)
		return nil, err
	}

	getUserIDCached := plugin.HydrateFunc(getUserID).WithCache()
	userIDCached, err := getUserIDCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	userID := userIDCached.(string)

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
		logger.Error("listMicrosoft365MyDriveFiles", "create_iterator_instance_error", err)
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
		logger.Error("listMicrosoft365MyDriveFiles", "paging_error", err)
		return nil, err
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getMicrosoft365MyDriveFile(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	driveID := d.EqualsQualString("drive_id")
	id := d.EqualsQualString("id")
	if driveID == "" || id == "" {
		return nil, nil
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_my_drive_file.getMicrosoft365MyDriveFile", "connection_error", err)
		return nil, err
	}

	getUserIDCached := plugin.HydrateFunc(getUserID).WithCache()
	userIDCached, err := getUserIDCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	userID := userIDCached.(string)

	result, err := client.Drives().ByDriveId(driveID).Items().ByDriveItemId(id).Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	return Microsoft365DriveItemInfo{result, driveID, userID}, nil
}
