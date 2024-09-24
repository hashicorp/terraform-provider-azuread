package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementHardwareConfigurationIdUserRunStateId{}

// DeviceManagementHardwareConfigurationIdUserRunStateId is a struct representing the Resource ID for a Device Management Hardware Configuration Id User Run State
type DeviceManagementHardwareConfigurationIdUserRunStateId struct {
	HardwareConfigurationId          string
	HardwareConfigurationUserStateId string
}

// NewDeviceManagementHardwareConfigurationIdUserRunStateID returns a new DeviceManagementHardwareConfigurationIdUserRunStateId struct
func NewDeviceManagementHardwareConfigurationIdUserRunStateID(hardwareConfigurationId string, hardwareConfigurationUserStateId string) DeviceManagementHardwareConfigurationIdUserRunStateId {
	return DeviceManagementHardwareConfigurationIdUserRunStateId{
		HardwareConfigurationId:          hardwareConfigurationId,
		HardwareConfigurationUserStateId: hardwareConfigurationUserStateId,
	}
}

// ParseDeviceManagementHardwareConfigurationIdUserRunStateID parses 'input' into a DeviceManagementHardwareConfigurationIdUserRunStateId
func ParseDeviceManagementHardwareConfigurationIdUserRunStateID(input string) (*DeviceManagementHardwareConfigurationIdUserRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementHardwareConfigurationIdUserRunStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementHardwareConfigurationIdUserRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementHardwareConfigurationIdUserRunStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementHardwareConfigurationIdUserRunStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementHardwareConfigurationIdUserRunStateIDInsensitively(input string) (*DeviceManagementHardwareConfigurationIdUserRunStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementHardwareConfigurationIdUserRunStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementHardwareConfigurationIdUserRunStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementHardwareConfigurationIdUserRunStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.HardwareConfigurationId, ok = input.Parsed["hardwareConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "hardwareConfigurationId", input)
	}

	if id.HardwareConfigurationUserStateId, ok = input.Parsed["hardwareConfigurationUserStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "hardwareConfigurationUserStateId", input)
	}

	return nil
}

// ValidateDeviceManagementHardwareConfigurationIdUserRunStateID checks that 'input' can be parsed as a Device Management Hardware Configuration Id User Run State ID
func ValidateDeviceManagementHardwareConfigurationIdUserRunStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementHardwareConfigurationIdUserRunStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Hardware Configuration Id User Run State ID
func (id DeviceManagementHardwareConfigurationIdUserRunStateId) ID() string {
	fmtString := "/deviceManagement/hardwareConfigurations/%s/userRunStates/%s"
	return fmt.Sprintf(fmtString, id.HardwareConfigurationId, id.HardwareConfigurationUserStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Hardware Configuration Id User Run State ID
func (id DeviceManagementHardwareConfigurationIdUserRunStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("hardwareConfigurations", "hardwareConfigurations", "hardwareConfigurations"),
		resourceids.UserSpecifiedSegment("hardwareConfigurationId", "hardwareConfigurationId"),
		resourceids.StaticSegment("userRunStates", "userRunStates", "userRunStates"),
		resourceids.UserSpecifiedSegment("hardwareConfigurationUserStateId", "hardwareConfigurationUserStateId"),
	}
}

// String returns a human-readable description of this Device Management Hardware Configuration Id User Run State ID
func (id DeviceManagementHardwareConfigurationIdUserRunStateId) String() string {
	components := []string{
		fmt.Sprintf("Hardware Configuration: %q", id.HardwareConfigurationId),
		fmt.Sprintf("Hardware Configuration User State: %q", id.HardwareConfigurationUserStateId),
	}
	return fmt.Sprintf("Device Management Hardware Configuration Id User Run State (%s)", strings.Join(components, "\n"))
}
