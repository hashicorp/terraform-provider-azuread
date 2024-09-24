package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EasEmailProfileConfigurationBase = IosEasEmailProfileConfiguration{}

type IosEasEmailProfileConfiguration struct {
	// Account name.
	AccountName *string `json:"accountName,omitempty"`

	// Authentication method for this Email profile. Possible values are: usernameAndPassword, certificate,
	// derivedCredential.
	AuthenticationMethod *EasAuthenticationMethod `json:"authenticationMethod,omitempty"`

	// Indicates whether or not to block moving messages to other email accounts.
	BlockMovingMessagesToOtherEmailAccounts nullable.Type[bool] `json:"blockMovingMessagesToOtherEmailAccounts,omitempty"`

	// Indicates whether or not to block sending email from third party apps.
	BlockSendingEmailFromThirdPartyApps nullable.Type[bool] `json:"blockSendingEmailFromThirdPartyApps,omitempty"`

	// Indicates whether or not to block syncing recently used email addresses, for instance - when composing new email.
	BlockSyncingRecentlyUsedEmailAddresses nullable.Type[bool] `json:"blockSyncingRecentlyUsedEmailAddresses,omitempty"`

	// Tenant level settings for the Derived Credentials to be used for authentication.
	DerivedCredentialSettings *DeviceManagementDerivedCredentialSettings `json:"derivedCredentialSettings,omitempty"`

	// Possible values for email sync duration.
	DurationOfEmailToSync *EmailSyncDuration `json:"durationOfEmailToSync,omitempty"`

	// Exchange data to sync. Possible values are: none, calendars, contacts, email, notes, reminders.
	EasServices *EasServices `json:"easServices,omitempty"`

	// Allow users to change sync settings.
	EasServicesUserOverrideEnabled nullable.Type[bool] `json:"easServicesUserOverrideEnabled,omitempty"`

	// Possible values for username source or email source.
	EmailAddressSource *UserEmailSource `json:"emailAddressSource,omitempty"`

	// Encryption Certificate type for this Email profile. Possible values are: none, certificate, derivedCredential.
	EncryptionCertificateType *EmailCertificateType `json:"encryptionCertificateType,omitempty"`

	// Exchange location that (URL) that the native mail app connects to.
	HostName *string `json:"hostName,omitempty"`

	// Identity certificate.
	IdentityCertificate *IosCertificateProfileBase `json:"identityCertificate,omitempty"`

	// Profile ID of the Per-App VPN policy to be used to access emails from the native Mail client
	PerAppVPNProfileId nullable.Type[string] `json:"perAppVPNProfileId,omitempty"`

	// Indicates whether or not to use S/MIME certificate.
	RequireSmime nullable.Type[bool] `json:"requireSmime,omitempty"`

	// Indicates whether or not to use SSL.
	RequireSsl *bool `json:"requireSsl,omitempty"`

	// Signing Certificate type for this Email profile. Possible values are: none, certificate, derivedCredential.
	SigningCertificateType *EmailCertificateType `json:"signingCertificateType,omitempty"`

	// Indicates whether or not to allow unencrypted emails.
	SmimeEnablePerMessageSwitch nullable.Type[bool] `json:"smimeEnablePerMessageSwitch,omitempty"`

	// If set to true S/MIME encryption is enabled by default.
	SmimeEncryptByDefaultEnabled nullable.Type[bool] `json:"smimeEncryptByDefaultEnabled,omitempty"`

	// If set to true, the user can toggle the encryption by default setting.
	SmimeEncryptByDefaultUserOverrideEnabled nullable.Type[bool] `json:"smimeEncryptByDefaultUserOverrideEnabled,omitempty"`

	// S/MIME encryption certificate.
	SmimeEncryptionCertificate *IosCertificateProfile `json:"smimeEncryptionCertificate,omitempty"`

	// If set to true the user can select the S/MIME encryption identity.
	SmimeEncryptionCertificateUserOverrideEnabled nullable.Type[bool] `json:"smimeEncryptionCertificateUserOverrideEnabled,omitempty"`

	// S/MIME signing certificate.
	SmimeSigningCertificate *IosCertificateProfile `json:"smimeSigningCertificate,omitempty"`

	// If set to true, the user can select the signing identity.
	SmimeSigningCertificateUserOverrideEnabled nullable.Type[bool] `json:"smimeSigningCertificateUserOverrideEnabled,omitempty"`

	// If set to true S/MIME signing is enabled for this account
	SmimeSigningEnabled nullable.Type[bool] `json:"smimeSigningEnabled,omitempty"`

	// If set to true, the user can toggle S/MIME signing on or off.
	SmimeSigningUserOverrideEnabled nullable.Type[bool] `json:"smimeSigningUserOverrideEnabled,omitempty"`

	// Specifies whether the connection should use OAuth for authentication.
	UseOAuth nullable.Type[bool] `json:"useOAuth,omitempty"`

	// Fields inherited from EasEmailProfileConfigurationBase

	// Custom domain name value used while generating an email profile before installing on the device.
	CustomDomainName nullable.Type[string] `json:"customDomainName,omitempty"`

	// UserDomainname attribute that is picked from AAD and injected into this profile before installing on the device.
	// Possible values are: fullDomainName, netBiosDomainName.
	UserDomainNameSource *DomainNameSource `json:"userDomainNameSource,omitempty"`

	// Name of the AAD field, that will be used to retrieve UserName for email profile. Possible values are:
	// userPrincipalName, primarySmtpAddress, samAccountName.
	UsernameAADSource *UsernameSource `json:"usernameAADSource,omitempty"`

	// Possible values for username source or email source.
	UsernameSource *UserEmailSource `json:"usernameSource,omitempty"`

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

func (s IosEasEmailProfileConfiguration) EasEmailProfileConfigurationBase() BaseEasEmailProfileConfigurationBaseImpl {
	return BaseEasEmailProfileConfigurationBaseImpl{
		CustomDomainName:     s.CustomDomainName,
		UserDomainNameSource: s.UserDomainNameSource,
		UsernameAADSource:    s.UsernameAADSource,
		UsernameSource:       s.UsernameSource,
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

func (s IosEasEmailProfileConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s IosEasEmailProfileConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IosEasEmailProfileConfiguration{}

func (s IosEasEmailProfileConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper IosEasEmailProfileConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosEasEmailProfileConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosEasEmailProfileConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iosEasEmailProfileConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosEasEmailProfileConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IosEasEmailProfileConfiguration{}

func (s *IosEasEmailProfileConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AccountName                                   *string                                      `json:"accountName,omitempty"`
		AuthenticationMethod                          *EasAuthenticationMethod                     `json:"authenticationMethod,omitempty"`
		BlockMovingMessagesToOtherEmailAccounts       nullable.Type[bool]                          `json:"blockMovingMessagesToOtherEmailAccounts,omitempty"`
		BlockSendingEmailFromThirdPartyApps           nullable.Type[bool]                          `json:"blockSendingEmailFromThirdPartyApps,omitempty"`
		BlockSyncingRecentlyUsedEmailAddresses        nullable.Type[bool]                          `json:"blockSyncingRecentlyUsedEmailAddresses,omitempty"`
		DerivedCredentialSettings                     *DeviceManagementDerivedCredentialSettings   `json:"derivedCredentialSettings,omitempty"`
		DurationOfEmailToSync                         *EmailSyncDuration                           `json:"durationOfEmailToSync,omitempty"`
		EasServices                                   *EasServices                                 `json:"easServices,omitempty"`
		EasServicesUserOverrideEnabled                nullable.Type[bool]                          `json:"easServicesUserOverrideEnabled,omitempty"`
		EmailAddressSource                            *UserEmailSource                             `json:"emailAddressSource,omitempty"`
		EncryptionCertificateType                     *EmailCertificateType                        `json:"encryptionCertificateType,omitempty"`
		HostName                                      *string                                      `json:"hostName,omitempty"`
		PerAppVPNProfileId                            nullable.Type[string]                        `json:"perAppVPNProfileId,omitempty"`
		RequireSmime                                  nullable.Type[bool]                          `json:"requireSmime,omitempty"`
		RequireSsl                                    *bool                                        `json:"requireSsl,omitempty"`
		SigningCertificateType                        *EmailCertificateType                        `json:"signingCertificateType,omitempty"`
		SmimeEnablePerMessageSwitch                   nullable.Type[bool]                          `json:"smimeEnablePerMessageSwitch,omitempty"`
		SmimeEncryptByDefaultEnabled                  nullable.Type[bool]                          `json:"smimeEncryptByDefaultEnabled,omitempty"`
		SmimeEncryptByDefaultUserOverrideEnabled      nullable.Type[bool]                          `json:"smimeEncryptByDefaultUserOverrideEnabled,omitempty"`
		SmimeEncryptionCertificateUserOverrideEnabled nullable.Type[bool]                          `json:"smimeEncryptionCertificateUserOverrideEnabled,omitempty"`
		SmimeSigningCertificateUserOverrideEnabled    nullable.Type[bool]                          `json:"smimeSigningCertificateUserOverrideEnabled,omitempty"`
		SmimeSigningEnabled                           nullable.Type[bool]                          `json:"smimeSigningEnabled,omitempty"`
		SmimeSigningUserOverrideEnabled               nullable.Type[bool]                          `json:"smimeSigningUserOverrideEnabled,omitempty"`
		UseOAuth                                      nullable.Type[bool]                          `json:"useOAuth,omitempty"`
		CustomDomainName                              nullable.Type[string]                        `json:"customDomainName,omitempty"`
		UserDomainNameSource                          *DomainNameSource                            `json:"userDomainNameSource,omitempty"`
		UsernameAADSource                             *UsernameSource                              `json:"usernameAADSource,omitempty"`
		UsernameSource                                *UserEmailSource                             `json:"usernameSource,omitempty"`
		Assignments                                   *[]DeviceConfigurationAssignment             `json:"assignments,omitempty"`
		CreatedDateTime                               *string                                      `json:"createdDateTime,omitempty"`
		Description                                   nullable.Type[string]                        `json:"description,omitempty"`
		DeviceManagementApplicabilityRuleDeviceMode   *DeviceManagementApplicabilityRuleDeviceMode `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
		DeviceManagementApplicabilityRuleOsEdition    *DeviceManagementApplicabilityRuleOsEdition  `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
		DeviceManagementApplicabilityRuleOsVersion    *DeviceManagementApplicabilityRuleOsVersion  `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
		DeviceSettingStateSummaries                   *[]SettingStateDeviceSummary                 `json:"deviceSettingStateSummaries,omitempty"`
		DeviceStatusOverview                          *DeviceConfigurationDeviceOverview           `json:"deviceStatusOverview,omitempty"`
		DeviceStatuses                                *[]DeviceConfigurationDeviceStatus           `json:"deviceStatuses,omitempty"`
		DisplayName                                   *string                                      `json:"displayName,omitempty"`
		GroupAssignments                              *[]DeviceConfigurationGroupAssignment        `json:"groupAssignments,omitempty"`
		LastModifiedDateTime                          *string                                      `json:"lastModifiedDateTime,omitempty"`
		RoleScopeTagIds                               *[]string                                    `json:"roleScopeTagIds,omitempty"`
		SupportsScopeTags                             *bool                                        `json:"supportsScopeTags,omitempty"`
		UserStatusOverview                            *DeviceConfigurationUserOverview             `json:"userStatusOverview,omitempty"`
		UserStatuses                                  *[]DeviceConfigurationUserStatus             `json:"userStatuses,omitempty"`
		Version                                       *int64                                       `json:"version,omitempty"`
		Id                                            *string                                      `json:"id,omitempty"`
		ODataId                                       *string                                      `json:"@odata.id,omitempty"`
		ODataType                                     *string                                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccountName = decoded.AccountName
	s.AuthenticationMethod = decoded.AuthenticationMethod
	s.BlockMovingMessagesToOtherEmailAccounts = decoded.BlockMovingMessagesToOtherEmailAccounts
	s.BlockSendingEmailFromThirdPartyApps = decoded.BlockSendingEmailFromThirdPartyApps
	s.BlockSyncingRecentlyUsedEmailAddresses = decoded.BlockSyncingRecentlyUsedEmailAddresses
	s.DerivedCredentialSettings = decoded.DerivedCredentialSettings
	s.DurationOfEmailToSync = decoded.DurationOfEmailToSync
	s.EasServices = decoded.EasServices
	s.EasServicesUserOverrideEnabled = decoded.EasServicesUserOverrideEnabled
	s.EmailAddressSource = decoded.EmailAddressSource
	s.EncryptionCertificateType = decoded.EncryptionCertificateType
	s.HostName = decoded.HostName
	s.PerAppVPNProfileId = decoded.PerAppVPNProfileId
	s.RequireSmime = decoded.RequireSmime
	s.RequireSsl = decoded.RequireSsl
	s.SigningCertificateType = decoded.SigningCertificateType
	s.SmimeEnablePerMessageSwitch = decoded.SmimeEnablePerMessageSwitch
	s.SmimeEncryptByDefaultEnabled = decoded.SmimeEncryptByDefaultEnabled
	s.SmimeEncryptByDefaultUserOverrideEnabled = decoded.SmimeEncryptByDefaultUserOverrideEnabled
	s.SmimeEncryptionCertificateUserOverrideEnabled = decoded.SmimeEncryptionCertificateUserOverrideEnabled
	s.SmimeSigningCertificateUserOverrideEnabled = decoded.SmimeSigningCertificateUserOverrideEnabled
	s.SmimeSigningEnabled = decoded.SmimeSigningEnabled
	s.SmimeSigningUserOverrideEnabled = decoded.SmimeSigningUserOverrideEnabled
	s.UseOAuth = decoded.UseOAuth
	s.Assignments = decoded.Assignments
	s.CreatedDateTime = decoded.CreatedDateTime
	s.CustomDomainName = decoded.CustomDomainName
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
	s.UserDomainNameSource = decoded.UserDomainNameSource
	s.UserStatusOverview = decoded.UserStatusOverview
	s.UserStatuses = decoded.UserStatuses
	s.UsernameAADSource = decoded.UsernameAADSource
	s.UsernameSource = decoded.UsernameSource
	s.Version = decoded.Version

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IosEasEmailProfileConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identityCertificate"]; ok {
		impl, err := UnmarshalIosCertificateProfileBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'IdentityCertificate' for 'IosEasEmailProfileConfiguration': %+v", err)
		}
		s.IdentityCertificate = &impl
	}

	if v, ok := temp["smimeEncryptionCertificate"]; ok {
		impl, err := UnmarshalIosCertificateProfileImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SmimeEncryptionCertificate' for 'IosEasEmailProfileConfiguration': %+v", err)
		}
		s.SmimeEncryptionCertificate = &impl
	}

	if v, ok := temp["smimeSigningCertificate"]; ok {
		impl, err := UnmarshalIosCertificateProfileImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SmimeSigningCertificate' for 'IosEasEmailProfileConfiguration': %+v", err)
		}
		s.SmimeSigningCertificate = &impl
	}

	return nil
}
