package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDeviceManagementRoleAssignmentId{}

// RoleManagementDeviceManagementRoleAssignmentId is a struct representing the Resource ID for a Role Management Device Management Role Assignment
type RoleManagementDeviceManagementRoleAssignmentId struct {
	UnifiedRoleAssignmentMultipleId string
}

// NewRoleManagementDeviceManagementRoleAssignmentID returns a new RoleManagementDeviceManagementRoleAssignmentId struct
func NewRoleManagementDeviceManagementRoleAssignmentID(unifiedRoleAssignmentMultipleId string) RoleManagementDeviceManagementRoleAssignmentId {
	return RoleManagementDeviceManagementRoleAssignmentId{
		UnifiedRoleAssignmentMultipleId: unifiedRoleAssignmentMultipleId,
	}
}

// ParseRoleManagementDeviceManagementRoleAssignmentID parses 'input' into a RoleManagementDeviceManagementRoleAssignmentId
func ParseRoleManagementDeviceManagementRoleAssignmentID(input string) (*RoleManagementDeviceManagementRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDeviceManagementRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDeviceManagementRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementDeviceManagementRoleAssignmentIDInsensitively parses 'input' case-insensitively into a RoleManagementDeviceManagementRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDeviceManagementRoleAssignmentIDInsensitively(input string) (*RoleManagementDeviceManagementRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementDeviceManagementRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementDeviceManagementRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementDeviceManagementRoleAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleAssignmentMultipleId, ok = input.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", input)
	}

	return nil
}

// ValidateRoleManagementDeviceManagementRoleAssignmentID checks that 'input' can be parsed as a Role Management Device Management Role Assignment ID
func ValidateRoleManagementDeviceManagementRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDeviceManagementRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Device Management Role Assignment ID
func (id RoleManagementDeviceManagementRoleAssignmentId) ID() string {
	fmtString := "/roleManagement/deviceManagement/roleAssignments/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentMultipleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Device Management Role Assignment ID
func (id RoleManagementDeviceManagementRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("roleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentMultipleId", "unifiedRoleAssignmentMultipleId"),
	}
}

// String returns a human-readable description of this Role Management Device Management Role Assignment ID
func (id RoleManagementDeviceManagementRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Multiple: %q", id.UnifiedRoleAssignmentMultipleId),
	}
	return fmt.Sprintf("Role Management Device Management Role Assignment (%s)", strings.Join(components, "\n"))
}
