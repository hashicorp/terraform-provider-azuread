package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId{}

// DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId is a struct representing the Resource ID for a Device Management User Experience Analytics Device Startup History
type DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId struct {
	UserExperienceAnalyticsDeviceStartupHistoryId string
}

// NewDeviceManagementUserExperienceAnalyticsDeviceStartupHistoryID returns a new DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId struct
func NewDeviceManagementUserExperienceAnalyticsDeviceStartupHistoryID(userExperienceAnalyticsDeviceStartupHistoryId string) DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId {
	return DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId{
		UserExperienceAnalyticsDeviceStartupHistoryId: userExperienceAnalyticsDeviceStartupHistoryId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsDeviceStartupHistoryID parses 'input' into a DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId
func ParseDeviceManagementUserExperienceAnalyticsDeviceStartupHistoryID(input string) (*DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsDeviceStartupHistoryIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsDeviceStartupHistoryIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsDeviceStartupHistoryId, ok = input.Parsed["userExperienceAnalyticsDeviceStartupHistoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsDeviceStartupHistoryId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsDeviceStartupHistoryID checks that 'input' can be parsed as a Device Management User Experience Analytics Device Startup History ID
func ValidateDeviceManagementUserExperienceAnalyticsDeviceStartupHistoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsDeviceStartupHistoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Device Startup History ID
func (id DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsDeviceStartupHistory/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsDeviceStartupHistoryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Device Startup History ID
func (id DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsDeviceStartupHistory", "userExperienceAnalyticsDeviceStartupHistory", "userExperienceAnalyticsDeviceStartupHistory"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsDeviceStartupHistoryId", "userExperienceAnalyticsDeviceStartupHistoryId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Device Startup History ID
func (id DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Device Startup History: %q", id.UserExperienceAnalyticsDeviceStartupHistoryId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Device Startup History (%s)", strings.Join(components, "\n"))
}
