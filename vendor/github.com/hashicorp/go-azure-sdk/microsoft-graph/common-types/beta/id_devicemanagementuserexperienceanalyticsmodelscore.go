package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsModelScoreId{}

// DeviceManagementUserExperienceAnalyticsModelScoreId is a struct representing the Resource ID for a Device Management User Experience Analytics Model Score
type DeviceManagementUserExperienceAnalyticsModelScoreId struct {
	UserExperienceAnalyticsModelScoresId string
}

// NewDeviceManagementUserExperienceAnalyticsModelScoreID returns a new DeviceManagementUserExperienceAnalyticsModelScoreId struct
func NewDeviceManagementUserExperienceAnalyticsModelScoreID(userExperienceAnalyticsModelScoresId string) DeviceManagementUserExperienceAnalyticsModelScoreId {
	return DeviceManagementUserExperienceAnalyticsModelScoreId{
		UserExperienceAnalyticsModelScoresId: userExperienceAnalyticsModelScoresId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsModelScoreID parses 'input' into a DeviceManagementUserExperienceAnalyticsModelScoreId
func ParseDeviceManagementUserExperienceAnalyticsModelScoreID(input string) (*DeviceManagementUserExperienceAnalyticsModelScoreId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsModelScoreId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsModelScoreId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsModelScoreIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsModelScoreId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsModelScoreIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsModelScoreId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsModelScoreId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsModelScoreId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsModelScoreId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsModelScoresId, ok = input.Parsed["userExperienceAnalyticsModelScoresId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsModelScoresId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsModelScoreID checks that 'input' can be parsed as a Device Management User Experience Analytics Model Score ID
func ValidateDeviceManagementUserExperienceAnalyticsModelScoreID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsModelScoreID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Model Score ID
func (id DeviceManagementUserExperienceAnalyticsModelScoreId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsModelScores/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsModelScoresId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Model Score ID
func (id DeviceManagementUserExperienceAnalyticsModelScoreId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsModelScores", "userExperienceAnalyticsModelScores", "userExperienceAnalyticsModelScores"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsModelScoresId", "userExperienceAnalyticsModelScoresId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Model Score ID
func (id DeviceManagementUserExperienceAnalyticsModelScoreId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Model Scores: %q", id.UserExperienceAnalyticsModelScoresId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Model Score (%s)", strings.Join(components, "\n"))
}
