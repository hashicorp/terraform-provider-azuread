package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDirectoryRoleAssignmentScheduleId{}

// RoleManagementDirectoryRoleAssignmentScheduleId is a struct representing the Resource ID for a Role Management Directory Role Assignment Schedule
type RoleManagementDirectoryRoleAssignmentScheduleId struct {
	UnifiedRoleAssignmentScheduleId string
}

// NewRoleManagementDirectoryRoleAssignmentScheduleID returns a new RoleManagementDirectoryRoleAssignmentScheduleId struct
func NewRoleManagementDirectoryRoleAssignmentScheduleID(unifiedRoleAssignmentScheduleId string) RoleManagementDirectoryRoleAssignmentScheduleId {
	return RoleManagementDirectoryRoleAssignmentScheduleId{
		UnifiedRoleAssignmentScheduleId: unifiedRoleAssignmentScheduleId,
	}
}

// ParseRoleManagementDirectoryRoleAssignmentScheduleID parses 'input' into a RoleManagementDirectoryRoleAssignmentScheduleId
func ParseRoleManagementDirectoryRoleAssignmentScheduleID(input string) (*RoleManagementDirectoryRoleAssignmentScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDirectoryRoleAssignmentScheduleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDirectoryRoleAssignmentScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDirectoryRoleAssignmentScheduleIDInsensitively parses 'input' case-insensitively into a RoleManagementDirectoryRoleAssignmentScheduleId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDirectoryRoleAssignmentScheduleIDInsensitively(input string) (*RoleManagementDirectoryRoleAssignmentScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDirectoryRoleAssignmentScheduleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDirectoryRoleAssignmentScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDirectoryRoleAssignmentScheduleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleAssignmentScheduleId, ok = input.Parsed["unifiedRoleAssignmentScheduleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentScheduleId", input)
	}

	return nil
}

// ValidateRoleManagementDirectoryRoleAssignmentScheduleID checks that 'input' can be parsed as a Role Management Directory Role Assignment Schedule ID
func ValidateRoleManagementDirectoryRoleAssignmentScheduleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDirectoryRoleAssignmentScheduleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Directory Role Assignment Schedule ID
func (id RoleManagementDirectoryRoleAssignmentScheduleId) ID() string {
	fmtString := "/roleManagement/directory/roleAssignmentSchedules/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentScheduleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Directory Role Assignment Schedule ID
func (id RoleManagementDirectoryRoleAssignmentScheduleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("roleAssignmentSchedules", "roleAssignmentSchedules", "roleAssignmentSchedules"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentScheduleId", "unifiedRoleAssignmentScheduleId"),
	}
}

// String returns a human-readable description of this Role Management Directory Role Assignment Schedule ID
func (id RoleManagementDirectoryRoleAssignmentScheduleId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Schedule: %q", id.UnifiedRoleAssignmentScheduleId),
	}
	return fmt.Sprintf("Role Management Directory Role Assignment Schedule (%s)", strings.Join(components, "\n"))
}
