package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId{}

// IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId is a struct representing the Resource ID for a Identity Governance Lifecycle Workflow Custom Task Extension
type IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId struct {
	CustomTaskExtensionId string
}

// NewIdentityGovernanceLifecycleWorkflowCustomTaskExtensionID returns a new IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId struct
func NewIdentityGovernanceLifecycleWorkflowCustomTaskExtensionID(customTaskExtensionId string) IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId {
	return IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId{
		CustomTaskExtensionId: customTaskExtensionId,
	}
}

// ParseIdentityGovernanceLifecycleWorkflowCustomTaskExtensionID parses 'input' into a IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId
func ParseIdentityGovernanceLifecycleWorkflowCustomTaskExtensionID(input string) (*IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceLifecycleWorkflowCustomTaskExtensionIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceLifecycleWorkflowCustomTaskExtensionIDInsensitively(input string) (*IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CustomTaskExtensionId, ok = input.Parsed["customTaskExtensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "customTaskExtensionId", input)
	}

	return nil
}

// ValidateIdentityGovernanceLifecycleWorkflowCustomTaskExtensionID checks that 'input' can be parsed as a Identity Governance Lifecycle Workflow Custom Task Extension ID
func ValidateIdentityGovernanceLifecycleWorkflowCustomTaskExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceLifecycleWorkflowCustomTaskExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Lifecycle Workflow Custom Task Extension ID
func (id IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId) ID() string {
	fmtString := "/identityGovernance/lifecycleWorkflows/customTaskExtensions/%s"
	return fmt.Sprintf(fmtString, id.CustomTaskExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Lifecycle Workflow Custom Task Extension ID
func (id IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("lifecycleWorkflows", "lifecycleWorkflows", "lifecycleWorkflows"),
		resourceids.StaticSegment("customTaskExtensions", "customTaskExtensions", "customTaskExtensions"),
		resourceids.UserSpecifiedSegment("customTaskExtensionId", "customTaskExtensionId"),
	}
}

// String returns a human-readable description of this Identity Governance Lifecycle Workflow Custom Task Extension ID
func (id IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Custom Task Extension: %q", id.CustomTaskExtensionId),
	}
	return fmt.Sprintf("Identity Governance Lifecycle Workflow Custom Task Extension (%s)", strings.Join(components, "\n"))
}
