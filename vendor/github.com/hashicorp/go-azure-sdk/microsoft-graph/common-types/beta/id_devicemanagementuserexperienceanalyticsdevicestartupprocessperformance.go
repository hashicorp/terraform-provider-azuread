package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId{}

// DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId is a struct representing the Resource ID for a Device Management User Experience Analytics Device Startup Process Performance
type DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId struct {
	UserExperienceAnalyticsDeviceStartupProcessPerformanceId string
}

// NewDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceID returns a new DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId struct
func NewDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceID(userExperienceAnalyticsDeviceStartupProcessPerformanceId string) DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId {
	return DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId{
		UserExperienceAnalyticsDeviceStartupProcessPerformanceId: userExperienceAnalyticsDeviceStartupProcessPerformanceId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceID parses 'input' into a DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId
func ParseDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceID(input string) (*DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsDeviceStartupProcessPerformanceId, ok = input.Parsed["userExperienceAnalyticsDeviceStartupProcessPerformanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsDeviceStartupProcessPerformanceId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceID checks that 'input' can be parsed as a Device Management User Experience Analytics Device Startup Process Performance ID
func ValidateDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Device Startup Process Performance ID
func (id DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsDeviceStartupProcessPerformance/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsDeviceStartupProcessPerformanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Device Startup Process Performance ID
func (id DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsDeviceStartupProcessPerformance", "userExperienceAnalyticsDeviceStartupProcessPerformance", "userExperienceAnalyticsDeviceStartupProcessPerformance"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsDeviceStartupProcessPerformanceId", "userExperienceAnalyticsDeviceStartupProcessPerformanceId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Device Startup Process Performance ID
func (id DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Device Startup Process Performance: %q", id.UserExperienceAnalyticsDeviceStartupProcessPerformanceId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Device Startup Process Performance (%s)", strings.Join(components, "\n"))
}
