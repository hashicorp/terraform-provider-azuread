package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId{}

// DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId is a struct representing the Resource ID for a Device Management Windows Autopilot Deployment Profile Id Assigned Device
type DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId struct {
	WindowsAutopilotDeploymentProfileId string
	WindowsAutopilotDeviceIdentityId    string
}

// NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID returns a new DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId struct
func NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID(windowsAutopilotDeploymentProfileId string, windowsAutopilotDeviceIdentityId string) DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId {
	return DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId{
		WindowsAutopilotDeploymentProfileId: windowsAutopilotDeploymentProfileId,
		WindowsAutopilotDeviceIdentityId:    windowsAutopilotDeviceIdentityId,
	}
}

// ParseDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID parses 'input' into a DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId
func ParseDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID(input string) (*DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceIDInsensitively parses 'input' case-insensitively into a DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceIDInsensitively(input string) (*DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WindowsAutopilotDeploymentProfileId, ok = input.Parsed["windowsAutopilotDeploymentProfileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsAutopilotDeploymentProfileId", input)
	}

	if id.WindowsAutopilotDeviceIdentityId, ok = input.Parsed["windowsAutopilotDeviceIdentityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsAutopilotDeviceIdentityId", input)
	}

	return nil
}

// ValidateDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID checks that 'input' can be parsed as a Device Management Windows Autopilot Deployment Profile Id Assigned Device ID
func ValidateDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Windows Autopilot Deployment Profile Id Assigned Device ID
func (id DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId) ID() string {
	fmtString := "/deviceManagement/windowsAutopilotDeploymentProfiles/%s/assignedDevices/%s"
	return fmt.Sprintf(fmtString, id.WindowsAutopilotDeploymentProfileId, id.WindowsAutopilotDeviceIdentityId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Windows Autopilot Deployment Profile Id Assigned Device ID
func (id DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("windowsAutopilotDeploymentProfiles", "windowsAutopilotDeploymentProfiles", "windowsAutopilotDeploymentProfiles"),
		resourceids.UserSpecifiedSegment("windowsAutopilotDeploymentProfileId", "windowsAutopilotDeploymentProfileId"),
		resourceids.StaticSegment("assignedDevices", "assignedDevices", "assignedDevices"),
		resourceids.UserSpecifiedSegment("windowsAutopilotDeviceIdentityId", "windowsAutopilotDeviceIdentityId"),
	}
}

// String returns a human-readable description of this Device Management Windows Autopilot Deployment Profile Id Assigned Device ID
func (id DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId) String() string {
	components := []string{
		fmt.Sprintf("Windows Autopilot Deployment Profile: %q", id.WindowsAutopilotDeploymentProfileId),
		fmt.Sprintf("Windows Autopilot Device Identity: %q", id.WindowsAutopilotDeviceIdentityId),
	}
	return fmt.Sprintf("Device Management Windows Autopilot Deployment Profile Id Assigned Device (%s)", strings.Join(components, "\n"))
}
