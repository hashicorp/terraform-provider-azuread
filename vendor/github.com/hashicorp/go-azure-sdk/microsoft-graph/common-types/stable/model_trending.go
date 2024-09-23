package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Trending{}

type Trending struct {
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Used for navigating to the trending document.
	Resource Entity `json:"resource"`

	// Reference properties of the trending document, such as the url and type of the document.
	ResourceReference *ResourceReference `json:"resourceReference,omitempty"`

	// Properties that you can use to visualize the document in your experience.
	ResourceVisualization *ResourceVisualization `json:"resourceVisualization,omitempty"`

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

func (s Trending) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Trending{}

func (s Trending) MarshalJSON() ([]byte, error) {
	type wrapper Trending
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Trending: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Trending: %+v", err)
	}

	delete(decoded, "resourceReference")
	delete(decoded, "resourceVisualization")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.trending"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Trending: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Trending{}

func (s *Trending) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		LastModifiedDateTime  nullable.Type[string]  `json:"lastModifiedDateTime,omitempty"`
		ResourceReference     *ResourceReference     `json:"resourceReference,omitempty"`
		ResourceVisualization *ResourceVisualization `json:"resourceVisualization,omitempty"`
		Id                    *string                `json:"id,omitempty"`
		ODataId               *string                `json:"@odata.id,omitempty"`
		ODataType             *string                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ResourceReference = decoded.ResourceReference
	s.ResourceVisualization = decoded.ResourceVisualization
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Trending into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["resource"]; ok {
		impl, err := UnmarshalEntityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Resource' for 'Trending': %+v", err)
		}
		s.Resource = impl
	}

	return nil
}
