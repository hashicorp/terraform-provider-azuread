package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId{}

// DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId is a struct representing the Resource ID for a Device Management User Experience Analytics App Health Device Performance
type DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId struct {
	UserExperienceAnalyticsAppHealthDevicePerformanceId string
}

// NewDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceID returns a new DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId struct
func NewDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceID(userExperienceAnalyticsAppHealthDevicePerformanceId string) DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId {
	return DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId{
		UserExperienceAnalyticsAppHealthDevicePerformanceId: userExperienceAnalyticsAppHealthDevicePerformanceId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceID parses 'input' into a DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId
func ParseDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceID(input string) (*DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsAppHealthDevicePerformanceId, ok = input.Parsed["userExperienceAnalyticsAppHealthDevicePerformanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsAppHealthDevicePerformanceId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceID checks that 'input' can be parsed as a Device Management User Experience Analytics App Health Device Performance ID
func ValidateDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics App Health Device Performance ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsAppHealthDevicePerformance/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsAppHealthDevicePerformanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics App Health Device Performance ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsAppHealthDevicePerformance", "userExperienceAnalyticsAppHealthDevicePerformance", "userExperienceAnalyticsAppHealthDevicePerformance"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsAppHealthDevicePerformanceId", "userExperienceAnalyticsAppHealthDevicePerformanceId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics App Health Device Performance ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics App Health Device Performance: %q", id.UserExperienceAnalyticsAppHealthDevicePerformanceId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics App Health Device Performance (%s)", strings.Join(components, "\n"))
}
