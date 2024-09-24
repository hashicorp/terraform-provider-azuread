package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsDisplayTemplate struct {
	// The text identifier for the display template; for example, contosoTickets. Maximum 16 characters. Only alphanumeric
	// characters allowed.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Defines the priority of a display template. A display template with priority 1 is evaluated before a template with
	// priority 4. Gaps in priority values are supported. Must be positive value.
	Priority *int64 `json:"priority,omitempty"`

	// Specifies additional rules for selecting this display template based on the item schema. Optional.
	Rules *[]ExternalConnectorsPropertyRule `json:"rules,omitempty"`
}
