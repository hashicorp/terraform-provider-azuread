package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsVpnConfiguration = Windows10VpnConfiguration{}

type Windows10VpnConfiguration struct {
	// Associated Apps. This collection can contain a maximum of 10000 elements.
	AssociatedApps *[]Windows10AssociatedApps `json:"associatedApps,omitempty"`

	// Windows 10 VPN connection types.
	AuthenticationMethod *Windows10VpnAuthenticationMethod `json:"authenticationMethod,omitempty"`

	// VPN connection types.
	ConnectionType *Windows10VpnConnectionType `json:"connectionType,omitempty"`

	// Cryptography Suite security settings for IKEv2 VPN in Windows10 and above
	CryptographySuite *CryptographySuite `json:"cryptographySuite,omitempty"`

	// DNS rules. This collection can contain a maximum of 1000 elements.
	DnsRules *[]VpnDnsRule `json:"dnsRules,omitempty"`

	// Specify DNS suffixes to add to the DNS search list to properly route short names.
	DnsSuffixes *[]string `json:"dnsSuffixes,omitempty"`

	// Extensible Authentication Protocol (EAP) XML. (UTF8 encoded byte array)
	EapXml nullable.Type[string] `json:"eapXml,omitempty"`

	// Enable Always On mode.
	EnableAlwaysOn nullable.Type[bool] `json:"enableAlwaysOn,omitempty"`

	// Enable conditional access.
	EnableConditionalAccess *bool `json:"enableConditionalAccess,omitempty"`

	// Enable device tunnel.
	EnableDeviceTunnel nullable.Type[bool] `json:"enableDeviceTunnel,omitempty"`

	// Enable IP address registration with internal DNS.
	EnableDnsRegistration nullable.Type[bool] `json:"enableDnsRegistration,omitempty"`

	// Enable single sign-on (SSO) with alternate certificate.
	EnableSingleSignOnWithAlternateCertificate *bool `json:"enableSingleSignOnWithAlternateCertificate,omitempty"`

	// Enable split tunneling.
	EnableSplitTunneling *bool `json:"enableSplitTunneling,omitempty"`

	// Identity certificate for client authentication when authentication method is certificate.
	IdentityCertificate *WindowsCertificateProfileBase `json:"identityCertificate,omitempty"`

	// ID of the Microsoft Tunnel site associated with the VPN profile.
	MicrosoftTunnelSiteId nullable.Type[string] `json:"microsoftTunnelSiteId,omitempty"`

	// Only associated Apps can use connection (per-app VPN).
	OnlyAssociatedAppsCanUseConnection nullable.Type[bool] `json:"onlyAssociatedAppsCanUseConnection,omitempty"`

	// Profile target type. Possible values are: user, device, autoPilotDevice.
	ProfileTarget *Windows10VpnProfileTarget `json:"profileTarget,omitempty"`

	// Proxy Server.
	ProxyServer *Windows10VpnProxyServer `json:"proxyServer,omitempty"`

	// Remember user credentials.
	RememberUserCredentials *bool `json:"rememberUserCredentials,omitempty"`

	// Routes (optional for third-party providers). This collection can contain a maximum of 1000 elements.
	Routes *[]VpnRoute `json:"routes,omitempty"`

	// Single sign-on Extended Key Usage (EKU).
	SingleSignOnEku *ExtendedKeyUsage `json:"singleSignOnEku,omitempty"`

	// Single sign-on issuer hash.
	SingleSignOnIssuerHash nullable.Type[string] `json:"singleSignOnIssuerHash,omitempty"`

	// Traffic rules. This collection can contain a maximum of 1000 elements.
	TrafficRules *[]VpnTrafficRule `json:"trafficRules,omitempty"`

	// Trusted Network Domains
	TrustedNetworkDomains *[]string `json:"trustedNetworkDomains,omitempty"`

	// Windows Information Protection (WIP) domain to associate with this connection.
	WindowsInformationProtectionDomain nullable.Type[string] `json:"windowsInformationProtectionDomain,omitempty"`

	// Fields inherited from WindowsVpnConfiguration

	// Connection name displayed to the user.
	ConnectionName *string `json:"connectionName,omitempty"`

	// Custom XML commands that configures the VPN connection. (UTF8 encoded byte array)
	CustomXml nullable.Type[string] `json:"customXml,omitempty"`

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

func (s Windows10VpnConfiguration) WindowsVpnConfiguration() BaseWindowsVpnConfigurationImpl {
	return BaseWindowsVpnConfigurationImpl{
		ConnectionName:  s.ConnectionName,
		CustomXml:       s.CustomXml,
		Servers:         s.Servers,
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

func (s Windows10VpnConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s Windows10VpnConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Windows10VpnConfiguration{}

func (s Windows10VpnConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper Windows10VpnConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Windows10VpnConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Windows10VpnConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windows10VpnConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Windows10VpnConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Windows10VpnConfiguration{}

func (s *Windows10VpnConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AssociatedApps                              *[]Windows10AssociatedApps                   `json:"associatedApps,omitempty"`
		AuthenticationMethod                        *Windows10VpnAuthenticationMethod            `json:"authenticationMethod,omitempty"`
		ConnectionType                              *Windows10VpnConnectionType                  `json:"connectionType,omitempty"`
		CryptographySuite                           *CryptographySuite                           `json:"cryptographySuite,omitempty"`
		DnsRules                                    *[]VpnDnsRule                                `json:"dnsRules,omitempty"`
		DnsSuffixes                                 *[]string                                    `json:"dnsSuffixes,omitempty"`
		EapXml                                      nullable.Type[string]                        `json:"eapXml,omitempty"`
		EnableAlwaysOn                              nullable.Type[bool]                          `json:"enableAlwaysOn,omitempty"`
		EnableConditionalAccess                     *bool                                        `json:"enableConditionalAccess,omitempty"`
		EnableDeviceTunnel                          nullable.Type[bool]                          `json:"enableDeviceTunnel,omitempty"`
		EnableDnsRegistration                       nullable.Type[bool]                          `json:"enableDnsRegistration,omitempty"`
		EnableSingleSignOnWithAlternateCertificate  *bool                                        `json:"enableSingleSignOnWithAlternateCertificate,omitempty"`
		EnableSplitTunneling                        *bool                                        `json:"enableSplitTunneling,omitempty"`
		MicrosoftTunnelSiteId                       nullable.Type[string]                        `json:"microsoftTunnelSiteId,omitempty"`
		OnlyAssociatedAppsCanUseConnection          nullable.Type[bool]                          `json:"onlyAssociatedAppsCanUseConnection,omitempty"`
		ProfileTarget                               *Windows10VpnProfileTarget                   `json:"profileTarget,omitempty"`
		ProxyServer                                 *Windows10VpnProxyServer                     `json:"proxyServer,omitempty"`
		RememberUserCredentials                     *bool                                        `json:"rememberUserCredentials,omitempty"`
		Routes                                      *[]VpnRoute                                  `json:"routes,omitempty"`
		SingleSignOnEku                             *ExtendedKeyUsage                            `json:"singleSignOnEku,omitempty"`
		SingleSignOnIssuerHash                      nullable.Type[string]                        `json:"singleSignOnIssuerHash,omitempty"`
		TrafficRules                                *[]VpnTrafficRule                            `json:"trafficRules,omitempty"`
		TrustedNetworkDomains                       *[]string                                    `json:"trustedNetworkDomains,omitempty"`
		WindowsInformationProtectionDomain          nullable.Type[string]                        `json:"windowsInformationProtectionDomain,omitempty"`
		ConnectionName                              *string                                      `json:"connectionName,omitempty"`
		CustomXml                                   nullable.Type[string]                        `json:"customXml,omitempty"`
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

	s.AssociatedApps = decoded.AssociatedApps
	s.AuthenticationMethod = decoded.AuthenticationMethod
	s.ConnectionType = decoded.ConnectionType
	s.CryptographySuite = decoded.CryptographySuite
	s.DnsRules = decoded.DnsRules
	s.DnsSuffixes = decoded.DnsSuffixes
	s.EapXml = decoded.EapXml
	s.EnableAlwaysOn = decoded.EnableAlwaysOn
	s.EnableConditionalAccess = decoded.EnableConditionalAccess
	s.EnableDeviceTunnel = decoded.EnableDeviceTunnel
	s.EnableDnsRegistration = decoded.EnableDnsRegistration
	s.EnableSingleSignOnWithAlternateCertificate = decoded.EnableSingleSignOnWithAlternateCertificate
	s.EnableSplitTunneling = decoded.EnableSplitTunneling
	s.MicrosoftTunnelSiteId = decoded.MicrosoftTunnelSiteId
	s.OnlyAssociatedAppsCanUseConnection = decoded.OnlyAssociatedAppsCanUseConnection
	s.ProfileTarget = decoded.ProfileTarget
	s.ProxyServer = decoded.ProxyServer
	s.RememberUserCredentials = decoded.RememberUserCredentials
	s.Routes = decoded.Routes
	s.SingleSignOnEku = decoded.SingleSignOnEku
	s.SingleSignOnIssuerHash = decoded.SingleSignOnIssuerHash
	s.TrafficRules = decoded.TrafficRules
	s.TrustedNetworkDomains = decoded.TrustedNetworkDomains
	s.WindowsInformationProtectionDomain = decoded.WindowsInformationProtectionDomain
	s.Assignments = decoded.Assignments
	s.ConnectionName = decoded.ConnectionName
	s.CreatedDateTime = decoded.CreatedDateTime
	s.CustomXml = decoded.CustomXml
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
	s.RoleScopeTagIds = decoded.RoleScopeTagIds
	s.Servers = decoded.Servers
	s.SupportsScopeTags = decoded.SupportsScopeTags
	s.UserStatusOverview = decoded.UserStatusOverview
	s.UserStatuses = decoded.UserStatuses
	s.Version = decoded.Version

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Windows10VpnConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identityCertificate"]; ok {
		impl, err := UnmarshalWindowsCertificateProfileBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'IdentityCertificate' for 'Windows10VpnConfiguration': %+v", err)
		}
		s.IdentityCertificate = &impl
	}

	return nil
}
