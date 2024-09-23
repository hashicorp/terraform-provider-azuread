package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDirectoryRoleAssignmentScheduleInstanceId{}

// RoleManagementDirectoryRoleAssignmentScheduleInstanceId is a struct representing the Resource ID for a Role Management Directory Role Assignment Schedule Instance
type RoleManagementDirectoryRoleAssignmentScheduleInstanceId struct {
	UnifiedRoleAssignmentScheduleInstanceId string
}

// NewRoleManagementDirectoryRoleAssignmentScheduleInstanceID returns a new RoleManagementDirectoryRoleAssignmentScheduleInstanceId struct
func NewRoleManagementDirectoryRoleAssignmentScheduleInstanceID(unifiedRoleAssignmentScheduleInstanceId string) RoleManagementDirectoryRoleAssignmentScheduleInstanceId {
	return RoleManagementDirectoryRoleAssignmentScheduleInstanceId{
		UnifiedRoleAssignmentScheduleInstanceId: unifiedRoleAssignmentScheduleInstanceId,
	}
}

// ParseRoleManagementDirectoryRoleAssignmentScheduleInstanceID parses 'input' into a RoleManagementDirectoryRoleAssignmentScheduleInstanceId
func ParseRoleManagementDirectoryRoleAssignmentScheduleInstanceID(input string) (*RoleManagementDirectoryRoleAssignmentScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDirectoryRoleAssignmentScheduleInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDirectoryRoleAssignmentScheduleInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDirectoryRoleAssignmentScheduleInstanceIDInsensitively parses 'input' case-insensitively into a RoleManagementDirectoryRoleAssignmentScheduleInstanceId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDirectoryRoleAssignmentScheduleInstanceIDInsensitively(input string) (*RoleManagementDirectoryRoleAssignmentScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDirectoryRoleAssignmentScheduleInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDirectoryRoleAssignmentScheduleInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDirectoryRoleAssignmentScheduleInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleAssignmentScheduleInstanceId, ok = input.Parsed["unifiedRoleAssignmentScheduleInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentScheduleInstanceId", input)
	}

	return nil
}

// ValidateRoleManagementDirectoryRoleAssignmentScheduleInstanceID checks that 'input' can be parsed as a Role Management Directory Role Assignment Schedule Instance ID
func ValidateRoleManagementDirectoryRoleAssignmentScheduleInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDirectoryRoleAssignmentScheduleInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Directory Role Assignment Schedule Instance ID
func (id RoleManagementDirectoryRoleAssignmentScheduleInstanceId) ID() string {
	fmtString := "/roleManagement/directory/roleAssignmentScheduleInstances/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentScheduleInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Directory Role Assignment Schedule Instance ID
func (id RoleManagementDirectoryRoleAssignmentScheduleInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("roleAssignmentScheduleInstances", "roleAssignmentScheduleInstances", "roleAssignmentScheduleInstances"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentScheduleInstanceId", "unifiedRoleAssignmentScheduleInstanceId"),
	}
}

// String returns a human-readable description of this Role Management Directory Role Assignment Schedule Instance ID
func (id RoleManagementDirectoryRoleAssignmentScheduleInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Schedule Instance: %q", id.UnifiedRoleAssignmentScheduleInstanceId),
	}
	return fmt.Sprintf("Role Management Directory Role Assignment Schedule Instance (%s)", strings.Join(components, "\n"))
}
