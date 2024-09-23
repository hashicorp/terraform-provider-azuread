package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowId{}

// IdentityGovernanceLifecycleWorkflowWorkflowId is a struct representing the Resource ID for a Identity Governance Lifecycle Workflow Workflow
type IdentityGovernanceLifecycleWorkflowWorkflowId struct {
	WorkflowId string
}

// NewIdentityGovernanceLifecycleWorkflowWorkflowID returns a new IdentityGovernanceLifecycleWorkflowWorkflowId struct
func NewIdentityGovernanceLifecycleWorkflowWorkflowID(workflowId string) IdentityGovernanceLifecycleWorkflowWorkflowId {
	return IdentityGovernanceLifecycleWorkflowWorkflowId{
		WorkflowId: workflowId,
	}
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowID parses 'input' into a IdentityGovernanceLifecycleWorkflowWorkflowId
func ParseIdentityGovernanceLifecycleWorkflowWorkflowID(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceLifecycleWorkflowWorkflowId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceLifecycleWorkflowWorkflowIDInsensitively(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceLifecycleWorkflowWorkflowId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WorkflowId, ok = input.Parsed["workflowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workflowId", input)
	}

	return nil
}

// ValidateIdentityGovernanceLifecycleWorkflowWorkflowID checks that 'input' can be parsed as a Identity Governance Lifecycle Workflow Workflow ID
func ValidateIdentityGovernanceLifecycleWorkflowWorkflowID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Lifecycle Workflow Workflow ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowId) ID() string {
	fmtString := "/identityGovernance/lifecycleWorkflows/workflows/%s"
	return fmt.Sprintf(fmtString, id.WorkflowId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Lifecycle Workflow Workflow ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("lifecycleWorkflows", "lifecycleWorkflows", "lifecycleWorkflows"),
		resourceids.StaticSegment("workflows", "workflows", "workflows"),
		resourceids.UserSpecifiedSegment("workflowId", "workflowId"),
	}
}

// String returns a human-readable description of this Identity Governance Lifecycle Workflow Workflow ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowId) String() string {
	components := []string{
		fmt.Sprintf("Workflow: %q", id.WorkflowId),
	}
	return fmt.Sprintf("Identity Governance Lifecycle Workflow Workflow (%s)", strings.Join(components, "\n"))
}
