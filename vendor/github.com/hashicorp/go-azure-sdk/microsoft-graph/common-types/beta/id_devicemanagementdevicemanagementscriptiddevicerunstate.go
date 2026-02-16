package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceManagementScriptIdDeviceRunStateId{}

// DeviceManagementDeviceManagementScriptIdDeviceRunStateId is a struct representing the Resource ID for a Device Management Device Management Script Id Device Run State
type DeviceManagementDeviceManagementScriptIdDeviceRunStateId struct {
	DeviceManagementScriptId            string
	DeviceManagementScriptDeviceStateId string
}

// NewDeviceManagementDeviceManagementScriptIdDeviceRunStateID returns a new DeviceManagementDeviceManagementScriptIdDeviceRunStateId struct
func NewDeviceManagementDeviceManagementScriptIdDeviceRunStateID(deviceManagementScriptId string, deviceManagementScriptDeviceStateId string) DeviceManagementDeviceManagementScriptIdDeviceRunStateId {
	return DeviceManagementDeviceManagementScriptIdDeviceRunStateId{
		DeviceManagementScriptId:            deviceManagementScriptId,
		DeviceManagementScriptDeviceStateId: deviceManagementScriptDeviceStateId,
	}
}

// ParseDeviceManagementDeviceManagementScriptIdDeviceRunStateID parses 'input' into a DeviceManagementDeviceManagementScriptIdDeviceRunStateId
func ParseDeviceManagementDeviceManagementScriptIdDeviceRunStateID(input string) (*DeviceManagementDeviceManagementScriptIdDeviceRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceManagementScriptIdDeviceRunStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceManagementScriptIdDeviceRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceManagementScriptIdDeviceRunStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceManagementScriptIdDeviceRunStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceManagementScriptIdDeviceRunStateIDInsensitively(input string) (*DeviceManagementDeviceManagementScriptIdDeviceRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceManagementScriptIdDeviceRunStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceManagementScriptIdDeviceRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceManagementScriptIdDeviceRunStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementScriptId, ok = input.Parsed["deviceManagementScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptId", input)
	}

	if id.DeviceManagementScriptDeviceStateId, ok = input.Parsed["deviceManagementScriptDeviceStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptDeviceStateId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceManagementScriptIdDeviceRunStateID checks that 'input' can be parsed as a Device Management Device Management Script Id Device Run State ID
func ValidateDeviceManagementDeviceManagementScriptIdDeviceRunStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceManagementScriptIdDeviceRunStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Management Script Id Device Run State ID
func (id DeviceManagementDeviceManagementScriptIdDeviceRunStateId) ID() string {
	fmtString := "/deviceManagement/deviceManagementScripts/%s/deviceRunStates/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementScriptId, id.DeviceManagementScriptDeviceStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Management Script Id Device Run State ID
func (id DeviceManagementDeviceManagementScriptIdDeviceRunStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceManagementScripts", "deviceManagementScripts", "deviceManagementScripts"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptId", "deviceManagementScriptId"),
		resourceids.StaticSegment("deviceRunStates", "deviceRunStates", "deviceRunStates"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptDeviceStateId", "deviceManagementScriptDeviceStateId"),
	}
}

// String returns a human-readable description of this Device Management Device Management Script Id Device Run State ID
func (id DeviceManagementDeviceManagementScriptIdDeviceRunStateId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Script: %q", id.DeviceManagementScriptId),
		fmt.Sprintf("Device Management Script Device State: %q", id.DeviceManagementScriptDeviceStateId),
	}
	return fmt.Sprintf("Device Management Device Management Script Id Device Run State (%s)", strings.Join(components, "\n"))
}
