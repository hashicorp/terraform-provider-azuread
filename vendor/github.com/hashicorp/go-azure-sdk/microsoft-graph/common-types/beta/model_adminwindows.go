package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AdminWindows{}

type AdminWindows struct {
	// Entity that acts as a container for all Windows Update for Business deployment service functionalities. Read-only.
	Updates *AdminWindowsUpdates `json:"updates,omitempty"`

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

func (s AdminWindows) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AdminWindows{}

func (s AdminWindows) MarshalJSON() ([]byte, error) {
	type wrapper AdminWindows
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AdminWindows: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AdminWindows: %+v", err)
	}

	delete(decoded, "updates")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.adminWindows"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AdminWindows: %+v", err)
	}

	return encoded, nil
}
