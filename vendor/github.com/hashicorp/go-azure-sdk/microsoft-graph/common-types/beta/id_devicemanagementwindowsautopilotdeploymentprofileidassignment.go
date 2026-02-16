package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId{}

// DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId is a struct representing the Resource ID for a Device Management Windows Autopilot Deployment Profile Id Assignment
type DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId struct {
	WindowsAutopilotDeploymentProfileId           string
	WindowsAutopilotDeploymentProfileAssignmentId string
}

// NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentID returns a new DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId struct
func NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentID(windowsAutopilotDeploymentProfileId string, windowsAutopilotDeploymentProfileAssignmentId string) DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId {
	return DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId{
		WindowsAutopilotDeploymentProfileId:           windowsAutopilotDeploymentProfileId,
		WindowsAutopilotDeploymentProfileAssignmentId: windowsAutopilotDeploymentProfileAssignmentId,
	}
}

// ParseDeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentID parses 'input' into a DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId
func ParseDeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentID(input string) (*DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentIDInsensitively(input string) (*DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WindowsAutopilotDeploymentProfileId, ok = input.Parsed["windowsAutopilotDeploymentProfileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsAutopilotDeploymentProfileId", input)
	}

	if id.WindowsAutopilotDeploymentProfileAssignmentId, ok = input.Parsed["windowsAutopilotDeploymentProfileAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsAutopilotDeploymentProfileAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentID checks that 'input' can be parsed as a Device Management Windows Autopilot Deployment Profile Id Assignment ID
func ValidateDeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Windows Autopilot Deployment Profile Id Assignment ID
func (id DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/windowsAutopilotDeploymentProfiles/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.WindowsAutopilotDeploymentProfileId, id.WindowsAutopilotDeploymentProfileAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Windows Autopilot Deployment Profile Id Assignment ID
func (id DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("windowsAutopilotDeploymentProfiles", "windowsAutopilotDeploymentProfiles", "windowsAutopilotDeploymentProfiles"),
		resourceids.UserSpecifiedSegment("windowsAutopilotDeploymentProfileId", "windowsAutopilotDeploymentProfileId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("windowsAutopilotDeploymentProfileAssignmentId", "windowsAutopilotDeploymentProfileAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Windows Autopilot Deployment Profile Id Assignment ID
func (id DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Windows Autopilot Deployment Profile: %q", id.WindowsAutopilotDeploymentProfileId),
		fmt.Sprintf("Windows Autopilot Deployment Profile Assignment: %q", id.WindowsAutopilotDeploymentProfileAssignmentId),
	}
	return fmt.Sprintf("Device Management Windows Autopilot Deployment Profile Id Assignment (%s)", strings.Join(components, "\n"))
}
