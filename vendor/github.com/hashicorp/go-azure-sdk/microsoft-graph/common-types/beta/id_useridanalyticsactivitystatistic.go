package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAnalyticsActivityStatisticId{}

// UserIdAnalyticsActivityStatisticId is a struct representing the Resource ID for a User Id Analytics Activity Statistic
type UserIdAnalyticsActivityStatisticId struct {
	UserId               string
	ActivityStatisticsId string
}

// NewUserIdAnalyticsActivityStatisticID returns a new UserIdAnalyticsActivityStatisticId struct
func NewUserIdAnalyticsActivityStatisticID(userId string, activityStatisticsId string) UserIdAnalyticsActivityStatisticId {
	return UserIdAnalyticsActivityStatisticId{
		UserId:               userId,
		ActivityStatisticsId: activityStatisticsId,
	}
}

// ParseUserIdAnalyticsActivityStatisticID parses 'input' into a UserIdAnalyticsActivityStatisticId
func ParseUserIdAnalyticsActivityStatisticID(input string) (*UserIdAnalyticsActivityStatisticId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAnalyticsActivityStatisticId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAnalyticsActivityStatisticId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdAnalyticsActivityStatisticIDInsensitively parses 'input' case-insensitively into a UserIdAnalyticsActivityStatisticId
// note: this method should only be used for API response data and not user input
func ParseUserIdAnalyticsActivityStatisticIDInsensitively(input string) (*UserIdAnalyticsActivityStatisticId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAnalyticsActivityStatisticId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAnalyticsActivityStatisticId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdAnalyticsActivityStatisticId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ActivityStatisticsId, ok = input.Parsed["activityStatisticsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "activityStatisticsId", input)
	}

	return nil
}

// ValidateUserIdAnalyticsActivityStatisticID checks that 'input' can be parsed as a User Id Analytics Activity Statistic ID
func ValidateUserIdAnalyticsActivityStatisticID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdAnalyticsActivityStatisticID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Analytics Activity Statistic ID
func (id UserIdAnalyticsActivityStatisticId) ID() string {
	fmtString := "/users/%s/analytics/activityStatistics/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ActivityStatisticsId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Analytics Activity Statistic ID
func (id UserIdAnalyticsActivityStatisticId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("analytics", "analytics", "analytics"),
		resourceids.StaticSegment("activityStatistics", "activityStatistics", "activityStatistics"),
		resourceids.UserSpecifiedSegment("activityStatisticsId", "activityStatisticsId"),
	}
}

// String returns a human-readable description of this User Id Analytics Activity Statistic ID
func (id UserIdAnalyticsActivityStatisticId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Activity Statistics: %q", id.ActivityStatisticsId),
	}
	return fmt.Sprintf("User Id Analytics Activity Statistic (%s)", strings.Join(components, "\n"))
}
