package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId{}

// IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId is a struct representing the Resource ID for a Identity Governance Lifecycle Workflow Deleted Item Workflow
type IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId struct {
	WorkflowId string
}

// NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowID returns a new IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId struct
func NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowID(workflowId string) IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId {
	return IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId{
		WorkflowId: workflowId,
	}
}

// ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowID parses 'input' into a IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId
func ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowID(input string) (*IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIDInsensitively(input string) (*IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WorkflowId, ok = input.Parsed["workflowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workflowId", input)
	}

	return nil
}

// ValidateIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowID checks that 'input' can be parsed as a Identity Governance Lifecycle Workflow Deleted Item Workflow ID
func ValidateIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Lifecycle Workflow Deleted Item Workflow ID
func (id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId) ID() string {
	fmtString := "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/%s"
	return fmt.Sprintf(fmtString, id.WorkflowId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Lifecycle Workflow Deleted Item Workflow ID
func (id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("lifecycleWorkflows", "lifecycleWorkflows", "lifecycleWorkflows"),
		resourceids.StaticSegment("deletedItems", "deletedItems", "deletedItems"),
		resourceids.StaticSegment("workflows", "workflows", "workflows"),
		resourceids.UserSpecifiedSegment("workflowId", "workflowId"),
	}
}

// String returns a human-readable description of this Identity Governance Lifecycle Workflow Deleted Item Workflow ID
func (id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId) String() string {
	components := []string{
		fmt.Sprintf("Workflow: %q", id.WorkflowId),
	}
	return fmt.Sprintf("Identity Governance Lifecycle Workflow Deleted Item Workflow (%s)", strings.Join(components, "\n"))
}
