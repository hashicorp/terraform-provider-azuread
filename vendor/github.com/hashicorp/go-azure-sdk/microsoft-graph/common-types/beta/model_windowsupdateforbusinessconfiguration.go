package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceConfiguration = WindowsUpdateForBusinessConfiguration{}

type WindowsUpdateForBusinessConfiguration struct {
	// When TRUE, allows eligible Windows 10 devices to upgrade to Windows 11. When FALSE, implies the device stays on the
	// existing operating system. Returned by default. Query parameters are not supported.
	AllowWindows11Upgrade *bool `json:"allowWindows11Upgrade,omitempty"`

	// Auto restart required notification dismissal method
	AutoRestartNotificationDismissal *AutoRestartNotificationDismissalMethod `json:"autoRestartNotificationDismissal,omitempty"`

	// Possible values for automatic update mode.
	AutomaticUpdateMode *AutomaticUpdateMode `json:"automaticUpdateMode,omitempty"`

	// Which branch devices will receive their updates from
	BusinessReadyUpdatesOnly *WindowsUpdateType `json:"businessReadyUpdatesOnly,omitempty"`

	// Number of days before feature updates are installed automatically with valid range from 0 to 30 days. Returned by
	// default. Query parameters are not supported.
	DeadlineForFeatureUpdatesInDays nullable.Type[int64] `json:"deadlineForFeatureUpdatesInDays,omitempty"`

	// Number of days before quality updates are installed automatically with valid range from 0 to 30 days. Returned by
	// default. Query parameters are not supported.
	DeadlineForQualityUpdatesInDays nullable.Type[int64] `json:"deadlineForQualityUpdatesInDays,omitempty"`

	// Number of days after deadline until restarts occur automatically with valid range from 0 to 7 days. Returned by
	// default. Query parameters are not supported.
	DeadlineGracePeriodInDays nullable.Type[int64] `json:"deadlineGracePeriodInDays,omitempty"`

	// Delivery optimization mode for peer distribution
	DeliveryOptimizationMode *WindowsDeliveryOptimizationMode `json:"deliveryOptimizationMode,omitempty"`

	// When TRUE, excludes Windows update Drivers. When FALSE, does not exclude Windows update Drivers. Returned by default.
	// Query parameters are not supported.
	DriversExcluded *bool `json:"driversExcluded,omitempty"`

	// Deadline in days before automatically scheduling and executing a pending restart outside of active hours, with valid
	// range from 2 to 30 days. Returned by default. Query parameters are not supported.
	EngagedRestartDeadlineInDays nullable.Type[int64] `json:"engagedRestartDeadlineInDays,omitempty"`

	// Number of days a user can snooze Engaged Restart reminder notifications with valid range from 1 to 3 days. Returned
	// by default. Query parameters are not supported.
	EngagedRestartSnoozeScheduleInDays nullable.Type[int64] `json:"engagedRestartSnoozeScheduleInDays,omitempty"`

	// Number of days before transitioning from Auto Restarts scheduled outside of active hours to Engaged Restart, which
	// requires the user to schedule, with valid range from 0 to 30 days. Returned by default. Query parameters are not
	// supported.
	EngagedRestartTransitionScheduleInDays nullable.Type[int64] `json:"engagedRestartTransitionScheduleInDays,omitempty"`

	// Defer Feature Updates by these many days with valid range from 0 to 30 days. Returned by default. Query parameters
	// are not supported.
	FeatureUpdatesDeferralPeriodInDays *int64 `json:"featureUpdatesDeferralPeriodInDays,omitempty"`

	// The Feature Updates Pause Expiry datetime. This value is 35 days from the time admin paused or extended the pause for
	// the ring. Returned by default. Query parameters are not supported.
	FeatureUpdatesPauseExpiryDateTime *string `json:"featureUpdatesPauseExpiryDateTime,omitempty"`

	// The Feature Updates Pause start date. This value is the time when the admin paused or extended the pause for the
	// ring. Returned by default. Query parameters are not supported. This property is read-only.
	FeatureUpdatesPauseStartDate nullable.Type[string] `json:"featureUpdatesPauseStartDate,omitempty"`

	// When TRUE, assigned devices are paused from receiving feature updates for up to 35 days from the time you pause the
	// ring. When FALSE, does not pause Feature Updates. Returned by default. Query parameters are not supported.s
	FeatureUpdatesPaused *bool `json:"featureUpdatesPaused,omitempty"`

	// The Feature Updates Rollback Start datetime.This value is the time when the admin rolled back the Feature update for
	// the ring.Returned by default.Query parameters are not supported.
	FeatureUpdatesRollbackStartDateTime *string `json:"featureUpdatesRollbackStartDateTime,omitempty"`

	// The number of days after a Feature Update for which a rollback is valid with valid range from 2 to 60 days. Returned
	// by default. Query parameters are not supported.
	FeatureUpdatesRollbackWindowInDays nullable.Type[int64] `json:"featureUpdatesRollbackWindowInDays,omitempty"`

	// When TRUE, rollback Feature Updates on the next device check in. When FALSE, do not rollback Feature Updates on the
	// next device check in. Returned by default.Query parameters are not supported.
	FeatureUpdatesWillBeRolledBack nullable.Type[bool] `json:"featureUpdatesWillBeRolledBack,omitempty"`

	// The Installation Schedule. Possible values are: ActiveHoursStart, ActiveHoursEnd, ScheduledInstallDay,
	// ScheduledInstallTime. Returned by default. Query parameters are not supported.
	InstallationSchedule WindowsUpdateInstallScheduleType `json:"installationSchedule"`

	// When TRUE, allows Microsoft Update Service. When FALSE, does not allow Microsoft Update Service. Returned by default.
	// Query parameters are not supported.
	MicrosoftUpdateServiceAllowed *bool `json:"microsoftUpdateServiceAllowed,omitempty"`

	// When TRUE the device should wait until deadline for rebooting outside of active hours. When FALSE the device should
	// not wait until deadline for rebooting outside of active hours. Returned by default. Query parameters are not
	// supported.
	PostponeRebootUntilAfterDeadline nullable.Type[bool] `json:"postponeRebootUntilAfterDeadline,omitempty"`

	// Possible values for pre-release features.
	PrereleaseFeatures *PrereleaseFeatures `json:"prereleaseFeatures,omitempty"`

	// Defer Quality Updates by these many days with valid range from 0 to 30 days. Returned by default. Query parameters
	// are not supported.
	QualityUpdatesDeferralPeriodInDays *int64 `json:"qualityUpdatesDeferralPeriodInDays,omitempty"`

	// The Quality Updates Pause Expiry datetime. This value is 35 days from the time admin paused or extended the pause for
	// the ring. Returned by default. Query parameters are not supported.
	QualityUpdatesPauseExpiryDateTime *string `json:"qualityUpdatesPauseExpiryDateTime,omitempty"`

	// The Quality Updates Pause start date. This value is the time when the admin paused or extended the pause for the
	// ring. Returned by default. Query parameters are not supported. This property is read-only.
	QualityUpdatesPauseStartDate nullable.Type[string] `json:"qualityUpdatesPauseStartDate,omitempty"`

	// When TRUE, assigned devices are paused from receiving quality updates for up to 35 days from the time you pause the
	// ring. When FALSE, does not pause Quality Updates. Returned by default. Query parameters are not supported.
	QualityUpdatesPaused *bool `json:"qualityUpdatesPaused,omitempty"`

	// The Quality Updates Rollback Start datetime. This value is the time when the admin rolled back the Quality update for
	// the ring. Returned by default. Query parameters are not supported.
	QualityUpdatesRollbackStartDateTime *string `json:"qualityUpdatesRollbackStartDateTime,omitempty"`

	// When TRUE, rollback Quality Updates on the next device check in. When FALSE, do not rollback Quality Updates on the
	// next device check in. Returned by default. Query parameters are not supported.
	QualityUpdatesWillBeRolledBack nullable.Type[bool] `json:"qualityUpdatesWillBeRolledBack,omitempty"`

	// Specify the period for auto-restart imminent warning notifications. Supported values: 15, 30 or 60 (minutes).
	// Returned by default. Query parameters are not supported.
	ScheduleImminentRestartWarningInMinutes nullable.Type[int64] `json:"scheduleImminentRestartWarningInMinutes,omitempty"`

	// Specify the period for auto-restart warning reminder notifications. Supported values: 2, 4, 8, 12 or 24 (hours).
	// Returned by default. Query parameters are not supported.
	ScheduleRestartWarningInHours nullable.Type[int64] `json:"scheduleRestartWarningInHours,omitempty"`

	// When TRUE, skips all checks before restart: Battery level = 40%, User presence, Display Needed, Presentation mode,
	// Full screen mode, phone call state, game mode etc. When FALSE, does not skip all checks before restart. Returned by
	// default. Query parameters are not supported.
	SkipChecksBeforeRestart *bool `json:"skipChecksBeforeRestart,omitempty"`

	// Windows Update Notification Display Options
	UpdateNotificationLevel *WindowsUpdateNotificationDisplayOption `json:"updateNotificationLevel,omitempty"`

	// Schedule the update installation on the weeks of the month. Possible values are: UserDefined, FirstWeek, SecondWeek,
	// ThirdWeek, FourthWeek, EveryWeek. Returned by default. Query parameters are not supported. Possible values are:
	// userDefined, firstWeek, secondWeek, thirdWeek, fourthWeek, everyWeek, unknownFutureValue.
	UpdateWeeks *WindowsUpdateForBusinessUpdateWeeks `json:"updateWeeks,omitempty"`

	// Possible values of a property
	UserPauseAccess *Enablement `json:"userPauseAccess,omitempty"`

	// Possible values of a property
	UserWindowsUpdateScanAccess *Enablement `json:"userWindowsUpdateScanAccess,omitempty"`

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

func (s WindowsUpdateForBusinessConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s WindowsUpdateForBusinessConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdateForBusinessConfiguration{}

func (s WindowsUpdateForBusinessConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdateForBusinessConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdateForBusinessConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdateForBusinessConfiguration: %+v", err)
	}

	delete(decoded, "featureUpdatesPauseStartDate")
	delete(decoded, "qualityUpdatesPauseStartDate")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdateForBusinessConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdateForBusinessConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsUpdateForBusinessConfiguration{}

func (s *WindowsUpdateForBusinessConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AllowWindows11Upgrade                       *bool                                        `json:"allowWindows11Upgrade,omitempty"`
		AutoRestartNotificationDismissal            *AutoRestartNotificationDismissalMethod      `json:"autoRestartNotificationDismissal,omitempty"`
		AutomaticUpdateMode                         *AutomaticUpdateMode                         `json:"automaticUpdateMode,omitempty"`
		BusinessReadyUpdatesOnly                    *WindowsUpdateType                           `json:"businessReadyUpdatesOnly,omitempty"`
		DeadlineForFeatureUpdatesInDays             nullable.Type[int64]                         `json:"deadlineForFeatureUpdatesInDays,omitempty"`
		DeadlineForQualityUpdatesInDays             nullable.Type[int64]                         `json:"deadlineForQualityUpdatesInDays,omitempty"`
		DeadlineGracePeriodInDays                   nullable.Type[int64]                         `json:"deadlineGracePeriodInDays,omitempty"`
		DeliveryOptimizationMode                    *WindowsDeliveryOptimizationMode             `json:"deliveryOptimizationMode,omitempty"`
		DriversExcluded                             *bool                                        `json:"driversExcluded,omitempty"`
		EngagedRestartDeadlineInDays                nullable.Type[int64]                         `json:"engagedRestartDeadlineInDays,omitempty"`
		EngagedRestartSnoozeScheduleInDays          nullable.Type[int64]                         `json:"engagedRestartSnoozeScheduleInDays,omitempty"`
		EngagedRestartTransitionScheduleInDays      nullable.Type[int64]                         `json:"engagedRestartTransitionScheduleInDays,omitempty"`
		FeatureUpdatesDeferralPeriodInDays          *int64                                       `json:"featureUpdatesDeferralPeriodInDays,omitempty"`
		FeatureUpdatesPauseExpiryDateTime           *string                                      `json:"featureUpdatesPauseExpiryDateTime,omitempty"`
		FeatureUpdatesPauseStartDate                nullable.Type[string]                        `json:"featureUpdatesPauseStartDate,omitempty"`
		FeatureUpdatesPaused                        *bool                                        `json:"featureUpdatesPaused,omitempty"`
		FeatureUpdatesRollbackStartDateTime         *string                                      `json:"featureUpdatesRollbackStartDateTime,omitempty"`
		FeatureUpdatesRollbackWindowInDays          nullable.Type[int64]                         `json:"featureUpdatesRollbackWindowInDays,omitempty"`
		FeatureUpdatesWillBeRolledBack              nullable.Type[bool]                          `json:"featureUpdatesWillBeRolledBack,omitempty"`
		MicrosoftUpdateServiceAllowed               *bool                                        `json:"microsoftUpdateServiceAllowed,omitempty"`
		PostponeRebootUntilAfterDeadline            nullable.Type[bool]                          `json:"postponeRebootUntilAfterDeadline,omitempty"`
		PrereleaseFeatures                          *PrereleaseFeatures                          `json:"prereleaseFeatures,omitempty"`
		QualityUpdatesDeferralPeriodInDays          *int64                                       `json:"qualityUpdatesDeferralPeriodInDays,omitempty"`
		QualityUpdatesPauseExpiryDateTime           *string                                      `json:"qualityUpdatesPauseExpiryDateTime,omitempty"`
		QualityUpdatesPauseStartDate                nullable.Type[string]                        `json:"qualityUpdatesPauseStartDate,omitempty"`
		QualityUpdatesPaused                        *bool                                        `json:"qualityUpdatesPaused,omitempty"`
		QualityUpdatesRollbackStartDateTime         *string                                      `json:"qualityUpdatesRollbackStartDateTime,omitempty"`
		QualityUpdatesWillBeRolledBack              nullable.Type[bool]                          `json:"qualityUpdatesWillBeRolledBack,omitempty"`
		ScheduleImminentRestartWarningInMinutes     nullable.Type[int64]                         `json:"scheduleImminentRestartWarningInMinutes,omitempty"`
		ScheduleRestartWarningInHours               nullable.Type[int64]                         `json:"scheduleRestartWarningInHours,omitempty"`
		SkipChecksBeforeRestart                     *bool                                        `json:"skipChecksBeforeRestart,omitempty"`
		UpdateNotificationLevel                     *WindowsUpdateNotificationDisplayOption      `json:"updateNotificationLevel,omitempty"`
		UpdateWeeks                                 *WindowsUpdateForBusinessUpdateWeeks         `json:"updateWeeks,omitempty"`
		UserPauseAccess                             *Enablement                                  `json:"userPauseAccess,omitempty"`
		UserWindowsUpdateScanAccess                 *Enablement                                  `json:"userWindowsUpdateScanAccess,omitempty"`
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

	s.AllowWindows11Upgrade = decoded.AllowWindows11Upgrade
	s.AutoRestartNotificationDismissal = decoded.AutoRestartNotificationDismissal
	s.AutomaticUpdateMode = decoded.AutomaticUpdateMode
	s.BusinessReadyUpdatesOnly = decoded.BusinessReadyUpdatesOnly
	s.DeadlineForFeatureUpdatesInDays = decoded.DeadlineForFeatureUpdatesInDays
	s.DeadlineForQualityUpdatesInDays = decoded.DeadlineForQualityUpdatesInDays
	s.DeadlineGracePeriodInDays = decoded.DeadlineGracePeriodInDays
	s.DeliveryOptimizationMode = decoded.DeliveryOptimizationMode
	s.DriversExcluded = decoded.DriversExcluded
	s.EngagedRestartDeadlineInDays = decoded.EngagedRestartDeadlineInDays
	s.EngagedRestartSnoozeScheduleInDays = decoded.EngagedRestartSnoozeScheduleInDays
	s.EngagedRestartTransitionScheduleInDays = decoded.EngagedRestartTransitionScheduleInDays
	s.FeatureUpdatesDeferralPeriodInDays = decoded.FeatureUpdatesDeferralPeriodInDays
	s.FeatureUpdatesPauseExpiryDateTime = decoded.FeatureUpdatesPauseExpiryDateTime
	s.FeatureUpdatesPauseStartDate = decoded.FeatureUpdatesPauseStartDate
	s.FeatureUpdatesPaused = decoded.FeatureUpdatesPaused
	s.FeatureUpdatesRollbackStartDateTime = decoded.FeatureUpdatesRollbackStartDateTime
	s.FeatureUpdatesRollbackWindowInDays = decoded.FeatureUpdatesRollbackWindowInDays
	s.FeatureUpdatesWillBeRolledBack = decoded.FeatureUpdatesWillBeRolledBack
	s.MicrosoftUpdateServiceAllowed = decoded.MicrosoftUpdateServiceAllowed
	s.PostponeRebootUntilAfterDeadline = decoded.PostponeRebootUntilAfterDeadline
	s.PrereleaseFeatures = decoded.PrereleaseFeatures
	s.QualityUpdatesDeferralPeriodInDays = decoded.QualityUpdatesDeferralPeriodInDays
	s.QualityUpdatesPauseExpiryDateTime = decoded.QualityUpdatesPauseExpiryDateTime
	s.QualityUpdatesPauseStartDate = decoded.QualityUpdatesPauseStartDate
	s.QualityUpdatesPaused = decoded.QualityUpdatesPaused
	s.QualityUpdatesRollbackStartDateTime = decoded.QualityUpdatesRollbackStartDateTime
	s.QualityUpdatesWillBeRolledBack = decoded.QualityUpdatesWillBeRolledBack
	s.ScheduleImminentRestartWarningInMinutes = decoded.ScheduleImminentRestartWarningInMinutes
	s.ScheduleRestartWarningInHours = decoded.ScheduleRestartWarningInHours
	s.SkipChecksBeforeRestart = decoded.SkipChecksBeforeRestart
	s.UpdateNotificationLevel = decoded.UpdateNotificationLevel
	s.UpdateWeeks = decoded.UpdateWeeks
	s.UserPauseAccess = decoded.UserPauseAccess
	s.UserWindowsUpdateScanAccess = decoded.UserWindowsUpdateScanAccess
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
		return fmt.Errorf("unmarshaling WindowsUpdateForBusinessConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["installationSchedule"]; ok {
		impl, err := UnmarshalWindowsUpdateInstallScheduleTypeImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'InstallationSchedule' for 'WindowsUpdateForBusinessConfiguration': %+v", err)
		}
		s.InstallationSchedule = impl
	}

	return nil
}
