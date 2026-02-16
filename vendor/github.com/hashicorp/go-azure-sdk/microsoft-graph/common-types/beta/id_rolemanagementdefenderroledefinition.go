package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDefenderRoleDefinitionId{}

// RoleManagementDefenderRoleDefinitionId is a struct representing the Resource ID for a Role Management Defender Role Definition
type RoleManagementDefenderRoleDefinitionId struct {
	UnifiedRoleDefinitionId string
}

// NewRoleManagementDefenderRoleDefinitionID returns a new RoleManagementDefenderRoleDefinitionId struct
func NewRoleManagementDefenderRoleDefinitionID(unifiedRoleDefinitionId string) RoleManagementDefenderRoleDefinitionId {
	return RoleManagementDefenderRoleDefinitionId{
		UnifiedRoleDefinitionId: unifiedRoleDefinitionId,
	}
}

// ParseRoleManagementDefenderRoleDefinitionID parses 'input' into a RoleManagementDefenderRoleDefinitionId
func ParseRoleManagementDefenderRoleDefinitionID(input string) (*RoleManagementDefenderRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDefenderRoleDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDefenderRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDefenderRoleDefinitionIDInsensitively parses 'input' case-insensitively into a RoleManagementDefenderRoleDefinitionId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDefenderRoleDefinitionIDInsensitively(input string) (*RoleManagementDefenderRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDefenderRoleDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDefenderRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDefenderRoleDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleDefinitionId, ok = input.Parsed["unifiedRoleDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId", input)
	}

	return nil
}

// ValidateRoleManagementDefenderRoleDefinitionID checks that 'input' can be parsed as a Role Management Defender Role Definition ID
func ValidateRoleManagementDefenderRoleDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDefenderRoleDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Defender Role Definition ID
func (id RoleManagementDefenderRoleDefinitionId) ID() string {
	fmtString := "/roleManagement/defender/roleDefinitions/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Defender Role Definition ID
func (id RoleManagementDefenderRoleDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("defender", "defender", "defender"),
		resourceids.StaticSegment("roleDefinitions", "roleDefinitions", "roleDefinitions"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId", "unifiedRoleDefinitionId"),
	}
}

// String returns a human-readable description of this Role Management Defender Role Definition ID
func (id RoleManagementDefenderRoleDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Definition: %q", id.UnifiedRoleDefinitionId),
	}
	return fmt.Sprintf("Role Management Defender Role Definition (%s)", strings.Join(components, "\n"))
}
