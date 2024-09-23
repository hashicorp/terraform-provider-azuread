package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightMonthlySummaryId{}

// ReportUserInsightMonthlySummaryId is a struct representing the Resource ID for a Report User Insight Monthly Summary
type ReportUserInsightMonthlySummaryId struct {
	InsightSummaryId string
}

// NewReportUserInsightMonthlySummaryID returns a new ReportUserInsightMonthlySummaryId struct
func NewReportUserInsightMonthlySummaryID(insightSummaryId string) ReportUserInsightMonthlySummaryId {
	return ReportUserInsightMonthlySummaryId{
		InsightSummaryId: insightSummaryId,
	}
}

// ParseReportUserInsightMonthlySummaryID parses 'input' into a ReportUserInsightMonthlySummaryId
func ParseReportUserInsightMonthlySummaryID(input string) (*ReportUserInsightMonthlySummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightMonthlySummaryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightMonthlySummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportUserInsightMonthlySummaryIDInsensitively parses 'input' case-insensitively into a ReportUserInsightMonthlySummaryId
// note: this method should only be used for API response data and not user input
func ParseReportUserInsightMonthlySummaryIDInsensitively(input string) (*ReportUserInsightMonthlySummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightMonthlySummaryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightMonthlySummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportUserInsightMonthlySummaryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.InsightSummaryId, ok = input.Parsed["insightSummaryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "insightSummaryId", input)
	}

	return nil
}

// ValidateReportUserInsightMonthlySummaryID checks that 'input' can be parsed as a Report User Insight Monthly Summary ID
func ValidateReportUserInsightMonthlySummaryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportUserInsightMonthlySummaryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report User Insight Monthly Summary ID
func (id ReportUserInsightMonthlySummaryId) ID() string {
	fmtString := "/reports/userInsights/monthly/summary/%s"
	return fmt.Sprintf(fmtString, id.InsightSummaryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report User Insight Monthly Summary ID
func (id ReportUserInsightMonthlySummaryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("userInsights", "userInsights", "userInsights"),
		resourceids.StaticSegment("monthly", "monthly", "monthly"),
		resourceids.StaticSegment("summary", "summary", "summary"),
		resourceids.UserSpecifiedSegment("insightSummaryId", "insightSummaryId"),
	}
}

// String returns a human-readable description of this Report User Insight Monthly Summary ID
func (id ReportUserInsightMonthlySummaryId) String() string {
	components := []string{
		fmt.Sprintf("Insight Summary: %q", id.InsightSummaryId),
	}
	return fmt.Sprintf("Report User Insight Monthly Summary (%s)", strings.Join(components, "\n"))
}
