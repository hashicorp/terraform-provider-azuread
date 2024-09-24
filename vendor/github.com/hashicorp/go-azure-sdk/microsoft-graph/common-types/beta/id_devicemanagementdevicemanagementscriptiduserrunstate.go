package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceManagementScriptIdUserRunStateId{}

// DeviceManagementDeviceManagementScriptIdUserRunStateId is a struct representing the Resource ID for a Device Management Device Management Script Id User Run State
type DeviceManagementDeviceManagementScriptIdUserRunStateId struct {
	DeviceManagementScriptId          string
	DeviceManagementScriptUserStateId string
}

// NewDeviceManagementDeviceManagementScriptIdUserRunStateID returns a new DeviceManagementDeviceManagementScriptIdUserRunStateId struct
func NewDeviceManagementDeviceManagementScriptIdUserRunStateID(deviceManagementScriptId string, deviceManagementScriptUserStateId string) DeviceManagementDeviceManagementScriptIdUserRunStateId {
	return DeviceManagementDeviceManagementScriptIdUserRunStateId{
		DeviceManagementScriptId:          deviceManagementScriptId,
		DeviceManagementScriptUserStateId: deviceManagementScriptUserStateId,
	}
}

// ParseDeviceManagementDeviceManagementScriptIdUserRunStateID parses 'input' into a DeviceManagementDeviceManagementScriptIdUserRunStateId
func ParseDeviceManagementDeviceManagementScriptIdUserRunStateID(input string) (*DeviceManagementDeviceManagementScriptIdUserRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceManagementScriptIdUserRunStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceManagementScriptIdUserRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceManagementScriptIdUserRunStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceManagementScriptIdUserRunStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceManagementScriptIdUserRunStateIDInsensitively(input string) (*DeviceManagementDeviceManagementScriptIdUserRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceManagementScriptIdUserRunStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceManagementScriptIdUserRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceManagementScriptIdUserRunStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementScriptId, ok = input.Parsed["deviceManagementScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptId", input)
	}

	if id.DeviceManagementScriptUserStateId, ok = input.Parsed["deviceManagementScriptUserStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptUserStateId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceManagementScriptIdUserRunStateID checks that 'input' can be parsed as a Device Management Device Management Script Id User Run State ID
func ValidateDeviceManagementDeviceManagementScriptIdUserRunStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceManagementScriptIdUserRunStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Management Script Id User Run State ID
func (id DeviceManagementDeviceManagementScriptIdUserRunStateId) ID() string {
	fmtString := "/deviceManagement/deviceManagementScripts/%s/userRunStates/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementScriptId, id.DeviceManagementScriptUserStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Management Script Id User Run State ID
func (id DeviceManagementDeviceManagementScriptIdUserRunStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceManagementScripts", "deviceManagementScripts", "deviceManagementScripts"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptId", "deviceManagementScriptId"),
		resourceids.StaticSegment("userRunStates", "userRunStates", "userRunStates"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptUserStateId", "deviceManagementScriptUserStateId"),
	}
}

// String returns a human-readable description of this Device Management Device Management Script Id User Run State ID
func (id DeviceManagementDeviceManagementScriptIdUserRunStateId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Script: %q", id.DeviceManagementScriptId),
		fmt.Sprintf("Device Management Script User State: %q", id.DeviceManagementScriptUserStateId),
	}
	return fmt.Sprintf("Device Management Device Management Script Id User Run State (%s)", strings.Join(components, "\n"))
}
