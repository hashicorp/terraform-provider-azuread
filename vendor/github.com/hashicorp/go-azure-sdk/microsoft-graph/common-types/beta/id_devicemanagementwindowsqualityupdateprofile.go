package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsQualityUpdateProfileId{}

// DeviceManagementWindowsQualityUpdateProfileId is a struct representing the Resource ID for a Device Management Windows Quality Update Profile
type DeviceManagementWindowsQualityUpdateProfileId struct {
	WindowsQualityUpdateProfileId string
}

// NewDeviceManagementWindowsQualityUpdateProfileID returns a new DeviceManagementWindowsQualityUpdateProfileId struct
func NewDeviceManagementWindowsQualityUpdateProfileID(windowsQualityUpdateProfileId string) DeviceManagementWindowsQualityUpdateProfileId {
	return DeviceManagementWindowsQualityUpdateProfileId{
		WindowsQualityUpdateProfileId: windowsQualityUpdateProfileId,
	}
}

// ParseDeviceManagementWindowsQualityUpdateProfileID parses 'input' into a DeviceManagementWindowsQualityUpdateProfileId
func ParseDeviceManagementWindowsQualityUpdateProfileID(input string) (*DeviceManagementWindowsQualityUpdateProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsQualityUpdateProfileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsQualityUpdateProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementWindowsQualityUpdateProfileIDInsensitively parses 'input' case-insensitively into a DeviceManagementWindowsQualityUpdateProfileId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementWindowsQualityUpdateProfileIDInsensitively(input string) (*DeviceManagementWindowsQualityUpdateProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsQualityUpdateProfileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsQualityUpdateProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementWindowsQualityUpdateProfileId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WindowsQualityUpdateProfileId, ok = input.Parsed["windowsQualityUpdateProfileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsQualityUpdateProfileId", input)
	}

	return nil
}

// ValidateDeviceManagementWindowsQualityUpdateProfileID checks that 'input' can be parsed as a Device Management Windows Quality Update Profile ID
func ValidateDeviceManagementWindowsQualityUpdateProfileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementWindowsQualityUpdateProfileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Windows Quality Update Profile ID
func (id DeviceManagementWindowsQualityUpdateProfileId) ID() string {
	fmtString := "/deviceManagement/windowsQualityUpdateProfiles/%s"
	return fmt.Sprintf(fmtString, id.WindowsQualityUpdateProfileId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Windows Quality Update Profile ID
func (id DeviceManagementWindowsQualityUpdateProfileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("windowsQualityUpdateProfiles", "windowsQualityUpdateProfiles", "windowsQualityUpdateProfiles"),
		resourceids.UserSpecifiedSegment("windowsQualityUpdateProfileId", "windowsQualityUpdateProfileId"),
	}
}

// String returns a human-readable description of this Device Management Windows Quality Update Profile ID
func (id DeviceManagementWindowsQualityUpdateProfileId) String() string {
	components := []string{
		fmt.Sprintf("Windows Quality Update Profile: %q", id.WindowsQualityUpdateProfileId),
	}
	return fmt.Sprintf("Device Management Windows Quality Update Profile (%s)", strings.Join(components, "\n"))
}
