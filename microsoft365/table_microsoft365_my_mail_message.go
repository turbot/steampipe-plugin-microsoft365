package microsoft365

import (
	"context"
	"strings"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/messages"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION

func tableMicrosoft365MyMailMessage(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_my_mail_message",
		Description: "Retrieves messages in the specified user's mailbox.",
		List: &plugin.ListConfig{
			Hydrate: listMicrosoft365MyMailMessages,
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
		Get: &plugin.GetConfig{
			Hydrate:    getMicrosoft365MyMailMessage,
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Columns: mailMessageColumns(),
	}
}

//// LIST FUNCTION

func listMicrosoft365MyMailMessages(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_my_mail_message.listMicrosoft365MyMailMessages", "connection_error", err)
		return nil, err
	}

	getUserIDCached := plugin.HydrateFunc(getUserID).WithCache()
	userIDCached, err := getUserIDCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	userID := userIDCached.(string)

	// List operations
	input := &messages.MessagesRequestBuilderGetQueryParameters{}

	// Minimum value is 1 (this function isn't run if "limit 0" is specified)
	// Maximum value is unknown (tested up to 9999)
	pageSize := int64(9999)
	limit := d.QueryContext.Limit
	if limit != nil && *limit < pageSize {
		pageSize = *limit
	}
	input.Top = Int32(int32(pageSize))

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

	result, err := client.UsersById(userID).Messages().Get(ctx, options)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateMessageCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listMicrosoft365MyMailMessages", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
		message := pageItem.(models.Messageable)

		d.StreamListItem(ctx, &Microsoft365MailMessageInfo{message, userID})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.QueryStatus.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listMicrosoft365MyMailMessages", "paging_error", err)
		return nil, err
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getMicrosoft365MyMailMessage(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := d.KeyColumnQualString("id")
	if id == "" {
		return nil, nil
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_my_mail_message.getMicrosoft365MyMailMessage", "connection_error", err)
		return nil, err
	}

	getUserIDCached := plugin.HydrateFunc(getUserID).WithCache()
	userIDCached, err := getUserIDCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	userID := userIDCached.(string)

	// List operations
	input := &item.MessageItemRequestBuilderGetQueryParameters{}

	// Check for query context and requests only for queried columns
	givenColumns := d.QueryContext.Columns
	selectColumns := buildMailMessageRequestFields(ctx, givenColumns)
	input.Select = selectColumns

	options := &item.MessageItemRequestBuilderGetRequestConfiguration{
		QueryParameters: input,
	}

	result, err := client.UsersById(userID).MessagesById(id).Get(ctx, options)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	return &Microsoft365MailMessageInfo{result, userID}, nil
}
