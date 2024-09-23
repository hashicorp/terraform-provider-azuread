package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceShellScriptIdGroupAssignmentId{}

// DeviceManagementDeviceShellScriptIdGroupAssignmentId is a struct representing the Resource ID for a Device Management Device Shell Script Id Group Assignment
type DeviceManagementDeviceShellScriptIdGroupAssignmentId struct {
	DeviceShellScriptId                     string
	DeviceManagementScriptGroupAssignmentId string
}

// NewDeviceManagementDeviceShellScriptIdGroupAssignmentID returns a new DeviceManagementDeviceShellScriptIdGroupAssignmentId struct
func NewDeviceManagementDeviceShellScriptIdGroupAssignmentID(deviceShellScriptId string, deviceManagementScriptGroupAssignmentId string) DeviceManagementDeviceShellScriptIdGroupAssignmentId {
	return DeviceManagementDeviceShellScriptIdGroupAssignmentId{
		DeviceShellScriptId:                     deviceShellScriptId,
		DeviceManagementScriptGroupAssignmentId: deviceManagementScriptGroupAssignmentId,
	}
}

// ParseDeviceManagementDeviceShellScriptIdGroupAssignmentID parses 'input' into a DeviceManagementDeviceShellScriptIdGroupAssignmentId
func ParseDeviceManagementDeviceShellScriptIdGroupAssignmentID(input string) (*DeviceManagementDeviceShellScriptIdGroupAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceShellScriptIdGroupAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceShellScriptIdGroupAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceShellScriptIdGroupAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceShellScriptIdGroupAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceShellScriptIdGroupAssignmentIDInsensitively(input string) (*DeviceManagementDeviceShellScriptIdGroupAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceShellScriptIdGroupAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceShellScriptIdGroupAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceShellScriptIdGroupAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceShellScriptId, ok = input.Parsed["deviceShellScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceShellScriptId", input)
	}

	if id.DeviceManagementScriptGroupAssignmentId, ok = input.Parsed["deviceManagementScriptGroupAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptGroupAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceShellScriptIdGroupAssignmentID checks that 'input' can be parsed as a Device Management Device Shell Script Id Group Assignment ID
func ValidateDeviceManagementDeviceShellScriptIdGroupAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceShellScriptIdGroupAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Shell Script Id Group Assignment ID
func (id DeviceManagementDeviceShellScriptIdGroupAssignmentId) ID() string {
	fmtString := "/deviceManagement/deviceShellScripts/%s/groupAssignments/%s"
	return fmt.Sprintf(fmtString, id.DeviceShellScriptId, id.DeviceManagementScriptGroupAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Shell Script Id Group Assignment ID
func (id DeviceManagementDeviceShellScriptIdGroupAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceShellScripts", "deviceShellScripts", "deviceShellScripts"),
		resourceids.UserSpecifiedSegment("deviceShellScriptId", "deviceShellScriptId"),
		resourceids.StaticSegment("groupAssignments", "groupAssignments", "groupAssignments"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptGroupAssignmentId", "deviceManagementScriptGroupAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Device Shell Script Id Group Assignment ID
func (id DeviceManagementDeviceShellScriptIdGroupAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Device Shell Script: %q", id.DeviceShellScriptId),
		fmt.Sprintf("Device Management Script Group Assignment: %q", id.DeviceManagementScriptGroupAssignmentId),
	}
	return fmt.Sprintf("Device Management Device Shell Script Id Group Assignment (%s)", strings.Join(components, "\n"))
}
