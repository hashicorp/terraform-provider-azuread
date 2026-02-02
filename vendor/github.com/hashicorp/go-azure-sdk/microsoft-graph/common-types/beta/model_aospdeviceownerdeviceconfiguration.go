package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceConfiguration = AospDeviceOwnerDeviceConfiguration{}

type AospDeviceOwnerDeviceConfiguration struct {
	// Indicates whether or not the user is allowed to enable unknown sources setting. When set to true, user is not allowed
	// to enable unknown sources settings.
	AppsBlockInstallFromUnknownSources nullable.Type[bool] `json:"appsBlockInstallFromUnknownSources,omitempty"`

	// Indicates whether or not to block a user from configuring bluetooth.
	BluetoothBlockConfiguration nullable.Type[bool] `json:"bluetoothBlockConfiguration,omitempty"`

	// Indicates whether or not to disable the use of bluetooth. When set to true, bluetooth cannot be enabled on the
	// device.
	BluetoothBlocked nullable.Type[bool] `json:"bluetoothBlocked,omitempty"`

	// Indicates whether or not to disable the use of the camera.
	CameraBlocked nullable.Type[bool] `json:"cameraBlocked,omitempty"`

	// Indicates whether or not the factory reset option in settings is disabled.
	FactoryResetBlocked nullable.Type[bool] `json:"factoryResetBlocked,omitempty"`

	// Indicates the minimum length of the password required on the device. Valid values 4 to 16
	PasswordMinimumLength nullable.Type[int64] `json:"passwordMinimumLength,omitempty"`

	// Minutes of inactivity before the screen times out.
	PasswordMinutesOfInactivityBeforeScreenTimeout nullable.Type[int64] `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`

	// Indicates the minimum password quality required on the device. Possible values are: deviceDefault, required, numeric,
	// numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
	PasswordRequiredType *AndroidDeviceOwnerRequiredPasswordType `json:"passwordRequiredType,omitempty"`

	// Indicates the number of times a user can enter an incorrect password before the device is wiped. Valid values 4 to 11
	PasswordSignInFailureCountBeforeFactoryReset nullable.Type[int64] `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`

	// Indicates whether or not to disable the capability to take screenshots.
	ScreenCaptureBlocked nullable.Type[bool] `json:"screenCaptureBlocked,omitempty"`

	// Indicates whether or not to block the user from enabling debugging features on the device.
	SecurityAllowDebuggingFeatures nullable.Type[bool] `json:"securityAllowDebuggingFeatures,omitempty"`

	// Indicates whether or not to block external media.
	StorageBlockExternalMedia nullable.Type[bool] `json:"storageBlockExternalMedia,omitempty"`

	// Indicates whether or not to block USB file transfer.
	StorageBlockUsbFileTransfer nullable.Type[bool] `json:"storageBlockUsbFileTransfer,omitempty"`

	// Indicates whether or not to block the user from editing the wifi connection settings.
	WifiBlockEditConfigurations nullable.Type[bool] `json:"wifiBlockEditConfigurations,omitempty"`

	// Fields inherited from DeviceConfiguration

	// The list of assignments for the device configuration profile.
	Assignments *[]DeviceConfigurationAssignment `json:"assignments,omitempty"`

	// DateTime the object was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Admin provided description of the Device Configuration.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The device mode applicability rule for this Policy.
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`

	// The OS edition applicability for this Policy.
	DeviceManagementApplicabilityRuleOsEdition *DeviceManagementApplicabilityRuleOsEdition `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`

	// The OS version applicability rule for this Policy.
	DeviceManagementApplicabilityRuleOsVersion *DeviceManagementApplicabilityRuleOsVersion `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`

	// Device Configuration Setting State Device Summary
	DeviceSettingStateSummaries *[]SettingStateDeviceSummary `json:"deviceSettingStateSummaries,omitempty"`

	// Device Configuration devices status overview
	DeviceStatusOverview *DeviceConfigurationDeviceOverview `json:"deviceStatusOverview,omitempty"`

	// Device configuration installation status by device.
	DeviceStatuses *[]DeviceConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`

	// Admin provided name of the device configuration.
	DisplayName *string `json:"displayName,omitempty"`

	// The list of group assignments for the device configuration profile.
	GroupAssignments *[]DeviceConfigurationGroupAssignment `json:"groupAssignments,omitempty"`

	// DateTime the object was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// List of Scope Tags for this Entity instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Indicates whether or not the underlying Device Configuration supports the assignment of scope tags. Assigning to the
	// ScopeTags property is not allowed when this value is false and entities will not be visible to scoped users. This
	// occurs for Legacy policies created in Silverlight and can be resolved by deleting and recreating the policy in the
	// Azure Portal. This property is read-only.
	SupportsScopeTags *bool `json:"supportsScopeTags,omitempty"`

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

func (s AospDeviceOwnerDeviceConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return BaseDeviceConfigurationImpl{
		Assignments:     s.Assignments,
		CreatedDateTime: s.CreatedDateTime,
		Description:     s.Description,
		DeviceManagementApplicabilityRuleDeviceMode: s.DeviceManagementApplicabilityRuleDeviceMode,
		DeviceManagementApplicabilityRuleOsEdition:  s.DeviceManagementApplicabilityRuleOsEdition,
		DeviceManagementApplicabilityRuleOsVersion:  s.DeviceManagementApplicabilityRuleOsVersion,
		DeviceSettingStateSummaries:                 s.DeviceSettingStateSummaries,
		DeviceStatusOverview:                        s.DeviceStatusOverview,
		DeviceStatuses:                              s.DeviceStatuses,
		DisplayName:                                 s.DisplayName,
		GroupAssignments:                            s.GroupAssignments,
		LastModifiedDateTime:                        s.LastModifiedDateTime,
		RoleScopeTagIds:                             s.RoleScopeTagIds,
		SupportsScopeTags:                           s.SupportsScopeTags,
		UserStatusOverview:                          s.UserStatusOverview,
		UserStatuses:                                s.UserStatuses,
		Version:                                     s.Version,
		Id:                                          s.Id,
		ODataId:                                     s.ODataId,
		ODataType:                                   s.ODataType,
	}
}

func (s AospDeviceOwnerDeviceConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AospDeviceOwnerDeviceConfiguration{}

func (s AospDeviceOwnerDeviceConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper AospDeviceOwnerDeviceConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AospDeviceOwnerDeviceConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AospDeviceOwnerDeviceConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.aospDeviceOwnerDeviceConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AospDeviceOwnerDeviceConfiguration: %+v", err)
	}

	return encoded, nil
}
