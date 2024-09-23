package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaDeploymentSettings struct {
	// Minimum battery level (%) required for both download and installation. Default: -1 (System defaults). Maximum is 100.
	BatteryRuleMinimumBatteryLevelPercentage nullable.Type[int64] `json:"batteryRuleMinimumBatteryLevelPercentage,omitempty"`

	// Flag indicating if charger is required. When set to false, the client can install updates whether the device is in or
	// out of the charger. Applied only for installation. Defaults to false.
	BatteryRuleRequireCharger nullable.Type[bool] `json:"batteryRuleRequireCharger,omitempty"`

	// Deploy update for devices with this model only.
	DeviceModel *string `json:"deviceModel,omitempty"`

	// Represents various network types for Zebra FOTA deployment.
	DownloadRuleNetworkType *ZebraFotaNetworkType `json:"downloadRuleNetworkType,omitempty"`

	// Date and time in the device time zone when the download will start (e.g., 2018-07-25T10:20:32). The default value is
	// UTC now and the maximum is 10 days from deployment creation.
	DownloadRuleStartDateTime nullable.Type[string] `json:"downloadRuleStartDateTime,omitempty"`

	// A description provided by Zebra for the the firmware artifact to update the device to (e.g.: LifeGuard Update 120
	// (released 29-June-2022).
	FirmwareTargetArtifactDescription nullable.Type[string] `json:"firmwareTargetArtifactDescription,omitempty"`

	// Deployment's Board Support Package (BSP. E.g.: '01.18.02.00'). Required only for custom update type.
	FirmwareTargetBoardSupportPackageVersion nullable.Type[string] `json:"firmwareTargetBoardSupportPackageVersion,omitempty"`

	// Target OS Version (e.g.: '8.1.0'). Required only for custom update type.
	FirmwareTargetOsVersion nullable.Type[string] `json:"firmwareTargetOsVersion,omitempty"`

	// Target patch name (e.g.: 'U06'). Required only for custom update type.
	FirmwareTargetPatch nullable.Type[string] `json:"firmwareTargetPatch,omitempty"`

	// Date and time in device time zone when the install will start. Default - download startDate if configured, otherwise
	// defaults to NOW. Ignored when deployment update type was set to auto.
	InstallRuleStartDateTime nullable.Type[string] `json:"installRuleStartDateTime,omitempty"`

	// Time of day after which the install cannot start. Possible range is 00:30:00 to 23:59:59. Should be greater than
	// 'installRuleWindowStartTime' by 30 mins. The time is expressed in a 24-hour format, as hh:mm, and is in the device
	// time zone. Default - 23:59:59. Respected for all values of update type, including AUTO.
	InstallRuleWindowEndTime nullable.Type[string] `json:"installRuleWindowEndTime,omitempty"`

	// Time of day (00:00:00 - 23:30:00) when installation should begin. The time is expressed in a 24-hour format, as
	// hh:mm, and is in the device time zone. Default - 00:00:00. Respected for all values of update type, including AUTO.
	InstallRuleWindowStartTime nullable.Type[string] `json:"installRuleWindowStartTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Maximum 28 days. Default is 28 days. Sequence of dates are: 1) Download start date. 2) Install start date. 3)
	// Schedule end date. If any of the values are not provided, the date provided in the preceding step of the sequence is
	// used. If no values are provided, the string value of the current UTC is used.
	ScheduleDurationInDays nullable.Type[int64] `json:"scheduleDurationInDays,omitempty"`

	// Represents various schedule modes for Zebra FOTA deployment.
	ScheduleMode *ZebraFotaScheduleMode `json:"scheduleMode,omitempty"`

	// This attribute indicates the deployment time offset (e.g.180 represents an offset of +03:00, and -270 represents an
	// offset of -04:30). The time offset is the time timezone where the devices are located. The deployment start and end
	// data uses this timezone
	TimeZoneOffsetInMinutes nullable.Type[int64] `json:"timeZoneOffsetInMinutes,omitempty"`

	// Represents various update types for Zebra FOTA deployment.
	UpdateType *ZebraFotaUpdateType `json:"updateType,omitempty"`
}
