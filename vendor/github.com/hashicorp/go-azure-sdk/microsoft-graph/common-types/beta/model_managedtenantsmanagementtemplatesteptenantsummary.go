package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsManagementTemplateStepTenantSummary{}

type ManagedTenantsManagementTemplateStepTenantSummary struct {
	AssignedTenantsCount                    nullable.Type[int64]  `json:"assignedTenantsCount,omitempty"`
	CompliantTenantsCount                   nullable.Type[int64]  `json:"compliantTenantsCount,omitempty"`
	CreatedByUserId                         nullable.Type[string] `json:"createdByUserId,omitempty"`
	CreatedDateTime                         nullable.Type[string] `json:"createdDateTime,omitempty"`
	DismissedTenantsCount                   nullable.Type[int64]  `json:"dismissedTenantsCount,omitempty"`
	IneligibleTenantsCount                  nullable.Type[int64]  `json:"ineligibleTenantsCount,omitempty"`
	LastActionByUserId                      nullable.Type[string] `json:"lastActionByUserId,omitempty"`
	LastActionDateTime                      nullable.Type[string] `json:"lastActionDateTime,omitempty"`
	ManagementTemplateCollectionDisplayName nullable.Type[string] `json:"managementTemplateCollectionDisplayName,omitempty"`
	ManagementTemplateCollectionId          nullable.Type[string] `json:"managementTemplateCollectionId,omitempty"`
	ManagementTemplateDisplayName           nullable.Type[string] `json:"managementTemplateDisplayName,omitempty"`
	ManagementTemplateId                    nullable.Type[string] `json:"managementTemplateId,omitempty"`
	ManagementTemplateStepDisplayName       nullable.Type[string] `json:"managementTemplateStepDisplayName,omitempty"`
	ManagementTemplateStepId                nullable.Type[string] `json:"managementTemplateStepId,omitempty"`
	NotCompliantTenantsCount                nullable.Type[int64]  `json:"notCompliantTenantsCount,omitempty"`

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

func (s ManagedTenantsManagementTemplateStepTenantSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsManagementTemplateStepTenantSummary{}

func (s ManagedTenantsManagementTemplateStepTenantSummary) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagementTemplateStepTenantSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagementTemplateStepTenantSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagementTemplateStepTenantSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.managementTemplateStepTenantSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagementTemplateStepTenantSummary: %+v", err)
	}

	return encoded, nil
}
