package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceShellScriptIdUserRunStateId{}

// DeviceManagementDeviceShellScriptIdUserRunStateId is a struct representing the Resource ID for a Device Management Device Shell Script Id User Run State
type DeviceManagementDeviceShellScriptIdUserRunStateId struct {
	DeviceShellScriptId               string
	DeviceManagementScriptUserStateId string
}

// NewDeviceManagementDeviceShellScriptIdUserRunStateID returns a new DeviceManagementDeviceShellScriptIdUserRunStateId struct
func NewDeviceManagementDeviceShellScriptIdUserRunStateID(deviceShellScriptId string, deviceManagementScriptUserStateId string) DeviceManagementDeviceShellScriptIdUserRunStateId {
	return DeviceManagementDeviceShellScriptIdUserRunStateId{
		DeviceShellScriptId:               deviceShellScriptId,
		DeviceManagementScriptUserStateId: deviceManagementScriptUserStateId,
	}
}

// ParseDeviceManagementDeviceShellScriptIdUserRunStateID parses 'input' into a DeviceManagementDeviceShellScriptIdUserRunStateId
func ParseDeviceManagementDeviceShellScriptIdUserRunStateID(input string) (*DeviceManagementDeviceShellScriptIdUserRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceShellScriptIdUserRunStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceShellScriptIdUserRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceShellScriptIdUserRunStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceShellScriptIdUserRunStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceShellScriptIdUserRunStateIDInsensitively(input string) (*DeviceManagementDeviceShellScriptIdUserRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceShellScriptIdUserRunStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceShellScriptIdUserRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceShellScriptIdUserRunStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceShellScriptId, ok = input.Parsed["deviceShellScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceShellScriptId", input)
	}

	if id.DeviceManagementScriptUserStateId, ok = input.Parsed["deviceManagementScriptUserStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptUserStateId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceShellScriptIdUserRunStateID checks that 'input' can be parsed as a Device Management Device Shell Script Id User Run State ID
func ValidateDeviceManagementDeviceShellScriptIdUserRunStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceShellScriptIdUserRunStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Shell Script Id User Run State ID
func (id DeviceManagementDeviceShellScriptIdUserRunStateId) ID() string {
	fmtString := "/deviceManagement/deviceShellScripts/%s/userRunStates/%s"
	return fmt.Sprintf(fmtString, id.DeviceShellScriptId, id.DeviceManagementScriptUserStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Shell Script Id User Run State ID
func (id DeviceManagementDeviceShellScriptIdUserRunStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceShellScripts", "deviceShellScripts", "deviceShellScripts"),
		resourceids.UserSpecifiedSegment("deviceShellScriptId", "deviceShellScriptId"),
		resourceids.StaticSegment("userRunStates", "userRunStates", "userRunStates"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptUserStateId", "deviceManagementScriptUserStateId"),
	}
}

// String returns a human-readable description of this Device Management Device Shell Script Id User Run State ID
func (id DeviceManagementDeviceShellScriptIdUserRunStateId) String() string {
	components := []string{
		fmt.Sprintf("Device Shell Script: %q", id.DeviceShellScriptId),
		fmt.Sprintf("Device Management Script User State: %q", id.DeviceManagementScriptUserStateId),
	}
	return fmt.Sprintf("Device Management Device Shell Script Id User Run State (%s)", strings.Join(components, "\n"))
}
