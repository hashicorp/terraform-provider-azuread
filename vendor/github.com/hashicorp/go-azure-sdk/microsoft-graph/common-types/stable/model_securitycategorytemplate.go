package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityFilePlanDescriptorTemplate = SecurityCategoryTemplate{}

type SecurityCategoryTemplate struct {
	// Represents all subcategories under a particular category.
	Subcategories *[]SecuritySubcategoryTemplate `json:"subcategories,omitempty"`

	// Fields inherited from SecurityFilePlanDescriptorTemplate

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

func (s SecurityCategoryTemplate) SecurityFilePlanDescriptorTemplate() BaseSecurityFilePlanDescriptorTemplateImpl {
	return BaseSecurityFilePlanDescriptorTemplateImpl{
		CreatedBy:       s.CreatedBy,
		CreatedDateTime: s.CreatedDateTime,
		DisplayName:     s.DisplayName,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s SecurityCategoryTemplate) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityCategoryTemplate{}

func (s SecurityCategoryTemplate) MarshalJSON() ([]byte, error) {
	type wrapper SecurityCategoryTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityCategoryTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityCategoryTemplate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.categoryTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityCategoryTemplate: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityCategoryTemplate{}

func (s *SecurityCategoryTemplate) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Subcategories   *[]SecuritySubcategoryTemplate `json:"subcategories,omitempty"`
		CreatedDateTime nullable.Type[string]          `json:"createdDateTime,omitempty"`
		DisplayName     nullable.Type[string]          `json:"displayName,omitempty"`
		Id              *string                        `json:"id,omitempty"`
		ODataId         *string                        `json:"@odata.id,omitempty"`
		ODataType       *string                        `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Subcategories = decoded.Subcategories
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityCategoryTemplate into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SecurityCategoryTemplate': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}
