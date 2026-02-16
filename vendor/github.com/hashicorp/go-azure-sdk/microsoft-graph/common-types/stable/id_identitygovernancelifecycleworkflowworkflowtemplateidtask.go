package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId{}

// IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId is a struct representing the Resource ID for a Identity Governance Lifecycle Workflow Workflow Template Id Task
type IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId struct {
	WorkflowTemplateId string
	TaskId             string
}

// NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskID returns a new IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId struct
func NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskID(workflowTemplateId string, taskId string) IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId {
	return IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId{
		WorkflowTemplateId: workflowTemplateId,
		TaskId:             taskId,
	}
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskID parses 'input' into a IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId
func ParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskID(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIDInsensitively(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WorkflowTemplateId, ok = input.Parsed["workflowTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workflowTemplateId", input)
	}

	if id.TaskId, ok = input.Parsed["taskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "taskId", input)
	}

	return nil
}

// ValidateIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskID checks that 'input' can be parsed as a Identity Governance Lifecycle Workflow Workflow Template Id Task ID
func ValidateIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Lifecycle Workflow Workflow Template Id Task ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId) ID() string {
	fmtString := "/identityGovernance/lifecycleWorkflows/workflowTemplates/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.WorkflowTemplateId, id.TaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Lifecycle Workflow Workflow Template Id Task ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("lifecycleWorkflows", "lifecycleWorkflows", "lifecycleWorkflows"),
		resourceids.StaticSegment("workflowTemplates", "workflowTemplates", "workflowTemplates"),
		resourceids.UserSpecifiedSegment("workflowTemplateId", "workflowTemplateId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("taskId", "taskId"),
	}
}

// String returns a human-readable description of this Identity Governance Lifecycle Workflow Workflow Template Id Task ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId) String() string {
	components := []string{
		fmt.Sprintf("Workflow Template: %q", id.WorkflowTemplateId),
		fmt.Sprintf("Task: %q", id.TaskId),
	}
	return fmt.Sprintf("Identity Governance Lifecycle Workflow Workflow Template Id Task (%s)", strings.Join(components, "\n"))
}
