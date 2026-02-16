package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId{}

// DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId is a struct representing the Resource ID for a Device Management Device Custom Attribute Shell Script Id Assignment
type DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId struct {
	DeviceCustomAttributeShellScriptId string
	DeviceManagementScriptAssignmentId string
}

// NewDeviceManagementDeviceCustomAttributeShellScriptIdAssignmentID returns a new DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId struct
func NewDeviceManagementDeviceCustomAttributeShellScriptIdAssignmentID(deviceCustomAttributeShellScriptId string, deviceManagementScriptAssignmentId string) DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId {
	return DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId{
		DeviceCustomAttributeShellScriptId: deviceCustomAttributeShellScriptId,
		DeviceManagementScriptAssignmentId: deviceManagementScriptAssignmentId,
	}
}

// ParseDeviceManagementDeviceCustomAttributeShellScriptIdAssignmentID parses 'input' into a DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId
func ParseDeviceManagementDeviceCustomAttributeShellScriptIdAssignmentID(input string) (*DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceCustomAttributeShellScriptIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceCustomAttributeShellScriptIdAssignmentIDInsensitively(input string) (*DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceCustomAttributeShellScriptId, ok = input.Parsed["deviceCustomAttributeShellScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceCustomAttributeShellScriptId", input)
	}

	if id.DeviceManagementScriptAssignmentId, ok = input.Parsed["deviceManagementScriptAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceCustomAttributeShellScriptIdAssignmentID checks that 'input' can be parsed as a Device Management Device Custom Attribute Shell Script Id Assignment ID
func ValidateDeviceManagementDeviceCustomAttributeShellScriptIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceCustomAttributeShellScriptIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Custom Attribute Shell Script Id Assignment ID
func (id DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/deviceCustomAttributeShellScripts/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.DeviceCustomAttributeShellScriptId, id.DeviceManagementScriptAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Custom Attribute Shell Script Id Assignment ID
func (id DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceCustomAttributeShellScripts", "deviceCustomAttributeShellScripts", "deviceCustomAttributeShellScripts"),
		resourceids.UserSpecifiedSegment("deviceCustomAttributeShellScriptId", "deviceCustomAttributeShellScriptId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptAssignmentId", "deviceManagementScriptAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Device Custom Attribute Shell Script Id Assignment ID
func (id DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Device Custom Attribute Shell Script: %q", id.DeviceCustomAttributeShellScriptId),
		fmt.Sprintf("Device Management Script Assignment: %q", id.DeviceManagementScriptAssignmentId),
	}
	return fmt.Sprintf("Device Management Device Custom Attribute Shell Script Id Assignment (%s)", strings.Join(components, "\n"))
}
