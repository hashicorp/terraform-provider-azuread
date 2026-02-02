package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CopilotAdmin{}

type CopilotAdmin struct {
	// Set of Microsoft 365 Copilot settings that can be added or modified. Read-only. Nullable.
	Settings *CopilotAdminSetting `json:"settings,omitempty"`

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

func (s CopilotAdmin) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CopilotAdmin{}

func (s CopilotAdmin) MarshalJSON() ([]byte, error) {
	type wrapper CopilotAdmin
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CopilotAdmin: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CopilotAdmin: %+v", err)
	}

	delete(decoded, "settings")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.copilotAdmin"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CopilotAdmin: %+v", err)
	}

	return encoded, nil
}
