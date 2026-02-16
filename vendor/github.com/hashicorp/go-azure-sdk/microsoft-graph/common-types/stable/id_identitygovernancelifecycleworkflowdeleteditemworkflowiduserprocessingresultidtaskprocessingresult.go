package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultId{}

// IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultId is a struct representing the Resource ID for a Identity Governance Lifecycle Workflow Deleted Item Workflow Id User Processing Result Id Task Processing Result
type IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultId struct {
	WorkflowId             string
	UserProcessingResultId string
	TaskProcessingResultId string
}

// NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultID returns a new IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultId struct
func NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultID(workflowId string, userProcessingResultId string, taskProcessingResultId string) IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultId {
	return IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultId{
		WorkflowId:             workflowId,
		UserProcessingResultId: userProcessingResultId,
		TaskProcessingResultId: taskProcessingResultId,
	}
}

// ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultID parses 'input' into a IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultId
func ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultID(input string) (*IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultIDInsensitively(input string) (*IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WorkflowId, ok = input.Parsed["workflowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workflowId", input)
	}

	if id.UserProcessingResultId, ok = input.Parsed["userProcessingResultId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userProcessingResultId", input)
	}

	if id.TaskProcessingResultId, ok = input.Parsed["taskProcessingResultId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "taskProcessingResultId", input)
	}

	return nil
}

// ValidateIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultID checks that 'input' can be parsed as a Identity Governance Lifecycle Workflow Deleted Item Workflow Id User Processing Result Id Task Processing Result ID
func ValidateIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Lifecycle Workflow Deleted Item Workflow Id User Processing Result Id Task Processing Result ID
func (id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultId) ID() string {
	fmtString := "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/%s/userProcessingResults/%s/taskProcessingResults/%s"
	return fmt.Sprintf(fmtString, id.WorkflowId, id.UserProcessingResultId, id.TaskProcessingResultId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Lifecycle Workflow Deleted Item Workflow Id User Processing Result Id Task Processing Result ID
func (id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("lifecycleWorkflows", "lifecycleWorkflows", "lifecycleWorkflows"),
		resourceids.StaticSegment("deletedItems", "deletedItems", "deletedItems"),
		resourceids.StaticSegment("workflows", "workflows", "workflows"),
		resourceids.UserSpecifiedSegment("workflowId", "workflowId"),
		resourceids.StaticSegment("userProcessingResults", "userProcessingResults", "userProcessingResults"),
		resourceids.UserSpecifiedSegment("userProcessingResultId", "userProcessingResultId"),
		resourceids.StaticSegment("taskProcessingResults", "taskProcessingResults", "taskProcessingResults"),
		resourceids.UserSpecifiedSegment("taskProcessingResultId", "taskProcessingResultId"),
	}
}

// String returns a human-readable description of this Identity Governance Lifecycle Workflow Deleted Item Workflow Id User Processing Result Id Task Processing Result ID
func (id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultId) String() string {
	components := []string{
		fmt.Sprintf("Workflow: %q", id.WorkflowId),
		fmt.Sprintf("User Processing Result: %q", id.UserProcessingResultId),
		fmt.Sprintf("Task Processing Result: %q", id.TaskProcessingResultId),
	}
	return fmt.Sprintf("Identity Governance Lifecycle Workflow Deleted Item Workflow Id User Processing Result Id Task Processing Result (%s)", strings.Join(components, "\n"))
}
