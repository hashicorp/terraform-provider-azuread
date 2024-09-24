package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceShellScriptIdDeviceRunStateId{}

// DeviceManagementDeviceShellScriptIdDeviceRunStateId is a struct representing the Resource ID for a Device Management Device Shell Script Id Device Run State
type DeviceManagementDeviceShellScriptIdDeviceRunStateId struct {
	DeviceShellScriptId                 string
	DeviceManagementScriptDeviceStateId string
}

// NewDeviceManagementDeviceShellScriptIdDeviceRunStateID returns a new DeviceManagementDeviceShellScriptIdDeviceRunStateId struct
func NewDeviceManagementDeviceShellScriptIdDeviceRunStateID(deviceShellScriptId string, deviceManagementScriptDeviceStateId string) DeviceManagementDeviceShellScriptIdDeviceRunStateId {
	return DeviceManagementDeviceShellScriptIdDeviceRunStateId{
		DeviceShellScriptId:                 deviceShellScriptId,
		DeviceManagementScriptDeviceStateId: deviceManagementScriptDeviceStateId,
	}
}

// ParseDeviceManagementDeviceShellScriptIdDeviceRunStateID parses 'input' into a DeviceManagementDeviceShellScriptIdDeviceRunStateId
func ParseDeviceManagementDeviceShellScriptIdDeviceRunStateID(input string) (*DeviceManagementDeviceShellScriptIdDeviceRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceShellScriptIdDeviceRunStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceShellScriptIdDeviceRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceShellScriptIdDeviceRunStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceShellScriptIdDeviceRunStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceShellScriptIdDeviceRunStateIDInsensitively(input string) (*DeviceManagementDeviceShellScriptIdDeviceRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceShellScriptIdDeviceRunStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceShellScriptIdDeviceRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceShellScriptIdDeviceRunStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceShellScriptId, ok = input.Parsed["deviceShellScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceShellScriptId", input)
	}

	if id.DeviceManagementScriptDeviceStateId, ok = input.Parsed["deviceManagementScriptDeviceStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptDeviceStateId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceShellScriptIdDeviceRunStateID checks that 'input' can be parsed as a Device Management Device Shell Script Id Device Run State ID
func ValidateDeviceManagementDeviceShellScriptIdDeviceRunStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceShellScriptIdDeviceRunStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Shell Script Id Device Run State ID
func (id DeviceManagementDeviceShellScriptIdDeviceRunStateId) ID() string {
	fmtString := "/deviceManagement/deviceShellScripts/%s/deviceRunStates/%s"
	return fmt.Sprintf(fmtString, id.DeviceShellScriptId, id.DeviceManagementScriptDeviceStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Shell Script Id Device Run State ID
func (id DeviceManagementDeviceShellScriptIdDeviceRunStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceShellScripts", "deviceShellScripts", "deviceShellScripts"),
		resourceids.UserSpecifiedSegment("deviceShellScriptId", "deviceShellScriptId"),
		resourceids.StaticSegment("deviceRunStates", "deviceRunStates", "deviceRunStates"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptDeviceStateId", "deviceManagementScriptDeviceStateId"),
	}
}

// String returns a human-readable description of this Device Management Device Shell Script Id Device Run State ID
func (id DeviceManagementDeviceShellScriptIdDeviceRunStateId) String() string {
	components := []string{
		fmt.Sprintf("Device Shell Script: %q", id.DeviceShellScriptId),
		fmt.Sprintf("Device Management Script Device State: %q", id.DeviceManagementScriptDeviceStateId),
	}
	return fmt.Sprintf("Device Management Device Shell Script Id Device Run State (%s)", strings.Join(components, "\n"))
}
