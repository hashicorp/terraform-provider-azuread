package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementAppleUserInitiatedEnrollmentProfileId{}

// DeviceManagementAppleUserInitiatedEnrollmentProfileId is a struct representing the Resource ID for a Device Management Apple User Initiated Enrollment Profile
type DeviceManagementAppleUserInitiatedEnrollmentProfileId struct {
	AppleUserInitiatedEnrollmentProfileId string
}

// NewDeviceManagementAppleUserInitiatedEnrollmentProfileID returns a new DeviceManagementAppleUserInitiatedEnrollmentProfileId struct
func NewDeviceManagementAppleUserInitiatedEnrollmentProfileID(appleUserInitiatedEnrollmentProfileId string) DeviceManagementAppleUserInitiatedEnrollmentProfileId {
	return DeviceManagementAppleUserInitiatedEnrollmentProfileId{
		AppleUserInitiatedEnrollmentProfileId: appleUserInitiatedEnrollmentProfileId,
	}
}

// ParseDeviceManagementAppleUserInitiatedEnrollmentProfileID parses 'input' into a DeviceManagementAppleUserInitiatedEnrollmentProfileId
func ParseDeviceManagementAppleUserInitiatedEnrollmentProfileID(input string) (*DeviceManagementAppleUserInitiatedEnrollmentProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementAppleUserInitiatedEnrollmentProfileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementAppleUserInitiatedEnrollmentProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementAppleUserInitiatedEnrollmentProfileIDInsensitively parses 'input' case-insensitively into a DeviceManagementAppleUserInitiatedEnrollmentProfileId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementAppleUserInitiatedEnrollmentProfileIDInsensitively(input string) (*DeviceManagementAppleUserInitiatedEnrollmentProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementAppleUserInitiatedEnrollmentProfileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementAppleUserInitiatedEnrollmentProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementAppleUserInitiatedEnrollmentProfileId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AppleUserInitiatedEnrollmentProfileId, ok = input.Parsed["appleUserInitiatedEnrollmentProfileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appleUserInitiatedEnrollmentProfileId", input)
	}

	return nil
}

// ValidateDeviceManagementAppleUserInitiatedEnrollmentProfileID checks that 'input' can be parsed as a Device Management Apple User Initiated Enrollment Profile ID
func ValidateDeviceManagementAppleUserInitiatedEnrollmentProfileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementAppleUserInitiatedEnrollmentProfileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Apple User Initiated Enrollment Profile ID
func (id DeviceManagementAppleUserInitiatedEnrollmentProfileId) ID() string {
	fmtString := "/deviceManagement/appleUserInitiatedEnrollmentProfiles/%s"
	return fmt.Sprintf(fmtString, id.AppleUserInitiatedEnrollmentProfileId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Apple User Initiated Enrollment Profile ID
func (id DeviceManagementAppleUserInitiatedEnrollmentProfileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("appleUserInitiatedEnrollmentProfiles", "appleUserInitiatedEnrollmentProfiles", "appleUserInitiatedEnrollmentProfiles"),
		resourceids.UserSpecifiedSegment("appleUserInitiatedEnrollmentProfileId", "appleUserInitiatedEnrollmentProfileId"),
	}
}

// String returns a human-readable description of this Device Management Apple User Initiated Enrollment Profile ID
func (id DeviceManagementAppleUserInitiatedEnrollmentProfileId) String() string {
	components := []string{
		fmt.Sprintf("Apple User Initiated Enrollment Profile: %q", id.AppleUserInitiatedEnrollmentProfileId),
	}
	return fmt.Sprintf("Device Management Apple User Initiated Enrollment Profile (%s)", strings.Join(components, "\n"))
}
