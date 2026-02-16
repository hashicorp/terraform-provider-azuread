package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsDriverUpdateProfileIdAssignmentId{}

// DeviceManagementWindowsDriverUpdateProfileIdAssignmentId is a struct representing the Resource ID for a Device Management Windows Driver Update Profile Id Assignment
type DeviceManagementWindowsDriverUpdateProfileIdAssignmentId struct {
	WindowsDriverUpdateProfileId           string
	WindowsDriverUpdateProfileAssignmentId string
}

// NewDeviceManagementWindowsDriverUpdateProfileIdAssignmentID returns a new DeviceManagementWindowsDriverUpdateProfileIdAssignmentId struct
func NewDeviceManagementWindowsDriverUpdateProfileIdAssignmentID(windowsDriverUpdateProfileId string, windowsDriverUpdateProfileAssignmentId string) DeviceManagementWindowsDriverUpdateProfileIdAssignmentId {
	return DeviceManagementWindowsDriverUpdateProfileIdAssignmentId{
		WindowsDriverUpdateProfileId:           windowsDriverUpdateProfileId,
		WindowsDriverUpdateProfileAssignmentId: windowsDriverUpdateProfileAssignmentId,
	}
}

// ParseDeviceManagementWindowsDriverUpdateProfileIdAssignmentID parses 'input' into a DeviceManagementWindowsDriverUpdateProfileIdAssignmentId
func ParseDeviceManagementWindowsDriverUpdateProfileIdAssignmentID(input string) (*DeviceManagementWindowsDriverUpdateProfileIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsDriverUpdateProfileIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsDriverUpdateProfileIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementWindowsDriverUpdateProfileIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementWindowsDriverUpdateProfileIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementWindowsDriverUpdateProfileIdAssignmentIDInsensitively(input string) (*DeviceManagementWindowsDriverUpdateProfileIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsDriverUpdateProfileIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsDriverUpdateProfileIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementWindowsDriverUpdateProfileIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WindowsDriverUpdateProfileId, ok = input.Parsed["windowsDriverUpdateProfileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsDriverUpdateProfileId", input)
	}

	if id.WindowsDriverUpdateProfileAssignmentId, ok = input.Parsed["windowsDriverUpdateProfileAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsDriverUpdateProfileAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementWindowsDriverUpdateProfileIdAssignmentID checks that 'input' can be parsed as a Device Management Windows Driver Update Profile Id Assignment ID
func ValidateDeviceManagementWindowsDriverUpdateProfileIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementWindowsDriverUpdateProfileIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Windows Driver Update Profile Id Assignment ID
func (id DeviceManagementWindowsDriverUpdateProfileIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/windowsDriverUpdateProfiles/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.WindowsDriverUpdateProfileId, id.WindowsDriverUpdateProfileAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Windows Driver Update Profile Id Assignment ID
func (id DeviceManagementWindowsDriverUpdateProfileIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("windowsDriverUpdateProfiles", "windowsDriverUpdateProfiles", "windowsDriverUpdateProfiles"),
		resourceids.UserSpecifiedSegment("windowsDriverUpdateProfileId", "windowsDriverUpdateProfileId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("windowsDriverUpdateProfileAssignmentId", "windowsDriverUpdateProfileAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Windows Driver Update Profile Id Assignment ID
func (id DeviceManagementWindowsDriverUpdateProfileIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Windows Driver Update Profile: %q", id.WindowsDriverUpdateProfileId),
		fmt.Sprintf("Windows Driver Update Profile Assignment: %q", id.WindowsDriverUpdateProfileAssignmentId),
	}
	return fmt.Sprintf("Device Management Windows Driver Update Profile Id Assignment (%s)", strings.Join(components, "\n"))
}
