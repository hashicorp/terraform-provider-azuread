package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateId{}

// DeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateId is a struct representing the Resource ID for a Device Management Device Management Script Id User Run State Id Device Run State
type DeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateId struct {
	DeviceManagementScriptId            string
	DeviceManagementScriptUserStateId   string
	DeviceManagementScriptDeviceStateId string
}

// NewDeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateID returns a new DeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateId struct
func NewDeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateID(deviceManagementScriptId string, deviceManagementScriptUserStateId string, deviceManagementScriptDeviceStateId string) DeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateId {
	return DeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateId{
		DeviceManagementScriptId:            deviceManagementScriptId,
		DeviceManagementScriptUserStateId:   deviceManagementScriptUserStateId,
		DeviceManagementScriptDeviceStateId: deviceManagementScriptDeviceStateId,
	}
}

// ParseDeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateID parses 'input' into a DeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateId
func ParseDeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateID(input string) (*DeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateIDInsensitively(input string) (*DeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementScriptId, ok = input.Parsed["deviceManagementScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptId", input)
	}

	if id.DeviceManagementScriptUserStateId, ok = input.Parsed["deviceManagementScriptUserStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptUserStateId", input)
	}

	if id.DeviceManagementScriptDeviceStateId, ok = input.Parsed["deviceManagementScriptDeviceStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptDeviceStateId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateID checks that 'input' can be parsed as a Device Management Device Management Script Id User Run State Id Device Run State ID
func ValidateDeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Management Script Id User Run State Id Device Run State ID
func (id DeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateId) ID() string {
	fmtString := "/deviceManagement/deviceManagementScripts/%s/userRunStates/%s/deviceRunStates/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementScriptId, id.DeviceManagementScriptUserStateId, id.DeviceManagementScriptDeviceStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Management Script Id User Run State Id Device Run State ID
func (id DeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceManagementScripts", "deviceManagementScripts", "deviceManagementScripts"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptId", "deviceManagementScriptId"),
		resourceids.StaticSegment("userRunStates", "userRunStates", "userRunStates"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptUserStateId", "deviceManagementScriptUserStateId"),
		resourceids.StaticSegment("deviceRunStates", "deviceRunStates", "deviceRunStates"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptDeviceStateId", "deviceManagementScriptDeviceStateId"),
	}
}

// String returns a human-readable description of this Device Management Device Management Script Id User Run State Id Device Run State ID
func (id DeviceManagementDeviceManagementScriptIdUserRunStateIdDeviceRunStateId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Script: %q", id.DeviceManagementScriptId),
		fmt.Sprintf("Device Management Script User State: %q", id.DeviceManagementScriptUserStateId),
		fmt.Sprintf("Device Management Script Device State: %q", id.DeviceManagementScriptDeviceStateId),
	}
	return fmt.Sprintf("Device Management Device Management Script Id User Run State Id Device Run State (%s)", strings.Join(components, "\n"))
}
