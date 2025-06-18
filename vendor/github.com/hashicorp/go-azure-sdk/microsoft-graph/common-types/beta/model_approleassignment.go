package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DirectoryObject = AppRoleAssignment{}

type AppRoleAssignment struct {
	// The identifier (id) for the app role that is assigned to the principal. This app role must be exposed in the appRoles
	// property on the resource application's service principal (resourceId). If the resource application hasn't declared
	// any app roles, a default app role ID of 00000000-0000-0000-0000-000000000000 can be specified to signal that the
	// principal is assigned to the resource app without any specific app roles. Required on create.
	AppRoleId *string `json:"appRoleId,omitempty"`

	// The time when the app role assignment was created. The Timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreationTimestamp nullable.Type[string] `json:"creationTimestamp,omitempty"`

	// The display name of the user, group, or service principal that was granted the app role assignment. Maximum length is
	// 256 characters. Read-only. Supports $filter (eq and startswith).
	PrincipalDisplayName nullable.Type[string] `json:"principalDisplayName,omitempty"`

	// The unique identifier (id) for the user, security group, or service principal being granted the app role. Security
	// groups with dynamic memberships are supported. Required on create.
	PrincipalId nullable.Type[string] `json:"principalId,omitempty"`

	// The type of the assigned principal. This can either be User, Group, or ServicePrincipal. Read-only.
	PrincipalType nullable.Type[string] `json:"principalType,omitempty"`

	// The display name of the resource app's service principal to which the assignment is made. Maximum length is 256
	// characters.
	ResourceDisplayName nullable.Type[string] `json:"resourceDisplayName,omitempty"`

	// The unique identifier (id) for the resource service principal for which the assignment is made. Required on create.
	// Supports $filter (eq only).
	ResourceId nullable.Type[string] `json:"resourceId,omitempty"`

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

func (s AppRoleAssignment) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s AppRoleAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AppRoleAssignment{}

func (s AppRoleAssignment) MarshalJSON() ([]byte, error) {
	type wrapper AppRoleAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AppRoleAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AppRoleAssignment: %+v", err)
	}

	delete(decoded, "creationTimestamp")
	delete(decoded, "principalDisplayName")
	delete(decoded, "principalType")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.appRoleAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AppRoleAssignment: %+v", err)
	}

	return encoded, nil
}
