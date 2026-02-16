package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeInsightTrendingId{}

// MeInsightTrendingId is a struct representing the Resource ID for a Me Insight Trending
type MeInsightTrendingId struct {
	TrendingId string
}

// NewMeInsightTrendingID returns a new MeInsightTrendingId struct
func NewMeInsightTrendingID(trendingId string) MeInsightTrendingId {
	return MeInsightTrendingId{
		TrendingId: trendingId,
	}
}

// ParseMeInsightTrendingID parses 'input' into a MeInsightTrendingId
func ParseMeInsightTrendingID(input string) (*MeInsightTrendingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeInsightTrendingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeInsightTrendingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeInsightTrendingIDInsensitively parses 'input' case-insensitively into a MeInsightTrendingId
// note: this method should only be used for API response data and not user input
func ParseMeInsightTrendingIDInsensitively(input string) (*MeInsightTrendingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeInsightTrendingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeInsightTrendingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeInsightTrendingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TrendingId, ok = input.Parsed["trendingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "trendingId", input)
	}

	return nil
}

// ValidateMeInsightTrendingID checks that 'input' can be parsed as a Me Insight Trending ID
func ValidateMeInsightTrendingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeInsightTrendingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Insight Trending ID
func (id MeInsightTrendingId) ID() string {
	fmtString := "/me/insights/trending/%s"
	return fmt.Sprintf(fmtString, id.TrendingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Insight Trending ID
func (id MeInsightTrendingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("insights", "insights", "insights"),
		resourceids.StaticSegment("trending", "trending", "trending"),
		resourceids.UserSpecifiedSegment("trendingId", "trendingId"),
	}
}

// String returns a human-readable description of this Me Insight Trending ID
func (id MeInsightTrendingId) String() string {
	components := []string{
		fmt.Sprintf("Trending: %q", id.TrendingId),
	}
	return fmt.Sprintf("Me Insight Trending (%s)", strings.Join(components, "\n"))
}
