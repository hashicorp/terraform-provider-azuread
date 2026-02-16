package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceEnrollmentConfigurationId{}

// DeviceManagementDeviceEnrollmentConfigurationId is a struct representing the Resource ID for a Device Management Device Enrollment Configuration
type DeviceManagementDeviceEnrollmentConfigurationId struct {
	DeviceEnrollmentConfigurationId string
}

// NewDeviceManagementDeviceEnrollmentConfigurationID returns a new DeviceManagementDeviceEnrollmentConfigurationId struct
func NewDeviceManagementDeviceEnrollmentConfigurationID(deviceEnrollmentConfigurationId string) DeviceManagementDeviceEnrollmentConfigurationId {
	return DeviceManagementDeviceEnrollmentConfigurationId{
		DeviceEnrollmentConfigurationId: deviceEnrollmentConfigurationId,
	}
}

// ParseDeviceManagementDeviceEnrollmentConfigurationID parses 'input' into a DeviceManagementDeviceEnrollmentConfigurationId
func ParseDeviceManagementDeviceEnrollmentConfigurationID(input string) (*DeviceManagementDeviceEnrollmentConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceEnrollmentConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceEnrollmentConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceEnrollmentConfigurationIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceEnrollmentConfigurationId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceEnrollmentConfigurationIDInsensitively(input string) (*DeviceManagementDeviceEnrollmentConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceEnrollmentConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceEnrollmentConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceEnrollmentConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceEnrollmentConfigurationId, ok = input.Parsed["deviceEnrollmentConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceEnrollmentConfigurationId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceEnrollmentConfigurationID checks that 'input' can be parsed as a Device Management Device Enrollment Configuration ID
func ValidateDeviceManagementDeviceEnrollmentConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceEnrollmentConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Enrollment Configuration ID
func (id DeviceManagementDeviceEnrollmentConfigurationId) ID() string {
	fmtString := "/deviceManagement/deviceEnrollmentConfigurations/%s"
	return fmt.Sprintf(fmtString, id.DeviceEnrollmentConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Enrollment Configuration ID
func (id DeviceManagementDeviceEnrollmentConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceEnrollmentConfigurations", "deviceEnrollmentConfigurations", "deviceEnrollmentConfigurations"),
		resourceids.UserSpecifiedSegment("deviceEnrollmentConfigurationId", "deviceEnrollmentConfigurationId"),
	}
}

// String returns a human-readable description of this Device Management Device Enrollment Configuration ID
func (id DeviceManagementDeviceEnrollmentConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Device Enrollment Configuration: %q", id.DeviceEnrollmentConfigurationId),
	}
	return fmt.Sprintf("Device Management Device Enrollment Configuration (%s)", strings.Join(components, "\n"))
}
