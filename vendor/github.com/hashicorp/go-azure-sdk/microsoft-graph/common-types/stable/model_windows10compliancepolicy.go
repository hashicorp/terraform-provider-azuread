package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceCompliancePolicy = Windows10CompliancePolicy{}

type Windows10CompliancePolicy struct {
	// Require devices to be reported healthy by Windows Device Health Attestation - bit locker is enabled
	BitLockerEnabled *bool `json:"bitLockerEnabled,omitempty"`

	// Require devices to be reported as healthy by Windows Device Health Attestation.
	CodeIntegrityEnabled *bool `json:"codeIntegrityEnabled,omitempty"`

	// Require devices to be reported as healthy by Windows Device Health Attestation - early launch antimalware driver is
	// enabled.
	EarlyLaunchAntiMalwareDriverEnabled *bool `json:"earlyLaunchAntiMalwareDriverEnabled,omitempty"`

	// Maximum Windows Phone version.
	MobileOsMaximumVersion nullable.Type[string] `json:"mobileOsMaximumVersion,omitempty"`

	// Minimum Windows Phone version.
	MobileOsMinimumVersion nullable.Type[string] `json:"mobileOsMinimumVersion,omitempty"`

	// Maximum Windows 10 version.
	OsMaximumVersion nullable.Type[string] `json:"osMaximumVersion,omitempty"`

	// Minimum Windows 10 version.
	OsMinimumVersion nullable.Type[string] `json:"osMinimumVersion,omitempty"`

	// Indicates whether or not to block simple password.
	PasswordBlockSimple *bool `json:"passwordBlockSimple,omitempty"`

	// The password expiration in days.
	PasswordExpirationDays nullable.Type[int64] `json:"passwordExpirationDays,omitempty"`

	// The number of character sets required in the password.
	PasswordMinimumCharacterSetCount nullable.Type[int64] `json:"passwordMinimumCharacterSetCount,omitempty"`

	// The minimum password length.
	PasswordMinimumLength nullable.Type[int64] `json:"passwordMinimumLength,omitempty"`

	// Minutes of inactivity before a password is required.
	PasswordMinutesOfInactivityBeforeLock nullable.Type[int64] `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`

	// The number of previous passwords to prevent re-use of.
	PasswordPreviousPasswordBlockCount nullable.Type[int64] `json:"passwordPreviousPasswordBlockCount,omitempty"`

	// Require a password to unlock Windows device.
	PasswordRequired *bool `json:"passwordRequired,omitempty"`

	// Require a password to unlock an idle device.
	PasswordRequiredToUnlockFromIdle *bool `json:"passwordRequiredToUnlockFromIdle,omitempty"`

	// Possible values of required passwords.
	PasswordRequiredType *RequiredPasswordType `json:"passwordRequiredType,omitempty"`

	// Require devices to be reported as healthy by Windows Device Health Attestation.
	RequireHealthyDeviceReport *bool `json:"requireHealthyDeviceReport,omitempty"`

	// Require devices to be reported as healthy by Windows Device Health Attestation - secure boot is enabled.
	SecureBootEnabled *bool `json:"secureBootEnabled,omitempty"`

	// Require encryption on windows devices.
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

func (s Windows10CompliancePolicy) DeviceCompliancePolicy() BaseDeviceCompliancePolicyImpl {
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

func (s Windows10CompliancePolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Windows10CompliancePolicy{}

func (s Windows10CompliancePolicy) MarshalJSON() ([]byte, error) {
	type wrapper Windows10CompliancePolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Windows10CompliancePolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Windows10CompliancePolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windows10CompliancePolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Windows10CompliancePolicy: %+v", err)
	}

	return encoded, nil
}
