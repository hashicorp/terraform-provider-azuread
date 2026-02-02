package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IdentityGovernanceTaskDefinition{}

type IdentityGovernanceTaskDefinition struct {
	Category *IdentityGovernanceLifecycleTaskCategory `json:"category,omitempty"`

	// Defines if the workflow will continue if the task has an error.
	ContinueOnError *bool `json:"continueOnError,omitempty"`

	// The description of the taskDefinition.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the taskDefinition.Supports $filter(eq, ne) and $orderby.
	DisplayName *string `json:"displayName,omitempty"`

	// The parameters that must be supplied when creating a workflow task object.Supports $filter(any).
	Parameters *[]IdentityGovernanceParameter `json:"parameters,omitempty"`

	// The version number of the taskDefinition. New records are pushed when we add support for new parameters.Supports
	// $filter(ge, gt, le, lt, eq, ne) and $orderby.
	Version *int64 `json:"version,omitempty"`

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

func (s IdentityGovernanceTaskDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityGovernanceTaskDefinition{}

func (s IdentityGovernanceTaskDefinition) MarshalJSON() ([]byte, error) {
	type wrapper IdentityGovernanceTaskDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityGovernanceTaskDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceTaskDefinition: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityGovernance.taskDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityGovernanceTaskDefinition: %+v", err)
	}

	return encoded, nil
}
