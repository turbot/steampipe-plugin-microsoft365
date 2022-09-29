package microsoft365

import (
	"context"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/contacts"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

func orgContactColumns() []*plugin.Column {
	return []*plugin.Column{
		{Name: "id", Type: proto.ColumnType_STRING, Description: "The contact's unique identifier.", Transform: transform.FromMethod("GetId")},
		{Name: "display_name", Type: proto.ColumnType_STRING, Description: "The contact's display name.", Transform: transform.FromMethod("GetDisplayName")},
		{Name: "given_name", Type: proto.ColumnType_STRING, Description: "The contact's given name.", Transform: transform.FromMethod("GetGivenName")},
		{Name: "surname", Type: proto.ColumnType_STRING, Description: "The contact's surname.", Transform: transform.FromMethod("GetSurname")},

		// Other columns
		{Name: "company_name", Type: proto.ColumnType_STRING, Description: "The name of the contact's company.", Transform: transform.FromMethod("GetCompanyName")},
		{Name: "department", Type: proto.ColumnType_STRING, Description: "The contact's department.", Transform: transform.FromMethod("GetDepartment")},
		{Name: "job_title", Type: proto.ColumnType_STRING, Description: "The contact’s job title.", Transform: transform.FromMethod("GetJobTitle")},
		{Name: "mail", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetMail")},
		{Name: "mail_nickname", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromMethod("GetMailNickname")},
		{Name: "on_premises_last_sync_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "", Transform: transform.FromMethod("GetOnPremisesLastSyncDateTime")},
		{Name: "on_premises_sync_enabled", Type: proto.ColumnType_BOOL, Description: "", Transform: transform.FromMethod("GetOnPremisesSyncEnabled")},

		// JSON columns
		{Name: "addresses", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("OrgContactAddresses")},
		{Name: "direct_reports", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("OrgContactDirectReports")},
		{Name: "manager", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("OrgContactManager")},
		{Name: "member_of", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("OrgContactMemberOf")},
		{Name: "on_premises_provisioning_errors", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("OrgContactOnPremisesProvisioningErrors")},
		{Name: "phones", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("OrgContactPhones")},
		{Name: "proxy_addresses", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("GetProxyAddresses")},
		{Name: "transitive_member_of", Type: proto.ColumnType_JSON, Description: "", Transform: transform.FromMethod("OrgContactTransitiveMemberOf")},

		// Standard columns
		{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromMethod("GetDisplayName")},
		{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Hydrate: plugin.HydrateFunc(getTenant).WithCache(), Transform: transform.FromValue()},
		{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Odata query to search for resources."},
	}
}

//// TABLE DEFINITION

func tableMicrosoft365OrgContact(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_org_contact",
		Description: "Retrieve the list of organizational contacts for this organization.",
		List: &plugin.ListConfig{
			Hydrate: listMicrosoft365OrgContacts,
			KeyColumns: plugin.KeyColumnSlice{
				// Key fields
				{Name: "display_name", Require: plugin.Optional},
				{Name: "filter", Require: plugin.Optional},
				{Name: "given_name", Require: plugin.Optional},
				{Name: "surname", Require: plugin.Optional},
			},
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getMicrosoft365OrgContact,
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"ResourceNotFound"}),
			},
		},
		Columns: orgContactColumns(),
	}
}

//// LIST FUNCTION

func listMicrosoft365OrgContacts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_contact.listMicrosoft365Contacts", "connection_error", err)
		return nil, err
	}

	// List operations
	input := &contacts.ContactsRequestBuilderGetQueryParameters{}

	var queryFilter string
	equalQuals := d.KeyColumnQuals
	filter := buildOrgContactQueryFilter(equalQuals)

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

	//userIdentifier := d.KeyColumnQuals["user_identifier"].GetStringValue()
	result, err := client.Contacts().Get(ctx, options)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator(result, adapter, models.CreateOrgContactCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listMicrosoft365Contacts", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem interface{}) bool {
		contact := pageItem.(models.OrgContactable)

		d.StreamListItem(ctx, &Microsoft365OrgContactInfo{contact})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.QueryStatus.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listMicrosoft365Contacts", "paging_error", err)
		return nil, err
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getMicrosoft365OrgContact(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	contactID := d.KeyColumnQualString("id")
	if contactID == "" {
		return nil, nil
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_contact.listMicrosoft365Contacts", "connection_error", err)
		return nil, err
	}

	result, err := client.ContactsById(contactID).Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	return &Microsoft365OrgContactInfo{result}, nil
}

func buildOrgContactQueryFilter(equalQuals plugin.KeyColumnEqualsQualMap) []string {
	filters := []string{}

	filterQuals := map[string]string{
		"display_name": "string",
		"given_name":   "string",
		"surname":      "string",
	}

	for qual, qualType := range filterQuals {
		switch qualType {
		case "string":
			if equalQuals[qual] != nil {
				filters = append(filters, fmt.Sprintf("%s eq '%s'", strcase.ToCamel(qual), equalQuals[qual].GetStringValue()))
			}
		}
	}
	return filters
}
