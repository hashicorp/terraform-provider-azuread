package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Finding = OpenAwsSecurityGroupFinding{}

type OpenAwsSecurityGroupFinding struct {
	// A set of AWS EC2 compute instances related to this open security group.
	AssignedComputeInstancesDetails *[]AssignedComputeInstanceDetails `json:"assignedComputeInstancesDetails,omitempty"`

	InboundPorts  InboundPorts                    `json:"inboundPorts"`
	SecurityGroup *AwsAuthorizationSystemResource `json:"securityGroup,omitempty"`

	// The number of storage buckets accessed by the assigned compute instances.
	TotalStorageBucketCount *int64 `json:"totalStorageBucketCount,omitempty"`

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

func (s OpenAwsSecurityGroupFinding) Finding() BaseFindingImpl {
	return BaseFindingImpl{
		CreatedDateTime: s.CreatedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s OpenAwsSecurityGroupFinding) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OpenAwsSecurityGroupFinding{}

func (s OpenAwsSecurityGroupFinding) MarshalJSON() ([]byte, error) {
	type wrapper OpenAwsSecurityGroupFinding
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OpenAwsSecurityGroupFinding: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OpenAwsSecurityGroupFinding: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.openAwsSecurityGroupFinding"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OpenAwsSecurityGroupFinding: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &OpenAwsSecurityGroupFinding{}

func (s *OpenAwsSecurityGroupFinding) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AssignedComputeInstancesDetails *[]AssignedComputeInstanceDetails `json:"assignedComputeInstancesDetails,omitempty"`
		SecurityGroup                   *AwsAuthorizationSystemResource   `json:"securityGroup,omitempty"`
		TotalStorageBucketCount         *int64                            `json:"totalStorageBucketCount,omitempty"`
		CreatedDateTime                 *string                           `json:"createdDateTime,omitempty"`
		Id                              *string                           `json:"id,omitempty"`
		ODataId                         *string                           `json:"@odata.id,omitempty"`
		ODataType                       *string                           `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AssignedComputeInstancesDetails = decoded.AssignedComputeInstancesDetails
	s.SecurityGroup = decoded.SecurityGroup
	s.TotalStorageBucketCount = decoded.TotalStorageBucketCount
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling OpenAwsSecurityGroupFinding into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["inboundPorts"]; ok {
		impl, err := UnmarshalInboundPortsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'InboundPorts' for 'OpenAwsSecurityGroupFinding': %+v", err)
		}
		s.InboundPorts = impl
	}

	return nil
}
