package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId{}

// DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId is a struct representing the Resource ID for a Device Management User Experience Analytics Battery Health App Impact
type DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId struct {
	UserExperienceAnalyticsBatteryHealthAppImpactId string
}

// NewDeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactID returns a new DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId struct
func NewDeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactID(userExperienceAnalyticsBatteryHealthAppImpactId string) DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId {
	return DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId{
		UserExperienceAnalyticsBatteryHealthAppImpactId: userExperienceAnalyticsBatteryHealthAppImpactId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactID parses 'input' into a DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId
func ParseDeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactID(input string) (*DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsBatteryHealthAppImpactId, ok = input.Parsed["userExperienceAnalyticsBatteryHealthAppImpactId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsBatteryHealthAppImpactId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactID checks that 'input' can be parsed as a Device Management User Experience Analytics Battery Health App Impact ID
func ValidateDeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Battery Health App Impact ID
func (id DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsBatteryHealthAppImpact/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsBatteryHealthAppImpactId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Battery Health App Impact ID
func (id DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsBatteryHealthAppImpact", "userExperienceAnalyticsBatteryHealthAppImpact", "userExperienceAnalyticsBatteryHealthAppImpact"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsBatteryHealthAppImpactId", "userExperienceAnalyticsBatteryHealthAppImpactId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Battery Health App Impact ID
func (id DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Battery Health App Impact: %q", id.UserExperienceAnalyticsBatteryHealthAppImpactId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Battery Health App Impact (%s)", strings.Join(components, "\n"))
}
