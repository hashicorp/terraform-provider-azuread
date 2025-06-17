package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDefenderRoleAssignmentId{}

// RoleManagementDefenderRoleAssignmentId is a struct representing the Resource ID for a Role Management Defender Role Assignment
type RoleManagementDefenderRoleAssignmentId struct {
	UnifiedRoleAssignmentMultipleId string
}

// NewRoleManagementDefenderRoleAssignmentID returns a new RoleManagementDefenderRoleAssignmentId struct
func NewRoleManagementDefenderRoleAssignmentID(unifiedRoleAssignmentMultipleId string) RoleManagementDefenderRoleAssignmentId {
	return RoleManagementDefenderRoleAssignmentId{
		UnifiedRoleAssignmentMultipleId: unifiedRoleAssignmentMultipleId,
	}
}

// ParseRoleManagementDefenderRoleAssignmentID parses 'input' into a RoleManagementDefenderRoleAssignmentId
func ParseRoleManagementDefenderRoleAssignmentID(input string) (*RoleManagementDefenderRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDefenderRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDefenderRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDefenderRoleAssignmentIDInsensitively parses 'input' case-insensitively into a RoleManagementDefenderRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDefenderRoleAssignmentIDInsensitively(input string) (*RoleManagementDefenderRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDefenderRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDefenderRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDefenderRoleAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleAssignmentMultipleId, ok = input.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", input)
	}

	return nil
}

// ValidateRoleManagementDefenderRoleAssignmentID checks that 'input' can be parsed as a Role Management Defender Role Assignment ID
func ValidateRoleManagementDefenderRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDefenderRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Defender Role Assignment ID
func (id RoleManagementDefenderRoleAssignmentId) ID() string {
	fmtString := "/roleManagement/defender/roleAssignments/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentMultipleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Defender Role Assignment ID
func (id RoleManagementDefenderRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("defender", "defender", "defender"),
		resourceids.StaticSegment("roleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentMultipleId", "unifiedRoleAssignmentMultipleId"),
	}
}

// String returns a human-readable description of this Role Management Defender Role Assignment ID
func (id RoleManagementDefenderRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Multiple: %q", id.UnifiedRoleAssignmentMultipleId),
	}
	return fmt.Sprintf("Role Management Defender Role Assignment (%s)", strings.Join(components, "\n"))
}
