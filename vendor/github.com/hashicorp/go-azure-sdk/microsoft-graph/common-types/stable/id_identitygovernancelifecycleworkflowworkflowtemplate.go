package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowTemplateId{}

// IdentityGovernanceLifecycleWorkflowWorkflowTemplateId is a struct representing the Resource ID for a Identity Governance Lifecycle Workflow Workflow Template
type IdentityGovernanceLifecycleWorkflowWorkflowTemplateId struct {
	WorkflowTemplateId string
}

// NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateID returns a new IdentityGovernanceLifecycleWorkflowWorkflowTemplateId struct
func NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateID(workflowTemplateId string) IdentityGovernanceLifecycleWorkflowWorkflowTemplateId {
	return IdentityGovernanceLifecycleWorkflowWorkflowTemplateId{
		WorkflowTemplateId: workflowTemplateId,
	}
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateID parses 'input' into a IdentityGovernanceLifecycleWorkflowWorkflowTemplateId
func ParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateID(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceLifecycleWorkflowWorkflowTemplateId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateIDInsensitively(input string) (*IdentityGovernanceLifecycleWorkflowWorkflowTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowWorkflowTemplateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowWorkflowTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceLifecycleWorkflowWorkflowTemplateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WorkflowTemplateId, ok = input.Parsed["workflowTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workflowTemplateId", input)
	}

	return nil
}

// ValidateIdentityGovernanceLifecycleWorkflowWorkflowTemplateID checks that 'input' can be parsed as a Identity Governance Lifecycle Workflow Workflow Template ID
func ValidateIdentityGovernanceLifecycleWorkflowWorkflowTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Lifecycle Workflow Workflow Template ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowTemplateId) ID() string {
	fmtString := "/identityGovernance/lifecycleWorkflows/workflowTemplates/%s"
	return fmt.Sprintf(fmtString, id.WorkflowTemplateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Lifecycle Workflow Workflow Template ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowTemplateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("lifecycleWorkflows", "lifecycleWorkflows", "lifecycleWorkflows"),
		resourceids.StaticSegment("workflowTemplates", "workflowTemplates", "workflowTemplates"),
		resourceids.UserSpecifiedSegment("workflowTemplateId", "workflowTemplateId"),
	}
}

// String returns a human-readable description of this Identity Governance Lifecycle Workflow Workflow Template ID
func (id IdentityGovernanceLifecycleWorkflowWorkflowTemplateId) String() string {
	components := []string{
		fmt.Sprintf("Workflow Template: %q", id.WorkflowTemplateId),
	}
	return fmt.Sprintf("Identity Governance Lifecycle Workflow Workflow Template (%s)", strings.Join(components, "\n"))
}
