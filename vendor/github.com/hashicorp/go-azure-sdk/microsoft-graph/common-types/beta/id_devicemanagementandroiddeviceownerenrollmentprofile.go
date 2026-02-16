package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementAndroidDeviceOwnerEnrollmentProfileId{}

// DeviceManagementAndroidDeviceOwnerEnrollmentProfileId is a struct representing the Resource ID for a Device Management Android Device Owner Enrollment Profile
type DeviceManagementAndroidDeviceOwnerEnrollmentProfileId struct {
	AndroidDeviceOwnerEnrollmentProfileId string
}

// NewDeviceManagementAndroidDeviceOwnerEnrollmentProfileID returns a new DeviceManagementAndroidDeviceOwnerEnrollmentProfileId struct
func NewDeviceManagementAndroidDeviceOwnerEnrollmentProfileID(androidDeviceOwnerEnrollmentProfileId string) DeviceManagementAndroidDeviceOwnerEnrollmentProfileId {
	return DeviceManagementAndroidDeviceOwnerEnrollmentProfileId{
		AndroidDeviceOwnerEnrollmentProfileId: androidDeviceOwnerEnrollmentProfileId,
	}
}

// ParseDeviceManagementAndroidDeviceOwnerEnrollmentProfileID parses 'input' into a DeviceManagementAndroidDeviceOwnerEnrollmentProfileId
func ParseDeviceManagementAndroidDeviceOwnerEnrollmentProfileID(input string) (*DeviceManagementAndroidDeviceOwnerEnrollmentProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementAndroidDeviceOwnerEnrollmentProfileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementAndroidDeviceOwnerEnrollmentProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementAndroidDeviceOwnerEnrollmentProfileIDInsensitively parses 'input' case-insensitively into a DeviceManagementAndroidDeviceOwnerEnrollmentProfileId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementAndroidDeviceOwnerEnrollmentProfileIDInsensitively(input string) (*DeviceManagementAndroidDeviceOwnerEnrollmentProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementAndroidDeviceOwnerEnrollmentProfileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementAndroidDeviceOwnerEnrollmentProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementAndroidDeviceOwnerEnrollmentProfileId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AndroidDeviceOwnerEnrollmentProfileId, ok = input.Parsed["androidDeviceOwnerEnrollmentProfileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "androidDeviceOwnerEnrollmentProfileId", input)
	}

	return nil
}

// ValidateDeviceManagementAndroidDeviceOwnerEnrollmentProfileID checks that 'input' can be parsed as a Device Management Android Device Owner Enrollment Profile ID
func ValidateDeviceManagementAndroidDeviceOwnerEnrollmentProfileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementAndroidDeviceOwnerEnrollmentProfileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Android Device Owner Enrollment Profile ID
func (id DeviceManagementAndroidDeviceOwnerEnrollmentProfileId) ID() string {
	fmtString := "/deviceManagement/androidDeviceOwnerEnrollmentProfiles/%s"
	return fmt.Sprintf(fmtString, id.AndroidDeviceOwnerEnrollmentProfileId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Android Device Owner Enrollment Profile ID
func (id DeviceManagementAndroidDeviceOwnerEnrollmentProfileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("androidDeviceOwnerEnrollmentProfiles", "androidDeviceOwnerEnrollmentProfiles", "androidDeviceOwnerEnrollmentProfiles"),
		resourceids.UserSpecifiedSegment("androidDeviceOwnerEnrollmentProfileId", "androidDeviceOwnerEnrollmentProfileId"),
	}
}

// String returns a human-readable description of this Device Management Android Device Owner Enrollment Profile ID
func (id DeviceManagementAndroidDeviceOwnerEnrollmentProfileId) String() string {
	components := []string{
		fmt.Sprintf("Android Device Owner Enrollment Profile: %q", id.AndroidDeviceOwnerEnrollmentProfileId),
	}
	return fmt.Sprintf("Device Management Android Device Owner Enrollment Profile (%s)", strings.Join(components, "\n"))
}
