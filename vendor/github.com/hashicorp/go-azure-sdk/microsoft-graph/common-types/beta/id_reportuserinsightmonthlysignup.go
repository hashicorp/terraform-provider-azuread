package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightMonthlySignUpId{}

// ReportUserInsightMonthlySignUpId is a struct representing the Resource ID for a Report User Insight Monthly Sign Up
type ReportUserInsightMonthlySignUpId struct {
	UserSignUpMetricId string
}

// NewReportUserInsightMonthlySignUpID returns a new ReportUserInsightMonthlySignUpId struct
func NewReportUserInsightMonthlySignUpID(userSignUpMetricId string) ReportUserInsightMonthlySignUpId {
	return ReportUserInsightMonthlySignUpId{
		UserSignUpMetricId: userSignUpMetricId,
	}
}

// ParseReportUserInsightMonthlySignUpID parses 'input' into a ReportUserInsightMonthlySignUpId
func ParseReportUserInsightMonthlySignUpID(input string) (*ReportUserInsightMonthlySignUpId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightMonthlySignUpId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightMonthlySignUpId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportUserInsightMonthlySignUpIDInsensitively parses 'input' case-insensitively into a ReportUserInsightMonthlySignUpId
// note: this method should only be used for API response data and not user input
func ParseReportUserInsightMonthlySignUpIDInsensitively(input string) (*ReportUserInsightMonthlySignUpId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightMonthlySignUpId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightMonthlySignUpId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportUserInsightMonthlySignUpId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserSignUpMetricId, ok = input.Parsed["userSignUpMetricId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userSignUpMetricId", input)
	}

	return nil
}

// ValidateReportUserInsightMonthlySignUpID checks that 'input' can be parsed as a Report User Insight Monthly Sign Up ID
func ValidateReportUserInsightMonthlySignUpID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportUserInsightMonthlySignUpID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report User Insight Monthly Sign Up ID
func (id ReportUserInsightMonthlySignUpId) ID() string {
	fmtString := "/reports/userInsights/monthly/signUps/%s"
	return fmt.Sprintf(fmtString, id.UserSignUpMetricId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report User Insight Monthly Sign Up ID
func (id ReportUserInsightMonthlySignUpId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("userInsights", "userInsights", "userInsights"),
		resourceids.StaticSegment("monthly", "monthly", "monthly"),
		resourceids.StaticSegment("signUps", "signUps", "signUps"),
		resourceids.UserSpecifiedSegment("userSignUpMetricId", "userSignUpMetricId"),
	}
}

// String returns a human-readable description of this Report User Insight Monthly Sign Up ID
func (id ReportUserInsightMonthlySignUpId) String() string {
	components := []string{
		fmt.Sprintf("User Sign Up Metric: %q", id.UserSignUpMetricId),
	}
	return fmt.Sprintf("Report User Insight Monthly Sign Up (%s)", strings.Join(components, "\n"))
}
