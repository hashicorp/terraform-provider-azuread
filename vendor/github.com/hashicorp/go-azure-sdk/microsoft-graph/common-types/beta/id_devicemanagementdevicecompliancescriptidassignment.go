package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceComplianceScriptIdAssignmentId{}

// DeviceManagementDeviceComplianceScriptIdAssignmentId is a struct representing the Resource ID for a Device Management Device Compliance Script Id Assignment
type DeviceManagementDeviceComplianceScriptIdAssignmentId struct {
	DeviceComplianceScriptId       string
	DeviceHealthScriptAssignmentId string
}

// NewDeviceManagementDeviceComplianceScriptIdAssignmentID returns a new DeviceManagementDeviceComplianceScriptIdAssignmentId struct
func NewDeviceManagementDeviceComplianceScriptIdAssignmentID(deviceComplianceScriptId string, deviceHealthScriptAssignmentId string) DeviceManagementDeviceComplianceScriptIdAssignmentId {
	return DeviceManagementDeviceComplianceScriptIdAssignmentId{
		DeviceComplianceScriptId:       deviceComplianceScriptId,
		DeviceHealthScriptAssignmentId: deviceHealthScriptAssignmentId,
	}
}

// ParseDeviceManagementDeviceComplianceScriptIdAssignmentID parses 'input' into a DeviceManagementDeviceComplianceScriptIdAssignmentId
func ParseDeviceManagementDeviceComplianceScriptIdAssignmentID(input string) (*DeviceManagementDeviceComplianceScriptIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceComplianceScriptIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceComplianceScriptIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceComplianceScriptIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceComplianceScriptIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceComplianceScriptIdAssignmentIDInsensitively(input string) (*DeviceManagementDeviceComplianceScriptIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceComplianceScriptIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceComplianceScriptIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceComplianceScriptIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceComplianceScriptId, ok = input.Parsed["deviceComplianceScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceComplianceScriptId", input)
	}

	if id.DeviceHealthScriptAssignmentId, ok = input.Parsed["deviceHealthScriptAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceHealthScriptAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceComplianceScriptIdAssignmentID checks that 'input' can be parsed as a Device Management Device Compliance Script Id Assignment ID
func ValidateDeviceManagementDeviceComplianceScriptIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceComplianceScriptIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Compliance Script Id Assignment ID
func (id DeviceManagementDeviceComplianceScriptIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/deviceComplianceScripts/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.DeviceComplianceScriptId, id.DeviceHealthScriptAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Compliance Script Id Assignment ID
func (id DeviceManagementDeviceComplianceScriptIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceComplianceScripts", "deviceComplianceScripts", "deviceComplianceScripts"),
		resourceids.UserSpecifiedSegment("deviceComplianceScriptId", "deviceComplianceScriptId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("deviceHealthScriptAssignmentId", "deviceHealthScriptAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Device Compliance Script Id Assignment ID
func (id DeviceManagementDeviceComplianceScriptIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Device Compliance Script: %q", id.DeviceComplianceScriptId),
		fmt.Sprintf("Device Health Script Assignment: %q", id.DeviceHealthScriptAssignmentId),
	}
	return fmt.Sprintf("Device Management Device Compliance Script Id Assignment (%s)", strings.Join(components, "\n"))
}
