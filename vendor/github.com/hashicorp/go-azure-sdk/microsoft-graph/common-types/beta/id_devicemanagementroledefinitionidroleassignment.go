package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementRoleDefinitionIdRoleAssignmentId{}

// DeviceManagementRoleDefinitionIdRoleAssignmentId is a struct representing the Resource ID for a Device Management Role Definition Id Role Assignment
type DeviceManagementRoleDefinitionIdRoleAssignmentId struct {
	RoleDefinitionId string
	RoleAssignmentId string
}

// NewDeviceManagementRoleDefinitionIdRoleAssignmentID returns a new DeviceManagementRoleDefinitionIdRoleAssignmentId struct
func NewDeviceManagementRoleDefinitionIdRoleAssignmentID(roleDefinitionId string, roleAssignmentId string) DeviceManagementRoleDefinitionIdRoleAssignmentId {
	return DeviceManagementRoleDefinitionIdRoleAssignmentId{
		RoleDefinitionId: roleDefinitionId,
		RoleAssignmentId: roleAssignmentId,
	}
}

// ParseDeviceManagementRoleDefinitionIdRoleAssignmentID parses 'input' into a DeviceManagementRoleDefinitionIdRoleAssignmentId
func ParseDeviceManagementRoleDefinitionIdRoleAssignmentID(input string) (*DeviceManagementRoleDefinitionIdRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementRoleDefinitionIdRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementRoleDefinitionIdRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementRoleDefinitionIdRoleAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementRoleDefinitionIdRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementRoleDefinitionIdRoleAssignmentIDInsensitively(input string) (*DeviceManagementRoleDefinitionIdRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementRoleDefinitionIdRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementRoleDefinitionIdRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementRoleDefinitionIdRoleAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.RoleDefinitionId, ok = input.Parsed["roleDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "roleDefinitionId", input)
	}

	if id.RoleAssignmentId, ok = input.Parsed["roleAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "roleAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementRoleDefinitionIdRoleAssignmentID checks that 'input' can be parsed as a Device Management Role Definition Id Role Assignment ID
func ValidateDeviceManagementRoleDefinitionIdRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementRoleDefinitionIdRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Role Definition Id Role Assignment ID
func (id DeviceManagementRoleDefinitionIdRoleAssignmentId) ID() string {
	fmtString := "/deviceManagement/roleDefinitions/%s/roleAssignments/%s"
	return fmt.Sprintf(fmtString, id.RoleDefinitionId, id.RoleAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Role Definition Id Role Assignment ID
func (id DeviceManagementRoleDefinitionIdRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("roleDefinitions", "roleDefinitions", "roleDefinitions"),
		resourceids.UserSpecifiedSegment("roleDefinitionId", "roleDefinitionId"),
		resourceids.StaticSegment("roleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("roleAssignmentId", "roleAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Role Definition Id Role Assignment ID
func (id DeviceManagementRoleDefinitionIdRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Role Definition: %q", id.RoleDefinitionId),
		fmt.Sprintf("Role Assignment: %q", id.RoleAssignmentId),
	}
	return fmt.Sprintf("Device Management Role Definition Id Role Assignment (%s)", strings.Join(components, "\n"))
}
