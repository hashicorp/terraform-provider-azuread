package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UnifiedRoleDefinition{}

type UnifiedRoleDefinition struct {
	// Types of principals that can be assigned the role. Read-only. The possible values are: user, servicePrincipal, group,
	// unknownFutureValue. This is a multi-valued enumeration that can contain up to three values as a comma-separated
	// string. For example, user, group. Supports $filter (eq).
	AllowedPrincipalTypes *AllowedRolePrincipalTypes `json:"allowedPrincipalTypes,omitempty"`

	// The description for the unifiedRoleDefinition. Read-only when isBuiltIn is true.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name for the unifiedRoleDefinition. Read-only when isBuiltIn is true. Required. Supports $filter (eq and
	// startsWith).
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Read-only collection of role definitions that the given role definition inherits from. Only Microsoft Entra built-in
	// roles support this attribute.
	InheritsPermissionsFrom *[]UnifiedRoleDefinition `json:"inheritsPermissionsFrom,omitempty"`

	// Flag indicating if the unifiedRoleDefinition is part of the default set included with the product or custom.
	// Read-only. Supports $filter (eq).
	IsBuiltIn nullable.Type[bool] `json:"isBuiltIn,omitempty"`

	// Flag indicating if the role is enabled for assignment. If false the role is not available for assignment. Read-only
	// when isBuiltIn is true.
	IsEnabled nullable.Type[bool] `json:"isEnabled,omitempty"`

	// Flag indicating if the role is privileged. Microsoft Entra ID defines a role as privileged if it contains at least
	// one sensitive resource action in the rolePermissions and allowedResourceActions objects. Applies only for actions in
	// the microsoft.directory resource namespace. Read-only. Supports $filter (eq).
	IsPrivileged nullable.Type[bool] `json:"isPrivileged,omitempty"`

	// List of scopes permissions granted by the role definition apply to. Currently only / is supported. Read-only when
	// isBuiltIn is true. DO NOT USE. This will be deprecated soon. Attach scope to role assignment.
	ResourceScopes *[]string `json:"resourceScopes,omitempty"`

	// List of permissions included in the role. Read-only when isBuiltIn is true. Required.
	RolePermissions []UnifiedRolePermission `json:"rolePermissions"`

	// Custom template identifier that can be set when isBuiltIn is false. This identifier is typically used if one needs an
	// identifier to be the same across different directories. Read-only when isBuiltIn is true.
	TemplateId nullable.Type[string] `json:"templateId,omitempty"`

	// Indicates the version of the unifiedRoleDefinition object. Read-only when isBuiltIn is true.
	Version nullable.Type[string] `json:"version,omitempty"`

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

func (s UnifiedRoleDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UnifiedRoleDefinition{}

func (s UnifiedRoleDefinition) MarshalJSON() ([]byte, error) {
	type wrapper UnifiedRoleDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UnifiedRoleDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UnifiedRoleDefinition: %+v", err)
	}

	delete(decoded, "allowedPrincipalTypes")
	delete(decoded, "isBuiltIn")
	delete(decoded, "isPrivileged")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unifiedRoleDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UnifiedRoleDefinition: %+v", err)
	}

	return encoded, nil
}
