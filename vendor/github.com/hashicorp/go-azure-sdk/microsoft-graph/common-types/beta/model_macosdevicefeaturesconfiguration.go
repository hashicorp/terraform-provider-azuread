package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AppleDeviceFeaturesConfigurationBase = MacOSDeviceFeaturesConfiguration{}

type MacOSDeviceFeaturesConfiguration struct {
	// Whether to show admin host information on the login window.
	AdminShowHostInfo *bool `json:"adminShowHostInfo,omitempty"`

	// Gets or sets a list that maps apps to their associated domains. Application identifiers must be unique. This
	// collection can contain a maximum of 500 elements.
	AppAssociatedDomains *[]MacOSAssociatedDomainsItem `json:"appAssociatedDomains,omitempty"`

	// DEPRECATED: use appAssociatedDomains instead. Gets or sets a list that maps apps to their associated domains. The key
	// should match the app's ID, and the value should be a string in the form of 'service:domain' where domain is a fully
	// qualified hostname (e.g. webcredentials:example.com). This collection can contain a maximum of 500 elements.
	AssociatedDomains *[]KeyValuePair `json:"associatedDomains,omitempty"`

	// Whether to show the name and password dialog or a list of users on the login window.
	AuthorizedUsersListHidden *bool `json:"authorizedUsersListHidden,omitempty"`

	// Whether to hide admin users in the authorized users list on the login window.
	AuthorizedUsersListHideAdminUsers *bool `json:"authorizedUsersListHideAdminUsers,omitempty"`

	// Whether to show only network and system users in the authorized users list on the login window.
	AuthorizedUsersListHideLocalUsers *bool `json:"authorizedUsersListHideLocalUsers,omitempty"`

	// Whether to hide mobile users in the authorized users list on the login window.
	AuthorizedUsersListHideMobileAccounts *bool `json:"authorizedUsersListHideMobileAccounts,omitempty"`

	// Whether to show network users in the authorized users list on the login window.
	AuthorizedUsersListIncludeNetworkUsers *bool `json:"authorizedUsersListIncludeNetworkUsers,omitempty"`

	// Whether to show other users in the authorized users list on the login window.
	AuthorizedUsersListShowOtherManagedUsers *bool `json:"authorizedUsersListShowOtherManagedUsers,omitempty"`

	// List of applications, files, folders, and other items to launch when the user logs in. This collection can contain a
	// maximum of 500 elements.
	AutoLaunchItems *[]MacOSLaunchItem `json:"autoLaunchItems,omitempty"`

	// Whether the Other user will disregard use of the console special user name.
	ConsoleAccessDisabled *bool `json:"consoleAccessDisabled,omitempty"`

	// Prevents content caches from purging content to free up disk space for other apps.
	ContentCachingBlockDeletion *bool `json:"contentCachingBlockDeletion,omitempty"`

	// A list of custom IP ranges content caches will use to listen for clients. This collection can contain a maximum of
	// 500 elements.
	ContentCachingClientListenRanges *[]IPRange `json:"contentCachingClientListenRanges,omitempty"`

	// Determines which clients a content cache will serve.
	ContentCachingClientPolicy *MacOSContentCachingClientPolicy `json:"contentCachingClientPolicy,omitempty"`

	// The path to the directory used to store cached content. The value must be (or end with) /Library/Application
	// Support/Apple/AssetCache/Data
	ContentCachingDataPath nullable.Type[string] `json:"contentCachingDataPath,omitempty"`

	// Disables internet connection sharing.
	ContentCachingDisableConnectionSharing *bool `json:"contentCachingDisableConnectionSharing,omitempty"`

	// Enables content caching and prevents it from being disabled by the user.
	ContentCachingEnabled *bool `json:"contentCachingEnabled,omitempty"`

	// Forces internet connection sharing. contentCachingDisableConnectionSharing overrides this setting.
	ContentCachingForceConnectionSharing *bool `json:"contentCachingForceConnectionSharing,omitempty"`

	// Prevent the device from sleeping if content caching is enabled.
	ContentCachingKeepAwake *bool `json:"contentCachingKeepAwake,omitempty"`

	// Enables logging of IP addresses and ports of clients that request cached content.
	ContentCachingLogClientIdentities *bool `json:"contentCachingLogClientIdentities,omitempty"`

	// The maximum number of bytes of disk space that will be used for the content cache. A value of 0 (default) indicates
	// unlimited disk space.
	ContentCachingMaxSizeBytes nullable.Type[int64] `json:"contentCachingMaxSizeBytes,omitempty"`

	// Determines how content caches select a parent cache.
	ContentCachingParentSelectionPolicy *MacOSContentCachingParentSelectionPolicy `json:"contentCachingParentSelectionPolicy,omitempty"`

	// A list of IP addresses representing parent content caches.
	ContentCachingParents *[]string `json:"contentCachingParents,omitempty"`

	// A list of custom IP ranges content caches will use to query for content from peers caches. This collection can
	// contain a maximum of 500 elements.
	ContentCachingPeerFilterRanges *[]IPRange `json:"contentCachingPeerFilterRanges,omitempty"`

	// A list of custom IP ranges content caches will use to listen for peer caches. This collection can contain a maximum
	// of 500 elements.
	ContentCachingPeerListenRanges *[]IPRange `json:"contentCachingPeerListenRanges,omitempty"`

	// Determines which content caches other content caches will peer with.
	ContentCachingPeerPolicy *MacOSContentCachingPeerPolicy `json:"contentCachingPeerPolicy,omitempty"`

	// Sets the port used for content caching. If the value is 0, a random available port will be selected. Valid values 0
	// to 65535
	ContentCachingPort nullable.Type[int64] `json:"contentCachingPort,omitempty"`

	// A list of custom IP ranges that Apple's content caching service should use to match clients to content caches. This
	// collection can contain a maximum of 500 elements.
	ContentCachingPublicRanges *[]IPRange `json:"contentCachingPublicRanges,omitempty"`

	// Display content caching alerts as system notifications.
	ContentCachingShowAlerts *bool `json:"contentCachingShowAlerts,omitempty"`

	// Indicates the type of content allowed to be cached by Apple's content caching service.
	ContentCachingType *MacOSContentCachingType `json:"contentCachingType,omitempty"`

	// Whether the Log Out menu item on the login window will be disabled while the user is logged in.
	LogOutDisabledWhileLoggedIn *bool `json:"logOutDisabledWhileLoggedIn,omitempty"`

	// Custom text to be displayed on the login window.
	LoginWindowText nullable.Type[string] `json:"loginWindowText,omitempty"`

	// Gets or sets a single sign-on extension profile.
	MacOSSingleSignOnExtension *MacOSSingleSignOnExtension `json:"macOSSingleSignOnExtension,omitempty"`

	// Whether the Power Off menu item on the login window will be disabled while the user is logged in.
	PowerOffDisabledWhileLoggedIn *bool `json:"powerOffDisabledWhileLoggedIn,omitempty"`

	// Whether to hide the Restart button item on the login window.
	RestartDisabled *bool `json:"restartDisabled,omitempty"`

	// Whether the Restart menu item on the login window will be disabled while the user is logged in.
	RestartDisabledWhileLoggedIn *bool `json:"restartDisabledWhileLoggedIn,omitempty"`

	// Whether to disable the immediate screen lock functions.
	ScreenLockDisableImmediate *bool `json:"screenLockDisableImmediate,omitempty"`

	// Whether to hide the Shut Down button item on the login window.
	ShutDownDisabled *bool `json:"shutDownDisabled,omitempty"`

	// Whether the Shut Down menu item on the login window will be disabled while the user is logged in.
	ShutDownDisabledWhileLoggedIn *bool `json:"shutDownDisabledWhileLoggedIn,omitempty"`

	// Gets or sets a single sign-on extension profile. Deprecated: use MacOSSingleSignOnExtension instead.
	SingleSignOnExtension SingleSignOnExtension `json:"singleSignOnExtension"`

	// PKINIT Certificate for the authentication with single sign-on extensions.
	SingleSignOnExtensionPkinitCertificate *MacOSCertificateProfileBase `json:"singleSignOnExtensionPkinitCertificate,omitempty"`

	// Whether to hide the Sleep menu item on the login window.
	SleepDisabled *bool `json:"sleepDisabled,omitempty"`

	// Fields inherited from AppleDeviceFeaturesConfigurationBase

	// An array of AirPrint printers that should always be shown. This collection can contain a maximum of 500 elements.
	AirPrintDestinations *[]AirPrintDestination `json:"airPrintDestinations,omitempty"`

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

func (s MacOSDeviceFeaturesConfiguration) AppleDeviceFeaturesConfigurationBase() BaseAppleDeviceFeaturesConfigurationBaseImpl {
	return BaseAppleDeviceFeaturesConfigurationBaseImpl{
		AirPrintDestinations: s.AirPrintDestinations,
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

func (s MacOSDeviceFeaturesConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s MacOSDeviceFeaturesConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MacOSDeviceFeaturesConfiguration{}

func (s MacOSDeviceFeaturesConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper MacOSDeviceFeaturesConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MacOSDeviceFeaturesConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MacOSDeviceFeaturesConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.macOSDeviceFeaturesConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MacOSDeviceFeaturesConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &MacOSDeviceFeaturesConfiguration{}

func (s *MacOSDeviceFeaturesConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AdminShowHostInfo                           *bool                                        `json:"adminShowHostInfo,omitempty"`
		AppAssociatedDomains                        *[]MacOSAssociatedDomainsItem                `json:"appAssociatedDomains,omitempty"`
		AssociatedDomains                           *[]KeyValuePair                              `json:"associatedDomains,omitempty"`
		AuthorizedUsersListHidden                   *bool                                        `json:"authorizedUsersListHidden,omitempty"`
		AuthorizedUsersListHideAdminUsers           *bool                                        `json:"authorizedUsersListHideAdminUsers,omitempty"`
		AuthorizedUsersListHideLocalUsers           *bool                                        `json:"authorizedUsersListHideLocalUsers,omitempty"`
		AuthorizedUsersListHideMobileAccounts       *bool                                        `json:"authorizedUsersListHideMobileAccounts,omitempty"`
		AuthorizedUsersListIncludeNetworkUsers      *bool                                        `json:"authorizedUsersListIncludeNetworkUsers,omitempty"`
		AuthorizedUsersListShowOtherManagedUsers    *bool                                        `json:"authorizedUsersListShowOtherManagedUsers,omitempty"`
		AutoLaunchItems                             *[]MacOSLaunchItem                           `json:"autoLaunchItems,omitempty"`
		ConsoleAccessDisabled                       *bool                                        `json:"consoleAccessDisabled,omitempty"`
		ContentCachingBlockDeletion                 *bool                                        `json:"contentCachingBlockDeletion,omitempty"`
		ContentCachingClientPolicy                  *MacOSContentCachingClientPolicy             `json:"contentCachingClientPolicy,omitempty"`
		ContentCachingDataPath                      nullable.Type[string]                        `json:"contentCachingDataPath,omitempty"`
		ContentCachingDisableConnectionSharing      *bool                                        `json:"contentCachingDisableConnectionSharing,omitempty"`
		ContentCachingEnabled                       *bool                                        `json:"contentCachingEnabled,omitempty"`
		ContentCachingForceConnectionSharing        *bool                                        `json:"contentCachingForceConnectionSharing,omitempty"`
		ContentCachingKeepAwake                     *bool                                        `json:"contentCachingKeepAwake,omitempty"`
		ContentCachingLogClientIdentities           *bool                                        `json:"contentCachingLogClientIdentities,omitempty"`
		ContentCachingMaxSizeBytes                  nullable.Type[int64]                         `json:"contentCachingMaxSizeBytes,omitempty"`
		ContentCachingParentSelectionPolicy         *MacOSContentCachingParentSelectionPolicy    `json:"contentCachingParentSelectionPolicy,omitempty"`
		ContentCachingParents                       *[]string                                    `json:"contentCachingParents,omitempty"`
		ContentCachingPeerPolicy                    *MacOSContentCachingPeerPolicy               `json:"contentCachingPeerPolicy,omitempty"`
		ContentCachingPort                          nullable.Type[int64]                         `json:"contentCachingPort,omitempty"`
		ContentCachingShowAlerts                    *bool                                        `json:"contentCachingShowAlerts,omitempty"`
		ContentCachingType                          *MacOSContentCachingType                     `json:"contentCachingType,omitempty"`
		LogOutDisabledWhileLoggedIn                 *bool                                        `json:"logOutDisabledWhileLoggedIn,omitempty"`
		LoginWindowText                             nullable.Type[string]                        `json:"loginWindowText,omitempty"`
		PowerOffDisabledWhileLoggedIn               *bool                                        `json:"powerOffDisabledWhileLoggedIn,omitempty"`
		RestartDisabled                             *bool                                        `json:"restartDisabled,omitempty"`
		RestartDisabledWhileLoggedIn                *bool                                        `json:"restartDisabledWhileLoggedIn,omitempty"`
		ScreenLockDisableImmediate                  *bool                                        `json:"screenLockDisableImmediate,omitempty"`
		ShutDownDisabled                            *bool                                        `json:"shutDownDisabled,omitempty"`
		ShutDownDisabledWhileLoggedIn               *bool                                        `json:"shutDownDisabledWhileLoggedIn,omitempty"`
		SleepDisabled                               *bool                                        `json:"sleepDisabled,omitempty"`
		AirPrintDestinations                        *[]AirPrintDestination                       `json:"airPrintDestinations,omitempty"`
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

	s.AdminShowHostInfo = decoded.AdminShowHostInfo
	s.AppAssociatedDomains = decoded.AppAssociatedDomains
	s.AssociatedDomains = decoded.AssociatedDomains
	s.AuthorizedUsersListHidden = decoded.AuthorizedUsersListHidden
	s.AuthorizedUsersListHideAdminUsers = decoded.AuthorizedUsersListHideAdminUsers
	s.AuthorizedUsersListHideLocalUsers = decoded.AuthorizedUsersListHideLocalUsers
	s.AuthorizedUsersListHideMobileAccounts = decoded.AuthorizedUsersListHideMobileAccounts
	s.AuthorizedUsersListIncludeNetworkUsers = decoded.AuthorizedUsersListIncludeNetworkUsers
	s.AuthorizedUsersListShowOtherManagedUsers = decoded.AuthorizedUsersListShowOtherManagedUsers
	s.AutoLaunchItems = decoded.AutoLaunchItems
	s.ConsoleAccessDisabled = decoded.ConsoleAccessDisabled
	s.ContentCachingBlockDeletion = decoded.ContentCachingBlockDeletion
	s.ContentCachingClientPolicy = decoded.ContentCachingClientPolicy
	s.ContentCachingDataPath = decoded.ContentCachingDataPath
	s.ContentCachingDisableConnectionSharing = decoded.ContentCachingDisableConnectionSharing
	s.ContentCachingEnabled = decoded.ContentCachingEnabled
	s.ContentCachingForceConnectionSharing = decoded.ContentCachingForceConnectionSharing
	s.ContentCachingKeepAwake = decoded.ContentCachingKeepAwake
	s.ContentCachingLogClientIdentities = decoded.ContentCachingLogClientIdentities
	s.ContentCachingMaxSizeBytes = decoded.ContentCachingMaxSizeBytes
	s.ContentCachingParentSelectionPolicy = decoded.ContentCachingParentSelectionPolicy
	s.ContentCachingParents = decoded.ContentCachingParents
	s.ContentCachingPeerPolicy = decoded.ContentCachingPeerPolicy
	s.ContentCachingPort = decoded.ContentCachingPort
	s.ContentCachingShowAlerts = decoded.ContentCachingShowAlerts
	s.ContentCachingType = decoded.ContentCachingType
	s.LogOutDisabledWhileLoggedIn = decoded.LogOutDisabledWhileLoggedIn
	s.LoginWindowText = decoded.LoginWindowText
	s.PowerOffDisabledWhileLoggedIn = decoded.PowerOffDisabledWhileLoggedIn
	s.RestartDisabled = decoded.RestartDisabled
	s.RestartDisabledWhileLoggedIn = decoded.RestartDisabledWhileLoggedIn
	s.ScreenLockDisableImmediate = decoded.ScreenLockDisableImmediate
	s.ShutDownDisabled = decoded.ShutDownDisabled
	s.ShutDownDisabledWhileLoggedIn = decoded.ShutDownDisabledWhileLoggedIn
	s.SleepDisabled = decoded.SleepDisabled
	s.AirPrintDestinations = decoded.AirPrintDestinations
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
		return fmt.Errorf("unmarshaling MacOSDeviceFeaturesConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["contentCachingClientListenRanges"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ContentCachingClientListenRanges into list []json.RawMessage: %+v", err)
		}

		output := make([]IPRange, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIPRangeImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ContentCachingClientListenRanges' for 'MacOSDeviceFeaturesConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ContentCachingClientListenRanges = &output
	}

	if v, ok := temp["contentCachingPeerFilterRanges"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ContentCachingPeerFilterRanges into list []json.RawMessage: %+v", err)
		}

		output := make([]IPRange, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIPRangeImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ContentCachingPeerFilterRanges' for 'MacOSDeviceFeaturesConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ContentCachingPeerFilterRanges = &output
	}

	if v, ok := temp["contentCachingPeerListenRanges"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ContentCachingPeerListenRanges into list []json.RawMessage: %+v", err)
		}

		output := make([]IPRange, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIPRangeImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ContentCachingPeerListenRanges' for 'MacOSDeviceFeaturesConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ContentCachingPeerListenRanges = &output
	}

	if v, ok := temp["contentCachingPublicRanges"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ContentCachingPublicRanges into list []json.RawMessage: %+v", err)
		}

		output := make([]IPRange, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIPRangeImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ContentCachingPublicRanges' for 'MacOSDeviceFeaturesConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ContentCachingPublicRanges = &output
	}

	if v, ok := temp["macOSSingleSignOnExtension"]; ok {
		impl, err := UnmarshalMacOSSingleSignOnExtensionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'MacOSSingleSignOnExtension' for 'MacOSDeviceFeaturesConfiguration': %+v", err)
		}
		s.MacOSSingleSignOnExtension = &impl
	}

	if v, ok := temp["singleSignOnExtension"]; ok {
		impl, err := UnmarshalSingleSignOnExtensionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SingleSignOnExtension' for 'MacOSDeviceFeaturesConfiguration': %+v", err)
		}
		s.SingleSignOnExtension = impl
	}

	if v, ok := temp["singleSignOnExtensionPkinitCertificate"]; ok {
		impl, err := UnmarshalMacOSCertificateProfileBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SingleSignOnExtensionPkinitCertificate' for 'MacOSDeviceFeaturesConfiguration': %+v", err)
		}
		s.SingleSignOnExtensionPkinitCertificate = &impl
	}

	return nil
}
