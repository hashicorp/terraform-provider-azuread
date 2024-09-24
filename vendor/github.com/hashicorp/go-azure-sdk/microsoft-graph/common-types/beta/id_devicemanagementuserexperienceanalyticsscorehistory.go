package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsScoreHistoryId{}

// DeviceManagementUserExperienceAnalyticsScoreHistoryId is a struct representing the Resource ID for a Device Management User Experience Analytics Score History
type DeviceManagementUserExperienceAnalyticsScoreHistoryId struct {
	UserExperienceAnalyticsScoreHistoryId string
}

// NewDeviceManagementUserExperienceAnalyticsScoreHistoryID returns a new DeviceManagementUserExperienceAnalyticsScoreHistoryId struct
func NewDeviceManagementUserExperienceAnalyticsScoreHistoryID(userExperienceAnalyticsScoreHistoryId string) DeviceManagementUserExperienceAnalyticsScoreHistoryId {
	return DeviceManagementUserExperienceAnalyticsScoreHistoryId{
		UserExperienceAnalyticsScoreHistoryId: userExperienceAnalyticsScoreHistoryId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsScoreHistoryID parses 'input' into a DeviceManagementUserExperienceAnalyticsScoreHistoryId
func ParseDeviceManagementUserExperienceAnalyticsScoreHistoryID(input string) (*DeviceManagementUserExperienceAnalyticsScoreHistoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsScoreHistoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsScoreHistoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsScoreHistoryIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsScoreHistoryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsScoreHistoryIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsScoreHistoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsScoreHistoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsScoreHistoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsScoreHistoryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsScoreHistoryId, ok = input.Parsed["userExperienceAnalyticsScoreHistoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsScoreHistoryId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsScoreHistoryID checks that 'input' can be parsed as a Device Management User Experience Analytics Score History ID
func ValidateDeviceManagementUserExperienceAnalyticsScoreHistoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsScoreHistoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Score History ID
func (id DeviceManagementUserExperienceAnalyticsScoreHistoryId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsScoreHistory/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsScoreHistoryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Score History ID
func (id DeviceManagementUserExperienceAnalyticsScoreHistoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsScoreHistory", "userExperienceAnalyticsScoreHistory", "userExperienceAnalyticsScoreHistory"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsScoreHistoryId", "userExperienceAnalyticsScoreHistoryId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Score History ID
func (id DeviceManagementUserExperienceAnalyticsScoreHistoryId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Score History: %q", id.UserExperienceAnalyticsScoreHistoryId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Score History (%s)", strings.Join(components, "\n"))
}
