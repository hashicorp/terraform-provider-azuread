package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsMetricHistoryId{}

// DeviceManagementUserExperienceAnalyticsMetricHistoryId is a struct representing the Resource ID for a Device Management User Experience Analytics Metric History
type DeviceManagementUserExperienceAnalyticsMetricHistoryId struct {
	UserExperienceAnalyticsMetricHistoryId string
}

// NewDeviceManagementUserExperienceAnalyticsMetricHistoryID returns a new DeviceManagementUserExperienceAnalyticsMetricHistoryId struct
func NewDeviceManagementUserExperienceAnalyticsMetricHistoryID(userExperienceAnalyticsMetricHistoryId string) DeviceManagementUserExperienceAnalyticsMetricHistoryId {
	return DeviceManagementUserExperienceAnalyticsMetricHistoryId{
		UserExperienceAnalyticsMetricHistoryId: userExperienceAnalyticsMetricHistoryId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsMetricHistoryID parses 'input' into a DeviceManagementUserExperienceAnalyticsMetricHistoryId
func ParseDeviceManagementUserExperienceAnalyticsMetricHistoryID(input string) (*DeviceManagementUserExperienceAnalyticsMetricHistoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsMetricHistoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsMetricHistoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsMetricHistoryIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsMetricHistoryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsMetricHistoryIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsMetricHistoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsMetricHistoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsMetricHistoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsMetricHistoryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsMetricHistoryId, ok = input.Parsed["userExperienceAnalyticsMetricHistoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsMetricHistoryId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsMetricHistoryID checks that 'input' can be parsed as a Device Management User Experience Analytics Metric History ID
func ValidateDeviceManagementUserExperienceAnalyticsMetricHistoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsMetricHistoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Metric History ID
func (id DeviceManagementUserExperienceAnalyticsMetricHistoryId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsMetricHistory/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsMetricHistoryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Metric History ID
func (id DeviceManagementUserExperienceAnalyticsMetricHistoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsMetricHistory", "userExperienceAnalyticsMetricHistory", "userExperienceAnalyticsMetricHistory"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsMetricHistoryId", "userExperienceAnalyticsMetricHistoryId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Metric History ID
func (id DeviceManagementUserExperienceAnalyticsMetricHistoryId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Metric History: %q", id.UserExperienceAnalyticsMetricHistoryId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Metric History (%s)", strings.Join(components, "\n"))
}
