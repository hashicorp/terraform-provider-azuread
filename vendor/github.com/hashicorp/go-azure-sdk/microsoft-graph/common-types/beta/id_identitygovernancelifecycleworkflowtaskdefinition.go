package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowTaskDefinitionId{}

// IdentityGovernanceLifecycleWorkflowTaskDefinitionId is a struct representing the Resource ID for a Identity Governance Lifecycle Workflow Task Definition
type IdentityGovernanceLifecycleWorkflowTaskDefinitionId struct {
	TaskDefinitionId string
}

// NewIdentityGovernanceLifecycleWorkflowTaskDefinitionID returns a new IdentityGovernanceLifecycleWorkflowTaskDefinitionId struct
func NewIdentityGovernanceLifecycleWorkflowTaskDefinitionID(taskDefinitionId string) IdentityGovernanceLifecycleWorkflowTaskDefinitionId {
	return IdentityGovernanceLifecycleWorkflowTaskDefinitionId{
		TaskDefinitionId: taskDefinitionId,
	}
}

// ParseIdentityGovernanceLifecycleWorkflowTaskDefinitionID parses 'input' into a IdentityGovernanceLifecycleWorkflowTaskDefinitionId
func ParseIdentityGovernanceLifecycleWorkflowTaskDefinitionID(input string) (*IdentityGovernanceLifecycleWorkflowTaskDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowTaskDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowTaskDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceLifecycleWorkflowTaskDefinitionIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceLifecycleWorkflowTaskDefinitionId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceLifecycleWorkflowTaskDefinitionIDInsensitively(input string) (*IdentityGovernanceLifecycleWorkflowTaskDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowTaskDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowTaskDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceLifecycleWorkflowTaskDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TaskDefinitionId, ok = input.Parsed["taskDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "taskDefinitionId", input)
	}

	return nil
}

// ValidateIdentityGovernanceLifecycleWorkflowTaskDefinitionID checks that 'input' can be parsed as a Identity Governance Lifecycle Workflow Task Definition ID
func ValidateIdentityGovernanceLifecycleWorkflowTaskDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceLifecycleWorkflowTaskDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Lifecycle Workflow Task Definition ID
func (id IdentityGovernanceLifecycleWorkflowTaskDefinitionId) ID() string {
	fmtString := "/identityGovernance/lifecycleWorkflows/taskDefinitions/%s"
	return fmt.Sprintf(fmtString, id.TaskDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Lifecycle Workflow Task Definition ID
func (id IdentityGovernanceLifecycleWorkflowTaskDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("lifecycleWorkflows", "lifecycleWorkflows", "lifecycleWorkflows"),
		resourceids.StaticSegment("taskDefinitions", "taskDefinitions", "taskDefinitions"),
		resourceids.UserSpecifiedSegment("taskDefinitionId", "taskDefinitionId"),
	}
}

// String returns a human-readable description of this Identity Governance Lifecycle Workflow Task Definition ID
func (id IdentityGovernanceLifecycleWorkflowTaskDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Task Definition: %q", id.TaskDefinitionId),
	}
	return fmt.Sprintf("Identity Governance Lifecycle Workflow Task Definition (%s)", strings.Join(components, "\n"))
}
