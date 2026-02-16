package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightMonthlyInactiveUsersByApplicationId{}

// ReportUserInsightMonthlyInactiveUsersByApplicationId is a struct representing the Resource ID for a Report User Insight Monthly Inactive Users By Application
type ReportUserInsightMonthlyInactiveUsersByApplicationId struct {
	MonthlyInactiveUsersByApplicationMetricId string
}

// NewReportUserInsightMonthlyInactiveUsersByApplicationID returns a new ReportUserInsightMonthlyInactiveUsersByApplicationId struct
func NewReportUserInsightMonthlyInactiveUsersByApplicationID(monthlyInactiveUsersByApplicationMetricId string) ReportUserInsightMonthlyInactiveUsersByApplicationId {
	return ReportUserInsightMonthlyInactiveUsersByApplicationId{
		MonthlyInactiveUsersByApplicationMetricId: monthlyInactiveUsersByApplicationMetricId,
	}
}

// ParseReportUserInsightMonthlyInactiveUsersByApplicationID parses 'input' into a ReportUserInsightMonthlyInactiveUsersByApplicationId
func ParseReportUserInsightMonthlyInactiveUsersByApplicationID(input string) (*ReportUserInsightMonthlyInactiveUsersByApplicationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightMonthlyInactiveUsersByApplicationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightMonthlyInactiveUsersByApplicationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportUserInsightMonthlyInactiveUsersByApplicationIDInsensitively parses 'input' case-insensitively into a ReportUserInsightMonthlyInactiveUsersByApplicationId
// note: this method should only be used for API response data and not user input
func ParseReportUserInsightMonthlyInactiveUsersByApplicationIDInsensitively(input string) (*ReportUserInsightMonthlyInactiveUsersByApplicationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightMonthlyInactiveUsersByApplicationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightMonthlyInactiveUsersByApplicationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportUserInsightMonthlyInactiveUsersByApplicationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MonthlyInactiveUsersByApplicationMetricId, ok = input.Parsed["monthlyInactiveUsersByApplicationMetricId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "monthlyInactiveUsersByApplicationMetricId", input)
	}

	return nil
}

// ValidateReportUserInsightMonthlyInactiveUsersByApplicationID checks that 'input' can be parsed as a Report User Insight Monthly Inactive Users By Application ID
func ValidateReportUserInsightMonthlyInactiveUsersByApplicationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportUserInsightMonthlyInactiveUsersByApplicationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report User Insight Monthly Inactive Users By Application ID
func (id ReportUserInsightMonthlyInactiveUsersByApplicationId) ID() string {
	fmtString := "/reports/userInsights/monthly/inactiveUsersByApplication/%s"
	return fmt.Sprintf(fmtString, id.MonthlyInactiveUsersByApplicationMetricId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report User Insight Monthly Inactive Users By Application ID
func (id ReportUserInsightMonthlyInactiveUsersByApplicationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("userInsights", "userInsights", "userInsights"),
		resourceids.StaticSegment("monthly", "monthly", "monthly"),
		resourceids.StaticSegment("inactiveUsersByApplication", "inactiveUsersByApplication", "inactiveUsersByApplication"),
		resourceids.UserSpecifiedSegment("monthlyInactiveUsersByApplicationMetricId", "monthlyInactiveUsersByApplicationMetricId"),
	}
}

// String returns a human-readable description of this Report User Insight Monthly Inactive Users By Application ID
func (id ReportUserInsightMonthlyInactiveUsersByApplicationId) String() string {
	components := []string{
		fmt.Sprintf("Monthly Inactive Users By Application Metric: %q", id.MonthlyInactiveUsersByApplicationMetricId),
	}
	return fmt.Sprintf("Report User Insight Monthly Inactive Users By Application (%s)", strings.Join(components, "\n"))
}
