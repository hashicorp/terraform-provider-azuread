package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EducationSubmissionResource{}

type EducationSubmissionResource struct {
	// Pointer to the assignment from which the resource was copied. If the value is null, the student uploaded the
	// resource.
	AssignmentResourceUrl nullable.Type[string] `json:"assignmentResourceUrl,omitempty"`

	DependentResources *[]EducationSubmissionResource `json:"dependentResources,omitempty"`

	// Resource object.
	Resource EducationResource `json:"resource"`

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

func (s EducationSubmissionResource) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationSubmissionResource{}

func (s EducationSubmissionResource) MarshalJSON() ([]byte, error) {
	type wrapper EducationSubmissionResource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationSubmissionResource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationSubmissionResource: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationSubmissionResource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationSubmissionResource: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EducationSubmissionResource{}

func (s *EducationSubmissionResource) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AssignmentResourceUrl nullable.Type[string]          `json:"assignmentResourceUrl,omitempty"`
		DependentResources    *[]EducationSubmissionResource `json:"dependentResources,omitempty"`
		Id                    *string                        `json:"id,omitempty"`
		ODataId               *string                        `json:"@odata.id,omitempty"`
		ODataType             *string                        `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AssignmentResourceUrl = decoded.AssignmentResourceUrl
	s.DependentResources = decoded.DependentResources
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EducationSubmissionResource into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["resource"]; ok {
		impl, err := UnmarshalEducationResourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Resource' for 'EducationSubmissionResource': %+v", err)
		}
		s.Resource = impl
	}

	return nil
}
