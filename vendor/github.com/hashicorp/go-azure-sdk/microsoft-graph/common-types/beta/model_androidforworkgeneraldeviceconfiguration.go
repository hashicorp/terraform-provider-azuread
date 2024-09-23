package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceConfiguration = AndroidForWorkGeneralDeviceConfiguration{}

type AndroidForWorkGeneralDeviceConfiguration struct {
	// Determine domains allow-list for accounts that can be added to work profile.
	AllowedGoogleAccountDomains *[]string `json:"allowedGoogleAccountDomains,omitempty"`

	// Prevent using unified password for unlocking device and work profile.
	BlockUnifiedPasswordForWorkProfile *bool `json:"blockUnifiedPasswordForWorkProfile,omitempty"`

	// Indicates whether or not to block face unlock.
	PasswordBlockFaceUnlock *bool `json:"passwordBlockFaceUnlock,omitempty"`

	// Indicates whether or not to block fingerprint unlock.
	PasswordBlockFingerprintUnlock *bool `json:"passwordBlockFingerprintUnlock,omitempty"`

	// Indicates whether or not to block iris unlock.
	PasswordBlockIrisUnlock *bool `json:"passwordBlockIrisUnlock,omitempty"`

	// Indicates whether or not to block Smart Lock and other trust agents.
	PasswordBlockTrustAgents *bool `json:"passwordBlockTrustAgents,omitempty"`

	// Number of days before the password expires. Valid values 1 to 365
	PasswordExpirationDays nullable.Type[int64] `json:"passwordExpirationDays,omitempty"`

	// Minimum length of passwords. Valid values 4 to 16
	PasswordMinimumLength nullable.Type[int64] `json:"passwordMinimumLength,omitempty"`

	// Minutes of inactivity before the screen times out.
	PasswordMinutesOfInactivityBeforeScreenTimeout nullable.Type[int64] `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`

	// Number of previous passwords to block. Valid values 0 to 24
	PasswordPreviousPasswordBlockCount nullable.Type[int64] `json:"passwordPreviousPasswordBlockCount,omitempty"`

	// Android For Work required password type.
	PasswordRequiredType *AndroidForWorkRequiredPasswordType `json:"passwordRequiredType,omitempty"`

	// Number of sign in failures allowed before factory reset. Valid values 1 to 16
	PasswordSignInFailureCountBeforeFactoryReset nullable.Type[int64] `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`

	// The password complexity types that can be set on Android. One of: NONE, LOW, MEDIUM, HIGH. This is an API targeted to
	// Android 11+.
	RequiredPasswordComplexity *AndroidRequiredPasswordComplexity `json:"requiredPasswordComplexity,omitempty"`

	// Require the Android Verify apps feature is turned on.
	SecurityRequireVerifyApps *bool `json:"securityRequireVerifyApps,omitempty"`

	// Enable lockdown mode for always-on VPN.
	VpnAlwaysOnPackageIdentifier nullable.Type[string] `json:"vpnAlwaysOnPackageIdentifier,omitempty"`

	// Enable lockdown mode for always-on VPN.
	VpnEnableAlwaysOnLockdownMode *bool `json:"vpnEnableAlwaysOnLockdownMode,omitempty"`

	// An enum representing possible values for account use in work profile.
	WorkProfileAccountUse *AndroidWorkProfileAccountUse `json:"workProfileAccountUse,omitempty"`

	// Allow widgets from work profile apps.
	WorkProfileAllowWidgets *bool `json:"workProfileAllowWidgets,omitempty"`

	// Block users from adding/removing accounts in work profile.
	WorkProfileBlockAddingAccounts *bool `json:"workProfileBlockAddingAccounts,omitempty"`

	// Block work profile camera.
	WorkProfileBlockCamera *bool `json:"workProfileBlockCamera,omitempty"`

	// Block display work profile caller ID in personal profile.
	WorkProfileBlockCrossProfileCallerId *bool `json:"workProfileBlockCrossProfileCallerId,omitempty"`

	// Block work profile contacts availability in personal profile.
	WorkProfileBlockCrossProfileContactsSearch *bool `json:"workProfileBlockCrossProfileContactsSearch,omitempty"`

	// Boolean that indicates if the setting disallow cross profile copy/paste is enabled.
	WorkProfileBlockCrossProfileCopyPaste *bool `json:"workProfileBlockCrossProfileCopyPaste,omitempty"`

	// Indicates whether or not to block notifications while device locked.
	WorkProfileBlockNotificationsWhileDeviceLocked *bool `json:"workProfileBlockNotificationsWhileDeviceLocked,omitempty"`

	// Prevent app installations from unknown sources in the personal profile.
	WorkProfileBlockPersonalAppInstallsFromUnknownSources *bool `json:"workProfileBlockPersonalAppInstallsFromUnknownSources,omitempty"`

	// Block screen capture in work profile.
	WorkProfileBlockScreenCapture *bool `json:"workProfileBlockScreenCapture,omitempty"`

	// Allow bluetooth devices to access enterprise contacts.
	WorkProfileBluetoothEnableContactSharing *bool `json:"workProfileBluetoothEnableContactSharing,omitempty"`

	// Android For Work cross profile data sharing type.
	WorkProfileDataSharingType *AndroidForWorkCrossProfileDataSharingType `json:"workProfileDataSharingType,omitempty"`

	// Android For Work default app permission policy type.
	WorkProfileDefaultAppPermissionPolicy *AndroidForWorkDefaultAppPermissionPolicyType `json:"workProfileDefaultAppPermissionPolicy,omitempty"`

	// Indicates whether or not to block face unlock for work profile.
	WorkProfilePasswordBlockFaceUnlock *bool `json:"workProfilePasswordBlockFaceUnlock,omitempty"`

	// Indicates whether or not to block fingerprint unlock for work profile.
	WorkProfilePasswordBlockFingerprintUnlock *bool `json:"workProfilePasswordBlockFingerprintUnlock,omitempty"`

	// Indicates whether or not to block iris unlock for work profile.
	WorkProfilePasswordBlockIrisUnlock *bool `json:"workProfilePasswordBlockIrisUnlock,omitempty"`

	// Indicates whether or not to block Smart Lock and other trust agents for work profile.
	WorkProfilePasswordBlockTrustAgents *bool `json:"workProfilePasswordBlockTrustAgents,omitempty"`

	// Number of days before the work profile password expires. Valid values 1 to 365
	WorkProfilePasswordExpirationDays nullable.Type[int64] `json:"workProfilePasswordExpirationDays,omitempty"`

	// Minimum # of letter characters required in work profile password. Valid values 1 to 10
	WorkProfilePasswordMinLetterCharacters nullable.Type[int64] `json:"workProfilePasswordMinLetterCharacters,omitempty"`

	// Minimum # of lower-case characters required in work profile password. Valid values 1 to 10
	WorkProfilePasswordMinLowerCaseCharacters nullable.Type[int64] `json:"workProfilePasswordMinLowerCaseCharacters,omitempty"`

	// Minimum # of non-letter characters required in work profile password. Valid values 1 to 10
	WorkProfilePasswordMinNonLetterCharacters nullable.Type[int64] `json:"workProfilePasswordMinNonLetterCharacters,omitempty"`

	// Minimum # of numeric characters required in work profile password. Valid values 1 to 10
	WorkProfilePasswordMinNumericCharacters nullable.Type[int64] `json:"workProfilePasswordMinNumericCharacters,omitempty"`

	// Minimum # of symbols required in work profile password. Valid values 1 to 10
	WorkProfilePasswordMinSymbolCharacters nullable.Type[int64] `json:"workProfilePasswordMinSymbolCharacters,omitempty"`

	// Minimum # of upper-case characters required in work profile password. Valid values 1 to 10
	WorkProfilePasswordMinUpperCaseCharacters nullable.Type[int64] `json:"workProfilePasswordMinUpperCaseCharacters,omitempty"`

	// Minimum length of work profile password. Valid values 4 to 16
	WorkProfilePasswordMinimumLength nullable.Type[int64] `json:"workProfilePasswordMinimumLength,omitempty"`

	// Minutes of inactivity before the screen times out.
	WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout nullable.Type[int64] `json:"workProfilePasswordMinutesOfInactivityBeforeScreenTimeout,omitempty"`

	// Number of previous work profile passwords to block. Valid values 0 to 24
	WorkProfilePasswordPreviousPasswordBlockCount nullable.Type[int64] `json:"workProfilePasswordPreviousPasswordBlockCount,omitempty"`

	// Android For Work required password type.
	WorkProfilePasswordRequiredType *AndroidForWorkRequiredPasswordType `json:"workProfilePasswordRequiredType,omitempty"`

	// Number of sign in failures allowed before work profile is removed and all corporate data deleted. Valid values 1 to
	// 16
	WorkProfilePasswordSignInFailureCountBeforeFactoryReset nullable.Type[int64] `json:"workProfilePasswordSignInFailureCountBeforeFactoryReset,omitempty"`

	// Password is required or not for work profile
	WorkProfileRequirePassword *bool `json:"workProfileRequirePassword,omitempty"`

	// The password complexity types that can be set on Android. One of: NONE, LOW, MEDIUM, HIGH. This is an API targeted to
	// Android 11+.
	WorkProfileRequiredPasswordComplexity *AndroidRequiredPasswordComplexity `json:"workProfileRequiredPasswordComplexity,omitempty"`

	// Fields inherited from DeviceConfiguration

	// The list of assignments for the device configuration profile.
	Assignments *[]DeviceConfigurationAssignment `json:"assignments,omitempty"`

	// DateTime the object was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Admin provided description of the Device Configuration.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The device mode applicability rule for this Policy.
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`

	// The OS edition applicability for this Policy.
	DeviceManagementApplicabilityRuleOsEdition *DeviceManagementApplicabilityRuleOsEdition `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`

	// The OS version applicability rule for this Policy.
	DeviceManagementApplicabilityRuleOsVersion *DeviceManagementApplicabilityRuleOsVersion `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`

	// Device Configuration Setting State Device Summary
	DeviceSettingStateSummaries *[]SettingStateDeviceSummary `json:"deviceSettingStateSummaries,omitempty"`

	// Device Configuration devices status overview
	DeviceStatusOverview *DeviceConfigurationDeviceOverview `json:"deviceStatusOverview,omitempty"`

	// Device configuration installation status by device.
	DeviceStatuses *[]DeviceConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`

	// Admin provided name of the device configuration.
	DisplayName *string `json:"displayName,omitempty"`

	// The list of group assignments for the device configuration profile.
	GroupAssignments *[]DeviceConfigurationGroupAssignment `json:"groupAssignments,omitempty"`

	// DateTime the object was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// List of Scope Tags for this Entity instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Indicates whether or not the underlying Device Configuration supports the assignment of scope tags. Assigning to the
	// ScopeTags property is not allowed when this value is false and entities will not be visible to scoped users. This
	// occurs for Legacy policies created in Silverlight and can be resolved by deleting and recreating the policy in the
	// Azure Portal. This property is read-only.
	SupportsScopeTags *bool `json:"supportsScopeTags,omitempty"`

	// Device Configuration users status overview
	UserStatusOverview *DeviceConfigurationUserOverview `json:"userStatusOverview,omitempty"`

	// Device configuration installation status by user.
	UserStatuses *[]DeviceConfigurationUserStatus `json:"userStatuses,omitempty"`

	// Version of the device configuration.
	Version *int64 `json:"version,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AndroidForWorkGeneralDeviceConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return BaseDeviceConfigurationImpl{
		Assignments:     s.Assignments,
		CreatedDateTime: s.CreatedDateTime,
		Description:     s.Description,
		DeviceManagementApplicabilityRuleDeviceMode: s.DeviceManagementApplicabilityRuleDeviceMode,
		DeviceManagementApplicabilityRuleOsEdition:  s.DeviceManagementApplicabilityRuleOsEdition,
		DeviceManagementApplicabilityRuleOsVersion:  s.DeviceManagementApplicabilityRuleOsVersion,
		DeviceSettingStateSummaries:                 s.DeviceSettingStateSummaries,
		DeviceStatusOverview:                        s.DeviceStatusOverview,
		DeviceStatuses:                              s.DeviceStatuses,
		DisplayName:                                 s.DisplayName,
		GroupAssignments:                            s.GroupAssignments,
		LastModifiedDateTime:                        s.LastModifiedDateTime,
		RoleScopeTagIds:                             s.RoleScopeTagIds,
		SupportsScopeTags:                           s.SupportsScopeTags,
		UserStatusOverview:                          s.UserStatusOverview,
		UserStatuses:                                s.UserStatuses,
		Version:                                     s.Version,
		Id:                                          s.Id,
		ODataId:                                     s.ODataId,
		ODataType:                                   s.ODataType,
	}
}

func (s AndroidForWorkGeneralDeviceConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidForWorkGeneralDeviceConfiguration{}

func (s AndroidForWorkGeneralDeviceConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper AndroidForWorkGeneralDeviceConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidForWorkGeneralDeviceConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidForWorkGeneralDeviceConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidForWorkGeneralDeviceConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidForWorkGeneralDeviceConfiguration: %+v", err)
	}

	return encoded, nil
}
