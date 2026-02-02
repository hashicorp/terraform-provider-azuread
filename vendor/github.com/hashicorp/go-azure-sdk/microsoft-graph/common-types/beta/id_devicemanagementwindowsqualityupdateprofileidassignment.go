package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsQualityUpdateProfileIdAssignmentId{}

// DeviceManagementWindowsQualityUpdateProfileIdAssignmentId is a struct representing the Resource ID for a Device Management Windows Quality Update Profile Id Assignment
type DeviceManagementWindowsQualityUpdateProfileIdAssignmentId struct {
	WindowsQualityUpdateProfileId           string
	WindowsQualityUpdateProfileAssignmentId string
}

// NewDeviceManagementWindowsQualityUpdateProfileIdAssignmentID returns a new DeviceManagementWindowsQualityUpdateProfileIdAssignmentId struct
func NewDeviceManagementWindowsQualityUpdateProfileIdAssignmentID(windowsQualityUpdateProfileId string, windowsQualityUpdateProfileAssignmentId string) DeviceManagementWindowsQualityUpdateProfileIdAssignmentId {
	return DeviceManagementWindowsQualityUpdateProfileIdAssignmentId{
		WindowsQualityUpdateProfileId:           windowsQualityUpdateProfileId,
		WindowsQualityUpdateProfileAssignmentId: windowsQualityUpdateProfileAssignmentId,
	}
}

// ParseDeviceManagementWindowsQualityUpdateProfileIdAssignmentID parses 'input' into a DeviceManagementWindowsQualityUpdateProfileIdAssignmentId
func ParseDeviceManagementWindowsQualityUpdateProfileIdAssignmentID(input string) (*DeviceManagementWindowsQualityUpdateProfileIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsQualityUpdateProfileIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsQualityUpdateProfileIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementWindowsQualityUpdateProfileIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementWindowsQualityUpdateProfileIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementWindowsQualityUpdateProfileIdAssignmentIDInsensitively(input string) (*DeviceManagementWindowsQualityUpdateProfileIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsQualityUpdateProfileIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsQualityUpdateProfileIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementWindowsQualityUpdateProfileIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WindowsQualityUpdateProfileId, ok = input.Parsed["windowsQualityUpdateProfileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsQualityUpdateProfileId", input)
	}

	if id.WindowsQualityUpdateProfileAssignmentId, ok = input.Parsed["windowsQualityUpdateProfileAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsQualityUpdateProfileAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementWindowsQualityUpdateProfileIdAssignmentID checks that 'input' can be parsed as a Device Management Windows Quality Update Profile Id Assignment ID
func ValidateDeviceManagementWindowsQualityUpdateProfileIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementWindowsQualityUpdateProfileIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Windows Quality Update Profile Id Assignment ID
func (id DeviceManagementWindowsQualityUpdateProfileIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/windowsQualityUpdateProfiles/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.WindowsQualityUpdateProfileId, id.WindowsQualityUpdateProfileAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Windows Quality Update Profile Id Assignment ID
func (id DeviceManagementWindowsQualityUpdateProfileIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("windowsQualityUpdateProfiles", "windowsQualityUpdateProfiles", "windowsQualityUpdateProfiles"),
		resourceids.UserSpecifiedSegment("windowsQualityUpdateProfileId", "windowsQualityUpdateProfileId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("windowsQualityUpdateProfileAssignmentId", "windowsQualityUpdateProfileAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Windows Quality Update Profile Id Assignment ID
func (id DeviceManagementWindowsQualityUpdateProfileIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Windows Quality Update Profile: %q", id.WindowsQualityUpdateProfileId),
		fmt.Sprintf("Windows Quality Update Profile Assignment: %q", id.WindowsQualityUpdateProfileAssignmentId),
	}
	return fmt.Sprintf("Device Management Windows Quality Update Profile Id Assignment (%s)", strings.Join(components, "\n"))
}
