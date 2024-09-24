package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId{}

// DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId is a struct representing the Resource ID for a Device Management User Experience Analytics App Health Overview Metric Value
type DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId struct {
	UserExperienceAnalyticsMetricId string
}

// NewDeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueID returns a new DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId struct
func NewDeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueID(userExperienceAnalyticsMetricId string) DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId {
	return DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId{
		UserExperienceAnalyticsMetricId: userExperienceAnalyticsMetricId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueID parses 'input' into a DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId
func ParseDeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueID(input string) (*DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsMetricId, ok = input.Parsed["userExperienceAnalyticsMetricId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsMetricId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueID checks that 'input' can be parsed as a Device Management User Experience Analytics App Health Overview Metric Value ID
func ValidateDeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics App Health Overview Metric Value ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsAppHealthOverview/metricValues/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsMetricId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics App Health Overview Metric Value ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsAppHealthOverview", "userExperienceAnalyticsAppHealthOverview", "userExperienceAnalyticsAppHealthOverview"),
		resourceids.StaticSegment("metricValues", "metricValues", "metricValues"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsMetricId", "userExperienceAnalyticsMetricId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics App Health Overview Metric Value ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthOverviewMetricValueId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Metric: %q", id.UserExperienceAnalyticsMetricId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics App Health Overview Metric Value (%s)", strings.Join(components, "\n"))
}
