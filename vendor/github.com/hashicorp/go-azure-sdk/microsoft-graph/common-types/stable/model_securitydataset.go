package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDataSet interface {
	Entity
	SecurityDataSet() BaseSecurityDataSetImpl
}

var _ SecurityDataSet = BaseSecurityDataSetImpl{}

type BaseSecurityDataSetImpl struct {
	// The user who created the data set. Read-only.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`

	// The date and time when the review set was created. The timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The description of the data set.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The name of the data set. The name is unique with a maximum limit of 64 characters.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

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

func (s BaseSecurityDataSetImpl) SecurityDataSet() BaseSecurityDataSetImpl {
	return s
}

func (s BaseSecurityDataSetImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ SecurityDataSet = RawSecurityDataSetImpl{}

// RawSecurityDataSetImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityDataSetImpl struct {
	securityDataSet BaseSecurityDataSetImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawSecurityDataSetImpl) SecurityDataSet() BaseSecurityDataSetImpl {
	return s.securityDataSet
}

func (s RawSecurityDataSetImpl) Entity() BaseEntityImpl {
	return s.securityDataSet.Entity()
}

var _ json.Marshaler = BaseSecurityDataSetImpl{}

func (s BaseSecurityDataSetImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseSecurityDataSetImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseSecurityDataSetImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseSecurityDataSetImpl: %+v", err)
	}

	delete(decoded, "createdBy")
	delete(decoded, "createdDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.dataSet"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseSecurityDataSetImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseSecurityDataSetImpl{}

func (s *BaseSecurityDataSetImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`
		Description     nullable.Type[string] `json:"description,omitempty"`
		DisplayName     nullable.Type[string] `json:"displayName,omitempty"`
		Id              *string               `json:"id,omitempty"`
		ODataId         *string               `json:"@odata.id,omitempty"`
		ODataType       *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseSecurityDataSetImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'BaseSecurityDataSetImpl': %+v", err)
		}
		s.CreatedBy = &impl
	}

	return nil
}

func UnmarshalSecurityDataSetImplementation(input []byte) (SecurityDataSet, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityDataSet into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryReviewSet") {
		var out SecurityEdiscoveryReviewSet
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryReviewSet: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityDataSetImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityDataSetImpl: %+v", err)
	}

	return RawSecurityDataSetImpl{
		securityDataSet: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
