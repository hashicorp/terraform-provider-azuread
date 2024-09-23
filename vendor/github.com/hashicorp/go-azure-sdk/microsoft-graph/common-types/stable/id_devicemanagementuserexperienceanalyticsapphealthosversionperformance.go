package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId{}

// DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId is a struct representing the Resource ID for a Device Management User Experience Analytics App Health O S Version Performance
type DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId struct {
	UserExperienceAnalyticsAppHealthOSVersionPerformanceId string
}

// NewDeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceID returns a new DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId struct
func NewDeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceID(userExperienceAnalyticsAppHealthOSVersionPerformanceId string) DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId {
	return DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId{
		UserExperienceAnalyticsAppHealthOSVersionPerformanceId: userExperienceAnalyticsAppHealthOSVersionPerformanceId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceID parses 'input' into a DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId
func ParseDeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceID(input string) (*DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsAppHealthOSVersionPerformanceId, ok = input.Parsed["userExperienceAnalyticsAppHealthOSVersionPerformanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsAppHealthOSVersionPerformanceId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceID checks that 'input' can be parsed as a Device Management User Experience Analytics App Health O S Version Performance ID
func ValidateDeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics App Health O S Version Performance ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsAppHealthOSVersionPerformance/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsAppHealthOSVersionPerformanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics App Health O S Version Performance ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsAppHealthOSVersionPerformance", "userExperienceAnalyticsAppHealthOSVersionPerformance", "userExperienceAnalyticsAppHealthOSVersionPerformance"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsAppHealthOSVersionPerformanceId", "userExperienceAnalyticsAppHealthOSVersionPerformanceId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics App Health O S Version Performance ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics App Health O S Version Performance: %q", id.UserExperienceAnalyticsAppHealthOSVersionPerformanceId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics App Health O S Version Performance (%s)", strings.Join(components, "\n"))
}
