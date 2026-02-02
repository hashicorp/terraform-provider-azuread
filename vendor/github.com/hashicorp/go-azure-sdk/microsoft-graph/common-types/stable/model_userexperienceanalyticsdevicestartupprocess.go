package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsDeviceStartupProcess{}

type UserExperienceAnalyticsDeviceStartupProcess struct {
	// The Intune device id of the device. Supports: $select, $OrderBy. Read-only.
	ManagedDeviceId nullable.Type[string] `json:"managedDeviceId,omitempty"`

	// The name of the process. Examples: outlook, excel. Supports: $select, $OrderBy. Read-only.
	ProcessName nullable.Type[string] `json:"processName,omitempty"`

	// The product name of the process. Examples: Microsoft Outlook, Microsoft Excel. Supports: $select, $OrderBy.
	// Read-only.
	ProductName nullable.Type[string] `json:"productName,omitempty"`

	// The publisher of the process. Examples: Microsoft Corporation, Contoso Corp. Supports: $select, $OrderBy. Read-only.
	Publisher nullable.Type[string] `json:"publisher,omitempty"`

	// The impact of startup process on device boot time in milliseconds. Supports: $select, $OrderBy. Read-only.
	StartupImpactInMs *int64 `json:"startupImpactInMs,omitempty"`

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

func (s UserExperienceAnalyticsDeviceStartupProcess) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsDeviceStartupProcess{}

func (s UserExperienceAnalyticsDeviceStartupProcess) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsDeviceStartupProcess
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsDeviceStartupProcess: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsDeviceStartupProcess: %+v", err)
	}

	delete(decoded, "managedDeviceId")
	delete(decoded, "processName")
	delete(decoded, "productName")
	delete(decoded, "publisher")
	delete(decoded, "startupImpactInMs")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsDeviceStartupProcess"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsDeviceStartupProcess: %+v", err)
	}

	return encoded, nil
}
