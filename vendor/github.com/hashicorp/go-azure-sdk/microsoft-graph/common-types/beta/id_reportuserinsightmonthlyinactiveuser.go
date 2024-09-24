package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightMonthlyInactiveUserId{}

// ReportUserInsightMonthlyInactiveUserId is a struct representing the Resource ID for a Report User Insight Monthly Inactive User
type ReportUserInsightMonthlyInactiveUserId struct {
	MonthlyInactiveUsersMetricId string
}

// NewReportUserInsightMonthlyInactiveUserID returns a new ReportUserInsightMonthlyInactiveUserId struct
func NewReportUserInsightMonthlyInactiveUserID(monthlyInactiveUsersMetricId string) ReportUserInsightMonthlyInactiveUserId {
	return ReportUserInsightMonthlyInactiveUserId{
		MonthlyInactiveUsersMetricId: monthlyInactiveUsersMetricId,
	}
}

// ParseReportUserInsightMonthlyInactiveUserID parses 'input' into a ReportUserInsightMonthlyInactiveUserId
func ParseReportUserInsightMonthlyInactiveUserID(input string) (*ReportUserInsightMonthlyInactiveUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightMonthlyInactiveUserId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightMonthlyInactiveUserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportUserInsightMonthlyInactiveUserIDInsensitively parses 'input' case-insensitively into a ReportUserInsightMonthlyInactiveUserId
// note: this method should only be used for API response data and not user input
func ParseReportUserInsightMonthlyInactiveUserIDInsensitively(input string) (*ReportUserInsightMonthlyInactiveUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightMonthlyInactiveUserId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightMonthlyInactiveUserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportUserInsightMonthlyInactiveUserId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MonthlyInactiveUsersMetricId, ok = input.Parsed["monthlyInactiveUsersMetricId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "monthlyInactiveUsersMetricId", input)
	}

	return nil
}

// ValidateReportUserInsightMonthlyInactiveUserID checks that 'input' can be parsed as a Report User Insight Monthly Inactive User ID
func ValidateReportUserInsightMonthlyInactiveUserID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportUserInsightMonthlyInactiveUserID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report User Insight Monthly Inactive User ID
func (id ReportUserInsightMonthlyInactiveUserId) ID() string {
	fmtString := "/reports/userInsights/monthly/inactiveUsers/%s"
	return fmt.Sprintf(fmtString, id.MonthlyInactiveUsersMetricId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report User Insight Monthly Inactive User ID
func (id ReportUserInsightMonthlyInactiveUserId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("userInsights", "userInsights", "userInsights"),
		resourceids.StaticSegment("monthly", "monthly", "monthly"),
		resourceids.StaticSegment("inactiveUsers", "inactiveUsers", "inactiveUsers"),
		resourceids.UserSpecifiedSegment("monthlyInactiveUsersMetricId", "monthlyInactiveUsersMetricId"),
	}
}

// String returns a human-readable description of this Report User Insight Monthly Inactive User ID
func (id ReportUserInsightMonthlyInactiveUserId) String() string {
	components := []string{
		fmt.Sprintf("Monthly Inactive Users Metric: %q", id.MonthlyInactiveUsersMetricId),
	}
	return fmt.Sprintf("Report User Insight Monthly Inactive User (%s)", strings.Join(components, "\n"))
}
