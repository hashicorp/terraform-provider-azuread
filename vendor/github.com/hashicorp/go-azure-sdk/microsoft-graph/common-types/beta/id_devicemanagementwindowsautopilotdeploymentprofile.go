package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsAutopilotDeploymentProfileId{}

// DeviceManagementWindowsAutopilotDeploymentProfileId is a struct representing the Resource ID for a Device Management Windows Autopilot Deployment Profile
type DeviceManagementWindowsAutopilotDeploymentProfileId struct {
	WindowsAutopilotDeploymentProfileId string
}

// NewDeviceManagementWindowsAutopilotDeploymentProfileID returns a new DeviceManagementWindowsAutopilotDeploymentProfileId struct
func NewDeviceManagementWindowsAutopilotDeploymentProfileID(windowsAutopilotDeploymentProfileId string) DeviceManagementWindowsAutopilotDeploymentProfileId {
	return DeviceManagementWindowsAutopilotDeploymentProfileId{
		WindowsAutopilotDeploymentProfileId: windowsAutopilotDeploymentProfileId,
	}
}

// ParseDeviceManagementWindowsAutopilotDeploymentProfileID parses 'input' into a DeviceManagementWindowsAutopilotDeploymentProfileId
func ParseDeviceManagementWindowsAutopilotDeploymentProfileID(input string) (*DeviceManagementWindowsAutopilotDeploymentProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsAutopilotDeploymentProfileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsAutopilotDeploymentProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementWindowsAutopilotDeploymentProfileIDInsensitively parses 'input' case-insensitively into a DeviceManagementWindowsAutopilotDeploymentProfileId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementWindowsAutopilotDeploymentProfileIDInsensitively(input string) (*DeviceManagementWindowsAutopilotDeploymentProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsAutopilotDeploymentProfileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsAutopilotDeploymentProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementWindowsAutopilotDeploymentProfileId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WindowsAutopilotDeploymentProfileId, ok = input.Parsed["windowsAutopilotDeploymentProfileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsAutopilotDeploymentProfileId", input)
	}

	return nil
}

// ValidateDeviceManagementWindowsAutopilotDeploymentProfileID checks that 'input' can be parsed as a Device Management Windows Autopilot Deployment Profile ID
func ValidateDeviceManagementWindowsAutopilotDeploymentProfileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementWindowsAutopilotDeploymentProfileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Windows Autopilot Deployment Profile ID
func (id DeviceManagementWindowsAutopilotDeploymentProfileId) ID() string {
	fmtString := "/deviceManagement/windowsAutopilotDeploymentProfiles/%s"
	return fmt.Sprintf(fmtString, id.WindowsAutopilotDeploymentProfileId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Windows Autopilot Deployment Profile ID
func (id DeviceManagementWindowsAutopilotDeploymentProfileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("windowsAutopilotDeploymentProfiles", "windowsAutopilotDeploymentProfiles", "windowsAutopilotDeploymentProfiles"),
		resourceids.UserSpecifiedSegment("windowsAutopilotDeploymentProfileId", "windowsAutopilotDeploymentProfileId"),
	}
}

// String returns a human-readable description of this Device Management Windows Autopilot Deployment Profile ID
func (id DeviceManagementWindowsAutopilotDeploymentProfileId) String() string {
	components := []string{
		fmt.Sprintf("Windows Autopilot Deployment Profile: %q", id.WindowsAutopilotDeploymentProfileId),
	}
	return fmt.Sprintf("Device Management Windows Autopilot Deployment Profile (%s)", strings.Join(components, "\n"))
}
