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
	"github.com/microsoftgraph/msgraph-sdk-go/contacts"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

func organizationContactColumns() []*plugin.Column {
	return commonColumns([]*plugin.Column{
		{Name: "display_name", Type: proto.ColumnType_STRING, Description: "The contact's display name.", Transform: transform.FromMethod("GetDisplayName")},
		{Name: "given_name", Type: proto.ColumnType_STRING, Description: "The contact's given name.", Transform: transform.FromMethod("GetGivenName")},
		{Name: "surname", Type: proto.ColumnType_STRING, Description: "The contact's surname.", Transform: transform.FromMethod("GetSurname")},
		{Name: "id", Type: proto.ColumnType_STRING, Description: "The contact's unique identifier.", Transform: transform.FromMethod("GetId")},

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
		{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Odata query to search for resources."},
	})
}

//// TABLE DEFINITION

func tableMicrosoft365OrganizationContact(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_organization_contact",
		Description: "Retrieve the list of organizational contacts for this organization.",
		List: &plugin.ListConfig{
			Hydrate: listMicrosoft365OrganizationContacts,
			KeyColumns: plugin.KeyColumnSlice{
				// Key fields
				{Name: "display_name", Require: plugin.Optional},
				{Name: "filter", Require: plugin.Optional},
				{Name: "given_name", Require: plugin.Optional},
				{Name: "surname", Require: plugin.Optional},
			},
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"Request_ResourceNotFound"}),
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getMicrosoft365OrganizationContact,
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"Request_ResourceNotFound"}),
			},
		},
		Columns: organizationContactColumns(),
	}
}

//// LIST FUNCTION

func listMicrosoft365OrganizationContacts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
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
	equalQuals := d.EqualsQuals
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

	// Minimum value is 1 (this function isn't run if "limit 0" is specified)
	// Maximum value is 999
	pageSize := int64(999)
	limit := d.QueryContext.Limit
	if limit != nil && *limit < pageSize {
		pageSize = *limit
	}
	input.Top = Int32(int32(pageSize))

	options := &contacts.ContactsRequestBuilderGetRequestConfiguration{
		QueryParameters: input,
	}

	result, err := client.Contacts().Get(ctx, options)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator[models.OrgContactable](result, adapter, models.CreateOrgContactCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listMicrosoft365Contacts", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem models.OrgContactable) bool {
		contact := pageItem

		d.StreamListItem(ctx, &Microsoft365OrgContactInfo{contact})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listMicrosoft365Contacts", "paging_error", err)
		return nil, err
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getMicrosoft365OrganizationContact(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	contactID := d.EqualsQualString("id")
	if contactID == "" {
		return nil, nil
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_contact.listMicrosoft365Contacts", "connection_error", err)
		return nil, err
	}

	result, err := client.Contacts().ByOrgContactId(contactID).Get(ctx, nil)
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
