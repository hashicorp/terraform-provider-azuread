package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId{}

// DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId is a struct representing the Resource ID for a Device Management Windows Quality Update Policy Id Assignment
type DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId struct {
	WindowsQualityUpdatePolicyId           string
	WindowsQualityUpdatePolicyAssignmentId string
}

// NewDeviceManagementWindowsQualityUpdatePolicyIdAssignmentID returns a new DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId struct
func NewDeviceManagementWindowsQualityUpdatePolicyIdAssignmentID(windowsQualityUpdatePolicyId string, windowsQualityUpdatePolicyAssignmentId string) DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId {
	return DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId{
		WindowsQualityUpdatePolicyId:           windowsQualityUpdatePolicyId,
		WindowsQualityUpdatePolicyAssignmentId: windowsQualityUpdatePolicyAssignmentId,
	}
}

// ParseDeviceManagementWindowsQualityUpdatePolicyIdAssignmentID parses 'input' into a DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId
func ParseDeviceManagementWindowsQualityUpdatePolicyIdAssignmentID(input string) (*DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementWindowsQualityUpdatePolicyIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementWindowsQualityUpdatePolicyIdAssignmentIDInsensitively(input string) (*DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WindowsQualityUpdatePolicyId, ok = input.Parsed["windowsQualityUpdatePolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsQualityUpdatePolicyId", input)
	}

	if id.WindowsQualityUpdatePolicyAssignmentId, ok = input.Parsed["windowsQualityUpdatePolicyAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsQualityUpdatePolicyAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementWindowsQualityUpdatePolicyIdAssignmentID checks that 'input' can be parsed as a Device Management Windows Quality Update Policy Id Assignment ID
func ValidateDeviceManagementWindowsQualityUpdatePolicyIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementWindowsQualityUpdatePolicyIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Windows Quality Update Policy Id Assignment ID
func (id DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/windowsQualityUpdatePolicies/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.WindowsQualityUpdatePolicyId, id.WindowsQualityUpdatePolicyAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Windows Quality Update Policy Id Assignment ID
func (id DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("windowsQualityUpdatePolicies", "windowsQualityUpdatePolicies", "windowsQualityUpdatePolicies"),
		resourceids.UserSpecifiedSegment("windowsQualityUpdatePolicyId", "windowsQualityUpdatePolicyId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("windowsQualityUpdatePolicyAssignmentId", "windowsQualityUpdatePolicyAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Windows Quality Update Policy Id Assignment ID
func (id DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Windows Quality Update Policy: %q", id.WindowsQualityUpdatePolicyId),
		fmt.Sprintf("Windows Quality Update Policy Assignment: %q", id.WindowsQualityUpdatePolicyAssignmentId),
	}
	return fmt.Sprintf("Device Management Windows Quality Update Policy Id Assignment (%s)", strings.Join(components, "\n"))
}
