package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId{}

// IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId is a struct representing the Resource ID for a Identity Governance Lifecycle Workflow Workflow Id Task Report Id Task Processing Result
type IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId struct {
	WorkflowId             string
	TaskReportId           string
	TaskProcessingResultId string
}

// NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultID returns a new IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId struct
func NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultID(workflowId string, taskReportId string, taskProcessingResultId string) IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId {
	return IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId{
		WorkflowId:             workflowId,
		TaskReportId:           taskReportId,
		TaskProcessingResultId: taskProcessingResultId,
	}
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultID parses 'input' into a IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId
func ParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultID(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultIDInsensitively(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WorkflowId, ok = input.Parsed["workflowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workflowId", input)
	}

	if id.TaskReportId, ok = input.Parsed["taskReportId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "taskReportId", input)
	}

	if id.TaskProcessingResultId, ok = input.Parsed["taskProcessingResultId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "taskProcessingResultId", input)
	}

	return nil
}

// ValidateIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultID checks that 'input' can be parsed as a Identity Governance Lifecycle Workflow Workflow Id Task Report Id Task Processing Result ID
func ValidateIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Lifecycle Workflow Workflow Id Task Report Id Task Processing Result ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId) ID() string {
	fmtString := "/identityGovernance/lifecycleWorkflows/workflows/%s/taskReports/%s/taskProcessingResults/%s"
	return fmt.Sprintf(fmtString, id.WorkflowId, id.TaskReportId, id.TaskProcessingResultId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Lifecycle Workflow Workflow Id Task Report Id Task Processing Result ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("lifecycleWorkflows", "lifecycleWorkflows", "lifecycleWorkflows"),
		resourceids.StaticSegment("workflows", "workflows", "workflows"),
		resourceids.UserSpecifiedSegment("workflowId", "workflowId"),
		resourceids.StaticSegment("taskReports", "taskReports", "taskReports"),
		resourceids.UserSpecifiedSegment("taskReportId", "taskReportId"),
		resourceids.StaticSegment("taskProcessingResults", "taskProcessingResults", "taskProcessingResults"),
		resourceids.UserSpecifiedSegment("taskProcessingResultId", "taskProcessingResultId"),
	}
}

// String returns a human-readable description of this Identity Governance Lifecycle Workflow Workflow Id Task Report Id Task Processing Result ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId) String() string {
	components := []string{
		fmt.Sprintf("Workflow: %q", id.WorkflowId),
		fmt.Sprintf("Task Report: %q", id.TaskReportId),
		fmt.Sprintf("Task Processing Result: %q", id.TaskProcessingResultId),
	}
	return fmt.Sprintf("Identity Governance Lifecycle Workflow Workflow Id Task Report Id Task Processing Result (%s)", strings.Join(components, "\n"))
}
