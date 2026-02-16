package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ExternalConnectorsSchema{}

type ExternalConnectorsSchema struct {
	// Must be set to microsoft.graph.externalConnector.externalItem. Required.
	BaseType string `json:"baseType"`

	// The properties defined for the items in the connection. The minimum number of properties is one, the maximum is 128.
	Properties *[]ExternalConnectorsProperty `json:"properties,omitempty"`

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

func (s ExternalConnectorsSchema) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ExternalConnectorsSchema{}

func (s ExternalConnectorsSchema) MarshalJSON() ([]byte, error) {
	type wrapper ExternalConnectorsSchema
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExternalConnectorsSchema: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExternalConnectorsSchema: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.externalConnectors.schema"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExternalConnectorsSchema: %+v", err)
	}

	return encoded, nil
}
