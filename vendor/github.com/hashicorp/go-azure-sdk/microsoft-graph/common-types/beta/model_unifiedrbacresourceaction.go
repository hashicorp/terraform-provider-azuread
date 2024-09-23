package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UnifiedRbacResourceAction{}

type UnifiedRbacResourceAction struct {
	// HTTP method for the action, such as DELETE, GET, PATCH, POST, PUT, or null. Supports $filter (eq) but not for null
	// values.
	ActionVerb nullable.Type[string] `json:"actionVerb,omitempty"`

	AuthenticationContext   *AuthenticationContextClassReference `json:"authenticationContext,omitempty"`
	AuthenticationContextId nullable.Type[string]                `json:"authenticationContextId,omitempty"`

	// Description for the action. Supports $filter (eq).
	Description nullable.Type[string] `json:"description,omitempty"`

	IsAuthenticationContextSettable nullable.Type[bool] `json:"isAuthenticationContextSettable,omitempty"`

	// Flag indicating if the action is a sensitive resource action. Applies only for actions in the microsoft.directory
	// resource namespace. Read-only. Supports $filter (eq).
	IsPrivileged nullable.Type[bool] `json:"isPrivileged,omitempty"`

	// Name for the action within the resource namespace, such as microsoft.insights/programs/update. Can include slash
	// character (/). Case insensitive. Required. Supports $filter (eq).
	Name string `json:"name"`

	ResourceScope *UnifiedRbacResourceScope `json:"resourceScope,omitempty"`

	// Not implemented.
	ResourceScopeId nullable.Type[string] `json:"resourceScopeId,omitempty"`

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

func (s UnifiedRbacResourceAction) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UnifiedRbacResourceAction{}

func (s UnifiedRbacResourceAction) MarshalJSON() ([]byte, error) {
	type wrapper UnifiedRbacResourceAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UnifiedRbacResourceAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UnifiedRbacResourceAction: %+v", err)
	}

	delete(decoded, "isPrivileged")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unifiedRbacResourceAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UnifiedRbacResourceAction: %+v", err)
	}

	return encoded, nil
}
