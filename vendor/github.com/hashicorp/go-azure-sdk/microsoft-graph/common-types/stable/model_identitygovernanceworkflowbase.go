package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceWorkflowBase interface {
	IdentityGovernanceWorkflowBase() BaseIdentityGovernanceWorkflowBaseImpl
}

var _ IdentityGovernanceWorkflowBase = BaseIdentityGovernanceWorkflowBaseImpl{}

type BaseIdentityGovernanceWorkflowBaseImpl struct {
	Category *IdentityGovernanceLifecycleWorkflowCategory `json:"category,omitempty"`

	// The user who created the workflow.
	CreatedBy *User `json:"createdBy,omitempty"`

	// When a workflow was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// A string that describes the purpose of the workflow.
	Description nullable.Type[string] `json:"description,omitempty"`

	// A string to identify the workflow.
	DisplayName *string `json:"displayName,omitempty"`

	// Defines when and for who the workflow will run.
	ExecutionConditions IdentityGovernanceWorkflowExecutionConditions `json:"executionConditions"`

	// Whether the workflow is enabled or disabled. If this setting is true, the workflow can be run on demand or on
	// schedule when isSchedulingEnabled is true.
	IsEnabled *bool `json:"isEnabled,omitempty"`

	// If true, the Lifecycle Workflow engine executes the workflow based on the schedule defined by tenant settings. Can't
	// be true for a disabled workflow (where isEnabled is false).
	IsSchedulingEnabled *bool `json:"isSchedulingEnabled,omitempty"`

	// The unique identifier of the Microsoft Entra identity that last modified the workflow.
	LastModifiedBy *User `json:"lastModifiedBy,omitempty"`

	// When the workflow was last modified.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The tasks in the workflow.
	Tasks *[]IdentityGovernanceTask `json:"tasks,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseIdentityGovernanceWorkflowBaseImpl) IdentityGovernanceWorkflowBase() BaseIdentityGovernanceWorkflowBaseImpl {
	return s
}

var _ IdentityGovernanceWorkflowBase = RawIdentityGovernanceWorkflowBaseImpl{}

// RawIdentityGovernanceWorkflowBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIdentityGovernanceWorkflowBaseImpl struct {
	identityGovernanceWorkflowBase BaseIdentityGovernanceWorkflowBaseImpl
	Type                           string
	Values                         map[string]interface{}
}

func (s RawIdentityGovernanceWorkflowBaseImpl) IdentityGovernanceWorkflowBase() BaseIdentityGovernanceWorkflowBaseImpl {
	return s.identityGovernanceWorkflowBase
}

var _ json.Unmarshaler = &BaseIdentityGovernanceWorkflowBaseImpl{}

func (s *BaseIdentityGovernanceWorkflowBaseImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Category             *IdentityGovernanceLifecycleWorkflowCategory `json:"category,omitempty"`
		CreatedBy            *User                                        `json:"createdBy,omitempty"`
		CreatedDateTime      nullable.Type[string]                        `json:"createdDateTime,omitempty"`
		Description          nullable.Type[string]                        `json:"description,omitempty"`
		DisplayName          *string                                      `json:"displayName,omitempty"`
		IsEnabled            *bool                                        `json:"isEnabled,omitempty"`
		IsSchedulingEnabled  *bool                                        `json:"isSchedulingEnabled,omitempty"`
		LastModifiedBy       *User                                        `json:"lastModifiedBy,omitempty"`
		LastModifiedDateTime nullable.Type[string]                        `json:"lastModifiedDateTime,omitempty"`
		ODataId              *string                                      `json:"@odata.id,omitempty"`
		ODataType            *string                                      `json:"@odata.type,omitempty"`
		Tasks                *[]IdentityGovernanceTask                    `json:"tasks,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Category = decoded.Category
	s.CreatedBy = decoded.CreatedBy
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.IsEnabled = decoded.IsEnabled
	s.IsSchedulingEnabled = decoded.IsSchedulingEnabled
	s.LastModifiedBy = decoded.LastModifiedBy
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Tasks = decoded.Tasks

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseIdentityGovernanceWorkflowBaseImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["executionConditions"]; ok {
		impl, err := UnmarshalIdentityGovernanceWorkflowExecutionConditionsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ExecutionConditions' for 'BaseIdentityGovernanceWorkflowBaseImpl': %+v", err)
		}
		s.ExecutionConditions = impl
	}

	return nil
}

func UnmarshalIdentityGovernanceWorkflowBaseImplementation(input []byte) (IdentityGovernanceWorkflowBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceWorkflowBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.workflow") {
		var out IdentityGovernanceWorkflow
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceWorkflow: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.workflowVersion") {
		var out IdentityGovernanceWorkflowVersion
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceWorkflowVersion: %+v", err)
		}
		return out, nil
	}

	var parent BaseIdentityGovernanceWorkflowBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIdentityGovernanceWorkflowBaseImpl: %+v", err)
	}

	return RawIdentityGovernanceWorkflowBaseImpl{
		identityGovernanceWorkflowBase: parent,
		Type:                           value,
		Values:                         temp,
	}, nil

}
