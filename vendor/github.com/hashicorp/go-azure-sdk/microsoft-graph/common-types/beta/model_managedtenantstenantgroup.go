package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsTenantGroup{}

type ManagedTenantsTenantGroup struct {
	// A flag indicating whether all managed tenant are included in the tenant group. Required. Read-only.
	AllTenantsIncluded *bool `json:"allTenantsIncluded,omitempty"`

	// The display name for the tenant group. Optional. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The collection of management action associated with the tenant group. Optional. Read-only.
	ManagementActions *[]ManagedTenantsManagementActionInfo `json:"managementActions,omitempty"`

	// The collection of management intents associated with the tenant group. Optional. Read-only.
	ManagementIntents *[]ManagedTenantsManagementIntentInfo `json:"managementIntents,omitempty"`

	// The collection of managed tenant identifiers include in the tenant group. Optional. Read-only.
	TenantIds *[]string `json:"tenantIds,omitempty"`

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

func (s ManagedTenantsTenantGroup) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsTenantGroup{}

func (s ManagedTenantsTenantGroup) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsTenantGroup
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsTenantGroup: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsTenantGroup: %+v", err)
	}

	delete(decoded, "allTenantsIncluded")
	delete(decoded, "displayName")
	delete(decoded, "managementActions")
	delete(decoded, "managementIntents")
	delete(decoded, "tenantIds")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.tenantGroup"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsTenantGroup: %+v", err)
	}

	return encoded, nil
}
