package microsoft365

import (
	"context"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION

func tableMicrosoft365DriveMyFile(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_drive_my_file",
		Description: "Retrieves file's metadata or content owned by an user.",
		List: &plugin.ListConfig{
			Hydrate:       listMicrosoft365DriveMyFiles,
			ParentHydrate: listMicrosoft365MyDrives,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getMicrosoft365DriveMyFile,
			KeyColumns: plugin.AllColumns([]string{"drive_id", "id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Columns: driveFileColumns(),
	}
}

//// LIST FUNCTION

func listMicrosoft365DriveMyFiles(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	driveData := h.Item.(*Microsoft365DriveInfo)

	var driveID string
	if driveData != nil {
		driveID = *driveData.GetId()
	}
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_drive_my_file.listMicrosoft365DriveMyFiles", "connection_error", err)
		return nil, err
	}

	getUserIdentifierCached := plugin.HydrateFunc(getUserIdentifier).WithCache()
	userID, err := getUserIdentifierCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	userIdentifier := userID.(string)

	result, err := client.UsersById(userIdentifier).DrivesById(driveID).Root().Children().Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateDriveItemCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listMicrosoft365DriveMyFiles", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
		var resultFiles []Microsoft365DriveItemInfo

		item := pageItem.(models.DriveItemable)

		resultFiles = append(resultFiles, Microsoft365DriveItemInfo{item, driveID, userIdentifier})
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
		logger.Error("listMicrosoft365DriveMyFiles", "paging_error", err)
		return nil, err
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getMicrosoft365DriveMyFile(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	driveID := d.KeyColumnQualString("drive_id")
	id := d.KeyColumnQualString("id")
	if driveID == "" || id == "" {
		return nil, nil
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_drive_my_file.getMicrosoft365DriveMyFile", "connection_error", err)
		return nil, err
	}

	getUserIdentifierCached := plugin.HydrateFunc(getUserIdentifier).WithCache()
	userID, err := getUserIdentifierCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	userIdentifier := userID.(string)

	result, err := client.UsersById(userIdentifier).DrivesById(driveID).ItemsById(id).Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	return Microsoft365DriveItemInfo{result, driveID, userIdentifier}, nil
}
