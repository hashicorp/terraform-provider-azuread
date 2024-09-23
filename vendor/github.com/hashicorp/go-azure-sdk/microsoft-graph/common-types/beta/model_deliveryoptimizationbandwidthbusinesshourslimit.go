package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeliveryOptimizationBandwidthBusinessHoursLimit struct {
	// Specifies the beginning of business hours using a 24-hour clock (0-23). Valid values 0 to 23
	BandwidthBeginBusinessHours *int64 `json:"bandwidthBeginBusinessHours,omitempty"`

	// Specifies the end of business hours using a 24-hour clock (0-23). Valid values 0 to 23
	BandwidthEndBusinessHours *int64 `json:"bandwidthEndBusinessHours,omitempty"`

	// Specifies the percentage of bandwidth to limit during business hours (0-100). Valid values 0 to 100
	BandwidthPercentageDuringBusinessHours *int64 `json:"bandwidthPercentageDuringBusinessHours,omitempty"`

	// Specifies the percentage of bandwidth to limit outsidse business hours (0-100). Valid values 0 to 100
	BandwidthPercentageOutsideBusinessHours *int64 `json:"bandwidthPercentageOutsideBusinessHours,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
