package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessRemoteNetworkConnectivityConfiguration struct {
	// List of connectivity configurations for deviceLink objects.
	Links *[]NetworkaccessConnectivityConfigurationLink `json:"links,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Unique identifier or a specific reference assigned to a branchSite. Key.
	RemoteNetworkId *string `json:"remoteNetworkId,omitempty"`

	// Display name assigned to a branchSite.
	RemoteNetworkName *string `json:"remoteNetworkName,omitempty"`
}
