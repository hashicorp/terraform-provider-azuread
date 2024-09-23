package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId{}

// IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId is a struct representing the Resource ID for a Identity Governance Lifecycle Workflow Workflow Id Task
type IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId struct {
	WorkflowId string
	TaskId     string
}

// NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskID returns a new IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId struct
func NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskID(workflowId string, taskId string) IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId {
	return IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId{
		WorkflowId: workflowId,
		TaskId:     taskId,
	}
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskID parses 'input' into a IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId
func ParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskID(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskIDInsensitively(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WorkflowId, ok = input.Parsed["workflowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workflowId", input)
	}

	if id.TaskId, ok = input.Parsed["taskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "taskId", input)
	}

	return nil
}

// ValidateIdentityGovernanceLifecycleWorkflowWorkflowIdTaskID checks that 'input' can be parsed as a Identity Governance Lifecycle Workflow Workflow Id Task ID
func ValidateIdentityGovernanceLifecycleWorkflowWorkflowIdTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Lifecycle Workflow Workflow Id Task ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId) ID() string {
	fmtString := "/identityGovernance/lifecycleWorkflows/workflows/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.WorkflowId, id.TaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Lifecycle Workflow Workflow Id Task ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("lifecycleWorkflows", "lifecycleWorkflows", "lifecycleWorkflows"),
		resourceids.StaticSegment("workflows", "workflows", "workflows"),
		resourceids.UserSpecifiedSegment("workflowId", "workflowId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("taskId", "taskId"),
	}
}

// String returns a human-readable description of this Identity Governance Lifecycle Workflow Workflow Id Task ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId) String() string {
	components := []string{
		fmt.Sprintf("Workflow: %q", id.WorkflowId),
		fmt.Sprintf("Task: %q", id.TaskId),
	}
	return fmt.Sprintf("Identity Governance Lifecycle Workflow Workflow Id Task (%s)", strings.Join(components, "\n"))
}
