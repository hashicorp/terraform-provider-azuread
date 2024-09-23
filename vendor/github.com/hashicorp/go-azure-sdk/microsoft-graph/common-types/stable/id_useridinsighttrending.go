package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdInsightTrendingId{}

// UserIdInsightTrendingId is a struct representing the Resource ID for a User Id Insight Trending
type UserIdInsightTrendingId struct {
	UserId     string
	TrendingId string
}

// NewUserIdInsightTrendingID returns a new UserIdInsightTrendingId struct
func NewUserIdInsightTrendingID(userId string, trendingId string) UserIdInsightTrendingId {
	return UserIdInsightTrendingId{
		UserId:     userId,
		TrendingId: trendingId,
	}
}

// ParseUserIdInsightTrendingID parses 'input' into a UserIdInsightTrendingId
func ParseUserIdInsightTrendingID(input string) (*UserIdInsightTrendingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdInsightTrendingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdInsightTrendingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdInsightTrendingIDInsensitively parses 'input' case-insensitively into a UserIdInsightTrendingId
// note: this method should only be used for API response data and not user input
func ParseUserIdInsightTrendingIDInsensitively(input string) (*UserIdInsightTrendingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdInsightTrendingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdInsightTrendingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdInsightTrendingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TrendingId, ok = input.Parsed["trendingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "trendingId", input)
	}

	return nil
}

// ValidateUserIdInsightTrendingID checks that 'input' can be parsed as a User Id Insight Trending ID
func ValidateUserIdInsightTrendingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdInsightTrendingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Insight Trending ID
func (id UserIdInsightTrendingId) ID() string {
	fmtString := "/users/%s/insights/trending/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TrendingId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Insight Trending ID
func (id UserIdInsightTrendingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("insights", "insights", "insights"),
		resourceids.StaticSegment("trending", "trending", "trending"),
		resourceids.UserSpecifiedSegment("trendingId", "trendingId"),
	}
}

// String returns a human-readable description of this User Id Insight Trending ID
func (id UserIdInsightTrendingId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Trending: %q", id.TrendingId),
	}
	return fmt.Sprintf("User Id Insight Trending (%s)", strings.Join(components, "\n"))
}
