package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId{}

// DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId is a struct representing the Resource ID for a Device Management User Experience Analytics Work From Anywhere Metric Id Metric Device
type DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId struct {
	UserExperienceAnalyticsWorkFromAnywhereMetricId string
	UserExperienceAnalyticsWorkFromAnywhereDeviceId string
}

// NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceID returns a new DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId struct
func NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceID(userExperienceAnalyticsWorkFromAnywhereMetricId string, userExperienceAnalyticsWorkFromAnywhereDeviceId string) DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId {
	return DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId{
		UserExperienceAnalyticsWorkFromAnywhereMetricId: userExperienceAnalyticsWorkFromAnywhereMetricId,
		UserExperienceAnalyticsWorkFromAnywhereDeviceId: userExperienceAnalyticsWorkFromAnywhereDeviceId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceID parses 'input' into a DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId
func ParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceID(input string) (*DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsWorkFromAnywhereMetricId, ok = input.Parsed["userExperienceAnalyticsWorkFromAnywhereMetricId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsWorkFromAnywhereMetricId", input)
	}

	if id.UserExperienceAnalyticsWorkFromAnywhereDeviceId, ok = input.Parsed["userExperienceAnalyticsWorkFromAnywhereDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsWorkFromAnywhereDeviceId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceID checks that 'input' can be parsed as a Device Management User Experience Analytics Work From Anywhere Metric Id Metric Device ID
func ValidateDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Work From Anywhere Metric Id Metric Device ID
func (id DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/%s/metricDevices/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsWorkFromAnywhereMetricId, id.UserExperienceAnalyticsWorkFromAnywhereDeviceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Work From Anywhere Metric Id Metric Device ID
func (id DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsWorkFromAnywhereMetrics", "userExperienceAnalyticsWorkFromAnywhereMetrics", "userExperienceAnalyticsWorkFromAnywhereMetrics"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsWorkFromAnywhereMetricId", "userExperienceAnalyticsWorkFromAnywhereMetricId"),
		resourceids.StaticSegment("metricDevices", "metricDevices", "metricDevices"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsWorkFromAnywhereDeviceId", "userExperienceAnalyticsWorkFromAnywhereDeviceId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Work From Anywhere Metric Id Metric Device ID
func (id DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Work From Anywhere Metric: %q", id.UserExperienceAnalyticsWorkFromAnywhereMetricId),
		fmt.Sprintf("User Experience Analytics Work From Anywhere Device: %q", id.UserExperienceAnalyticsWorkFromAnywhereDeviceId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Work From Anywhere Metric Id Metric Device (%s)", strings.Join(components, "\n"))
}
