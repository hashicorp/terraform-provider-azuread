package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsFeatureUpdateProfileId{}

// DeviceManagementWindowsFeatureUpdateProfileId is a struct representing the Resource ID for a Device Management Windows Feature Update Profile
type DeviceManagementWindowsFeatureUpdateProfileId struct {
	WindowsFeatureUpdateProfileId string
}

// NewDeviceManagementWindowsFeatureUpdateProfileID returns a new DeviceManagementWindowsFeatureUpdateProfileId struct
func NewDeviceManagementWindowsFeatureUpdateProfileID(windowsFeatureUpdateProfileId string) DeviceManagementWindowsFeatureUpdateProfileId {
	return DeviceManagementWindowsFeatureUpdateProfileId{
		WindowsFeatureUpdateProfileId: windowsFeatureUpdateProfileId,
	}
}

// ParseDeviceManagementWindowsFeatureUpdateProfileID parses 'input' into a DeviceManagementWindowsFeatureUpdateProfileId
func ParseDeviceManagementWindowsFeatureUpdateProfileID(input string) (*DeviceManagementWindowsFeatureUpdateProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsFeatureUpdateProfileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsFeatureUpdateProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementWindowsFeatureUpdateProfileIDInsensitively parses 'input' case-insensitively into a DeviceManagementWindowsFeatureUpdateProfileId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementWindowsFeatureUpdateProfileIDInsensitively(input string) (*DeviceManagementWindowsFeatureUpdateProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsFeatureUpdateProfileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsFeatureUpdateProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementWindowsFeatureUpdateProfileId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WindowsFeatureUpdateProfileId, ok = input.Parsed["windowsFeatureUpdateProfileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsFeatureUpdateProfileId", input)
	}

	return nil
}

// ValidateDeviceManagementWindowsFeatureUpdateProfileID checks that 'input' can be parsed as a Device Management Windows Feature Update Profile ID
func ValidateDeviceManagementWindowsFeatureUpdateProfileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementWindowsFeatureUpdateProfileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Windows Feature Update Profile ID
func (id DeviceManagementWindowsFeatureUpdateProfileId) ID() string {
	fmtString := "/deviceManagement/windowsFeatureUpdateProfiles/%s"
	return fmt.Sprintf(fmtString, id.WindowsFeatureUpdateProfileId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Windows Feature Update Profile ID
func (id DeviceManagementWindowsFeatureUpdateProfileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("windowsFeatureUpdateProfiles", "windowsFeatureUpdateProfiles", "windowsFeatureUpdateProfiles"),
		resourceids.UserSpecifiedSegment("windowsFeatureUpdateProfileId", "windowsFeatureUpdateProfileId"),
	}
}

// String returns a human-readable description of this Device Management Windows Feature Update Profile ID
func (id DeviceManagementWindowsFeatureUpdateProfileId) String() string {
	components := []string{
		fmt.Sprintf("Windows Feature Update Profile: %q", id.WindowsFeatureUpdateProfileId),
	}
	return fmt.Sprintf("Device Management Windows Feature Update Profile (%s)", strings.Join(components, "\n"))
}
