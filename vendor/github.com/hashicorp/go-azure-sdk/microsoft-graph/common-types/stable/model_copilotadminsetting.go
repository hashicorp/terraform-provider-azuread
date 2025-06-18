package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CopilotAdminSetting{}

type CopilotAdminSetting struct {
	// Represents a setting that controls whether users of Microsoft 365 Copilot in Teams meetings can receive responses to
	// sentiment-related prompts. Read-only. Nullable.
	LimitedMode *CopilotAdminLimitedMode `json:"limitedMode,omitempty"`

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

func (s CopilotAdminSetting) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CopilotAdminSetting{}

func (s CopilotAdminSetting) MarshalJSON() ([]byte, error) {
	type wrapper CopilotAdminSetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CopilotAdminSetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CopilotAdminSetting: %+v", err)
	}

	delete(decoded, "limitedMode")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.copilotAdminSetting"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CopilotAdminSetting: %+v", err)
	}

	return encoded, nil
}
