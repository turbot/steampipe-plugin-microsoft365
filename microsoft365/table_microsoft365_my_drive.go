package microsoft365

import (
	"context"
	"strings"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableMicrosoft365MyDrive(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_my_drive",
		Description: "Drives defined user's shared drives in the Google Drive.",
		List: &plugin.ListConfig{
			Hydrate: listMicrosoft365MyDrives,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getMicrosoft365MyDrive,
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"itemNotFound"}),
			},
		},
		Columns: driveColumns(),
	}
}

//// LIST FUNCTION

func listMicrosoft365MyDrives(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_my_drive.listMicrosoft365MyDrives", "connection_error", err)
		return nil, err
	}

	getUserIDCached := plugin.HydrateFunc(getUserID).WithCache()
	userIDCached, err := getUserIDCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	userID := userIDCached.(string)

	input := &users.ItemDrivesRequestBuilderGetQueryParameters{}

	// Minimum value is 1 (this function isn't run if "limit 0" is specified)
	// Maximum value is unknown (tested up to 9999)
	pageSize := int64(9999)
	limit := d.QueryContext.Limit
	if limit != nil && *limit < pageSize {
		pageSize = *limit
	}
	input.Top = Int32(int32(pageSize))

	var queryFilter string
	equalQuals := d.EqualsQuals
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

	options := &users.ItemDrivesRequestBuilderGetRequestConfiguration{
		QueryParameters: input,
	}

	result, err := client.Users().ByUserId(userID).Drives().Get(ctx, options)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator[models.Driveable](result, adapter, models.CreateDriveCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listMicrosoft365MyDrives", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem models.Driveable) bool {
		drive := pageItem

		d.StreamListItem(ctx, &Microsoft365DriveInfo{drive, userID})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listMicrosoft365MyDrives", "paging_error", err)
		return nil, err
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getMicrosoft365MyDrive(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	driveID := d.EqualsQualString("id")
	if driveID == "" {
		return nil, nil
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_my_drive.getMicrosoft365MyDrive", "connection_error", err)
		return nil, err
	}

	getUserIDCached := plugin.HydrateFunc(getUserID).WithCache()
	userIDCached, err := getUserIDCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	userID := userIDCached.(string)

	result, err := client.Users().ByUserId(userID).Drives().ByDriveId(driveID).Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	return &Microsoft365DriveInfo{result, userID}, nil
}
