package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportId{}

// IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportId is a struct representing the Resource ID for a Identity Governance Lifecycle Workflow Workflow Id Task Report
type IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportId struct {
	WorkflowId   string
	TaskReportId string
}

// NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportID returns a new IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportId struct
func NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportID(workflowId string, taskReportId string) IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportId {
	return IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportId{
		WorkflowId:   workflowId,
		TaskReportId: taskReportId,
	}
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportID parses 'input' into a IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportId
func ParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportID(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIDInsensitively(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WorkflowId, ok = input.Parsed["workflowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workflowId", input)
	}

	if id.TaskReportId, ok = input.Parsed["taskReportId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "taskReportId", input)
	}

	return nil
}

// ValidateIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportID checks that 'input' can be parsed as a Identity Governance Lifecycle Workflow Workflow Id Task Report ID
func ValidateIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Lifecycle Workflow Workflow Id Task Report ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportId) ID() string {
	fmtString := "/identityGovernance/lifecycleWorkflows/workflows/%s/taskReports/%s"
	return fmt.Sprintf(fmtString, id.WorkflowId, id.TaskReportId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Lifecycle Workflow Workflow Id Task Report ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("lifecycleWorkflows", "lifecycleWorkflows", "lifecycleWorkflows"),
		resourceids.StaticSegment("workflows", "workflows", "workflows"),
		resourceids.UserSpecifiedSegment("workflowId", "workflowId"),
		resourceids.StaticSegment("taskReports", "taskReports", "taskReports"),
		resourceids.UserSpecifiedSegment("taskReportId", "taskReportId"),
	}
}

// String returns a human-readable description of this Identity Governance Lifecycle Workflow Workflow Id Task Report ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportId) String() string {
	components := []string{
		fmt.Sprintf("Workflow: %q", id.WorkflowId),
		fmt.Sprintf("Task Report: %q", id.TaskReportId),
	}
	return fmt.Sprintf("Identity Governance Lifecycle Workflow Workflow Id Task Report (%s)", strings.Join(components, "\n"))
}
