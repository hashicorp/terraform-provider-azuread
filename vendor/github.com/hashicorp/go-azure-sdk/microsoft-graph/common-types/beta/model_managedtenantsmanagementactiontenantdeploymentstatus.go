package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsManagementActionTenantDeploymentStatus{}

type ManagedTenantsManagementActionTenantDeploymentStatus struct {
	// The collection of deployment status for each instance of a management action. Optional.
	Statuses *[]ManagedTenantsManagementActionDeploymentStatus `json:"statuses,omitempty"`

	// The identifier for the tenant group that is associated with the management action. Required. Read-only.
	TenantGroupId nullable.Type[string] `json:"tenantGroupId,omitempty"`

	// The Microsoft Entra tenant identifier for the managed tenant. Required. Read-only.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

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

func (s ManagedTenantsManagementActionTenantDeploymentStatus) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsManagementActionTenantDeploymentStatus{}

func (s ManagedTenantsManagementActionTenantDeploymentStatus) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagementActionTenantDeploymentStatus
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagementActionTenantDeploymentStatus: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagementActionTenantDeploymentStatus: %+v", err)
	}

	delete(decoded, "tenantGroupId")
	delete(decoded, "tenantId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.managementActionTenantDeploymentStatus"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagementActionTenantDeploymentStatus: %+v", err)
	}

	return encoded, nil
}
