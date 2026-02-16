package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AndroidManagedStoreAppConfigurationSchema{}

type AndroidManagedStoreAppConfigurationSchema struct {
	// UTF8 encoded byte array containing example JSON string conforming to this schema that demonstrates how to set the
	// configuration for this app
	ExampleJson nullable.Type[string] `json:"exampleJson,omitempty"`

	// Collection of items each representing a named configuration option in the schema. It contains a flat list of all
	// configuration.
	NestedSchemaItems *[]AndroidManagedStoreAppConfigurationSchemaItem `json:"nestedSchemaItems,omitempty"`

	// Collection of items each representing a named configuration option in the schema. It only contains the root-level
	// configuration.
	SchemaItems *[]AndroidManagedStoreAppConfigurationSchemaItem `json:"schemaItems,omitempty"`

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

func (s AndroidManagedStoreAppConfigurationSchema) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidManagedStoreAppConfigurationSchema{}

func (s AndroidManagedStoreAppConfigurationSchema) MarshalJSON() ([]byte, error) {
	type wrapper AndroidManagedStoreAppConfigurationSchema
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidManagedStoreAppConfigurationSchema: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidManagedStoreAppConfigurationSchema: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidManagedStoreAppConfigurationSchema"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidManagedStoreAppConfigurationSchema: %+v", err)
	}

	return encoded, nil
}
