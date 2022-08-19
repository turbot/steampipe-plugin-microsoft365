package office365

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/contacts"
)

func contactColumns() []*plugin.Column {
	return []*plugin.Column{
		{Name: "id", Type: proto.ColumnType_STRING, Description: "The contact's unique identifier.", Transform: transform.FromMethod("GetId")},
		{Name: "display_name", Type: proto.ColumnType_STRING, Description: "The contact's display name.", Transform: transform.FromMethod("GetDisplayName")},
		{Name: "given_name", Type: proto.ColumnType_STRING, Description: "The contact's given name.", Transform: transform.FromMethod("GetGivenName")},
		{Name: "middle_name", Type: proto.ColumnType_STRING, Description: "The contact's middle name.", Transform: transform.FromMethod("GetMiddleName")},
		{Name: "surname", Type: proto.ColumnType_STRING, Description: "The contact's surname.", Transform: transform.FromMethod("GetSurname")},
		{Name: "mobile_phone", Type: proto.ColumnType_STRING, Description: "The contact's mobile phone number.", Transform: transform.FromMethod("GetMobilePhone")},

		// Other columns
		{Name: "assistant_name", Type: proto.ColumnType_STRING, Description: "The name of the contact's assistant.", Transform: transform.FromMethod("GetAssistantName")},
		{Name: "birthday", Type: proto.ColumnType_TIMESTAMP, Description: "The contact's birthday. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time.", Transform: transform.FromMethod("GetBirthday")},
		{Name: "business_home_page", Type: proto.ColumnType_STRING, Description: "The business home page of the contact.", Transform: transform.FromMethod("GetBusinessHomePage")},
		{Name: "company_name", Type: proto.ColumnType_STRING, Description: "The name of the contact's company.", Transform: transform.FromMethod("GetCompanyName")},
		{Name: "department", Type: proto.ColumnType_STRING, Description: "The contact's department.", Transform: transform.FromMethod("GetDepartment")},
		{Name: "file_as", Type: proto.ColumnType_STRING, Description: "The name the contact is filed under.", Transform: transform.FromMethod("GetFileAs")},
		{Name: "generation", Type: proto.ColumnType_STRING, Description: "The contact's generation.", Transform: transform.FromMethod("GetGeneration")},
		{Name: "initials", Type: proto.ColumnType_STRING, Description: "The contact's initials.", Transform: transform.FromMethod("GetInitials")},
		{Name: "job_title", Type: proto.ColumnType_STRING, Description: "The contact’s job title.", Transform: transform.FromMethod("GetJobTitle")},
		{Name: "manager", Type: proto.ColumnType_STRING, Description: "The name of the contact's manager.", Transform: transform.FromMethod("GetManager")},
		{Name: "nick_name", Type: proto.ColumnType_STRING, Description: "The contact's nickname.", Transform: transform.FromMethod("GetNickName")},
		{Name: "parent_folder_id", Type: proto.ColumnType_STRING, Description: "The ID of the contact's parent folder.", Transform: transform.FromMethod("GetParentFolderId")},
		{Name: "office_location", Type: proto.ColumnType_STRING, Description: "The location of the contact's office.", Transform: transform.FromMethod("GetOfficeLocation")},
		{Name: "personal_notes", Type: proto.ColumnType_STRING, Description: "The user's notes about the contact.", Transform: transform.FromMethod("GetPersonalNotes")},
		{Name: "profession", Type: proto.ColumnType_STRING, Description: "The contact's profession.", Transform: transform.FromMethod("GetProfession")},
		{Name: "spouse_name", Type: proto.ColumnType_STRING, Description: "The name of the contact's spouse/partner.", Transform: transform.FromMethod("GetSpouseName")},
		{Name: "yomi_company_name", Type: proto.ColumnType_STRING, Description: "The phonetic Japanese company name of the contact.", Transform: transform.FromMethod("GetYomiCompanyName")},
		{Name: "yomi_given_name", Type: proto.ColumnType_STRING, Description: "The phonetic Japanese given name (first name) of the contact.", Transform: transform.FromMethod("GetYomiGivenName")},
		{Name: "yomi_surname", Type: proto.ColumnType_STRING, Description: "The phonetic Japanese surname (last name) of the contact.", Transform: transform.FromMethod("GetYomiSurname")},

		// JSON columns
		{Name: "business_address", Type: proto.ColumnType_STRING, Description: "The contact's business address.", Transform: transform.FromMethod("ContactBusinessAddress")},
		{Name: "business_phones", Type: proto.ColumnType_STRING, Description: "The contact's business phone numbers.", Transform: transform.FromMethod("GetBusinessPhones")},
		{Name: "children", Type: proto.ColumnType_STRING, Description: "The names of the contact's children.", Transform: transform.FromMethod("GetChildren")},
		{Name: "email_addresses", Type: proto.ColumnType_JSON, Description: "The contact's email addresses.", Transform: transform.FromMethod("ContactEmailAddresses")},
		{Name: "home_address", Type: proto.ColumnType_STRING, Description: "The contact's home address.", Transform: transform.FromMethod("ContactHomeAddress")},
		{Name: "home_phones", Type: proto.ColumnType_STRING, Description: "The contact's home phone numbers.", Transform: transform.FromMethod("GetHomePhones")},
		{Name: "im_addresses", Type: proto.ColumnType_JSON, Description: "The contact's instant messaging (IM) addresses.", Transform: transform.FromMethod("GetImAddresses")},
		{Name: "other_address", Type: proto.ColumnType_STRING, Description: "Other addresses for the contact.", Transform: transform.FromMethod("ContactOtherAddress")},

		// Standard columns
		{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromMethod("GetDisplayName")},
		{Name: "user_identifier", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromQual("user_identifier")},
		{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Hydrate: plugin.HydrateFunc(getTenant).WithCache(), Transform: transform.FromValue()},
	}
}

//// TABLE DEFINITION

func tableOffice365Contact(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "office365_contact",
		Description: "Contacts owned by the specified user.",
		List: &plugin.ListConfig{
			Hydrate: listOffice365Contacts,
			KeyColumns: plugin.KeyColumnSlice{
				// Key fields
				{Name: "user_identifier", Require: plugin.Required},
			},
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Columns: contactColumns(),
	}
}

//// LIST FUNCTION

func listOffice365Contacts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("office365_contact.listOffice365Contacts", "connection_error", err)
		return nil, err
	}

	// List operations
	input := &contacts.ContactsRequestBuilderGetQueryParameters{}

	// Restrict the limit value to be passed in the query parameter which is not between 1 and 999, otherwise API will throw an error as follow
	limit := d.QueryContext.Limit
	if limit != nil {
		if *limit > 0 && *limit <= 999 {
			l := int32(*limit)
			input.Top = &l
		}
	}

	options := &contacts.ContactsRequestBuilderGetRequestConfiguration{
		QueryParameters: input,
	}

	userIdentifier := d.KeyColumnQuals["user_identifier"].GetStringValue()
	result, err := client.UsersById(userIdentifier).Contacts().GetWithRequestConfigurationAndResponseHandler(options, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateContactCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listOffice365Contacts", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(func(pageItem interface{}) bool {
		contact := pageItem.(models.Contactable)

		d.StreamListItem(ctx, &Office365ContactInfo{contact})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.QueryStatus.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listOffice365Contacts", "paging_error", err)
		return nil, err
	}

	return nil, nil
}
