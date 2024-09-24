package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutOfBoxExperienceSettings struct {
	DeviceUsageType *WindowsDeviceUsageType `json:"deviceUsageType,omitempty"`

	// Show or hide EULA to user
	HideEULA *bool `json:"hideEULA,omitempty"`

	// If set to true, then the user can't start over with different account, on company sign-in
	HideEscapeLink *bool `json:"hideEscapeLink,omitempty"`

	// Show or hide privacy settings to user
	HidePrivacySettings *bool `json:"hidePrivacySettings,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// If set, then skip the keyboard selection page if Language and Region are set
	SkipKeyboardSelectionPage *bool `json:"skipKeyboardSelectionPage,omitempty"`

	UserType *WindowsUserType `json:"userType,omitempty"`
}
