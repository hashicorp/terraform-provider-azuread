package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PrivilegedAccess{}

type PrivilegedAccess struct {
	// The display name of the provider managed by PIM.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// A collection of resources for the provider.
	Resources *[]GovernanceResource `json:"resources,omitempty"`

	// A collection of role assignment requests for the provider.
	RoleAssignmentRequests *[]GovernanceRoleAssignmentRequest `json:"roleAssignmentRequests,omitempty"`

	// A collection of role assignments for the provider.
	RoleAssignments *[]GovernanceRoleAssignment `json:"roleAssignments,omitempty"`

	// A collection of role definitions for the provider.
	RoleDefinitions *[]GovernanceRoleDefinition `json:"roleDefinitions,omitempty"`

	// A collection of role settings for the provider.
	RoleSettings *[]GovernanceRoleSetting `json:"roleSettings,omitempty"`

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

func (s PrivilegedAccess) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrivilegedAccess{}

func (s PrivilegedAccess) MarshalJSON() ([]byte, error) {
	type wrapper PrivilegedAccess
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrivilegedAccess: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrivilegedAccess: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.privilegedAccess"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrivilegedAccess: %+v", err)
	}

	return encoded, nil
}
