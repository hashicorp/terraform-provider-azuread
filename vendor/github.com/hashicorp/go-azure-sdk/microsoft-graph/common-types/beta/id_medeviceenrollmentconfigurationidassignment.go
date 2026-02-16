package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDeviceEnrollmentConfigurationIdAssignmentId{}

// MeDeviceEnrollmentConfigurationIdAssignmentId is a struct representing the Resource ID for a Me Device Enrollment Configuration Id Assignment
type MeDeviceEnrollmentConfigurationIdAssignmentId struct {
	DeviceEnrollmentConfigurationId     string
	EnrollmentConfigurationAssignmentId string
}

// NewMeDeviceEnrollmentConfigurationIdAssignmentID returns a new MeDeviceEnrollmentConfigurationIdAssignmentId struct
func NewMeDeviceEnrollmentConfigurationIdAssignmentID(deviceEnrollmentConfigurationId string, enrollmentConfigurationAssignmentId string) MeDeviceEnrollmentConfigurationIdAssignmentId {
	return MeDeviceEnrollmentConfigurationIdAssignmentId{
		DeviceEnrollmentConfigurationId:     deviceEnrollmentConfigurationId,
		EnrollmentConfigurationAssignmentId: enrollmentConfigurationAssignmentId,
	}
}

// ParseMeDeviceEnrollmentConfigurationIdAssignmentID parses 'input' into a MeDeviceEnrollmentConfigurationIdAssignmentId
func ParseMeDeviceEnrollmentConfigurationIdAssignmentID(input string) (*MeDeviceEnrollmentConfigurationIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDeviceEnrollmentConfigurationIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDeviceEnrollmentConfigurationIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDeviceEnrollmentConfigurationIdAssignmentIDInsensitively parses 'input' case-insensitively into a MeDeviceEnrollmentConfigurationIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseMeDeviceEnrollmentConfigurationIdAssignmentIDInsensitively(input string) (*MeDeviceEnrollmentConfigurationIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDeviceEnrollmentConfigurationIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDeviceEnrollmentConfigurationIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDeviceEnrollmentConfigurationIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceEnrollmentConfigurationId, ok = input.Parsed["deviceEnrollmentConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceEnrollmentConfigurationId", input)
	}

	if id.EnrollmentConfigurationAssignmentId, ok = input.Parsed["enrollmentConfigurationAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "enrollmentConfigurationAssignmentId", input)
	}

	return nil
}

// ValidateMeDeviceEnrollmentConfigurationIdAssignmentID checks that 'input' can be parsed as a Me Device Enrollment Configuration Id Assignment ID
func ValidateMeDeviceEnrollmentConfigurationIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDeviceEnrollmentConfigurationIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Device Enrollment Configuration Id Assignment ID
func (id MeDeviceEnrollmentConfigurationIdAssignmentId) ID() string {
	fmtString := "/me/deviceEnrollmentConfigurations/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.DeviceEnrollmentConfigurationId, id.EnrollmentConfigurationAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Device Enrollment Configuration Id Assignment ID
func (id MeDeviceEnrollmentConfigurationIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("deviceEnrollmentConfigurations", "deviceEnrollmentConfigurations", "deviceEnrollmentConfigurations"),
		resourceids.UserSpecifiedSegment("deviceEnrollmentConfigurationId", "deviceEnrollmentConfigurationId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("enrollmentConfigurationAssignmentId", "enrollmentConfigurationAssignmentId"),
	}
}

// String returns a human-readable description of this Me Device Enrollment Configuration Id Assignment ID
func (id MeDeviceEnrollmentConfigurationIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Device Enrollment Configuration: %q", id.DeviceEnrollmentConfigurationId),
		fmt.Sprintf("Enrollment Configuration Assignment: %q", id.EnrollmentConfigurationAssignmentId),
	}
	return fmt.Sprintf("Me Device Enrollment Configuration Id Assignment (%s)", strings.Join(components, "\n"))
}
