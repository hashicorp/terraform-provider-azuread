package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsManagementAction{}

type ManagedTenantsManagementAction struct {
	Category *ManagedTenantsManagementCategory `json:"category,omitempty"`

	// The description for the management action. Optional. Read-only.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name for the management action. Optional. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The reference for the management template used to generate the management action. Required. Read-only.
	ReferenceTemplateId *string `json:"referenceTemplateId,omitempty"`

	ReferenceTemplateVersion *int64 `json:"referenceTemplateVersion,omitempty"`

	// The collection of workload actions associated with the management action. Required. Read-only.
	WorkloadActions *[]ManagedTenantsWorkloadAction `json:"workloadActions,omitempty"`

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

func (s ManagedTenantsManagementAction) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsManagementAction{}

func (s ManagedTenantsManagementAction) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagementAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagementAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagementAction: %+v", err)
	}

	delete(decoded, "description")
	delete(decoded, "displayName")
	delete(decoded, "referenceTemplateId")
	delete(decoded, "workloadActions")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.managementAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagementAction: %+v", err)
	}

	return encoded, nil
}
