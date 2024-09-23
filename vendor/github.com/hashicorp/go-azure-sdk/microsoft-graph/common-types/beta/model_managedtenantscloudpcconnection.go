package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsCloudPCConnection{}

type ManagedTenantsCloudPCConnection struct {
	// The display name of the cloud PC connection. Required. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The health status of the cloud PC connection. Possible values are: pending, running, passed, failed,
	// unknownFutureValue. Required. Read-only.
	HealthCheckStatus nullable.Type[string] `json:"healthCheckStatus,omitempty"`

	// Date and time the entity was last updated in the multi-tenant management platform. Required. Read-only.
	LastRefreshedDateTime nullable.Type[string] `json:"lastRefreshedDateTime,omitempty"`

	// The display name for the managed tenant. Required. Read-only.
	TenantDisplayName nullable.Type[string] `json:"tenantDisplayName,omitempty"`

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

func (s ManagedTenantsCloudPCConnection) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsCloudPCConnection{}

func (s ManagedTenantsCloudPCConnection) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsCloudPCConnection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsCloudPCConnection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsCloudPCConnection: %+v", err)
	}

	delete(decoded, "displayName")
	delete(decoded, "healthCheckStatus")
	delete(decoded, "lastRefreshedDateTime")
	delete(decoded, "tenantDisplayName")
	delete(decoded, "tenantId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.cloudPcConnection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsCloudPCConnection: %+v", err)
	}

	return encoded, nil
}
