package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId{}

// DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId is a struct representing the Resource ID for a Device Management User Experience Analytics App Health Application Performance By App Version Device Id
type DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId struct {
	UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId string
}

// NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdID returns a new DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId struct
func NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdID(userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId string) DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId {
	return DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId{
		UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId: userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdID parses 'input' into a DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId
func ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdID(input string) (*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId, ok = input.Parsed["userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdID checks that 'input' can be parsed as a Device Management User Experience Analytics App Health Application Performance By App Version Device Id ID
func ValidateDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics App Health Application Performance By App Version Device Id ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics App Health Application Performance By App Version Device Id ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId", "userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId", "userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId", "userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics App Health Application Performance By App Version Device Id ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics App Health App Performance By App Version Device Id: %q", id.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics App Health Application Performance By App Version Device Id (%s)", strings.Join(components, "\n"))
}
