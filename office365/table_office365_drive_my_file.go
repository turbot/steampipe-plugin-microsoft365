package office365

import (
	"context"
	"fmt"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION

func tableOffice365DriveMyFile(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "office365_drive_my_file",
		Description: "Retrieves file's metadata or content owned by an user.",
		List: &plugin.ListConfig{
			Hydrate:       listOffice365DriveMyFiles,
			ParentHydrate: listOffice365MyDrives,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Columns: driveFileColumns(),
	}
}

//// LIST FUNCTION

func listOffice365DriveMyFiles(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	driveData := h.Item.(*Office365DriveInfo)

	var driveID string
	if driveData != nil {
		driveID = *driveData.GetId()
	}
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("office365_drive_my_file.listOffice365DriveMyFiles", "connection_error", err)
		return nil, err
	}

	userIdentifier := getUserFromConfig(ctx, d, h)
	if userIdentifier == "" {
		return nil, fmt.Errorf("userIdentifier must be set in the connection configuration")
	}

	result, err := client.UsersById(userIdentifier).DrivesById(driveID).Root().Children().Get()
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateDriveItemCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listOffice365DriveMyFiles", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(func(pageItem interface{}) bool {
		var resultFiles []Office365DriveItemInfo

		item := pageItem.(models.DriveItemable)

		resultFiles = append(resultFiles, Office365DriveItemInfo{item, driveID, userIdentifier})
		if item.GetFolder() != nil && item.GetFolder().GetChildCount() != nil && *item.GetFolder().GetChildCount() != 0 {
			childData, err := expandDriveFolders(ctx, client, adapter, item, userIdentifier, driveID)
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
		logger.Error("listOffice365DriveMyFiles", "paging_error", err)
		return nil, err
	}

	return nil, nil
}
