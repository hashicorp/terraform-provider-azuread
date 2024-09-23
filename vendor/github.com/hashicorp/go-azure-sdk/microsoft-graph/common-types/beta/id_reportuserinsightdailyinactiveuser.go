package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightDailyInactiveUserId{}

// ReportUserInsightDailyInactiveUserId is a struct representing the Resource ID for a Report User Insight Daily Inactive User
type ReportUserInsightDailyInactiveUserId struct {
	DailyInactiveUsersMetricId string
}

// NewReportUserInsightDailyInactiveUserID returns a new ReportUserInsightDailyInactiveUserId struct
func NewReportUserInsightDailyInactiveUserID(dailyInactiveUsersMetricId string) ReportUserInsightDailyInactiveUserId {
	return ReportUserInsightDailyInactiveUserId{
		DailyInactiveUsersMetricId: dailyInactiveUsersMetricId,
	}
}

// ParseReportUserInsightDailyInactiveUserID parses 'input' into a ReportUserInsightDailyInactiveUserId
func ParseReportUserInsightDailyInactiveUserID(input string) (*ReportUserInsightDailyInactiveUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightDailyInactiveUserId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightDailyInactiveUserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportUserInsightDailyInactiveUserIDInsensitively parses 'input' case-insensitively into a ReportUserInsightDailyInactiveUserId
// note: this method should only be used for API response data and not user input
func ParseReportUserInsightDailyInactiveUserIDInsensitively(input string) (*ReportUserInsightDailyInactiveUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightDailyInactiveUserId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightDailyInactiveUserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportUserInsightDailyInactiveUserId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DailyInactiveUsersMetricId, ok = input.Parsed["dailyInactiveUsersMetricId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dailyInactiveUsersMetricId", input)
	}

	return nil
}

// ValidateReportUserInsightDailyInactiveUserID checks that 'input' can be parsed as a Report User Insight Daily Inactive User ID
func ValidateReportUserInsightDailyInactiveUserID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportUserInsightDailyInactiveUserID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report User Insight Daily Inactive User ID
func (id ReportUserInsightDailyInactiveUserId) ID() string {
	fmtString := "/reports/userInsights/daily/inactiveUsers/%s"
	return fmt.Sprintf(fmtString, id.DailyInactiveUsersMetricId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report User Insight Daily Inactive User ID
func (id ReportUserInsightDailyInactiveUserId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("userInsights", "userInsights", "userInsights"),
		resourceids.StaticSegment("daily", "daily", "daily"),
		resourceids.StaticSegment("inactiveUsers", "inactiveUsers", "inactiveUsers"),
		resourceids.UserSpecifiedSegment("dailyInactiveUsersMetricId", "dailyInactiveUsersMetricId"),
	}
}

// String returns a human-readable description of this Report User Insight Daily Inactive User ID
func (id ReportUserInsightDailyInactiveUserId) String() string {
	components := []string{
		fmt.Sprintf("Daily Inactive Users Metric: %q", id.DailyInactiveUsersMetricId),
	}
	return fmt.Sprintf("Report User Insight Daily Inactive User (%s)", strings.Join(components, "\n"))
}
