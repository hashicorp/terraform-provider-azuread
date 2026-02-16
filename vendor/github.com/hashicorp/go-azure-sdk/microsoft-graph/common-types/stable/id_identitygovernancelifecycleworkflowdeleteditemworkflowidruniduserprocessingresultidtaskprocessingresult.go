package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId{}

// IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId is a struct representing the Resource ID for a Identity Governance Lifecycle Workflow Deleted Item Workflow Id Run Id User Processing Result Id Task Processing Result
type IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId struct {
	WorkflowId             string
	RunId                  string
	UserProcessingResultId string
	TaskProcessingResultId string
}

// NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultID returns a new IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId struct
func NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultID(workflowId string, runId string, userProcessingResultId string, taskProcessingResultId string) IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId {
	return IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId{
		WorkflowId:             workflowId,
		RunId:                  runId,
		UserProcessingResultId: userProcessingResultId,
		TaskProcessingResultId: taskProcessingResultId,
	}
}

// ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultID parses 'input' into a IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId
func ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultID(input string) (*IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultIDInsensitively(input string) (*IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WorkflowId, ok = input.Parsed["workflowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workflowId", input)
	}

	if id.RunId, ok = input.Parsed["runId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "runId", input)
	}

	if id.UserProcessingResultId, ok = input.Parsed["userProcessingResultId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userProcessingResultId", input)
	}

	if id.TaskProcessingResultId, ok = input.Parsed["taskProcessingResultId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "taskProcessingResultId", input)
	}

	return nil
}

// ValidateIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultID checks that 'input' can be parsed as a Identity Governance Lifecycle Workflow Deleted Item Workflow Id Run Id User Processing Result Id Task Processing Result ID
func ValidateIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Lifecycle Workflow Deleted Item Workflow Id Run Id User Processing Result Id Task Processing Result ID
func (id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId) ID() string {
	fmtString := "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/%s/runs/%s/userProcessingResults/%s/taskProcessingResults/%s"
	return fmt.Sprintf(fmtString, id.WorkflowId, id.RunId, id.UserProcessingResultId, id.TaskProcessingResultId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Lifecycle Workflow Deleted Item Workflow Id Run Id User Processing Result Id Task Processing Result ID
func (id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("lifecycleWorkflows", "lifecycleWorkflows", "lifecycleWorkflows"),
		resourceids.StaticSegment("deletedItems", "deletedItems", "deletedItems"),
		resourceids.StaticSegment("workflows", "workflows", "workflows"),
		resourceids.UserSpecifiedSegment("workflowId", "workflowId"),
		resourceids.StaticSegment("runs", "runs", "runs"),
		resourceids.UserSpecifiedSegment("runId", "runId"),
		resourceids.StaticSegment("userProcessingResults", "userProcessingResults", "userProcessingResults"),
		resourceids.UserSpecifiedSegment("userProcessingResultId", "userProcessingResultId"),
		resourceids.StaticSegment("taskProcessingResults", "taskProcessingResults", "taskProcessingResults"),
		resourceids.UserSpecifiedSegment("taskProcessingResultId", "taskProcessingResultId"),
	}
}

// String returns a human-readable description of this Identity Governance Lifecycle Workflow Deleted Item Workflow Id Run Id User Processing Result Id Task Processing Result ID
func (id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId) String() string {
	components := []string{
		fmt.Sprintf("Workflow: %q", id.WorkflowId),
		fmt.Sprintf("Run: %q", id.RunId),
		fmt.Sprintf("User Processing Result: %q", id.UserProcessingResultId),
		fmt.Sprintf("Task Processing Result: %q", id.TaskProcessingResultId),
	}
	return fmt.Sprintf("Identity Governance Lifecycle Workflow Deleted Item Workflow Id Run Id User Processing Result Id Task Processing Result (%s)", strings.Join(components, "\n"))
}
