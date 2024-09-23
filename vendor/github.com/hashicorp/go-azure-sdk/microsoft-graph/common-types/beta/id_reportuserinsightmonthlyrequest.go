package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightMonthlyRequestId{}

// ReportUserInsightMonthlyRequestId is a struct representing the Resource ID for a Report User Insight Monthly Request
type ReportUserInsightMonthlyRequestId struct {
	UserRequestsMetricId string
}

// NewReportUserInsightMonthlyRequestID returns a new ReportUserInsightMonthlyRequestId struct
func NewReportUserInsightMonthlyRequestID(userRequestsMetricId string) ReportUserInsightMonthlyRequestId {
	return ReportUserInsightMonthlyRequestId{
		UserRequestsMetricId: userRequestsMetricId,
	}
}

// ParseReportUserInsightMonthlyRequestID parses 'input' into a ReportUserInsightMonthlyRequestId
func ParseReportUserInsightMonthlyRequestID(input string) (*ReportUserInsightMonthlyRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightMonthlyRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightMonthlyRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportUserInsightMonthlyRequestIDInsensitively parses 'input' case-insensitively into a ReportUserInsightMonthlyRequestId
// note: this method should only be used for API response data and not user input
func ParseReportUserInsightMonthlyRequestIDInsensitively(input string) (*ReportUserInsightMonthlyRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightMonthlyRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightMonthlyRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportUserInsightMonthlyRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserRequestsMetricId, ok = input.Parsed["userRequestsMetricId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userRequestsMetricId", input)
	}

	return nil
}

// ValidateReportUserInsightMonthlyRequestID checks that 'input' can be parsed as a Report User Insight Monthly Request ID
func ValidateReportUserInsightMonthlyRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportUserInsightMonthlyRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report User Insight Monthly Request ID
func (id ReportUserInsightMonthlyRequestId) ID() string {
	fmtString := "/reports/userInsights/monthly/requests/%s"
	return fmt.Sprintf(fmtString, id.UserRequestsMetricId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report User Insight Monthly Request ID
func (id ReportUserInsightMonthlyRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("userInsights", "userInsights", "userInsights"),
		resourceids.StaticSegment("monthly", "monthly", "monthly"),
		resourceids.StaticSegment("requests", "requests", "requests"),
		resourceids.UserSpecifiedSegment("userRequestsMetricId", "userRequestsMetricId"),
	}
}

// String returns a human-readable description of this Report User Insight Monthly Request ID
func (id ReportUserInsightMonthlyRequestId) String() string {
	components := []string{
		fmt.Sprintf("User Requests Metric: %q", id.UserRequestsMetricId),
	}
	return fmt.Sprintf("Report User Insight Monthly Request (%s)", strings.Join(components, "\n"))
}
