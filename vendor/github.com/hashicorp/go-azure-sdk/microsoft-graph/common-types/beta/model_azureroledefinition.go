package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AzureRoleDefinition{}

type AzureRoleDefinition struct {
	// Scopes at which the Azure role can be assigned. For more information about common patterns, see Understand Azure role
	// definitions: AssignableScopes. Supports $filter (eq).
	AssignableScopes *[]string `json:"assignableScopes,omitempty"`

	AzureRoleDefinitionType *AzureRoleDefinitionType `json:"azureRoleDefinitionType,omitempty"`

	// Name of the Azure role. Supports $filter (eq, contains).
	DisplayName *string `json:"displayName,omitempty"`

	// Identifier of an Azure role defined by Microsoft Azure. Alternate key. Supports $filter (eq).
	ExternalId *string `json:"externalId,omitempty"`

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

func (s AzureRoleDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AzureRoleDefinition{}

func (s AzureRoleDefinition) MarshalJSON() ([]byte, error) {
	type wrapper AzureRoleDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureRoleDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureRoleDefinition: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.azureRoleDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureRoleDefinition: %+v", err)
	}

	return encoded, nil
}
