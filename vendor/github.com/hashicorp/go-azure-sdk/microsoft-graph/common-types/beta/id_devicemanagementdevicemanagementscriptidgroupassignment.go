package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceManagementScriptIdGroupAssignmentId{}

// DeviceManagementDeviceManagementScriptIdGroupAssignmentId is a struct representing the Resource ID for a Device Management Device Management Script Id Group Assignment
type DeviceManagementDeviceManagementScriptIdGroupAssignmentId struct {
	DeviceManagementScriptId                string
	DeviceManagementScriptGroupAssignmentId string
}

// NewDeviceManagementDeviceManagementScriptIdGroupAssignmentID returns a new DeviceManagementDeviceManagementScriptIdGroupAssignmentId struct
func NewDeviceManagementDeviceManagementScriptIdGroupAssignmentID(deviceManagementScriptId string, deviceManagementScriptGroupAssignmentId string) DeviceManagementDeviceManagementScriptIdGroupAssignmentId {
	return DeviceManagementDeviceManagementScriptIdGroupAssignmentId{
		DeviceManagementScriptId:                deviceManagementScriptId,
		DeviceManagementScriptGroupAssignmentId: deviceManagementScriptGroupAssignmentId,
	}
}

// ParseDeviceManagementDeviceManagementScriptIdGroupAssignmentID parses 'input' into a DeviceManagementDeviceManagementScriptIdGroupAssignmentId
func ParseDeviceManagementDeviceManagementScriptIdGroupAssignmentID(input string) (*DeviceManagementDeviceManagementScriptIdGroupAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceManagementScriptIdGroupAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceManagementScriptIdGroupAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceManagementScriptIdGroupAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceManagementScriptIdGroupAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceManagementScriptIdGroupAssignmentIDInsensitively(input string) (*DeviceManagementDeviceManagementScriptIdGroupAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceManagementScriptIdGroupAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceManagementScriptIdGroupAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceManagementScriptIdGroupAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementScriptId, ok = input.Parsed["deviceManagementScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptId", input)
	}

	if id.DeviceManagementScriptGroupAssignmentId, ok = input.Parsed["deviceManagementScriptGroupAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptGroupAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceManagementScriptIdGroupAssignmentID checks that 'input' can be parsed as a Device Management Device Management Script Id Group Assignment ID
func ValidateDeviceManagementDeviceManagementScriptIdGroupAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceManagementScriptIdGroupAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Management Script Id Group Assignment ID
func (id DeviceManagementDeviceManagementScriptIdGroupAssignmentId) ID() string {
	fmtString := "/deviceManagement/deviceManagementScripts/%s/groupAssignments/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementScriptId, id.DeviceManagementScriptGroupAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Management Script Id Group Assignment ID
func (id DeviceManagementDeviceManagementScriptIdGroupAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceManagementScripts", "deviceManagementScripts", "deviceManagementScripts"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptId", "deviceManagementScriptId"),
		resourceids.StaticSegment("groupAssignments", "groupAssignments", "groupAssignments"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptGroupAssignmentId", "deviceManagementScriptGroupAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Device Management Script Id Group Assignment ID
func (id DeviceManagementDeviceManagementScriptIdGroupAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Script: %q", id.DeviceManagementScriptId),
		fmt.Sprintf("Device Management Script Group Assignment: %q", id.DeviceManagementScriptGroupAssignmentId),
	}
	return fmt.Sprintf("Device Management Device Management Script Id Group Assignment (%s)", strings.Join(components, "\n"))
}
