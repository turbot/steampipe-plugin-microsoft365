package office365

import (
	"context"
	"fmt"
	"strings"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/drives"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION

func tableOffice365MyDrive(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "office365_my_drive",
		Description: "Drives defined user's shared drives in the Google Drive.",
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
		logger.Error("office365_my_drive.listOffice365MyDrives", "connection_error", err)
		return nil, err
	}

	userIdentifier := getUserFromConfig(ctx, d, h)
	if userIdentifier == "" {
		return nil, fmt.Errorf("user_identifier must be set in the connection configuration")
	}

	input := &drives.DrivesRequestBuilderGetQueryParameters{}

	var queryFilter string
	equalQuals := d.KeyColumnQuals
	filter := buildDriveQueryFilter(equalQuals)

	if equalQuals["filter"] != nil {
		queryFilter = equalQuals["filter"].GetStringValue()
	}

	if queryFilter != "" {
		filter = append(filter, queryFilter)
	}

	if len(filter) > 0 {
		joinStr := strings.Join(filter, " and ")
		input.Filter = &joinStr
	}

	options := &drives.DrivesRequestBuilderGetRequestConfiguration{
		QueryParameters: input,
	}

	result, err := client.UsersById(userIdentifier).Drives().Get(ctx, options)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateDriveCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listOffice365MyDrives", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
		drive := pageItem.(models.Driveable)

		d.StreamListItem(ctx, &Office365DriveInfo{drive, userIdentifier})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.QueryStatus.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listOffice365MyDrives", "paging_error", err)
		return nil, err
	}

	return nil, nil
}
