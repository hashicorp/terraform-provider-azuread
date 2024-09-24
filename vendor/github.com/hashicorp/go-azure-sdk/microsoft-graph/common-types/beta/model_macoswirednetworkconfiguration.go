package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceConfiguration = MacOSWiredNetworkConfiguration{}

type MacOSWiredNetworkConfiguration struct {
	// Authentication Method when EAP Type is configured to PEAP or EAP-TTLS. Possible values are: certificate,
	// usernameAndPassword, derivedCredential.
	AuthenticationMethod *WiFiAuthenticationMethod `json:"authenticationMethod,omitempty"`

	// Indicates the deployment channel type used to deploy the configuration profile. Possible values are deviceChannel,
	// userChannel. Possible values are: deviceChannel, userChannel, unknownFutureValue.
	DeploymentChannel *AppleDeploymentChannel `json:"deploymentChannel,omitempty"`

	// EAP-FAST Configuration Option when EAP-FAST is the selected EAP Type. Possible values are:
	// noProtectedAccessCredential, useProtectedAccessCredential, useProtectedAccessCredentialAndProvision,
	// useProtectedAccessCredentialAndProvisionAnonymously.
	EapFastConfiguration *EapFastConfiguration `json:"eapFastConfiguration,omitempty"`

	// Extensible Authentication Protocol (EAP) configuration types.
	EapType *EapType `json:"eapType,omitempty"`

	// Enable identity privacy (Outer Identity) when EAP Type is configured to EAP-TTLS, EAP-FAST or PEAP. This property
	// masks usernames with the text you enter. For example, if you use 'anonymous', each user that authenticates with this
	// wired network using their real username is displayed as 'anonymous'.
	EnableOuterIdentityPrivacy nullable.Type[string] `json:"enableOuterIdentityPrivacy,omitempty"`

	// Identity Certificate for client authentication when EAP Type is configured to EAP-TLS, EAP-TTLS (with Certificate
	// Authentication), or PEAP (with Certificate Authentication).
	IdentityCertificateForClientAuthentication *MacOSCertificateProfileBase `json:"identityCertificateForClientAuthentication,omitempty"`

	// Apple network interface type.
	NetworkInterface *WiredNetworkInterface `json:"networkInterface,omitempty"`

	// Network Name
	NetworkName *string `json:"networkName,omitempty"`

	// Non-EAP Method for Authentication (Inner Identity) when EAP Type is EAP-TTLS and Authenticationmethod is Username and
	// Password. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap,
	// microsoftChapVersionTwo.
	NonEapAuthenticationMethodForEapTtls *NonEapAuthenticationMethodForEapTtlsType `json:"nonEapAuthenticationMethodForEapTtls,omitempty"`

	// Trusted Root Certificate for Server Validation when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP.
	RootCertificateForServerValidation *MacOSTrustedRootCertificate `json:"rootCertificateForServerValidation,omitempty"`

	// Trusted server certificate names when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. This is the common name
	// used in the certificates issued by your trusted certificate authority (CA). If you provide this information, you can
	// bypass the dynamic trust dialog that is displayed on end users devices when they connect to this wired network.
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

func (s MacOSWiredNetworkConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s MacOSWiredNetworkConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MacOSWiredNetworkConfiguration{}

func (s MacOSWiredNetworkConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper MacOSWiredNetworkConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MacOSWiredNetworkConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MacOSWiredNetworkConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.macOSWiredNetworkConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MacOSWiredNetworkConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &MacOSWiredNetworkConfiguration{}

func (s *MacOSWiredNetworkConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AuthenticationMethod                        *WiFiAuthenticationMethod                    `json:"authenticationMethod,omitempty"`
		DeploymentChannel                           *AppleDeploymentChannel                      `json:"deploymentChannel,omitempty"`
		EapFastConfiguration                        *EapFastConfiguration                        `json:"eapFastConfiguration,omitempty"`
		EapType                                     *EapType                                     `json:"eapType,omitempty"`
		EnableOuterIdentityPrivacy                  nullable.Type[string]                        `json:"enableOuterIdentityPrivacy,omitempty"`
		NetworkInterface                            *WiredNetworkInterface                       `json:"networkInterface,omitempty"`
		NetworkName                                 *string                                      `json:"networkName,omitempty"`
		NonEapAuthenticationMethodForEapTtls        *NonEapAuthenticationMethodForEapTtlsType    `json:"nonEapAuthenticationMethodForEapTtls,omitempty"`
		RootCertificateForServerValidation          *MacOSTrustedRootCertificate                 `json:"rootCertificateForServerValidation,omitempty"`
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

	s.AuthenticationMethod = decoded.AuthenticationMethod
	s.DeploymentChannel = decoded.DeploymentChannel
	s.EapFastConfiguration = decoded.EapFastConfiguration
	s.EapType = decoded.EapType
	s.EnableOuterIdentityPrivacy = decoded.EnableOuterIdentityPrivacy
	s.NetworkInterface = decoded.NetworkInterface
	s.NetworkName = decoded.NetworkName
	s.NonEapAuthenticationMethodForEapTtls = decoded.NonEapAuthenticationMethodForEapTtls
	s.RootCertificateForServerValidation = decoded.RootCertificateForServerValidation
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
		return fmt.Errorf("unmarshaling MacOSWiredNetworkConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identityCertificateForClientAuthentication"]; ok {
		impl, err := UnmarshalMacOSCertificateProfileBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'IdentityCertificateForClientAuthentication' for 'MacOSWiredNetworkConfiguration': %+v", err)
		}
		s.IdentityCertificateForClientAuthentication = &impl
	}

	return nil
}
