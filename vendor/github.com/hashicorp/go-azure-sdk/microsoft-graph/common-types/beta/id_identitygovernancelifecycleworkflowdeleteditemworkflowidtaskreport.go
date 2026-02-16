package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId{}

// IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId is a struct representing the Resource ID for a Identity Governance Lifecycle Workflow Deleted Item Workflow Id Task Report
type IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId struct {
	WorkflowId   string
	TaskReportId string
}

// NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportID returns a new IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId struct
func NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportID(workflowId string, taskReportId string) IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId {
	return IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId{
		WorkflowId:   workflowId,
		TaskReportId: taskReportId,
	}
}

// ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportID parses 'input' into a IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId
func ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportID(input string) (*IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportIDInsensitively(input string) (*IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WorkflowId, ok = input.Parsed["workflowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workflowId", input)
	}

	if id.TaskReportId, ok = input.Parsed["taskReportId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "taskReportId", input)
	}

	return nil
}

// ValidateIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportID checks that 'input' can be parsed as a Identity Governance Lifecycle Workflow Deleted Item Workflow Id Task Report ID
func ValidateIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Lifecycle Workflow Deleted Item Workflow Id Task Report ID
func (id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId) ID() string {
	fmtString := "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/%s/taskReports/%s"
	return fmt.Sprintf(fmtString, id.WorkflowId, id.TaskReportId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Lifecycle Workflow Deleted Item Workflow Id Task Report ID
func (id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("lifecycleWorkflows", "lifecycleWorkflows", "lifecycleWorkflows"),
		resourceids.StaticSegment("deletedItems", "deletedItems", "deletedItems"),
		resourceids.StaticSegment("workflows", "workflows", "workflows"),
		resourceids.UserSpecifiedSegment("workflowId", "workflowId"),
		resourceids.StaticSegment("taskReports", "taskReports", "taskReports"),
		resourceids.UserSpecifiedSegment("taskReportId", "taskReportId"),
	}
}

// String returns a human-readable description of this Identity Governance Lifecycle Workflow Deleted Item Workflow Id Task Report ID
func (id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId) String() string {
	components := []string{
		fmt.Sprintf("Workflow: %q", id.WorkflowId),
		fmt.Sprintf("Task Report: %q", id.TaskReportId),
	}
	return fmt.Sprintf("Identity Governance Lifecycle Workflow Deleted Item Workflow Id Task Report (%s)", strings.Join(components, "\n"))
}
