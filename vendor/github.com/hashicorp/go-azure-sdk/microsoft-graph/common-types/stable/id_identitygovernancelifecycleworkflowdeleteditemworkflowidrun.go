package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId{}

// IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId is a struct representing the Resource ID for a Identity Governance Lifecycle Workflow Deleted Item Workflow Id Run
type IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId struct {
	WorkflowId string
	RunId      string
}

// NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunID returns a new IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId struct
func NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunID(workflowId string, runId string) IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId {
	return IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId{
		WorkflowId: workflowId,
		RunId:      runId,
	}
}

// ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunID parses 'input' into a IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId
func ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunID(input string) (*IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIDInsensitively(input string) (*IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WorkflowId, ok = input.Parsed["workflowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workflowId", input)
	}

	if id.RunId, ok = input.Parsed["runId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "runId", input)
	}

	return nil
}

// ValidateIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunID checks that 'input' can be parsed as a Identity Governance Lifecycle Workflow Deleted Item Workflow Id Run ID
func ValidateIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Lifecycle Workflow Deleted Item Workflow Id Run ID
func (id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId) ID() string {
	fmtString := "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/%s/runs/%s"
	return fmt.Sprintf(fmtString, id.WorkflowId, id.RunId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Lifecycle Workflow Deleted Item Workflow Id Run ID
func (id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("lifecycleWorkflows", "lifecycleWorkflows", "lifecycleWorkflows"),
		resourceids.StaticSegment("deletedItems", "deletedItems", "deletedItems"),
		resourceids.StaticSegment("workflows", "workflows", "workflows"),
		resourceids.UserSpecifiedSegment("workflowId", "workflowId"),
		resourceids.StaticSegment("runs", "runs", "runs"),
		resourceids.UserSpecifiedSegment("runId", "runId"),
	}
}

// String returns a human-readable description of this Identity Governance Lifecycle Workflow Deleted Item Workflow Id Run ID
func (id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId) String() string {
	components := []string{
		fmt.Sprintf("Workflow: %q", id.WorkflowId),
		fmt.Sprintf("Run: %q", id.RunId),
	}
	return fmt.Sprintf("Identity Governance Lifecycle Workflow Deleted Item Workflow Id Run (%s)", strings.Join(components, "\n"))
}
