package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SystemCredentialPreferences struct {
	// Users and groups excluded from the preferred authentication method experience of the system.
	ExcludeTargets *[]ExcludeTarget `json:"excludeTargets,omitempty"`

	// Users and groups included in the preferred authentication method experience of the system.
	IncludeTargets *[]IncludeTarget `json:"includeTargets,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	State *AdvancedConfigState `json:"state,omitempty"`
}
