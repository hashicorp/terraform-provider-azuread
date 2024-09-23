package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SynchronizationSchema{}

type SynchronizationSchema struct {
	// Contains the collection of directories and all of their objects.
	Directories *[]DirectoryDefinition `json:"directories,omitempty"`

	// A collection of synchronization rules configured for the synchronizationJob or synchronizationTemplate.
	SynchronizationRules *[]SynchronizationRule `json:"synchronizationRules,omitempty"`

	// The version of the schema, updated automatically with every schema change.
	Version nullable.Type[string] `json:"version,omitempty"`

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

func (s SynchronizationSchema) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SynchronizationSchema{}

func (s SynchronizationSchema) MarshalJSON() ([]byte, error) {
	type wrapper SynchronizationSchema
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SynchronizationSchema: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SynchronizationSchema: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.synchronizationSchema"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SynchronizationSchema: %+v", err)
	}

	return encoded, nil
}
