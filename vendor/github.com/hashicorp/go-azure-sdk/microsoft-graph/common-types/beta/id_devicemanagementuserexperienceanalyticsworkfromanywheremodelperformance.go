package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId{}

// DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId is a struct representing the Resource ID for a Device Management User Experience Analytics Work From Anywhere Model Performance
type DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId struct {
	UserExperienceAnalyticsWorkFromAnywhereModelPerformanceId string
}

// NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceID returns a new DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId struct
func NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceID(userExperienceAnalyticsWorkFromAnywhereModelPerformanceId string) DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId {
	return DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId{
		UserExperienceAnalyticsWorkFromAnywhereModelPerformanceId: userExperienceAnalyticsWorkFromAnywhereModelPerformanceId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceID parses 'input' into a DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId
func ParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceID(input string) (*DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceId, ok = input.Parsed["userExperienceAnalyticsWorkFromAnywhereModelPerformanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsWorkFromAnywhereModelPerformanceId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceID checks that 'input' can be parsed as a Device Management User Experience Analytics Work From Anywhere Model Performance ID
func ValidateDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Work From Anywhere Model Performance ID
func (id DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereModelPerformance/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Work From Anywhere Model Performance ID
func (id DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsWorkFromAnywhereModelPerformance", "userExperienceAnalyticsWorkFromAnywhereModelPerformance", "userExperienceAnalyticsWorkFromAnywhereModelPerformance"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsWorkFromAnywhereModelPerformanceId", "userExperienceAnalyticsWorkFromAnywhereModelPerformanceId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Work From Anywhere Model Performance ID
func (id DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Work From Anywhere Model Performance: %q", id.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Work From Anywhere Model Performance (%s)", strings.Join(components, "\n"))
}
