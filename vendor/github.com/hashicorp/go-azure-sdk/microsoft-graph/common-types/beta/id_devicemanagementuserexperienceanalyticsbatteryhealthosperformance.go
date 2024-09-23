package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId{}

// DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId is a struct representing the Resource ID for a Device Management User Experience Analytics Battery Health Os Performance
type DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId struct {
	UserExperienceAnalyticsBatteryHealthOsPerformanceId string
}

// NewDeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceID returns a new DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId struct
func NewDeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceID(userExperienceAnalyticsBatteryHealthOsPerformanceId string) DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId {
	return DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId{
		UserExperienceAnalyticsBatteryHealthOsPerformanceId: userExperienceAnalyticsBatteryHealthOsPerformanceId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceID parses 'input' into a DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId
func ParseDeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceID(input string) (*DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsBatteryHealthOsPerformanceId, ok = input.Parsed["userExperienceAnalyticsBatteryHealthOsPerformanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsBatteryHealthOsPerformanceId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceID checks that 'input' can be parsed as a Device Management User Experience Analytics Battery Health Os Performance ID
func ValidateDeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Battery Health Os Performance ID
func (id DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsBatteryHealthOsPerformance/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsBatteryHealthOsPerformanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Battery Health Os Performance ID
func (id DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsBatteryHealthOsPerformance", "userExperienceAnalyticsBatteryHealthOsPerformance", "userExperienceAnalyticsBatteryHealthOsPerformance"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsBatteryHealthOsPerformanceId", "userExperienceAnalyticsBatteryHealthOsPerformanceId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Battery Health Os Performance ID
func (id DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Battery Health Os Performance: %q", id.UserExperienceAnalyticsBatteryHealthOsPerformanceId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Battery Health Os Performance (%s)", strings.Join(components, "\n"))
}
