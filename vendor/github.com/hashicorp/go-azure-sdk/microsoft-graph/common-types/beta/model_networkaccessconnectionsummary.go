package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessConnectionSummary struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Total number of connections for the specified traffic type.
	TotalCount *int64 `json:"totalCount,omitempty"`

	TrafficType *NetworkaccessTrafficType `json:"trafficType,omitempty"`
}
