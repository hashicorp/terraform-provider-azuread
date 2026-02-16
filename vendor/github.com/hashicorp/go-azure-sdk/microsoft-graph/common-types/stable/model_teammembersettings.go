package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamMemberSettings struct {
	// If set to true, members can add and remove apps.
	AllowAddRemoveApps nullable.Type[bool] `json:"allowAddRemoveApps,omitempty"`

	// If set to true, members can add and update private channels.
	AllowCreatePrivateChannels nullable.Type[bool] `json:"allowCreatePrivateChannels,omitempty"`

	// If set to true, members can add and update channels.
	AllowCreateUpdateChannels nullable.Type[bool] `json:"allowCreateUpdateChannels,omitempty"`

	// If set to true, members can add, update, and remove connectors.
	AllowCreateUpdateRemoveConnectors nullable.Type[bool] `json:"allowCreateUpdateRemoveConnectors,omitempty"`

	// If set to true, members can add, update, and remove tabs.
	AllowCreateUpdateRemoveTabs nullable.Type[bool] `json:"allowCreateUpdateRemoveTabs,omitempty"`

	// If set to true, members can delete channels.
	AllowDeleteChannels nullable.Type[bool] `json:"allowDeleteChannels,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
