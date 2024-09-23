package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsWifiConfiguration = WindowsWifiEnterpriseEAPConfiguration{}

type WindowsWifiEnterpriseEAPConfiguration struct {
	// Specify the authentication method. Possible values are: certificate, usernameAndPassword, derivedCredential.
	AuthenticationMethod *WiFiAuthenticationMethod `json:"authenticationMethod,omitempty"`

	// Specify the number of seconds for the client to wait after an authentication attempt before failing. Valid range
	// 1-3600.
	AuthenticationPeriodInSeconds nullable.Type[int64] `json:"authenticationPeriodInSeconds,omitempty"`

	// Specify the number of seconds between a failed authentication and the next authentication attempt. Valid range
	// 1-3600.
	AuthenticationRetryDelayPeriodInSeconds nullable.Type[int64] `json:"authenticationRetryDelayPeriodInSeconds,omitempty"`

	// Specify whether to authenticate the user, the device, either, or to use guest authentication (none). If you’re
	// using certificate authentication, make sure the certificate type matches the authentication type. Possible values
	// are: none, user, machine, machineOrUser, guest.
	AuthenticationType *WifiAuthenticationType `json:"authenticationType,omitempty"`

	// Specify whether to cache user credentials on the device so that users don’t need to keep entering them each time
	// they connect.
	CacheCredentials nullable.Type[bool] `json:"cacheCredentials,omitempty"`

	// Specify whether to prevent the user from being prompted to authorize new servers for trusted certification
	// authorities when EAP type is selected as PEAP.
	DisableUserPromptForServerValidation nullable.Type[bool] `json:"disableUserPromptForServerValidation,omitempty"`

	// Extensible Authentication Protocol (EAP) configuration types.
	EapType *EapType `json:"eapType,omitempty"`

	// Specify the number of seconds to wait before sending an EAPOL (Extensible Authentication Protocol over LAN) Start
	// message. Valid range 1-3600.
	EapolStartPeriodInSeconds nullable.Type[int64] `json:"eapolStartPeriodInSeconds,omitempty"`

	// Specify whether the wifi connection should enable pairwise master key caching.
	EnablePairwiseMasterKeyCaching nullable.Type[bool] `json:"enablePairwiseMasterKeyCaching,omitempty"`

	// Specify whether pre-authentication should be enabled.
	EnablePreAuthentication nullable.Type[bool] `json:"enablePreAuthentication,omitempty"`

	// Specify identity certificate for client authentication.
	IdentityCertificateForClientAuthentication *WindowsCertificateProfileBase `json:"identityCertificateForClientAuthentication,omitempty"`

	// Specify inner authentication protocol for EAP TTLS. Possible values are: unencryptedPassword,
	// challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
	InnerAuthenticationProtocolForEAPTTLS *NonEapAuthenticationMethodForEapTtlsType `json:"innerAuthenticationProtocolForEAPTTLS,omitempty"`

	// Specify the maximum authentication failures allowed for a set of credentials. Valid range 1-100.
	MaximumAuthenticationFailures nullable.Type[int64] `json:"maximumAuthenticationFailures,omitempty"`

	// Specify maximum authentication timeout (in seconds). Valid range: 1-120
	MaximumAuthenticationTimeoutInSeconds nullable.Type[int64] `json:"maximumAuthenticationTimeoutInSeconds,omitempty"`

	// Specifiy the maximum number of EAPOL (Extensible Authentication Protocol over LAN) Start messages to be sent before
	// returning failure. Valid range 1-100.
	MaximumEAPOLStartMessages nullable.Type[int64] `json:"maximumEAPOLStartMessages,omitempty"`

	// Specify maximum number of pairwise master keys in cache. Valid range: 1-255
	MaximumNumberOfPairwiseMasterKeysInCache nullable.Type[int64] `json:"maximumNumberOfPairwiseMasterKeysInCache,omitempty"`

	// Specify maximum pairwise master key cache time (in minutes). Valid range: 5-1440
	MaximumPairwiseMasterKeyCacheTimeInMinutes nullable.Type[int64] `json:"maximumPairwiseMasterKeyCacheTimeInMinutes,omitempty"`

	// Specify maximum pre-authentication attempts. Valid range: 1-16
	MaximumPreAuthenticationAttempts nullable.Type[int64] `json:"maximumPreAuthenticationAttempts,omitempty"`

	// Specify the network single sign on type. Possible values are: disabled, prelogon, postlogon.
	NetworkSingleSignOn *NetworkSingleSignOnType `json:"networkSingleSignOn,omitempty"`

	// Specify the string to replace usernames for privacy when using EAP TTLS or PEAP.
	OuterIdentityPrivacyTemporaryValue nullable.Type[string] `json:"outerIdentityPrivacyTemporaryValue,omitempty"`

	// Specify whether to enable verification of server's identity by validating the certificate when EAP type is selected
	// as PEAP.
	PerformServerValidation nullable.Type[bool] `json:"performServerValidation,omitempty"`

	// Specify whether the wifi connection should prompt for additional authentication credentials.
	PromptForAdditionalAuthenticationCredentials nullable.Type[bool] `json:"promptForAdditionalAuthenticationCredentials,omitempty"`

	// Specify whether to enable cryptographic binding when EAP type is selected as PEAP.
	RequireCryptographicBinding nullable.Type[bool] `json:"requireCryptographicBinding,omitempty"`

	// Specify root certificate for client validation.
	RootCertificateForClientValidation *Windows81TrustedRootCertificate `json:"rootCertificateForClientValidation,omitempty"`

	// Specify root certificate for server validation. This collection can contain a maximum of 500 elements.
	RootCertificatesForServerValidation *[]Windows81TrustedRootCertificate `json:"rootCertificatesForServerValidation,omitempty"`

	// Specify trusted server certificate names.
	TrustedServerCertificateNames *[]string `json:"trustedServerCertificateNames,omitempty"`

	// Specifiy whether to change the virtual LAN used by the device based on the user’s credentials. Cannot be used when
	// NetworkSingleSignOnType is set to ​Disabled.
	UserBasedVirtualLan nullable.Type[bool] `json:"userBasedVirtualLan,omitempty"`

	// Fields inherited from WindowsWifiConfiguration

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

func (s WindowsWifiEnterpriseEAPConfiguration) WindowsWifiConfiguration() BaseWindowsWifiConfigurationImpl {
	return BaseWindowsWifiConfigurationImpl{
		ConnectAutomatically:           s.ConnectAutomatically,
		ConnectToPreferredNetwork:      s.ConnectToPreferredNetwork,
		ConnectWhenNetworkNameIsHidden: s.ConnectWhenNetworkNameIsHidden,
		ForceFIPSCompliance:            s.ForceFIPSCompliance,
		MeteredConnectionLimit:         s.MeteredConnectionLimit,
		NetworkName:                    s.NetworkName,
		PreSharedKey:                   s.PreSharedKey,
		ProxyAutomaticConfigurationUrl: s.ProxyAutomaticConfigurationUrl,
		ProxyManualAddress:             s.ProxyManualAddress,
		ProxyManualPort:                s.ProxyManualPort,
		ProxySetting:                   s.ProxySetting,
		Ssid:                           s.Ssid,
		WifiSecurityType:               s.WifiSecurityType,
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

func (s WindowsWifiEnterpriseEAPConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s WindowsWifiEnterpriseEAPConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsWifiEnterpriseEAPConfiguration{}

func (s WindowsWifiEnterpriseEAPConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper WindowsWifiEnterpriseEAPConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsWifiEnterpriseEAPConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsWifiEnterpriseEAPConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsWifiEnterpriseEAPConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsWifiEnterpriseEAPConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsWifiEnterpriseEAPConfiguration{}

func (s *WindowsWifiEnterpriseEAPConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AuthenticationMethod                         *WiFiAuthenticationMethod                    `json:"authenticationMethod,omitempty"`
		AuthenticationPeriodInSeconds                nullable.Type[int64]                         `json:"authenticationPeriodInSeconds,omitempty"`
		AuthenticationRetryDelayPeriodInSeconds      nullable.Type[int64]                         `json:"authenticationRetryDelayPeriodInSeconds,omitempty"`
		AuthenticationType                           *WifiAuthenticationType                      `json:"authenticationType,omitempty"`
		CacheCredentials                             nullable.Type[bool]                          `json:"cacheCredentials,omitempty"`
		DisableUserPromptForServerValidation         nullable.Type[bool]                          `json:"disableUserPromptForServerValidation,omitempty"`
		EapType                                      *EapType                                     `json:"eapType,omitempty"`
		EapolStartPeriodInSeconds                    nullable.Type[int64]                         `json:"eapolStartPeriodInSeconds,omitempty"`
		EnablePairwiseMasterKeyCaching               nullable.Type[bool]                          `json:"enablePairwiseMasterKeyCaching,omitempty"`
		EnablePreAuthentication                      nullable.Type[bool]                          `json:"enablePreAuthentication,omitempty"`
		InnerAuthenticationProtocolForEAPTTLS        *NonEapAuthenticationMethodForEapTtlsType    `json:"innerAuthenticationProtocolForEAPTTLS,omitempty"`
		MaximumAuthenticationFailures                nullable.Type[int64]                         `json:"maximumAuthenticationFailures,omitempty"`
		MaximumAuthenticationTimeoutInSeconds        nullable.Type[int64]                         `json:"maximumAuthenticationTimeoutInSeconds,omitempty"`
		MaximumEAPOLStartMessages                    nullable.Type[int64]                         `json:"maximumEAPOLStartMessages,omitempty"`
		MaximumNumberOfPairwiseMasterKeysInCache     nullable.Type[int64]                         `json:"maximumNumberOfPairwiseMasterKeysInCache,omitempty"`
		MaximumPairwiseMasterKeyCacheTimeInMinutes   nullable.Type[int64]                         `json:"maximumPairwiseMasterKeyCacheTimeInMinutes,omitempty"`
		MaximumPreAuthenticationAttempts             nullable.Type[int64]                         `json:"maximumPreAuthenticationAttempts,omitempty"`
		NetworkSingleSignOn                          *NetworkSingleSignOnType                     `json:"networkSingleSignOn,omitempty"`
		OuterIdentityPrivacyTemporaryValue           nullable.Type[string]                        `json:"outerIdentityPrivacyTemporaryValue,omitempty"`
		PerformServerValidation                      nullable.Type[bool]                          `json:"performServerValidation,omitempty"`
		PromptForAdditionalAuthenticationCredentials nullable.Type[bool]                          `json:"promptForAdditionalAuthenticationCredentials,omitempty"`
		RequireCryptographicBinding                  nullable.Type[bool]                          `json:"requireCryptographicBinding,omitempty"`
		RootCertificateForClientValidation           *Windows81TrustedRootCertificate             `json:"rootCertificateForClientValidation,omitempty"`
		RootCertificatesForServerValidation          *[]Windows81TrustedRootCertificate           `json:"rootCertificatesForServerValidation,omitempty"`
		TrustedServerCertificateNames                *[]string                                    `json:"trustedServerCertificateNames,omitempty"`
		UserBasedVirtualLan                          nullable.Type[bool]                          `json:"userBasedVirtualLan,omitempty"`
		ConnectAutomatically                         nullable.Type[bool]                          `json:"connectAutomatically,omitempty"`
		ConnectToPreferredNetwork                    nullable.Type[bool]                          `json:"connectToPreferredNetwork,omitempty"`
		ConnectWhenNetworkNameIsHidden               nullable.Type[bool]                          `json:"connectWhenNetworkNameIsHidden,omitempty"`
		ForceFIPSCompliance                          nullable.Type[bool]                          `json:"forceFIPSCompliance,omitempty"`
		MeteredConnectionLimit                       *MeteredConnectionLimitType                  `json:"meteredConnectionLimit,omitempty"`
		NetworkName                                  nullable.Type[string]                        `json:"networkName,omitempty"`
		PreSharedKey                                 nullable.Type[string]                        `json:"preSharedKey,omitempty"`
		ProxyAutomaticConfigurationUrl               nullable.Type[string]                        `json:"proxyAutomaticConfigurationUrl,omitempty"`
		ProxyManualAddress                           nullable.Type[string]                        `json:"proxyManualAddress,omitempty"`
		ProxyManualPort                              nullable.Type[int64]                         `json:"proxyManualPort,omitempty"`
		ProxySetting                                 *WiFiProxySetting                            `json:"proxySetting,omitempty"`
		Ssid                                         nullable.Type[string]                        `json:"ssid,omitempty"`
		WifiSecurityType                             *WiFiSecurityType                            `json:"wifiSecurityType,omitempty"`
		Assignments                                  *[]DeviceConfigurationAssignment             `json:"assignments,omitempty"`
		CreatedDateTime                              *string                                      `json:"createdDateTime,omitempty"`
		Description                                  nullable.Type[string]                        `json:"description,omitempty"`
		DeviceManagementApplicabilityRuleDeviceMode  *DeviceManagementApplicabilityRuleDeviceMode `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
		DeviceManagementApplicabilityRuleOsEdition   *DeviceManagementApplicabilityRuleOsEdition  `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
		DeviceManagementApplicabilityRuleOsVersion   *DeviceManagementApplicabilityRuleOsVersion  `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
		DeviceSettingStateSummaries                  *[]SettingStateDeviceSummary                 `json:"deviceSettingStateSummaries,omitempty"`
		DeviceStatusOverview                         *DeviceConfigurationDeviceOverview           `json:"deviceStatusOverview,omitempty"`
		DeviceStatuses                               *[]DeviceConfigurationDeviceStatus           `json:"deviceStatuses,omitempty"`
		DisplayName                                  *string                                      `json:"displayName,omitempty"`
		GroupAssignments                             *[]DeviceConfigurationGroupAssignment        `json:"groupAssignments,omitempty"`
		LastModifiedDateTime                         *string                                      `json:"lastModifiedDateTime,omitempty"`
		RoleScopeTagIds                              *[]string                                    `json:"roleScopeTagIds,omitempty"`
		SupportsScopeTags                            *bool                                        `json:"supportsScopeTags,omitempty"`
		UserStatusOverview                           *DeviceConfigurationUserOverview             `json:"userStatusOverview,omitempty"`
		UserStatuses                                 *[]DeviceConfigurationUserStatus             `json:"userStatuses,omitempty"`
		Version                                      *int64                                       `json:"version,omitempty"`
		Id                                           *string                                      `json:"id,omitempty"`
		ODataId                                      *string                                      `json:"@odata.id,omitempty"`
		ODataType                                    *string                                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AuthenticationMethod = decoded.AuthenticationMethod
	s.AuthenticationPeriodInSeconds = decoded.AuthenticationPeriodInSeconds
	s.AuthenticationRetryDelayPeriodInSeconds = decoded.AuthenticationRetryDelayPeriodInSeconds
	s.AuthenticationType = decoded.AuthenticationType
	s.CacheCredentials = decoded.CacheCredentials
	s.DisableUserPromptForServerValidation = decoded.DisableUserPromptForServerValidation
	s.EapType = decoded.EapType
	s.EapolStartPeriodInSeconds = decoded.EapolStartPeriodInSeconds
	s.EnablePairwiseMasterKeyCaching = decoded.EnablePairwiseMasterKeyCaching
	s.EnablePreAuthentication = decoded.EnablePreAuthentication
	s.InnerAuthenticationProtocolForEAPTTLS = decoded.InnerAuthenticationProtocolForEAPTTLS
	s.MaximumAuthenticationFailures = decoded.MaximumAuthenticationFailures
	s.MaximumAuthenticationTimeoutInSeconds = decoded.MaximumAuthenticationTimeoutInSeconds
	s.MaximumEAPOLStartMessages = decoded.MaximumEAPOLStartMessages
	s.MaximumNumberOfPairwiseMasterKeysInCache = decoded.MaximumNumberOfPairwiseMasterKeysInCache
	s.MaximumPairwiseMasterKeyCacheTimeInMinutes = decoded.MaximumPairwiseMasterKeyCacheTimeInMinutes
	s.MaximumPreAuthenticationAttempts = decoded.MaximumPreAuthenticationAttempts
	s.NetworkSingleSignOn = decoded.NetworkSingleSignOn
	s.OuterIdentityPrivacyTemporaryValue = decoded.OuterIdentityPrivacyTemporaryValue
	s.PerformServerValidation = decoded.PerformServerValidation
	s.PromptForAdditionalAuthenticationCredentials = decoded.PromptForAdditionalAuthenticationCredentials
	s.RequireCryptographicBinding = decoded.RequireCryptographicBinding
	s.RootCertificateForClientValidation = decoded.RootCertificateForClientValidation
	s.RootCertificatesForServerValidation = decoded.RootCertificatesForServerValidation
	s.TrustedServerCertificateNames = decoded.TrustedServerCertificateNames
	s.UserBasedVirtualLan = decoded.UserBasedVirtualLan
	s.Assignments = decoded.Assignments
	s.ConnectAutomatically = decoded.ConnectAutomatically
	s.ConnectToPreferredNetwork = decoded.ConnectToPreferredNetwork
	s.ConnectWhenNetworkNameIsHidden = decoded.ConnectWhenNetworkNameIsHidden
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DeviceManagementApplicabilityRuleDeviceMode = decoded.DeviceManagementApplicabilityRuleDeviceMode
	s.DeviceManagementApplicabilityRuleOsEdition = decoded.DeviceManagementApplicabilityRuleOsEdition
	s.DeviceManagementApplicabilityRuleOsVersion = decoded.DeviceManagementApplicabilityRuleOsVersion
	s.DeviceSettingStateSummaries = decoded.DeviceSettingStateSummaries
	s.DeviceStatusOverview = decoded.DeviceStatusOverview
	s.DeviceStatuses = decoded.DeviceStatuses
	s.DisplayName = decoded.DisplayName
	s.ForceFIPSCompliance = decoded.ForceFIPSCompliance
	s.GroupAssignments = decoded.GroupAssignments
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.MeteredConnectionLimit = decoded.MeteredConnectionLimit
	s.NetworkName = decoded.NetworkName
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PreSharedKey = decoded.PreSharedKey
	s.ProxyAutomaticConfigurationUrl = decoded.ProxyAutomaticConfigurationUrl
	s.ProxyManualAddress = decoded.ProxyManualAddress
	s.ProxyManualPort = decoded.ProxyManualPort
	s.ProxySetting = decoded.ProxySetting
	s.RoleScopeTagIds = decoded.RoleScopeTagIds
	s.Ssid = decoded.Ssid
	s.SupportsScopeTags = decoded.SupportsScopeTags
	s.UserStatusOverview = decoded.UserStatusOverview
	s.UserStatuses = decoded.UserStatuses
	s.Version = decoded.Version
	s.WifiSecurityType = decoded.WifiSecurityType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsWifiEnterpriseEAPConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identityCertificateForClientAuthentication"]; ok {
		impl, err := UnmarshalWindowsCertificateProfileBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'IdentityCertificateForClientAuthentication' for 'WindowsWifiEnterpriseEAPConfiguration': %+v", err)
		}
		s.IdentityCertificateForClientAuthentication = &impl
	}

	return nil
}
