package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsManagementTemplate{}

type ManagedTenantsManagementTemplate struct {
	// The management category for the management template. Possible values are: custom, devices, identity,
	// unknownFutureValue. Required. Read-only.
	Category *ManagedTenantsManagementCategory `json:"category,omitempty"`

	CreatedByUserId nullable.Type[string] `json:"createdByUserId,omitempty"`
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The description for the management template. Optional. Read-only.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name for the management template. Required. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	InformationLinks              *[]ActionUrl                                  `json:"informationLinks,omitempty"`
	LastActionByUserId            nullable.Type[string]                         `json:"lastActionByUserId,omitempty"`
	LastActionDateTime            nullable.Type[string]                         `json:"lastActionDateTime,omitempty"`
	ManagementTemplateCollections *[]ManagedTenantsManagementTemplateCollection `json:"managementTemplateCollections,omitempty"`
	ManagementTemplateSteps       *[]ManagedTenantsManagementTemplateStep       `json:"managementTemplateSteps,omitempty"`

	// The collection of parameters used by the management template. Optional. Read-only.
	Parameters *[]ManagedTenantsTemplateParameter `json:"parameters,omitempty"`

	Priority   *int64                            `json:"priority,omitempty"`
	Provider   *ManagedTenantsManagementProvider `json:"provider,omitempty"`
	UserImpact nullable.Type[string]             `json:"userImpact,omitempty"`
	Version    nullable.Type[int64]              `json:"version,omitempty"`

	// The collection of workload actions associated with the management template. Optional. Read-only.
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

func (s ManagedTenantsManagementTemplate) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsManagementTemplate{}

func (s ManagedTenantsManagementTemplate) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagementTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagementTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagementTemplate: %+v", err)
	}

	delete(decoded, "category")
	delete(decoded, "description")
	delete(decoded, "displayName")
	delete(decoded, "parameters")
	delete(decoded, "workloadActions")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.managementTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagementTemplate: %+v", err)
	}

	return encoded, nil
}
