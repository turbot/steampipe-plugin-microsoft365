package microsoft365

import (
	"context"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/contacts"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION

func tableMicrosoft365MyContact(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_my_contact",
		Description: "Contacts owned by the specified user.",
		List: &plugin.ListConfig{
			Hydrate: listMicrosoft365MyContacts,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getMicrosoft365MyContact,
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Columns: contactColumns(),
	}
}

//// LIST FUNCTION

func listMicrosoft365MyContacts(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_my_contact.listMicrosoft365MyContacts", "connection_error", err)
		return nil, err
	}

	getUserIdentifierCached := plugin.HydrateFunc(getUserIdentifier).WithCache()
	userID, err := getUserIdentifierCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	userIdentifier := userID.(string)

	// List operations
	input := &contacts.ContactsRequestBuilderGetQueryParameters{}

	// Minimum value is 1 (this function isn't run if "limit 0" is specified)
	// Maximum value is unknown (tested up to 9999)
	pageSize := int64(9999)
	limit := d.QueryContext.Limit
	if limit != nil && *limit < pageSize {
		pageSize = *limit
	}

	pageSize32 := int32(pageSize)
	input.Top = &pageSize32

	options := &contacts.ContactsRequestBuilderGetRequestConfiguration{
		QueryParameters: input,
	}

	result, err := client.UsersById(userIdentifier).Contacts().Get(ctx, options)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateContactCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listMicrosoft365MyContacts", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
		contact := pageItem.(models.Contactable)

		d.StreamListItem(ctx, &Microsoft365ContactInfo{contact, userIdentifier})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.QueryStatus.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listMicrosoft365MyContacts", "paging_error", err)
		return nil, err
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getMicrosoft365MyContact(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	contactID := d.KeyColumnQualString("id")
	if contactID == "" {
		return nil, nil
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_my_contact.getMicrosoft365MyContact", "connection_error", err)
		return nil, err
	}

	getUserIdentifierCached := plugin.HydrateFunc(getUserIdentifier).WithCache()
	userID, err := getUserIdentifierCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	userIdentifier := userID.(string)

	result, err := client.UsersById(userIdentifier).ContactsById(contactID).Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	return &Microsoft365ContactInfo{result, userIdentifier}, nil
}
