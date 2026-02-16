package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementHardwareConfigurationId{}

// DeviceManagementHardwareConfigurationId is a struct representing the Resource ID for a Device Management Hardware Configuration
type DeviceManagementHardwareConfigurationId struct {
	HardwareConfigurationId string
}

// NewDeviceManagementHardwareConfigurationID returns a new DeviceManagementHardwareConfigurationId struct
func NewDeviceManagementHardwareConfigurationID(hardwareConfigurationId string) DeviceManagementHardwareConfigurationId {
	return DeviceManagementHardwareConfigurationId{
		HardwareConfigurationId: hardwareConfigurationId,
	}
}

// ParseDeviceManagementHardwareConfigurationID parses 'input' into a DeviceManagementHardwareConfigurationId
func ParseDeviceManagementHardwareConfigurationID(input string) (*DeviceManagementHardwareConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementHardwareConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementHardwareConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementHardwareConfigurationIDInsensitively parses 'input' case-insensitively into a DeviceManagementHardwareConfigurationId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementHardwareConfigurationIDInsensitively(input string) (*DeviceManagementHardwareConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementHardwareConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementHardwareConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementHardwareConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.HardwareConfigurationId, ok = input.Parsed["hardwareConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "hardwareConfigurationId", input)
	}

	return nil
}

// ValidateDeviceManagementHardwareConfigurationID checks that 'input' can be parsed as a Device Management Hardware Configuration ID
func ValidateDeviceManagementHardwareConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementHardwareConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Hardware Configuration ID
func (id DeviceManagementHardwareConfigurationId) ID() string {
	fmtString := "/deviceManagement/hardwareConfigurations/%s"
	return fmt.Sprintf(fmtString, id.HardwareConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Hardware Configuration ID
func (id DeviceManagementHardwareConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("hardwareConfigurations", "hardwareConfigurations", "hardwareConfigurations"),
		resourceids.UserSpecifiedSegment("hardwareConfigurationId", "hardwareConfigurationId"),
	}
}

// String returns a human-readable description of this Device Management Hardware Configuration ID
func (id DeviceManagementHardwareConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Hardware Configuration: %q", id.HardwareConfigurationId),
	}
	return fmt.Sprintf("Device Management Hardware Configuration (%s)", strings.Join(components, "\n"))
}
