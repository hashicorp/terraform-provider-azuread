package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutOfBoxExperienceSetting struct {
	DeviceUsageType *WindowsDeviceUsageType `json:"deviceUsageType,omitempty"`

	// When TRUE, the link that allows user to start over with a different account on company sign-in is hidden. When false,
	// the link that allows user to start over with a different account on company sign-in is available. Default value is
	// FALSE.
	EscapeLinkHidden *bool `json:"escapeLinkHidden,omitempty"`

	// When TRUE, EULA is hidden to the end user during OOBE. When FALSE, EULA is shown to the end user during OOBE. Default
	// value is FALSE.
	EulaHidden *bool `json:"eulaHidden,omitempty"`

	// When TRUE, the keyboard selection page is hidden to the end user during OOBE if Language and Region are set. When
	// FALSE, the keyboard selection page is skipped during OOBE.
	KeyboardSelectionPageSkipped *bool `json:"keyboardSelectionPageSkipped,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// When TRUE, privacy settings is hidden to the end user during OOBE. When FALSE, privacy settings is shown to the end
	// user during OOBE. Default value is FALSE.
	PrivacySettingsHidden *bool `json:"privacySettingsHidden,omitempty"`

	UserType *WindowsUserType `json:"userType,omitempty"`
}
