package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = GovernanceRoleDefinition{}

type GovernanceRoleDefinition struct {
	// The display name of the role definition.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The external ID of the role definition.
	ExternalId nullable.Type[string] `json:"externalId,omitempty"`

	// Read-only. The associated resource for the role definition.
	Resource *GovernanceResource `json:"resource,omitempty"`

	// Required. The ID of the resource associated with the role definition.
	ResourceId nullable.Type[string] `json:"resourceId,omitempty"`

	// The associated role setting for the role definition.
	RoleSetting *GovernanceRoleSetting `json:"roleSetting,omitempty"`

	// The unique identifier for the template.
	TemplateId nullable.Type[string] `json:"templateId,omitempty"`

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

func (s GovernanceRoleDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GovernanceRoleDefinition{}

func (s GovernanceRoleDefinition) MarshalJSON() ([]byte, error) {
	type wrapper GovernanceRoleDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GovernanceRoleDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GovernanceRoleDefinition: %+v", err)
	}

	delete(decoded, "resource")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.governanceRoleDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GovernanceRoleDefinition: %+v", err)
	}

	return encoded, nil
}
