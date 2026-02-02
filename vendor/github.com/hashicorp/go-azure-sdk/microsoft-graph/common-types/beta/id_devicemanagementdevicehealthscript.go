package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceHealthScriptId{}

// DeviceManagementDeviceHealthScriptId is a struct representing the Resource ID for a Device Management Device Health Script
type DeviceManagementDeviceHealthScriptId struct {
	DeviceHealthScriptId string
}

// NewDeviceManagementDeviceHealthScriptID returns a new DeviceManagementDeviceHealthScriptId struct
func NewDeviceManagementDeviceHealthScriptID(deviceHealthScriptId string) DeviceManagementDeviceHealthScriptId {
	return DeviceManagementDeviceHealthScriptId{
		DeviceHealthScriptId: deviceHealthScriptId,
	}
}

// ParseDeviceManagementDeviceHealthScriptID parses 'input' into a DeviceManagementDeviceHealthScriptId
func ParseDeviceManagementDeviceHealthScriptID(input string) (*DeviceManagementDeviceHealthScriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceHealthScriptId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceHealthScriptId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceHealthScriptIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceHealthScriptId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceHealthScriptIDInsensitively(input string) (*DeviceManagementDeviceHealthScriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceHealthScriptId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceHealthScriptId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceHealthScriptId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceHealthScriptId, ok = input.Parsed["deviceHealthScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceHealthScriptId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceHealthScriptID checks that 'input' can be parsed as a Device Management Device Health Script ID
func ValidateDeviceManagementDeviceHealthScriptID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceHealthScriptID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Health Script ID
func (id DeviceManagementDeviceHealthScriptId) ID() string {
	fmtString := "/deviceManagement/deviceHealthScripts/%s"
	return fmt.Sprintf(fmtString, id.DeviceHealthScriptId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Health Script ID
func (id DeviceManagementDeviceHealthScriptId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceHealthScripts", "deviceHealthScripts", "deviceHealthScripts"),
		resourceids.UserSpecifiedSegment("deviceHealthScriptId", "deviceHealthScriptId"),
	}
}

// String returns a human-readable description of this Device Management Device Health Script ID
func (id DeviceManagementDeviceHealthScriptId) String() string {
	components := []string{
		fmt.Sprintf("Device Health Script: %q", id.DeviceHealthScriptId),
	}
	return fmt.Sprintf("Device Management Device Health Script (%s)", strings.Join(components, "\n"))
}
