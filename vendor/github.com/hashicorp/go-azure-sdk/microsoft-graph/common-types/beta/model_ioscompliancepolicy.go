package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceCompliancePolicy = IosCompliancePolicy{}

type IosCompliancePolicy struct {
	// Device threat protection levels for the Device Threat Protection API.
	AdvancedThreatProtectionRequiredSecurityLevel *DeviceThreatProtectionLevel `json:"advancedThreatProtectionRequiredSecurityLevel,omitempty"`

	// Require that devices have enabled device threat protection .
	DeviceThreatProtectionEnabled *bool `json:"deviceThreatProtectionEnabled,omitempty"`

	// Device threat protection levels for the Device Threat Protection API.
	DeviceThreatProtectionRequiredSecurityLevel *DeviceThreatProtectionLevel `json:"deviceThreatProtectionRequiredSecurityLevel,omitempty"`

	// Indicates whether or not to require a managed email profile.
	ManagedEmailProfileRequired *bool `json:"managedEmailProfileRequired,omitempty"`

	// Maximum IOS build version.
	OsMaximumBuildVersion nullable.Type[string] `json:"osMaximumBuildVersion,omitempty"`

	// Maximum IOS version.
	OsMaximumVersion nullable.Type[string] `json:"osMaximumVersion,omitempty"`

	// Minimum IOS build version.
	OsMinimumBuildVersion nullable.Type[string] `json:"osMinimumBuildVersion,omitempty"`

	// Minimum IOS version.
	OsMinimumVersion nullable.Type[string] `json:"osMinimumVersion,omitempty"`

	// Indicates whether or not to block simple passcodes.
	PasscodeBlockSimple *bool `json:"passcodeBlockSimple,omitempty"`

	// Number of days before the passcode expires. Valid values 1 to 65535
	PasscodeExpirationDays nullable.Type[int64] `json:"passcodeExpirationDays,omitempty"`

	// The number of character sets required in the password.
	PasscodeMinimumCharacterSetCount nullable.Type[int64] `json:"passcodeMinimumCharacterSetCount,omitempty"`

	// Minimum length of passcode. Valid values 4 to 14
	PasscodeMinimumLength nullable.Type[int64] `json:"passcodeMinimumLength,omitempty"`

	// Minutes of inactivity before a passcode is required.
	PasscodeMinutesOfInactivityBeforeLock nullable.Type[int64] `json:"passcodeMinutesOfInactivityBeforeLock,omitempty"`

	// Minutes of inactivity before the screen times out.
	PasscodeMinutesOfInactivityBeforeScreenTimeout nullable.Type[int64] `json:"passcodeMinutesOfInactivityBeforeScreenTimeout,omitempty"`

	// Number of previous passcodes to block. Valid values 1 to 24
	PasscodePreviousPasscodeBlockCount nullable.Type[int64] `json:"passcodePreviousPasscodeBlockCount,omitempty"`

	// Indicates whether or not to require a passcode.
	PasscodeRequired *bool `json:"passcodeRequired,omitempty"`

	// Possible values of required passwords.
	PasscodeRequiredType *RequiredPasswordType `json:"passcodeRequiredType,omitempty"`

	// Require the device to not have the specified apps installed. This collection can contain a maximum of 100 elements.
	RestrictedApps *[]AppListItem `json:"restrictedApps,omitempty"`

	// Indicates the device should not be jailbroken. When TRUE, if the device is detected as jailbroken it will be reported
	// non-compliant. When FALSE, the device is not reported as non-compliant regardless of device jailbroken state. Default
	// is FALSE.
	SecurityBlockJailbrokenDevices *bool `json:"securityBlockJailbrokenDevices,omitempty"`

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

func (s IosCompliancePolicy) DeviceCompliancePolicy() BaseDeviceCompliancePolicyImpl {
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

func (s IosCompliancePolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IosCompliancePolicy{}

func (s IosCompliancePolicy) MarshalJSON() ([]byte, error) {
	type wrapper IosCompliancePolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosCompliancePolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosCompliancePolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iosCompliancePolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosCompliancePolicy: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IosCompliancePolicy{}

func (s *IosCompliancePolicy) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AdvancedThreatProtectionRequiredSecurityLevel  *DeviceThreatProtectionLevel              `json:"advancedThreatProtectionRequiredSecurityLevel,omitempty"`
		DeviceThreatProtectionEnabled                  *bool                                     `json:"deviceThreatProtectionEnabled,omitempty"`
		DeviceThreatProtectionRequiredSecurityLevel    *DeviceThreatProtectionLevel              `json:"deviceThreatProtectionRequiredSecurityLevel,omitempty"`
		ManagedEmailProfileRequired                    *bool                                     `json:"managedEmailProfileRequired,omitempty"`
		OsMaximumBuildVersion                          nullable.Type[string]                     `json:"osMaximumBuildVersion,omitempty"`
		OsMaximumVersion                               nullable.Type[string]                     `json:"osMaximumVersion,omitempty"`
		OsMinimumBuildVersion                          nullable.Type[string]                     `json:"osMinimumBuildVersion,omitempty"`
		OsMinimumVersion                               nullable.Type[string]                     `json:"osMinimumVersion,omitempty"`
		PasscodeBlockSimple                            *bool                                     `json:"passcodeBlockSimple,omitempty"`
		PasscodeExpirationDays                         nullable.Type[int64]                      `json:"passcodeExpirationDays,omitempty"`
		PasscodeMinimumCharacterSetCount               nullable.Type[int64]                      `json:"passcodeMinimumCharacterSetCount,omitempty"`
		PasscodeMinimumLength                          nullable.Type[int64]                      `json:"passcodeMinimumLength,omitempty"`
		PasscodeMinutesOfInactivityBeforeLock          nullable.Type[int64]                      `json:"passcodeMinutesOfInactivityBeforeLock,omitempty"`
		PasscodeMinutesOfInactivityBeforeScreenTimeout nullable.Type[int64]                      `json:"passcodeMinutesOfInactivityBeforeScreenTimeout,omitempty"`
		PasscodePreviousPasscodeBlockCount             nullable.Type[int64]                      `json:"passcodePreviousPasscodeBlockCount,omitempty"`
		PasscodeRequired                               *bool                                     `json:"passcodeRequired,omitempty"`
		PasscodeRequiredType                           *RequiredPasswordType                     `json:"passcodeRequiredType,omitempty"`
		SecurityBlockJailbrokenDevices                 *bool                                     `json:"securityBlockJailbrokenDevices,omitempty"`
		Assignments                                    *[]DeviceCompliancePolicyAssignment       `json:"assignments,omitempty"`
		CreatedDateTime                                *string                                   `json:"createdDateTime,omitempty"`
		Description                                    nullable.Type[string]                     `json:"description,omitempty"`
		DeviceSettingStateSummaries                    *[]SettingStateDeviceSummary              `json:"deviceSettingStateSummaries,omitempty"`
		DeviceStatusOverview                           *DeviceComplianceDeviceOverview           `json:"deviceStatusOverview,omitempty"`
		DeviceStatuses                                 *[]DeviceComplianceDeviceStatus           `json:"deviceStatuses,omitempty"`
		DisplayName                                    *string                                   `json:"displayName,omitempty"`
		LastModifiedDateTime                           *string                                   `json:"lastModifiedDateTime,omitempty"`
		RoleScopeTagIds                                *[]string                                 `json:"roleScopeTagIds,omitempty"`
		ScheduledActionsForRule                        *[]DeviceComplianceScheduledActionForRule `json:"scheduledActionsForRule,omitempty"`
		UserStatusOverview                             *DeviceComplianceUserOverview             `json:"userStatusOverview,omitempty"`
		UserStatuses                                   *[]DeviceComplianceUserStatus             `json:"userStatuses,omitempty"`
		Version                                        *int64                                    `json:"version,omitempty"`
		Id                                             *string                                   `json:"id,omitempty"`
		ODataId                                        *string                                   `json:"@odata.id,omitempty"`
		ODataType                                      *string                                   `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AdvancedThreatProtectionRequiredSecurityLevel = decoded.AdvancedThreatProtectionRequiredSecurityLevel
	s.DeviceThreatProtectionEnabled = decoded.DeviceThreatProtectionEnabled
	s.DeviceThreatProtectionRequiredSecurityLevel = decoded.DeviceThreatProtectionRequiredSecurityLevel
	s.ManagedEmailProfileRequired = decoded.ManagedEmailProfileRequired
	s.OsMaximumBuildVersion = decoded.OsMaximumBuildVersion
	s.OsMaximumVersion = decoded.OsMaximumVersion
	s.OsMinimumBuildVersion = decoded.OsMinimumBuildVersion
	s.OsMinimumVersion = decoded.OsMinimumVersion
	s.PasscodeBlockSimple = decoded.PasscodeBlockSimple
	s.PasscodeExpirationDays = decoded.PasscodeExpirationDays
	s.PasscodeMinimumCharacterSetCount = decoded.PasscodeMinimumCharacterSetCount
	s.PasscodeMinimumLength = decoded.PasscodeMinimumLength
	s.PasscodeMinutesOfInactivityBeforeLock = decoded.PasscodeMinutesOfInactivityBeforeLock
	s.PasscodeMinutesOfInactivityBeforeScreenTimeout = decoded.PasscodeMinutesOfInactivityBeforeScreenTimeout
	s.PasscodePreviousPasscodeBlockCount = decoded.PasscodePreviousPasscodeBlockCount
	s.PasscodeRequired = decoded.PasscodeRequired
	s.PasscodeRequiredType = decoded.PasscodeRequiredType
	s.SecurityBlockJailbrokenDevices = decoded.SecurityBlockJailbrokenDevices
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
		return fmt.Errorf("unmarshaling IosCompliancePolicy into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'RestrictedApps' for 'IosCompliancePolicy': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.RestrictedApps = &output
	}

	return nil
}
