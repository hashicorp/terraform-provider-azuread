package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityFilePlanDescriptorTemplate = SecurityAuthorityTemplate{}

type SecurityAuthorityTemplate struct {

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

func (s SecurityAuthorityTemplate) SecurityFilePlanDescriptorTemplate() BaseSecurityFilePlanDescriptorTemplateImpl {
	return BaseSecurityFilePlanDescriptorTemplateImpl{
		CreatedBy:       s.CreatedBy,
		CreatedDateTime: s.CreatedDateTime,
		DisplayName:     s.DisplayName,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s SecurityAuthorityTemplate) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityAuthorityTemplate{}

func (s SecurityAuthorityTemplate) MarshalJSON() ([]byte, error) {
	type wrapper SecurityAuthorityTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityAuthorityTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityAuthorityTemplate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.authorityTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityAuthorityTemplate: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityAuthorityTemplate{}

func (s *SecurityAuthorityTemplate) UnmarshalJSON(bytes []byte) error {
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
		return fmt.Errorf("unmarshaling SecurityAuthorityTemplate into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SecurityAuthorityTemplate': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}
