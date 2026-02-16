package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCase interface {
	Entity
	SecurityCase() BaseSecurityCaseImpl
}

var _ SecurityCase = BaseSecurityCaseImpl{}

type BaseSecurityCaseImpl struct {
	CreatedDateTime      nullable.Type[string] `json:"createdDateTime,omitempty"`
	Description          nullable.Type[string] `json:"description,omitempty"`
	DisplayName          nullable.Type[string] `json:"displayName,omitempty"`
	LastModifiedBy       IdentitySet           `json:"lastModifiedBy"`
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
	Status               *SecurityCaseStatus   `json:"status,omitempty"`

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

func (s BaseSecurityCaseImpl) SecurityCase() BaseSecurityCaseImpl {
	return s
}

func (s BaseSecurityCaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ SecurityCase = RawSecurityCaseImpl{}

// RawSecurityCaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityCaseImpl struct {
	securityCase BaseSecurityCaseImpl
	Type         string
	Values       map[string]interface{}
}

func (s RawSecurityCaseImpl) SecurityCase() BaseSecurityCaseImpl {
	return s.securityCase
}

func (s RawSecurityCaseImpl) Entity() BaseEntityImpl {
	return s.securityCase.Entity()
}

var _ json.Marshaler = BaseSecurityCaseImpl{}

func (s BaseSecurityCaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseSecurityCaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseSecurityCaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseSecurityCaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.case"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseSecurityCaseImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseSecurityCaseImpl{}

func (s *BaseSecurityCaseImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime      nullable.Type[string] `json:"createdDateTime,omitempty"`
		Description          nullable.Type[string] `json:"description,omitempty"`
		DisplayName          nullable.Type[string] `json:"displayName,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		Status               *SecurityCaseStatus   `json:"status,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseSecurityCaseImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BaseSecurityCaseImpl': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}

func UnmarshalSecurityCaseImplementation(input []byte) (SecurityCase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityCase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryCase") {
		var out SecurityEdiscoveryCase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryCase: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityCaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityCaseImpl: %+v", err)
	}

	return RawSecurityCaseImpl{
		securityCase: parent,
		Type:         value,
		Values:       temp,
	}, nil

}
