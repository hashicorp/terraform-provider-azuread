package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LinkScopeAbilities struct {
	// The blockDownload link abilities.
	BlockDownloadLinkAbilities *LinkRoleAbilities `json:"blockDownloadLinkAbilities,omitempty"`

	EditLinkAbilities *LinkRoleAbilities `json:"editLinkAbilities,omitempty"`

	// The manageList link abilities.
	ManageListLinkAbilities *LinkRoleAbilities `json:"manageListLinkAbilities,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	ReadLinkAbilities *LinkRoleAbilities `json:"readLinkAbilities,omitempty"`

	// The review link abilities.
	ReviewLinkAbilities *LinkRoleAbilities `json:"reviewLinkAbilities,omitempty"`

	// The submitOnly link abilities.
	SubmitOnlyLinkAbilities *LinkRoleAbilities `json:"submitOnlyLinkAbilities,omitempty"`
}
