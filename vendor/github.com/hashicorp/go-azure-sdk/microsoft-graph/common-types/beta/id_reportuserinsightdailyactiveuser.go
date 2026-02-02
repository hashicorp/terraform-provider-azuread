package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightDailyActiveUserId{}

// ReportUserInsightDailyActiveUserId is a struct representing the Resource ID for a Report User Insight Daily Active User
type ReportUserInsightDailyActiveUserId struct {
	ActiveUsersMetricId string
}

// NewReportUserInsightDailyActiveUserID returns a new ReportUserInsightDailyActiveUserId struct
func NewReportUserInsightDailyActiveUserID(activeUsersMetricId string) ReportUserInsightDailyActiveUserId {
	return ReportUserInsightDailyActiveUserId{
		ActiveUsersMetricId: activeUsersMetricId,
	}
}

// ParseReportUserInsightDailyActiveUserID parses 'input' into a ReportUserInsightDailyActiveUserId
func ParseReportUserInsightDailyActiveUserID(input string) (*ReportUserInsightDailyActiveUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightDailyActiveUserId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightDailyActiveUserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportUserInsightDailyActiveUserIDInsensitively parses 'input' case-insensitively into a ReportUserInsightDailyActiveUserId
// note: this method should only be used for API response data and not user input
func ParseReportUserInsightDailyActiveUserIDInsensitively(input string) (*ReportUserInsightDailyActiveUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightDailyActiveUserId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightDailyActiveUserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportUserInsightDailyActiveUserId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ActiveUsersMetricId, ok = input.Parsed["activeUsersMetricId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "activeUsersMetricId", input)
	}

	return nil
}

// ValidateReportUserInsightDailyActiveUserID checks that 'input' can be parsed as a Report User Insight Daily Active User ID
func ValidateReportUserInsightDailyActiveUserID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportUserInsightDailyActiveUserID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report User Insight Daily Active User ID
func (id ReportUserInsightDailyActiveUserId) ID() string {
	fmtString := "/reports/userInsights/daily/activeUsers/%s"
	return fmt.Sprintf(fmtString, id.ActiveUsersMetricId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report User Insight Daily Active User ID
func (id ReportUserInsightDailyActiveUserId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("userInsights", "userInsights", "userInsights"),
		resourceids.StaticSegment("daily", "daily", "daily"),
		resourceids.StaticSegment("activeUsers", "activeUsers", "activeUsers"),
		resourceids.UserSpecifiedSegment("activeUsersMetricId", "activeUsersMetricId"),
	}
}

// String returns a human-readable description of this Report User Insight Daily Active User ID
func (id ReportUserInsightDailyActiveUserId) String() string {
	components := []string{
		fmt.Sprintf("Active Users Metric: %q", id.ActiveUsersMetricId),
	}
	return fmt.Sprintf("Report User Insight Daily Active User (%s)", strings.Join(components, "\n"))
}
