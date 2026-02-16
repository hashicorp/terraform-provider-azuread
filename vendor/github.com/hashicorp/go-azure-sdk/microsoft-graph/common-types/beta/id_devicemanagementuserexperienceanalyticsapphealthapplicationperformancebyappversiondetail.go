package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId{}

// DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId is a struct representing the Resource ID for a Device Management User Experience Analytics App Health Application Performance By App Version Detail
type DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId struct {
	UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId string
}

// NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailID returns a new DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId struct
func NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailID(userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId string) DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId {
	return DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId{
		UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId: userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailID parses 'input' into a DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId
func ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailID(input string) (*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId, ok = input.Parsed["userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailID checks that 'input' can be parsed as a Device Management User Experience Analytics App Health Application Performance By App Version Detail ID
func ValidateDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics App Health Application Performance By App Version Detail ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics App Health Application Performance By App Version Detail ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails", "userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails", "userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId", "userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics App Health Application Performance By App Version Detail ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics App Health App Performance By App Version Details: %q", id.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics App Health Application Performance By App Version Detail (%s)", strings.Join(components, "\n"))
}
