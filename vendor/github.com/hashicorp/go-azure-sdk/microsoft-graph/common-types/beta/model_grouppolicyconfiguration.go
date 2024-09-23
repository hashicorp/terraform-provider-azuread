package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = GroupPolicyConfiguration{}

type GroupPolicyConfiguration struct {
	// The list of group assignments for the configuration.
	Assignments *[]GroupPolicyConfigurationAssignment `json:"assignments,omitempty"`

	// The date and time the object was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The list of enabled or disabled group policy definition values for the configuration.
	DefinitionValues *[]GroupPolicyDefinitionValue `json:"definitionValues,omitempty"`

	// User provided description for the resource object.
	Description nullable.Type[string] `json:"description,omitempty"`

	// User provided name for the resource object.
	DisplayName *string `json:"displayName,omitempty"`

	// The date and time the entity was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Group Policy Configuration Ingestion Type
	PolicyConfigurationIngestionType *GroupPolicyConfigurationIngestionType `json:"policyConfigurationIngestionType,omitempty"`

	// The list of scope tags for the configuration.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

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

func (s GroupPolicyConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GroupPolicyConfiguration{}

func (s GroupPolicyConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper GroupPolicyConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GroupPolicyConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupPolicyConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.groupPolicyConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GroupPolicyConfiguration: %+v", err)
	}

	return encoded, nil
}
