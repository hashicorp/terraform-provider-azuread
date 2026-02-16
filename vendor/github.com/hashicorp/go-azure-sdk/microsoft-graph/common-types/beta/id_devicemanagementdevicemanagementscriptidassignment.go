package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceManagementScriptIdAssignmentId{}

// DeviceManagementDeviceManagementScriptIdAssignmentId is a struct representing the Resource ID for a Device Management Device Management Script Id Assignment
type DeviceManagementDeviceManagementScriptIdAssignmentId struct {
	DeviceManagementScriptId           string
	DeviceManagementScriptAssignmentId string
}

// NewDeviceManagementDeviceManagementScriptIdAssignmentID returns a new DeviceManagementDeviceManagementScriptIdAssignmentId struct
func NewDeviceManagementDeviceManagementScriptIdAssignmentID(deviceManagementScriptId string, deviceManagementScriptAssignmentId string) DeviceManagementDeviceManagementScriptIdAssignmentId {
	return DeviceManagementDeviceManagementScriptIdAssignmentId{
		DeviceManagementScriptId:           deviceManagementScriptId,
		DeviceManagementScriptAssignmentId: deviceManagementScriptAssignmentId,
	}
}

// ParseDeviceManagementDeviceManagementScriptIdAssignmentID parses 'input' into a DeviceManagementDeviceManagementScriptIdAssignmentId
func ParseDeviceManagementDeviceManagementScriptIdAssignmentID(input string) (*DeviceManagementDeviceManagementScriptIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceManagementScriptIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceManagementScriptIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceManagementScriptIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceManagementScriptIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceManagementScriptIdAssignmentIDInsensitively(input string) (*DeviceManagementDeviceManagementScriptIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceManagementScriptIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceManagementScriptIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceManagementScriptIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementScriptId, ok = input.Parsed["deviceManagementScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptId", input)
	}

	if id.DeviceManagementScriptAssignmentId, ok = input.Parsed["deviceManagementScriptAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceManagementScriptIdAssignmentID checks that 'input' can be parsed as a Device Management Device Management Script Id Assignment ID
func ValidateDeviceManagementDeviceManagementScriptIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceManagementScriptIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Management Script Id Assignment ID
func (id DeviceManagementDeviceManagementScriptIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/deviceManagementScripts/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementScriptId, id.DeviceManagementScriptAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Management Script Id Assignment ID
func (id DeviceManagementDeviceManagementScriptIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceManagementScripts", "deviceManagementScripts", "deviceManagementScripts"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptId", "deviceManagementScriptId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptAssignmentId", "deviceManagementScriptAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Device Management Script Id Assignment ID
func (id DeviceManagementDeviceManagementScriptIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Script: %q", id.DeviceManagementScriptId),
		fmt.Sprintf("Device Management Script Assignment: %q", id.DeviceManagementScriptAssignmentId),
	}
	return fmt.Sprintf("Device Management Device Management Script Id Assignment (%s)", strings.Join(components, "\n"))
}
