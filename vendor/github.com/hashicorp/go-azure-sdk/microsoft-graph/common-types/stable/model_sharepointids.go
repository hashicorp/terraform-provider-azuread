package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharepointIds struct {
	// The unique identifier (guid) for the item's list in SharePoint.
	ListId nullable.Type[string] `json:"listId,omitempty"`

	// An integer identifier for the item within the containing list.
	ListItemId nullable.Type[string] `json:"listItemId,omitempty"`

	// The unique identifier (guid) for the item within OneDrive for Business or a SharePoint site.
	ListItemUniqueId nullable.Type[string] `json:"listItemUniqueId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The unique identifier (guid) for the item's site collection (SPSite).
	SiteId nullable.Type[string] `json:"siteId,omitempty"`

	// The SharePoint URL for the site that contains the item.
	SiteUrl nullable.Type[string] `json:"siteUrl,omitempty"`

	// The unique identifier (guid) for the tenancy.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

	// The unique identifier (guid) for the item's site (SPWeb).
	WebId nullable.Type[string] `json:"webId,omitempty"`
}
