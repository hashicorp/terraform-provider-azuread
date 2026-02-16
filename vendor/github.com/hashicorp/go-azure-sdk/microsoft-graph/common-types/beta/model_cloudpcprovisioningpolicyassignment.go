package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudPCProvisioningPolicyAssignment{}

type CloudPCProvisioningPolicyAssignment struct {
	// The assignment targeted users for the provisioning policy. This list of users is computed based on assignments,
	// licenses, group memberships, and policies. This property is read-only. Supports$expand.
	AssignedUsers *[]User `json:"assignedUsers,omitempty"`

	// The assignment target for the provisioning policy. Currently, the only target supported for this policy is a user
	// group. For details, see cloudPcManagementGroupAssignmentTarget.
	Target CloudPCManagementAssignmentTarget `json:"target"`

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

func (s CloudPCProvisioningPolicyAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudPCProvisioningPolicyAssignment{}

func (s CloudPCProvisioningPolicyAssignment) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCProvisioningPolicyAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCProvisioningPolicyAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCProvisioningPolicyAssignment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPcProvisioningPolicyAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCProvisioningPolicyAssignment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &CloudPCProvisioningPolicyAssignment{}

func (s *CloudPCProvisioningPolicyAssignment) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AssignedUsers *[]User `json:"assignedUsers,omitempty"`
		Id            *string `json:"id,omitempty"`
		ODataId       *string `json:"@odata.id,omitempty"`
		ODataType     *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AssignedUsers = decoded.AssignedUsers
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CloudPCProvisioningPolicyAssignment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["target"]; ok {
		impl, err := UnmarshalCloudPCManagementAssignmentTargetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Target' for 'CloudPCProvisioningPolicyAssignment': %+v", err)
		}
		s.Target = impl
	}

	return nil
}
