package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightDailyMfaTelecomFraudId{}

// ReportUserInsightDailyMfaTelecomFraudId is a struct representing the Resource ID for a Report User Insight Daily Mfa Telecom Fraud
type ReportUserInsightDailyMfaTelecomFraudId struct {
	MfaTelecomFraudMetricId string
}

// NewReportUserInsightDailyMfaTelecomFraudID returns a new ReportUserInsightDailyMfaTelecomFraudId struct
func NewReportUserInsightDailyMfaTelecomFraudID(mfaTelecomFraudMetricId string) ReportUserInsightDailyMfaTelecomFraudId {
	return ReportUserInsightDailyMfaTelecomFraudId{
		MfaTelecomFraudMetricId: mfaTelecomFraudMetricId,
	}
}

// ParseReportUserInsightDailyMfaTelecomFraudID parses 'input' into a ReportUserInsightDailyMfaTelecomFraudId
func ParseReportUserInsightDailyMfaTelecomFraudID(input string) (*ReportUserInsightDailyMfaTelecomFraudId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightDailyMfaTelecomFraudId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightDailyMfaTelecomFraudId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportUserInsightDailyMfaTelecomFraudIDInsensitively parses 'input' case-insensitively into a ReportUserInsightDailyMfaTelecomFraudId
// note: this method should only be used for API response data and not user input
func ParseReportUserInsightDailyMfaTelecomFraudIDInsensitively(input string) (*ReportUserInsightDailyMfaTelecomFraudId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightDailyMfaTelecomFraudId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightDailyMfaTelecomFraudId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportUserInsightDailyMfaTelecomFraudId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MfaTelecomFraudMetricId, ok = input.Parsed["mfaTelecomFraudMetricId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mfaTelecomFraudMetricId", input)
	}

	return nil
}

// ValidateReportUserInsightDailyMfaTelecomFraudID checks that 'input' can be parsed as a Report User Insight Daily Mfa Telecom Fraud ID
func ValidateReportUserInsightDailyMfaTelecomFraudID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportUserInsightDailyMfaTelecomFraudID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report User Insight Daily Mfa Telecom Fraud ID
func (id ReportUserInsightDailyMfaTelecomFraudId) ID() string {
	fmtString := "/reports/userInsights/daily/mfaTelecomFraud/%s"
	return fmt.Sprintf(fmtString, id.MfaTelecomFraudMetricId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report User Insight Daily Mfa Telecom Fraud ID
func (id ReportUserInsightDailyMfaTelecomFraudId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("userInsights", "userInsights", "userInsights"),
		resourceids.StaticSegment("daily", "daily", "daily"),
		resourceids.StaticSegment("mfaTelecomFraud", "mfaTelecomFraud", "mfaTelecomFraud"),
		resourceids.UserSpecifiedSegment("mfaTelecomFraudMetricId", "mfaTelecomFraudMetricId"),
	}
}

// String returns a human-readable description of this Report User Insight Daily Mfa Telecom Fraud ID
func (id ReportUserInsightDailyMfaTelecomFraudId) String() string {
	components := []string{
		fmt.Sprintf("Mfa Telecom Fraud Metric: %q", id.MfaTelecomFraudMetricId),
	}
	return fmt.Sprintf("Report User Insight Daily Mfa Telecom Fraud (%s)", strings.Join(components, "\n"))
}
