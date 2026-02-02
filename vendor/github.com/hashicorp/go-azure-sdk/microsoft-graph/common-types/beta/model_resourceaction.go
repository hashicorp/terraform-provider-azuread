package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceAction struct {
	// Allowed Actions
	AllowedResourceActions *[]string `json:"allowedResourceActions,omitempty"`

	// Not Allowed Actions.
	NotAllowedResourceActions *[]string `json:"notAllowedResourceActions,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
