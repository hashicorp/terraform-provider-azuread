package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAutonomousSystem struct {
	// The name of the autonomous system.
	Name *string `json:"name,omitempty"`

	// The autonomous system number, assigned by IANA.
	Number *int64 `json:"number,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The name of the autonomous system organization.
	Organization *string `json:"organization,omitempty"`

	// A displayable value for these autonomous system details.
	Value *string `json:"value,omitempty"`
}
