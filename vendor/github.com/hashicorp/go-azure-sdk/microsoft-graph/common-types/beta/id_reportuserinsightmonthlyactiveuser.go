package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightMonthlyActiveUserId{}

// ReportUserInsightMonthlyActiveUserId is a struct representing the Resource ID for a Report User Insight Monthly Active User
type ReportUserInsightMonthlyActiveUserId struct {
	ActiveUsersMetricId string
}

// NewReportUserInsightMonthlyActiveUserID returns a new ReportUserInsightMonthlyActiveUserId struct
func NewReportUserInsightMonthlyActiveUserID(activeUsersMetricId string) ReportUserInsightMonthlyActiveUserId {
	return ReportUserInsightMonthlyActiveUserId{
		ActiveUsersMetricId: activeUsersMetricId,
	}
}

// ParseReportUserInsightMonthlyActiveUserID parses 'input' into a ReportUserInsightMonthlyActiveUserId
func ParseReportUserInsightMonthlyActiveUserID(input string) (*ReportUserInsightMonthlyActiveUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightMonthlyActiveUserId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightMonthlyActiveUserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportUserInsightMonthlyActiveUserIDInsensitively parses 'input' case-insensitively into a ReportUserInsightMonthlyActiveUserId
// note: this method should only be used for API response data and not user input
func ParseReportUserInsightMonthlyActiveUserIDInsensitively(input string) (*ReportUserInsightMonthlyActiveUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightMonthlyActiveUserId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightMonthlyActiveUserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportUserInsightMonthlyActiveUserId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ActiveUsersMetricId, ok = input.Parsed["activeUsersMetricId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "activeUsersMetricId", input)
	}

	return nil
}

// ValidateReportUserInsightMonthlyActiveUserID checks that 'input' can be parsed as a Report User Insight Monthly Active User ID
func ValidateReportUserInsightMonthlyActiveUserID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportUserInsightMonthlyActiveUserID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report User Insight Monthly Active User ID
func (id ReportUserInsightMonthlyActiveUserId) ID() string {
	fmtString := "/reports/userInsights/monthly/activeUsers/%s"
	return fmt.Sprintf(fmtString, id.ActiveUsersMetricId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report User Insight Monthly Active User ID
func (id ReportUserInsightMonthlyActiveUserId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("userInsights", "userInsights", "userInsights"),
		resourceids.StaticSegment("monthly", "monthly", "monthly"),
		resourceids.StaticSegment("activeUsers", "activeUsers", "activeUsers"),
		resourceids.UserSpecifiedSegment("activeUsersMetricId", "activeUsersMetricId"),
	}
}

// String returns a human-readable description of this Report User Insight Monthly Active User ID
func (id ReportUserInsightMonthlyActiveUserId) String() string {
	components := []string{
		fmt.Sprintf("Active Users Metric: %q", id.ActiveUsersMetricId),
	}
	return fmt.Sprintf("Report User Insight Monthly Active User (%s)", strings.Join(components, "\n"))
}
