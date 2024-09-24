package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId{}

// DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId is a struct representing the Resource ID for a Device Management User Experience Analytics App Health Device Model Performance
type DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId struct {
	UserExperienceAnalyticsAppHealthDeviceModelPerformanceId string
}

// NewDeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceID returns a new DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId struct
func NewDeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceID(userExperienceAnalyticsAppHealthDeviceModelPerformanceId string) DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId {
	return DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId{
		UserExperienceAnalyticsAppHealthDeviceModelPerformanceId: userExperienceAnalyticsAppHealthDeviceModelPerformanceId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceID parses 'input' into a DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId
func ParseDeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceID(input string) (*DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsAppHealthDeviceModelPerformanceId, ok = input.Parsed["userExperienceAnalyticsAppHealthDeviceModelPerformanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsAppHealthDeviceModelPerformanceId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceID checks that 'input' can be parsed as a Device Management User Experience Analytics App Health Device Model Performance ID
func ValidateDeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics App Health Device Model Performance ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsAppHealthDeviceModelPerformance/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsAppHealthDeviceModelPerformanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics App Health Device Model Performance ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsAppHealthDeviceModelPerformance", "userExperienceAnalyticsAppHealthDeviceModelPerformance", "userExperienceAnalyticsAppHealthDeviceModelPerformance"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsAppHealthDeviceModelPerformanceId", "userExperienceAnalyticsAppHealthDeviceModelPerformanceId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics App Health Device Model Performance ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics App Health Device Model Performance: %q", id.UserExperienceAnalyticsAppHealthDeviceModelPerformanceId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics App Health Device Model Performance (%s)", strings.Join(components, "\n"))
}
