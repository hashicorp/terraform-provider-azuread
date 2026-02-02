package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultId{}

// IdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultId is a struct representing the Resource ID for a Identity Governance Lifecycle Workflow Workflow Id Task Id Task Processing Result
type IdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultId struct {
	WorkflowId             string
	TaskId                 string
	TaskProcessingResultId string
}

// NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultID returns a new IdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultId struct
func NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultID(workflowId string, taskId string, taskProcessingResultId string) IdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultId {
	return IdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultId{
		WorkflowId:             workflowId,
		TaskId:                 taskId,
		TaskProcessingResultId: taskProcessingResultId,
	}
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultID parses 'input' into a IdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultId
func ParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultID(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultIDInsensitively(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WorkflowId, ok = input.Parsed["workflowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workflowId", input)
	}

	if id.TaskId, ok = input.Parsed["taskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "taskId", input)
	}

	if id.TaskProcessingResultId, ok = input.Parsed["taskProcessingResultId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "taskProcessingResultId", input)
	}

	return nil
}

// ValidateIdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultID checks that 'input' can be parsed as a Identity Governance Lifecycle Workflow Workflow Id Task Id Task Processing Result ID
func ValidateIdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Lifecycle Workflow Workflow Id Task Id Task Processing Result ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultId) ID() string {
	fmtString := "/identityGovernance/lifecycleWorkflows/workflows/%s/tasks/%s/taskProcessingResults/%s"
	return fmt.Sprintf(fmtString, id.WorkflowId, id.TaskId, id.TaskProcessingResultId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Lifecycle Workflow Workflow Id Task Id Task Processing Result ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("lifecycleWorkflows", "lifecycleWorkflows", "lifecycleWorkflows"),
		resourceids.StaticSegment("workflows", "workflows", "workflows"),
		resourceids.UserSpecifiedSegment("workflowId", "workflowId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("taskId", "taskId"),
		resourceids.StaticSegment("taskProcessingResults", "taskProcessingResults", "taskProcessingResults"),
		resourceids.UserSpecifiedSegment("taskProcessingResultId", "taskProcessingResultId"),
	}
}

// String returns a human-readable description of this Identity Governance Lifecycle Workflow Workflow Id Task Id Task Processing Result ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultId) String() string {
	components := []string{
		fmt.Sprintf("Workflow: %q", id.WorkflowId),
		fmt.Sprintf("Task: %q", id.TaskId),
		fmt.Sprintf("Task Processing Result: %q", id.TaskProcessingResultId),
	}
	return fmt.Sprintf("Identity Governance Lifecycle Workflow Workflow Id Task Id Task Processing Result (%s)", strings.Join(components, "\n"))
}
