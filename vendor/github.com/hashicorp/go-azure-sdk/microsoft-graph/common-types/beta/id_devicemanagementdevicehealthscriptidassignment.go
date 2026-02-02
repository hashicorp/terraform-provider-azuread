package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceHealthScriptIdAssignmentId{}

// DeviceManagementDeviceHealthScriptIdAssignmentId is a struct representing the Resource ID for a Device Management Device Health Script Id Assignment
type DeviceManagementDeviceHealthScriptIdAssignmentId struct {
	DeviceHealthScriptId           string
	DeviceHealthScriptAssignmentId string
}

// NewDeviceManagementDeviceHealthScriptIdAssignmentID returns a new DeviceManagementDeviceHealthScriptIdAssignmentId struct
func NewDeviceManagementDeviceHealthScriptIdAssignmentID(deviceHealthScriptId string, deviceHealthScriptAssignmentId string) DeviceManagementDeviceHealthScriptIdAssignmentId {
	return DeviceManagementDeviceHealthScriptIdAssignmentId{
		DeviceHealthScriptId:           deviceHealthScriptId,
		DeviceHealthScriptAssignmentId: deviceHealthScriptAssignmentId,
	}
}

// ParseDeviceManagementDeviceHealthScriptIdAssignmentID parses 'input' into a DeviceManagementDeviceHealthScriptIdAssignmentId
func ParseDeviceManagementDeviceHealthScriptIdAssignmentID(input string) (*DeviceManagementDeviceHealthScriptIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceHealthScriptIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceHealthScriptIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceHealthScriptIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceHealthScriptIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceHealthScriptIdAssignmentIDInsensitively(input string) (*DeviceManagementDeviceHealthScriptIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceHealthScriptIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceHealthScriptIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceHealthScriptIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceHealthScriptId, ok = input.Parsed["deviceHealthScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceHealthScriptId", input)
	}

	if id.DeviceHealthScriptAssignmentId, ok = input.Parsed["deviceHealthScriptAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceHealthScriptAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceHealthScriptIdAssignmentID checks that 'input' can be parsed as a Device Management Device Health Script Id Assignment ID
func ValidateDeviceManagementDeviceHealthScriptIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceHealthScriptIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Health Script Id Assignment ID
func (id DeviceManagementDeviceHealthScriptIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/deviceHealthScripts/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.DeviceHealthScriptId, id.DeviceHealthScriptAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Health Script Id Assignment ID
func (id DeviceManagementDeviceHealthScriptIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceHealthScripts", "deviceHealthScripts", "deviceHealthScripts"),
		resourceids.UserSpecifiedSegment("deviceHealthScriptId", "deviceHealthScriptId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("deviceHealthScriptAssignmentId", "deviceHealthScriptAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Device Health Script Id Assignment ID
func (id DeviceManagementDeviceHealthScriptIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Device Health Script: %q", id.DeviceHealthScriptId),
		fmt.Sprintf("Device Health Script Assignment: %q", id.DeviceHealthScriptAssignmentId),
	}
	return fmt.Sprintf("Device Management Device Health Script Id Assignment (%s)", strings.Join(components, "\n"))
}
