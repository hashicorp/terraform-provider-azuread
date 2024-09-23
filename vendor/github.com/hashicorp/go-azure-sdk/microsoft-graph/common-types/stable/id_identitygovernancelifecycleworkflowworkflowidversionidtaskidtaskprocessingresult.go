package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId{}

// IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId is a struct representing the Resource ID for a Identity Governance Lifecycle Workflow Workflow Id Version Id Task Id Task Processing Result
type IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId struct {
	WorkflowId                   string
	WorkflowVersionVersionNumber string
	TaskId                       string
	TaskProcessingResultId       string
}

// NewIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultID returns a new IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId struct
func NewIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultID(workflowId string, workflowVersionVersionNumber string, taskId string, taskProcessingResultId string) IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId {
	return IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId{
		WorkflowId:                   workflowId,
		WorkflowVersionVersionNumber: workflowVersionVersionNumber,
		TaskId:                       taskId,
		TaskProcessingResultId:       taskProcessingResultId,
	}
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultID parses 'input' into a IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId
func ParseIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultID(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultIDInsensitively(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WorkflowId, ok = input.Parsed["workflowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workflowId", input)
	}

	if id.WorkflowVersionVersionNumber, ok = input.Parsed["workflowVersionVersionNumber"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workflowVersionVersionNumber", input)
	}

	if id.TaskId, ok = input.Parsed["taskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "taskId", input)
	}

	if id.TaskProcessingResultId, ok = input.Parsed["taskProcessingResultId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "taskProcessingResultId", input)
	}

	return nil
}

// ValidateIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultID checks that 'input' can be parsed as a Identity Governance Lifecycle Workflow Workflow Id Version Id Task Id Task Processing Result ID
func ValidateIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Lifecycle Workflow Workflow Id Version Id Task Id Task Processing Result ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId) ID() string {
	fmtString := "/identityGovernance/lifecycleWorkflows/workflows/%s/versions/%s/tasks/%s/taskProcessingResults/%s"
	return fmt.Sprintf(fmtString, id.WorkflowId, id.WorkflowVersionVersionNumber, id.TaskId, id.TaskProcessingResultId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Lifecycle Workflow Workflow Id Version Id Task Id Task Processing Result ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("lifecycleWorkflows", "lifecycleWorkflows", "lifecycleWorkflows"),
		resourceids.StaticSegment("workflows", "workflows", "workflows"),
		resourceids.UserSpecifiedSegment("workflowId", "workflowId"),
		resourceids.StaticSegment("versions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("workflowVersionVersionNumber", "workflowVersionVersionNumber"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("taskId", "taskId"),
		resourceids.StaticSegment("taskProcessingResults", "taskProcessingResults", "taskProcessingResults"),
		resourceids.UserSpecifiedSegment("taskProcessingResultId", "taskProcessingResultId"),
	}
}

// String returns a human-readable description of this Identity Governance Lifecycle Workflow Workflow Id Version Id Task Id Task Processing Result ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId) String() string {
	components := []string{
		fmt.Sprintf("Workflow: %q", id.WorkflowId),
		fmt.Sprintf("Workflow Version Version Number: %q", id.WorkflowVersionVersionNumber),
		fmt.Sprintf("Task: %q", id.TaskId),
		fmt.Sprintf("Task Processing Result: %q", id.TaskProcessingResultId),
	}
	return fmt.Sprintf("Identity Governance Lifecycle Workflow Workflow Id Version Id Task Id Task Processing Result (%s)", strings.Join(components, "\n"))
}
