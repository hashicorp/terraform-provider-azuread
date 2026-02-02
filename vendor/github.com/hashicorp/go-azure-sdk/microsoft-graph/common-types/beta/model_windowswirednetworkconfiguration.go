package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceConfiguration = WindowsWiredNetworkConfiguration{}

type WindowsWiredNetworkConfiguration struct {
	// Specify the duration for which automatic authentication attempts will be blocked from occuring after a failed
	// authentication attempt.
	AuthenticationBlockPeriodInMinutes nullable.Type[int64] `json:"authenticationBlockPeriodInMinutes,omitempty"`

	// Specify the authentication method. Possible values are: certificate, usernameAndPassword, derivedCredential. Possible
	// values are: certificate, usernameAndPassword, derivedCredential, unknownFutureValue.
	AuthenticationMethod *WiredNetworkAuthenticationMethod `json:"authenticationMethod,omitempty"`

	// Specify the number of seconds for the client to wait after an authentication attempt before failing. Valid range
	// 1-3600.
	AuthenticationPeriodInSeconds nullable.Type[int64] `json:"authenticationPeriodInSeconds,omitempty"`

	// Specify the number of seconds between a failed authentication and the next authentication attempt. Valid range
	// 1-3600.
	AuthenticationRetryDelayPeriodInSeconds nullable.Type[int64] `json:"authenticationRetryDelayPeriodInSeconds,omitempty"`

	// Specify whether to authenticate the user, the device, either, or to use guest authentication (none). If you're using
	// certificate authentication, make sure the certificate type matches the authentication type. Possible values are:
	// none, user, machine, machineOrUser, guest. Possible values are: none, user, machine, machineOrUser, guest,
	// unknownFutureValue.
	AuthenticationType *WiredNetworkAuthenticationType `json:"authenticationType,omitempty"`

	// When TRUE, caches user credentials on the device so that users don't need to keep entering them each time they
	// connect. When FALSE, do not cache credentials. Default value is FALSE.
	CacheCredentials nullable.Type[bool] `json:"cacheCredentials,omitempty"`

	// When TRUE, prevents the user from being prompted to authorize new servers for trusted certification authorities when
	// EAP type is selected as PEAP. When FALSE, does not prevent the user from being prompted. Default value is FALSE.
	DisableUserPromptForServerValidation nullable.Type[bool] `json:"disableUserPromptForServerValidation,omitempty"`

	// Extensible Authentication Protocol (EAP) configuration types.
	EapType *EapType `json:"eapType,omitempty"`

	// Specify the number of seconds to wait before sending an EAPOL (Extensible Authentication Protocol over LAN) Start
	// message. Valid range 1-3600.
	EapolStartPeriodInSeconds nullable.Type[int64] `json:"eapolStartPeriodInSeconds,omitempty"`

	// When TRUE, the automatic configuration service for wired networks requires the use of 802.1X for port authentication.
	// When FALSE, 802.1X is not required. Default value is FALSE.
	Enforce8021X nullable.Type[bool] `json:"enforce8021X,omitempty"`

	// When TRUE, forces FIPS compliance. When FALSE, does not enable FIPS compliance. Default value is FALSE.
	ForceFIPSCompliance nullable.Type[bool] `json:"forceFIPSCompliance,omitempty"`

	// Specify identity certificate for client authentication.
	IdentityCertificateForClientAuthentication *WindowsCertificateProfileBase `json:"identityCertificateForClientAuthentication,omitempty"`

	// Specify inner authentication protocol for EAP TTLS. Possible values are: unencryptedPassword,
	// challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo. Possible values are:
	// unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
	InnerAuthenticationProtocolForEAPTTLS *NonEapAuthenticationMethodForEapTtlsType `json:"innerAuthenticationProtocolForEAPTTLS,omitempty"`

	// Specify the maximum authentication failures allowed for a set of credentials. Valid range 1-100.
	MaximumAuthenticationFailures nullable.Type[int64] `json:"maximumAuthenticationFailures,omitempty"`

	// Specify the maximum number of EAPOL (Extensible Authentication Protocol over LAN) Start messages to be sent before
	// returning failure. Valid range 1-100.
	MaximumEAPOLStartMessages nullable.Type[int64] `json:"maximumEAPOLStartMessages,omitempty"`

	// Specify the string to replace usernames for privacy when using EAP TTLS or PEAP.
	OuterIdentityPrivacyTemporaryValue nullable.Type[string] `json:"outerIdentityPrivacyTemporaryValue,omitempty"`

	// When TRUE, enables verification of server's identity by validating the certificate when EAP type is selected as PEAP.
	// When FALSE, the certificate is not validated. Default value is TRUE.
	PerformServerValidation nullable.Type[bool] `json:"performServerValidation,omitempty"`

	// When TRUE, enables cryptographic binding when EAP type is selected as PEAP. When FALSE, does not enable cryptogrpahic
	// binding. Default value is TRUE.
	RequireCryptographicBinding nullable.Type[bool] `json:"requireCryptographicBinding,omitempty"`

	// Specify root certificate for client validation.
	RootCertificateForClientValidation *Windows81TrustedRootCertificate `json:"rootCertificateForClientValidation,omitempty"`

	// Specify root certificates for server validation. This collection can contain a maximum of 500 elements.
	RootCertificatesForServerValidation *[]Windows81TrustedRootCertificate `json:"rootCertificatesForServerValidation,omitempty"`

	// Specify the secondary authentication method. Possible values are: certificate, usernameAndPassword,
	// derivedCredential. Possible values are: certificate, usernameAndPassword, derivedCredential, unknownFutureValue.
	SecondaryAuthenticationMethod *WiredNetworkAuthenticationMethod `json:"secondaryAuthenticationMethod,omitempty"`

	// Specify secondary identity certificate for client authentication.
	SecondaryIdentityCertificateForClientAuthentication *WindowsCertificateProfileBase `json:"secondaryIdentityCertificateForClientAuthentication,omitempty"`

	// Specify secondary root certificate for client validation.
	SecondaryRootCertificateForClientValidation *Windows81TrustedRootCertificate `json:"secondaryRootCertificateForClientValidation,omitempty"`

	// Specify trusted server certificate names.
	TrustedServerCertificateNames *[]string `json:"trustedServerCertificateNames,omitempty"`

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

func (s WindowsWiredNetworkConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s WindowsWiredNetworkConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsWiredNetworkConfiguration{}

func (s WindowsWiredNetworkConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper WindowsWiredNetworkConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsWiredNetworkConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsWiredNetworkConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsWiredNetworkConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsWiredNetworkConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsWiredNetworkConfiguration{}

func (s *WindowsWiredNetworkConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AuthenticationBlockPeriodInMinutes          nullable.Type[int64]                         `json:"authenticationBlockPeriodInMinutes,omitempty"`
		AuthenticationMethod                        *WiredNetworkAuthenticationMethod            `json:"authenticationMethod,omitempty"`
		AuthenticationPeriodInSeconds               nullable.Type[int64]                         `json:"authenticationPeriodInSeconds,omitempty"`
		AuthenticationRetryDelayPeriodInSeconds     nullable.Type[int64]                         `json:"authenticationRetryDelayPeriodInSeconds,omitempty"`
		AuthenticationType                          *WiredNetworkAuthenticationType              `json:"authenticationType,omitempty"`
		CacheCredentials                            nullable.Type[bool]                          `json:"cacheCredentials,omitempty"`
		DisableUserPromptForServerValidation        nullable.Type[bool]                          `json:"disableUserPromptForServerValidation,omitempty"`
		EapType                                     *EapType                                     `json:"eapType,omitempty"`
		EapolStartPeriodInSeconds                   nullable.Type[int64]                         `json:"eapolStartPeriodInSeconds,omitempty"`
		Enforce8021X                                nullable.Type[bool]                          `json:"enforce8021X,omitempty"`
		ForceFIPSCompliance                         nullable.Type[bool]                          `json:"forceFIPSCompliance,omitempty"`
		InnerAuthenticationProtocolForEAPTTLS       *NonEapAuthenticationMethodForEapTtlsType    `json:"innerAuthenticationProtocolForEAPTTLS,omitempty"`
		MaximumAuthenticationFailures               nullable.Type[int64]                         `json:"maximumAuthenticationFailures,omitempty"`
		MaximumEAPOLStartMessages                   nullable.Type[int64]                         `json:"maximumEAPOLStartMessages,omitempty"`
		OuterIdentityPrivacyTemporaryValue          nullable.Type[string]                        `json:"outerIdentityPrivacyTemporaryValue,omitempty"`
		PerformServerValidation                     nullable.Type[bool]                          `json:"performServerValidation,omitempty"`
		RequireCryptographicBinding                 nullable.Type[bool]                          `json:"requireCryptographicBinding,omitempty"`
		RootCertificateForClientValidation          *Windows81TrustedRootCertificate             `json:"rootCertificateForClientValidation,omitempty"`
		RootCertificatesForServerValidation         *[]Windows81TrustedRootCertificate           `json:"rootCertificatesForServerValidation,omitempty"`
		SecondaryAuthenticationMethod               *WiredNetworkAuthenticationMethod            `json:"secondaryAuthenticationMethod,omitempty"`
		SecondaryRootCertificateForClientValidation *Windows81TrustedRootCertificate             `json:"secondaryRootCertificateForClientValidation,omitempty"`
		TrustedServerCertificateNames               *[]string                                    `json:"trustedServerCertificateNames,omitempty"`
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

	s.AuthenticationBlockPeriodInMinutes = decoded.AuthenticationBlockPeriodInMinutes
	s.AuthenticationMethod = decoded.AuthenticationMethod
	s.AuthenticationPeriodInSeconds = decoded.AuthenticationPeriodInSeconds
	s.AuthenticationRetryDelayPeriodInSeconds = decoded.AuthenticationRetryDelayPeriodInSeconds
	s.AuthenticationType = decoded.AuthenticationType
	s.CacheCredentials = decoded.CacheCredentials
	s.DisableUserPromptForServerValidation = decoded.DisableUserPromptForServerValidation
	s.EapType = decoded.EapType
	s.EapolStartPeriodInSeconds = decoded.EapolStartPeriodInSeconds
	s.Enforce8021X = decoded.Enforce8021X
	s.ForceFIPSCompliance = decoded.ForceFIPSCompliance
	s.InnerAuthenticationProtocolForEAPTTLS = decoded.InnerAuthenticationProtocolForEAPTTLS
	s.MaximumAuthenticationFailures = decoded.MaximumAuthenticationFailures
	s.MaximumEAPOLStartMessages = decoded.MaximumEAPOLStartMessages
	s.OuterIdentityPrivacyTemporaryValue = decoded.OuterIdentityPrivacyTemporaryValue
	s.PerformServerValidation = decoded.PerformServerValidation
	s.RequireCryptographicBinding = decoded.RequireCryptographicBinding
	s.RootCertificateForClientValidation = decoded.RootCertificateForClientValidation
	s.RootCertificatesForServerValidation = decoded.RootCertificatesForServerValidation
	s.SecondaryAuthenticationMethod = decoded.SecondaryAuthenticationMethod
	s.SecondaryRootCertificateForClientValidation = decoded.SecondaryRootCertificateForClientValidation
	s.TrustedServerCertificateNames = decoded.TrustedServerCertificateNames
	s.Assignments = decoded.Assignments
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
	s.RoleScopeTagIds = decoded.RoleScopeTagIds
	s.SupportsScopeTags = decoded.SupportsScopeTags
	s.UserStatusOverview = decoded.UserStatusOverview
	s.UserStatuses = decoded.UserStatuses
	s.Version = decoded.Version

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsWiredNetworkConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identityCertificateForClientAuthentication"]; ok {
		impl, err := UnmarshalWindowsCertificateProfileBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'IdentityCertificateForClientAuthentication' for 'WindowsWiredNetworkConfiguration': %+v", err)
		}
		s.IdentityCertificateForClientAuthentication = &impl
	}

	if v, ok := temp["secondaryIdentityCertificateForClientAuthentication"]; ok {
		impl, err := UnmarshalWindowsCertificateProfileBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SecondaryIdentityCertificateForClientAuthentication' for 'WindowsWiredNetworkConfiguration': %+v", err)
		}
		s.SecondaryIdentityCertificateForClientAuthentication = &impl
	}

	return nil
}
