package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAttributeCollectionOptionConfiguration struct {
	// The label of the option that will be displayed to user, unless overridden.
	Label *string `json:"label,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The value of the option that will be stored.
	Value *string `json:"value,omitempty"`
}
