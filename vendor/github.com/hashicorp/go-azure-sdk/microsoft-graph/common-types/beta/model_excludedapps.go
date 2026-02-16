package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExcludedApps struct {
	// The value for if MS Office Access should be excluded or not.
	Access *bool `json:"access,omitempty"`

	// The value for if Microsoft Search as default should be excluded or not.
	Bing *bool `json:"bing,omitempty"`

	// The value for if MS Office Excel should be excluded or not.
	Excel *bool `json:"excel,omitempty"`

	// The value for if MS Office OneDrive for Business - Groove should be excluded or not.
	Groove *bool `json:"groove,omitempty"`

	// The value for if MS Office InfoPath should be excluded or not.
	InfoPath *bool `json:"infoPath,omitempty"`

	// The value for if MS Office Skype for Business - Lync should be excluded or not.
	Lync *bool `json:"lync,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The value for if MS Office OneDrive should be excluded or not.
	OneDrive *bool `json:"oneDrive,omitempty"`

	// The value for if MS Office OneNote should be excluded or not.
	OneNote *bool `json:"oneNote,omitempty"`

	// The value for if MS Office Outlook should be excluded or not.
	Outlook *bool `json:"outlook,omitempty"`

	// The value for if MS Office PowerPoint should be excluded or not.
	PowerPoint *bool `json:"powerPoint,omitempty"`

	// The value for if MS Office Publisher should be excluded or not.
	Publisher *bool `json:"publisher,omitempty"`

	// The value for if MS Office SharePointDesigner should be excluded or not.
	SharePointDesigner *bool `json:"sharePointDesigner,omitempty"`

	// The value for if MS Office Teams should be excluded or not.
	Teams *bool `json:"teams,omitempty"`

	// The value for if MS Office Visio should be excluded or not.
	Visio *bool `json:"visio,omitempty"`

	// The value for if MS Office Word should be excluded or not.
	Word *bool `json:"word,omitempty"`
}
