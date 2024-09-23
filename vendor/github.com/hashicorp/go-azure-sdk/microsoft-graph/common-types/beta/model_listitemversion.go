package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListItemVersion interface {
	Entity
	BaseItemVersion
	ListItemVersion() BaseListItemVersionImpl
}

var _ ListItemVersion = BaseListItemVersionImpl{}

type BaseListItemVersionImpl struct {
	// A collection of the fields and values for this version of the list item.
	Fields *FieldValueSet `json:"fields,omitempty"`

	// Fields inherited from BaseItemVersion

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

func (s BaseListItemVersionImpl) ListItemVersion() BaseListItemVersionImpl {
	return s
}

func (s BaseListItemVersionImpl) BaseItemVersion() BaseBaseItemVersionImpl {
	return BaseBaseItemVersionImpl{
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Publication:          s.Publication,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s BaseListItemVersionImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ListItemVersion = RawListItemVersionImpl{}

// RawListItemVersionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawListItemVersionImpl struct {
	listItemVersion BaseListItemVersionImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawListItemVersionImpl) ListItemVersion() BaseListItemVersionImpl {
	return s.listItemVersion
}

func (s RawListItemVersionImpl) BaseItemVersion() BaseBaseItemVersionImpl {
	return s.listItemVersion.BaseItemVersion()
}

func (s RawListItemVersionImpl) Entity() BaseEntityImpl {
	return s.listItemVersion.Entity()
}

var _ json.Marshaler = BaseListItemVersionImpl{}

func (s BaseListItemVersionImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseListItemVersionImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseListItemVersionImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseListItemVersionImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.listItemVersion"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseListItemVersionImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseListItemVersionImpl{}

func (s *BaseListItemVersionImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Fields               *FieldValueSet        `json:"fields,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		Publication          *PublicationFacet     `json:"publication,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Fields = decoded.Fields
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Publication = decoded.Publication

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseListItemVersionImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BaseListItemVersionImpl': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	return nil
}

func UnmarshalListItemVersionImplementation(input []byte) (ListItemVersion, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ListItemVersion into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.documentSetVersion") {
		var out DocumentSetVersion
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DocumentSetVersion: %+v", err)
		}
		return out, nil
	}

	var parent BaseListItemVersionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseListItemVersionImpl: %+v", err)
	}

	return RawListItemVersionImpl{
		listItemVersion: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
