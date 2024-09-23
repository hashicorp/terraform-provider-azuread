package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsResourcePerformanceId{}

// DeviceManagementUserExperienceAnalyticsResourcePerformanceId is a struct representing the Resource ID for a Device Management User Experience Analytics Resource Performance
type DeviceManagementUserExperienceAnalyticsResourcePerformanceId struct {
	UserExperienceAnalyticsResourcePerformanceId string
}

// NewDeviceManagementUserExperienceAnalyticsResourcePerformanceID returns a new DeviceManagementUserExperienceAnalyticsResourcePerformanceId struct
func NewDeviceManagementUserExperienceAnalyticsResourcePerformanceID(userExperienceAnalyticsResourcePerformanceId string) DeviceManagementUserExperienceAnalyticsResourcePerformanceId {
	return DeviceManagementUserExperienceAnalyticsResourcePerformanceId{
		UserExperienceAnalyticsResourcePerformanceId: userExperienceAnalyticsResourcePerformanceId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsResourcePerformanceID parses 'input' into a DeviceManagementUserExperienceAnalyticsResourcePerformanceId
func ParseDeviceManagementUserExperienceAnalyticsResourcePerformanceID(input string) (*DeviceManagementUserExperienceAnalyticsResourcePerformanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsResourcePerformanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsResourcePerformanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsResourcePerformanceIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsResourcePerformanceId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsResourcePerformanceIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsResourcePerformanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsResourcePerformanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsResourcePerformanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsResourcePerformanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsResourcePerformanceId, ok = input.Parsed["userExperienceAnalyticsResourcePerformanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsResourcePerformanceId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsResourcePerformanceID checks that 'input' can be parsed as a Device Management User Experience Analytics Resource Performance ID
func ValidateDeviceManagementUserExperienceAnalyticsResourcePerformanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsResourcePerformanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Resource Performance ID
func (id DeviceManagementUserExperienceAnalyticsResourcePerformanceId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsResourcePerformance/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsResourcePerformanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Resource Performance ID
func (id DeviceManagementUserExperienceAnalyticsResourcePerformanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsResourcePerformance", "userExperienceAnalyticsResourcePerformance", "userExperienceAnalyticsResourcePerformance"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsResourcePerformanceId", "userExperienceAnalyticsResourcePerformanceId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Resource Performance ID
func (id DeviceManagementUserExperienceAnalyticsResourcePerformanceId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Resource Performance: %q", id.UserExperienceAnalyticsResourcePerformanceId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Resource Performance (%s)", strings.Join(components, "\n"))
}
