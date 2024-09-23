package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsAutopilotDeviceIdentityId{}

// DeviceManagementWindowsAutopilotDeviceIdentityId is a struct representing the Resource ID for a Device Management Windows Autopilot Device Identity
type DeviceManagementWindowsAutopilotDeviceIdentityId struct {
	WindowsAutopilotDeviceIdentityId string
}

// NewDeviceManagementWindowsAutopilotDeviceIdentityID returns a new DeviceManagementWindowsAutopilotDeviceIdentityId struct
func NewDeviceManagementWindowsAutopilotDeviceIdentityID(windowsAutopilotDeviceIdentityId string) DeviceManagementWindowsAutopilotDeviceIdentityId {
	return DeviceManagementWindowsAutopilotDeviceIdentityId{
		WindowsAutopilotDeviceIdentityId: windowsAutopilotDeviceIdentityId,
	}
}

// ParseDeviceManagementWindowsAutopilotDeviceIdentityID parses 'input' into a DeviceManagementWindowsAutopilotDeviceIdentityId
func ParseDeviceManagementWindowsAutopilotDeviceIdentityID(input string) (*DeviceManagementWindowsAutopilotDeviceIdentityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsAutopilotDeviceIdentityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsAutopilotDeviceIdentityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementWindowsAutopilotDeviceIdentityIDInsensitively parses 'input' case-insensitively into a DeviceManagementWindowsAutopilotDeviceIdentityId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementWindowsAutopilotDeviceIdentityIDInsensitively(input string) (*DeviceManagementWindowsAutopilotDeviceIdentityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsAutopilotDeviceIdentityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsAutopilotDeviceIdentityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementWindowsAutopilotDeviceIdentityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WindowsAutopilotDeviceIdentityId, ok = input.Parsed["windowsAutopilotDeviceIdentityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsAutopilotDeviceIdentityId", input)
	}

	return nil
}

// ValidateDeviceManagementWindowsAutopilotDeviceIdentityID checks that 'input' can be parsed as a Device Management Windows Autopilot Device Identity ID
func ValidateDeviceManagementWindowsAutopilotDeviceIdentityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementWindowsAutopilotDeviceIdentityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Windows Autopilot Device Identity ID
func (id DeviceManagementWindowsAutopilotDeviceIdentityId) ID() string {
	fmtString := "/deviceManagement/windowsAutopilotDeviceIdentities/%s"
	return fmt.Sprintf(fmtString, id.WindowsAutopilotDeviceIdentityId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Windows Autopilot Device Identity ID
func (id DeviceManagementWindowsAutopilotDeviceIdentityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("windowsAutopilotDeviceIdentities", "windowsAutopilotDeviceIdentities", "windowsAutopilotDeviceIdentities"),
		resourceids.UserSpecifiedSegment("windowsAutopilotDeviceIdentityId", "windowsAutopilotDeviceIdentityId"),
	}
}

// String returns a human-readable description of this Device Management Windows Autopilot Device Identity ID
func (id DeviceManagementWindowsAutopilotDeviceIdentityId) String() string {
	components := []string{
		fmt.Sprintf("Windows Autopilot Device Identity: %q", id.WindowsAutopilotDeviceIdentityId),
	}
	return fmt.Sprintf("Device Management Windows Autopilot Device Identity (%s)", strings.Join(components, "\n"))
}
