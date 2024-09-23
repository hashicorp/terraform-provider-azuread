package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PolicySet{}

type PolicySet struct {
	// Assignments of the PolicySet.
	Assignments *[]PolicySetAssignment `json:"assignments,omitempty"`

	// Creation time of the PolicySet.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Description of the PolicySet.
	Description nullable.Type[string] `json:"description,omitempty"`

	// DisplayName of the PolicySet.
	DisplayName *string `json:"displayName,omitempty"`

	ErrorCode *ErrorCode `json:"errorCode,omitempty"`

	// Tags of the guided deployment
	GuidedDeploymentTags *[]string `json:"guidedDeploymentTags,omitempty"`

	// Items of the PolicySet with maximum count 100.
	Items *[]PolicySetItem `json:"items,omitempty"`

	// Last modified time of the PolicySet.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// RoleScopeTags of the PolicySet
	RoleScopeTags *[]string `json:"roleScopeTags,omitempty"`

	// The enum to specify the status of PolicySet.
	Status *PolicySetStatus `json:"status,omitempty"`

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

func (s PolicySet) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PolicySet{}

func (s PolicySet) MarshalJSON() ([]byte, error) {
	type wrapper PolicySet
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PolicySet: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PolicySet: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.policySet"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PolicySet: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PolicySet{}

func (s *PolicySet) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Assignments          *[]PolicySetAssignment `json:"assignments,omitempty"`
		CreatedDateTime      *string                `json:"createdDateTime,omitempty"`
		Description          nullable.Type[string]  `json:"description,omitempty"`
		DisplayName          *string                `json:"displayName,omitempty"`
		ErrorCode            *ErrorCode             `json:"errorCode,omitempty"`
		GuidedDeploymentTags *[]string              `json:"guidedDeploymentTags,omitempty"`
		LastModifiedDateTime *string                `json:"lastModifiedDateTime,omitempty"`
		RoleScopeTags        *[]string              `json:"roleScopeTags,omitempty"`
		Status               *PolicySetStatus       `json:"status,omitempty"`
		Id                   *string                `json:"id,omitempty"`
		ODataId              *string                `json:"@odata.id,omitempty"`
		ODataType            *string                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Assignments = decoded.Assignments
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.ErrorCode = decoded.ErrorCode
	s.GuidedDeploymentTags = decoded.GuidedDeploymentTags
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.RoleScopeTags = decoded.RoleScopeTags
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PolicySet into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["items"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Items into list []json.RawMessage: %+v", err)
		}

		output := make([]PolicySetItem, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalPolicySetItemImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Items' for 'PolicySet': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Items = &output
	}

	return nil
}
