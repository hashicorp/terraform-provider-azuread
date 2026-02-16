package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsDeviceScoreId{}

// DeviceManagementUserExperienceAnalyticsDeviceScoreId is a struct representing the Resource ID for a Device Management User Experience Analytics Device Score
type DeviceManagementUserExperienceAnalyticsDeviceScoreId struct {
	UserExperienceAnalyticsDeviceScoresId string
}

// NewDeviceManagementUserExperienceAnalyticsDeviceScoreID returns a new DeviceManagementUserExperienceAnalyticsDeviceScoreId struct
func NewDeviceManagementUserExperienceAnalyticsDeviceScoreID(userExperienceAnalyticsDeviceScoresId string) DeviceManagementUserExperienceAnalyticsDeviceScoreId {
	return DeviceManagementUserExperienceAnalyticsDeviceScoreId{
		UserExperienceAnalyticsDeviceScoresId: userExperienceAnalyticsDeviceScoresId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsDeviceScoreID parses 'input' into a DeviceManagementUserExperienceAnalyticsDeviceScoreId
func ParseDeviceManagementUserExperienceAnalyticsDeviceScoreID(input string) (*DeviceManagementUserExperienceAnalyticsDeviceScoreId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsDeviceScoreId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsDeviceScoreId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsDeviceScoreIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsDeviceScoreId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsDeviceScoreIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsDeviceScoreId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsDeviceScoreId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsDeviceScoreId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsDeviceScoreId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsDeviceScoresId, ok = input.Parsed["userExperienceAnalyticsDeviceScoresId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsDeviceScoresId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsDeviceScoreID checks that 'input' can be parsed as a Device Management User Experience Analytics Device Score ID
func ValidateDeviceManagementUserExperienceAnalyticsDeviceScoreID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsDeviceScoreID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Device Score ID
func (id DeviceManagementUserExperienceAnalyticsDeviceScoreId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsDeviceScores/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsDeviceScoresId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Device Score ID
func (id DeviceManagementUserExperienceAnalyticsDeviceScoreId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsDeviceScores", "userExperienceAnalyticsDeviceScores", "userExperienceAnalyticsDeviceScores"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsDeviceScoresId", "userExperienceAnalyticsDeviceScoresId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Device Score ID
func (id DeviceManagementUserExperienceAnalyticsDeviceScoreId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Device Scores: %q", id.UserExperienceAnalyticsDeviceScoresId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Device Score (%s)", strings.Join(components, "\n"))
}
