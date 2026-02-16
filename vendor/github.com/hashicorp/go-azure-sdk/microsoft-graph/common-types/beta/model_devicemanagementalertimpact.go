package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAlertImpact struct {
	// The aggregation type of the impact. The possible values are: count, percentage, affectedCloudPcCount,
	// affectedCloudPcPercentage, unknownFutureValue.
	AggregationType *DeviceManagementAggregationType `json:"aggregationType,omitempty"`

	// The detail information of the impact. For example, if the Frontline Cloud PCs near concurrency limit alert is
	// triggered, the details contain the impacted Frontline license SKU name, such as Windows 365 Frontline 2
	// vCPU/8GB/128GB, and the corresponding impacted value.
	AlertImpactDetails *[]KeyValuePair `json:"alertImpactDetails,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The number value of the impact. For the aggregation types of count and affectedCloudPcCount, the value indicates the
	// number of affected instances. For example, 6 affectedCloudPcCount means that six Cloud PCs are affected. For the
	// aggregation types of percentage and affectedCloudPcPercentage, the value indicates the percent of affected instances.
	// For example, 12 affectedCloudPcPercentage means that 12% of Cloud PCs are affected.
	Value nullable.Type[int64] `json:"value,omitempty"`
}
