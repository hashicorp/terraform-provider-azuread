package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId{}

// DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId is a struct representing the Resource ID for a Device Management User Experience Analytics Battery Health Model Performance
type DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId struct {
	UserExperienceAnalyticsBatteryHealthModelPerformanceId string
}

// NewDeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceID returns a new DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId struct
func NewDeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceID(userExperienceAnalyticsBatteryHealthModelPerformanceId string) DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId {
	return DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId{
		UserExperienceAnalyticsBatteryHealthModelPerformanceId: userExperienceAnalyticsBatteryHealthModelPerformanceId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceID parses 'input' into a DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId
func ParseDeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceID(input string) (*DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsBatteryHealthModelPerformanceId, ok = input.Parsed["userExperienceAnalyticsBatteryHealthModelPerformanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsBatteryHealthModelPerformanceId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceID checks that 'input' can be parsed as a Device Management User Experience Analytics Battery Health Model Performance ID
func ValidateDeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Battery Health Model Performance ID
func (id DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsBatteryHealthModelPerformance/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsBatteryHealthModelPerformanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Battery Health Model Performance ID
func (id DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsBatteryHealthModelPerformance", "userExperienceAnalyticsBatteryHealthModelPerformance", "userExperienceAnalyticsBatteryHealthModelPerformance"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsBatteryHealthModelPerformanceId", "userExperienceAnalyticsBatteryHealthModelPerformanceId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Battery Health Model Performance ID
func (id DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Battery Health Model Performance: %q", id.UserExperienceAnalyticsBatteryHealthModelPerformanceId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Battery Health Model Performance (%s)", strings.Join(components, "\n"))
}
