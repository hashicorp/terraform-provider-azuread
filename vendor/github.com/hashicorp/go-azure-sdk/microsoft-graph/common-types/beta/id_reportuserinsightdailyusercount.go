package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightDailyUserCountId{}

// ReportUserInsightDailyUserCountId is a struct representing the Resource ID for a Report User Insight Daily User Count
type ReportUserInsightDailyUserCountId struct {
	UserCountMetricId string
}

// NewReportUserInsightDailyUserCountID returns a new ReportUserInsightDailyUserCountId struct
func NewReportUserInsightDailyUserCountID(userCountMetricId string) ReportUserInsightDailyUserCountId {
	return ReportUserInsightDailyUserCountId{
		UserCountMetricId: userCountMetricId,
	}
}

// ParseReportUserInsightDailyUserCountID parses 'input' into a ReportUserInsightDailyUserCountId
func ParseReportUserInsightDailyUserCountID(input string) (*ReportUserInsightDailyUserCountId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightDailyUserCountId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightDailyUserCountId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportUserInsightDailyUserCountIDInsensitively parses 'input' case-insensitively into a ReportUserInsightDailyUserCountId
// note: this method should only be used for API response data and not user input
func ParseReportUserInsightDailyUserCountIDInsensitively(input string) (*ReportUserInsightDailyUserCountId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightDailyUserCountId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightDailyUserCountId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportUserInsightDailyUserCountId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserCountMetricId, ok = input.Parsed["userCountMetricId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userCountMetricId", input)
	}

	return nil
}

// ValidateReportUserInsightDailyUserCountID checks that 'input' can be parsed as a Report User Insight Daily User Count ID
func ValidateReportUserInsightDailyUserCountID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportUserInsightDailyUserCountID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report User Insight Daily User Count ID
func (id ReportUserInsightDailyUserCountId) ID() string {
	fmtString := "/reports/userInsights/daily/userCount/%s"
	return fmt.Sprintf(fmtString, id.UserCountMetricId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report User Insight Daily User Count ID
func (id ReportUserInsightDailyUserCountId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("userInsights", "userInsights", "userInsights"),
		resourceids.StaticSegment("daily", "daily", "daily"),
		resourceids.StaticSegment("userCount", "userCount", "userCount"),
		resourceids.UserSpecifiedSegment("userCountMetricId", "userCountMetricId"),
	}
}

// String returns a human-readable description of this Report User Insight Daily User Count ID
func (id ReportUserInsightDailyUserCountId) String() string {
	components := []string{
		fmt.Sprintf("User Count Metric: %q", id.UserCountMetricId),
	}
	return fmt.Sprintf("Report User Insight Daily User Count (%s)", strings.Join(components, "\n"))
}
