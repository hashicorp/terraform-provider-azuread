package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkWiFiConfiguration interface {
	Entity
	DeviceConfiguration
	AndroidForWorkWiFiConfiguration() BaseAndroidForWorkWiFiConfigurationImpl
}

var _ AndroidForWorkWiFiConfiguration = BaseAndroidForWorkWiFiConfigurationImpl{}

type BaseAndroidForWorkWiFiConfigurationImpl struct {
	// Connect automatically when this network is in range. Setting this to true will skip the user prompt and automatically
	// connect the device to Wi-Fi network.
	ConnectAutomatically *bool `json:"connectAutomatically,omitempty"`

	// When set to true, this profile forces the device to connect to a network that doesn't broadcast its SSID to all
	// devices.
	ConnectWhenNetworkNameIsHidden *bool `json:"connectWhenNetworkNameIsHidden,omitempty"`

	// Network Name
	NetworkName *string `json:"networkName,omitempty"`

	// This is the name of the Wi-Fi network that is broadcast to all devices.
	Ssid *string `json:"ssid,omitempty"`

	// The possible security types for Android Wi-Fi profiles. Default value 'Open', indicates no authentication required
	// for the network. The security protocols supported are WEP, WPA and WPA2. 'WpaEnterprise' and 'Wpa2Enterprise' options
	// are available for Enterprise Wi-Fi profiles. 'Wep' and 'WpaPersonal' (supports WPA and WPA2) options are available
	// for Basic Wi-Fi profiles.
	WiFiSecurityType *AndroidWiFiSecurityType `json:"wiFiSecurityType,omitempty"`

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

func (s BaseAndroidForWorkWiFiConfigurationImpl) AndroidForWorkWiFiConfiguration() BaseAndroidForWorkWiFiConfigurationImpl {
	return s
}

func (s BaseAndroidForWorkWiFiConfigurationImpl) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s BaseAndroidForWorkWiFiConfigurationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AndroidForWorkWiFiConfiguration = RawAndroidForWorkWiFiConfigurationImpl{}

// RawAndroidForWorkWiFiConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAndroidForWorkWiFiConfigurationImpl struct {
	androidForWorkWiFiConfiguration BaseAndroidForWorkWiFiConfigurationImpl
	Type                            string
	Values                          map[string]interface{}
}

func (s RawAndroidForWorkWiFiConfigurationImpl) AndroidForWorkWiFiConfiguration() BaseAndroidForWorkWiFiConfigurationImpl {
	return s.androidForWorkWiFiConfiguration
}

func (s RawAndroidForWorkWiFiConfigurationImpl) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return s.androidForWorkWiFiConfiguration.DeviceConfiguration()
}

func (s RawAndroidForWorkWiFiConfigurationImpl) Entity() BaseEntityImpl {
	return s.androidForWorkWiFiConfiguration.Entity()
}

var _ json.Marshaler = BaseAndroidForWorkWiFiConfigurationImpl{}

func (s BaseAndroidForWorkWiFiConfigurationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAndroidForWorkWiFiConfigurationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAndroidForWorkWiFiConfigurationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAndroidForWorkWiFiConfigurationImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidForWorkWiFiConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAndroidForWorkWiFiConfigurationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAndroidForWorkWiFiConfigurationImplementation(input []byte) (AndroidForWorkWiFiConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidForWorkWiFiConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.androidForWorkEnterpriseWiFiConfiguration") {
		var out AndroidForWorkEnterpriseWiFiConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidForWorkEnterpriseWiFiConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseAndroidForWorkWiFiConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAndroidForWorkWiFiConfigurationImpl: %+v", err)
	}

	return RawAndroidForWorkWiFiConfigurationImpl{
		androidForWorkWiFiConfiguration: parent,
		Type:                            value,
		Values:                          temp,
	}, nil

}
