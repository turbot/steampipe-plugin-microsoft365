package microsoft365

import (
	"context"
	"fmt"
	"strings"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/messages"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION

func tableMicrosoft365MailMyMessage(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_mail_my_message",
		Description: "Retrieves messages in the specified user's mailbox.",
		List: &plugin.ListConfig{
			Hydrate: listMicrosoft365MailMyMessages,
			KeyColumns: plugin.KeyColumnSlice{
				// Key fields
				{Name: "filter", Require: plugin.Optional},

				// Other fields for filtering OData
				{Name: "subject", Require: plugin.Optional},
				{Name: "has_attachments", Require: plugin.Optional, Operators: []string{"<>", "="}},
				{Name: "is_read", Require: plugin.Optional, Operators: []string{"<>", "="}},
				{Name: "is_draft", Require: plugin.Optional, Operators: []string{"<>", "="}},
			},
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Columns: mailMessageColumns(),
	}
}

//// LIST FUNCTION

func listMicrosoft365MailMyMessages(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_mail_my_message.listMicrosoft365MailMyMessages", "connection_error", err)
		return nil, err
	}

	userIdentifier := getUserFromConfig(ctx, d, h)
	if userIdentifier == "" {
		return nil, fmt.Errorf("user_identifier must be set in the connection configuration")
	}

	// List operations
	input := &messages.MessagesRequestBuilderGetQueryParameters{}

	// Restrict the limit value to be passed in the query parameter which is not between 1 and 999, otherwise API will throw an error as follow
	limit := d.QueryContext.Limit
	if limit != nil {
		if *limit > 0 && *limit <= 999 {
			l := int32(*limit)
			input.Top = &l
		}
	}

	// Check for query context and requests only for queried columns
	givenColumns := d.QueryContext.Columns
	selectColumns := buildMailMessageRequestFields(ctx, givenColumns)
	input.Select = selectColumns

	equalQuals := d.KeyColumnQuals
	quals := d.Quals

	var queryFilter string
	filter := buildQueryFilter(equalQuals)
	filter = append(filter, buildBoolNEFilter(quals)...)

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

	options := &messages.MessagesRequestBuilderGetRequestConfiguration{
		QueryParameters: input,
	}

	result, err := client.UsersById(userIdentifier).Messages().Get(ctx, options)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateMessageCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listMicrosoft365MailMyMessages", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
		message := pageItem.(models.Messageable)

		d.StreamListItem(ctx, &Microsoft365MailMessageInfo{message, userIdentifier})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.QueryStatus.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listMicrosoft365MailMyMessages", "paging_error", err)
		return nil, err
	}

	return nil, nil
}
