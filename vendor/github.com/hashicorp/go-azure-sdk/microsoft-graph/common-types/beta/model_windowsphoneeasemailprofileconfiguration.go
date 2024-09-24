package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EasEmailProfileConfigurationBase = WindowsPhoneEASEmailProfileConfiguration{}

type WindowsPhoneEASEmailProfileConfiguration struct {
	// Account name.
	AccountName *string `json:"accountName,omitempty"`

	// Value indicating whether this policy only applies to Windows 8.1. This property is read-only.
	ApplyOnlyToWindowsPhone81 *bool `json:"applyOnlyToWindowsPhone81,omitempty"`

	// Possible values for email sync duration.
	DurationOfEmailToSync *EmailSyncDuration `json:"durationOfEmailToSync,omitempty"`

	// Email attribute that is picked from AAD and injected into this profile before installing on the device. Possible
	// values are: userPrincipalName, primarySmtpAddress.
	EmailAddressSource *UserEmailSource `json:"emailAddressSource,omitempty"`

	// Possible values for email sync schedule.
	EmailSyncSchedule *EmailSyncSchedule `json:"emailSyncSchedule,omitempty"`

	// Exchange location that (URL) that the native mail app connects to.
	HostName *string `json:"hostName,omitempty"`

	// Indicates whether or not to use SSL.
	RequireSsl *bool `json:"requireSsl,omitempty"`

	// Whether or not to sync the calendar.
	SyncCalendar *bool `json:"syncCalendar,omitempty"`

	// Whether or not to sync contacts.
	SyncContacts *bool `json:"syncContacts,omitempty"`

	// Whether or not to sync tasks.
	SyncTasks *bool `json:"syncTasks,omitempty"`

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

func (s WindowsPhoneEASEmailProfileConfiguration) EasEmailProfileConfigurationBase() BaseEasEmailProfileConfigurationBaseImpl {
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

func (s WindowsPhoneEASEmailProfileConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s WindowsPhoneEASEmailProfileConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsPhoneEASEmailProfileConfiguration{}

func (s WindowsPhoneEASEmailProfileConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper WindowsPhoneEASEmailProfileConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsPhoneEASEmailProfileConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsPhoneEASEmailProfileConfiguration: %+v", err)
	}

	delete(decoded, "applyOnlyToWindowsPhone81")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsPhoneEASEmailProfileConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsPhoneEASEmailProfileConfiguration: %+v", err)
	}

	return encoded, nil
}
