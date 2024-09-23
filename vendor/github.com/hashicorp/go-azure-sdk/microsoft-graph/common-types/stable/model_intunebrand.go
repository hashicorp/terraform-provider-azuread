package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntuneBrand struct {
	// Email address of the person/organization responsible for IT support.
	ContactITEmailAddress nullable.Type[string] `json:"contactITEmailAddress,omitempty"`

	// Name of the person/organization responsible for IT support.
	ContactITName nullable.Type[string] `json:"contactITName,omitempty"`

	// Text comments regarding the person/organization responsible for IT support.
	ContactITNotes nullable.Type[string] `json:"contactITNotes,omitempty"`

	// Phone number of the person/organization responsible for IT support.
	ContactITPhoneNumber nullable.Type[string] `json:"contactITPhoneNumber,omitempty"`

	// Logo image displayed in Company Portal apps which have a dark background behind the logo.
	DarkBackgroundLogo *MimeContent `json:"darkBackgroundLogo,omitempty"`

	// Company/organization name that is displayed to end users.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Logo image displayed in Company Portal apps which have a light background behind the logo.
	LightBackgroundLogo *MimeContent `json:"lightBackgroundLogo,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Display name of the company/organization’s IT helpdesk site.
	OnlineSupportSiteName nullable.Type[string] `json:"onlineSupportSiteName,omitempty"`

	// URL to the company/organization’s IT helpdesk site.
	OnlineSupportSiteUrl nullable.Type[string] `json:"onlineSupportSiteUrl,omitempty"`

	// URL to the company/organization’s privacy policy.
	PrivacyUrl nullable.Type[string] `json:"privacyUrl,omitempty"`

	// Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
	ShowDisplayNameNextToLogo *bool `json:"showDisplayNameNextToLogo,omitempty"`

	// Boolean that represents whether the administrator-supplied logo images are shown or not shown.
	ShowLogo *bool `json:"showLogo,omitempty"`

	// Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
	ShowNameNextToLogo *bool `json:"showNameNextToLogo,omitempty"`

	// Primary theme color used in the Company Portal applications and web portal.
	ThemeColor *RgbColor `json:"themeColor,omitempty"`
}
