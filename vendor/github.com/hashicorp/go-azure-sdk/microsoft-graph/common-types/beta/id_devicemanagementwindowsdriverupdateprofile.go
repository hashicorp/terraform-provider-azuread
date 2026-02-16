package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsDriverUpdateProfileId{}

// DeviceManagementWindowsDriverUpdateProfileId is a struct representing the Resource ID for a Device Management Windows Driver Update Profile
type DeviceManagementWindowsDriverUpdateProfileId struct {
	WindowsDriverUpdateProfileId string
}

// NewDeviceManagementWindowsDriverUpdateProfileID returns a new DeviceManagementWindowsDriverUpdateProfileId struct
func NewDeviceManagementWindowsDriverUpdateProfileID(windowsDriverUpdateProfileId string) DeviceManagementWindowsDriverUpdateProfileId {
	return DeviceManagementWindowsDriverUpdateProfileId{
		WindowsDriverUpdateProfileId: windowsDriverUpdateProfileId,
	}
}

// ParseDeviceManagementWindowsDriverUpdateProfileID parses 'input' into a DeviceManagementWindowsDriverUpdateProfileId
func ParseDeviceManagementWindowsDriverUpdateProfileID(input string) (*DeviceManagementWindowsDriverUpdateProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsDriverUpdateProfileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsDriverUpdateProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementWindowsDriverUpdateProfileIDInsensitively parses 'input' case-insensitively into a DeviceManagementWindowsDriverUpdateProfileId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementWindowsDriverUpdateProfileIDInsensitively(input string) (*DeviceManagementWindowsDriverUpdateProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsDriverUpdateProfileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsDriverUpdateProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementWindowsDriverUpdateProfileId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WindowsDriverUpdateProfileId, ok = input.Parsed["windowsDriverUpdateProfileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsDriverUpdateProfileId", input)
	}

	return nil
}

// ValidateDeviceManagementWindowsDriverUpdateProfileID checks that 'input' can be parsed as a Device Management Windows Driver Update Profile ID
func ValidateDeviceManagementWindowsDriverUpdateProfileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementWindowsDriverUpdateProfileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Windows Driver Update Profile ID
func (id DeviceManagementWindowsDriverUpdateProfileId) ID() string {
	fmtString := "/deviceManagement/windowsDriverUpdateProfiles/%s"
	return fmt.Sprintf(fmtString, id.WindowsDriverUpdateProfileId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Windows Driver Update Profile ID
func (id DeviceManagementWindowsDriverUpdateProfileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("windowsDriverUpdateProfiles", "windowsDriverUpdateProfiles", "windowsDriverUpdateProfiles"),
		resourceids.UserSpecifiedSegment("windowsDriverUpdateProfileId", "windowsDriverUpdateProfileId"),
	}
}

// String returns a human-readable description of this Device Management Windows Driver Update Profile ID
func (id DeviceManagementWindowsDriverUpdateProfileId) String() string {
	components := []string{
		fmt.Sprintf("Windows Driver Update Profile: %q", id.WindowsDriverUpdateProfileId),
	}
	return fmt.Sprintf("Device Management Windows Driver Update Profile (%s)", strings.Join(components, "\n"))
}
