package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsManagementTemplateCollectionTenantSummary{}

type ManagedTenantsManagementTemplateCollectionTenantSummary struct {
	CompleteStepsCount                      nullable.Type[int64]  `json:"completeStepsCount,omitempty"`
	CompleteUsersCount                      nullable.Type[int64]  `json:"completeUsersCount,omitempty"`
	CreatedByUserId                         nullable.Type[string] `json:"createdByUserId,omitempty"`
	CreatedDateTime                         nullable.Type[string] `json:"createdDateTime,omitempty"`
	DismissedStepsCount                     nullable.Type[int64]  `json:"dismissedStepsCount,omitempty"`
	ExcludedUsersCount                      nullable.Type[int64]  `json:"excludedUsersCount,omitempty"`
	ExcludedUsersDistinctCount              nullable.Type[int64]  `json:"excludedUsersDistinctCount,omitempty"`
	IncompleteStepsCount                    nullable.Type[int64]  `json:"incompleteStepsCount,omitempty"`
	IncompleteUsersCount                    nullable.Type[int64]  `json:"incompleteUsersCount,omitempty"`
	IneligibleStepsCount                    nullable.Type[int64]  `json:"ineligibleStepsCount,omitempty"`
	IsComplete                              nullable.Type[bool]   `json:"isComplete,omitempty"`
	LastActionByUserId                      nullable.Type[string] `json:"lastActionByUserId,omitempty"`
	LastActionDateTime                      nullable.Type[string] `json:"lastActionDateTime,omitempty"`
	ManagementTemplateCollectionDisplayName nullable.Type[string] `json:"managementTemplateCollectionDisplayName,omitempty"`
	ManagementTemplateCollectionId          nullable.Type[string] `json:"managementTemplateCollectionId,omitempty"`
	RegressedStepsCount                     nullable.Type[int64]  `json:"regressedStepsCount,omitempty"`
	RegressedUsersCount                     nullable.Type[int64]  `json:"regressedUsersCount,omitempty"`
	TenantId                                nullable.Type[string] `json:"tenantId,omitempty"`
	UnlicensedUsersCount                    nullable.Type[int64]  `json:"unlicensedUsersCount,omitempty"`

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

func (s ManagedTenantsManagementTemplateCollectionTenantSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsManagementTemplateCollectionTenantSummary{}

func (s ManagedTenantsManagementTemplateCollectionTenantSummary) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagementTemplateCollectionTenantSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagementTemplateCollectionTenantSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagementTemplateCollectionTenantSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.managementTemplateCollectionTenantSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagementTemplateCollectionTenantSummary: %+v", err)
	}

	return encoded, nil
}
