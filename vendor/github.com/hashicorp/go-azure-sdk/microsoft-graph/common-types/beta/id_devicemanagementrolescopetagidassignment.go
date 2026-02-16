package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementRoleScopeTagIdAssignmentId{}

// DeviceManagementRoleScopeTagIdAssignmentId is a struct representing the Resource ID for a Device Management Role Scope Tag Id Assignment
type DeviceManagementRoleScopeTagIdAssignmentId struct {
	RoleScopeTagId               string
	RoleScopeTagAutoAssignmentId string
}

// NewDeviceManagementRoleScopeTagIdAssignmentID returns a new DeviceManagementRoleScopeTagIdAssignmentId struct
func NewDeviceManagementRoleScopeTagIdAssignmentID(roleScopeTagId string, roleScopeTagAutoAssignmentId string) DeviceManagementRoleScopeTagIdAssignmentId {
	return DeviceManagementRoleScopeTagIdAssignmentId{
		RoleScopeTagId:               roleScopeTagId,
		RoleScopeTagAutoAssignmentId: roleScopeTagAutoAssignmentId,
	}
}

// ParseDeviceManagementRoleScopeTagIdAssignmentID parses 'input' into a DeviceManagementRoleScopeTagIdAssignmentId
func ParseDeviceManagementRoleScopeTagIdAssignmentID(input string) (*DeviceManagementRoleScopeTagIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementRoleScopeTagIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementRoleScopeTagIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementRoleScopeTagIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementRoleScopeTagIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementRoleScopeTagIdAssignmentIDInsensitively(input string) (*DeviceManagementRoleScopeTagIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementRoleScopeTagIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementRoleScopeTagIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementRoleScopeTagIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.RoleScopeTagId, ok = input.Parsed["roleScopeTagId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "roleScopeTagId", input)
	}

	if id.RoleScopeTagAutoAssignmentId, ok = input.Parsed["roleScopeTagAutoAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "roleScopeTagAutoAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementRoleScopeTagIdAssignmentID checks that 'input' can be parsed as a Device Management Role Scope Tag Id Assignment ID
func ValidateDeviceManagementRoleScopeTagIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementRoleScopeTagIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Role Scope Tag Id Assignment ID
func (id DeviceManagementRoleScopeTagIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/roleScopeTags/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.RoleScopeTagId, id.RoleScopeTagAutoAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Role Scope Tag Id Assignment ID
func (id DeviceManagementRoleScopeTagIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("roleScopeTags", "roleScopeTags", "roleScopeTags"),
		resourceids.UserSpecifiedSegment("roleScopeTagId", "roleScopeTagId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("roleScopeTagAutoAssignmentId", "roleScopeTagAutoAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Role Scope Tag Id Assignment ID
func (id DeviceManagementRoleScopeTagIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Role Scope Tag: %q", id.RoleScopeTagId),
		fmt.Sprintf("Role Scope Tag Auto Assignment: %q", id.RoleScopeTagAutoAssignmentId),
	}
	return fmt.Sprintf("Device Management Role Scope Tag Id Assignment (%s)", strings.Join(components, "\n"))
}
