package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamedLocation interface {
	Entity
	NamedLocation() BaseNamedLocationImpl
}

var _ NamedLocation = BaseNamedLocationImpl{}

type BaseNamedLocationImpl struct {
	// The Timestamp type represents creation date and time of the location using ISO 8601 format and is always in UTC time.
	// For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Human-readable name of the location.
	DisplayName *string `json:"displayName,omitempty"`

	// The Timestamp type represents last modified date and time of the location using ISO 8601 format and is always in UTC
	// time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	ModifiedDateTime nullable.Type[string] `json:"modifiedDateTime,omitempty"`

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

func (s BaseNamedLocationImpl) NamedLocation() BaseNamedLocationImpl {
	return s
}

func (s BaseNamedLocationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ NamedLocation = RawNamedLocationImpl{}

// RawNamedLocationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawNamedLocationImpl struct {
	namedLocation BaseNamedLocationImpl
	Type          string
	Values        map[string]interface{}
}

func (s RawNamedLocationImpl) NamedLocation() BaseNamedLocationImpl {
	return s.namedLocation
}

func (s RawNamedLocationImpl) Entity() BaseEntityImpl {
	return s.namedLocation.Entity()
}

var _ json.Marshaler = BaseNamedLocationImpl{}

func (s BaseNamedLocationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseNamedLocationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseNamedLocationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseNamedLocationImpl: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "modifiedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.namedLocation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseNamedLocationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalNamedLocationImplementation(input []byte) (NamedLocation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling NamedLocation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.countryNamedLocation") {
		var out CountryNamedLocation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CountryNamedLocation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ipNamedLocation") {
		var out IPNamedLocation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IPNamedLocation: %+v", err)
		}
		return out, nil
	}

	var parent BaseNamedLocationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseNamedLocationImpl: %+v", err)
	}

	return RawNamedLocationImpl{
		namedLocation: parent,
		Type:          value,
		Values:        temp,
	}, nil

}
