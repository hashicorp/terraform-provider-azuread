package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceShellScriptIdAssignmentId{}

// DeviceManagementDeviceShellScriptIdAssignmentId is a struct representing the Resource ID for a Device Management Device Shell Script Id Assignment
type DeviceManagementDeviceShellScriptIdAssignmentId struct {
	DeviceShellScriptId                string
	DeviceManagementScriptAssignmentId string
}

// NewDeviceManagementDeviceShellScriptIdAssignmentID returns a new DeviceManagementDeviceShellScriptIdAssignmentId struct
func NewDeviceManagementDeviceShellScriptIdAssignmentID(deviceShellScriptId string, deviceManagementScriptAssignmentId string) DeviceManagementDeviceShellScriptIdAssignmentId {
	return DeviceManagementDeviceShellScriptIdAssignmentId{
		DeviceShellScriptId:                deviceShellScriptId,
		DeviceManagementScriptAssignmentId: deviceManagementScriptAssignmentId,
	}
}

// ParseDeviceManagementDeviceShellScriptIdAssignmentID parses 'input' into a DeviceManagementDeviceShellScriptIdAssignmentId
func ParseDeviceManagementDeviceShellScriptIdAssignmentID(input string) (*DeviceManagementDeviceShellScriptIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceShellScriptIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceShellScriptIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceShellScriptIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceShellScriptIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceShellScriptIdAssignmentIDInsensitively(input string) (*DeviceManagementDeviceShellScriptIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceShellScriptIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceShellScriptIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceShellScriptIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceShellScriptId, ok = input.Parsed["deviceShellScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceShellScriptId", input)
	}

	if id.DeviceManagementScriptAssignmentId, ok = input.Parsed["deviceManagementScriptAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceShellScriptIdAssignmentID checks that 'input' can be parsed as a Device Management Device Shell Script Id Assignment ID
func ValidateDeviceManagementDeviceShellScriptIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceShellScriptIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Shell Script Id Assignment ID
func (id DeviceManagementDeviceShellScriptIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/deviceShellScripts/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.DeviceShellScriptId, id.DeviceManagementScriptAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Shell Script Id Assignment ID
func (id DeviceManagementDeviceShellScriptIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceShellScripts", "deviceShellScripts", "deviceShellScripts"),
		resourceids.UserSpecifiedSegment("deviceShellScriptId", "deviceShellScriptId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptAssignmentId", "deviceManagementScriptAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Device Shell Script Id Assignment ID
func (id DeviceManagementDeviceShellScriptIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Device Shell Script: %q", id.DeviceShellScriptId),
		fmt.Sprintf("Device Management Script Assignment: %q", id.DeviceManagementScriptAssignmentId),
	}
	return fmt.Sprintf("Device Management Device Shell Script Id Assignment (%s)", strings.Join(components, "\n"))
}
