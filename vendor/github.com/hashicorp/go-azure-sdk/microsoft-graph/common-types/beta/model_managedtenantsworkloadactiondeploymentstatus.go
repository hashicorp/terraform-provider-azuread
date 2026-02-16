package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsWorkloadActionDeploymentStatus struct {
	// The unique identifier for the workload action. Required. Read-only.
	ActionId *string `json:"actionId,omitempty"`

	// The identifier of any policy that was created by applying the workload action. Optional. Read-only.
	DeployedPolicyId nullable.Type[string] `json:"deployedPolicyId,omitempty"`

	// The detailed information for exceptions that occur when deploying the workload action. Optional. Required.
	Error GenericError `json:"error"`

	ExcludeGroups   *[]string `json:"excludeGroups,omitempty"`
	IncludeAllUsers *bool     `json:"includeAllUsers,omitempty"`
	IncludeGroups   *[]string `json:"includeGroups,omitempty"`

	// The date and time the workload action was last deployed. Optional.
	LastDeploymentDateTime nullable.Type[string] `json:"lastDeploymentDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Status *ManagedTenantsWorkloadActionStatus `json:"status,omitempty"`
}

var _ json.Marshaler = ManagedTenantsWorkloadActionDeploymentStatus{}

func (s ManagedTenantsWorkloadActionDeploymentStatus) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsWorkloadActionDeploymentStatus
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsWorkloadActionDeploymentStatus: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsWorkloadActionDeploymentStatus: %+v", err)
	}

	delete(decoded, "actionId")
	delete(decoded, "deployedPolicyId")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsWorkloadActionDeploymentStatus: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ManagedTenantsWorkloadActionDeploymentStatus{}

func (s *ManagedTenantsWorkloadActionDeploymentStatus) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ActionId               *string                             `json:"actionId,omitempty"`
		DeployedPolicyId       nullable.Type[string]               `json:"deployedPolicyId,omitempty"`
		ExcludeGroups          *[]string                           `json:"excludeGroups,omitempty"`
		IncludeAllUsers        *bool                               `json:"includeAllUsers,omitempty"`
		IncludeGroups          *[]string                           `json:"includeGroups,omitempty"`
		LastDeploymentDateTime nullable.Type[string]               `json:"lastDeploymentDateTime,omitempty"`
		ODataId                *string                             `json:"@odata.id,omitempty"`
		ODataType              *string                             `json:"@odata.type,omitempty"`
		Status                 *ManagedTenantsWorkloadActionStatus `json:"status,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ActionId = decoded.ActionId
	s.DeployedPolicyId = decoded.DeployedPolicyId
	s.ExcludeGroups = decoded.ExcludeGroups
	s.IncludeAllUsers = decoded.IncludeAllUsers
	s.IncludeGroups = decoded.IncludeGroups
	s.LastDeploymentDateTime = decoded.LastDeploymentDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ManagedTenantsWorkloadActionDeploymentStatus into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["error"]; ok {
		impl, err := UnmarshalGenericErrorImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Error' for 'ManagedTenantsWorkloadActionDeploymentStatus': %+v", err)
		}
		s.Error = impl
	}

	return nil
}
