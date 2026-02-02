package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerFilter struct {
	// The identifiers of containers, such as organizational units, that are in scope for a synchronization rule. For Active
	// Directory organizational units, use the distinguished names. An empty list means no container filtering is
	// configured.
	IncludedContainers *[]string `json:"includedContainers,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
