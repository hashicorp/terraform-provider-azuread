package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightDailySummaryId{}

// ReportUserInsightDailySummaryId is a struct representing the Resource ID for a Report User Insight Daily Summary
type ReportUserInsightDailySummaryId struct {
	InsightSummaryId string
}

// NewReportUserInsightDailySummaryID returns a new ReportUserInsightDailySummaryId struct
func NewReportUserInsightDailySummaryID(insightSummaryId string) ReportUserInsightDailySummaryId {
	return ReportUserInsightDailySummaryId{
		InsightSummaryId: insightSummaryId,
	}
}

// ParseReportUserInsightDailySummaryID parses 'input' into a ReportUserInsightDailySummaryId
func ParseReportUserInsightDailySummaryID(input string) (*ReportUserInsightDailySummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightDailySummaryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightDailySummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportUserInsightDailySummaryIDInsensitively parses 'input' case-insensitively into a ReportUserInsightDailySummaryId
// note: this method should only be used for API response data and not user input
func ParseReportUserInsightDailySummaryIDInsensitively(input string) (*ReportUserInsightDailySummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightDailySummaryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightDailySummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportUserInsightDailySummaryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.InsightSummaryId, ok = input.Parsed["insightSummaryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "insightSummaryId", input)
	}

	return nil
}

// ValidateReportUserInsightDailySummaryID checks that 'input' can be parsed as a Report User Insight Daily Summary ID
func ValidateReportUserInsightDailySummaryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportUserInsightDailySummaryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report User Insight Daily Summary ID
func (id ReportUserInsightDailySummaryId) ID() string {
	fmtString := "/reports/userInsights/daily/summary/%s"
	return fmt.Sprintf(fmtString, id.InsightSummaryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report User Insight Daily Summary ID
func (id ReportUserInsightDailySummaryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("userInsights", "userInsights", "userInsights"),
		resourceids.StaticSegment("daily", "daily", "daily"),
		resourceids.StaticSegment("summary", "summary", "summary"),
		resourceids.UserSpecifiedSegment("insightSummaryId", "insightSummaryId"),
	}
}

// String returns a human-readable description of this Report User Insight Daily Summary ID
func (id ReportUserInsightDailySummaryId) String() string {
	components := []string{
		fmt.Sprintf("Insight Summary: %q", id.InsightSummaryId),
	}
	return fmt.Sprintf("Report User Insight Daily Summary (%s)", strings.Join(components, "\n"))
}
