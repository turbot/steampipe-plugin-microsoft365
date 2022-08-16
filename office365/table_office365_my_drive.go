package office365

import (
	"context"
	"fmt"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION

func tableOffice365MyDrive(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "office365_my_drive",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listOffice365MyDrives,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Columns: driveColumns(),
	}
}

//// LIST FUNCTION

func listOffice365MyDrives(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
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

	result, err := client.UsersById(userIdentifier).Drives().Get()
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateDriveCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listOffice365MyDrives", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(func(pageItem interface{}) bool {
		drive := pageItem.(models.Driveable)

		d.StreamListItem(ctx, &Office365DriveInfo{drive})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.QueryStatus.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listOffice365MyDrives", "paging_error", err)
		return nil, err
	}

	return nil, nil
}
