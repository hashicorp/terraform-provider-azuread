package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultId{}

// IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultId is a struct representing the Resource ID for a Identity Governance Lifecycle Workflow Deleted Item Workflow Id Run Id User Processing Result
type IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultId struct {
	WorkflowId             string
	RunId                  string
	UserProcessingResultId string
}

// NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultID returns a new IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultId struct
func NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultID(workflowId string, runId string, userProcessingResultId string) IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultId {
	return IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultId{
		WorkflowId:             workflowId,
		RunId:                  runId,
		UserProcessingResultId: userProcessingResultId,
	}
}

// ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultID parses 'input' into a IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultId
func ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultID(input string) (*IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIDInsensitively(input string) (*IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultID checks that 'input' can be parsed as a Identity Governance Lifecycle Workflow Deleted Item Workflow Id Run Id User Processing Result ID
func ValidateIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Lifecycle Workflow Deleted Item Workflow Id Run Id User Processing Result ID
func (id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultId) ID() string {
	fmtString := "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/%s/runs/%s/userProcessingResults/%s"
	return fmt.Sprintf(fmtString, id.WorkflowId, id.RunId, id.UserProcessingResultId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Lifecycle Workflow Deleted Item Workflow Id Run Id User Processing Result ID
func (id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this Identity Governance Lifecycle Workflow Deleted Item Workflow Id Run Id User Processing Result ID
func (id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultId) String() string {
	components := []string{
		fmt.Sprintf("Workflow: %q", id.WorkflowId),
		fmt.Sprintf("Run: %q", id.RunId),
		fmt.Sprintf("User Processing Result: %q", id.UserProcessingResultId),
	}
	return fmt.Sprintf("Identity Governance Lifecycle Workflow Deleted Item Workflow Id Run Id User Processing Result (%s)", strings.Join(components, "\n"))
}
