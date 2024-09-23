package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceCompliancePolicy = AndroidDeviceOwnerCompliancePolicy{}

type AndroidDeviceOwnerCompliancePolicy struct {
	// MDATP Require Mobile Threat Protection minimum risk level to report noncompliance. Possible values are: unavailable,
	// secured, low, medium, high, notSet.
	AdvancedThreatProtectionRequiredSecurityLevel *DeviceThreatProtectionLevel `json:"advancedThreatProtectionRequiredSecurityLevel,omitempty"`

	// Require that devices have enabled device threat protection.
	DeviceThreatProtectionEnabled nullable.Type[bool] `json:"deviceThreatProtectionEnabled,omitempty"`

	// Require Mobile Threat Protection minimum risk level to report noncompliance. Possible values are: unavailable,
	// secured, low, medium, high, notSet.
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

	// Indicates the minimum number of letter characters required for device password. Valid values 1 to 16
	PasswordMinimumLetterCharacters nullable.Type[int64] `json:"passwordMinimumLetterCharacters,omitempty"`

	// Indicates the minimum number of lower case characters required for device password. Valid values 1 to 16
	PasswordMinimumLowerCaseCharacters nullable.Type[int64] `json:"passwordMinimumLowerCaseCharacters,omitempty"`

	// Indicates the minimum number of non-letter characters required for device password. Valid values 1 to 16
	PasswordMinimumNonLetterCharacters nullable.Type[int64] `json:"passwordMinimumNonLetterCharacters,omitempty"`

	// Indicates the minimum number of numeric characters required for device password. Valid values 1 to 16
	PasswordMinimumNumericCharacters nullable.Type[int64] `json:"passwordMinimumNumericCharacters,omitempty"`

	// Indicates the minimum number of symbol characters required for device password. Valid values 1 to 16
	PasswordMinimumSymbolCharacters nullable.Type[int64] `json:"passwordMinimumSymbolCharacters,omitempty"`

	// Indicates the minimum number of upper case letter characters required for device password. Valid values 1 to 16
	PasswordMinimumUpperCaseCharacters nullable.Type[int64] `json:"passwordMinimumUpperCaseCharacters,omitempty"`

	// Minutes of inactivity before a password is required.
	PasswordMinutesOfInactivityBeforeLock nullable.Type[int64] `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`

	// Number of previous passwords to block. Valid values 1 to 24
	PasswordPreviousPasswordCountToBlock nullable.Type[int64] `json:"passwordPreviousPasswordCountToBlock,omitempty"`

	// Require a password to unlock device.
	PasswordRequired nullable.Type[bool] `json:"passwordRequired,omitempty"`

	// Type of characters in password. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic,
	// alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
	PasswordRequiredType *AndroidDeviceOwnerRequiredPasswordType `json:"passwordRequiredType,omitempty"`

	// Require device to have no pending Android system updates.
	RequireNoPendingSystemUpdates nullable.Type[bool] `json:"requireNoPendingSystemUpdates,omitempty"`

	// If setting is set to true, checks that the Intune app installed on fully managed, dedicated, or corporate-owned work
	// profile Android Enterprise enrolled devices, is the one provided by Microsoft from the Managed Google Playstore. If
	// the check fails, the device will be reported as non-compliant.
	SecurityRequireIntuneAppIntegrity nullable.Type[bool] `json:"securityRequireIntuneAppIntegrity,omitempty"`

	// Require the device to pass the Play Integrity basic integrity check.
	SecurityRequireSafetyNetAttestationBasicIntegrity nullable.Type[bool] `json:"securityRequireSafetyNetAttestationBasicIntegrity,omitempty"`

	// Require the device to pass the Play Integrity device integrity check.
	SecurityRequireSafetyNetAttestationCertifiedDevice nullable.Type[bool] `json:"securityRequireSafetyNetAttestationCertifiedDevice,omitempty"`

	// Require a specific Play Integrity evaluation type for compliance. Possible values are: basic, hardwareBacked.
	SecurityRequiredAndroidSafetyNetEvaluationType *AndroidSafetyNetEvaluationType `json:"securityRequiredAndroidSafetyNetEvaluationType,omitempty"`

	// Require encryption on Android devices.
	StorageRequireEncryption nullable.Type[bool] `json:"storageRequireEncryption,omitempty"`

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

	// The list of scheduled action for this rule
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

func (s AndroidDeviceOwnerCompliancePolicy) DeviceCompliancePolicy() BaseDeviceCompliancePolicyImpl {
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

func (s AndroidDeviceOwnerCompliancePolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidDeviceOwnerCompliancePolicy{}

func (s AndroidDeviceOwnerCompliancePolicy) MarshalJSON() ([]byte, error) {
	type wrapper AndroidDeviceOwnerCompliancePolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidDeviceOwnerCompliancePolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidDeviceOwnerCompliancePolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidDeviceOwnerCompliancePolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidDeviceOwnerCompliancePolicy: %+v", err)
	}

	return encoded, nil
}
