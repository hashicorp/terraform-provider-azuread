package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceComplianceScriptIdDeviceRunStateId{}

// DeviceManagementDeviceComplianceScriptIdDeviceRunStateId is a struct representing the Resource ID for a Device Management Device Compliance Script Id Device Run State
type DeviceManagementDeviceComplianceScriptIdDeviceRunStateId struct {
	DeviceComplianceScriptId            string
	DeviceComplianceScriptDeviceStateId string
}

// NewDeviceManagementDeviceComplianceScriptIdDeviceRunStateID returns a new DeviceManagementDeviceComplianceScriptIdDeviceRunStateId struct
func NewDeviceManagementDeviceComplianceScriptIdDeviceRunStateID(deviceComplianceScriptId string, deviceComplianceScriptDeviceStateId string) DeviceManagementDeviceComplianceScriptIdDeviceRunStateId {
	return DeviceManagementDeviceComplianceScriptIdDeviceRunStateId{
		DeviceComplianceScriptId:            deviceComplianceScriptId,
		DeviceComplianceScriptDeviceStateId: deviceComplianceScriptDeviceStateId,
	}
}

// ParseDeviceManagementDeviceComplianceScriptIdDeviceRunStateID parses 'input' into a DeviceManagementDeviceComplianceScriptIdDeviceRunStateId
func ParseDeviceManagementDeviceComplianceScriptIdDeviceRunStateID(input string) (*DeviceManagementDeviceComplianceScriptIdDeviceRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceComplianceScriptIdDeviceRunStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceComplianceScriptIdDeviceRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceComplianceScriptIdDeviceRunStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceComplianceScriptIdDeviceRunStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceComplianceScriptIdDeviceRunStateIDInsensitively(input string) (*DeviceManagementDeviceComplianceScriptIdDeviceRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceComplianceScriptIdDeviceRunStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceComplianceScriptIdDeviceRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceComplianceScriptIdDeviceRunStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceComplianceScriptId, ok = input.Parsed["deviceComplianceScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceComplianceScriptId", input)
	}

	if id.DeviceComplianceScriptDeviceStateId, ok = input.Parsed["deviceComplianceScriptDeviceStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceComplianceScriptDeviceStateId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceComplianceScriptIdDeviceRunStateID checks that 'input' can be parsed as a Device Management Device Compliance Script Id Device Run State ID
func ValidateDeviceManagementDeviceComplianceScriptIdDeviceRunStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceComplianceScriptIdDeviceRunStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Compliance Script Id Device Run State ID
func (id DeviceManagementDeviceComplianceScriptIdDeviceRunStateId) ID() string {
	fmtString := "/deviceManagement/deviceComplianceScripts/%s/deviceRunStates/%s"
	return fmt.Sprintf(fmtString, id.DeviceComplianceScriptId, id.DeviceComplianceScriptDeviceStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Compliance Script Id Device Run State ID
func (id DeviceManagementDeviceComplianceScriptIdDeviceRunStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceComplianceScripts", "deviceComplianceScripts", "deviceComplianceScripts"),
		resourceids.UserSpecifiedSegment("deviceComplianceScriptId", "deviceComplianceScriptId"),
		resourceids.StaticSegment("deviceRunStates", "deviceRunStates", "deviceRunStates"),
		resourceids.UserSpecifiedSegment("deviceComplianceScriptDeviceStateId", "deviceComplianceScriptDeviceStateId"),
	}
}

// String returns a human-readable description of this Device Management Device Compliance Script Id Device Run State ID
func (id DeviceManagementDeviceComplianceScriptIdDeviceRunStateId) String() string {
	components := []string{
		fmt.Sprintf("Device Compliance Script: %q", id.DeviceComplianceScriptId),
		fmt.Sprintf("Device Compliance Script Device State: %q", id.DeviceComplianceScriptDeviceStateId),
	}
	return fmt.Sprintf("Device Management Device Compliance Script Id Device Run State (%s)", strings.Join(components, "\n"))
}
