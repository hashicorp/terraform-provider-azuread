package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceConfiguration = Windows10TeamGeneralConfiguration{}

type Windows10TeamGeneralConfiguration struct {
	// Indicates whether or not to Block Azure Operational Insights.
	AzureOperationalInsightsBlockTelemetry *bool `json:"azureOperationalInsightsBlockTelemetry,omitempty"`

	// The Azure Operational Insights workspace id.
	AzureOperationalInsightsWorkspaceId nullable.Type[string] `json:"azureOperationalInsightsWorkspaceId,omitempty"`

	// The Azure Operational Insights Workspace key.
	AzureOperationalInsightsWorkspaceKey nullable.Type[string] `json:"azureOperationalInsightsWorkspaceKey,omitempty"`

	// Specifies whether to automatically launch the Connect app whenever a projection is initiated.
	ConnectAppBlockAutoLaunch *bool `json:"connectAppBlockAutoLaunch,omitempty"`

	// Indicates whether or not to Block setting a maintenance window for device updates.
	MaintenanceWindowBlocked *bool `json:"maintenanceWindowBlocked,omitempty"`

	// Maintenance window duration for device updates. Valid values 0 to 5
	MaintenanceWindowDurationInHours nullable.Type[int64] `json:"maintenanceWindowDurationInHours,omitempty"`

	// Maintenance window start time for device updates.
	MaintenanceWindowStartTime nullable.Type[string] `json:"maintenanceWindowStartTime,omitempty"`

	// Indicates whether or not to Block wireless projection.
	MiracastBlocked *bool `json:"miracastBlocked,omitempty"`

	// Possible values for Miracast channel.
	MiracastChannel *MiracastChannel `json:"miracastChannel,omitempty"`

	// Indicates whether or not to require a pin for wireless projection.
	MiracastRequirePin *bool `json:"miracastRequirePin,omitempty"`

	// Specifies whether to disable the 'My meetings and files' feature in the Start menu, which shows the signed-in user's
	// meetings and files from Office 365.
	SettingsBlockMyMeetingsAndFiles *bool `json:"settingsBlockMyMeetingsAndFiles,omitempty"`

	// Specifies whether to allow the ability to resume a session when the session times out.
	SettingsBlockSessionResume *bool `json:"settingsBlockSessionResume,omitempty"`

	// Specifies whether to disable auto-populating of the sign-in dialog with invitees from scheduled meetings.
	SettingsBlockSigninSuggestions *bool `json:"settingsBlockSigninSuggestions,omitempty"`

	// Specifies the default volume value for a new session. Permitted values are 0-100. The default is 45. Valid values 0
	// to 100
	SettingsDefaultVolume nullable.Type[int64] `json:"settingsDefaultVolume,omitempty"`

	// Specifies the number of minutes until the Hub screen turns off.
	SettingsScreenTimeoutInMinutes nullable.Type[int64] `json:"settingsScreenTimeoutInMinutes,omitempty"`

	// Specifies the number of minutes until the session times out.
	SettingsSessionTimeoutInMinutes nullable.Type[int64] `json:"settingsSessionTimeoutInMinutes,omitempty"`

	// Specifies the number of minutes until the Hub enters sleep mode.
	SettingsSleepTimeoutInMinutes nullable.Type[int64] `json:"settingsSleepTimeoutInMinutes,omitempty"`

	// The welcome screen background image URL. The URL must use the HTTPS protocol and return a PNG image.
	WelcomeScreenBackgroundImageUrl nullable.Type[string] `json:"welcomeScreenBackgroundImageUrl,omitempty"`

	// Indicates whether or not to Block the welcome screen from waking up automatically when someone enters the room.
	WelcomeScreenBlockAutomaticWakeUp *bool `json:"welcomeScreenBlockAutomaticWakeUp,omitempty"`

	// Possible values for welcome screen meeting information.
	WelcomeScreenMeetingInformation *WelcomeScreenMeetingInformation `json:"welcomeScreenMeetingInformation,omitempty"`

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

func (s Windows10TeamGeneralConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s Windows10TeamGeneralConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Windows10TeamGeneralConfiguration{}

func (s Windows10TeamGeneralConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper Windows10TeamGeneralConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Windows10TeamGeneralConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Windows10TeamGeneralConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windows10TeamGeneralConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Windows10TeamGeneralConfiguration: %+v", err)
	}

	return encoded, nil
}
