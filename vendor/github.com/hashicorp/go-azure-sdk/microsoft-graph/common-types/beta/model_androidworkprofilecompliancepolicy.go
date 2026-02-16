package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceCompliancePolicy = AndroidWorkProfileCompliancePolicy{}

type AndroidWorkProfileCompliancePolicy struct {
	// Device threat protection levels for the Device Threat Protection API.
	AdvancedThreatProtectionRequiredSecurityLevel *DeviceThreatProtectionLevel `json:"advancedThreatProtectionRequiredSecurityLevel,omitempty"`

	// Require that devices have enabled device threat protection.
	DeviceThreatProtectionEnabled *bool `json:"deviceThreatProtectionEnabled,omitempty"`

	// Device threat protection levels for the Device Threat Protection API.
	DeviceThreatProtectionRequiredSecurityLevel *DeviceThreatProtectionLevel `json:"deviceThreatProtectionRequiredSecurityLevel,omitempty"`

	// Minimum Android security patch level.
	MinAndroidSecurityPatchLevel nullable.Type[string] `json:"minAndroidSecurityPatchLevel,omitempty"`

	// Maximum Android version.
	OsMaximumVersion nullable.Type[string] `json:"osMaximumVersion,omitempty"`

	// Minimum Android version.
	OsMinimumVersion nullable.Type[string] `json:"osMinimumVersion,omitempty"`

	// Number of days before the password expires. Valid values 1 to 365
	PasswordExpirationDays nullable.Type[int64] `json:"passwordExpirationDays,omitempty"`

	// Minimum password length. Valid values 4 to 16
	PasswordMinimumLength nullable.Type[int64] `json:"passwordMinimumLength,omitempty"`

	// Minutes of inactivity before a password is required.
	PasswordMinutesOfInactivityBeforeLock nullable.Type[int64] `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`

	// Number of previous passwords to block. Valid values 1 to 24
	PasswordPreviousPasswordBlockCount nullable.Type[int64] `json:"passwordPreviousPasswordBlockCount,omitempty"`

	// Require a password to unlock device.
	PasswordRequired *bool `json:"passwordRequired,omitempty"`

	// Android required password type.
	PasswordRequiredType *AndroidRequiredPasswordType `json:"passwordRequiredType,omitempty"`

	// Number of sign-in failures allowed before factory reset. Valid values 1 to 16
	PasswordSignInFailureCountBeforeFactoryReset nullable.Type[int64] `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`

	// The password complexity types that can be set on Android. One of: NONE, LOW, MEDIUM, HIGH. This is an API targeted to
	// Android 11+.
	RequiredPasswordComplexity *AndroidRequiredPasswordComplexity `json:"requiredPasswordComplexity,omitempty"`

	// Indicates the device should not be rooted. When TRUE, if the device is detected as rooted it will be reported
	// non-compliant. When FALSE, the device is not reported as non-compliant regardless of device rooted state. Default is
	// FALSE.
	SecurityBlockJailbrokenDevices *bool `json:"securityBlockJailbrokenDevices,omitempty"`

	// Disable USB debugging on Android devices.
	SecurityDisableUsbDebugging *bool `json:"securityDisableUsbDebugging,omitempty"`

	// Require that devices disallow installation of apps from unknown sources.
	SecurityPreventInstallAppsFromUnknownSources *bool `json:"securityPreventInstallAppsFromUnknownSources,omitempty"`

	// Require the device to pass the Company Portal client app runtime integrity check.
	SecurityRequireCompanyPortalAppIntegrity *bool `json:"securityRequireCompanyPortalAppIntegrity,omitempty"`

	// Require Google Play Services to be installed and enabled on the device.
	SecurityRequireGooglePlayServices *bool `json:"securityRequireGooglePlayServices,omitempty"`

	// Require the device to pass the Play Integrity basic integrity check.
	SecurityRequireSafetyNetAttestationBasicIntegrity *bool `json:"securityRequireSafetyNetAttestationBasicIntegrity,omitempty"`

	// Require the device to pass the Play Integrity device integrity check.
	SecurityRequireSafetyNetAttestationCertifiedDevice *bool `json:"securityRequireSafetyNetAttestationCertifiedDevice,omitempty"`

	// Require the device to have up to date security providers. The device will require Google Play Services to be enabled
	// and up to date.
	SecurityRequireUpToDateSecurityProviders *bool `json:"securityRequireUpToDateSecurityProviders,omitempty"`

	// Require the Android Verify apps feature is turned on.
	SecurityRequireVerifyApps *bool `json:"securityRequireVerifyApps,omitempty"`

	// An enum representing the Android Play Integrity API evaluation types.
	SecurityRequiredAndroidSafetyNetEvaluationType *AndroidSafetyNetEvaluationType `json:"securityRequiredAndroidSafetyNetEvaluationType,omitempty"`

	// Require encryption on Android devices.
	StorageRequireEncryption *bool `json:"storageRequireEncryption,omitempty"`

	// Minutes of inactivity before the screen times out.
	WorkProfileInactiveBeforeScreenLockInMinutes nullable.Type[int64] `json:"workProfileInactiveBeforeScreenLockInMinutes,omitempty"`

	// Number of days before the work profile password expires. Valid values 1 to 365
	WorkProfilePasswordExpirationInDays nullable.Type[int64] `json:"workProfilePasswordExpirationInDays,omitempty"`

	// Minimum length of work profile password. Valid values 4 to 16
	WorkProfilePasswordMinimumLength nullable.Type[int64] `json:"workProfilePasswordMinimumLength,omitempty"`

	// Android Work Profile required password type.
	WorkProfilePasswordRequiredType *AndroidWorkProfileRequiredPasswordType `json:"workProfilePasswordRequiredType,omitempty"`

	// Number of previous work profile passwords to block. Valid values 0 to 24
	WorkProfilePreviousPasswordBlockCount nullable.Type[int64] `json:"workProfilePreviousPasswordBlockCount,omitempty"`

	// Password is required or not for work profile
	WorkProfileRequirePassword *bool `json:"workProfileRequirePassword,omitempty"`

	// The password complexity types that can be set on Android. One of: NONE, LOW, MEDIUM, HIGH. This is an API targeted to
	// Android 11+.
	WorkProfileRequiredPasswordComplexity *AndroidRequiredPasswordComplexity `json:"workProfileRequiredPasswordComplexity,omitempty"`

	// Fields inherited from DeviceCompliancePolicy

	// The collection of assignments for this compliance policy.
	Assignments *[]DeviceCompliancePolicyAssignment `json:"assignments,omitempty"`

	// DateTime the object was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Admin provided description of the Device Configuration.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Compliance Setting State Device Summary
	DeviceSettingStateSummaries *[]SettingStateDeviceSummary `json:"deviceSettingStateSummaries,omitempty"`

	// Device compliance devices status overview
	DeviceStatusOverview *DeviceComplianceDeviceOverview `json:"deviceStatusOverview,omitempty"`

	// List of DeviceComplianceDeviceStatus.
	DeviceStatuses *[]DeviceComplianceDeviceStatus `json:"deviceStatuses,omitempty"`

	// Admin provided name of the device configuration.
	DisplayName *string `json:"displayName,omitempty"`

	// DateTime the object was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// List of Scope Tags for this Entity instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// The list of scheduled action per rule for this compliance policy. This is a required property when creating any
	// individual per-platform compliance policies.
	ScheduledActionsForRule *[]DeviceComplianceScheduledActionForRule `json:"scheduledActionsForRule,omitempty"`

	// Device compliance users status overview
	UserStatusOverview *DeviceComplianceUserOverview `json:"userStatusOverview,omitempty"`

	// List of DeviceComplianceUserStatus.
	UserStatuses *[]DeviceComplianceUserStatus `json:"userStatuses,omitempty"`

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

func (s AndroidWorkProfileCompliancePolicy) DeviceCompliancePolicy() BaseDeviceCompliancePolicyImpl {
	return BaseDeviceCompliancePolicyImpl{
		Assignments:                 s.Assignments,
		CreatedDateTime:             s.CreatedDateTime,
		Description:                 s.Description,
		DeviceSettingStateSummaries: s.DeviceSettingStateSummaries,
		DeviceStatusOverview:        s.DeviceStatusOverview,
		DeviceStatuses:              s.DeviceStatuses,
		DisplayName:                 s.DisplayName,
		LastModifiedDateTime:        s.LastModifiedDateTime,
		RoleScopeTagIds:             s.RoleScopeTagIds,
		ScheduledActionsForRule:     s.ScheduledActionsForRule,
		UserStatusOverview:          s.UserStatusOverview,
		UserStatuses:                s.UserStatuses,
		Version:                     s.Version,
		Id:                          s.Id,
		ODataId:                     s.ODataId,
		ODataType:                   s.ODataType,
	}
}

func (s AndroidWorkProfileCompliancePolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidWorkProfileCompliancePolicy{}

func (s AndroidWorkProfileCompliancePolicy) MarshalJSON() ([]byte, error) {
	type wrapper AndroidWorkProfileCompliancePolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidWorkProfileCompliancePolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidWorkProfileCompliancePolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidWorkProfileCompliancePolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidWorkProfileCompliancePolicy: %+v", err)
	}

	return encoded, nil
}
