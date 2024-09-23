package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceConfiguration = AndroidEasEmailProfileConfiguration{}

type AndroidEasEmailProfileConfiguration struct {
	// Exchange ActiveSync account name, displayed to users as name of EAS (this) profile.
	AccountName *string `json:"accountName,omitempty"`

	// Exchange Active Sync authentication method.
	AuthenticationMethod *EasAuthenticationMethod `json:"authenticationMethod,omitempty"`

	// Custom domain name value used while generating an email profile before installing on the device.
	CustomDomainName nullable.Type[string] `json:"customDomainName,omitempty"`

	// Possible values for email sync duration.
	DurationOfEmailToSync *EmailSyncDuration `json:"durationOfEmailToSync,omitempty"`

	// Possible values for username source or email source.
	EmailAddressSource *UserEmailSource `json:"emailAddressSource,omitempty"`

	// Possible values for email sync schedule.
	EmailSyncSchedule *EmailSyncSchedule `json:"emailSyncSchedule,omitempty"`

	// Exchange location (URL) that the native mail app connects to.
	HostName *string `json:"hostName,omitempty"`

	// Identity certificate.
	IdentityCertificate *AndroidCertificateProfileBase `json:"identityCertificate,omitempty"`

	// Indicates whether or not to use S/MIME certificate.
	RequireSmime *bool `json:"requireSmime,omitempty"`

	// Indicates whether or not to use SSL.
	RequireSsl *bool `json:"requireSsl,omitempty"`

	// S/MIME signing certificate.
	SmimeSigningCertificate *AndroidCertificateProfileBase `json:"smimeSigningCertificate,omitempty"`

	// Toggles syncing the calendar. If set to false calendar is turned off on the device.
	SyncCalendar *bool `json:"syncCalendar,omitempty"`

	// Toggles syncing contacts. If set to false contacts are turned off on the device.
	SyncContacts *bool `json:"syncContacts,omitempty"`

	// Toggles syncing notes. If set to false notes are turned off on the device.
	SyncNotes *bool `json:"syncNotes,omitempty"`

	// Toggles syncing tasks. If set to false tasks are turned off on the device.
	SyncTasks *bool `json:"syncTasks,omitempty"`

	// UserDomainname attribute that is picked from AAD and injected into this profile before installing on the device.
	// Possible values are: fullDomainName, netBiosDomainName.
	UserDomainNameSource *DomainNameSource `json:"userDomainNameSource,omitempty"`

	// Android username source.
	UsernameSource *AndroidUsernameSource `json:"usernameSource,omitempty"`

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

func (s AndroidEasEmailProfileConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s AndroidEasEmailProfileConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidEasEmailProfileConfiguration{}

func (s AndroidEasEmailProfileConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper AndroidEasEmailProfileConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidEasEmailProfileConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidEasEmailProfileConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidEasEmailProfileConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidEasEmailProfileConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AndroidEasEmailProfileConfiguration{}

func (s *AndroidEasEmailProfileConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AccountName                                 *string                                      `json:"accountName,omitempty"`
		AuthenticationMethod                        *EasAuthenticationMethod                     `json:"authenticationMethod,omitempty"`
		CustomDomainName                            nullable.Type[string]                        `json:"customDomainName,omitempty"`
		DurationOfEmailToSync                       *EmailSyncDuration                           `json:"durationOfEmailToSync,omitempty"`
		EmailAddressSource                          *UserEmailSource                             `json:"emailAddressSource,omitempty"`
		EmailSyncSchedule                           *EmailSyncSchedule                           `json:"emailSyncSchedule,omitempty"`
		HostName                                    *string                                      `json:"hostName,omitempty"`
		RequireSmime                                *bool                                        `json:"requireSmime,omitempty"`
		RequireSsl                                  *bool                                        `json:"requireSsl,omitempty"`
		SyncCalendar                                *bool                                        `json:"syncCalendar,omitempty"`
		SyncContacts                                *bool                                        `json:"syncContacts,omitempty"`
		SyncNotes                                   *bool                                        `json:"syncNotes,omitempty"`
		SyncTasks                                   *bool                                        `json:"syncTasks,omitempty"`
		UserDomainNameSource                        *DomainNameSource                            `json:"userDomainNameSource,omitempty"`
		UsernameSource                              *AndroidUsernameSource                       `json:"usernameSource,omitempty"`
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

	s.AccountName = decoded.AccountName
	s.AuthenticationMethod = decoded.AuthenticationMethod
	s.CustomDomainName = decoded.CustomDomainName
	s.DurationOfEmailToSync = decoded.DurationOfEmailToSync
	s.EmailAddressSource = decoded.EmailAddressSource
	s.EmailSyncSchedule = decoded.EmailSyncSchedule
	s.HostName = decoded.HostName
	s.RequireSmime = decoded.RequireSmime
	s.RequireSsl = decoded.RequireSsl
	s.SyncCalendar = decoded.SyncCalendar
	s.SyncContacts = decoded.SyncContacts
	s.SyncNotes = decoded.SyncNotes
	s.SyncTasks = decoded.SyncTasks
	s.UserDomainNameSource = decoded.UserDomainNameSource
	s.UsernameSource = decoded.UsernameSource
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
		return fmt.Errorf("unmarshaling AndroidEasEmailProfileConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identityCertificate"]; ok {
		impl, err := UnmarshalAndroidCertificateProfileBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'IdentityCertificate' for 'AndroidEasEmailProfileConfiguration': %+v", err)
		}
		s.IdentityCertificate = &impl
	}

	if v, ok := temp["smimeSigningCertificate"]; ok {
		impl, err := UnmarshalAndroidCertificateProfileBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SmimeSigningCertificate' for 'AndroidEasEmailProfileConfiguration': %+v", err)
		}
		s.SmimeSigningCertificate = &impl
	}

	return nil
}
