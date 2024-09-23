package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeId{}

// IdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeId is a struct representing the Resource ID for a Identity Governance Lifecycle Workflow Workflow Id Execution Scope
type IdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeId struct {
	WorkflowId             string
	UserProcessingResultId string
}

// NewIdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeID returns a new IdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeId struct
func NewIdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeID(workflowId string, userProcessingResultId string) IdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeId {
	return IdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeId{
		WorkflowId:             workflowId,
		UserProcessingResultId: userProcessingResultId,
	}
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeID parses 'input' into a IdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeId
func ParseIdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeID(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeIDInsensitively(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WorkflowId, ok = input.Parsed["workflowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workflowId", input)
	}

	if id.UserProcessingResultId, ok = input.Parsed["userProcessingResultId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userProcessingResultId", input)
	}

	return nil
}

// ValidateIdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeID checks that 'input' can be parsed as a Identity Governance Lifecycle Workflow Workflow Id Execution Scope ID
func ValidateIdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Lifecycle Workflow Workflow Id Execution Scope ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeId) ID() string {
	fmtString := "/identityGovernance/lifecycleWorkflows/workflows/%s/executionScope/%s"
	return fmt.Sprintf(fmtString, id.WorkflowId, id.UserProcessingResultId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Lifecycle Workflow Workflow Id Execution Scope ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("lifecycleWorkflows", "lifecycleWorkflows", "lifecycleWorkflows"),
		resourceids.StaticSegment("workflows", "workflows", "workflows"),
		resourceids.UserSpecifiedSegment("workflowId", "workflowId"),
		resourceids.StaticSegment("executionScope", "executionScope", "executionScope"),
		resourceids.UserSpecifiedSegment("userProcessingResultId", "userProcessingResultId"),
	}
}

// String returns a human-readable description of this Identity Governance Lifecycle Workflow Workflow Id Execution Scope ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeId) String() string {
	components := []string{
		fmt.Sprintf("Workflow: %q", id.WorkflowId),
		fmt.Sprintf("User Processing Result: %q", id.UserProcessingResultId),
	}
	return fmt.Sprintf("Identity Governance Lifecycle Workflow Workflow Id Execution Scope (%s)", strings.Join(components, "\n"))
}
