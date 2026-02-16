package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IdentityGovernanceLifecycleWorkflowsContainer{}

type IdentityGovernanceLifecycleWorkflowsContainer struct {
	// The customTaskExtension instance.
	CustomTaskExtensions *[]IdentityGovernanceCustomTaskExtension `json:"customTaskExtensions,omitempty"`

	// Deleted workflows in your lifecycle workflows instance.
	DeletedItems *DeletedItemContainer `json:"deletedItems,omitempty"`

	// The insight container holding workflow insight summaries for a tenant.
	Insights *IdentityGovernanceInsights `json:"insights,omitempty"`

	Settings *IdentityGovernanceLifecycleManagementSettings `json:"settings,omitempty"`

	// The definition of tasks within the lifecycle workflows instance.
	TaskDefinitions *[]IdentityGovernanceTaskDefinition `json:"taskDefinitions,omitempty"`

	// The workflow templates in the lifecycle workflow instance.
	WorkflowTemplates *[]IdentityGovernanceWorkflowTemplate `json:"workflowTemplates,omitempty"`

	// The workflows in the lifecycle workflows instance.
	Workflows *[]IdentityGovernanceWorkflow `json:"workflows,omitempty"`

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

func (s IdentityGovernanceLifecycleWorkflowsContainer) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityGovernanceLifecycleWorkflowsContainer{}

func (s IdentityGovernanceLifecycleWorkflowsContainer) MarshalJSON() ([]byte, error) {
	type wrapper IdentityGovernanceLifecycleWorkflowsContainer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityGovernanceLifecycleWorkflowsContainer: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceLifecycleWorkflowsContainer: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityGovernance.lifecycleWorkflowsContainer"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityGovernanceLifecycleWorkflowsContainer: %+v", err)
	}

	return encoded, nil
}
