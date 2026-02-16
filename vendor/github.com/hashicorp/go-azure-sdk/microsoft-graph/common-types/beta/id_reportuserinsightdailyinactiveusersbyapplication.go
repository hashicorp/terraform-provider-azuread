package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightDailyInactiveUsersByApplicationId{}

// ReportUserInsightDailyInactiveUsersByApplicationId is a struct representing the Resource ID for a Report User Insight Daily Inactive Users By Application
type ReportUserInsightDailyInactiveUsersByApplicationId struct {
	DailyInactiveUsersByApplicationMetricId string
}

// NewReportUserInsightDailyInactiveUsersByApplicationID returns a new ReportUserInsightDailyInactiveUsersByApplicationId struct
func NewReportUserInsightDailyInactiveUsersByApplicationID(dailyInactiveUsersByApplicationMetricId string) ReportUserInsightDailyInactiveUsersByApplicationId {
	return ReportUserInsightDailyInactiveUsersByApplicationId{
		DailyInactiveUsersByApplicationMetricId: dailyInactiveUsersByApplicationMetricId,
	}
}

// ParseReportUserInsightDailyInactiveUsersByApplicationID parses 'input' into a ReportUserInsightDailyInactiveUsersByApplicationId
func ParseReportUserInsightDailyInactiveUsersByApplicationID(input string) (*ReportUserInsightDailyInactiveUsersByApplicationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightDailyInactiveUsersByApplicationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightDailyInactiveUsersByApplicationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportUserInsightDailyInactiveUsersByApplicationIDInsensitively parses 'input' case-insensitively into a ReportUserInsightDailyInactiveUsersByApplicationId
// note: this method should only be used for API response data and not user input
func ParseReportUserInsightDailyInactiveUsersByApplicationIDInsensitively(input string) (*ReportUserInsightDailyInactiveUsersByApplicationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightDailyInactiveUsersByApplicationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightDailyInactiveUsersByApplicationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportUserInsightDailyInactiveUsersByApplicationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DailyInactiveUsersByApplicationMetricId, ok = input.Parsed["dailyInactiveUsersByApplicationMetricId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dailyInactiveUsersByApplicationMetricId", input)
	}

	return nil
}

// ValidateReportUserInsightDailyInactiveUsersByApplicationID checks that 'input' can be parsed as a Report User Insight Daily Inactive Users By Application ID
func ValidateReportUserInsightDailyInactiveUsersByApplicationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportUserInsightDailyInactiveUsersByApplicationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report User Insight Daily Inactive Users By Application ID
func (id ReportUserInsightDailyInactiveUsersByApplicationId) ID() string {
	fmtString := "/reports/userInsights/daily/inactiveUsersByApplication/%s"
	return fmt.Sprintf(fmtString, id.DailyInactiveUsersByApplicationMetricId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report User Insight Daily Inactive Users By Application ID
func (id ReportUserInsightDailyInactiveUsersByApplicationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("userInsights", "userInsights", "userInsights"),
		resourceids.StaticSegment("daily", "daily", "daily"),
		resourceids.StaticSegment("inactiveUsersByApplication", "inactiveUsersByApplication", "inactiveUsersByApplication"),
		resourceids.UserSpecifiedSegment("dailyInactiveUsersByApplicationMetricId", "dailyInactiveUsersByApplicationMetricId"),
	}
}

// String returns a human-readable description of this Report User Insight Daily Inactive Users By Application ID
func (id ReportUserInsightDailyInactiveUsersByApplicationId) String() string {
	components := []string{
		fmt.Sprintf("Daily Inactive Users By Application Metric: %q", id.DailyInactiveUsersByApplicationMetricId),
	}
	return fmt.Sprintf("Report User Insight Daily Inactive Users By Application (%s)", strings.Join(components, "\n"))
}
