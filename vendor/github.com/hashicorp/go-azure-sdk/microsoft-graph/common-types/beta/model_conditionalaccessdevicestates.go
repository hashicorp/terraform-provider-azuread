package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessDeviceStates struct {
	// States excluded from the scope of the policy. Possible values: Compliant, DomainJoined.
	ExcludeStates *[]string `json:"excludeStates,omitempty"`

	// States in the scope of the policy. All is the only allowed value.
	IncludeStates *[]string `json:"includeStates,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
