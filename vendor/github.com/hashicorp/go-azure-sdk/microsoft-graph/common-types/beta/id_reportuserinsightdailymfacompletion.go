package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightDailyMfaCompletionId{}

// ReportUserInsightDailyMfaCompletionId is a struct representing the Resource ID for a Report User Insight Daily Mfa Completion
type ReportUserInsightDailyMfaCompletionId struct {
	MfaCompletionMetricId string
}

// NewReportUserInsightDailyMfaCompletionID returns a new ReportUserInsightDailyMfaCompletionId struct
func NewReportUserInsightDailyMfaCompletionID(mfaCompletionMetricId string) ReportUserInsightDailyMfaCompletionId {
	return ReportUserInsightDailyMfaCompletionId{
		MfaCompletionMetricId: mfaCompletionMetricId,
	}
}

// ParseReportUserInsightDailyMfaCompletionID parses 'input' into a ReportUserInsightDailyMfaCompletionId
func ParseReportUserInsightDailyMfaCompletionID(input string) (*ReportUserInsightDailyMfaCompletionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightDailyMfaCompletionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightDailyMfaCompletionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportUserInsightDailyMfaCompletionIDInsensitively parses 'input' case-insensitively into a ReportUserInsightDailyMfaCompletionId
// note: this method should only be used for API response data and not user input
func ParseReportUserInsightDailyMfaCompletionIDInsensitively(input string) (*ReportUserInsightDailyMfaCompletionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightDailyMfaCompletionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightDailyMfaCompletionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportUserInsightDailyMfaCompletionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MfaCompletionMetricId, ok = input.Parsed["mfaCompletionMetricId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mfaCompletionMetricId", input)
	}

	return nil
}

// ValidateReportUserInsightDailyMfaCompletionID checks that 'input' can be parsed as a Report User Insight Daily Mfa Completion ID
func ValidateReportUserInsightDailyMfaCompletionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportUserInsightDailyMfaCompletionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report User Insight Daily Mfa Completion ID
func (id ReportUserInsightDailyMfaCompletionId) ID() string {
	fmtString := "/reports/userInsights/daily/mfaCompletions/%s"
	return fmt.Sprintf(fmtString, id.MfaCompletionMetricId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report User Insight Daily Mfa Completion ID
func (id ReportUserInsightDailyMfaCompletionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("userInsights", "userInsights", "userInsights"),
		resourceids.StaticSegment("daily", "daily", "daily"),
		resourceids.StaticSegment("mfaCompletions", "mfaCompletions", "mfaCompletions"),
		resourceids.UserSpecifiedSegment("mfaCompletionMetricId", "mfaCompletionMetricId"),
	}
}

// String returns a human-readable description of this Report User Insight Daily Mfa Completion ID
func (id ReportUserInsightDailyMfaCompletionId) String() string {
	components := []string{
		fmt.Sprintf("Mfa Completion Metric: %q", id.MfaCompletionMetricId),
	}
	return fmt.Sprintf("Report User Insight Daily Mfa Completion (%s)", strings.Join(components, "\n"))
}
