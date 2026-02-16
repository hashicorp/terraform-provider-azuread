package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ VpnConfiguration = AndroidDeviceOwnerVpnConfiguration{}

type AndroidDeviceOwnerVpnConfiguration struct {
	// Whether or not to enable always-on VPN connection.
	AlwaysOn nullable.Type[bool] `json:"alwaysOn,omitempty"`

	// If always-on VPN connection is enabled, whether or not to lock network traffic when that VPN is disconnected.
	AlwaysOnLockdown nullable.Type[bool] `json:"alwaysOnLockdown,omitempty"`

	// Android VPN connection type.
	ConnectionType *AndroidVpnConnectionType `json:"connectionType,omitempty"`

	// Custom data to define key/value pairs specific to a VPN provider. This collection can contain a maximum of 25
	// elements.
	CustomData *[]KeyValue `json:"customData,omitempty"`

	// Custom data to define key/value pairs specific to a VPN provider. This collection can contain a maximum of 25
	// elements.
	CustomKeyValueData *[]KeyValuePair `json:"customKeyValueData,omitempty"`

	// Tenant level settings for the Derived Credentials to be used for authentication.
	DerivedCredentialSettings *DeviceManagementDerivedCredentialSettings `json:"derivedCredentialSettings,omitempty"`

	// Identity certificate for client authentication when authentication method is certificate.
	IdentityCertificate *AndroidDeviceOwnerCertificateProfileBase `json:"identityCertificate,omitempty"`

	// Microsoft Tunnel site ID.
	MicrosoftTunnelSiteId nullable.Type[string] `json:"microsoftTunnelSiteId,omitempty"`

	// List of hosts to exclude using the proxy on connections for. These hosts can use wildcards such as .example.com.
	ProxyExclusionList *[]string `json:"proxyExclusionList,omitempty"`

	// Proxy server.
	ProxyServer VpnProxyServer `json:"proxyServer"`

	// Targeted mobile apps. This collection can contain a maximum of 500 elements.
	TargetedMobileApps *[]AppListItem `json:"targetedMobileApps,omitempty"`

	// Targeted App package IDs.
	TargetedPackageIds *[]string `json:"targetedPackageIds,omitempty"`

	// Fields inherited from VpnConfiguration

	// VPN Authentication Method.
	AuthenticationMethod *VpnAuthenticationMethod `json:"authenticationMethod,omitempty"`

	// Connection name displayed to the user.
	ConnectionName *string `json:"connectionName,omitempty"`

	// Realm when connection type is set to Pulse Secure.
	Realm nullable.Type[string] `json:"realm,omitempty"`

	// Role when connection type is set to Pulse Secure.
	Role nullable.Type[string] `json:"role,omitempty"`

	// List of VPN Servers on the network. Make sure end users can access these network locations. This collection can
	// contain a maximum of 500 elements.
	Servers *[]VpnServer `json:"servers,omitempty"`

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

func (s AndroidDeviceOwnerVpnConfiguration) VpnConfiguration() BaseVpnConfigurationImpl {
	return BaseVpnConfigurationImpl{
		AuthenticationMethod: s.AuthenticationMethod,
		ConnectionName:       s.ConnectionName,
		Realm:                s.Realm,
		Role:                 s.Role,
		Servers:              s.Servers,
		Assignments:          s.Assignments,
		CreatedDateTime:      s.CreatedDateTime,
		Description:          s.Description,
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

func (s AndroidDeviceOwnerVpnConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s AndroidDeviceOwnerVpnConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidDeviceOwnerVpnConfiguration{}

func (s AndroidDeviceOwnerVpnConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper AndroidDeviceOwnerVpnConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidDeviceOwnerVpnConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidDeviceOwnerVpnConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidDeviceOwnerVpnConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidDeviceOwnerVpnConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AndroidDeviceOwnerVpnConfiguration{}

func (s *AndroidDeviceOwnerVpnConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AlwaysOn                                    nullable.Type[bool]                          `json:"alwaysOn,omitempty"`
		AlwaysOnLockdown                            nullable.Type[bool]                          `json:"alwaysOnLockdown,omitempty"`
		ConnectionType                              *AndroidVpnConnectionType                    `json:"connectionType,omitempty"`
		CustomData                                  *[]KeyValue                                  `json:"customData,omitempty"`
		CustomKeyValueData                          *[]KeyValuePair                              `json:"customKeyValueData,omitempty"`
		DerivedCredentialSettings                   *DeviceManagementDerivedCredentialSettings   `json:"derivedCredentialSettings,omitempty"`
		MicrosoftTunnelSiteId                       nullable.Type[string]                        `json:"microsoftTunnelSiteId,omitempty"`
		ProxyExclusionList                          *[]string                                    `json:"proxyExclusionList,omitempty"`
		TargetedPackageIds                          *[]string                                    `json:"targetedPackageIds,omitempty"`
		AuthenticationMethod                        *VpnAuthenticationMethod                     `json:"authenticationMethod,omitempty"`
		ConnectionName                              *string                                      `json:"connectionName,omitempty"`
		Realm                                       nullable.Type[string]                        `json:"realm,omitempty"`
		Role                                        nullable.Type[string]                        `json:"role,omitempty"`
		Servers                                     *[]VpnServer                                 `json:"servers,omitempty"`
		Assignments                                 *[]DeviceConfigurationAssignment             `json:"assignments,omitempty"`
		CreatedDateTime                             *string                                      `json:"createdDateTime,omitempty"`
		Description                                 nullable.Type[string]                        `json:"description,omitempty"`
		DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
		DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition  `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
		DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion  `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
		DeviceSettingStateSummaries                 *[]SettingStateDeviceSummary                 `json:"deviceSettingStateSummaries,omitempty"`
		DeviceStatusOverview                        *DeviceConfigurationDeviceOverview           `json:"deviceStatusOverview,omitempty"`
		DeviceStatuses                              *[]DeviceConfigurationDeviceStatus           `json:"deviceStatuses,omitempty"`
		DisplayName                                 *string                                      `json:"displayName,omitempty"`
		GroupAssignments                            *[]DeviceConfigurationGroupAssignment        `json:"groupAssignments,omitempty"`
		LastModifiedDateTime                        *string                                      `json:"lastModifiedDateTime,omitempty"`
		RoleScopeTagIds                             *[]string                                    `json:"roleScopeTagIds,omitempty"`
		SupportsScopeTags                           *bool                                        `json:"supportsScopeTags,omitempty"`
		UserStatusOverview                          *DeviceConfigurationUserOverview             `json:"userStatusOverview,omitempty"`
		UserStatuses                                *[]DeviceConfigurationUserStatus             `json:"userStatuses,omitempty"`
		Version                                     *int64                                       `json:"version,omitempty"`
		Id                                          *string                                      `json:"id,omitempty"`
		ODataId                                     *string                                      `json:"@odata.id,omitempty"`
		ODataType                                   *string                                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AlwaysOn = decoded.AlwaysOn
	s.AlwaysOnLockdown = decoded.AlwaysOnLockdown
	s.ConnectionType = decoded.ConnectionType
	s.CustomData = decoded.CustomData
	s.CustomKeyValueData = decoded.CustomKeyValueData
	s.DerivedCredentialSettings = decoded.DerivedCredentialSettings
	s.MicrosoftTunnelSiteId = decoded.MicrosoftTunnelSiteId
	s.ProxyExclusionList = decoded.ProxyExclusionList
	s.TargetedPackageIds = decoded.TargetedPackageIds
	s.Assignments = decoded.Assignments
	s.AuthenticationMethod = decoded.AuthenticationMethod
	s.ConnectionName = decoded.ConnectionName
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DeviceManagementApplicabilityRuleDeviceMode = decoded.DeviceManagementApplicabilityRuleDeviceMode
	s.DeviceManagementApplicabilityRuleOsEdition = decoded.DeviceManagementApplicabilityRuleOsEdition
	s.DeviceManagementApplicabilityRuleOsVersion = decoded.DeviceManagementApplicabilityRuleOsVersion
	s.DeviceSettingStateSummaries = decoded.DeviceSettingStateSummaries
	s.DeviceStatusOverview = decoded.DeviceStatusOverview
	s.DeviceStatuses = decoded.DeviceStatuses
	s.DisplayName = decoded.DisplayName
	s.GroupAssignments = decoded.GroupAssignments
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Realm = decoded.Realm
	s.Role = decoded.Role
	s.RoleScopeTagIds = decoded.RoleScopeTagIds
	s.Servers = decoded.Servers
	s.SupportsScopeTags = decoded.SupportsScopeTags
	s.UserStatusOverview = decoded.UserStatusOverview
	s.UserStatuses = decoded.UserStatuses
	s.Version = decoded.Version

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AndroidDeviceOwnerVpnConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identityCertificate"]; ok {
		impl, err := UnmarshalAndroidDeviceOwnerCertificateProfileBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'IdentityCertificate' for 'AndroidDeviceOwnerVpnConfiguration': %+v", err)
		}
		s.IdentityCertificate = &impl
	}

	if v, ok := temp["proxyServer"]; ok {
		impl, err := UnmarshalVpnProxyServerImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ProxyServer' for 'AndroidDeviceOwnerVpnConfiguration': %+v", err)
		}
		s.ProxyServer = impl
	}

	if v, ok := temp["targetedMobileApps"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling TargetedMobileApps into list []json.RawMessage: %+v", err)
		}

		output := make([]AppListItem, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAppListItemImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'TargetedMobileApps' for 'AndroidDeviceOwnerVpnConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.TargetedMobileApps = &output
	}

	return nil
}
