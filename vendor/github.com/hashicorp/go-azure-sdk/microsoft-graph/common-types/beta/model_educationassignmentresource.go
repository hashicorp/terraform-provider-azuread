package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EducationAssignmentResource{}

type EducationAssignmentResource struct {
	DependentResources *[]EducationAssignmentResource `json:"dependentResources,omitempty"`

	// Indicates whether this resource should be copied to each student submission for modification and submission. Required
	DistributeForStudentWork nullable.Type[bool] `json:"distributeForStudentWork,omitempty"`

	// Resource object that is associated with this assignment.
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

func (s EducationAssignmentResource) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationAssignmentResource{}

func (s EducationAssignmentResource) MarshalJSON() ([]byte, error) {
	type wrapper EducationAssignmentResource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationAssignmentResource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationAssignmentResource: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationAssignmentResource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationAssignmentResource: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EducationAssignmentResource{}

func (s *EducationAssignmentResource) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DependentResources       *[]EducationAssignmentResource `json:"dependentResources,omitempty"`
		DistributeForStudentWork nullable.Type[bool]            `json:"distributeForStudentWork,omitempty"`
		Id                       *string                        `json:"id,omitempty"`
		ODataId                  *string                        `json:"@odata.id,omitempty"`
		ODataType                *string                        `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DependentResources = decoded.DependentResources
	s.DistributeForStudentWork = decoded.DistributeForStudentWork
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EducationAssignmentResource into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["resource"]; ok {
		impl, err := UnmarshalEducationResourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Resource' for 'EducationAssignmentResource': %+v", err)
		}
		s.Resource = impl
	}

	return nil
}
