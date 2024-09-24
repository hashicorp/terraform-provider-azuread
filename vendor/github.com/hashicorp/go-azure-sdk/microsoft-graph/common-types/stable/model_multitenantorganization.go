package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MultiTenantOrganization{}

type MultiTenantOrganization struct {
	// Date when multitenant organization was created. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Description of the multitenant organization.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display name of the multitenant organization.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Defines the status of a tenant joining a multitenant organization.
	JoinRequest *MultiTenantOrganizationJoinRequestRecord `json:"joinRequest,omitempty"`

	// State of the multitenant organization. The possible values are: active, inactive, unknownFutureValue. active
	// indicates the multitenant organization is created. inactive indicates the multitenant organization isn't created.
	// Read-only.
	State *MultiTenantOrganizationState `json:"state,omitempty"`

	// Defines tenants added to a multitenant organization.
	Tenants *[]MultiTenantOrganizationMember `json:"tenants,omitempty"`

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

func (s MultiTenantOrganization) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MultiTenantOrganization{}

func (s MultiTenantOrganization) MarshalJSON() ([]byte, error) {
	type wrapper MultiTenantOrganization
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MultiTenantOrganization: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MultiTenantOrganization: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "state")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.multiTenantOrganization"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MultiTenantOrganization: %+v", err)
	}

	return encoded, nil
}
