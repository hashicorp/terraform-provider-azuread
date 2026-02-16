package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsBaselineId{}

// DeviceManagementUserExperienceAnalyticsBaselineId is a struct representing the Resource ID for a Device Management User Experience Analytics Baseline
type DeviceManagementUserExperienceAnalyticsBaselineId struct {
	UserExperienceAnalyticsBaselineId string
}

// NewDeviceManagementUserExperienceAnalyticsBaselineID returns a new DeviceManagementUserExperienceAnalyticsBaselineId struct
func NewDeviceManagementUserExperienceAnalyticsBaselineID(userExperienceAnalyticsBaselineId string) DeviceManagementUserExperienceAnalyticsBaselineId {
	return DeviceManagementUserExperienceAnalyticsBaselineId{
		UserExperienceAnalyticsBaselineId: userExperienceAnalyticsBaselineId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsBaselineID parses 'input' into a DeviceManagementUserExperienceAnalyticsBaselineId
func ParseDeviceManagementUserExperienceAnalyticsBaselineID(input string) (*DeviceManagementUserExperienceAnalyticsBaselineId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsBaselineId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsBaselineId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsBaselineIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsBaselineId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsBaselineIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsBaselineId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsBaselineId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsBaselineId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsBaselineId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsBaselineId, ok = input.Parsed["userExperienceAnalyticsBaselineId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsBaselineId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsBaselineID checks that 'input' can be parsed as a Device Management User Experience Analytics Baseline ID
func ValidateDeviceManagementUserExperienceAnalyticsBaselineID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsBaselineID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Baseline ID
func (id DeviceManagementUserExperienceAnalyticsBaselineId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsBaselines/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsBaselineId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Baseline ID
func (id DeviceManagementUserExperienceAnalyticsBaselineId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsBaselines", "userExperienceAnalyticsBaselines", "userExperienceAnalyticsBaselines"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsBaselineId", "userExperienceAnalyticsBaselineId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Baseline ID
func (id DeviceManagementUserExperienceAnalyticsBaselineId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Baseline: %q", id.UserExperienceAnalyticsBaselineId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Baseline (%s)", strings.Join(components, "\n"))
}
