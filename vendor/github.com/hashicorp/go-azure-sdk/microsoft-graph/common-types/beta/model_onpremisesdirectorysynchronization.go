package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = OnPremisesDirectorySynchronization{}

type OnPremisesDirectorySynchronization struct {
	// Consists of configurations that can be fine-tuned and impact the on-premises directory synchronization process for a
	// tenant. Nullable.
	Configuration *OnPremisesDirectorySynchronizationConfiguration `json:"configuration,omitempty"`

	Features *OnPremisesDirectorySynchronizationFeature `json:"features,omitempty"`

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

func (s OnPremisesDirectorySynchronization) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OnPremisesDirectorySynchronization{}

func (s OnPremisesDirectorySynchronization) MarshalJSON() ([]byte, error) {
	type wrapper OnPremisesDirectorySynchronization
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OnPremisesDirectorySynchronization: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OnPremisesDirectorySynchronization: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.onPremisesDirectorySynchronization"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OnPremisesDirectorySynchronization: %+v", err)
	}

	return encoded, nil
}
