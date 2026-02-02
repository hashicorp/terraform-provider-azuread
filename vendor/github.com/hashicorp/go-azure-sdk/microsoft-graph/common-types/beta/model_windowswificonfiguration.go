package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWifiConfiguration interface {
	Entity
	DeviceConfiguration
	WindowsWifiConfiguration() BaseWindowsWifiConfigurationImpl
}

var _ WindowsWifiConfiguration = BaseWindowsWifiConfigurationImpl{}

type BaseWindowsWifiConfigurationImpl struct {
	// Specify whether the wifi connection should connect automatically when in range.
	ConnectAutomatically nullable.Type[bool] `json:"connectAutomatically,omitempty"`

	// Specify whether the wifi connection should connect to more preferred networks when already connected to this one.
	// Requires ConnectAutomatically to be true.
	ConnectToPreferredNetwork nullable.Type[bool] `json:"connectToPreferredNetwork,omitempty"`

	// Specify whether the wifi connection should connect automatically even when the SSID is not broadcasting.
	ConnectWhenNetworkNameIsHidden nullable.Type[bool] `json:"connectWhenNetworkNameIsHidden,omitempty"`

	// Specify whether to force FIPS compliance.
	ForceFIPSCompliance nullable.Type[bool] `json:"forceFIPSCompliance,omitempty"`

	// Specify the metered connection limit type for the wifi connection. Possible values are: unrestricted, fixed,
	// variable.
	MeteredConnectionLimit *MeteredConnectionLimitType `json:"meteredConnectionLimit,omitempty"`

	// Specify the network configuration name.
	NetworkName nullable.Type[string] `json:"networkName,omitempty"`

	// This is the pre-shared key for WPA Personal Wi-Fi network.
	PreSharedKey nullable.Type[string] `json:"preSharedKey,omitempty"`

	// Specify the URL for the proxy server configuration script.
	ProxyAutomaticConfigurationUrl nullable.Type[string] `json:"proxyAutomaticConfigurationUrl,omitempty"`

	// Specify the IP address for the proxy server.
	ProxyManualAddress nullable.Type[string] `json:"proxyManualAddress,omitempty"`

	// Specify the port for the proxy server.
	ProxyManualPort nullable.Type[int64] `json:"proxyManualPort,omitempty"`

	// Specify the proxy setting for Wi-Fi configuration. Possible values are: none, manual, automatic, unknownFutureValue.
	ProxySetting *WiFiProxySetting `json:"proxySetting,omitempty"`

	// Specify the SSID of the wifi connection.
	Ssid nullable.Type[string] `json:"ssid,omitempty"`

	// Specify the Wifi Security Type. Possible values are: open, wpaPersonal, wpaEnterprise, wep, wpa2Personal,
	// wpa2Enterprise.
	WifiSecurityType *WiFiSecurityType `json:"wifiSecurityType,omitempty"`

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

func (s BaseWindowsWifiConfigurationImpl) WindowsWifiConfiguration() BaseWindowsWifiConfigurationImpl {
	return s
}

func (s BaseWindowsWifiConfigurationImpl) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s BaseWindowsWifiConfigurationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ WindowsWifiConfiguration = RawWindowsWifiConfigurationImpl{}

// RawWindowsWifiConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindowsWifiConfigurationImpl struct {
	windowsWifiConfiguration BaseWindowsWifiConfigurationImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawWindowsWifiConfigurationImpl) WindowsWifiConfiguration() BaseWindowsWifiConfigurationImpl {
	return s.windowsWifiConfiguration
}

func (s RawWindowsWifiConfigurationImpl) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return s.windowsWifiConfiguration.DeviceConfiguration()
}

func (s RawWindowsWifiConfigurationImpl) Entity() BaseEntityImpl {
	return s.windowsWifiConfiguration.Entity()
}

var _ json.Marshaler = BaseWindowsWifiConfigurationImpl{}

func (s BaseWindowsWifiConfigurationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseWindowsWifiConfigurationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseWindowsWifiConfigurationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseWindowsWifiConfigurationImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsWifiConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseWindowsWifiConfigurationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalWindowsWifiConfigurationImplementation(input []byte) (WindowsWifiConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsWifiConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsWifiEnterpriseEAPConfiguration") {
		var out WindowsWifiEnterpriseEAPConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsWifiEnterpriseEAPConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsWifiConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsWifiConfigurationImpl: %+v", err)
	}

	return RawWindowsWifiConfigurationImpl{
		windowsWifiConfiguration: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
