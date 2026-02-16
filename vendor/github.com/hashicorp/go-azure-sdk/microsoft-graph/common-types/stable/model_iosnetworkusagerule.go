package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosNetworkUsageRule struct {
	// If set to true, corresponding managed apps will not be allowed to use cellular data when roaming.
	CellularDataBlockWhenRoaming *bool `json:"cellularDataBlockWhenRoaming,omitempty"`

	// If set to true, corresponding managed apps will not be allowed to use cellular data at any time.
	CellularDataBlocked *bool `json:"cellularDataBlocked,omitempty"`

	// Information about the managed apps that this rule is going to apply to. This collection can contain a maximum of 500
	// elements.
	ManagedApps *[]AppListItem `json:"managedApps,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
