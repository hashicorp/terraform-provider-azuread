package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AndroidWorkProfileEasEmailProfileBase = AndroidWorkProfileNineWorkEasConfiguration{}

type AndroidWorkProfileNineWorkEasConfiguration struct {
	// Toggles syncing the calendar. If set to false the calendar is turned off on the device.
	SyncCalendar *bool `json:"syncCalendar,omitempty"`

	// Toggles syncing contacts. If set to false contacts are turned off on the device.
	SyncContacts *bool `json:"syncContacts,omitempty"`

	// Toggles syncing tasks. If set to false tasks are turned off on the device.
	SyncTasks *bool `json:"syncTasks,omitempty"`

	// Fields inherited from AndroidWorkProfileEasEmailProfileBase

	// Exchange Active Sync authentication method.
	AuthenticationMethod *EasAuthenticationMethod `json:"authenticationMethod,omitempty"`

	// Possible values for email sync duration.
	DurationOfEmailToSync *EmailSyncDuration `json:"durationOfEmailToSync,omitempty"`

	// Possible values for username source or email source.
	EmailAddressSource *UserEmailSource `json:"emailAddressSource,omitempty"`

	// Exchange location (URL) that the mail app connects to.
	HostName *string `json:"hostName,omitempty"`

	// Identity certificate.
	IdentityCertificate *AndroidWorkProfileCertificateProfileBase `json:"identityCertificate,omitempty"`

	// Indicates whether or not to use SSL.
	RequireSsl *bool `json:"requireSsl,omitempty"`

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

func (s AndroidWorkProfileNineWorkEasConfiguration) AndroidWorkProfileEasEmailProfileBase() BaseAndroidWorkProfileEasEmailProfileBaseImpl {
	return BaseAndroidWorkProfileEasEmailProfileBaseImpl{
		AuthenticationMethod:  s.AuthenticationMethod,
		DurationOfEmailToSync: s.DurationOfEmailToSync,
		EmailAddressSource:    s.EmailAddressSource,
		HostName:              s.HostName,
		IdentityCertificate:   s.IdentityCertificate,
		RequireSsl:            s.RequireSsl,
		UsernameSource:        s.UsernameSource,
		Assignments:           s.Assignments,
		CreatedDateTime:       s.CreatedDateTime,
		Description:           s.Description,
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

func (s AndroidWorkProfileNineWorkEasConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s AndroidWorkProfileNineWorkEasConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidWorkProfileNineWorkEasConfiguration{}

func (s AndroidWorkProfileNineWorkEasConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper AndroidWorkProfileNineWorkEasConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidWorkProfileNineWorkEasConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidWorkProfileNineWorkEasConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidWorkProfileNineWorkEasConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidWorkProfileNineWorkEasConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AndroidWorkProfileNineWorkEasConfiguration{}

func (s *AndroidWorkProfileNineWorkEasConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		SyncCalendar                                *bool                                        `json:"syncCalendar,omitempty"`
		SyncContacts                                *bool                                        `json:"syncContacts,omitempty"`
		SyncTasks                                   *bool                                        `json:"syncTasks,omitempty"`
		AuthenticationMethod                        *EasAuthenticationMethod                     `json:"authenticationMethod,omitempty"`
		DurationOfEmailToSync                       *EmailSyncDuration                           `json:"durationOfEmailToSync,omitempty"`
		EmailAddressSource                          *UserEmailSource                             `json:"emailAddressSource,omitempty"`
		HostName                                    *string                                      `json:"hostName,omitempty"`
		RequireSsl                                  *bool                                        `json:"requireSsl,omitempty"`
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

	s.SyncCalendar = decoded.SyncCalendar
	s.SyncContacts = decoded.SyncContacts
	s.SyncTasks = decoded.SyncTasks
	s.Assignments = decoded.Assignments
	s.AuthenticationMethod = decoded.AuthenticationMethod
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DeviceManagementApplicabilityRuleDeviceMode = decoded.DeviceManagementApplicabilityRuleDeviceMode
	s.DeviceManagementApplicabilityRuleOsEdition = decoded.DeviceManagementApplicabilityRuleOsEdition
	s.DeviceManagementApplicabilityRuleOsVersion = decoded.DeviceManagementApplicabilityRuleOsVersion
	s.DeviceSettingStateSummaries = decoded.DeviceSettingStateSummaries
	s.DeviceStatusOverview = decoded.DeviceStatusOverview
	s.DeviceStatuses = decoded.DeviceStatuses
	s.DisplayName = decoded.DisplayName
	s.DurationOfEmailToSync = decoded.DurationOfEmailToSync
	s.EmailAddressSource = decoded.EmailAddressSource
	s.GroupAssignments = decoded.GroupAssignments
	s.HostName = decoded.HostName
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RequireSsl = decoded.RequireSsl
	s.RoleScopeTagIds = decoded.RoleScopeTagIds
	s.SupportsScopeTags = decoded.SupportsScopeTags
	s.UserStatusOverview = decoded.UserStatusOverview
	s.UserStatuses = decoded.UserStatuses
	s.UsernameSource = decoded.UsernameSource
	s.Version = decoded.Version

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AndroidWorkProfileNineWorkEasConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identityCertificate"]; ok {
		impl, err := UnmarshalAndroidWorkProfileCertificateProfileBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'IdentityCertificate' for 'AndroidWorkProfileNineWorkEasConfiguration': %+v", err)
		}
		s.IdentityCertificate = &impl
	}

	return nil
}
