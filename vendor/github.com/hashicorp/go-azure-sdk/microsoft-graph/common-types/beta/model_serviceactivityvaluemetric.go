package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceActivityValueMetric struct {
	// The starting date and time (UTC) of the interval.
	IntervalStartDateTime *string `json:"intervalStartDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The aggregated value over the given aggregation interval starting from the intervalStartDateTime. The value is
	// caculated at the minute level. The value at the starting minute of the intervalStartDateTime is included. The value
	// at the last minute of the given interval is excluded. For example, if intervalStartDateTime is 2023-09-20T18:00:00Z
	// and aggregation interval is 5 minutes, then the value is aggregated from 2023-09-20T18:00:00Z(inclusive) to
	// 2023-09-20T18:05:00Z(exclusive).
	Value *int64 `json:"value,omitempty"`
}
