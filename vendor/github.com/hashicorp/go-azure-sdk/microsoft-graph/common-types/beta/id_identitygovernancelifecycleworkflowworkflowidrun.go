package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowIdRunId{}

// IdentityGovernanceLifecycleWorkflowWorkflowIdRunId is a struct representing the Resource ID for a Identity Governance Lifecycle Workflow Workflow Id Run
type IdentityGovernanceLifecycleWorkflowWorkflowIdRunId struct {
	WorkflowId string
	RunId      string
}

// NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunID returns a new IdentityGovernanceLifecycleWorkflowWorkflowIdRunId struct
func NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunID(workflowId string, runId string) IdentityGovernanceLifecycleWorkflowWorkflowIdRunId {
	return IdentityGovernanceLifecycleWorkflowWorkflowIdRunId{
		WorkflowId: workflowId,
		RunId:      runId,
	}
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowIdRunID parses 'input' into a IdentityGovernanceLifecycleWorkflowWorkflowIdRunId
func ParseIdentityGovernanceLifecycleWorkflowWorkflowIdRunID(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowIdRunId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowIdRunId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowIdRunId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowIdRunIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceLifecycleWorkflowWorkflowIdRunId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceLifecycleWorkflowWorkflowIdRunIDInsensitively(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowIdRunId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowIdRunId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowIdRunId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceLifecycleWorkflowWorkflowIdRunId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WorkflowId, ok = input.Parsed["workflowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workflowId", input)
	}

	if id.RunId, ok = input.Parsed["runId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "runId", input)
	}

	return nil
}

// ValidateIdentityGovernanceLifecycleWorkflowWorkflowIdRunID checks that 'input' can be parsed as a Identity Governance Lifecycle Workflow Workflow Id Run ID
func ValidateIdentityGovernanceLifecycleWorkflowWorkflowIdRunID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowIdRunID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Lifecycle Workflow Workflow Id Run ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdRunId) ID() string {
	fmtString := "/identityGovernance/lifecycleWorkflows/workflows/%s/runs/%s"
	return fmt.Sprintf(fmtString, id.WorkflowId, id.RunId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Lifecycle Workflow Workflow Id Run ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdRunId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("lifecycleWorkflows", "lifecycleWorkflows", "lifecycleWorkflows"),
		resourceids.StaticSegment("workflows", "workflows", "workflows"),
		resourceids.UserSpecifiedSegment("workflowId", "workflowId"),
		resourceids.StaticSegment("runs", "runs", "runs"),
		resourceids.UserSpecifiedSegment("runId", "runId"),
	}
}

// String returns a human-readable description of this Identity Governance Lifecycle Workflow Workflow Id Run ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdRunId) String() string {
	components := []string{
		fmt.Sprintf("Workflow: %q", id.WorkflowId),
		fmt.Sprintf("Run: %q", id.RunId),
	}
	return fmt.Sprintf("Identity Governance Lifecycle Workflow Workflow Id Run (%s)", strings.Join(components, "\n"))
}
