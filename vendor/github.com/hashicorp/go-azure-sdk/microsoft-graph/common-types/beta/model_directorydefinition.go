package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DirectoryDefinition{}

type DirectoryDefinition struct {
	Discoverabilities *DirectoryDefinitionDiscoverabilities `json:"discoverabilities,omitempty"`

	// Represents the discovery date and time using ISO 8601 format and is always in UTC time. For example, midnight UTC on
	// Jan 1, 2014 is 2014-01-01T00:00:00Z.
	DiscoveryDateTime nullable.Type[string] `json:"discoveryDateTime,omitempty"`

	// Name of the directory. Must be unique within the synchronization schema. Not nullable.
	Name nullable.Type[string] `json:"name,omitempty"`

	// Collection of objects supported by the directory.
	Objects *[]ObjectDefinition `json:"objects,omitempty"`

	// Whether this object is read-only.
	ReadOnly *bool `json:"readOnly,omitempty"`

	// Read only value that indicates version discovered. null if discovery hasn't yet occurred.
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

func (s DirectoryDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DirectoryDefinition{}

func (s DirectoryDefinition) MarshalJSON() ([]byte, error) {
	type wrapper DirectoryDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DirectoryDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DirectoryDefinition: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.directoryDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DirectoryDefinition: %+v", err)
	}

	return encoded, nil
}
