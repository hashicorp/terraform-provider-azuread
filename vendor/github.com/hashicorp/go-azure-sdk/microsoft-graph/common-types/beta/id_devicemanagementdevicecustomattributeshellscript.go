package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCustomAttributeShellScriptId{}

// DeviceManagementDeviceCustomAttributeShellScriptId is a struct representing the Resource ID for a Device Management Device Custom Attribute Shell Script
type DeviceManagementDeviceCustomAttributeShellScriptId struct {
	DeviceCustomAttributeShellScriptId string
}

// NewDeviceManagementDeviceCustomAttributeShellScriptID returns a new DeviceManagementDeviceCustomAttributeShellScriptId struct
func NewDeviceManagementDeviceCustomAttributeShellScriptID(deviceCustomAttributeShellScriptId string) DeviceManagementDeviceCustomAttributeShellScriptId {
	return DeviceManagementDeviceCustomAttributeShellScriptId{
		DeviceCustomAttributeShellScriptId: deviceCustomAttributeShellScriptId,
	}
}

// ParseDeviceManagementDeviceCustomAttributeShellScriptID parses 'input' into a DeviceManagementDeviceCustomAttributeShellScriptId
func ParseDeviceManagementDeviceCustomAttributeShellScriptID(input string) (*DeviceManagementDeviceCustomAttributeShellScriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCustomAttributeShellScriptId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCustomAttributeShellScriptId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceCustomAttributeShellScriptIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceCustomAttributeShellScriptId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceCustomAttributeShellScriptIDInsensitively(input string) (*DeviceManagementDeviceCustomAttributeShellScriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCustomAttributeShellScriptId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCustomAttributeShellScriptId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceCustomAttributeShellScriptId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceCustomAttributeShellScriptId, ok = input.Parsed["deviceCustomAttributeShellScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceCustomAttributeShellScriptId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceCustomAttributeShellScriptID checks that 'input' can be parsed as a Device Management Device Custom Attribute Shell Script ID
func ValidateDeviceManagementDeviceCustomAttributeShellScriptID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceCustomAttributeShellScriptID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Custom Attribute Shell Script ID
func (id DeviceManagementDeviceCustomAttributeShellScriptId) ID() string {
	fmtString := "/deviceManagement/deviceCustomAttributeShellScripts/%s"
	return fmt.Sprintf(fmtString, id.DeviceCustomAttributeShellScriptId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Custom Attribute Shell Script ID
func (id DeviceManagementDeviceCustomAttributeShellScriptId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceCustomAttributeShellScripts", "deviceCustomAttributeShellScripts", "deviceCustomAttributeShellScripts"),
		resourceids.UserSpecifiedSegment("deviceCustomAttributeShellScriptId", "deviceCustomAttributeShellScriptId"),
	}
}

// String returns a human-readable description of this Device Management Device Custom Attribute Shell Script ID
func (id DeviceManagementDeviceCustomAttributeShellScriptId) String() string {
	components := []string{
		fmt.Sprintf("Device Custom Attribute Shell Script: %q", id.DeviceCustomAttributeShellScriptId),
	}
	return fmt.Sprintf("Device Management Device Custom Attribute Shell Script (%s)", strings.Join(components, "\n"))
}
