package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceCompliancePolicy = Windows81CompliancePolicy{}

type Windows81CompliancePolicy struct {
	// Maximum Windows 8.1 version.
	OsMaximumVersion nullable.Type[string] `json:"osMaximumVersion,omitempty"`

	// Minimum Windows 8.1 version.
	OsMinimumVersion nullable.Type[string] `json:"osMinimumVersion,omitempty"`

	// Indicates whether or not to block simple password.
	PasswordBlockSimple *bool `json:"passwordBlockSimple,omitempty"`

	// Password expiration in days.
	PasswordExpirationDays nullable.Type[int64] `json:"passwordExpirationDays,omitempty"`

	// The number of character sets required in the password.
	PasswordMinimumCharacterSetCount nullable.Type[int64] `json:"passwordMinimumCharacterSetCount,omitempty"`

	// The minimum password length.
	PasswordMinimumLength nullable.Type[int64] `json:"passwordMinimumLength,omitempty"`

	// Minutes of inactivity before a password is required.
	PasswordMinutesOfInactivityBeforeLock nullable.Type[int64] `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`

	// The number of previous passwords to prevent re-use of. Valid values 0 to 24
	PasswordPreviousPasswordBlockCount nullable.Type[int64] `json:"passwordPreviousPasswordBlockCount,omitempty"`

	// Require a password to unlock Windows device.
	PasswordRequired *bool `json:"passwordRequired,omitempty"`

	// Possible values of required passwords.
	PasswordRequiredType *RequiredPasswordType `json:"passwordRequiredType,omitempty"`

	// Indicates whether or not to require encryption on a windows 8.1 device.
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

func (s Windows81CompliancePolicy) DeviceCompliancePolicy() BaseDeviceCompliancePolicyImpl {
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

func (s Windows81CompliancePolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Windows81CompliancePolicy{}

func (s Windows81CompliancePolicy) MarshalJSON() ([]byte, error) {
	type wrapper Windows81CompliancePolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Windows81CompliancePolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Windows81CompliancePolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windows81CompliancePolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Windows81CompliancePolicy: %+v", err)
	}

	return encoded, nil
}
