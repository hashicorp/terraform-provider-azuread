package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementRoleAssignmentIdRoleScopeTagId{}

// DeviceManagementRoleAssignmentIdRoleScopeTagId is a struct representing the Resource ID for a Device Management Role Assignment Id Role Scope Tag
type DeviceManagementRoleAssignmentIdRoleScopeTagId struct {
	DeviceAndAppManagementRoleAssignmentId string
	RoleScopeTagId                         string
}

// NewDeviceManagementRoleAssignmentIdRoleScopeTagID returns a new DeviceManagementRoleAssignmentIdRoleScopeTagId struct
func NewDeviceManagementRoleAssignmentIdRoleScopeTagID(deviceAndAppManagementRoleAssignmentId string, roleScopeTagId string) DeviceManagementRoleAssignmentIdRoleScopeTagId {
	return DeviceManagementRoleAssignmentIdRoleScopeTagId{
		DeviceAndAppManagementRoleAssignmentId: deviceAndAppManagementRoleAssignmentId,
		RoleScopeTagId:                         roleScopeTagId,
	}
}

// ParseDeviceManagementRoleAssignmentIdRoleScopeTagID parses 'input' into a DeviceManagementRoleAssignmentIdRoleScopeTagId
func ParseDeviceManagementRoleAssignmentIdRoleScopeTagID(input string) (*DeviceManagementRoleAssignmentIdRoleScopeTagId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementRoleAssignmentIdRoleScopeTagId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementRoleAssignmentIdRoleScopeTagId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementRoleAssignmentIdRoleScopeTagIDInsensitively parses 'input' case-insensitively into a DeviceManagementRoleAssignmentIdRoleScopeTagId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementRoleAssignmentIdRoleScopeTagIDInsensitively(input string) (*DeviceManagementRoleAssignmentIdRoleScopeTagId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementRoleAssignmentIdRoleScopeTagId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementRoleAssignmentIdRoleScopeTagId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementRoleAssignmentIdRoleScopeTagId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceAndAppManagementRoleAssignmentId, ok = input.Parsed["deviceAndAppManagementRoleAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceAndAppManagementRoleAssignmentId", input)
	}

	if id.RoleScopeTagId, ok = input.Parsed["roleScopeTagId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "roleScopeTagId", input)
	}

	return nil
}

// ValidateDeviceManagementRoleAssignmentIdRoleScopeTagID checks that 'input' can be parsed as a Device Management Role Assignment Id Role Scope Tag ID
func ValidateDeviceManagementRoleAssignmentIdRoleScopeTagID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementRoleAssignmentIdRoleScopeTagID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Role Assignment Id Role Scope Tag ID
func (id DeviceManagementRoleAssignmentIdRoleScopeTagId) ID() string {
	fmtString := "/deviceManagement/roleAssignments/%s/roleScopeTags/%s"
	return fmt.Sprintf(fmtString, id.DeviceAndAppManagementRoleAssignmentId, id.RoleScopeTagId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Role Assignment Id Role Scope Tag ID
func (id DeviceManagementRoleAssignmentIdRoleScopeTagId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("roleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("deviceAndAppManagementRoleAssignmentId", "deviceAndAppManagementRoleAssignmentId"),
		resourceids.StaticSegment("roleScopeTags", "roleScopeTags", "roleScopeTags"),
		resourceids.UserSpecifiedSegment("roleScopeTagId", "roleScopeTagId"),
	}
}

// String returns a human-readable description of this Device Management Role Assignment Id Role Scope Tag ID
func (id DeviceManagementRoleAssignmentIdRoleScopeTagId) String() string {
	components := []string{
		fmt.Sprintf("Device And App Management Role Assignment: %q", id.DeviceAndAppManagementRoleAssignmentId),
		fmt.Sprintf("Role Scope Tag: %q", id.RoleScopeTagId),
	}
	return fmt.Sprintf("Device Management Role Assignment Id Role Scope Tag (%s)", strings.Join(components, "\n"))
}
