package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAnomalyDeviceId{}

// DeviceManagementUserExperienceAnalyticsAnomalyDeviceId is a struct representing the Resource ID for a Device Management User Experience Analytics Anomaly Device
type DeviceManagementUserExperienceAnalyticsAnomalyDeviceId struct {
	UserExperienceAnalyticsAnomalyDeviceId string
}

// NewDeviceManagementUserExperienceAnalyticsAnomalyDeviceID returns a new DeviceManagementUserExperienceAnalyticsAnomalyDeviceId struct
func NewDeviceManagementUserExperienceAnalyticsAnomalyDeviceID(userExperienceAnalyticsAnomalyDeviceId string) DeviceManagementUserExperienceAnalyticsAnomalyDeviceId {
	return DeviceManagementUserExperienceAnalyticsAnomalyDeviceId{
		UserExperienceAnalyticsAnomalyDeviceId: userExperienceAnalyticsAnomalyDeviceId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsAnomalyDeviceID parses 'input' into a DeviceManagementUserExperienceAnalyticsAnomalyDeviceId
func ParseDeviceManagementUserExperienceAnalyticsAnomalyDeviceID(input string) (*DeviceManagementUserExperienceAnalyticsAnomalyDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAnomalyDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAnomalyDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsAnomalyDeviceIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsAnomalyDeviceId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsAnomalyDeviceIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsAnomalyDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAnomalyDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAnomalyDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsAnomalyDeviceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsAnomalyDeviceId, ok = input.Parsed["userExperienceAnalyticsAnomalyDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsAnomalyDeviceId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsAnomalyDeviceID checks that 'input' can be parsed as a Device Management User Experience Analytics Anomaly Device ID
func ValidateDeviceManagementUserExperienceAnalyticsAnomalyDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsAnomalyDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Anomaly Device ID
func (id DeviceManagementUserExperienceAnalyticsAnomalyDeviceId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsAnomalyDevice/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsAnomalyDeviceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Anomaly Device ID
func (id DeviceManagementUserExperienceAnalyticsAnomalyDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsAnomalyDevice", "userExperienceAnalyticsAnomalyDevice", "userExperienceAnalyticsAnomalyDevice"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsAnomalyDeviceId", "userExperienceAnalyticsAnomalyDeviceId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Anomaly Device ID
func (id DeviceManagementUserExperienceAnalyticsAnomalyDeviceId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Anomaly Device: %q", id.UserExperienceAnalyticsAnomalyDeviceId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Anomaly Device (%s)", strings.Join(components, "\n"))
}
