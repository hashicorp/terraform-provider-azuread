package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Finding = VirtualMachineWithAwsStorageBucketAccessFinding{}

type VirtualMachineWithAwsStorageBucketAccessFinding struct {
	// The total number of storage buckets that the EC2 instance can access using the role.
	AccessibleCount *int64 `json:"accessibleCount,omitempty"`

	// The total number of storage buckets in the authorization system that hosts the EC2 instance.
	BucketCount *int64 `json:"bucketCount,omitempty"`

	Ec2Instance           *AuthorizationSystemResource `json:"ec2Instance,omitempty"`
	PermissionsCreepIndex *PermissionsCreepIndex       `json:"permissionsCreepIndex,omitempty"`
	Role                  *AwsRole                     `json:"role,omitempty"`

	// Fields inherited from Finding

	// Defines when the finding was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

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

func (s VirtualMachineWithAwsStorageBucketAccessFinding) Finding() BaseFindingImpl {
	return BaseFindingImpl{
		CreatedDateTime: s.CreatedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s VirtualMachineWithAwsStorageBucketAccessFinding) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = VirtualMachineWithAwsStorageBucketAccessFinding{}

func (s VirtualMachineWithAwsStorageBucketAccessFinding) MarshalJSON() ([]byte, error) {
	type wrapper VirtualMachineWithAwsStorageBucketAccessFinding
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VirtualMachineWithAwsStorageBucketAccessFinding: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VirtualMachineWithAwsStorageBucketAccessFinding: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.virtualMachineWithAwsStorageBucketAccessFinding"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VirtualMachineWithAwsStorageBucketAccessFinding: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &VirtualMachineWithAwsStorageBucketAccessFinding{}

func (s *VirtualMachineWithAwsStorageBucketAccessFinding) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AccessibleCount       *int64                 `json:"accessibleCount,omitempty"`
		BucketCount           *int64                 `json:"bucketCount,omitempty"`
		PermissionsCreepIndex *PermissionsCreepIndex `json:"permissionsCreepIndex,omitempty"`
		Role                  *AwsRole               `json:"role,omitempty"`
		CreatedDateTime       *string                `json:"createdDateTime,omitempty"`
		Id                    *string                `json:"id,omitempty"`
		ODataId               *string                `json:"@odata.id,omitempty"`
		ODataType             *string                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccessibleCount = decoded.AccessibleCount
	s.BucketCount = decoded.BucketCount
	s.PermissionsCreepIndex = decoded.PermissionsCreepIndex
	s.Role = decoded.Role
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling VirtualMachineWithAwsStorageBucketAccessFinding into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["ec2Instance"]; ok {
		impl, err := UnmarshalAuthorizationSystemResourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Ec2Instance' for 'VirtualMachineWithAwsStorageBucketAccessFinding': %+v", err)
		}
		s.Ec2Instance = &impl
	}

	return nil
}
