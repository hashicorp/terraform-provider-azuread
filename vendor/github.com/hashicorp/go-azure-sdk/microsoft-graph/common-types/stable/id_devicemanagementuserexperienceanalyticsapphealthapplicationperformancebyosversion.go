package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId{}

// DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId is a struct representing the Resource ID for a Device Management User Experience Analytics App Health Application Performance By OS Version
type DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId struct {
	UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionId string
}

// NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionID returns a new DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId struct
func NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionID(userExperienceAnalyticsAppHealthAppPerformanceByOSVersionId string) DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId {
	return DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId{
		UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionId: userExperienceAnalyticsAppHealthAppPerformanceByOSVersionId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionID parses 'input' into a DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId
func ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionID(input string) (*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionId, ok = input.Parsed["userExperienceAnalyticsAppHealthAppPerformanceByOSVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsAppHealthAppPerformanceByOSVersionId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionID checks that 'input' can be parsed as a Device Management User Experience Analytics App Health Application Performance By OS Version ID
func ValidateDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics App Health Application Performance By OS Version ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics App Health Application Performance By OS Version ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion", "userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion", "userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsAppHealthAppPerformanceByOSVersionId", "userExperienceAnalyticsAppHealthAppPerformanceByOSVersionId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics App Health Application Performance By OS Version ID
func (id DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics App Health App Performance By OS Version: %q", id.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics App Health Application Performance By OS Version (%s)", strings.Join(components, "\n"))
}
