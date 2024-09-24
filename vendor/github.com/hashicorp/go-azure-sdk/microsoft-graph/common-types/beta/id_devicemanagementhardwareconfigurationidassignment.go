package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementHardwareConfigurationIdAssignmentId{}

// DeviceManagementHardwareConfigurationIdAssignmentId is a struct representing the Resource ID for a Device Management Hardware Configuration Id Assignment
type DeviceManagementHardwareConfigurationIdAssignmentId struct {
	HardwareConfigurationId           string
	HardwareConfigurationAssignmentId string
}

// NewDeviceManagementHardwareConfigurationIdAssignmentID returns a new DeviceManagementHardwareConfigurationIdAssignmentId struct
func NewDeviceManagementHardwareConfigurationIdAssignmentID(hardwareConfigurationId string, hardwareConfigurationAssignmentId string) DeviceManagementHardwareConfigurationIdAssignmentId {
	return DeviceManagementHardwareConfigurationIdAssignmentId{
		HardwareConfigurationId:           hardwareConfigurationId,
		HardwareConfigurationAssignmentId: hardwareConfigurationAssignmentId,
	}
}

// ParseDeviceManagementHardwareConfigurationIdAssignmentID parses 'input' into a DeviceManagementHardwareConfigurationIdAssignmentId
func ParseDeviceManagementHardwareConfigurationIdAssignmentID(input string) (*DeviceManagementHardwareConfigurationIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementHardwareConfigurationIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementHardwareConfigurationIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementHardwareConfigurationIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementHardwareConfigurationIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementHardwareConfigurationIdAssignmentIDInsensitively(input string) (*DeviceManagementHardwareConfigurationIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementHardwareConfigurationIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementHardwareConfigurationIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementHardwareConfigurationIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.HardwareConfigurationId, ok = input.Parsed["hardwareConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "hardwareConfigurationId", input)
	}

	if id.HardwareConfigurationAssignmentId, ok = input.Parsed["hardwareConfigurationAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "hardwareConfigurationAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementHardwareConfigurationIdAssignmentID checks that 'input' can be parsed as a Device Management Hardware Configuration Id Assignment ID
func ValidateDeviceManagementHardwareConfigurationIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementHardwareConfigurationIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Hardware Configuration Id Assignment ID
func (id DeviceManagementHardwareConfigurationIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/hardwareConfigurations/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.HardwareConfigurationId, id.HardwareConfigurationAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Hardware Configuration Id Assignment ID
func (id DeviceManagementHardwareConfigurationIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("hardwareConfigurations", "hardwareConfigurations", "hardwareConfigurations"),
		resourceids.UserSpecifiedSegment("hardwareConfigurationId", "hardwareConfigurationId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("hardwareConfigurationAssignmentId", "hardwareConfigurationAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Hardware Configuration Id Assignment ID
func (id DeviceManagementHardwareConfigurationIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Hardware Configuration: %q", id.HardwareConfigurationId),
		fmt.Sprintf("Hardware Configuration Assignment: %q", id.HardwareConfigurationAssignmentId),
	}
	return fmt.Sprintf("Device Management Hardware Configuration Id Assignment (%s)", strings.Join(components, "\n"))
}
