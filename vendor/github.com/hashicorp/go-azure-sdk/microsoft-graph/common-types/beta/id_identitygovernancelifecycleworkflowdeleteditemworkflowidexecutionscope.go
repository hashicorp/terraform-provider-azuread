package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeId{}

// IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeId is a struct representing the Resource ID for a Identity Governance Lifecycle Workflow Deleted Item Workflow Id Execution Scope
type IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeId struct {
	WorkflowId             string
	UserProcessingResultId string
}

// NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeID returns a new IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeId struct
func NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeID(workflowId string, userProcessingResultId string) IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeId {
	return IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeId{
		WorkflowId:             workflowId,
		UserProcessingResultId: userProcessingResultId,
	}
}

// ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeID parses 'input' into a IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeId
func ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeID(input string) (*IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeIDInsensitively(input string) (*IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WorkflowId, ok = input.Parsed["workflowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workflowId", input)
	}

	if id.UserProcessingResultId, ok = input.Parsed["userProcessingResultId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userProcessingResultId", input)
	}

	return nil
}

// ValidateIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeID checks that 'input' can be parsed as a Identity Governance Lifecycle Workflow Deleted Item Workflow Id Execution Scope ID
func ValidateIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Lifecycle Workflow Deleted Item Workflow Id Execution Scope ID
func (id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeId) ID() string {
	fmtString := "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/%s/executionScope/%s"
	return fmt.Sprintf(fmtString, id.WorkflowId, id.UserProcessingResultId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Lifecycle Workflow Deleted Item Workflow Id Execution Scope ID
func (id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("lifecycleWorkflows", "lifecycleWorkflows", "lifecycleWorkflows"),
		resourceids.StaticSegment("deletedItems", "deletedItems", "deletedItems"),
		resourceids.StaticSegment("workflows", "workflows", "workflows"),
		resourceids.UserSpecifiedSegment("workflowId", "workflowId"),
		resourceids.StaticSegment("executionScope", "executionScope", "executionScope"),
		resourceids.UserSpecifiedSegment("userProcessingResultId", "userProcessingResultId"),
	}
}

// String returns a human-readable description of this Identity Governance Lifecycle Workflow Deleted Item Workflow Id Execution Scope ID
func (id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdExecutionScopeId) String() string {
	components := []string{
		fmt.Sprintf("Workflow: %q", id.WorkflowId),
		fmt.Sprintf("User Processing Result: %q", id.UserProcessingResultId),
	}
	return fmt.Sprintf("Identity Governance Lifecycle Workflow Deleted Item Workflow Id Execution Scope (%s)", strings.Join(components, "\n"))
}
