package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceCompliancePolicy = AospDeviceOwnerCompliancePolicy{}

type AospDeviceOwnerCompliancePolicy struct {
	// Minimum Android security patch level.
	MinAndroidSecurityPatchLevel nullable.Type[string] `json:"minAndroidSecurityPatchLevel,omitempty"`

	// Maximum Android version.
	OsMaximumVersion nullable.Type[string] `json:"osMaximumVersion,omitempty"`

	// Minimum Android version.
	OsMinimumVersion nullable.Type[string] `json:"osMinimumVersion,omitempty"`

	// Minimum password length. Valid values 4 to 16
	PasswordMinimumLength nullable.Type[int64] `json:"passwordMinimumLength,omitempty"`

	// Minutes of inactivity before a password is required. Valid values 1 to 8640
	PasswordMinutesOfInactivityBeforeLock nullable.Type[int64] `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`

	// Require a password to unlock device.
	PasswordRequired nullable.Type[bool] `json:"passwordRequired,omitempty"`

	// Type of characters in password. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic,
	// alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
	PasswordRequiredType *AndroidDeviceOwnerRequiredPasswordType `json:"passwordRequiredType,omitempty"`

	// Indicates the device should not be rooted. When TRUE, if the device is detected as rooted it will be reported
	// non-compliant. When FALSE, the device is not reported as non-compliant regardless of device rooted state. Default is
	// FALSE.
	SecurityBlockJailbrokenDevices nullable.Type[bool] `json:"securityBlockJailbrokenDevices,omitempty"`

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

func (s AospDeviceOwnerCompliancePolicy) DeviceCompliancePolicy() BaseDeviceCompliancePolicyImpl {
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

func (s AospDeviceOwnerCompliancePolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AospDeviceOwnerCompliancePolicy{}

func (s AospDeviceOwnerCompliancePolicy) MarshalJSON() ([]byte, error) {
	type wrapper AospDeviceOwnerCompliancePolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AospDeviceOwnerCompliancePolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AospDeviceOwnerCompliancePolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.aospDeviceOwnerCompliancePolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AospDeviceOwnerCompliancePolicy: %+v", err)
	}

	return encoded, nil
}
