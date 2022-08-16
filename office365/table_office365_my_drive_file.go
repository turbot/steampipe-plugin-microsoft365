package office365

import (
	"context"
	"fmt"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION

func tableOffice365MyDriveFile(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "office365_my_drive_file",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate:       listOffice365MyDriveFiles,
			ParentHydrate: listOffice365MyDrives,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Columns: driveFileColumns(),
	}
}

//// LIST FUNCTION

func listOffice365MyDriveFiles(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
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
		logger.Error("listOffice365MyDriveFiles", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(func(pageItem interface{}) bool {
		item := pageItem.(models.DriveItemable)

		d.StreamLeafListItem(ctx, &Office365DriveItemInfo{item, driveID})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.QueryStatus.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listOffice365MyDriveFiles", "paging_error", err)
		return nil, err
	}

	return nil, nil
}
