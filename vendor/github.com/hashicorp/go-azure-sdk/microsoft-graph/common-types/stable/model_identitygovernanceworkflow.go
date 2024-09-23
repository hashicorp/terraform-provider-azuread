package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentityGovernanceWorkflowBase = IdentityGovernanceWorkflow{}

type IdentityGovernanceWorkflow struct {
	// When the workflow was deleted.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
	DeletedDateTime nullable.Type[string] `json:"deletedDateTime,omitempty"`

	// The unique identifier of the Microsoft Entra identity that last modified the workflow object.
	ExecutionScope *[]IdentityGovernanceUserProcessingResult `json:"executionScope,omitempty"`

	// Identifier used for individually addressing a specific workflow.Supports $filter(eq, ne) and $orderby.
	Id *string `json:"id,omitempty"`

	// The date time when the workflow is expected to run next based on the schedule interval, if there are any users
	// matching the execution conditions. Supports $filter(lt,gt) and $orderby.
	NextScheduleRunDateTime nullable.Type[string] `json:"nextScheduleRunDateTime,omitempty"`

	// Workflow runs.
	Runs *[]IdentityGovernanceRun `json:"runs,omitempty"`

	// Represents the aggregation of task execution data for tasks within a workflow object.
	TaskReports *[]IdentityGovernanceTaskReport `json:"taskReports,omitempty"`

	// Per-user workflow execution results.
	UserProcessingResults *[]IdentityGovernanceUserProcessingResult `json:"userProcessingResults,omitempty"`

	// The current version number of the workflow. Value is 1 when the workflow is first created.Supports $filter(lt, le,
	// gt, ge, eq, ne) and $orderby.
	Version nullable.Type[int64] `json:"version,omitempty"`

	// The workflow versions that are available.
	Versions *[]IdentityGovernanceWorkflowVersion `json:"versions,omitempty"`

	// Fields inherited from IdentityGovernanceWorkflowBase

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

func (s IdentityGovernanceWorkflow) IdentityGovernanceWorkflowBase() BaseIdentityGovernanceWorkflowBaseImpl {
	return BaseIdentityGovernanceWorkflowBaseImpl{
		Category:             s.Category,
		CreatedBy:            s.CreatedBy,
		CreatedDateTime:      s.CreatedDateTime,
		Description:          s.Description,
		DisplayName:          s.DisplayName,
		ExecutionConditions:  s.ExecutionConditions,
		IsEnabled:            s.IsEnabled,
		IsSchedulingEnabled:  s.IsSchedulingEnabled,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
		Tasks:                s.Tasks,
	}
}

var _ json.Marshaler = IdentityGovernanceWorkflow{}

func (s IdentityGovernanceWorkflow) MarshalJSON() ([]byte, error) {
	type wrapper IdentityGovernanceWorkflow
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityGovernanceWorkflow: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceWorkflow: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityGovernance.workflow"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityGovernanceWorkflow: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IdentityGovernanceWorkflow{}

func (s *IdentityGovernanceWorkflow) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DeletedDateTime         nullable.Type[string]                        `json:"deletedDateTime,omitempty"`
		ExecutionScope          *[]IdentityGovernanceUserProcessingResult    `json:"executionScope,omitempty"`
		Id                      *string                                      `json:"id,omitempty"`
		NextScheduleRunDateTime nullable.Type[string]                        `json:"nextScheduleRunDateTime,omitempty"`
		Runs                    *[]IdentityGovernanceRun                     `json:"runs,omitempty"`
		TaskReports             *[]IdentityGovernanceTaskReport              `json:"taskReports,omitempty"`
		UserProcessingResults   *[]IdentityGovernanceUserProcessingResult    `json:"userProcessingResults,omitempty"`
		Version                 nullable.Type[int64]                         `json:"version,omitempty"`
		Versions                *[]IdentityGovernanceWorkflowVersion         `json:"versions,omitempty"`
		Category                *IdentityGovernanceLifecycleWorkflowCategory `json:"category,omitempty"`
		CreatedBy               *User                                        `json:"createdBy,omitempty"`
		CreatedDateTime         nullable.Type[string]                        `json:"createdDateTime,omitempty"`
		Description             nullable.Type[string]                        `json:"description,omitempty"`
		DisplayName             *string                                      `json:"displayName,omitempty"`
		IsEnabled               *bool                                        `json:"isEnabled,omitempty"`
		IsSchedulingEnabled     *bool                                        `json:"isSchedulingEnabled,omitempty"`
		LastModifiedBy          *User                                        `json:"lastModifiedBy,omitempty"`
		LastModifiedDateTime    nullable.Type[string]                        `json:"lastModifiedDateTime,omitempty"`
		ODataId                 *string                                      `json:"@odata.id,omitempty"`
		ODataType               *string                                      `json:"@odata.type,omitempty"`
		Tasks                   *[]IdentityGovernanceTask                    `json:"tasks,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DeletedDateTime = decoded.DeletedDateTime
	s.ExecutionScope = decoded.ExecutionScope
	s.Id = decoded.Id
	s.NextScheduleRunDateTime = decoded.NextScheduleRunDateTime
	s.Runs = decoded.Runs
	s.TaskReports = decoded.TaskReports
	s.UserProcessingResults = decoded.UserProcessingResults
	s.Version = decoded.Version
	s.Versions = decoded.Versions
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
		return fmt.Errorf("unmarshaling IdentityGovernanceWorkflow into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["executionConditions"]; ok {
		impl, err := UnmarshalIdentityGovernanceWorkflowExecutionConditionsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ExecutionConditions' for 'IdentityGovernanceWorkflow': %+v", err)
		}
		s.ExecutionConditions = impl
	}

	return nil
}
