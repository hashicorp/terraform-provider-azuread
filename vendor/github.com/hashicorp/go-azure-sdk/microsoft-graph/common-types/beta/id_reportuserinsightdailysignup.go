package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightDailySignUpId{}

// ReportUserInsightDailySignUpId is a struct representing the Resource ID for a Report User Insight Daily Sign Up
type ReportUserInsightDailySignUpId struct {
	UserSignUpMetricId string
}

// NewReportUserInsightDailySignUpID returns a new ReportUserInsightDailySignUpId struct
func NewReportUserInsightDailySignUpID(userSignUpMetricId string) ReportUserInsightDailySignUpId {
	return ReportUserInsightDailySignUpId{
		UserSignUpMetricId: userSignUpMetricId,
	}
}

// ParseReportUserInsightDailySignUpID parses 'input' into a ReportUserInsightDailySignUpId
func ParseReportUserInsightDailySignUpID(input string) (*ReportUserInsightDailySignUpId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightDailySignUpId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightDailySignUpId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportUserInsightDailySignUpIDInsensitively parses 'input' case-insensitively into a ReportUserInsightDailySignUpId
// note: this method should only be used for API response data and not user input
func ParseReportUserInsightDailySignUpIDInsensitively(input string) (*ReportUserInsightDailySignUpId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightDailySignUpId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightDailySignUpId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportUserInsightDailySignUpId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserSignUpMetricId, ok = input.Parsed["userSignUpMetricId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userSignUpMetricId", input)
	}

	return nil
}

// ValidateReportUserInsightDailySignUpID checks that 'input' can be parsed as a Report User Insight Daily Sign Up ID
func ValidateReportUserInsightDailySignUpID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportUserInsightDailySignUpID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report User Insight Daily Sign Up ID
func (id ReportUserInsightDailySignUpId) ID() string {
	fmtString := "/reports/userInsights/daily/signUps/%s"
	return fmt.Sprintf(fmtString, id.UserSignUpMetricId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report User Insight Daily Sign Up ID
func (id ReportUserInsightDailySignUpId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("userInsights", "userInsights", "userInsights"),
		resourceids.StaticSegment("daily", "daily", "daily"),
		resourceids.StaticSegment("signUps", "signUps", "signUps"),
		resourceids.UserSpecifiedSegment("userSignUpMetricId", "userSignUpMetricId"),
	}
}

// String returns a human-readable description of this Report User Insight Daily Sign Up ID
func (id ReportUserInsightDailySignUpId) String() string {
	components := []string{
		fmt.Sprintf("User Sign Up Metric: %q", id.UserSignUpMetricId),
	}
	return fmt.Sprintf("Report User Insight Daily Sign Up (%s)", strings.Join(components, "\n"))
}
