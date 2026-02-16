package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId{}

// DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId is a struct representing the Resource ID for a Device Management Device Custom Attribute Shell Script Id Group Assignment
type DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId struct {
	DeviceCustomAttributeShellScriptId      string
	DeviceManagementScriptGroupAssignmentId string
}

// NewDeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentID returns a new DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId struct
func NewDeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentID(deviceCustomAttributeShellScriptId string, deviceManagementScriptGroupAssignmentId string) DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId {
	return DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId{
		DeviceCustomAttributeShellScriptId:      deviceCustomAttributeShellScriptId,
		DeviceManagementScriptGroupAssignmentId: deviceManagementScriptGroupAssignmentId,
	}
}

// ParseDeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentID parses 'input' into a DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId
func ParseDeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentID(input string) (*DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentIDInsensitively(input string) (*DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceCustomAttributeShellScriptId, ok = input.Parsed["deviceCustomAttributeShellScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceCustomAttributeShellScriptId", input)
	}

	if id.DeviceManagementScriptGroupAssignmentId, ok = input.Parsed["deviceManagementScriptGroupAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptGroupAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentID checks that 'input' can be parsed as a Device Management Device Custom Attribute Shell Script Id Group Assignment ID
func ValidateDeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Custom Attribute Shell Script Id Group Assignment ID
func (id DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId) ID() string {
	fmtString := "/deviceManagement/deviceCustomAttributeShellScripts/%s/groupAssignments/%s"
	return fmt.Sprintf(fmtString, id.DeviceCustomAttributeShellScriptId, id.DeviceManagementScriptGroupAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Custom Attribute Shell Script Id Group Assignment ID
func (id DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceCustomAttributeShellScripts", "deviceCustomAttributeShellScripts", "deviceCustomAttributeShellScripts"),
		resourceids.UserSpecifiedSegment("deviceCustomAttributeShellScriptId", "deviceCustomAttributeShellScriptId"),
		resourceids.StaticSegment("groupAssignments", "groupAssignments", "groupAssignments"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptGroupAssignmentId", "deviceManagementScriptGroupAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Device Custom Attribute Shell Script Id Group Assignment ID
func (id DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Device Custom Attribute Shell Script: %q", id.DeviceCustomAttributeShellScriptId),
		fmt.Sprintf("Device Management Script Group Assignment: %q", id.DeviceManagementScriptGroupAssignmentId),
	}
	return fmt.Sprintf("Device Management Device Custom Attribute Shell Script Id Group Assignment (%s)", strings.Join(components, "\n"))
}
