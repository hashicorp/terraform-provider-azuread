package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessUsers struct {
	// Group IDs excluded from scope of policy.
	ExcludeGroups *[]string `json:"excludeGroups,omitempty"`

	// Internal guests or external users excluded from the policy scope. Optionally populated.
	ExcludeGuestsOrExternalUsers *ConditionalAccessGuestsOrExternalUsers `json:"excludeGuestsOrExternalUsers"`

	// Role IDs excluded from scope of policy.
	ExcludeRoles *[]string `json:"excludeRoles,omitempty"`

	// User IDs excluded from scope of policy and/or GuestsOrExternalUsers.
	ExcludeUsers *[]string `json:"excludeUsers,omitempty"`

	// Group IDs in scope of policy unless explicitly excluded.
	IncludeGroups *[]string `json:"includeGroups,omitempty"`

	// Internal guests or external users included in the policy scope. Optionally populated.
	IncludeGuestsOrExternalUsers *ConditionalAccessGuestsOrExternalUsers `json:"includeGuestsOrExternalUsers"`

	// Role IDs in scope of policy unless explicitly excluded.
	IncludeRoles *[]string `json:"includeRoles,omitempty"`

	// User IDs in scope of policy unless explicitly excluded, None, All, or GuestsOrExternalUsers.
	IncludeUsers *[]string `json:"includeUsers,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
