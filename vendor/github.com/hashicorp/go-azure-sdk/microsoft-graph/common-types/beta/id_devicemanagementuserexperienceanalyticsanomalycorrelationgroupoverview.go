package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId{}

// DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId is a struct representing the Resource ID for a Device Management User Experience Analytics Anomaly Correlation Group Overview
type DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId struct {
	UserExperienceAnalyticsAnomalyCorrelationGroupOverviewId string
}

// NewDeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewID returns a new DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId struct
func NewDeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewID(userExperienceAnalyticsAnomalyCorrelationGroupOverviewId string) DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId {
	return DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId{
		UserExperienceAnalyticsAnomalyCorrelationGroupOverviewId: userExperienceAnalyticsAnomalyCorrelationGroupOverviewId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewID parses 'input' into a DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId
func ParseDeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewID(input string) (*DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewId, ok = input.Parsed["userExperienceAnalyticsAnomalyCorrelationGroupOverviewId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsAnomalyCorrelationGroupOverviewId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewID checks that 'input' can be parsed as a Device Management User Experience Analytics Anomaly Correlation Group Overview ID
func ValidateDeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Anomaly Correlation Group Overview ID
func (id DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsAnomalyCorrelationGroupOverview/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Anomaly Correlation Group Overview ID
func (id DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsAnomalyCorrelationGroupOverview", "userExperienceAnalyticsAnomalyCorrelationGroupOverview", "userExperienceAnalyticsAnomalyCorrelationGroupOverview"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsAnomalyCorrelationGroupOverviewId", "userExperienceAnalyticsAnomalyCorrelationGroupOverviewId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Anomaly Correlation Group Overview ID
func (id DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Anomaly Correlation Group Overview: %q", id.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Anomaly Correlation Group Overview (%s)", strings.Join(components, "\n"))
}
