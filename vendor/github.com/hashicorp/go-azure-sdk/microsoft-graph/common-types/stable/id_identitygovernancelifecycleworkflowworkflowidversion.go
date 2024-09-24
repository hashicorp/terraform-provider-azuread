package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId{}

// IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId is a struct representing the Resource ID for a Identity Governance Lifecycle Workflow Workflow Id Version
type IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId struct {
	WorkflowId                   string
	WorkflowVersionVersionNumber string
}

// NewIdentityGovernanceLifecycleWorkflowWorkflowIdVersionID returns a new IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId struct
func NewIdentityGovernanceLifecycleWorkflowWorkflowIdVersionID(workflowId string, workflowVersionVersionNumber string) IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId {
	return IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId{
		WorkflowId:                   workflowId,
		WorkflowVersionVersionNumber: workflowVersionVersionNumber,
	}
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowIdVersionID parses 'input' into a IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId
func ParseIdentityGovernanceLifecycleWorkflowWorkflowIdVersionID(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIDInsensitively(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WorkflowId, ok = input.Parsed["workflowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workflowId", input)
	}

	if id.WorkflowVersionVersionNumber, ok = input.Parsed["workflowVersionVersionNumber"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workflowVersionVersionNumber", input)
	}

	return nil
}

// ValidateIdentityGovernanceLifecycleWorkflowWorkflowIdVersionID checks that 'input' can be parsed as a Identity Governance Lifecycle Workflow Workflow Id Version ID
func ValidateIdentityGovernanceLifecycleWorkflowWorkflowIdVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowIdVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Lifecycle Workflow Workflow Id Version ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId) ID() string {
	fmtString := "/identityGovernance/lifecycleWorkflows/workflows/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.WorkflowId, id.WorkflowVersionVersionNumber)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Lifecycle Workflow Workflow Id Version ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("lifecycleWorkflows", "lifecycleWorkflows", "lifecycleWorkflows"),
		resourceids.StaticSegment("workflows", "workflows", "workflows"),
		resourceids.UserSpecifiedSegment("workflowId", "workflowId"),
		resourceids.StaticSegment("versions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("workflowVersionVersionNumber", "workflowVersionVersionNumber"),
	}
}

// String returns a human-readable description of this Identity Governance Lifecycle Workflow Workflow Id Version ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId) String() string {
	components := []string{
		fmt.Sprintf("Workflow: %q", id.WorkflowId),
		fmt.Sprintf("Workflow Version Version Number: %q", id.WorkflowVersionVersionNumber),
	}
	return fmt.Sprintf("Identity Governance Lifecycle Workflow Workflow Id Version (%s)", strings.Join(components, "\n"))
}
