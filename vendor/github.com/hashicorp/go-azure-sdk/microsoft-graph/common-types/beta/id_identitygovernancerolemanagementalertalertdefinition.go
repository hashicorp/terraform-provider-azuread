package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceRoleManagementAlertAlertDefinitionId{}

// IdentityGovernanceRoleManagementAlertAlertDefinitionId is a struct representing the Resource ID for a Identity Governance Role Management Alert Alert Definition
type IdentityGovernanceRoleManagementAlertAlertDefinitionId struct {
	UnifiedRoleManagementAlertDefinitionId string
}

// NewIdentityGovernanceRoleManagementAlertAlertDefinitionID returns a new IdentityGovernanceRoleManagementAlertAlertDefinitionId struct
func NewIdentityGovernanceRoleManagementAlertAlertDefinitionID(unifiedRoleManagementAlertDefinitionId string) IdentityGovernanceRoleManagementAlertAlertDefinitionId {
	return IdentityGovernanceRoleManagementAlertAlertDefinitionId{
		UnifiedRoleManagementAlertDefinitionId: unifiedRoleManagementAlertDefinitionId,
	}
}

// ParseIdentityGovernanceRoleManagementAlertAlertDefinitionID parses 'input' into a IdentityGovernanceRoleManagementAlertAlertDefinitionId
func ParseIdentityGovernanceRoleManagementAlertAlertDefinitionID(input string) (*IdentityGovernanceRoleManagementAlertAlertDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceRoleManagementAlertAlertDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceRoleManagementAlertAlertDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceRoleManagementAlertAlertDefinitionIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceRoleManagementAlertAlertDefinitionId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceRoleManagementAlertAlertDefinitionIDInsensitively(input string) (*IdentityGovernanceRoleManagementAlertAlertDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceRoleManagementAlertAlertDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceRoleManagementAlertAlertDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceRoleManagementAlertAlertDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleManagementAlertDefinitionId, ok = input.Parsed["unifiedRoleManagementAlertDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleManagementAlertDefinitionId", input)
	}

	return nil
}

// ValidateIdentityGovernanceRoleManagementAlertAlertDefinitionID checks that 'input' can be parsed as a Identity Governance Role Management Alert Alert Definition ID
func ValidateIdentityGovernanceRoleManagementAlertAlertDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceRoleManagementAlertAlertDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Role Management Alert Alert Definition ID
func (id IdentityGovernanceRoleManagementAlertAlertDefinitionId) ID() string {
	fmtString := "/identityGovernance/roleManagementAlerts/alertDefinitions/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleManagementAlertDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Role Management Alert Alert Definition ID
func (id IdentityGovernanceRoleManagementAlertAlertDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("roleManagementAlerts", "roleManagementAlerts", "roleManagementAlerts"),
		resourceids.StaticSegment("alertDefinitions", "alertDefinitions", "alertDefinitions"),
		resourceids.UserSpecifiedSegment("unifiedRoleManagementAlertDefinitionId", "unifiedRoleManagementAlertDefinitionId"),
	}
}

// String returns a human-readable description of this Identity Governance Role Management Alert Alert Definition ID
func (id IdentityGovernanceRoleManagementAlertAlertDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Management Alert Definition: %q", id.UnifiedRoleManagementAlertDefinitionId),
	}
	return fmt.Sprintf("Identity Governance Role Management Alert Alert Definition (%s)", strings.Join(components, "\n"))
}
