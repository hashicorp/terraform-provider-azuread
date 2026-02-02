package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DirectoryObject = MultiTenantOrganizationMember{}

type MultiTenantOrganizationMember struct {
	// Tenant ID of the tenant that added the tenant to the multitenant organization. Read-only.
	AddedByTenantId nullable.Type[string] `json:"addedByTenantId,omitempty"`

	// Date and time when the tenant was added to the multitenant organization. Read-only.
	AddedDateTime nullable.Type[string] `json:"addedDateTime,omitempty"`

	// Display name of the tenant added to the multitenant organization.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Date and time when the tenant joined the multitenant organization. Read-only.
	JoinedDateTime nullable.Type[string] `json:"joinedDateTime,omitempty"`

	// Role of the tenant in the multitenant organization. The possible values are: owner, member (default),
	// unknownFutureValue. Tenants with the owner role can manage the multitenant organization but tenants with the member
	// role can only participate in a multitenant organization. There can be multiple tenants with the owner role in a
	// multitenant organization.
	Role *MultiTenantOrganizationMemberRole `json:"role,omitempty"`

	// State of the tenant in the multitenant organization. The possible values are: pending, active, removed,
	// unknownFutureValue. Tenants in the pending state must join the multitenant organization to participate in the
	// multitenant organization. Tenants in the active state can participate in the multitenant organization. Tenants in the
	// removed state are in the process of being removed from the multitenant organization. Read-only.
	State *MultiTenantOrganizationMemberState `json:"state,omitempty"`

	// Tenant ID of the Microsoft Entra tenant added to the multitenant organization. Set at the time tenant is
	// added.Supports $filter. Key.
	TenantId *string `json:"tenantId,omitempty"`

	// Details of the processing status for a tenant in a multitenant organization. Read-only. Nullable.
	TransitionDetails *MultiTenantOrganizationMemberTransitionDetails `json:"transitionDetails,omitempty"`

	// Fields inherited from DirectoryObject

	// Date and time when this object was deleted. Always null when the object hasn't been deleted.
	DeletedDateTime nullable.Type[string] `json:"deletedDateTime,omitempty"`

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

func (s MultiTenantOrganizationMember) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s MultiTenantOrganizationMember) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MultiTenantOrganizationMember{}

func (s MultiTenantOrganizationMember) MarshalJSON() ([]byte, error) {
	type wrapper MultiTenantOrganizationMember
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MultiTenantOrganizationMember: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MultiTenantOrganizationMember: %+v", err)
	}

	delete(decoded, "addedByTenantId")
	delete(decoded, "addedDateTime")
	delete(decoded, "joinedDateTime")
	delete(decoded, "state")
	delete(decoded, "transitionDetails")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.multiTenantOrganizationMember"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MultiTenantOrganizationMember: %+v", err)
	}

	return encoded, nil
}
