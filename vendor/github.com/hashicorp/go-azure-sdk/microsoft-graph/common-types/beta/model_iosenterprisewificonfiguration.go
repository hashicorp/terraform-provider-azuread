package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IosWiFiConfiguration = IosEnterpriseWiFiConfiguration{}

type IosEnterpriseWiFiConfiguration struct {
	// Authentication Method when EAP Type is configured to PEAP or EAP-TTLS. Possible values are: certificate,
	// usernameAndPassword, derivedCredential.
	AuthenticationMethod *WiFiAuthenticationMethod `json:"authenticationMethod,omitempty"`

	// Tenant level settings for the Derived Credentials to be used for authentication.
	DerivedCredentialSettings *DeviceManagementDerivedCredentialSettings `json:"derivedCredentialSettings,omitempty"`

	// EAP-FAST Configuration Option when EAP-FAST is the selected EAP Type. Possible values are:
	// noProtectedAccessCredential, useProtectedAccessCredential, useProtectedAccessCredentialAndProvision,
	// useProtectedAccessCredentialAndProvisionAnonymously.
	EapFastConfiguration *EapFastConfiguration `json:"eapFastConfiguration,omitempty"`

	// Extensible Authentication Protocol (EAP) configuration types.
	EapType *EapType `json:"eapType,omitempty"`

	// Identity Certificate for client authentication when EAP Type is configured to EAP-TLS, EAP-TTLS (with Certificate
	// Authentication), or PEAP (with Certificate Authentication).
	IdentityCertificateForClientAuthentication *IosCertificateProfileBase `json:"identityCertificateForClientAuthentication,omitempty"`

	// Non-EAP Method for Authentication when EAP Type is EAP-TTLS and Authenticationmethod is Username and Password.
	// Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap,
	// microsoftChapVersionTwo.
	InnerAuthenticationProtocolForEapTtls *NonEapAuthenticationMethodForEapTtlsType `json:"innerAuthenticationProtocolForEapTtls,omitempty"`

	// Enable identity privacy (Outer Identity) when EAP Type is configured to EAP - TTLS, EAP - FAST or PEAP. This property
	// masks usernames with the text you enter. For example, if you use 'anonymous', each user that authenticates with this
	// Wi-Fi connection using their real username is displayed as 'anonymous'.
	OuterIdentityPrivacyTemporaryValue nullable.Type[string] `json:"outerIdentityPrivacyTemporaryValue,omitempty"`

	// Password format string used to build the password to connect to wifi
	PasswordFormatString nullable.Type[string] `json:"passwordFormatString,omitempty"`

	// Trusted Root Certificates for Server Validation when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. If you
	// provide this value you do not need to provide trustedServerCertificateNames, and vice versa. This collection can
	// contain a maximum of 500 elements.
	RootCertificatesForServerValidation *[]IosTrustedRootCertificate `json:"rootCertificatesForServerValidation,omitempty"`

	// Trusted server certificate names when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. This is the common name
	// used in the certificates issued by your trusted certificate authority (CA). If you provide this information, you can
	// bypass the dynamic trust dialog that is displayed on end users' devices when they connect to this Wi-Fi network.
	TrustedServerCertificateNames *[]string `json:"trustedServerCertificateNames,omitempty"`

	// Username format string used to build the username to connect to wifi
	UsernameFormatString nullable.Type[string] `json:"usernameFormatString,omitempty"`

	// Fields inherited from IosWiFiConfiguration

	// Connect automatically when this network is in range. Setting this to true will skip the user prompt and automatically
	// connect the device to Wi-Fi network.
	ConnectAutomatically *bool `json:"connectAutomatically,omitempty"`

	// Connect when the network is not broadcasting its name (SSID). When set to true, this profile forces the device to
	// connect to a network that doesn't broadcast its SSID to all devices.
	ConnectWhenNetworkNameIsHidden *bool `json:"connectWhenNetworkNameIsHidden,omitempty"`

	// If set to true, forces devices connecting using this Wi-Fi profile to present their actual Wi-Fi MAC address instead
	// of a random MAC address. Applies to iOS 14 and later.
	DisableMacAddressRandomization nullable.Type[bool] `json:"disableMacAddressRandomization,omitempty"`

	// Network Name
	NetworkName *string `json:"networkName,omitempty"`

	// This is the pre-shared key for WPA Personal Wi-Fi network.
	PreSharedKey nullable.Type[string] `json:"preSharedKey,omitempty"`

	// URL of the proxy server automatic configuration script when automatic configuration is selected. This URL is
	// typically the location of PAC (Proxy Auto Configuration) file.
	ProxyAutomaticConfigurationUrl nullable.Type[string] `json:"proxyAutomaticConfigurationUrl,omitempty"`

	// IP Address or DNS hostname of the proxy server when manual configuration is selected.
	ProxyManualAddress nullable.Type[string] `json:"proxyManualAddress,omitempty"`

	// Port of the proxy server when manual configuration is selected.
	ProxyManualPort nullable.Type[int64] `json:"proxyManualPort,omitempty"`

	// Wi-Fi Proxy Settings.
	ProxySettings *WiFiProxySetting `json:"proxySettings,omitempty"`

	// This is the name of the Wi-Fi network that is broadcast to all devices.
	Ssid *string `json:"ssid,omitempty"`

	// Wi-Fi Security Types.
	WiFiSecurityType *WiFiSecurityType `json:"wiFiSecurityType,omitempty"`

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

func (s IosEnterpriseWiFiConfiguration) IosWiFiConfiguration() BaseIosWiFiConfigurationImpl {
	return BaseIosWiFiConfigurationImpl{
		ConnectAutomatically:           s.ConnectAutomatically,
		ConnectWhenNetworkNameIsHidden: s.ConnectWhenNetworkNameIsHidden,
		DisableMacAddressRandomization: s.DisableMacAddressRandomization,
		NetworkName:                    s.NetworkName,
		PreSharedKey:                   s.PreSharedKey,
		ProxyAutomaticConfigurationUrl: s.ProxyAutomaticConfigurationUrl,
		ProxyManualAddress:             s.ProxyManualAddress,
		ProxyManualPort:                s.ProxyManualPort,
		ProxySettings:                  s.ProxySettings,
		Ssid:                           s.Ssid,
		WiFiSecurityType:               s.WiFiSecurityType,
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

func (s IosEnterpriseWiFiConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s IosEnterpriseWiFiConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IosEnterpriseWiFiConfiguration{}

func (s IosEnterpriseWiFiConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper IosEnterpriseWiFiConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosEnterpriseWiFiConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosEnterpriseWiFiConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iosEnterpriseWiFiConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosEnterpriseWiFiConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IosEnterpriseWiFiConfiguration{}

func (s *IosEnterpriseWiFiConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AuthenticationMethod                        *WiFiAuthenticationMethod                    `json:"authenticationMethod,omitempty"`
		DerivedCredentialSettings                   *DeviceManagementDerivedCredentialSettings   `json:"derivedCredentialSettings,omitempty"`
		EapFastConfiguration                        *EapFastConfiguration                        `json:"eapFastConfiguration,omitempty"`
		EapType                                     *EapType                                     `json:"eapType,omitempty"`
		InnerAuthenticationProtocolForEapTtls       *NonEapAuthenticationMethodForEapTtlsType    `json:"innerAuthenticationProtocolForEapTtls,omitempty"`
		OuterIdentityPrivacyTemporaryValue          nullable.Type[string]                        `json:"outerIdentityPrivacyTemporaryValue,omitempty"`
		PasswordFormatString                        nullable.Type[string]                        `json:"passwordFormatString,omitempty"`
		RootCertificatesForServerValidation         *[]IosTrustedRootCertificate                 `json:"rootCertificatesForServerValidation,omitempty"`
		TrustedServerCertificateNames               *[]string                                    `json:"trustedServerCertificateNames,omitempty"`
		UsernameFormatString                        nullable.Type[string]                        `json:"usernameFormatString,omitempty"`
		ConnectAutomatically                        *bool                                        `json:"connectAutomatically,omitempty"`
		ConnectWhenNetworkNameIsHidden              *bool                                        `json:"connectWhenNetworkNameIsHidden,omitempty"`
		DisableMacAddressRandomization              nullable.Type[bool]                          `json:"disableMacAddressRandomization,omitempty"`
		NetworkName                                 *string                                      `json:"networkName,omitempty"`
		PreSharedKey                                nullable.Type[string]                        `json:"preSharedKey,omitempty"`
		ProxyAutomaticConfigurationUrl              nullable.Type[string]                        `json:"proxyAutomaticConfigurationUrl,omitempty"`
		ProxyManualAddress                          nullable.Type[string]                        `json:"proxyManualAddress,omitempty"`
		ProxyManualPort                             nullable.Type[int64]                         `json:"proxyManualPort,omitempty"`
		ProxySettings                               *WiFiProxySetting                            `json:"proxySettings,omitempty"`
		Ssid                                        *string                                      `json:"ssid,omitempty"`
		WiFiSecurityType                            *WiFiSecurityType                            `json:"wiFiSecurityType,omitempty"`
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

	s.AuthenticationMethod = decoded.AuthenticationMethod
	s.DerivedCredentialSettings = decoded.DerivedCredentialSettings
	s.EapFastConfiguration = decoded.EapFastConfiguration
	s.EapType = decoded.EapType
	s.InnerAuthenticationProtocolForEapTtls = decoded.InnerAuthenticationProtocolForEapTtls
	s.OuterIdentityPrivacyTemporaryValue = decoded.OuterIdentityPrivacyTemporaryValue
	s.PasswordFormatString = decoded.PasswordFormatString
	s.RootCertificatesForServerValidation = decoded.RootCertificatesForServerValidation
	s.TrustedServerCertificateNames = decoded.TrustedServerCertificateNames
	s.UsernameFormatString = decoded.UsernameFormatString
	s.Assignments = decoded.Assignments
	s.ConnectAutomatically = decoded.ConnectAutomatically
	s.ConnectWhenNetworkNameIsHidden = decoded.ConnectWhenNetworkNameIsHidden
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DeviceManagementApplicabilityRuleDeviceMode = decoded.DeviceManagementApplicabilityRuleDeviceMode
	s.DeviceManagementApplicabilityRuleOsEdition = decoded.DeviceManagementApplicabilityRuleOsEdition
	s.DeviceManagementApplicabilityRuleOsVersion = decoded.DeviceManagementApplicabilityRuleOsVersion
	s.DeviceSettingStateSummaries = decoded.DeviceSettingStateSummaries
	s.DeviceStatusOverview = decoded.DeviceStatusOverview
	s.DeviceStatuses = decoded.DeviceStatuses
	s.DisableMacAddressRandomization = decoded.DisableMacAddressRandomization
	s.DisplayName = decoded.DisplayName
	s.GroupAssignments = decoded.GroupAssignments
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.NetworkName = decoded.NetworkName
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PreSharedKey = decoded.PreSharedKey
	s.ProxyAutomaticConfigurationUrl = decoded.ProxyAutomaticConfigurationUrl
	s.ProxyManualAddress = decoded.ProxyManualAddress
	s.ProxyManualPort = decoded.ProxyManualPort
	s.ProxySettings = decoded.ProxySettings
	s.RoleScopeTagIds = decoded.RoleScopeTagIds
	s.Ssid = decoded.Ssid
	s.SupportsScopeTags = decoded.SupportsScopeTags
	s.UserStatusOverview = decoded.UserStatusOverview
	s.UserStatuses = decoded.UserStatuses
	s.Version = decoded.Version
	s.WiFiSecurityType = decoded.WiFiSecurityType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IosEnterpriseWiFiConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identityCertificateForClientAuthentication"]; ok {
		impl, err := UnmarshalIosCertificateProfileBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'IdentityCertificateForClientAuthentication' for 'IosEnterpriseWiFiConfiguration': %+v", err)
		}
		s.IdentityCertificateForClientAuthentication = &impl
	}

	return nil
}
