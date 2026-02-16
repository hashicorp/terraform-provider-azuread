package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateId{}

// DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateId is a struct representing the Resource ID for a Device Management Device Custom Attribute Shell Script Id User Run State
type DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateId struct {
	DeviceCustomAttributeShellScriptId string
	DeviceManagementScriptUserStateId  string
}

// NewDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateID returns a new DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateId struct
func NewDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateID(deviceCustomAttributeShellScriptId string, deviceManagementScriptUserStateId string) DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateId {
	return DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateId{
		DeviceCustomAttributeShellScriptId: deviceCustomAttributeShellScriptId,
		DeviceManagementScriptUserStateId:  deviceManagementScriptUserStateId,
	}
}

// ParseDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateID parses 'input' into a DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateId
func ParseDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateID(input string) (*DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIDInsensitively(input string) (*DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceCustomAttributeShellScriptId, ok = input.Parsed["deviceCustomAttributeShellScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceCustomAttributeShellScriptId", input)
	}

	if id.DeviceManagementScriptUserStateId, ok = input.Parsed["deviceManagementScriptUserStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptUserStateId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateID checks that 'input' can be parsed as a Device Management Device Custom Attribute Shell Script Id User Run State ID
func ValidateDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Custom Attribute Shell Script Id User Run State ID
func (id DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateId) ID() string {
	fmtString := "/deviceManagement/deviceCustomAttributeShellScripts/%s/userRunStates/%s"
	return fmt.Sprintf(fmtString, id.DeviceCustomAttributeShellScriptId, id.DeviceManagementScriptUserStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Custom Attribute Shell Script Id User Run State ID
func (id DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceCustomAttributeShellScripts", "deviceCustomAttributeShellScripts", "deviceCustomAttributeShellScripts"),
		resourceids.UserSpecifiedSegment("deviceCustomAttributeShellScriptId", "deviceCustomAttributeShellScriptId"),
		resourceids.StaticSegment("userRunStates", "userRunStates", "userRunStates"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptUserStateId", "deviceManagementScriptUserStateId"),
	}
}

// String returns a human-readable description of this Device Management Device Custom Attribute Shell Script Id User Run State ID
func (id DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateId) String() string {
	components := []string{
		fmt.Sprintf("Device Custom Attribute Shell Script: %q", id.DeviceCustomAttributeShellScriptId),
		fmt.Sprintf("Device Management Script User State: %q", id.DeviceManagementScriptUserStateId),
	}
	return fmt.Sprintf("Device Management Device Custom Attribute Shell Script Id User Run State (%s)", strings.Join(components, "\n"))
}
