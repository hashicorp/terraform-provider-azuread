package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId{}

// DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId is a struct representing the Resource ID for a Device Management User Experience Analytics Work From Anywhere Metric
type DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId struct {
	UserExperienceAnalyticsWorkFromAnywhereMetricId string
}

// NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricID returns a new DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId struct
func NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricID(userExperienceAnalyticsWorkFromAnywhereMetricId string) DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId {
	return DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId{
		UserExperienceAnalyticsWorkFromAnywhereMetricId: userExperienceAnalyticsWorkFromAnywhereMetricId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricID parses 'input' into a DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId
func ParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricID(input string) (*DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsWorkFromAnywhereMetricId, ok = input.Parsed["userExperienceAnalyticsWorkFromAnywhereMetricId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsWorkFromAnywhereMetricId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricID checks that 'input' can be parsed as a Device Management User Experience Analytics Work From Anywhere Metric ID
func ValidateDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Work From Anywhere Metric ID
func (id DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsWorkFromAnywhereMetricId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Work From Anywhere Metric ID
func (id DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsWorkFromAnywhereMetrics", "userExperienceAnalyticsWorkFromAnywhereMetrics", "userExperienceAnalyticsWorkFromAnywhereMetrics"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsWorkFromAnywhereMetricId", "userExperienceAnalyticsWorkFromAnywhereMetricId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Work From Anywhere Metric ID
func (id DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Work From Anywhere Metric: %q", id.UserExperienceAnalyticsWorkFromAnywhereMetricId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Work From Anywhere Metric (%s)", strings.Join(components, "\n"))
}
