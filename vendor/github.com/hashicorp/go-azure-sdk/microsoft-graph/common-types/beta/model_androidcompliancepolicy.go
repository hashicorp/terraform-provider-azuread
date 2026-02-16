package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceCompliancePolicy = AndroidCompliancePolicy{}

type AndroidCompliancePolicy struct {
	// Device threat protection levels for the Device Threat Protection API.
	AdvancedThreatProtectionRequiredSecurityLevel *DeviceThreatProtectionLevel `json:"advancedThreatProtectionRequiredSecurityLevel,omitempty"`

	// Condition statement id.
	ConditionStatementId nullable.Type[string] `json:"conditionStatementId,omitempty"`

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

	// Require the device to not have the specified apps installed. This collection can contain a maximum of 100 elements.
	RestrictedApps *[]AppListItem `json:"restrictedApps,omitempty"`

	// Block device administrator managed devices.
	SecurityBlockDeviceAdministratorManagedDevices *bool `json:"securityBlockDeviceAdministratorManagedDevices,omitempty"`

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

	// Require the device to pass the SafetyNet basic integrity check.
	SecurityRequireSafetyNetAttestationBasicIntegrity *bool `json:"securityRequireSafetyNetAttestationBasicIntegrity,omitempty"`

	// Require the device to pass the SafetyNet certified device check.
	SecurityRequireSafetyNetAttestationCertifiedDevice *bool `json:"securityRequireSafetyNetAttestationCertifiedDevice,omitempty"`

	// Require the device to have up to date security providers. The device will require Google Play Services to be enabled
	// and up to date.
	SecurityRequireUpToDateSecurityProviders *bool `json:"securityRequireUpToDateSecurityProviders,omitempty"`

	// Require the Android Verify apps feature is turned on.
	SecurityRequireVerifyApps *bool `json:"securityRequireVerifyApps,omitempty"`

	// Require encryption on Android devices.
	StorageRequireEncryption *bool `json:"storageRequireEncryption,omitempty"`

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

func (s AndroidCompliancePolicy) DeviceCompliancePolicy() BaseDeviceCompliancePolicyImpl {
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

func (s AndroidCompliancePolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidCompliancePolicy{}

func (s AndroidCompliancePolicy) MarshalJSON() ([]byte, error) {
	type wrapper AndroidCompliancePolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidCompliancePolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidCompliancePolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidCompliancePolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidCompliancePolicy: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AndroidCompliancePolicy{}

func (s *AndroidCompliancePolicy) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AdvancedThreatProtectionRequiredSecurityLevel      *DeviceThreatProtectionLevel              `json:"advancedThreatProtectionRequiredSecurityLevel,omitempty"`
		ConditionStatementId                               nullable.Type[string]                     `json:"conditionStatementId,omitempty"`
		DeviceThreatProtectionEnabled                      *bool                                     `json:"deviceThreatProtectionEnabled,omitempty"`
		DeviceThreatProtectionRequiredSecurityLevel        *DeviceThreatProtectionLevel              `json:"deviceThreatProtectionRequiredSecurityLevel,omitempty"`
		MinAndroidSecurityPatchLevel                       nullable.Type[string]                     `json:"minAndroidSecurityPatchLevel,omitempty"`
		OsMaximumVersion                                   nullable.Type[string]                     `json:"osMaximumVersion,omitempty"`
		OsMinimumVersion                                   nullable.Type[string]                     `json:"osMinimumVersion,omitempty"`
		PasswordExpirationDays                             nullable.Type[int64]                      `json:"passwordExpirationDays,omitempty"`
		PasswordMinimumLength                              nullable.Type[int64]                      `json:"passwordMinimumLength,omitempty"`
		PasswordMinutesOfInactivityBeforeLock              nullable.Type[int64]                      `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`
		PasswordPreviousPasswordBlockCount                 nullable.Type[int64]                      `json:"passwordPreviousPasswordBlockCount,omitempty"`
		PasswordRequired                                   *bool                                     `json:"passwordRequired,omitempty"`
		PasswordRequiredType                               *AndroidRequiredPasswordType              `json:"passwordRequiredType,omitempty"`
		PasswordSignInFailureCountBeforeFactoryReset       nullable.Type[int64]                      `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`
		RequiredPasswordComplexity                         *AndroidRequiredPasswordComplexity        `json:"requiredPasswordComplexity,omitempty"`
		SecurityBlockDeviceAdministratorManagedDevices     *bool                                     `json:"securityBlockDeviceAdministratorManagedDevices,omitempty"`
		SecurityBlockJailbrokenDevices                     *bool                                     `json:"securityBlockJailbrokenDevices,omitempty"`
		SecurityDisableUsbDebugging                        *bool                                     `json:"securityDisableUsbDebugging,omitempty"`
		SecurityPreventInstallAppsFromUnknownSources       *bool                                     `json:"securityPreventInstallAppsFromUnknownSources,omitempty"`
		SecurityRequireCompanyPortalAppIntegrity           *bool                                     `json:"securityRequireCompanyPortalAppIntegrity,omitempty"`
		SecurityRequireGooglePlayServices                  *bool                                     `json:"securityRequireGooglePlayServices,omitempty"`
		SecurityRequireSafetyNetAttestationBasicIntegrity  *bool                                     `json:"securityRequireSafetyNetAttestationBasicIntegrity,omitempty"`
		SecurityRequireSafetyNetAttestationCertifiedDevice *bool                                     `json:"securityRequireSafetyNetAttestationCertifiedDevice,omitempty"`
		SecurityRequireUpToDateSecurityProviders           *bool                                     `json:"securityRequireUpToDateSecurityProviders,omitempty"`
		SecurityRequireVerifyApps                          *bool                                     `json:"securityRequireVerifyApps,omitempty"`
		StorageRequireEncryption                           *bool                                     `json:"storageRequireEncryption,omitempty"`
		Assignments                                        *[]DeviceCompliancePolicyAssignment       `json:"assignments,omitempty"`
		CreatedDateTime                                    *string                                   `json:"createdDateTime,omitempty"`
		Description                                        nullable.Type[string]                     `json:"description,omitempty"`
		DeviceSettingStateSummaries                        *[]SettingStateDeviceSummary              `json:"deviceSettingStateSummaries,omitempty"`
		DeviceStatusOverview                               *DeviceComplianceDeviceOverview           `json:"deviceStatusOverview,omitempty"`
		DeviceStatuses                                     *[]DeviceComplianceDeviceStatus           `json:"deviceStatuses,omitempty"`
		DisplayName                                        *string                                   `json:"displayName,omitempty"`
		LastModifiedDateTime                               *string                                   `json:"lastModifiedDateTime,omitempty"`
		RoleScopeTagIds                                    *[]string                                 `json:"roleScopeTagIds,omitempty"`
		ScheduledActionsForRule                            *[]DeviceComplianceScheduledActionForRule `json:"scheduledActionsForRule,omitempty"`
		UserStatusOverview                                 *DeviceComplianceUserOverview             `json:"userStatusOverview,omitempty"`
		UserStatuses                                       *[]DeviceComplianceUserStatus             `json:"userStatuses,omitempty"`
		Version                                            *int64                                    `json:"version,omitempty"`
		Id                                                 *string                                   `json:"id,omitempty"`
		ODataId                                            *string                                   `json:"@odata.id,omitempty"`
		ODataType                                          *string                                   `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AdvancedThreatProtectionRequiredSecurityLevel = decoded.AdvancedThreatProtectionRequiredSecurityLevel
	s.ConditionStatementId = decoded.ConditionStatementId
	s.DeviceThreatProtectionEnabled = decoded.DeviceThreatProtectionEnabled
	s.DeviceThreatProtectionRequiredSecurityLevel = decoded.DeviceThreatProtectionRequiredSecurityLevel
	s.MinAndroidSecurityPatchLevel = decoded.MinAndroidSecurityPatchLevel
	s.OsMaximumVersion = decoded.OsMaximumVersion
	s.OsMinimumVersion = decoded.OsMinimumVersion
	s.PasswordExpirationDays = decoded.PasswordExpirationDays
	s.PasswordMinimumLength = decoded.PasswordMinimumLength
	s.PasswordMinutesOfInactivityBeforeLock = decoded.PasswordMinutesOfInactivityBeforeLock
	s.PasswordPreviousPasswordBlockCount = decoded.PasswordPreviousPasswordBlockCount
	s.PasswordRequired = decoded.PasswordRequired
	s.PasswordRequiredType = decoded.PasswordRequiredType
	s.PasswordSignInFailureCountBeforeFactoryReset = decoded.PasswordSignInFailureCountBeforeFactoryReset
	s.RequiredPasswordComplexity = decoded.RequiredPasswordComplexity
	s.SecurityBlockDeviceAdministratorManagedDevices = decoded.SecurityBlockDeviceAdministratorManagedDevices
	s.SecurityBlockJailbrokenDevices = decoded.SecurityBlockJailbrokenDevices
	s.SecurityDisableUsbDebugging = decoded.SecurityDisableUsbDebugging
	s.SecurityPreventInstallAppsFromUnknownSources = decoded.SecurityPreventInstallAppsFromUnknownSources
	s.SecurityRequireCompanyPortalAppIntegrity = decoded.SecurityRequireCompanyPortalAppIntegrity
	s.SecurityRequireGooglePlayServices = decoded.SecurityRequireGooglePlayServices
	s.SecurityRequireSafetyNetAttestationBasicIntegrity = decoded.SecurityRequireSafetyNetAttestationBasicIntegrity
	s.SecurityRequireSafetyNetAttestationCertifiedDevice = decoded.SecurityRequireSafetyNetAttestationCertifiedDevice
	s.SecurityRequireUpToDateSecurityProviders = decoded.SecurityRequireUpToDateSecurityProviders
	s.SecurityRequireVerifyApps = decoded.SecurityRequireVerifyApps
	s.StorageRequireEncryption = decoded.StorageRequireEncryption
	s.Assignments = decoded.Assignments
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DeviceSettingStateSummaries = decoded.DeviceSettingStateSummaries
	s.DeviceStatusOverview = decoded.DeviceStatusOverview
	s.DeviceStatuses = decoded.DeviceStatuses
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RoleScopeTagIds = decoded.RoleScopeTagIds
	s.ScheduledActionsForRule = decoded.ScheduledActionsForRule
	s.UserStatusOverview = decoded.UserStatusOverview
	s.UserStatuses = decoded.UserStatuses
	s.Version = decoded.Version

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AndroidCompliancePolicy into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["restrictedApps"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling RestrictedApps into list []json.RawMessage: %+v", err)
		}

		output := make([]AppListItem, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAppListItemImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'RestrictedApps' for 'AndroidCompliancePolicy': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.RestrictedApps = &output
	}

	return nil
}
