package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAnalyticsActivityStatisticId{}

// MeAnalyticsActivityStatisticId is a struct representing the Resource ID for a Me Analytics Activity Statistic
type MeAnalyticsActivityStatisticId struct {
	ActivityStatisticsId string
}

// NewMeAnalyticsActivityStatisticID returns a new MeAnalyticsActivityStatisticId struct
func NewMeAnalyticsActivityStatisticID(activityStatisticsId string) MeAnalyticsActivityStatisticId {
	return MeAnalyticsActivityStatisticId{
		ActivityStatisticsId: activityStatisticsId,
	}
}

// ParseMeAnalyticsActivityStatisticID parses 'input' into a MeAnalyticsActivityStatisticId
func ParseMeAnalyticsActivityStatisticID(input string) (*MeAnalyticsActivityStatisticId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAnalyticsActivityStatisticId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAnalyticsActivityStatisticId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeAnalyticsActivityStatisticIDInsensitively parses 'input' case-insensitively into a MeAnalyticsActivityStatisticId
// note: this method should only be used for API response data and not user input
func ParseMeAnalyticsActivityStatisticIDInsensitively(input string) (*MeAnalyticsActivityStatisticId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAnalyticsActivityStatisticId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAnalyticsActivityStatisticId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeAnalyticsActivityStatisticId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ActivityStatisticsId, ok = input.Parsed["activityStatisticsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "activityStatisticsId", input)
	}

	return nil
}

// ValidateMeAnalyticsActivityStatisticID checks that 'input' can be parsed as a Me Analytics Activity Statistic ID
func ValidateMeAnalyticsActivityStatisticID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAnalyticsActivityStatisticID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Analytics Activity Statistic ID
func (id MeAnalyticsActivityStatisticId) ID() string {
	fmtString := "/me/analytics/activityStatistics/%s"
	return fmt.Sprintf(fmtString, id.ActivityStatisticsId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Analytics Activity Statistic ID
func (id MeAnalyticsActivityStatisticId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("analytics", "analytics", "analytics"),
		resourceids.StaticSegment("activityStatistics", "activityStatistics", "activityStatistics"),
		resourceids.UserSpecifiedSegment("activityStatisticsId", "activityStatisticsId"),
	}
}

// String returns a human-readable description of this Me Analytics Activity Statistic ID
func (id MeAnalyticsActivityStatisticId) String() string {
	components := []string{
		fmt.Sprintf("Activity Statistics: %q", id.ActivityStatisticsId),
	}
	return fmt.Sprintf("Me Analytics Activity Statistic (%s)", strings.Join(components, "\n"))
}
