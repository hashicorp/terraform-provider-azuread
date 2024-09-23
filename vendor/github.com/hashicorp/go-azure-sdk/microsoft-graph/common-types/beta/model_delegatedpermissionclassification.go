package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DelegatedPermissionClassification{}

type DelegatedPermissionClassification struct {
	// The classification value. Possible values: low, medium (preview), high (preview). Doesn't support $filter.
	Classification *PermissionClassificationType `json:"classification,omitempty"`

	// The unique identifier (id) for the delegated permission listed in the publishedPermissionScopes collection of the
	// servicePrincipal. Required on create. Doesn't support $filter.
	PermissionId nullable.Type[string] `json:"permissionId,omitempty"`

	// The claim value (value) for the delegated permission listed in the publishedPermissionScopes collection of the
	// servicePrincipal. Doesn't support $filter.
	PermissionName nullable.Type[string] `json:"permissionName,omitempty"`

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

func (s DelegatedPermissionClassification) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DelegatedPermissionClassification{}

func (s DelegatedPermissionClassification) MarshalJSON() ([]byte, error) {
	type wrapper DelegatedPermissionClassification
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DelegatedPermissionClassification: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DelegatedPermissionClassification: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.delegatedPermissionClassification"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DelegatedPermissionClassification: %+v", err)
	}

	return encoded, nil
}
