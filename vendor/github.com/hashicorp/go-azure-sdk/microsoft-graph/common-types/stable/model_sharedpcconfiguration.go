package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceConfiguration = SharedPCConfiguration{}

type SharedPCConfiguration struct {
	// Specifies how accounts are managed on a shared PC. Only applies when disableAccountManager is false.
	AccountManagerPolicy *SharedPCAccountManagerPolicy `json:"accountManagerPolicy,omitempty"`

	// Specifies whether local storage is allowed on a shared PC.
	AllowLocalStorage *bool `json:"allowLocalStorage,omitempty"`

	// Type of accounts that are allowed to share the PC.
	AllowedAccounts *SharedPCAllowedAccountType `json:"allowedAccounts,omitempty"`

	// Disables the account manager for shared PC mode.
	DisableAccountManager *bool `json:"disableAccountManager,omitempty"`

	// Specifies whether the default shared PC education environment policies should be disabled. For Windows 10 RS2 and
	// later, this policy will be applied without setting Enabled to true.
	DisableEduPolicies *bool `json:"disableEduPolicies,omitempty"`

	// Specifies whether the default shared PC power policies should be disabled.
	DisablePowerPolicies *bool `json:"disablePowerPolicies,omitempty"`

	// Disables the requirement to sign in whenever the device wakes up from sleep mode.
	DisableSignInOnResume *bool `json:"disableSignInOnResume,omitempty"`

	// Enables shared PC mode and applies the shared pc policies.
	Enabled *bool `json:"enabled,omitempty"`

	// Specifies the time in seconds that a device must sit idle before the PC goes to sleep. Setting this value to 0
	// prevents the sleep timeout from occurring.
	IdleTimeBeforeSleepInSeconds nullable.Type[int64] `json:"idleTimeBeforeSleepInSeconds,omitempty"`

	// Specifies the display text for the account shown on the sign-in screen which launches the app specified by
	// SetKioskAppUserModelId. Only applies when KioskAppUserModelId is set.
	KioskAppDisplayName nullable.Type[string] `json:"kioskAppDisplayName,omitempty"`

	// Specifies the application user model ID of the app to use with assigned access.
	KioskAppUserModelId nullable.Type[string] `json:"kioskAppUserModelId,omitempty"`

	// Specifies the daily start time of maintenance hour.
	MaintenanceStartTime nullable.Type[string] `json:"maintenanceStartTime,omitempty"`

	// Fields inherited from DeviceConfiguration

	// The list of assignments for the device configuration profile.
	Assignments *[]DeviceConfigurationAssignment `json:"assignments,omitempty"`

	// DateTime the object was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Admin provided description of the Device Configuration.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Device Configuration Setting State Device Summary
	DeviceSettingStateSummaries *[]SettingStateDeviceSummary `json:"deviceSettingStateSummaries,omitempty"`

	// Device Configuration devices status overview
	DeviceStatusOverview *DeviceConfigurationDeviceOverview `json:"deviceStatusOverview,omitempty"`

	// Device configuration installation status by device.
	DeviceStatuses *[]DeviceConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`

	// Admin provided name of the device configuration.
	DisplayName *string `json:"displayName,omitempty"`

	// DateTime the object was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

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

func (s SharedPCConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return BaseDeviceConfigurationImpl{
		Assignments:                 s.Assignments,
		CreatedDateTime:             s.CreatedDateTime,
		Description:                 s.Description,
		DeviceSettingStateSummaries: s.DeviceSettingStateSummaries,
		DeviceStatusOverview:        s.DeviceStatusOverview,
		DeviceStatuses:              s.DeviceStatuses,
		DisplayName:                 s.DisplayName,
		LastModifiedDateTime:        s.LastModifiedDateTime,
		UserStatusOverview:          s.UserStatusOverview,
		UserStatuses:                s.UserStatuses,
		Version:                     s.Version,
		Id:                          s.Id,
		ODataId:                     s.ODataId,
		ODataType:                   s.ODataType,
	}
}

func (s SharedPCConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SharedPCConfiguration{}

func (s SharedPCConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper SharedPCConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SharedPCConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SharedPCConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.sharedPCConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SharedPCConfiguration: %+v", err)
	}

	return encoded, nil
}
