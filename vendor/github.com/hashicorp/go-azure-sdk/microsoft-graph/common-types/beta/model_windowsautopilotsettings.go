package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsAutopilotSettings{}

type WindowsAutopilotSettings struct {
	// Last data sync date time with DDS service.
	LastManualSyncTriggerDateTime *string `json:"lastManualSyncTriggerDateTime,omitempty"`

	// Last data sync date time with DDS service.
	LastSyncDateTime *string `json:"lastSyncDateTime,omitempty"`

	SyncStatus *WindowsAutopilotSyncStatus `json:"syncStatus,omitempty"`

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

func (s WindowsAutopilotSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsAutopilotSettings{}

func (s WindowsAutopilotSettings) MarshalJSON() ([]byte, error) {
	type wrapper WindowsAutopilotSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsAutopilotSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsAutopilotSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsAutopilotSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsAutopilotSettings: %+v", err)
	}

	return encoded, nil
}
