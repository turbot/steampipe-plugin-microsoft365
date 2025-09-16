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
	"github.com/microsoftgraph/msgraph-sdk-go/sites"
)

func siteColumns() []*plugin.Column {
	return commonColumns([]*plugin.Column{
		{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the item.", Transform: transform.FromMethod("GetName")},
		{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier of the site.", Transform: transform.FromMethod("GetId")},
		{Name: "display_name", Type: proto.ColumnType_STRING, Description: "The full title for the site.", Transform: transform.FromMethod("GetDisplayName")},
		{Name: "web_url", Type: proto.ColumnType_STRING, Description: "URL that displays the resource in the browser.", Transform: transform.FromMethod("GetWebUrl")},
		{Name: "description", Type: proto.ColumnType_STRING, Description: "Provide a user-visible description of the site.", Transform: transform.FromMethod("GetDescription")},
		{Name: "created_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time of item creation.", Transform: transform.FromMethod("GetCreatedDateTime")},
		{Name: "last_modified_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time the item was last modified.", Transform: transform.FromMethod("GetLastModifiedDateTime")},

		// SharePoint specific fields
		{Name: "is_personal_site", Type: proto.ColumnType_BOOL, Description: "Indicates whether this is a personal site.", Transform: transform.FromMethod("GetIsPersonalSite")},

		// JSON Fields
		{Name: "created_by", Type: proto.ColumnType_JSON, Description: "Identity of the user, device, or application which created the item.", Transform: transform.FromMethod("SiteCreatedBy")},
		{Name: "last_modified_by", Type: proto.ColumnType_JSON, Description: "Identity of the user, device, and application which last modified the item.", Transform: transform.FromMethod("SiteLastModifiedBy")},
		{Name: "parent_reference", Type: proto.ColumnType_JSON, Description: "Parent information, if the site has a parent.", Transform: transform.FromMethod("SiteParentReference")},
		{Name: "sharepoint_ids", Type: proto.ColumnType_JSON, Description: "Returns identifiers useful for SharePoint REST compatibility.", Transform: transform.FromMethod("SiteSharepointIds")},
		{Name: "site_collection", Type: proto.ColumnType_JSON, Description: "Provides details about the site's site collection.", Transform: transform.FromMethod("SiteSiteCollection")},
		{Name: "root", Type: proto.ColumnType_JSON, Description: "If present, represents the root of the site.", Transform: transform.FromMethod("SiteRoot")},

		// Additional SharePoint-specific fields
		{Name: "analytics", Type: proto.ColumnType_JSON, Description: "Analytics about the view activities that took place in this site.", Transform: transform.FromMethod("SiteAnalytics")},
		{Name: "columns", Type: proto.ColumnType_JSON, Description: "The collection of column definitions reusable across lists under this site.", Transform: transform.FromMethod("SiteColumns")},
		{Name: "content_types", Type: proto.ColumnType_JSON, Description: "The collection of content types defined for this site.", Transform: transform.FromMethod("SiteContentTypes")},
		{Name: "drives", Type: proto.ColumnType_JSON, Description: "The collection of drives (document libraries) under this site.", Transform: transform.FromMethod("SiteDrives")},
		{Name: "lists", Type: proto.ColumnType_JSON, Description: "The collection of lists under this site.", Transform: transform.FromMethod("SiteLists")},
		{Name: "operations", Type: proto.ColumnType_JSON, Description: "The collection of long-running operations on the site.", Transform: transform.FromMethod("SiteOperations")},
		{Name: "pages", Type: proto.ColumnType_JSON, Description: "The collection of pages in the SitePages library in this site.", Transform: transform.FromMethod("SitePages")},
		{Name: "permissions", Type: proto.ColumnType_JSON, Description: "The permissions associated with the site.", Transform: transform.FromMethod("SitePermissions")},
		{Name: "sites", Type: proto.ColumnType_JSON, Description: "The collection of the sub-sites under this site.", Transform: transform.FromMethod("SiteSites")},

		// Standard columns
		{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromMethod("GetDisplayName")},
		{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Odata query to search for resources."},
	})
}

//// TABLE DEFINITION

func tableMicrosoft365Site(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_site",
		Description: "SharePoint sites in Microsoft 365.",
		List: &plugin.ListConfig{
			Hydrate: listMicrosoft365Sites,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "filter", Require: plugin.Optional},
				{Name: "id", Require: plugin.Optional},
				{Name: "display_name", Require: plugin.Optional},
				{Name: "is_personal_site", Require: plugin.Optional},
			},
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"itemNotFound"}),
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getMicrosoft365Site,
			KeyColumns: plugin.AllColumns([]string{"id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"itemNotFound"}),
			},
		},
		Columns: siteColumns(),
	}
}

//// LIST FUNCTION

func listMicrosoft365Sites(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_site.listMicrosoft365Sites", "connection_error", err)
		return nil, err
	}

	input := &sites.GetAllSitesRequestBuilderGetQueryParameters{}

	// Minimum value is 1 (this function isn't run if "limit 0" is specified)
	// Maximum value is unknown (tested up to 9999)
	pageSize := int64(9999)
	limit := d.QueryContext.Limit
	if limit != nil && *limit < pageSize {
		pageSize = *limit
	}
	input.Top = Int32(int32(pageSize))

	var queryFilter string
	equalQuals := d.EqualsQuals
	filter := buildSiteQueryFilter(equalQuals)

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

	options := &sites.GetAllSitesRequestBuilderGetRequestConfiguration{
		QueryParameters: input,
	}

	result, err := client.Sites().GetAllSites().GetAsGetAllSitesGetResponse(ctx, options)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator[models.Siteable](result, adapter, models.CreateSiteCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listMicrosoft365Sites", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem models.Siteable) bool {
		site := pageItem

		d.StreamListItem(ctx, &Microsoft365SiteInfo{site})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		return d.RowsRemaining(ctx) != 0
	})
	if err != nil {
		logger.Error("listMicrosoft365Sites", "paging_error", err)
		return nil, err
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getMicrosoft365Site(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	siteID := d.EqualsQualString("id")
	if siteID == "" {
		return nil, nil
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_site.getMicrosoft365Site", "connection_error", err)
		return nil, err
	}

	result, err := client.Sites().BySiteId(siteID).Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	return &Microsoft365SiteInfo{result}, nil
}

func buildSiteQueryFilter(equalQuals plugin.KeyColumnEqualsQualMap) []string {
	filters := []string{}

	filterQuals := map[string]string{
		"display_name":     "string",
		"id":               "string",
		"is_personal_site": "bool",
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
