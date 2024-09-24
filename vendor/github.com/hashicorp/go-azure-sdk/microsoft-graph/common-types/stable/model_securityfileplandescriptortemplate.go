package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFilePlanDescriptorTemplate interface {
	Entity
	SecurityFilePlanDescriptorTemplate() BaseSecurityFilePlanDescriptorTemplateImpl
}

var _ SecurityFilePlanDescriptorTemplate = BaseSecurityFilePlanDescriptorTemplateImpl{}

type BaseSecurityFilePlanDescriptorTemplateImpl struct {
	// Represents the user who created the filePlanDescriptorTemplate column.
	CreatedBy IdentitySet `json:"createdBy"`

	// Represents the date and time in which the filePlanDescriptorTemplate is created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Unique string that defines a filePlanDescriptorTemplate name.
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

func (s BaseSecurityFilePlanDescriptorTemplateImpl) SecurityFilePlanDescriptorTemplate() BaseSecurityFilePlanDescriptorTemplateImpl {
	return s
}

func (s BaseSecurityFilePlanDescriptorTemplateImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ SecurityFilePlanDescriptorTemplate = RawSecurityFilePlanDescriptorTemplateImpl{}

// RawSecurityFilePlanDescriptorTemplateImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityFilePlanDescriptorTemplateImpl struct {
	securityFilePlanDescriptorTemplate BaseSecurityFilePlanDescriptorTemplateImpl
	Type                               string
	Values                             map[string]interface{}
}

func (s RawSecurityFilePlanDescriptorTemplateImpl) SecurityFilePlanDescriptorTemplate() BaseSecurityFilePlanDescriptorTemplateImpl {
	return s.securityFilePlanDescriptorTemplate
}

func (s RawSecurityFilePlanDescriptorTemplateImpl) Entity() BaseEntityImpl {
	return s.securityFilePlanDescriptorTemplate.Entity()
}

var _ json.Marshaler = BaseSecurityFilePlanDescriptorTemplateImpl{}

func (s BaseSecurityFilePlanDescriptorTemplateImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseSecurityFilePlanDescriptorTemplateImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseSecurityFilePlanDescriptorTemplateImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseSecurityFilePlanDescriptorTemplateImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.filePlanDescriptorTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseSecurityFilePlanDescriptorTemplateImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseSecurityFilePlanDescriptorTemplateImpl{}

func (s *BaseSecurityFilePlanDescriptorTemplateImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`
		DisplayName     nullable.Type[string] `json:"displayName,omitempty"`
		Id              *string               `json:"id,omitempty"`
		ODataId         *string               `json:"@odata.id,omitempty"`
		ODataType       *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseSecurityFilePlanDescriptorTemplateImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'BaseSecurityFilePlanDescriptorTemplateImpl': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}

func UnmarshalSecurityFilePlanDescriptorTemplateImplementation(input []byte) (SecurityFilePlanDescriptorTemplate, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityFilePlanDescriptorTemplate into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.authorityTemplate") {
		var out SecurityAuthorityTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAuthorityTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.categoryTemplate") {
		var out SecurityCategoryTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCategoryTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.citationTemplate") {
		var out SecurityCitationTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCitationTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.departmentTemplate") {
		var out SecurityDepartmentTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDepartmentTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.filePlanReferenceTemplate") {
		var out SecurityFilePlanReferenceTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityFilePlanReferenceTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.subcategoryTemplate") {
		var out SecuritySubcategoryTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySubcategoryTemplate: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityFilePlanDescriptorTemplateImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityFilePlanDescriptorTemplateImpl: %+v", err)
	}

	return RawSecurityFilePlanDescriptorTemplateImpl{
		securityFilePlanDescriptorTemplate: parent,
		Type:                               value,
		Values:                             temp,
	}, nil

}
