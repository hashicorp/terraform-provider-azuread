package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCRestorePointSetting struct {
	// The time interval in hours to take snapshots (restore points) of a Cloud PC automatically. Possible values are 4, 6,
	// 12, 16, and 24. The default frequency is 12 hours. The frequencyInHours property is deprecated and will stop
	// returning data on January 31, 2024. Going forward, use the frequencyType property.
	FrequencyInHours nullable.Type[int64] `json:"frequencyInHours,omitempty"`

	// The time interval in hours to take snapshots (restore points) of a Cloud PC automatically. Possible values are:
	// default, fourHours, sixHours, twelveHours, sixteenHours, twentyFourHours, unknownFutureValue. The default value is
	// default that indicates that the time interval for automatic capturing of restore point snapshots is set to 12 hours.
	FrequencyType *CloudPCRestorePointFrequencyType `json:"frequencyType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// If true, the user has the ability to use snapshots to restore Cloud PCs. If false, non-admin users can't use
	// snapshots to restore the Cloud PC.
	UserRestoreEnabled nullable.Type[bool] `json:"userRestoreEnabled,omitempty"`
}
