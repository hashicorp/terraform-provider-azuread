package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementActionDeploymentStatus struct {
	// The identifier for the management action. Required. Read-only.
	ManagementActionId *string `json:"managementActionId,omitempty"`

	// The management template identifier that was used to generate the management action. Required. Read-only.
	ManagementTemplateId *string `json:"managementTemplateId,omitempty"`

	ManagementTemplateVersion *int64 `json:"managementTemplateVersion,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Status *ManagedTenantsManagementActionStatus `json:"status,omitempty"`

	// The collection of workload action deployment statues for the given management action. Optional.
	WorkloadActionDeploymentStatuses *[]ManagedTenantsWorkloadActionDeploymentStatus `json:"workloadActionDeploymentStatuses,omitempty"`
}

var _ json.Marshaler = ManagedTenantsManagementActionDeploymentStatus{}

func (s ManagedTenantsManagementActionDeploymentStatus) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagementActionDeploymentStatus
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagementActionDeploymentStatus: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagementActionDeploymentStatus: %+v", err)
	}

	delete(decoded, "managementActionId")
	delete(decoded, "managementTemplateId")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagementActionDeploymentStatus: %+v", err)
	}

	return encoded, nil
}
