package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsDeviceStartupProcessPerformance{}

type UserExperienceAnalyticsDeviceStartupProcessPerformance struct {
	// The count of devices which initiated this process on startup. Supports: $filter, $select, $OrderBy. Read-only.
	DeviceCount *int64 `json:"deviceCount,omitempty"`

	// The median impact of startup process on device boot time in milliseconds. Supports: $filter, $select, $OrderBy.
	// Read-only.
	MedianImpactInMs *int64 `json:"medianImpactInMs,omitempty"`

	// The name of the startup process. Examples: outlook, excel. Supports: $select, $OrderBy. Read-only.
	ProcessName nullable.Type[string] `json:"processName,omitempty"`

	// The product name of the startup process. Examples: Microsoft Outlook, Microsoft Excel. Supports: $select, $OrderBy.
	// Read-only.
	ProductName nullable.Type[string] `json:"productName,omitempty"`

	// The publisher of the startup process. Examples: Microsoft Corporation, Contoso Corp. Supports: $select, $OrderBy.
	// Read-only.
	Publisher nullable.Type[string] `json:"publisher,omitempty"`

	// The total impact of startup process on device boot time in milliseconds. Supports: $filter, $select, $OrderBy.
	// Read-only.
	TotalImpactInMs *int64 `json:"totalImpactInMs,omitempty"`

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

func (s UserExperienceAnalyticsDeviceStartupProcessPerformance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsDeviceStartupProcessPerformance{}

func (s UserExperienceAnalyticsDeviceStartupProcessPerformance) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsDeviceStartupProcessPerformance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsDeviceStartupProcessPerformance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsDeviceStartupProcessPerformance: %+v", err)
	}

	delete(decoded, "deviceCount")
	delete(decoded, "medianImpactInMs")
	delete(decoded, "processName")
	delete(decoded, "productName")
	delete(decoded, "publisher")
	delete(decoded, "totalImpactInMs")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsDeviceStartupProcessPerformance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsDeviceStartupProcessPerformance: %+v", err)
	}

	return encoded, nil
}
