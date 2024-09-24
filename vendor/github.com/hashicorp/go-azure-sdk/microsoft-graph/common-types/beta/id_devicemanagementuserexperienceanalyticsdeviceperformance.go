package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsDevicePerformanceId{}

// DeviceManagementUserExperienceAnalyticsDevicePerformanceId is a struct representing the Resource ID for a Device Management User Experience Analytics Device Performance
type DeviceManagementUserExperienceAnalyticsDevicePerformanceId struct {
	UserExperienceAnalyticsDevicePerformanceId string
}

// NewDeviceManagementUserExperienceAnalyticsDevicePerformanceID returns a new DeviceManagementUserExperienceAnalyticsDevicePerformanceId struct
func NewDeviceManagementUserExperienceAnalyticsDevicePerformanceID(userExperienceAnalyticsDevicePerformanceId string) DeviceManagementUserExperienceAnalyticsDevicePerformanceId {
	return DeviceManagementUserExperienceAnalyticsDevicePerformanceId{
		UserExperienceAnalyticsDevicePerformanceId: userExperienceAnalyticsDevicePerformanceId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsDevicePerformanceID parses 'input' into a DeviceManagementUserExperienceAnalyticsDevicePerformanceId
func ParseDeviceManagementUserExperienceAnalyticsDevicePerformanceID(input string) (*DeviceManagementUserExperienceAnalyticsDevicePerformanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsDevicePerformanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsDevicePerformanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsDevicePerformanceIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsDevicePerformanceId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsDevicePerformanceIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsDevicePerformanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsDevicePerformanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsDevicePerformanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsDevicePerformanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsDevicePerformanceId, ok = input.Parsed["userExperienceAnalyticsDevicePerformanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsDevicePerformanceId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsDevicePerformanceID checks that 'input' can be parsed as a Device Management User Experience Analytics Device Performance ID
func ValidateDeviceManagementUserExperienceAnalyticsDevicePerformanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsDevicePerformanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Device Performance ID
func (id DeviceManagementUserExperienceAnalyticsDevicePerformanceId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsDevicePerformance/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsDevicePerformanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Device Performance ID
func (id DeviceManagementUserExperienceAnalyticsDevicePerformanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsDevicePerformance", "userExperienceAnalyticsDevicePerformance", "userExperienceAnalyticsDevicePerformance"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsDevicePerformanceId", "userExperienceAnalyticsDevicePerformanceId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Device Performance ID
func (id DeviceManagementUserExperienceAnalyticsDevicePerformanceId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Device Performance: %q", id.UserExperienceAnalyticsDevicePerformanceId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Device Performance (%s)", strings.Join(components, "\n"))
}
