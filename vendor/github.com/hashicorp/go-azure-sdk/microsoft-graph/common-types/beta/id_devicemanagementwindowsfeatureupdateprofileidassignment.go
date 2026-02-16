package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId{}

// DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId is a struct representing the Resource ID for a Device Management Windows Feature Update Profile Id Assignment
type DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId struct {
	WindowsFeatureUpdateProfileId           string
	WindowsFeatureUpdateProfileAssignmentId string
}

// NewDeviceManagementWindowsFeatureUpdateProfileIdAssignmentID returns a new DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId struct
func NewDeviceManagementWindowsFeatureUpdateProfileIdAssignmentID(windowsFeatureUpdateProfileId string, windowsFeatureUpdateProfileAssignmentId string) DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId {
	return DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId{
		WindowsFeatureUpdateProfileId:           windowsFeatureUpdateProfileId,
		WindowsFeatureUpdateProfileAssignmentId: windowsFeatureUpdateProfileAssignmentId,
	}
}

// ParseDeviceManagementWindowsFeatureUpdateProfileIdAssignmentID parses 'input' into a DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId
func ParseDeviceManagementWindowsFeatureUpdateProfileIdAssignmentID(input string) (*DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementWindowsFeatureUpdateProfileIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementWindowsFeatureUpdateProfileIdAssignmentIDInsensitively(input string) (*DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WindowsFeatureUpdateProfileId, ok = input.Parsed["windowsFeatureUpdateProfileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsFeatureUpdateProfileId", input)
	}

	if id.WindowsFeatureUpdateProfileAssignmentId, ok = input.Parsed["windowsFeatureUpdateProfileAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsFeatureUpdateProfileAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementWindowsFeatureUpdateProfileIdAssignmentID checks that 'input' can be parsed as a Device Management Windows Feature Update Profile Id Assignment ID
func ValidateDeviceManagementWindowsFeatureUpdateProfileIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementWindowsFeatureUpdateProfileIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Windows Feature Update Profile Id Assignment ID
func (id DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/windowsFeatureUpdateProfiles/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.WindowsFeatureUpdateProfileId, id.WindowsFeatureUpdateProfileAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Windows Feature Update Profile Id Assignment ID
func (id DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("windowsFeatureUpdateProfiles", "windowsFeatureUpdateProfiles", "windowsFeatureUpdateProfiles"),
		resourceids.UserSpecifiedSegment("windowsFeatureUpdateProfileId", "windowsFeatureUpdateProfileId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("windowsFeatureUpdateProfileAssignmentId", "windowsFeatureUpdateProfileAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Windows Feature Update Profile Id Assignment ID
func (id DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Windows Feature Update Profile: %q", id.WindowsFeatureUpdateProfileId),
		fmt.Sprintf("Windows Feature Update Profile Assignment: %q", id.WindowsFeatureUpdateProfileAssignmentId),
	}
	return fmt.Sprintf("Device Management Windows Feature Update Profile Id Assignment (%s)", strings.Join(components, "\n"))
}
