package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceHealthScriptIdDeviceRunStateId{}

// DeviceManagementDeviceHealthScriptIdDeviceRunStateId is a struct representing the Resource ID for a Device Management Device Health Script Id Device Run State
type DeviceManagementDeviceHealthScriptIdDeviceRunStateId struct {
	DeviceHealthScriptId            string
	DeviceHealthScriptDeviceStateId string
}

// NewDeviceManagementDeviceHealthScriptIdDeviceRunStateID returns a new DeviceManagementDeviceHealthScriptIdDeviceRunStateId struct
func NewDeviceManagementDeviceHealthScriptIdDeviceRunStateID(deviceHealthScriptId string, deviceHealthScriptDeviceStateId string) DeviceManagementDeviceHealthScriptIdDeviceRunStateId {
	return DeviceManagementDeviceHealthScriptIdDeviceRunStateId{
		DeviceHealthScriptId:            deviceHealthScriptId,
		DeviceHealthScriptDeviceStateId: deviceHealthScriptDeviceStateId,
	}
}

// ParseDeviceManagementDeviceHealthScriptIdDeviceRunStateID parses 'input' into a DeviceManagementDeviceHealthScriptIdDeviceRunStateId
func ParseDeviceManagementDeviceHealthScriptIdDeviceRunStateID(input string) (*DeviceManagementDeviceHealthScriptIdDeviceRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceHealthScriptIdDeviceRunStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceHealthScriptIdDeviceRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceHealthScriptIdDeviceRunStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceHealthScriptIdDeviceRunStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceHealthScriptIdDeviceRunStateIDInsensitively(input string) (*DeviceManagementDeviceHealthScriptIdDeviceRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceHealthScriptIdDeviceRunStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceHealthScriptIdDeviceRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceHealthScriptIdDeviceRunStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceHealthScriptId, ok = input.Parsed["deviceHealthScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceHealthScriptId", input)
	}

	if id.DeviceHealthScriptDeviceStateId, ok = input.Parsed["deviceHealthScriptDeviceStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceHealthScriptDeviceStateId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceHealthScriptIdDeviceRunStateID checks that 'input' can be parsed as a Device Management Device Health Script Id Device Run State ID
func ValidateDeviceManagementDeviceHealthScriptIdDeviceRunStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceHealthScriptIdDeviceRunStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Health Script Id Device Run State ID
func (id DeviceManagementDeviceHealthScriptIdDeviceRunStateId) ID() string {
	fmtString := "/deviceManagement/deviceHealthScripts/%s/deviceRunStates/%s"
	return fmt.Sprintf(fmtString, id.DeviceHealthScriptId, id.DeviceHealthScriptDeviceStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Health Script Id Device Run State ID
func (id DeviceManagementDeviceHealthScriptIdDeviceRunStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceHealthScripts", "deviceHealthScripts", "deviceHealthScripts"),
		resourceids.UserSpecifiedSegment("deviceHealthScriptId", "deviceHealthScriptId"),
		resourceids.StaticSegment("deviceRunStates", "deviceRunStates", "deviceRunStates"),
		resourceids.UserSpecifiedSegment("deviceHealthScriptDeviceStateId", "deviceHealthScriptDeviceStateId"),
	}
}

// String returns a human-readable description of this Device Management Device Health Script Id Device Run State ID
func (id DeviceManagementDeviceHealthScriptIdDeviceRunStateId) String() string {
	components := []string{
		fmt.Sprintf("Device Health Script: %q", id.DeviceHealthScriptId),
		fmt.Sprintf("Device Health Script Device State: %q", id.DeviceHealthScriptDeviceStateId),
	}
	return fmt.Sprintf("Device Management Device Health Script Id Device Run State (%s)", strings.Join(components, "\n"))
}
