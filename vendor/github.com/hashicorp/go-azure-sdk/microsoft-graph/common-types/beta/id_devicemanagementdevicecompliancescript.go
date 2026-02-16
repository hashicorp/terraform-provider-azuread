package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceComplianceScriptId{}

// DeviceManagementDeviceComplianceScriptId is a struct representing the Resource ID for a Device Management Device Compliance Script
type DeviceManagementDeviceComplianceScriptId struct {
	DeviceComplianceScriptId string
}

// NewDeviceManagementDeviceComplianceScriptID returns a new DeviceManagementDeviceComplianceScriptId struct
func NewDeviceManagementDeviceComplianceScriptID(deviceComplianceScriptId string) DeviceManagementDeviceComplianceScriptId {
	return DeviceManagementDeviceComplianceScriptId{
		DeviceComplianceScriptId: deviceComplianceScriptId,
	}
}

// ParseDeviceManagementDeviceComplianceScriptID parses 'input' into a DeviceManagementDeviceComplianceScriptId
func ParseDeviceManagementDeviceComplianceScriptID(input string) (*DeviceManagementDeviceComplianceScriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceComplianceScriptId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceComplianceScriptId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceComplianceScriptIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceComplianceScriptId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceComplianceScriptIDInsensitively(input string) (*DeviceManagementDeviceComplianceScriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceComplianceScriptId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceComplianceScriptId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceComplianceScriptId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceComplianceScriptId, ok = input.Parsed["deviceComplianceScriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceComplianceScriptId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceComplianceScriptID checks that 'input' can be parsed as a Device Management Device Compliance Script ID
func ValidateDeviceManagementDeviceComplianceScriptID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceComplianceScriptID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Compliance Script ID
func (id DeviceManagementDeviceComplianceScriptId) ID() string {
	fmtString := "/deviceManagement/deviceComplianceScripts/%s"
	return fmt.Sprintf(fmtString, id.DeviceComplianceScriptId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Compliance Script ID
func (id DeviceManagementDeviceComplianceScriptId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceComplianceScripts", "deviceComplianceScripts", "deviceComplianceScripts"),
		resourceids.UserSpecifiedSegment("deviceComplianceScriptId", "deviceComplianceScriptId"),
	}
}

// String returns a human-readable description of this Device Management Device Compliance Script ID
func (id DeviceManagementDeviceComplianceScriptId) String() string {
	components := []string{
		fmt.Sprintf("Device Compliance Script: %q", id.DeviceComplianceScriptId),
	}
	return fmt.Sprintf("Device Management Device Compliance Script (%s)", strings.Join(components, "\n"))
}
