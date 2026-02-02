package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetricTimeSeriesDataPoint struct {
	// Time of the metric time series data point
	DateTime *string `json:"dateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Value of the metric time series data point
	Value *int64 `json:"value,omitempty"`
}
