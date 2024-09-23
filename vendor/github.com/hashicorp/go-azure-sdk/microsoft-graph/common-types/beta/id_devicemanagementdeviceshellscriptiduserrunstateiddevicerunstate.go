package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateId{}

// DeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateId is a struct representing the Resource ID for a Device Management Device Shell Script Id User Run State Id Device Run State
type DeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateId struct {
	DeviceShellScriptId                 string
	DeviceManagementScriptUserStateId   string
	DeviceManagementScriptDeviceStateId string
}

// NewDeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateID returns a new DeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateId struct
func NewDeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateID(deviceShellScriptId string, deviceManagementScriptUserStateId string, deviceManagementScriptDeviceStateId string) DeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateId {
	return DeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateId{
		DeviceShellScriptId:                 deviceShellScriptId,
		DeviceManagementScriptUserStateId:   deviceManagementScriptUserStateId,
		DeviceManagementScriptDeviceStateId: deviceManagementScriptDeviceStateId,
	}
}

// ParseDeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateID parses 'input' into a DeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateId
func ParseDeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateID(input string) (*DeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateIDInsensitively(input string) (*DeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceShellScriptId, ok = input.Parsed["deviceShellScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceShellScriptId", input)
	}

	if id.DeviceManagementScriptUserStateId, ok = input.Parsed["deviceManagementScriptUserStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptUserStateId", input)
	}

	if id.DeviceManagementScriptDeviceStateId, ok = input.Parsed["deviceManagementScriptDeviceStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptDeviceStateId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateID checks that 'input' can be parsed as a Device Management Device Shell Script Id User Run State Id Device Run State ID
func ValidateDeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Shell Script Id User Run State Id Device Run State ID
func (id DeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateId) ID() string {
	fmtString := "/deviceManagement/deviceShellScripts/%s/userRunStates/%s/deviceRunStates/%s"
	return fmt.Sprintf(fmtString, id.DeviceShellScriptId, id.DeviceManagementScriptUserStateId, id.DeviceManagementScriptDeviceStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Shell Script Id User Run State Id Device Run State ID
func (id DeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceShellScripts", "deviceShellScripts", "deviceShellScripts"),
		resourceids.UserSpecifiedSegment("deviceShellScriptId", "deviceShellScriptId"),
		resourceids.StaticSegment("userRunStates", "userRunStates", "userRunStates"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptUserStateId", "deviceManagementScriptUserStateId"),
		resourceids.StaticSegment("deviceRunStates", "deviceRunStates", "deviceRunStates"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptDeviceStateId", "deviceManagementScriptDeviceStateId"),
	}
}

// String returns a human-readable description of this Device Management Device Shell Script Id User Run State Id Device Run State ID
func (id DeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateId) String() string {
	components := []string{
		fmt.Sprintf("Device Shell Script: %q", id.DeviceShellScriptId),
		fmt.Sprintf("Device Management Script User State: %q", id.DeviceManagementScriptUserStateId),
		fmt.Sprintf("Device Management Script Device State: %q", id.DeviceManagementScriptDeviceStateId),
	}
	return fmt.Sprintf("Device Management Device Shell Script Id User Run State Id Device Run State (%s)", strings.Join(components, "\n"))
}
