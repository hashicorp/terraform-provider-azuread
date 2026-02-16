package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AssignedComputeInstanceDetails{}

type AssignedComputeInstanceDetails struct {
	// Represents a set of S3 buckets accessed by this EC2 instance.
	AccessedStorageBuckets *[]AuthorizationSystemResource `json:"accessedStorageBuckets,omitempty"`

	// assigned EC2 instance.
	AssignedComputeInstance *AuthorizationSystemResource `json:"assignedComputeInstance,omitempty"`

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

func (s AssignedComputeInstanceDetails) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AssignedComputeInstanceDetails{}

func (s AssignedComputeInstanceDetails) MarshalJSON() ([]byte, error) {
	type wrapper AssignedComputeInstanceDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AssignedComputeInstanceDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AssignedComputeInstanceDetails: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.assignedComputeInstanceDetails"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AssignedComputeInstanceDetails: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AssignedComputeInstanceDetails{}

func (s *AssignedComputeInstanceDetails) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Id        *string `json:"id,omitempty"`
		ODataId   *string `json:"@odata.id,omitempty"`
		ODataType *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AssignedComputeInstanceDetails into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["accessedStorageBuckets"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AccessedStorageBuckets into list []json.RawMessage: %+v", err)
		}

		output := make([]AuthorizationSystemResource, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAuthorizationSystemResourceImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AccessedStorageBuckets' for 'AssignedComputeInstanceDetails': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AccessedStorageBuckets = &output
	}

	if v, ok := temp["assignedComputeInstance"]; ok {
		impl, err := UnmarshalAuthorizationSystemResourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AssignedComputeInstance' for 'AssignedComputeInstanceDetails': %+v", err)
		}
		s.AssignedComputeInstance = &impl
	}

	return nil
}
