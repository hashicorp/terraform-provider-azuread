package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementHardwareConfigurationIdDeviceRunStateId{}

// DeviceManagementHardwareConfigurationIdDeviceRunStateId is a struct representing the Resource ID for a Device Management Hardware Configuration Id Device Run State
type DeviceManagementHardwareConfigurationIdDeviceRunStateId struct {
	HardwareConfigurationId            string
	HardwareConfigurationDeviceStateId string
}

// NewDeviceManagementHardwareConfigurationIdDeviceRunStateID returns a new DeviceManagementHardwareConfigurationIdDeviceRunStateId struct
func NewDeviceManagementHardwareConfigurationIdDeviceRunStateID(hardwareConfigurationId string, hardwareConfigurationDeviceStateId string) DeviceManagementHardwareConfigurationIdDeviceRunStateId {
	return DeviceManagementHardwareConfigurationIdDeviceRunStateId{
		HardwareConfigurationId:            hardwareConfigurationId,
		HardwareConfigurationDeviceStateId: hardwareConfigurationDeviceStateId,
	}
}

// ParseDeviceManagementHardwareConfigurationIdDeviceRunStateID parses 'input' into a DeviceManagementHardwareConfigurationIdDeviceRunStateId
func ParseDeviceManagementHardwareConfigurationIdDeviceRunStateID(input string) (*DeviceManagementHardwareConfigurationIdDeviceRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementHardwareConfigurationIdDeviceRunStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementHardwareConfigurationIdDeviceRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementHardwareConfigurationIdDeviceRunStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementHardwareConfigurationIdDeviceRunStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementHardwareConfigurationIdDeviceRunStateIDInsensitively(input string) (*DeviceManagementHardwareConfigurationIdDeviceRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementHardwareConfigurationIdDeviceRunStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementHardwareConfigurationIdDeviceRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementHardwareConfigurationIdDeviceRunStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.HardwareConfigurationId, ok = input.Parsed["hardwareConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "hardwareConfigurationId", input)
	}

	if id.HardwareConfigurationDeviceStateId, ok = input.Parsed["hardwareConfigurationDeviceStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "hardwareConfigurationDeviceStateId", input)
	}

	return nil
}

// ValidateDeviceManagementHardwareConfigurationIdDeviceRunStateID checks that 'input' can be parsed as a Device Management Hardware Configuration Id Device Run State ID
func ValidateDeviceManagementHardwareConfigurationIdDeviceRunStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementHardwareConfigurationIdDeviceRunStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Hardware Configuration Id Device Run State ID
func (id DeviceManagementHardwareConfigurationIdDeviceRunStateId) ID() string {
	fmtString := "/deviceManagement/hardwareConfigurations/%s/deviceRunStates/%s"
	return fmt.Sprintf(fmtString, id.HardwareConfigurationId, id.HardwareConfigurationDeviceStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Hardware Configuration Id Device Run State ID
func (id DeviceManagementHardwareConfigurationIdDeviceRunStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("hardwareConfigurations", "hardwareConfigurations", "hardwareConfigurations"),
		resourceids.UserSpecifiedSegment("hardwareConfigurationId", "hardwareConfigurationId"),
		resourceids.StaticSegment("deviceRunStates", "deviceRunStates", "deviceRunStates"),
		resourceids.UserSpecifiedSegment("hardwareConfigurationDeviceStateId", "hardwareConfigurationDeviceStateId"),
	}
}

// String returns a human-readable description of this Device Management Hardware Configuration Id Device Run State ID
func (id DeviceManagementHardwareConfigurationIdDeviceRunStateId) String() string {
	components := []string{
		fmt.Sprintf("Hardware Configuration: %q", id.HardwareConfigurationId),
		fmt.Sprintf("Hardware Configuration Device State: %q", id.HardwareConfigurationDeviceStateId),
	}
	return fmt.Sprintf("Device Management Hardware Configuration Id Device Run State (%s)", strings.Join(components, "\n"))
}
