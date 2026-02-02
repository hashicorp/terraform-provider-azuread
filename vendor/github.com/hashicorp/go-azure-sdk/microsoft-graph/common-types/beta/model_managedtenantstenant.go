package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsTenant{}

type ManagedTenantsTenant struct {
	// The relationship details for the tenant with the managing entity.
	Contract *ManagedTenantsTenantContract `json:"contract,omitempty"`

	// The date and time the tenant was created in the multi-tenant management platform. Optional. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The display name for the tenant. Required. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The date and time the tenant was last updated within the multi-tenant management platform. Optional. Read-only.
	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`

	// The Microsoft Entra tenant identifier for the managed tenant. Optional. Read-only.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

	// The onboarding status information for the tenant. Optional. Read-only.
	TenantStatusInformation *ManagedTenantsTenantStatusInformation `json:"tenantStatusInformation,omitempty"`

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

func (s ManagedTenantsTenant) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsTenant{}

func (s ManagedTenantsTenant) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsTenant
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsTenant: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsTenant: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "displayName")
	delete(decoded, "lastUpdatedDateTime")
	delete(decoded, "tenantId")
	delete(decoded, "tenantStatusInformation")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.tenant"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsTenant: %+v", err)
	}

	return encoded, nil
}
