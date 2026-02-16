package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId{}

// DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId is a struct representing the Resource ID for a Device Management Device Custom Attribute Shell Script Id User Run State Id Device Run State
type DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId struct {
	DeviceCustomAttributeShellScriptId  string
	DeviceManagementScriptUserStateId   string
	DeviceManagementScriptDeviceStateId string
}

// NewDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateID returns a new DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId struct
func NewDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateID(deviceCustomAttributeShellScriptId string, deviceManagementScriptUserStateId string, deviceManagementScriptDeviceStateId string) DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId {
	return DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId{
		DeviceCustomAttributeShellScriptId:  deviceCustomAttributeShellScriptId,
		DeviceManagementScriptUserStateId:   deviceManagementScriptUserStateId,
		DeviceManagementScriptDeviceStateId: deviceManagementScriptDeviceStateId,
	}
}

// ParseDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateID parses 'input' into a DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId
func ParseDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateID(input string) (*DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateIDInsensitively(input string) (*DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceCustomAttributeShellScriptId, ok = input.Parsed["deviceCustomAttributeShellScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceCustomAttributeShellScriptId", input)
	}

	if id.DeviceManagementScriptUserStateId, ok = input.Parsed["deviceManagementScriptUserStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptUserStateId", input)
	}

	if id.DeviceManagementScriptDeviceStateId, ok = input.Parsed["deviceManagementScriptDeviceStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptDeviceStateId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateID checks that 'input' can be parsed as a Device Management Device Custom Attribute Shell Script Id User Run State Id Device Run State ID
func ValidateDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Custom Attribute Shell Script Id User Run State Id Device Run State ID
func (id DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId) ID() string {
	fmtString := "/deviceManagement/deviceCustomAttributeShellScripts/%s/userRunStates/%s/deviceRunStates/%s"
	return fmt.Sprintf(fmtString, id.DeviceCustomAttributeShellScriptId, id.DeviceManagementScriptUserStateId, id.DeviceManagementScriptDeviceStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Custom Attribute Shell Script Id User Run State Id Device Run State ID
func (id DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceCustomAttributeShellScripts", "deviceCustomAttributeShellScripts", "deviceCustomAttributeShellScripts"),
		resourceids.UserSpecifiedSegment("deviceCustomAttributeShellScriptId", "deviceCustomAttributeShellScriptId"),
		resourceids.StaticSegment("userRunStates", "userRunStates", "userRunStates"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptUserStateId", "deviceManagementScriptUserStateId"),
		resourceids.StaticSegment("deviceRunStates", "deviceRunStates", "deviceRunStates"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptDeviceStateId", "deviceManagementScriptDeviceStateId"),
	}
}

// String returns a human-readable description of this Device Management Device Custom Attribute Shell Script Id User Run State Id Device Run State ID
func (id DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId) String() string {
	components := []string{
		fmt.Sprintf("Device Custom Attribute Shell Script: %q", id.DeviceCustomAttributeShellScriptId),
		fmt.Sprintf("Device Management Script User State: %q", id.DeviceManagementScriptUserStateId),
		fmt.Sprintf("Device Management Script Device State: %q", id.DeviceManagementScriptDeviceStateId),
	}
	return fmt.Sprintf("Device Management Device Custom Attribute Shell Script Id User Run State Id Device Run State (%s)", strings.Join(components, "\n"))
}
