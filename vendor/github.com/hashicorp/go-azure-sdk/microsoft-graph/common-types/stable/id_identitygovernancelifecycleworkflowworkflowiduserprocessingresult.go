package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultId{}

// IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultId is a struct representing the Resource ID for a Identity Governance Lifecycle Workflow Workflow Id User Processing Result
type IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultId struct {
	WorkflowId             string
	UserProcessingResultId string
}

// NewIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultID returns a new IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultId struct
func NewIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultID(workflowId string, userProcessingResultId string) IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultId {
	return IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultId{
		WorkflowId:             workflowId,
		UserProcessingResultId: userProcessingResultId,
	}
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultID parses 'input' into a IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultId
func ParseIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultID(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIDInsensitively(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WorkflowId, ok = input.Parsed["workflowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workflowId", input)
	}

	if id.UserProcessingResultId, ok = input.Parsed["userProcessingResultId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userProcessingResultId", input)
	}

	return nil
}

// ValidateIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultID checks that 'input' can be parsed as a Identity Governance Lifecycle Workflow Workflow Id User Processing Result ID
func ValidateIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Lifecycle Workflow Workflow Id User Processing Result ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultId) ID() string {
	fmtString := "/identityGovernance/lifecycleWorkflows/workflows/%s/userProcessingResults/%s"
	return fmt.Sprintf(fmtString, id.WorkflowId, id.UserProcessingResultId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Lifecycle Workflow Workflow Id User Processing Result ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("lifecycleWorkflows", "lifecycleWorkflows", "lifecycleWorkflows"),
		resourceids.StaticSegment("workflows", "workflows", "workflows"),
		resourceids.UserSpecifiedSegment("workflowId", "workflowId"),
		resourceids.StaticSegment("userProcessingResults", "userProcessingResults", "userProcessingResults"),
		resourceids.UserSpecifiedSegment("userProcessingResultId", "userProcessingResultId"),
	}
}

// String returns a human-readable description of this Identity Governance Lifecycle Workflow Workflow Id User Processing Result ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultId) String() string {
	components := []string{
		fmt.Sprintf("Workflow: %q", id.WorkflowId),
		fmt.Sprintf("User Processing Result: %q", id.UserProcessingResultId),
	}
	return fmt.Sprintf("Identity Governance Lifecycle Workflow Workflow Id User Processing Result (%s)", strings.Join(components, "\n"))
}
