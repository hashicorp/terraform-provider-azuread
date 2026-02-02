package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightMonthlyMfaCompletionId{}

// ReportUserInsightMonthlyMfaCompletionId is a struct representing the Resource ID for a Report User Insight Monthly Mfa Completion
type ReportUserInsightMonthlyMfaCompletionId struct {
	MfaCompletionMetricId string
}

// NewReportUserInsightMonthlyMfaCompletionID returns a new ReportUserInsightMonthlyMfaCompletionId struct
func NewReportUserInsightMonthlyMfaCompletionID(mfaCompletionMetricId string) ReportUserInsightMonthlyMfaCompletionId {
	return ReportUserInsightMonthlyMfaCompletionId{
		MfaCompletionMetricId: mfaCompletionMetricId,
	}
}

// ParseReportUserInsightMonthlyMfaCompletionID parses 'input' into a ReportUserInsightMonthlyMfaCompletionId
func ParseReportUserInsightMonthlyMfaCompletionID(input string) (*ReportUserInsightMonthlyMfaCompletionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightMonthlyMfaCompletionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightMonthlyMfaCompletionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportUserInsightMonthlyMfaCompletionIDInsensitively parses 'input' case-insensitively into a ReportUserInsightMonthlyMfaCompletionId
// note: this method should only be used for API response data and not user input
func ParseReportUserInsightMonthlyMfaCompletionIDInsensitively(input string) (*ReportUserInsightMonthlyMfaCompletionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightMonthlyMfaCompletionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightMonthlyMfaCompletionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportUserInsightMonthlyMfaCompletionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MfaCompletionMetricId, ok = input.Parsed["mfaCompletionMetricId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mfaCompletionMetricId", input)
	}

	return nil
}

// ValidateReportUserInsightMonthlyMfaCompletionID checks that 'input' can be parsed as a Report User Insight Monthly Mfa Completion ID
func ValidateReportUserInsightMonthlyMfaCompletionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportUserInsightMonthlyMfaCompletionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report User Insight Monthly Mfa Completion ID
func (id ReportUserInsightMonthlyMfaCompletionId) ID() string {
	fmtString := "/reports/userInsights/monthly/mfaCompletions/%s"
	return fmt.Sprintf(fmtString, id.MfaCompletionMetricId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report User Insight Monthly Mfa Completion ID
func (id ReportUserInsightMonthlyMfaCompletionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("userInsights", "userInsights", "userInsights"),
		resourceids.StaticSegment("monthly", "monthly", "monthly"),
		resourceids.StaticSegment("mfaCompletions", "mfaCompletions", "mfaCompletions"),
		resourceids.UserSpecifiedSegment("mfaCompletionMetricId", "mfaCompletionMetricId"),
	}
}

// String returns a human-readable description of this Report User Insight Monthly Mfa Completion ID
func (id ReportUserInsightMonthlyMfaCompletionId) String() string {
	components := []string{
		fmt.Sprintf("Mfa Completion Metric: %q", id.MfaCompletionMetricId),
	}
	return fmt.Sprintf("Report User Insight Monthly Mfa Completion (%s)", strings.Join(components, "\n"))
}
