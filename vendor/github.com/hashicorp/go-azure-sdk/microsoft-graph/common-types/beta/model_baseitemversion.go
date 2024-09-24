package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BaseItemVersion interface {
	Entity
	BaseItemVersion() BaseBaseItemVersionImpl
}

var _ BaseItemVersion = BaseBaseItemVersionImpl{}

type BaseBaseItemVersionImpl struct {
	// Identity of the user that last modified the version. Read-only.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// Date and time when the version was last modified. Read-only.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Indicates the publication status of this particular version. Read-only.
	Publication *PublicationFacet `json:"publication,omitempty"`

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

func (s BaseBaseItemVersionImpl) BaseItemVersion() BaseBaseItemVersionImpl {
	return s
}

func (s BaseBaseItemVersionImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ BaseItemVersion = RawBaseItemVersionImpl{}

// RawBaseItemVersionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawBaseItemVersionImpl struct {
	baseItemVersion BaseBaseItemVersionImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawBaseItemVersionImpl) BaseItemVersion() BaseBaseItemVersionImpl {
	return s.baseItemVersion
}

func (s RawBaseItemVersionImpl) Entity() BaseEntityImpl {
	return s.baseItemVersion.Entity()
}

var _ json.Marshaler = BaseBaseItemVersionImpl{}

func (s BaseBaseItemVersionImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseBaseItemVersionImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseBaseItemVersionImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseBaseItemVersionImpl: %+v", err)
	}

	delete(decoded, "lastModifiedBy")
	delete(decoded, "lastModifiedDateTime")
	delete(decoded, "publication")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.baseItemVersion"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseBaseItemVersionImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseBaseItemVersionImpl{}

func (s *BaseBaseItemVersionImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		Publication          *PublicationFacet     `json:"publication,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Publication = decoded.Publication
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseBaseItemVersionImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BaseBaseItemVersionImpl': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	return nil
}

func UnmarshalBaseItemVersionImplementation(input []byte) (BaseItemVersion, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseItemVersion into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.driveItemVersion") {
		var out DriveItemVersion
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DriveItemVersion: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.listItemVersion") {
		var out ListItemVersion
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ListItemVersion: %+v", err)
		}
		return out, nil
	}

	var parent BaseBaseItemVersionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseBaseItemVersionImpl: %+v", err)
	}

	return RawBaseItemVersionImpl{
		baseItemVersion: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
