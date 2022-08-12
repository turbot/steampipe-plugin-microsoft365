package office365

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users/item/contacts"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION

func tableOffice365Contact(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "office365_contact",
		Description: "",
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

		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetId")},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetDisplayName")},
			{Name: "given_name", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetGivenName")},
			{Name: "middle_name", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetMiddleName")},
			{Name: "surname", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetSurname")},
			{Name: "mobile_phone", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetMobilePhone")},

			// Other columns
			{Name: "assistant_name", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetAssistantName")},
			{Name: "birthday", Type: proto.ColumnType_TIMESTAMP, Description: "", Transform: transform.FromMethod("GetBirthday")},
			{Name: "business_home_page", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetBusinessHomePage")},
			{Name: "company_name", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetCompanyName")},
			{Name: "department", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetDepartment")},
			{Name: "file_as", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetFileAs")},
			{Name: "generation", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetGeneration")},
			{Name: "initials", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetInitials")},
			{Name: "job_title", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetJobTitle")},
			{Name: "manager", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetManager")},
			{Name: "nick_name", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetNickName")},
			{Name: "parent_folder_id", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetParentFolderId")},
			{Name: "office_location", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetOfficeLocation")},
			{Name: "personal_notes", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetPersonalNotes")},
			{Name: "profession", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetProfession")},
			{Name: "spouse_name", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetSpouseName")},
			{Name: "yomi_company_name", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetYomiCompanyName")},
			{Name: "yomi_given_name", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetYomiGivenName")},
			{Name: "yomi_surname", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetYomiSurname")},

			// JSON columns
			{Name: "business_address", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("ContactBusinessAddress")},
			{Name: "business_phones", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetBusinessPhones")},
			{Name: "children", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetChildren")},
			{Name: "email_addresses", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("ContactEmailAddresses")},
			{Name: "home_address", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("ContactHomeAddress")},
			{Name: "home_phones", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetHomePhones")},
			{Name: "im_addresses", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("GetImAddresses")},
			{Name: "other_address", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("ContactOtherAddress")},

			// Standard columns
			{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromMethod("GetDisplayName")},
			{Name: "user_identifier", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromQual("user_identifier")},
			{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Hydrate: plugin.HydrateFunc(getTenant).WithCache(), Transform: transform.FromValue()},
		},
	}
}

//// LIST FUNCTION

func listOffice365Contacts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("error creating client: %v", err)
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
