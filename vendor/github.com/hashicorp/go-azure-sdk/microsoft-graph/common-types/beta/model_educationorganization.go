package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationOrganization interface {
	Entity
	EducationOrganization() BaseEducationOrganizationImpl
}

var _ EducationOrganization = BaseEducationOrganizationImpl{}

type BaseEducationOrganizationImpl struct {
	// Organization description.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Organization display name.
	DisplayName *string `json:"displayName,omitempty"`

	// Where this user was created from. Possible values are: sis, lms, or manual.
	ExternalSource *EducationExternalSource `json:"externalSource,omitempty"`

	ExternalSourceDetail nullable.Type[string] `json:"externalSourceDetail,omitempty"`

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

func (s BaseEducationOrganizationImpl) EducationOrganization() BaseEducationOrganizationImpl {
	return s
}

func (s BaseEducationOrganizationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ EducationOrganization = RawEducationOrganizationImpl{}

// RawEducationOrganizationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEducationOrganizationImpl struct {
	educationOrganization BaseEducationOrganizationImpl
	Type                  string
	Values                map[string]interface{}
}

func (s RawEducationOrganizationImpl) EducationOrganization() BaseEducationOrganizationImpl {
	return s.educationOrganization
}

func (s RawEducationOrganizationImpl) Entity() BaseEntityImpl {
	return s.educationOrganization.Entity()
}

var _ json.Marshaler = BaseEducationOrganizationImpl{}

func (s BaseEducationOrganizationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseEducationOrganizationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseEducationOrganizationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseEducationOrganizationImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationOrganization"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseEducationOrganizationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalEducationOrganizationImplementation(input []byte) (EducationOrganization, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationOrganization into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.educationSchool") {
		var out EducationSchool
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationSchool: %+v", err)
		}
		return out, nil
	}

	var parent BaseEducationOrganizationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEducationOrganizationImpl: %+v", err)
	}

	return RawEducationOrganizationImpl{
		educationOrganization: parent,
		Type:                  value,
		Values:                temp,
	}, nil

}
