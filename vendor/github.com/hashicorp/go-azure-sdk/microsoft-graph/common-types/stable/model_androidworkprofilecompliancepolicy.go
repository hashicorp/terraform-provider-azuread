package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceCompliancePolicy = AndroidWorkProfileCompliancePolicy{}

type AndroidWorkProfileCompliancePolicy struct {
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

	// Devices must not be jailbroken or rooted.
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
