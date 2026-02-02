package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementAndroidForWorkEnrollmentProfileId{}

// DeviceManagementAndroidForWorkEnrollmentProfileId is a struct representing the Resource ID for a Device Management Android For Work Enrollment Profile
type DeviceManagementAndroidForWorkEnrollmentProfileId struct {
	AndroidForWorkEnrollmentProfileId string
}

// NewDeviceManagementAndroidForWorkEnrollmentProfileID returns a new DeviceManagementAndroidForWorkEnrollmentProfileId struct
func NewDeviceManagementAndroidForWorkEnrollmentProfileID(androidForWorkEnrollmentProfileId string) DeviceManagementAndroidForWorkEnrollmentProfileId {
	return DeviceManagementAndroidForWorkEnrollmentProfileId{
		AndroidForWorkEnrollmentProfileId: androidForWorkEnrollmentProfileId,
	}
}

// ParseDeviceManagementAndroidForWorkEnrollmentProfileID parses 'input' into a DeviceManagementAndroidForWorkEnrollmentProfileId
func ParseDeviceManagementAndroidForWorkEnrollmentProfileID(input string) (*DeviceManagementAndroidForWorkEnrollmentProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementAndroidForWorkEnrollmentProfileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementAndroidForWorkEnrollmentProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementAndroidForWorkEnrollmentProfileIDInsensitively parses 'input' case-insensitively into a DeviceManagementAndroidForWorkEnrollmentProfileId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementAndroidForWorkEnrollmentProfileIDInsensitively(input string) (*DeviceManagementAndroidForWorkEnrollmentProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementAndroidForWorkEnrollmentProfileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementAndroidForWorkEnrollmentProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementAndroidForWorkEnrollmentProfileId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AndroidForWorkEnrollmentProfileId, ok = input.Parsed["androidForWorkEnrollmentProfileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "androidForWorkEnrollmentProfileId", input)
	}

	return nil
}

// ValidateDeviceManagementAndroidForWorkEnrollmentProfileID checks that 'input' can be parsed as a Device Management Android For Work Enrollment Profile ID
func ValidateDeviceManagementAndroidForWorkEnrollmentProfileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementAndroidForWorkEnrollmentProfileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Android For Work Enrollment Profile ID
func (id DeviceManagementAndroidForWorkEnrollmentProfileId) ID() string {
	fmtString := "/deviceManagement/androidForWorkEnrollmentProfiles/%s"
	return fmt.Sprintf(fmtString, id.AndroidForWorkEnrollmentProfileId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Android For Work Enrollment Profile ID
func (id DeviceManagementAndroidForWorkEnrollmentProfileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("androidForWorkEnrollmentProfiles", "androidForWorkEnrollmentProfiles", "androidForWorkEnrollmentProfiles"),
		resourceids.UserSpecifiedSegment("androidForWorkEnrollmentProfileId", "androidForWorkEnrollmentProfileId"),
	}
}

// String returns a human-readable description of this Device Management Android For Work Enrollment Profile ID
func (id DeviceManagementAndroidForWorkEnrollmentProfileId) String() string {
	components := []string{
		fmt.Sprintf("Android For Work Enrollment Profile: %q", id.AndroidForWorkEnrollmentProfileId),
	}
	return fmt.Sprintf("Device Management Android For Work Enrollment Profile (%s)", strings.Join(components, "\n"))
}
