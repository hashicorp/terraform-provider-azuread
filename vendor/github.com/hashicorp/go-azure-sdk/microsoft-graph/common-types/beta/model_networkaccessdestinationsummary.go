package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessDestinationSummary struct {
	// The number of the destinationSummary objects, aggregated by Global Secure Access service.
	Count *int64 `json:"count,omitempty"`

	// The IP address or FQDN of the destination.
	Destination *string `json:"destination,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The traffic classification. The allowed values are internet, private, microsoft365, all, and unknownFutureValue.
	TrafficType *NetworkaccessTrafficType `json:"trafficType,omitempty"`
}
