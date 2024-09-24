package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentityGovernanceWorkflowBase = IdentityGovernanceWorkflowVersion{}

type IdentityGovernanceWorkflowVersion struct {
	// The version of the workflow.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
	VersionNumber *int64 `json:"versionNumber,omitempty"`

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

func (s IdentityGovernanceWorkflowVersion) IdentityGovernanceWorkflowBase() BaseIdentityGovernanceWorkflowBaseImpl {
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

var _ json.Marshaler = IdentityGovernanceWorkflowVersion{}

func (s IdentityGovernanceWorkflowVersion) MarshalJSON() ([]byte, error) {
	type wrapper IdentityGovernanceWorkflowVersion
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityGovernanceWorkflowVersion: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceWorkflowVersion: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityGovernance.workflowVersion"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityGovernanceWorkflowVersion: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IdentityGovernanceWorkflowVersion{}

func (s *IdentityGovernanceWorkflowVersion) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		VersionNumber        *int64                                       `json:"versionNumber,omitempty"`
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

	s.VersionNumber = decoded.VersionNumber
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
		return fmt.Errorf("unmarshaling IdentityGovernanceWorkflowVersion into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["executionConditions"]; ok {
		impl, err := UnmarshalIdentityGovernanceWorkflowExecutionConditionsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ExecutionConditions' for 'IdentityGovernanceWorkflowVersion': %+v", err)
		}
		s.ExecutionConditions = impl
	}

	return nil
}
