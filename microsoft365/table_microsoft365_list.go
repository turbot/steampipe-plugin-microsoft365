package microsoft365

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

func listColumns() []*plugin.Column {
	return commonColumns([]*plugin.Column{
		{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the list.", Transform: transform.FromMethod("GetName")},
		{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier of the list.", Transform: transform.FromMethod("GetId")},
		{Name: "display_name", Type: proto.ColumnType_STRING, Description: "The displayable title of the list.", Transform: transform.FromMethod("GetDisplayName")},
		{Name: "description", Type: proto.ColumnType_STRING, Description: "Provides a user-visible description of the list.", Transform: transform.FromMethod("GetDescription")},
		{Name: "created_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time of item creation.", Transform: transform.FromMethod("GetCreatedDateTime")},
		{Name: "last_modified_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time the item was last modified.", Transform: transform.FromMethod("GetLastModifiedDateTime")},
		{Name: "web_url", Type: proto.ColumnType_STRING, Description: "URL that displays the resource in the browser.", Transform: transform.FromMethod("GetWebUrl")},

		// JSON Fields
		{Name: "created_by", Type: proto.ColumnType_JSON, Description: "Identity of the user, device, or application which created the item.", Transform: transform.FromMethod("ListCreatedBy")},
		{Name: "last_modified_by", Type: proto.ColumnType_JSON, Description: "Identity of the user, device, and application which last modified the item.", Transform: transform.FromMethod("ListLastModifiedBy")},
		{Name: "parent_reference", Type: proto.ColumnType_JSON, Description: "Parent information, if the list has a parent.", Transform: transform.FromMethod("ListParentReference")},
		{Name: "sharepoint_ids", Type: proto.ColumnType_JSON, Description: "Returns identifiers useful for SharePoint REST compatibility.", Transform: transform.FromMethod("ListSharepointIds")},

		// List-specific fields
		{Name: "columns", Type: proto.ColumnType_JSON, Description: "The collection of field definitions for this list.", Transform: transform.FromMethod("ListColumns")},
		{Name: "content_types", Type: proto.ColumnType_JSON, Description: "The collection of content types present in this list.", Transform: transform.FromMethod("ListContentTypes")},
		{Name: "drive", Type: proto.ColumnType_JSON, Description: "Only present on document libraries. Allows access to the list as a drive resource with driveItems.", Transform: transform.FromMethod("ListDrive")},
		{Name: "items", Type: proto.ColumnType_JSON, Description: "All items contained in the list.", Transform: transform.FromMethod("ListItems")},
		{Name: "list_info", Type: proto.ColumnType_JSON, Description: "Provides additional details about the list.", Transform: transform.FromMethod("ListInfo")},
		{Name: "operations", Type: proto.ColumnType_JSON, Description: "The collection of long-running operations on the list.", Transform: transform.FromMethod("ListOperations")},
		{Name: "subscriptions", Type: proto.ColumnType_JSON, Description: "The set of subscriptions on the list.", Transform: transform.FromMethod("ListSubscriptions")},
		{Name: "system", Type: proto.ColumnType_JSON, Description: "If present, indicates that this is a system-managed list.", Transform: transform.FromMethod("ListSystem")},

		// Standard columns
		{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.FromMethod("GetDisplayName")},
		{Name: "site_id", Type: proto.ColumnType_STRING, Description: "The ID of the site this list belongs to."},
		{Name: "filter", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter"), Description: "Odata query to search for resources."},
	})
}

//// TABLE DEFINITION

func tableMicrosoft365List(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "microsoft365_list",
		Description: "SharePoint lists in Microsoft 365.",
		List: &plugin.ListConfig{
			Hydrate:       listMicrosoft365Lists,
			ParentHydrate: listMicrosoft365Sites,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"itemNotFound"}),
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getMicrosoft365List,
			KeyColumns: plugin.AllColumns([]string{"site_id", "id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"itemNotFound"}),
			},
		},
		Columns: listColumns(),
	}
}

//// LIST FUNCTION

func listMicrosoft365Lists(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	siteData := h.Item.(*Microsoft365SiteInfo)

	var siteID string
	if siteData != nil {
		siteID = *siteData.GetId()
	}
	logger := plugin.Logger(ctx)

	// Create client
	client, adapter, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_list.listMicrosoft365Lists", "connection_error", err)
		return nil, err
	}

	result, err := client.Sites().BySiteId(siteID).Lists().Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	pageIterator, err := msgraphcore.NewPageIterator[models.Listable](result, adapter, models.CreateListCollectionResponseFromDiscriminatorValue)
	if err != nil {
		logger.Error("listMicrosoft365Lists", "create_iterator_instance_error", err)
		return nil, err
	}

	err = pageIterator.Iterate(ctx, func(pageItem models.Listable) bool {
		d.StreamListItem(ctx, Microsoft365ListInfo{pageItem, siteID})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return false
		}

		return true
	})
	if err != nil {
		logger.Error("listMicrosoft365Lists", "paging_error", err)
		return nil, err
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getMicrosoft365List(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	siteID := d.EqualsQualString("site_id")
	id := d.EqualsQualString("id")
	if siteID == "" || id == "" {
		return nil, nil
	}

	// Create client
	client, _, err := GetGraphClient(ctx, d)
	if err != nil {
		logger.Error("microsoft365_list.getMicrosoft365List", "connection_error", err)
		return nil, err
	}

	result, err := client.Sites().BySiteId(siteID).Lists().ByListId(id).Get(ctx, nil)
	if err != nil {
		errObj := getErrorObject(err)
		return nil, errObj
	}

	return Microsoft365ListInfo{result, siteID}, nil
}
