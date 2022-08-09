package office365

import (
	"context"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/messages"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION

func tableOffice365MailMessage() *plugin.Table {
	return &plugin.Table{
		Name:        "office365_mail_message",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listOffice365MailMessages,
			KeyColumns: plugin.KeyColumnSlice{
				// Key fields
				{Name: "user_identifier", Require: plugin.Required},
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

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetId")},
			{Name: "subject", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetSubject")},
			{Name: "body_preview", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetBodyPreview")},
			{Name: "is_read", Type: proto.ColumnType_BOOL, Description: "", Transform: transform.FromMethod("GetIsRead")},
			{Name: "has_attachments", Type: proto.ColumnType_BOOL, Description: "", Transform: transform.FromMethod("GetHasAttachments")},
			{Name: "is_draft", Type: proto.ColumnType_BOOL, Description: "", Transform: transform.FromMethod("GetIsDraft")},
			{Name: "received_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "", Transform: transform.FromMethod("GetReceivedDateTime")},
			{Name: "sent_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "", Transform: transform.FromMethod("GetSentDateTime")},

			// Other columns
			{Name: "importance", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("MessageImportance")},
			{Name: "inference_classification", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("MessageInferenceClassification")},
			{Name: "change_key", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetChangeKey")},
			{Name: "conversation_id", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetConversationId")},
			{Name: "internet_message_id", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetInternetMessageId")},
			{Name: "is_read_receipt_requested", Type: proto.ColumnType_BOOL, Description: "", Transform: transform.FromMethod("GetIsReadReceiptRequested")},
			{Name: "is_delivery_receipt_requested", Type: proto.ColumnType_BOOL, Description: "", Transform: transform.FromMethod("GetIsDeliveryReceiptRequested")},
			{Name: "parent_folder_id", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetParentFolderId")},

			// Other fields
			{Name: "created_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "", Transform: transform.FromMethod("GetCreatedDateTime")},
			{Name: "last_modified_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "", Transform: transform.FromMethod("GetLastModifiedDateTime")},
			{Name: "web_link", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetWebLink")},

			// JSON fields
			{Name: "sender", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("MessageSender")},
			{Name: "from", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("MessageFrom")},
			{Name: "body", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("MessageBody")},
			{Name: "categories", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("GetCategories")},
			{Name: "reply_to", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("MessageReplyTo")},
			{Name: "to_recipients", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("MessageToRecipients")},
			{Name: "attachments", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("MessageAttachments")},
			{Name: "bcc_recipients", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("MessageBccRecipients")},
			{Name: "cc_recipients", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("MessageCcRecipients")},

			// Standard columns
			{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromMethod("GetSubject")},
			{Name: "user_identifier", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromQual("user_identifier")},
			{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Odata query to search for resources."},
			// {Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Hydrate: plugin.HydrateFunc(getTenant).WithCache(), Transform: transform.FromValue()},
		},
	}
}

//// LIST FUNCTION

func listOffice365MailMessages(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("error creating client: %v", err)
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
		input.Filter = &queryFilter
	} else if len(filter) > 0 {
		joinStr := strings.Join(filter, " and ")
		input.Filter = &joinStr
	}

	options := &messages.MessagesRequestBuilderGetRequestConfiguration{
		QueryParameters: input,
	}

	userIdentifier := d.KeyColumnQuals["user_identifier"].GetStringValue()
	result, err := client.UsersById(userIdentifier).Messages().GetWithRequestConfigurationAndResponseHandler(options, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateMessageCollectionResponseFromDiscriminatorValue)

	err = pageIterator.Iterate(func(pageItem interface{}) bool {
		message := pageItem.(models.Messageable)

		d.StreamListItem(ctx, &Office365MailMessageInfo{message})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return false
		}

		return true
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func buildMailMessageRequestFields(ctx context.Context, queryColumns []string) []string {
	var selectColumns []string

	for _, columnName := range queryColumns {
		if columnName == "title" || columnName == "filter" || columnName == "user_identifier" || columnName == "_ctx" {
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
