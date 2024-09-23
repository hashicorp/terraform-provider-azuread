package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId{}

// DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId is a struct representing the Resource ID for a Device Management User Experience Analytics App Health Application Performance By App Version
type DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId struct {
	UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionId string
}

// NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionID returns a new DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId struct
func NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionID(userExperienceAnalyticsAppHealthAppPerformanceByAppVersionId string) DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId {
	return DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId{
		UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionId: userExperienceAnalyticsAppHealthAppPerformanceByAppVersionId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionID parses 'input' into a DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId
func ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionID(input string) (*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionId, ok = input.Parsed["userExperienceAnalyticsAppHealthAppPerformanceByAppVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsAppHealthAppPerformanceByAppVersionId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionID checks that 'input' can be parsed as a Device Management User Experience Analytics App Health Application Performance By App Version ID
func ValidateDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics App Health Application Performance By App Version ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics App Health Application Performance By App Version ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion", "userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion", "userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsAppHealthAppPerformanceByAppVersionId", "userExperienceAnalyticsAppHealthAppPerformanceByAppVersionId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics App Health Application Performance By App Version ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics App Health App Performance By App Version: %q", id.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics App Health Application Performance By App Version (%s)", strings.Join(components, "\n"))
}
