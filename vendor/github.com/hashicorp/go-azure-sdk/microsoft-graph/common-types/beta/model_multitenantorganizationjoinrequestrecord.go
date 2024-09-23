package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MultiTenantOrganizationJoinRequestRecord{}

type MultiTenantOrganizationJoinRequestRecord struct {
	// Tenant ID of the Microsoft Entra tenant that added a tenant to the multi-tenant organization. To reset a failed join
	// request, set addedByTenantId to 00000000-0000-0000-0000-000000000000. Required.
	AddedByTenantId nullable.Type[string] `json:"addedByTenantId,omitempty"`

	// State of the tenant in the multi-tenant organization. The possible values are: pending, active, removed,
	// unknownFutureValue. Tenants in the pending state must join the multi-tenant organization to participate in the
	// multi-tenant organization. Tenants in the active state can participate in the multi-tenant organization. Tenants in
	// the removed state are in the process of being removed from the multi-tenant organization. Read-only.
	MemberState *MultiTenantOrganizationMemberState `json:"memberState,omitempty"`

	// Role of the tenant in the multi-tenant organization. The possible values are: owner, member (default),
	// unknownFutureValue. Tenants with the owner role can manage the multi-tenant organization. There can be multiple
	// tenants with the owner role in a multi-tenant organization. Tenants with the member role can participate in a
	// multi-tenant organization.
	Role *MultiTenantOrganizationMemberRole `json:"role,omitempty"`

	// Details of the processing status for a tenant joining a multi-tenant organization. Read-only.
	TransitionDetails *MultiTenantOrganizationJoinRequestTransitionDetails `json:"transitionDetails,omitempty"`

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

func (s MultiTenantOrganizationJoinRequestRecord) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MultiTenantOrganizationJoinRequestRecord{}

func (s MultiTenantOrganizationJoinRequestRecord) MarshalJSON() ([]byte, error) {
	type wrapper MultiTenantOrganizationJoinRequestRecord
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MultiTenantOrganizationJoinRequestRecord: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MultiTenantOrganizationJoinRequestRecord: %+v", err)
	}

	delete(decoded, "memberState")
	delete(decoded, "transitionDetails")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.multiTenantOrganizationJoinRequestRecord"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MultiTenantOrganizationJoinRequestRecord: %+v", err)
	}

	return encoded, nil
}
