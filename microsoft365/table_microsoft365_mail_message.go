package microsoft365

import (
	"context"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/messages"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item"
)

func mailMessageColumns() []*plugin.Column {
	return commonColumns([]*plugin.Column{
		{Name: "subject", Type: proto.ColumnType_STRING, Description: "The subject of the message.", Transform: transform.FromMethod("GetSubject")},
		{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier for the message.", Transform: transform.FromMethod("GetId")},
		{Name: "body_preview", Type: proto.ColumnType_STRING, Description: "The first 255 characters of the message body in text format.", Transform: transform.FromMethod("GetBodyPreview")},
		{Name: "is_read", Type: proto.ColumnType_BOOL, Description: "Indicates whether the message has been read.", Transform: transform.FromMethod("GetIsRead")},
		{Name: "has_attachments", Type: proto.ColumnType_BOOL, Description: "Indicates whether the message has attachments.", Transform: transform.FromMethod("GetHasAttachments")},
		{Name: "is_draft", Type: proto.ColumnType_BOOL, Description: "Indicates whether the message is a draft. A message is a draft if it hasn't been sent yet.", Transform: transform.FromMethod("GetIsDraft")},
		{Name: "received_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time the message was received.", Transform: transform.FromMethod("GetReceivedDateTime")},
		{Name: "sent_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time the message was sent.", Transform: transform.FromMethod("GetSentDateTime")},

		// Other columns
		{Name: "importance", Type: proto.ColumnType_STRING, Description: "The importance of the message. The possible values are: low, normal, and high.", Transform: transform.FromMethod("MessageImportance")},
		{Name: "inference_classification", Type: proto.ColumnType_STRING, Description: "The classification of the message for the user, based on inferred relevance or importance, or on an explicit override. The possible values are: focused or other.", Transform: transform.FromMethod("MessageInferenceClassification")},
		{Name: "change_key", Type: proto.ColumnType_STRING, Description: "The version of the message.", Transform: transform.FromMethod("GetChangeKey")},
		{Name: "conversation_id", Type: proto.ColumnType_STRING, Description: "The ID of the conversation the email belongs to.", Transform: transform.FromMethod("GetConversationId")},
		{Name: "internet_message_id", Type: proto.ColumnType_STRING, Description: "The message ID in the format specified by RFC2822.", Transform: transform.FromMethod("GetInternetMessageId")},
		{Name: "is_read_receipt_requested", Type: proto.ColumnType_BOOL, Description: "Indicates whether a read receipt is requested for the message.", Transform: transform.FromMethod("GetIsReadReceiptRequested")},
		{Name: "is_delivery_receipt_requested", Type: proto.ColumnType_BOOL, Description: "Indicates whether a read receipt is requested for the message.", Transform: transform.FromMethod("GetIsDeliveryReceiptRequested")},
		{Name: "parent_folder_id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the message's parent mailFolder.", Transform: transform.FromMethod("GetParentFolderId")},

		// Other fields
		{Name: "created_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "", Transform: transform.FromMethod("GetCreatedDateTime")},
		{Name: "last_modified_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "", Transform: transform.FromMethod("GetLastModifiedDateTime")},
		{Name: "web_link", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetWebLink")},

		// JSON fields
		{Name: "sender", Type: proto.ColumnType_JSON, Description: "The date and time the message was created.", Transform: transform.FromMethod("MessageSender")},
		{Name: "from", Type: proto.ColumnType_JSON, Description: "The owner of the mailbox from which the message is sent.", Transform: transform.FromMethod("MessageFrom")},
		{Name: "body", Type: proto.ColumnType_JSON, Description: "The body of the message. It can be in HTML or text format.", Transform: transform.FromMethod("MessageBody")},
		{Name: "categories", Type: proto.ColumnType_JSON, Description: "The categories associated with the message.", Transform: transform.FromMethod("GetCategories")},
		{Name: "reply_to", Type: proto.ColumnType_JSON, Description: "The email addresses to use when replying.", Transform: transform.FromMethod("MessageReplyTo")},
		{Name: "to_recipients", Type: proto.ColumnType_JSON, Description: "The To: recipients for the message.", Transform: transform.FromMethod("MessageToRecipients")},
		{Name: "attachments", Type: proto.ColumnType_JSON, Description: "The attachments of the message.", Transform: transform.FromMethod("MessageAttachments")},
		{Name: "bcc_recipients", Type: proto.ColumnType_JSON, Description: "The Bcc: recipients for the message.", Transform: transform.FromMethod("MessageBccRecipients")},
		{Name: "cc_recipients", Type: proto.ColumnType_JSON, Description: "The Cc: recipients for the message.", Transform: transform.FromMethod("MessageCcRecipients")},

		// Standard columns
		{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromMethod("GetSubject")},
		{Name: "user_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionUserID},
		{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Odata query to search for resources."},
	})
}

//// TABLE DEFINITION

func tableMicrosoft365MailMessage(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_mail_message",
		Description: "Retrieves messages in the specified user's mailbox.",
		List: &plugin.ListConfig{
			Hydrate: listMicrosoft365MailMessages,
			KeyColumns: plugin.KeyColumnSlice{
				// Key fields
				{Name: "user_id", Require: plugin.Required},
				{Name: "filter", Require: plugin.Optional},

				// Other fields for filtering OData
				{Name: "subject", Require: plugin.Optional},
				{Name: "has_attachments", Require: plugin.Optional, Operators: []string{"<>", "="}},
				{Name: "is_read", Require: plugin.Optional, Operators: []string{"<>", "="}},
				{Name: "is_draft", Require: plugin.Optional, Operators: []string{"<>", "="}},
			},
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ErrorItemNotFound"}),
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getMicrosoft365MailMessage,
			KeyColumns: plugin.AllColumns([]string{"user_id", "id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ErrorItemNotFound"}),
			},
		},
		Columns: mailMessageColumns(),
	}
}

//// LIST FUNCTION

func listMicrosoft365MailMessages(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_mail_message.listMicrosoft365MailMessages", "connection_error", err)
		return nil, err
	}

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

	equalQuals := d.EqualsQuals
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

	userID := d.EqualsQuals["user_id"].GetStringValue()
	result, err := client.UsersById(userID).Messages().Get(ctx, options)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateMessageCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listMicrosoft365MailMessages", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
		message := pageItem.(models.Messageable)

		d.StreamListItem(ctx, &Microsoft365MailMessageInfo{message, userID})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listMicrosoft365MailMessages", "paging_error", err)
		return nil, err
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getMicrosoft365MailMessage(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	userID := d.EqualsQualString("user_id")
	id := d.EqualsQualString("id")
	if userID == "" || id == "" {
		return nil, nil
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_mail_message.listMicrosoft365MailMessages", "connection_error", err)
		return nil, err
	}

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

func buildMailMessageRequestFields(ctx context.Context, queryColumns []string) []string {
	var selectColumns []string

	for _, columnName := range queryColumns {
		if columnName == "title" || columnName == "filter" || columnName == "user_id" || columnName == "_ctx" || columnName == "tenant_id" {
			continue
		}
		selectColumns = append(selectColumns, strcase.ToLowerCamel(columnName))
	}

	return selectColumns
}

func buildQueryFilter(equalQuals plugin.KeyColumnEqualsQualMap) []string {
	filters := []string{}

	filterQuals := map[string]string{
		"subject":         "string",
		"has_attachments": "bool",
		"is_read":         "bool",
		"is_draft":        "bool",
	}

	for qual, qualType := range filterQuals {
		switch qualType {
		case "string":
			if equalQuals[qual] != nil {
				filters = append(filters, fmt.Sprintf("%s eq '%s'", strcase.ToCamel(qual), equalQuals[qual].GetStringValue()))
			}
		case "bool":
			if equalQuals[qual] != nil {
				filters = append(filters, fmt.Sprintf("%s eq %t", strcase.ToCamel(qual), equalQuals[qual].GetBoolValue()))
			}
		}
	}
	return filters
}

func buildBoolNEFilter(quals plugin.KeyColumnQualMap) []string {
	filters := []string{}

	filterQuals := []string{
		"has_attachments",
		"is_read",
		"is_draft",
	}

	for _, qual := range filterQuals {
		if quals[qual] != nil {
			for _, q := range quals[qual].Quals {
				value := q.Value.GetBoolValue()
				if q.Operator == "<>" {
					filters = append(filters, fmt.Sprintf("%s eq %t", strcase.ToCamel(qual), !value))
					break
				}
			}
		}
	}
	return filters
}
