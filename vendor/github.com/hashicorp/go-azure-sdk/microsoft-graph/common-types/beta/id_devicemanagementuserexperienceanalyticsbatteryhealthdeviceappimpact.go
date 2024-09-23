package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId{}

// DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId is a struct representing the Resource ID for a Device Management User Experience Analytics Battery Health Device App Impact
type DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId struct {
	UserExperienceAnalyticsBatteryHealthDeviceAppImpactId string
}

// NewDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactID returns a new DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId struct
func NewDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactID(userExperienceAnalyticsBatteryHealthDeviceAppImpactId string) DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId {
	return DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId{
		UserExperienceAnalyticsBatteryHealthDeviceAppImpactId: userExperienceAnalyticsBatteryHealthDeviceAppImpactId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactID parses 'input' into a DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId
func ParseDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactID(input string) (*DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsBatteryHealthDeviceAppImpactId, ok = input.Parsed["userExperienceAnalyticsBatteryHealthDeviceAppImpactId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsBatteryHealthDeviceAppImpactId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactID checks that 'input' can be parsed as a Device Management User Experience Analytics Battery Health Device App Impact ID
func ValidateDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Battery Health Device App Impact ID
func (id DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceAppImpact/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsBatteryHealthDeviceAppImpactId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Battery Health Device App Impact ID
func (id DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsBatteryHealthDeviceAppImpact", "userExperienceAnalyticsBatteryHealthDeviceAppImpact", "userExperienceAnalyticsBatteryHealthDeviceAppImpact"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsBatteryHealthDeviceAppImpactId", "userExperienceAnalyticsBatteryHealthDeviceAppImpactId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Battery Health Device App Impact ID
func (id DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Battery Health Device App Impact: %q", id.UserExperienceAnalyticsBatteryHealthDeviceAppImpactId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Battery Health Device App Impact (%s)", strings.Join(components, "\n"))
}
