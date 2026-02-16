package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId{}

// DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId is a struct representing the Resource ID for a Device Management Device Custom Attribute Shell Script Id Device Run State
type DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId struct {
	DeviceCustomAttributeShellScriptId  string
	DeviceManagementScriptDeviceStateId string
}

// NewDeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateID returns a new DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId struct
func NewDeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateID(deviceCustomAttributeShellScriptId string, deviceManagementScriptDeviceStateId string) DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId {
	return DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId{
		DeviceCustomAttributeShellScriptId:  deviceCustomAttributeShellScriptId,
		DeviceManagementScriptDeviceStateId: deviceManagementScriptDeviceStateId,
	}
}

// ParseDeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateID parses 'input' into a DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId
func ParseDeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateID(input string) (*DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateIDInsensitively(input string) (*DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceCustomAttributeShellScriptId, ok = input.Parsed["deviceCustomAttributeShellScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceCustomAttributeShellScriptId", input)
	}

	if id.DeviceManagementScriptDeviceStateId, ok = input.Parsed["deviceManagementScriptDeviceStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptDeviceStateId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateID checks that 'input' can be parsed as a Device Management Device Custom Attribute Shell Script Id Device Run State ID
func ValidateDeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Custom Attribute Shell Script Id Device Run State ID
func (id DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId) ID() string {
	fmtString := "/deviceManagement/deviceCustomAttributeShellScripts/%s/deviceRunStates/%s"
	return fmt.Sprintf(fmtString, id.DeviceCustomAttributeShellScriptId, id.DeviceManagementScriptDeviceStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Custom Attribute Shell Script Id Device Run State ID
func (id DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceCustomAttributeShellScripts", "deviceCustomAttributeShellScripts", "deviceCustomAttributeShellScripts"),
		resourceids.UserSpecifiedSegment("deviceCustomAttributeShellScriptId", "deviceCustomAttributeShellScriptId"),
		resourceids.StaticSegment("deviceRunStates", "deviceRunStates", "deviceRunStates"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptDeviceStateId", "deviceManagementScriptDeviceStateId"),
	}
}

// String returns a human-readable description of this Device Management Device Custom Attribute Shell Script Id Device Run State ID
func (id DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId) String() string {
	components := []string{
		fmt.Sprintf("Device Custom Attribute Shell Script: %q", id.DeviceCustomAttributeShellScriptId),
		fmt.Sprintf("Device Management Script Device State: %q", id.DeviceManagementScriptDeviceStateId),
	}
	return fmt.Sprintf("Device Management Device Custom Attribute Shell Script Id Device Run State (%s)", strings.Join(components, "\n"))
}
