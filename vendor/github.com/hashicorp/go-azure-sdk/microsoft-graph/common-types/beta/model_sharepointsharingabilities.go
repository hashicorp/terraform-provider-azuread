package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharePointSharingAbilities struct {
	// The anyone link abilities.
	AnyoneLinkAbilities *LinkScopeAbilities `json:"anyoneLinkAbilities,omitempty"`

	// The direct sharing abilities.
	DirectSharingAbilities *DirectSharingAbilities `json:"directSharingAbilities,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The organization link abilities.
	OrganizationLinkAbilities *LinkScopeAbilities `json:"organizationLinkAbilities,omitempty"`

	// The specificPeople link abilities.
	SpecificPeopleLinkAbilities *LinkScopeAbilities `json:"specificPeopleLinkAbilities,omitempty"`
}
