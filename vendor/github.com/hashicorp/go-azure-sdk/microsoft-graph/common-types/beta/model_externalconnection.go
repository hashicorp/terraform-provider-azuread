package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ExternalConnection{}

type ExternalConnection struct {
	Configuration *Configuration         `json:"configuration,omitempty"`
	Description   nullable.Type[string]  `json:"description,omitempty"`
	Groups        *[]ExternalGroup       `json:"groups,omitempty"`
	Items         *[]ExternalItem        `json:"items,omitempty"`
	Name          nullable.Type[string]  `json:"name,omitempty"`
	Operations    *[]ConnectionOperation `json:"operations,omitempty"`
	Schema        *Schema                `json:"schema,omitempty"`
	State         *ConnectionState       `json:"state,omitempty"`

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

func (s ExternalConnection) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ExternalConnection{}

func (s ExternalConnection) MarshalJSON() ([]byte, error) {
	type wrapper ExternalConnection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExternalConnection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExternalConnection: %+v", err)
	}

	delete(decoded, "state")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.externalConnection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExternalConnection: %+v", err)
	}

	return encoded, nil
}
