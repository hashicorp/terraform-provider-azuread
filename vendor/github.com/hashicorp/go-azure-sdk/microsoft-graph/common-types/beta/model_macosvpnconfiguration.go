package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AppleVpnConfiguration = MacOSVpnConfiguration{}

type MacOSVpnConfiguration struct {
	// Indicates the deployment channel type used to deploy the configuration profile. Possible values are deviceChannel,
	// userChannel. Possible values are: deviceChannel, userChannel, unknownFutureValue.
	DeploymentChannel *AppleDeploymentChannel `json:"deploymentChannel,omitempty"`

	// Identity certificate for client authentication when authentication method is certificate.
	IdentityCertificate *MacOSCertificateProfileBase `json:"identityCertificate,omitempty"`

	// Fields inherited from AppleVpnConfiguration

	// Associated Domains
	AssociatedDomains *[]string `json:"associatedDomains,omitempty"`

	// VPN Authentication Method.
	AuthenticationMethod *VpnAuthenticationMethod `json:"authenticationMethod,omitempty"`

	// Connection name displayed to the user.
	ConnectionName *string `json:"connectionName,omitempty"`

	// Apple VPN connection type.
	ConnectionType *AppleVpnConnectionType `json:"connectionType,omitempty"`

	// Custom data when connection type is set to Custom VPN. Use this field to enable functionality not supported by
	// Intune, but available in your VPN solution. Contact your VPN vendor to learn how to add these key/value pairs. This
	// collection can contain a maximum of 25 elements.
	CustomData *[]KeyValue `json:"customData,omitempty"`

	// Custom data when connection type is set to Custom VPN. Use this field to enable functionality not supported by
	// Intune, but available in your VPN solution. Contact your VPN vendor to learn how to add these key/value pairs. This
	// collection can contain a maximum of 25 elements.
	CustomKeyValueData *[]KeyValuePair `json:"customKeyValueData,omitempty"`

	// Toggle to prevent user from disabling automatic VPN in the Settings app
	DisableOnDemandUserOverride nullable.Type[bool] `json:"disableOnDemandUserOverride,omitempty"`

	// Whether to disconnect after on-demand connection idles
	DisconnectOnIdle nullable.Type[bool] `json:"disconnectOnIdle,omitempty"`

	// The length of time in seconds to wait before disconnecting an on-demand connection. Valid values 0 to 65535
	DisconnectOnIdleTimerInSeconds nullable.Type[int64] `json:"disconnectOnIdleTimerInSeconds,omitempty"`

	// Setting this to true creates Per-App VPN payload which can later be associated with Apps that can trigger this VPN
	// conneciton on the end user's iOS device.
	EnablePerApp nullable.Type[bool] `json:"enablePerApp,omitempty"`

	// Send all network traffic through VPN.
	EnableSplitTunneling *bool `json:"enableSplitTunneling,omitempty"`

	// Domains that are accessed through the public internet instead of through VPN, even when per-app VPN is activated
	ExcludedDomains *[]string `json:"excludedDomains,omitempty"`

	// Identifier provided by VPN vendor when connection type is set to Custom VPN. For example: Cisco AnyConnect uses an
	// identifier of the form com.cisco.anyconnect.applevpn.plugin
	Identifier nullable.Type[string] `json:"identifier,omitempty"`

	// Login group or domain when connection type is set to Dell SonicWALL Mobile Connection.
	LoginGroupOrDomain nullable.Type[string] `json:"loginGroupOrDomain,omitempty"`

	// On-Demand Rules. This collection can contain a maximum of 500 elements.
	OnDemandRules *[]VpnOnDemandRule `json:"onDemandRules,omitempty"`

	// Opt-In to sharing the device's Id to third-party vpn clients for use during network access control validation.
	OptInToDeviceIdSharing nullable.Type[bool] `json:"optInToDeviceIdSharing,omitempty"`

	// Provider type for per-app VPN. Possible values are: notConfigured, appProxy, packetTunnel.
	ProviderType *VpnProviderType `json:"providerType,omitempty"`

	// Proxy Server.
	ProxyServer VpnProxyServer `json:"proxyServer"`

	// Realm when connection type is set to Pulse Secure.
	Realm nullable.Type[string] `json:"realm,omitempty"`

	// Role when connection type is set to Pulse Secure.
	Role nullable.Type[string] `json:"role,omitempty"`

	// Safari domains when this VPN per App setting is enabled. In addition to the apps associated with this VPN, Safari
	// domains specified here will also be able to trigger this VPN connection.
	SafariDomains *[]string `json:"safariDomains,omitempty"`

	// VPN Server definition.
	Server *VpnServer `json:"server,omitempty"`

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

func (s MacOSVpnConfiguration) AppleVpnConfiguration() BaseAppleVpnConfigurationImpl {
	return BaseAppleVpnConfigurationImpl{
		AssociatedDomains:              s.AssociatedDomains,
		AuthenticationMethod:           s.AuthenticationMethod,
		ConnectionName:                 s.ConnectionName,
		ConnectionType:                 s.ConnectionType,
		CustomData:                     s.CustomData,
		CustomKeyValueData:             s.CustomKeyValueData,
		DisableOnDemandUserOverride:    s.DisableOnDemandUserOverride,
		DisconnectOnIdle:               s.DisconnectOnIdle,
		DisconnectOnIdleTimerInSeconds: s.DisconnectOnIdleTimerInSeconds,
		EnablePerApp:                   s.EnablePerApp,
		EnableSplitTunneling:           s.EnableSplitTunneling,
		ExcludedDomains:                s.ExcludedDomains,
		Identifier:                     s.Identifier,
		LoginGroupOrDomain:             s.LoginGroupOrDomain,
		OnDemandRules:                  s.OnDemandRules,
		OptInToDeviceIdSharing:         s.OptInToDeviceIdSharing,
		ProviderType:                   s.ProviderType,
		ProxyServer:                    s.ProxyServer,
		Realm:                          s.Realm,
		Role:                           s.Role,
		SafariDomains:                  s.SafariDomains,
		Server:                         s.Server,
		Assignments:                    s.Assignments,
		CreatedDateTime:                s.CreatedDateTime,
		Description:                    s.Description,
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

func (s MacOSVpnConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s MacOSVpnConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MacOSVpnConfiguration{}

func (s MacOSVpnConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper MacOSVpnConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MacOSVpnConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MacOSVpnConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.macOSVpnConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MacOSVpnConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &MacOSVpnConfiguration{}

func (s *MacOSVpnConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DeploymentChannel                           *AppleDeploymentChannel                      `json:"deploymentChannel,omitempty"`
		AssociatedDomains                           *[]string                                    `json:"associatedDomains,omitempty"`
		AuthenticationMethod                        *VpnAuthenticationMethod                     `json:"authenticationMethod,omitempty"`
		ConnectionName                              *string                                      `json:"connectionName,omitempty"`
		ConnectionType                              *AppleVpnConnectionType                      `json:"connectionType,omitempty"`
		CustomData                                  *[]KeyValue                                  `json:"customData,omitempty"`
		CustomKeyValueData                          *[]KeyValuePair                              `json:"customKeyValueData,omitempty"`
		DisableOnDemandUserOverride                 nullable.Type[bool]                          `json:"disableOnDemandUserOverride,omitempty"`
		DisconnectOnIdle                            nullable.Type[bool]                          `json:"disconnectOnIdle,omitempty"`
		DisconnectOnIdleTimerInSeconds              nullable.Type[int64]                         `json:"disconnectOnIdleTimerInSeconds,omitempty"`
		EnablePerApp                                nullable.Type[bool]                          `json:"enablePerApp,omitempty"`
		EnableSplitTunneling                        *bool                                        `json:"enableSplitTunneling,omitempty"`
		ExcludedDomains                             *[]string                                    `json:"excludedDomains,omitempty"`
		Identifier                                  nullable.Type[string]                        `json:"identifier,omitempty"`
		LoginGroupOrDomain                          nullable.Type[string]                        `json:"loginGroupOrDomain,omitempty"`
		OnDemandRules                               *[]VpnOnDemandRule                           `json:"onDemandRules,omitempty"`
		OptInToDeviceIdSharing                      nullable.Type[bool]                          `json:"optInToDeviceIdSharing,omitempty"`
		ProviderType                                *VpnProviderType                             `json:"providerType,omitempty"`
		Realm                                       nullable.Type[string]                        `json:"realm,omitempty"`
		Role                                        nullable.Type[string]                        `json:"role,omitempty"`
		SafariDomains                               *[]string                                    `json:"safariDomains,omitempty"`
		Server                                      *VpnServer                                   `json:"server,omitempty"`
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

	s.DeploymentChannel = decoded.DeploymentChannel
	s.Assignments = decoded.Assignments
	s.AssociatedDomains = decoded.AssociatedDomains
	s.AuthenticationMethod = decoded.AuthenticationMethod
	s.ConnectionName = decoded.ConnectionName
	s.ConnectionType = decoded.ConnectionType
	s.CreatedDateTime = decoded.CreatedDateTime
	s.CustomData = decoded.CustomData
	s.CustomKeyValueData = decoded.CustomKeyValueData
	s.Description = decoded.Description
	s.DeviceManagementApplicabilityRuleDeviceMode = decoded.DeviceManagementApplicabilityRuleDeviceMode
	s.DeviceManagementApplicabilityRuleOsEdition = decoded.DeviceManagementApplicabilityRuleOsEdition
	s.DeviceManagementApplicabilityRuleOsVersion = decoded.DeviceManagementApplicabilityRuleOsVersion
	s.DeviceSettingStateSummaries = decoded.DeviceSettingStateSummaries
	s.DeviceStatusOverview = decoded.DeviceStatusOverview
	s.DeviceStatuses = decoded.DeviceStatuses
	s.DisableOnDemandUserOverride = decoded.DisableOnDemandUserOverride
	s.DisconnectOnIdle = decoded.DisconnectOnIdle
	s.DisconnectOnIdleTimerInSeconds = decoded.DisconnectOnIdleTimerInSeconds
	s.DisplayName = decoded.DisplayName
	s.EnablePerApp = decoded.EnablePerApp
	s.EnableSplitTunneling = decoded.EnableSplitTunneling
	s.ExcludedDomains = decoded.ExcludedDomains
	s.GroupAssignments = decoded.GroupAssignments
	s.Id = decoded.Id
	s.Identifier = decoded.Identifier
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.LoginGroupOrDomain = decoded.LoginGroupOrDomain
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.OnDemandRules = decoded.OnDemandRules
	s.OptInToDeviceIdSharing = decoded.OptInToDeviceIdSharing
	s.ProviderType = decoded.ProviderType
	s.Realm = decoded.Realm
	s.Role = decoded.Role
	s.RoleScopeTagIds = decoded.RoleScopeTagIds
	s.SafariDomains = decoded.SafariDomains
	s.Server = decoded.Server
	s.SupportsScopeTags = decoded.SupportsScopeTags
	s.UserStatusOverview = decoded.UserStatusOverview
	s.UserStatuses = decoded.UserStatuses
	s.Version = decoded.Version

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MacOSVpnConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identityCertificate"]; ok {
		impl, err := UnmarshalMacOSCertificateProfileBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'IdentityCertificate' for 'MacOSVpnConfiguration': %+v", err)
		}
		s.IdentityCertificate = &impl
	}

	if v, ok := temp["proxyServer"]; ok {
		impl, err := UnmarshalVpnProxyServerImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ProxyServer' for 'MacOSVpnConfiguration': %+v", err)
		}
		s.ProxyServer = impl
	}

	return nil
}
