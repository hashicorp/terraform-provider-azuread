package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceShellScriptId{}

// DeviceManagementDeviceShellScriptId is a struct representing the Resource ID for a Device Management Device Shell Script
type DeviceManagementDeviceShellScriptId struct {
	DeviceShellScriptId string
}

// NewDeviceManagementDeviceShellScriptID returns a new DeviceManagementDeviceShellScriptId struct
func NewDeviceManagementDeviceShellScriptID(deviceShellScriptId string) DeviceManagementDeviceShellScriptId {
	return DeviceManagementDeviceShellScriptId{
		DeviceShellScriptId: deviceShellScriptId,
	}
}

// ParseDeviceManagementDeviceShellScriptID parses 'input' into a DeviceManagementDeviceShellScriptId
func ParseDeviceManagementDeviceShellScriptID(input string) (*DeviceManagementDeviceShellScriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceShellScriptId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceShellScriptId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceShellScriptIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceShellScriptId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceShellScriptIDInsensitively(input string) (*DeviceManagementDeviceShellScriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceShellScriptId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceShellScriptId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceShellScriptId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceShellScriptId, ok = input.Parsed["deviceShellScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceShellScriptId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceShellScriptID checks that 'input' can be parsed as a Device Management Device Shell Script ID
func ValidateDeviceManagementDeviceShellScriptID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceShellScriptID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Shell Script ID
func (id DeviceManagementDeviceShellScriptId) ID() string {
	fmtString := "/deviceManagement/deviceShellScripts/%s"
	return fmt.Sprintf(fmtString, id.DeviceShellScriptId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Shell Script ID
func (id DeviceManagementDeviceShellScriptId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceShellScripts", "deviceShellScripts", "deviceShellScripts"),
		resourceids.UserSpecifiedSegment("deviceShellScriptId", "deviceShellScriptId"),
	}
}

// String returns a human-readable description of this Device Management Device Shell Script ID
func (id DeviceManagementDeviceShellScriptId) String() string {
	components := []string{
		fmt.Sprintf("Device Shell Script: %q", id.DeviceShellScriptId),
	}
	return fmt.Sprintf("Device Management Device Shell Script (%s)", strings.Join(components, "\n"))
}
