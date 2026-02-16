package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessEntitiesSummary struct {
	// The number of unique devices that were seen.
	DeviceCount *int64 `json:"deviceCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	TrafficType *NetworkaccessTrafficType `json:"trafficType,omitempty"`

	// The number of unique Microsoft Entra ID users that were seen.
	UserCount *int64 `json:"userCount,omitempty"`

	// The number of unique target workloads/hosts that were seen.
	WorkloadCount *int64 `json:"workloadCount,omitempty"`
}
