package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightMonthlyMfaRegisteredUserId{}

// ReportUserInsightMonthlyMfaRegisteredUserId is a struct representing the Resource ID for a Report User Insight Monthly Mfa Registered User
type ReportUserInsightMonthlyMfaRegisteredUserId struct {
	MfaUserCountMetricId string
}

// NewReportUserInsightMonthlyMfaRegisteredUserID returns a new ReportUserInsightMonthlyMfaRegisteredUserId struct
func NewReportUserInsightMonthlyMfaRegisteredUserID(mfaUserCountMetricId string) ReportUserInsightMonthlyMfaRegisteredUserId {
	return ReportUserInsightMonthlyMfaRegisteredUserId{
		MfaUserCountMetricId: mfaUserCountMetricId,
	}
}

// ParseReportUserInsightMonthlyMfaRegisteredUserID parses 'input' into a ReportUserInsightMonthlyMfaRegisteredUserId
func ParseReportUserInsightMonthlyMfaRegisteredUserID(input string) (*ReportUserInsightMonthlyMfaRegisteredUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightMonthlyMfaRegisteredUserId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightMonthlyMfaRegisteredUserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportUserInsightMonthlyMfaRegisteredUserIDInsensitively parses 'input' case-insensitively into a ReportUserInsightMonthlyMfaRegisteredUserId
// note: this method should only be used for API response data and not user input
func ParseReportUserInsightMonthlyMfaRegisteredUserIDInsensitively(input string) (*ReportUserInsightMonthlyMfaRegisteredUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightMonthlyMfaRegisteredUserId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightMonthlyMfaRegisteredUserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportUserInsightMonthlyMfaRegisteredUserId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MfaUserCountMetricId, ok = input.Parsed["mfaUserCountMetricId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mfaUserCountMetricId", input)
	}

	return nil
}

// ValidateReportUserInsightMonthlyMfaRegisteredUserID checks that 'input' can be parsed as a Report User Insight Monthly Mfa Registered User ID
func ValidateReportUserInsightMonthlyMfaRegisteredUserID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportUserInsightMonthlyMfaRegisteredUserID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report User Insight Monthly Mfa Registered User ID
func (id ReportUserInsightMonthlyMfaRegisteredUserId) ID() string {
	fmtString := "/reports/userInsights/monthly/mfaRegisteredUsers/%s"
	return fmt.Sprintf(fmtString, id.MfaUserCountMetricId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report User Insight Monthly Mfa Registered User ID
func (id ReportUserInsightMonthlyMfaRegisteredUserId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("userInsights", "userInsights", "userInsights"),
		resourceids.StaticSegment("monthly", "monthly", "monthly"),
		resourceids.StaticSegment("mfaRegisteredUsers", "mfaRegisteredUsers", "mfaRegisteredUsers"),
		resourceids.UserSpecifiedSegment("mfaUserCountMetricId", "mfaUserCountMetricId"),
	}
}

// String returns a human-readable description of this Report User Insight Monthly Mfa Registered User ID
func (id ReportUserInsightMonthlyMfaRegisteredUserId) String() string {
	components := []string{
		fmt.Sprintf("Mfa User Count Metric: %q", id.MfaUserCountMetricId),
	}
	return fmt.Sprintf("Report User Insight Monthly Mfa Registered User (%s)", strings.Join(components, "\n"))
}
