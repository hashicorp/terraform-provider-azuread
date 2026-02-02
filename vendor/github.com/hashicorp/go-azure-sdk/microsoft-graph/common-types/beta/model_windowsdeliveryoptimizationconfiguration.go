package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceConfiguration = WindowsDeliveryOptimizationConfiguration{}

type WindowsDeliveryOptimizationConfiguration struct {
	// Specifies number of seconds to delay an HTTP source in a background download that is allowed to use peer-to-peer.
	// Valid values 0 to 4294967295
	BackgroundDownloadFromHttpDelayInSeconds *int64 `json:"backgroundDownloadFromHttpDelayInSeconds,omitempty"`

	// Specifies foreground and background bandwidth usage using percentages, absolutes, or hours.
	BandwidthMode DeliveryOptimizationBandwidth `json:"bandwidthMode"`

	// Specifies number of seconds to delay a fall back from cache servers to an HTTP source for a background download.
	// Valid values 0 to 2592000.
	CacheServerBackgroundDownloadFallbackToHttpDelayInSeconds *int64 `json:"cacheServerBackgroundDownloadFallbackToHttpDelayInSeconds,omitempty"`

	// Specifies number of seconds to delay a fall back from cache servers to an HTTP source for a foreground download.
	// Valid values 0 to 2592000.​
	CacheServerForegroundDownloadFallbackToHttpDelayInSeconds *int64 `json:"cacheServerForegroundDownloadFallbackToHttpDelayInSeconds,omitempty"`

	// Specifies cache servers host names.
	CacheServerHostNames *[]string `json:"cacheServerHostNames,omitempty"`

	// Delivery optimization mode for peer distribution
	DeliveryOptimizationMode *WindowsDeliveryOptimizationMode `json:"deliveryOptimizationMode,omitempty"`

	// Specifies number of seconds to delay an HTTP source in a foreground download that is allowed to use peer-to-peer
	// (0-86400). Valid values 0 to 86400
	ForegroundDownloadFromHttpDelayInSeconds nullable.Type[int64] `json:"foregroundDownloadFromHttpDelayInSeconds,omitempty"`

	// Specifies to restrict peer selection to a specfic source.
	GroupIdSource DeliveryOptimizationGroupIdSource `json:"groupIdSource"`

	// Specifies the maximum time in days that each file is held in the Delivery Optimization cache after downloading
	// successfully (0-3650). Valid values 0 to 3650
	MaximumCacheAgeInDays nullable.Type[int64] `json:"maximumCacheAgeInDays,omitempty"`

	// Specifies the maximum cache size that Delivery Optimization either as a percentage or in GB.
	MaximumCacheSize DeliveryOptimizationMaxCacheSize `json:"maximumCacheSize"`

	// Specifies the minimum battery percentage to allow the device to upload data (0-100). Valid values 0 to 100
	MinimumBatteryPercentageAllowedToUpload nullable.Type[int64] `json:"minimumBatteryPercentageAllowedToUpload,omitempty"`

	// Specifies the minimum disk size in GB to use Peer Caching (1-100000). Valid values 1 to 100000
	MinimumDiskSizeAllowedToPeerInGigabytes nullable.Type[int64] `json:"minimumDiskSizeAllowedToPeerInGigabytes,omitempty"`

	// Specifies the minimum content file size in MB enabled to use Peer Caching (1-100000). Valid values 1 to 100000
	MinimumFileSizeToCacheInMegabytes nullable.Type[int64] `json:"minimumFileSizeToCacheInMegabytes,omitempty"`

	// Specifies the minimum RAM size in GB to use Peer Caching (1-100000). Valid values 1 to 100000
	MinimumRamAllowedToPeerInGigabytes nullable.Type[int64] `json:"minimumRamAllowedToPeerInGigabytes,omitempty"`

	// Specifies the drive that Delivery Optimization should use for its cache.
	ModifyCacheLocation nullable.Type[string] `json:"modifyCacheLocation,omitempty"`

	// Values to restrict peer selection by.
	RestrictPeerSelectionBy *DeliveryOptimizationRestrictPeerSelectionByOptions `json:"restrictPeerSelectionBy,omitempty"`

	// Possible values of a property
	VpnPeerCaching *Enablement `json:"vpnPeerCaching,omitempty"`

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

func (s WindowsDeliveryOptimizationConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s WindowsDeliveryOptimizationConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsDeliveryOptimizationConfiguration{}

func (s WindowsDeliveryOptimizationConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper WindowsDeliveryOptimizationConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsDeliveryOptimizationConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsDeliveryOptimizationConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsDeliveryOptimizationConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsDeliveryOptimizationConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsDeliveryOptimizationConfiguration{}

func (s *WindowsDeliveryOptimizationConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		BackgroundDownloadFromHttpDelayInSeconds                  *int64                                              `json:"backgroundDownloadFromHttpDelayInSeconds,omitempty"`
		CacheServerBackgroundDownloadFallbackToHttpDelayInSeconds *int64                                              `json:"cacheServerBackgroundDownloadFallbackToHttpDelayInSeconds,omitempty"`
		CacheServerForegroundDownloadFallbackToHttpDelayInSeconds *int64                                              `json:"cacheServerForegroundDownloadFallbackToHttpDelayInSeconds,omitempty"`
		CacheServerHostNames                                      *[]string                                           `json:"cacheServerHostNames,omitempty"`
		DeliveryOptimizationMode                                  *WindowsDeliveryOptimizationMode                    `json:"deliveryOptimizationMode,omitempty"`
		ForegroundDownloadFromHttpDelayInSeconds                  nullable.Type[int64]                                `json:"foregroundDownloadFromHttpDelayInSeconds,omitempty"`
		MaximumCacheAgeInDays                                     nullable.Type[int64]                                `json:"maximumCacheAgeInDays,omitempty"`
		MinimumBatteryPercentageAllowedToUpload                   nullable.Type[int64]                                `json:"minimumBatteryPercentageAllowedToUpload,omitempty"`
		MinimumDiskSizeAllowedToPeerInGigabytes                   nullable.Type[int64]                                `json:"minimumDiskSizeAllowedToPeerInGigabytes,omitempty"`
		MinimumFileSizeToCacheInMegabytes                         nullable.Type[int64]                                `json:"minimumFileSizeToCacheInMegabytes,omitempty"`
		MinimumRamAllowedToPeerInGigabytes                        nullable.Type[int64]                                `json:"minimumRamAllowedToPeerInGigabytes,omitempty"`
		ModifyCacheLocation                                       nullable.Type[string]                               `json:"modifyCacheLocation,omitempty"`
		RestrictPeerSelectionBy                                   *DeliveryOptimizationRestrictPeerSelectionByOptions `json:"restrictPeerSelectionBy,omitempty"`
		VpnPeerCaching                                            *Enablement                                         `json:"vpnPeerCaching,omitempty"`
		Assignments                                               *[]DeviceConfigurationAssignment                    `json:"assignments,omitempty"`
		CreatedDateTime                                           *string                                             `json:"createdDateTime,omitempty"`
		Description                                               nullable.Type[string]                               `json:"description,omitempty"`
		DeviceManagementApplicabilityRuleDeviceMode               *DeviceManagementApplicabilityRuleDeviceMode        `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
		DeviceManagementApplicabilityRuleOsEdition                *DeviceManagementApplicabilityRuleOsEdition         `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
		DeviceManagementApplicabilityRuleOsVersion                *DeviceManagementApplicabilityRuleOsVersion         `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
		DeviceSettingStateSummaries                               *[]SettingStateDeviceSummary                        `json:"deviceSettingStateSummaries,omitempty"`
		DeviceStatusOverview                                      *DeviceConfigurationDeviceOverview                  `json:"deviceStatusOverview,omitempty"`
		DeviceStatuses                                            *[]DeviceConfigurationDeviceStatus                  `json:"deviceStatuses,omitempty"`
		DisplayName                                               *string                                             `json:"displayName,omitempty"`
		GroupAssignments                                          *[]DeviceConfigurationGroupAssignment               `json:"groupAssignments,omitempty"`
		LastModifiedDateTime                                      *string                                             `json:"lastModifiedDateTime,omitempty"`
		RoleScopeTagIds                                           *[]string                                           `json:"roleScopeTagIds,omitempty"`
		SupportsScopeTags                                         *bool                                               `json:"supportsScopeTags,omitempty"`
		UserStatusOverview                                        *DeviceConfigurationUserOverview                    `json:"userStatusOverview,omitempty"`
		UserStatuses                                              *[]DeviceConfigurationUserStatus                    `json:"userStatuses,omitempty"`
		Version                                                   *int64                                              `json:"version,omitempty"`
		Id                                                        *string                                             `json:"id,omitempty"`
		ODataId                                                   *string                                             `json:"@odata.id,omitempty"`
		ODataType                                                 *string                                             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.BackgroundDownloadFromHttpDelayInSeconds = decoded.BackgroundDownloadFromHttpDelayInSeconds
	s.CacheServerBackgroundDownloadFallbackToHttpDelayInSeconds = decoded.CacheServerBackgroundDownloadFallbackToHttpDelayInSeconds
	s.CacheServerForegroundDownloadFallbackToHttpDelayInSeconds = decoded.CacheServerForegroundDownloadFallbackToHttpDelayInSeconds
	s.CacheServerHostNames = decoded.CacheServerHostNames
	s.DeliveryOptimizationMode = decoded.DeliveryOptimizationMode
	s.ForegroundDownloadFromHttpDelayInSeconds = decoded.ForegroundDownloadFromHttpDelayInSeconds
	s.MaximumCacheAgeInDays = decoded.MaximumCacheAgeInDays
	s.MinimumBatteryPercentageAllowedToUpload = decoded.MinimumBatteryPercentageAllowedToUpload
	s.MinimumDiskSizeAllowedToPeerInGigabytes = decoded.MinimumDiskSizeAllowedToPeerInGigabytes
	s.MinimumFileSizeToCacheInMegabytes = decoded.MinimumFileSizeToCacheInMegabytes
	s.MinimumRamAllowedToPeerInGigabytes = decoded.MinimumRamAllowedToPeerInGigabytes
	s.ModifyCacheLocation = decoded.ModifyCacheLocation
	s.RestrictPeerSelectionBy = decoded.RestrictPeerSelectionBy
	s.VpnPeerCaching = decoded.VpnPeerCaching
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
		return fmt.Errorf("unmarshaling WindowsDeliveryOptimizationConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["bandwidthMode"]; ok {
		impl, err := UnmarshalDeliveryOptimizationBandwidthImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'BandwidthMode' for 'WindowsDeliveryOptimizationConfiguration': %+v", err)
		}
		s.BandwidthMode = impl
	}

	if v, ok := temp["groupIdSource"]; ok {
		impl, err := UnmarshalDeliveryOptimizationGroupIdSourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'GroupIdSource' for 'WindowsDeliveryOptimizationConfiguration': %+v", err)
		}
		s.GroupIdSource = impl
	}

	if v, ok := temp["maximumCacheSize"]; ok {
		impl, err := UnmarshalDeliveryOptimizationMaxCacheSizeImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'MaximumCacheSize' for 'WindowsDeliveryOptimizationConfiguration': %+v", err)
		}
		s.MaximumCacheSize = impl
	}

	return nil
}
