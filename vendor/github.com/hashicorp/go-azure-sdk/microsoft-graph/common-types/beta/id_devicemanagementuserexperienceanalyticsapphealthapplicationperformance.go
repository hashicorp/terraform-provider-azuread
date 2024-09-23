package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId{}

// DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId is a struct representing the Resource ID for a Device Management User Experience Analytics App Health Application Performance
type DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId struct {
	UserExperienceAnalyticsAppHealthApplicationPerformanceId string
}

// NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceID returns a new DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId struct
func NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceID(userExperienceAnalyticsAppHealthApplicationPerformanceId string) DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId {
	return DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId{
		UserExperienceAnalyticsAppHealthApplicationPerformanceId: userExperienceAnalyticsAppHealthApplicationPerformanceId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceID parses 'input' into a DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId
func ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceID(input string) (*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsAppHealthApplicationPerformanceId, ok = input.Parsed["userExperienceAnalyticsAppHealthApplicationPerformanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsAppHealthApplicationPerformanceId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceID checks that 'input' can be parsed as a Device Management User Experience Analytics App Health Application Performance ID
func ValidateDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics App Health Application Performance ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformance/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsAppHealthApplicationPerformanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics App Health Application Performance ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsAppHealthApplicationPerformance", "userExperienceAnalyticsAppHealthApplicationPerformance", "userExperienceAnalyticsAppHealthApplicationPerformance"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsAppHealthApplicationPerformanceId", "userExperienceAnalyticsAppHealthApplicationPerformanceId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics App Health Application Performance ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics App Health Application Performance: %q", id.UserExperienceAnalyticsAppHealthApplicationPerformanceId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics App Health Application Performance (%s)", strings.Join(components, "\n"))
}
