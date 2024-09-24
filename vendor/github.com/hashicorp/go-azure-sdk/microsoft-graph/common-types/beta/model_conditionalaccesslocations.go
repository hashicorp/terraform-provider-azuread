package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessLocations struct {
	// Location IDs excluded from scope of policy.
	ExcludeLocations *[]string `json:"excludeLocations,omitempty"`

	// Location IDs in scope of policy unless explicitly excluded, All, or AllTrusted.
	IncludeLocations *[]string `json:"includeLocations,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
