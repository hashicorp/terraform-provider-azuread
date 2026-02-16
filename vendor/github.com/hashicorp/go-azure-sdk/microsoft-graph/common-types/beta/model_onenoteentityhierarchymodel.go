package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteEntityHierarchyModel interface {
	Entity
	OnenoteEntityBaseModel
	OnenoteEntitySchemaObjectModel
	OnenoteEntityHierarchyModel() BaseOnenoteEntityHierarchyModelImpl
}

var _ OnenoteEntityHierarchyModel = BaseOnenoteEntityHierarchyModelImpl{}

type BaseOnenoteEntityHierarchyModelImpl struct {
	CreatedBy            IdentitySet           `json:"createdBy"`
	DisplayName          nullable.Type[string] `json:"displayName,omitempty"`
	LastModifiedBy       IdentitySet           `json:"lastModifiedBy"`
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Fields inherited from OnenoteEntitySchemaObjectModel

	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Fields inherited from OnenoteEntityBaseModel

	Self nullable.Type[string] `json:"self,omitempty"`

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

func (s BaseOnenoteEntityHierarchyModelImpl) OnenoteEntityHierarchyModel() BaseOnenoteEntityHierarchyModelImpl {
	return s
}

func (s BaseOnenoteEntityHierarchyModelImpl) OnenoteEntitySchemaObjectModel() BaseOnenoteEntitySchemaObjectModelImpl {
	return BaseOnenoteEntitySchemaObjectModelImpl{
		CreatedDateTime: s.CreatedDateTime,
		Self:            s.Self,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s BaseOnenoteEntityHierarchyModelImpl) OnenoteEntityBaseModel() BaseOnenoteEntityBaseModelImpl {
	return BaseOnenoteEntityBaseModelImpl{
		Self:      s.Self,
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s BaseOnenoteEntityHierarchyModelImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ OnenoteEntityHierarchyModel = RawOnenoteEntityHierarchyModelImpl{}

// RawOnenoteEntityHierarchyModelImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawOnenoteEntityHierarchyModelImpl struct {
	onenoteEntityHierarchyModel BaseOnenoteEntityHierarchyModelImpl
	Type                        string
	Values                      map[string]interface{}
}

func (s RawOnenoteEntityHierarchyModelImpl) OnenoteEntityHierarchyModel() BaseOnenoteEntityHierarchyModelImpl {
	return s.onenoteEntityHierarchyModel
}

func (s RawOnenoteEntityHierarchyModelImpl) OnenoteEntitySchemaObjectModel() BaseOnenoteEntitySchemaObjectModelImpl {
	return s.onenoteEntityHierarchyModel.OnenoteEntitySchemaObjectModel()
}

func (s RawOnenoteEntityHierarchyModelImpl) OnenoteEntityBaseModel() BaseOnenoteEntityBaseModelImpl {
	return s.onenoteEntityHierarchyModel.OnenoteEntityBaseModel()
}

func (s RawOnenoteEntityHierarchyModelImpl) Entity() BaseEntityImpl {
	return s.onenoteEntityHierarchyModel.Entity()
}

var _ json.Marshaler = BaseOnenoteEntityHierarchyModelImpl{}

func (s BaseOnenoteEntityHierarchyModelImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseOnenoteEntityHierarchyModelImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseOnenoteEntityHierarchyModelImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseOnenoteEntityHierarchyModelImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.onenoteEntityHierarchyModel"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseOnenoteEntityHierarchyModelImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseOnenoteEntityHierarchyModelImpl{}

func (s *BaseOnenoteEntityHierarchyModelImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DisplayName          nullable.Type[string] `json:"displayName,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		CreatedDateTime      nullable.Type[string] `json:"createdDateTime,omitempty"`
		Self                 nullable.Type[string] `json:"self,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DisplayName = decoded.DisplayName
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Self = decoded.Self

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseOnenoteEntityHierarchyModelImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'BaseOnenoteEntityHierarchyModelImpl': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BaseOnenoteEntityHierarchyModelImpl': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}

func UnmarshalOnenoteEntityHierarchyModelImplementation(input []byte) (OnenoteEntityHierarchyModel, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling OnenoteEntityHierarchyModel into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.notebook") {
		var out Notebook
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Notebook: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onenoteSection") {
		var out OnenoteSection
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnenoteSection: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sectionGroup") {
		var out SectionGroup
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SectionGroup: %+v", err)
		}
		return out, nil
	}

	var parent BaseOnenoteEntityHierarchyModelImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOnenoteEntityHierarchyModelImpl: %+v", err)
	}

	return RawOnenoteEntityHierarchyModelImpl{
		onenoteEntityHierarchyModel: parent,
		Type:                        value,
		Values:                      temp,
	}, nil

}
