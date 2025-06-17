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
	// Indicates the Microsoft Defender for Endpoint (also referred to Microsoft Defender Advanced Threat Protection
	// (MDATP)) minimum risk level to report noncompliance. Possible values are: unavailable, secured, low, medium, high,
	// notSet. Possible values are: unavailable, secured, low, medium, high, notSet.
	AdvancedThreatProtectionRequiredSecurityLevel *DeviceThreatProtectionLevel `json:"advancedThreatProtectionRequiredSecurityLevel,omitempty"`

	// Indicates whether the policy requires devices have device threat protection enabled. When TRUE, threat protection is
	// enabled. When FALSE, threat protection is not enabled. Default is FALSE.
	DeviceThreatProtectionEnabled nullable.Type[bool] `json:"deviceThreatProtectionEnabled,omitempty"`

	// Indicates the minimum mobile threat protection risk level to that results in Intune reporting device noncompliance.
	// Possible values are: unavailable, secured, low, medium, high, notSet. Possible values are: unavailable, secured, low,
	// medium, high, notSet.
	DeviceThreatProtectionRequiredSecurityLevel *DeviceThreatProtectionLevel `json:"deviceThreatProtectionRequiredSecurityLevel,omitempty"`

	// Indicates the minimum Android security patch level required to mark the device as compliant. For example: 'February
	// 1, 2025'
	MinAndroidSecurityPatchLevel nullable.Type[string] `json:"minAndroidSecurityPatchLevel,omitempty"`

	// Indicates the maximum Android version required to mark the device as compliant. For example: '15'
	OsMaximumVersion nullable.Type[string] `json:"osMaximumVersion,omitempty"`

	// Indicates the minimum Android version required to mark the device as compliant. For example: '14'
	OsMinimumVersion nullable.Type[string] `json:"osMinimumVersion,omitempty"`

	// Indicates the number of days before the password expires. Valid values 1 to 365.
	PasswordExpirationDays nullable.Type[int64] `json:"passwordExpirationDays,omitempty"`

	// Indicates the minimum password length required to mark the device as compliant. Valid values are 4 to 16, inclusive.
	// Valid values 4 to 16
	PasswordMinimumLength nullable.Type[int64] `json:"passwordMinimumLength,omitempty"`

	// Indicates the minimum number of letter characters required for device password for the device to be marked compliant.
	// Valid values 1 to 16.
	PasswordMinimumLetterCharacters nullable.Type[int64] `json:"passwordMinimumLetterCharacters,omitempty"`

	// Indicates the minimum number of lower case characters required for device password for the device to be marked
	// compliant. Valid values 1 to 16.
	PasswordMinimumLowerCaseCharacters nullable.Type[int64] `json:"passwordMinimumLowerCaseCharacters,omitempty"`

	// Indicates the minimum number of non-letter characters required for device password for the device to be marked
	// compliant. Valid values 1 to 16.
	PasswordMinimumNonLetterCharacters nullable.Type[int64] `json:"passwordMinimumNonLetterCharacters,omitempty"`

	// Indicates the minimum number of numeric characters required for device password for the device to be marked
	// compliant. Valid values 1 to 16.
	PasswordMinimumNumericCharacters nullable.Type[int64] `json:"passwordMinimumNumericCharacters,omitempty"`

	// Indicates the minimum number of symbol characters required for device password for the device to be marked compliant.
	// Valid values 1 to 16.
	PasswordMinimumSymbolCharacters nullable.Type[int64] `json:"passwordMinimumSymbolCharacters,omitempty"`

	// Indicates the minimum number of upper case letter characters required for device password for the device to be marked
	// compliant. Valid values 1 to 16.
	PasswordMinimumUpperCaseCharacters nullable.Type[int64] `json:"passwordMinimumUpperCaseCharacters,omitempty"`

	// Indicates the number of minutes of inactivity before a password is required.
	PasswordMinutesOfInactivityBeforeLock nullable.Type[int64] `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`

	// Indicates the number of previous passwords to block. Valid values 1 to 24.
	PasswordPreviousPasswordCountToBlock nullable.Type[int64] `json:"passwordPreviousPasswordCountToBlock,omitempty"`

	// Indicates whether a password is required to unlock the device. When TRUE, there must be a password set that unlocks
	// the device for the device to be marked as compliant. When FALSE, a device is marked as compliant whether or not a
	// password is set as required to unlock the device. Default is FALSE.
	PasswordRequired nullable.Type[bool] `json:"passwordRequired,omitempty"`

	// Indicates the password complexity requirement for the device to be marked compliant. Possible values are:
	// deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols,
	// lowSecurityBiometric, customPassword. Possible values are: deviceDefault, required, numeric, numericComplex,
	// alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
	PasswordRequiredType *AndroidDeviceOwnerRequiredPasswordType `json:"passwordRequiredType,omitempty"`

	// Indicates whether the device has pending security or OS updates and sets the compliance state accordingly. When TRUE,
	// checks if there are any pending system updates on each check in and if there are any pending security or OS version
	// updates (System Updates), the device will be reported as non-compliant. If set to FALSE, then checks for any pending
	// security or OS version updates (System Updates) are done without impact to device compliance state. Default is FALSE.
	RequireNoPendingSystemUpdates nullable.Type[bool] `json:"requireNoPendingSystemUpdates,omitempty"`

	// Indicates the device should not be rooted. When TRUE, if the device is detected as rooted it will be reported
	// non-compliant. When FALSE, the device is not reported as non-compliant regardless of device rooted state. Default is
	// FALSE.
	SecurityBlockJailbrokenDevices *bool `json:"securityBlockJailbrokenDevices,omitempty"`

	// Indicates whether Intune application integrity is required to mark the device as compliant. When TRUE, Intune checks
	// that the Intune app installed on fully managed, dedicated, or corporate-owned work profile Android Enterprise
	// enrolled devices, is the one provided by Microsoft from the Managed Google Play store. If the check fails, the device
	// will be reported as non-compliant. Default is FALSE.
	SecurityRequireIntuneAppIntegrity nullable.Type[bool] `json:"securityRequireIntuneAppIntegrity,omitempty"`

	// Indicates whether the compliance check will validate the Google Play Integrity check. When TRUE, the Google Play
	// integrity basic check must pass to consider the device compliant. When FALSE, the Google Play integrity basic check
	// can pass or fail and the device will be considered compliant. Default is FALSE.
	SecurityRequireSafetyNetAttestationBasicIntegrity nullable.Type[bool] `json:"securityRequireSafetyNetAttestationBasicIntegrity,omitempty"`

	// Indicates whether the compliance check will validate the Google Play Integrity check. When TRUE, the Google Play
	// integrity device check must pass to consider the device compliant. When FALSE, the Google Play integrity device check
	// can pass or fail and the device will be considered compliant. Default is FALSE.
	SecurityRequireSafetyNetAttestationCertifiedDevice nullable.Type[bool] `json:"securityRequireSafetyNetAttestationCertifiedDevice,omitempty"`

	// Indicates the types of measurements and reference data used to evaluate the device SafetyNet evaluation. Evaluation
	// is completed on the device to assess device integrity based on checks defined by Android and built into the device
	// hardware, for example, compromised OS version or root detection. Possible values are: basic, hardwareBacked, with
	// default value of basic. Possible values are: basic, hardwareBacked.
	SecurityRequiredAndroidSafetyNetEvaluationType *AndroidSafetyNetEvaluationType `json:"securityRequiredAndroidSafetyNetEvaluationType,omitempty"`

	// Indicates whether encryption on Android devices is required to mark the device as compliant.
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
