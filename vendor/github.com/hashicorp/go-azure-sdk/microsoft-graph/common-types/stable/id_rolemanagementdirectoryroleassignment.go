package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDirectoryRoleAssignmentId{}

// RoleManagementDirectoryRoleAssignmentId is a struct representing the Resource ID for a Role Management Directory Role Assignment
type RoleManagementDirectoryRoleAssignmentId struct {
	UnifiedRoleAssignmentId string
}

// NewRoleManagementDirectoryRoleAssignmentID returns a new RoleManagementDirectoryRoleAssignmentId struct
func NewRoleManagementDirectoryRoleAssignmentID(unifiedRoleAssignmentId string) RoleManagementDirectoryRoleAssignmentId {
	return RoleManagementDirectoryRoleAssignmentId{
		UnifiedRoleAssignmentId: unifiedRoleAssignmentId,
	}
}

// ParseRoleManagementDirectoryRoleAssignmentID parses 'input' into a RoleManagementDirectoryRoleAssignmentId
func ParseRoleManagementDirectoryRoleAssignmentID(input string) (*RoleManagementDirectoryRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDirectoryRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDirectoryRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDirectoryRoleAssignmentIDInsensitively parses 'input' case-insensitively into a RoleManagementDirectoryRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDirectoryRoleAssignmentIDInsensitively(input string) (*RoleManagementDirectoryRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDirectoryRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDirectoryRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDirectoryRoleAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleAssignmentId, ok = input.Parsed["unifiedRoleAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentId", input)
	}

	return nil
}

// ValidateRoleManagementDirectoryRoleAssignmentID checks that 'input' can be parsed as a Role Management Directory Role Assignment ID
func ValidateRoleManagementDirectoryRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDirectoryRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Directory Role Assignment ID
func (id RoleManagementDirectoryRoleAssignmentId) ID() string {
	fmtString := "/roleManagement/directory/roleAssignments/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Directory Role Assignment ID
func (id RoleManagementDirectoryRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("roleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentId", "unifiedRoleAssignmentId"),
	}
}

// String returns a human-readable description of this Role Management Directory Role Assignment ID
func (id RoleManagementDirectoryRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment: %q", id.UnifiedRoleAssignmentId),
	}
	return fmt.Sprintf("Role Management Directory Role Assignment (%s)", strings.Join(components, "\n"))
}
