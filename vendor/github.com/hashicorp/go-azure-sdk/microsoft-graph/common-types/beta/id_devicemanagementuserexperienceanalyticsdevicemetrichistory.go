package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId{}

// DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId is a struct representing the Resource ID for a Device Management User Experience Analytics Device Metric History
type DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId struct {
	UserExperienceAnalyticsMetricHistoryId string
}

// NewDeviceManagementUserExperienceAnalyticsDeviceMetricHistoryID returns a new DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId struct
func NewDeviceManagementUserExperienceAnalyticsDeviceMetricHistoryID(userExperienceAnalyticsMetricHistoryId string) DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId {
	return DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId{
		UserExperienceAnalyticsMetricHistoryId: userExperienceAnalyticsMetricHistoryId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsDeviceMetricHistoryID parses 'input' into a DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId
func ParseDeviceManagementUserExperienceAnalyticsDeviceMetricHistoryID(input string) (*DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsDeviceMetricHistoryIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsDeviceMetricHistoryIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsMetricHistoryId, ok = input.Parsed["userExperienceAnalyticsMetricHistoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsMetricHistoryId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsDeviceMetricHistoryID checks that 'input' can be parsed as a Device Management User Experience Analytics Device Metric History ID
func ValidateDeviceManagementUserExperienceAnalyticsDeviceMetricHistoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsDeviceMetricHistoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Device Metric History ID
func (id DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsDeviceMetricHistory/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsMetricHistoryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Device Metric History ID
func (id DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsDeviceMetricHistory", "userExperienceAnalyticsDeviceMetricHistory", "userExperienceAnalyticsDeviceMetricHistory"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsMetricHistoryId", "userExperienceAnalyticsMetricHistoryId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Device Metric History ID
func (id DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Metric History: %q", id.UserExperienceAnalyticsMetricHistoryId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Device Metric History (%s)", strings.Join(components, "\n"))
}
