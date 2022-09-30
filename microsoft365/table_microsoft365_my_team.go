package microsoft365

import (
	"context"
	"fmt"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION

func tableMicrosoft365MyTeam(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_my_team",
		Description: "Retrieves the teams that the specified user is a direct member of.",
		List: &plugin.ListConfig{
			Hydrate: listMicrosoft365MyTeams,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Columns: teamColumns(),
	}
}

//// LIST FUNCTION

func listMicrosoft365MyTeams(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_my_team.listMicrosoft365MyTeams", "connection_error", err)
		return nil, err
	}

	userIdentifier := getUserFromConfig(ctx, d, h)
	if userIdentifier == "" {
		return nil, fmt.Errorf("user_identifier must be set in the connection configuration")
	}

	result, err := client.UsersById(userIdentifier).JoinedTeams().Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateTeamCollectionResponseFromDiscriminatorValue)
	if err != nil {
		plugin.Logger(ctx).Error("microsoft365_teams.listMicrosoft365MyTeams", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
		team := pageItem.(models.Teamable)

		d.StreamListItem(ctx, &Microsoft365TeamInfo{team, userIdentifier})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.QueryStatus.RowsRemaining(ctx) != 0
	})
	if err != nil {
		plugin.Logger(ctx).Error("microsoft365_teams.listMicrosoft365MyTeams", "paging_error", err)
		return nil, err
	}

	return nil, nil
}