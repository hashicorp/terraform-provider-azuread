package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementRoleAssignmentId{}

// DeviceManagementRoleAssignmentId is a struct representing the Resource ID for a Device Management Role Assignment
type DeviceManagementRoleAssignmentId struct {
	DeviceAndAppManagementRoleAssignmentId string
}

// NewDeviceManagementRoleAssignmentID returns a new DeviceManagementRoleAssignmentId struct
func NewDeviceManagementRoleAssignmentID(deviceAndAppManagementRoleAssignmentId string) DeviceManagementRoleAssignmentId {
	return DeviceManagementRoleAssignmentId{
		DeviceAndAppManagementRoleAssignmentId: deviceAndAppManagementRoleAssignmentId,
	}
}

// ParseDeviceManagementRoleAssignmentID parses 'input' into a DeviceManagementRoleAssignmentId
func ParseDeviceManagementRoleAssignmentID(input string) (*DeviceManagementRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementRoleAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementRoleAssignmentIDInsensitively(input string) (*DeviceManagementRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementRoleAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceAndAppManagementRoleAssignmentId, ok = input.Parsed["deviceAndAppManagementRoleAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceAndAppManagementRoleAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementRoleAssignmentID checks that 'input' can be parsed as a Device Management Role Assignment ID
func ValidateDeviceManagementRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Role Assignment ID
func (id DeviceManagementRoleAssignmentId) ID() string {
	fmtString := "/deviceManagement/roleAssignments/%s"
	return fmt.Sprintf(fmtString, id.DeviceAndAppManagementRoleAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Role Assignment ID
func (id DeviceManagementRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("roleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("deviceAndAppManagementRoleAssignmentId", "deviceAndAppManagementRoleAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Role Assignment ID
func (id DeviceManagementRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Device And App Management Role Assignment: %q", id.DeviceAndAppManagementRoleAssignmentId),
	}
	return fmt.Sprintf("Device Management Role Assignment (%s)", strings.Join(components, "\n"))
}
