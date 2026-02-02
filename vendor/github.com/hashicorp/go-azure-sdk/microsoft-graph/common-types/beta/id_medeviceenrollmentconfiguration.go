package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDeviceEnrollmentConfigurationId{}

// MeDeviceEnrollmentConfigurationId is a struct representing the Resource ID for a Me Device Enrollment Configuration
type MeDeviceEnrollmentConfigurationId struct {
	DeviceEnrollmentConfigurationId string
}

// NewMeDeviceEnrollmentConfigurationID returns a new MeDeviceEnrollmentConfigurationId struct
func NewMeDeviceEnrollmentConfigurationID(deviceEnrollmentConfigurationId string) MeDeviceEnrollmentConfigurationId {
	return MeDeviceEnrollmentConfigurationId{
		DeviceEnrollmentConfigurationId: deviceEnrollmentConfigurationId,
	}
}

// ParseMeDeviceEnrollmentConfigurationID parses 'input' into a MeDeviceEnrollmentConfigurationId
func ParseMeDeviceEnrollmentConfigurationID(input string) (*MeDeviceEnrollmentConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDeviceEnrollmentConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDeviceEnrollmentConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDeviceEnrollmentConfigurationIDInsensitively parses 'input' case-insensitively into a MeDeviceEnrollmentConfigurationId
// note: this method should only be used for API response data and not user input
func ParseMeDeviceEnrollmentConfigurationIDInsensitively(input string) (*MeDeviceEnrollmentConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDeviceEnrollmentConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDeviceEnrollmentConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDeviceEnrollmentConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceEnrollmentConfigurationId, ok = input.Parsed["deviceEnrollmentConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceEnrollmentConfigurationId", input)
	}

	return nil
}

// ValidateMeDeviceEnrollmentConfigurationID checks that 'input' can be parsed as a Me Device Enrollment Configuration ID
func ValidateMeDeviceEnrollmentConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDeviceEnrollmentConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Device Enrollment Configuration ID
func (id MeDeviceEnrollmentConfigurationId) ID() string {
	fmtString := "/me/deviceEnrollmentConfigurations/%s"
	return fmt.Sprintf(fmtString, id.DeviceEnrollmentConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Device Enrollment Configuration ID
func (id MeDeviceEnrollmentConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("deviceEnrollmentConfigurations", "deviceEnrollmentConfigurations", "deviceEnrollmentConfigurations"),
		resourceids.UserSpecifiedSegment("deviceEnrollmentConfigurationId", "deviceEnrollmentConfigurationId"),
	}
}

// String returns a human-readable description of this Me Device Enrollment Configuration ID
func (id MeDeviceEnrollmentConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Device Enrollment Configuration: %q", id.DeviceEnrollmentConfigurationId),
	}
	return fmt.Sprintf("Me Device Enrollment Configuration (%s)", strings.Join(components, "\n"))
}
