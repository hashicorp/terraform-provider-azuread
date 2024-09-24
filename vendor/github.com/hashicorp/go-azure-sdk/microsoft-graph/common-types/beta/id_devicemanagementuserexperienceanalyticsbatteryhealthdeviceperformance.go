package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId{}

// DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId is a struct representing the Resource ID for a Device Management User Experience Analytics Battery Health Device Performance
type DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId struct {
	UserExperienceAnalyticsBatteryHealthDevicePerformanceId string
}

// NewDeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceID returns a new DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId struct
func NewDeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceID(userExperienceAnalyticsBatteryHealthDevicePerformanceId string) DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId {
	return DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId{
		UserExperienceAnalyticsBatteryHealthDevicePerformanceId: userExperienceAnalyticsBatteryHealthDevicePerformanceId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceID parses 'input' into a DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId
func ParseDeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceID(input string) (*DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsBatteryHealthDevicePerformanceId, ok = input.Parsed["userExperienceAnalyticsBatteryHealthDevicePerformanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsBatteryHealthDevicePerformanceId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceID checks that 'input' can be parsed as a Device Management User Experience Analytics Battery Health Device Performance ID
func ValidateDeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Battery Health Device Performance ID
func (id DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsBatteryHealthDevicePerformance/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsBatteryHealthDevicePerformanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Battery Health Device Performance ID
func (id DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsBatteryHealthDevicePerformance", "userExperienceAnalyticsBatteryHealthDevicePerformance", "userExperienceAnalyticsBatteryHealthDevicePerformance"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsBatteryHealthDevicePerformanceId", "userExperienceAnalyticsBatteryHealthDevicePerformanceId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Battery Health Device Performance ID
func (id DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Battery Health Device Performance: %q", id.UserExperienceAnalyticsBatteryHealthDevicePerformanceId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Battery Health Device Performance (%s)", strings.Join(components, "\n"))
}
