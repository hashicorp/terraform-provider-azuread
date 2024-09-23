package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId{}

// DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId is a struct representing the Resource ID for a Device Management User Experience Analytics Category Id Metric Value
type DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId struct {
	UserExperienceAnalyticsCategoryId string
	UserExperienceAnalyticsMetricId   string
}

// NewDeviceManagementUserExperienceAnalyticsCategoryIdMetricValueID returns a new DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId struct
func NewDeviceManagementUserExperienceAnalyticsCategoryIdMetricValueID(userExperienceAnalyticsCategoryId string, userExperienceAnalyticsMetricId string) DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId {
	return DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId{
		UserExperienceAnalyticsCategoryId: userExperienceAnalyticsCategoryId,
		UserExperienceAnalyticsMetricId:   userExperienceAnalyticsMetricId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsCategoryIdMetricValueID parses 'input' into a DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId
func ParseDeviceManagementUserExperienceAnalyticsCategoryIdMetricValueID(input string) (*DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsCategoryIdMetricValueIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsCategoryIdMetricValueIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsCategoryId, ok = input.Parsed["userExperienceAnalyticsCategoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsCategoryId", input)
	}

	if id.UserExperienceAnalyticsMetricId, ok = input.Parsed["userExperienceAnalyticsMetricId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsMetricId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsCategoryIdMetricValueID checks that 'input' can be parsed as a Device Management User Experience Analytics Category Id Metric Value ID
func ValidateDeviceManagementUserExperienceAnalyticsCategoryIdMetricValueID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsCategoryIdMetricValueID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Category Id Metric Value ID
func (id DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsCategories/%s/metricValues/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsCategoryId, id.UserExperienceAnalyticsMetricId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Category Id Metric Value ID
func (id DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsCategories", "userExperienceAnalyticsCategories", "userExperienceAnalyticsCategories"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsCategoryId", "userExperienceAnalyticsCategoryId"),
		resourceids.StaticSegment("metricValues", "metricValues", "metricValues"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsMetricId", "userExperienceAnalyticsMetricId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Category Id Metric Value ID
func (id DeviceManagementUserExperienceAnalyticsCategoryIdMetricValueId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Category: %q", id.UserExperienceAnalyticsCategoryId),
		fmt.Sprintf("User Experience Analytics Metric: %q", id.UserExperienceAnalyticsMetricId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Category Id Metric Value (%s)", strings.Join(components, "\n"))
}
