package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceEnrollmentConfigurationIdAssignmentId{}

// DeviceManagementDeviceEnrollmentConfigurationIdAssignmentId is a struct representing the Resource ID for a Device Management Device Enrollment Configuration Id Assignment
type DeviceManagementDeviceEnrollmentConfigurationIdAssignmentId struct {
	DeviceEnrollmentConfigurationId     string
	EnrollmentConfigurationAssignmentId string
}

// NewDeviceManagementDeviceEnrollmentConfigurationIdAssignmentID returns a new DeviceManagementDeviceEnrollmentConfigurationIdAssignmentId struct
func NewDeviceManagementDeviceEnrollmentConfigurationIdAssignmentID(deviceEnrollmentConfigurationId string, enrollmentConfigurationAssignmentId string) DeviceManagementDeviceEnrollmentConfigurationIdAssignmentId {
	return DeviceManagementDeviceEnrollmentConfigurationIdAssignmentId{
		DeviceEnrollmentConfigurationId:     deviceEnrollmentConfigurationId,
		EnrollmentConfigurationAssignmentId: enrollmentConfigurationAssignmentId,
	}
}

// ParseDeviceManagementDeviceEnrollmentConfigurationIdAssignmentID parses 'input' into a DeviceManagementDeviceEnrollmentConfigurationIdAssignmentId
func ParseDeviceManagementDeviceEnrollmentConfigurationIdAssignmentID(input string) (*DeviceManagementDeviceEnrollmentConfigurationIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceEnrollmentConfigurationIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceEnrollmentConfigurationIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceEnrollmentConfigurationIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceEnrollmentConfigurationIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceEnrollmentConfigurationIdAssignmentIDInsensitively(input string) (*DeviceManagementDeviceEnrollmentConfigurationIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceEnrollmentConfigurationIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceEnrollmentConfigurationIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceEnrollmentConfigurationIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceEnrollmentConfigurationId, ok = input.Parsed["deviceEnrollmentConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceEnrollmentConfigurationId", input)
	}

	if id.EnrollmentConfigurationAssignmentId, ok = input.Parsed["enrollmentConfigurationAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "enrollmentConfigurationAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceEnrollmentConfigurationIdAssignmentID checks that 'input' can be parsed as a Device Management Device Enrollment Configuration Id Assignment ID
func ValidateDeviceManagementDeviceEnrollmentConfigurationIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceEnrollmentConfigurationIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Enrollment Configuration Id Assignment ID
func (id DeviceManagementDeviceEnrollmentConfigurationIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/deviceEnrollmentConfigurations/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.DeviceEnrollmentConfigurationId, id.EnrollmentConfigurationAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Enrollment Configuration Id Assignment ID
func (id DeviceManagementDeviceEnrollmentConfigurationIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceEnrollmentConfigurations", "deviceEnrollmentConfigurations", "deviceEnrollmentConfigurations"),
		resourceids.UserSpecifiedSegment("deviceEnrollmentConfigurationId", "deviceEnrollmentConfigurationId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("enrollmentConfigurationAssignmentId", "enrollmentConfigurationAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Device Enrollment Configuration Id Assignment ID
func (id DeviceManagementDeviceEnrollmentConfigurationIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Device Enrollment Configuration: %q", id.DeviceEnrollmentConfigurationId),
		fmt.Sprintf("Enrollment Configuration Assignment: %q", id.EnrollmentConfigurationAssignmentId),
	}
	return fmt.Sprintf("Device Management Device Enrollment Configuration Id Assignment (%s)", strings.Join(components, "\n"))
}
