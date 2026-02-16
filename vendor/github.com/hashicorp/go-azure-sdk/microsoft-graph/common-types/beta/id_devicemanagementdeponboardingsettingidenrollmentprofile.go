package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDepOnboardingSettingIdEnrollmentProfileId{}

// DeviceManagementDepOnboardingSettingIdEnrollmentProfileId is a struct representing the Resource ID for a Device Management Dep Onboarding Setting Id Enrollment Profile
type DeviceManagementDepOnboardingSettingIdEnrollmentProfileId struct {
	DepOnboardingSettingId string
	EnrollmentProfileId    string
}

// NewDeviceManagementDepOnboardingSettingIdEnrollmentProfileID returns a new DeviceManagementDepOnboardingSettingIdEnrollmentProfileId struct
func NewDeviceManagementDepOnboardingSettingIdEnrollmentProfileID(depOnboardingSettingId string, enrollmentProfileId string) DeviceManagementDepOnboardingSettingIdEnrollmentProfileId {
	return DeviceManagementDepOnboardingSettingIdEnrollmentProfileId{
		DepOnboardingSettingId: depOnboardingSettingId,
		EnrollmentProfileId:    enrollmentProfileId,
	}
}

// ParseDeviceManagementDepOnboardingSettingIdEnrollmentProfileID parses 'input' into a DeviceManagementDepOnboardingSettingIdEnrollmentProfileId
func ParseDeviceManagementDepOnboardingSettingIdEnrollmentProfileID(input string) (*DeviceManagementDepOnboardingSettingIdEnrollmentProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDepOnboardingSettingIdEnrollmentProfileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDepOnboardingSettingIdEnrollmentProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDepOnboardingSettingIdEnrollmentProfileIDInsensitively parses 'input' case-insensitively into a DeviceManagementDepOnboardingSettingIdEnrollmentProfileId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDepOnboardingSettingIdEnrollmentProfileIDInsensitively(input string) (*DeviceManagementDepOnboardingSettingIdEnrollmentProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDepOnboardingSettingIdEnrollmentProfileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDepOnboardingSettingIdEnrollmentProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDepOnboardingSettingIdEnrollmentProfileId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DepOnboardingSettingId, ok = input.Parsed["depOnboardingSettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "depOnboardingSettingId", input)
	}

	if id.EnrollmentProfileId, ok = input.Parsed["enrollmentProfileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "enrollmentProfileId", input)
	}

	return nil
}

// ValidateDeviceManagementDepOnboardingSettingIdEnrollmentProfileID checks that 'input' can be parsed as a Device Management Dep Onboarding Setting Id Enrollment Profile ID
func ValidateDeviceManagementDepOnboardingSettingIdEnrollmentProfileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDepOnboardingSettingIdEnrollmentProfileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Dep Onboarding Setting Id Enrollment Profile ID
func (id DeviceManagementDepOnboardingSettingIdEnrollmentProfileId) ID() string {
	fmtString := "/deviceManagement/depOnboardingSettings/%s/enrollmentProfiles/%s"
	return fmt.Sprintf(fmtString, id.DepOnboardingSettingId, id.EnrollmentProfileId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Dep Onboarding Setting Id Enrollment Profile ID
func (id DeviceManagementDepOnboardingSettingIdEnrollmentProfileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("depOnboardingSettings", "depOnboardingSettings", "depOnboardingSettings"),
		resourceids.UserSpecifiedSegment("depOnboardingSettingId", "depOnboardingSettingId"),
		resourceids.StaticSegment("enrollmentProfiles", "enrollmentProfiles", "enrollmentProfiles"),
		resourceids.UserSpecifiedSegment("enrollmentProfileId", "enrollmentProfileId"),
	}
}

// String returns a human-readable description of this Device Management Dep Onboarding Setting Id Enrollment Profile ID
func (id DeviceManagementDepOnboardingSettingIdEnrollmentProfileId) String() string {
	components := []string{
		fmt.Sprintf("Dep Onboarding Setting: %q", id.DepOnboardingSettingId),
		fmt.Sprintf("Enrollment Profile: %q", id.EnrollmentProfileId),
	}
	return fmt.Sprintf("Device Management Dep Onboarding Setting Id Enrollment Profile (%s)", strings.Join(components, "\n"))
}
