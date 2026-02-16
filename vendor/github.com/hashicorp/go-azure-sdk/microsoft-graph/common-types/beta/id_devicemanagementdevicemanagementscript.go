package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceManagementScriptId{}

// DeviceManagementDeviceManagementScriptId is a struct representing the Resource ID for a Device Management Device Management Script
type DeviceManagementDeviceManagementScriptId struct {
	DeviceManagementScriptId string
}

// NewDeviceManagementDeviceManagementScriptID returns a new DeviceManagementDeviceManagementScriptId struct
func NewDeviceManagementDeviceManagementScriptID(deviceManagementScriptId string) DeviceManagementDeviceManagementScriptId {
	return DeviceManagementDeviceManagementScriptId{
		DeviceManagementScriptId: deviceManagementScriptId,
	}
}

// ParseDeviceManagementDeviceManagementScriptID parses 'input' into a DeviceManagementDeviceManagementScriptId
func ParseDeviceManagementDeviceManagementScriptID(input string) (*DeviceManagementDeviceManagementScriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceManagementScriptId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceManagementScriptId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceManagementScriptIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceManagementScriptId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceManagementScriptIDInsensitively(input string) (*DeviceManagementDeviceManagementScriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceManagementScriptId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceManagementScriptId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceManagementScriptId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementScriptId, ok = input.Parsed["deviceManagementScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementScriptId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceManagementScriptID checks that 'input' can be parsed as a Device Management Device Management Script ID
func ValidateDeviceManagementDeviceManagementScriptID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceManagementScriptID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Management Script ID
func (id DeviceManagementDeviceManagementScriptId) ID() string {
	fmtString := "/deviceManagement/deviceManagementScripts/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementScriptId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Management Script ID
func (id DeviceManagementDeviceManagementScriptId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceManagementScripts", "deviceManagementScripts", "deviceManagementScripts"),
		resourceids.UserSpecifiedSegment("deviceManagementScriptId", "deviceManagementScriptId"),
	}
}

// String returns a human-readable description of this Device Management Device Management Script ID
func (id DeviceManagementDeviceManagementScriptId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Script: %q", id.DeviceManagementScriptId),
	}
	return fmt.Sprintf("Device Management Device Management Script (%s)", strings.Join(components, "\n"))
}
