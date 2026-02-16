package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightDailyAuthenticationId{}

// ReportUserInsightDailyAuthenticationId is a struct representing the Resource ID for a Report User Insight Daily Authentication
type ReportUserInsightDailyAuthenticationId struct {
	AuthenticationsMetricId string
}

// NewReportUserInsightDailyAuthenticationID returns a new ReportUserInsightDailyAuthenticationId struct
func NewReportUserInsightDailyAuthenticationID(authenticationsMetricId string) ReportUserInsightDailyAuthenticationId {
	return ReportUserInsightDailyAuthenticationId{
		AuthenticationsMetricId: authenticationsMetricId,
	}
}

// ParseReportUserInsightDailyAuthenticationID parses 'input' into a ReportUserInsightDailyAuthenticationId
func ParseReportUserInsightDailyAuthenticationID(input string) (*ReportUserInsightDailyAuthenticationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightDailyAuthenticationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightDailyAuthenticationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportUserInsightDailyAuthenticationIDInsensitively parses 'input' case-insensitively into a ReportUserInsightDailyAuthenticationId
// note: this method should only be used for API response data and not user input
func ParseReportUserInsightDailyAuthenticationIDInsensitively(input string) (*ReportUserInsightDailyAuthenticationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightDailyAuthenticationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightDailyAuthenticationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportUserInsightDailyAuthenticationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AuthenticationsMetricId, ok = input.Parsed["authenticationsMetricId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationsMetricId", input)
	}

	return nil
}

// ValidateReportUserInsightDailyAuthenticationID checks that 'input' can be parsed as a Report User Insight Daily Authentication ID
func ValidateReportUserInsightDailyAuthenticationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportUserInsightDailyAuthenticationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report User Insight Daily Authentication ID
func (id ReportUserInsightDailyAuthenticationId) ID() string {
	fmtString := "/reports/userInsights/daily/authentications/%s"
	return fmt.Sprintf(fmtString, id.AuthenticationsMetricId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report User Insight Daily Authentication ID
func (id ReportUserInsightDailyAuthenticationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("userInsights", "userInsights", "userInsights"),
		resourceids.StaticSegment("daily", "daily", "daily"),
		resourceids.StaticSegment("authentications", "authentications", "authentications"),
		resourceids.UserSpecifiedSegment("authenticationsMetricId", "authenticationsMetricId"),
	}
}

// String returns a human-readable description of this Report User Insight Daily Authentication ID
func (id ReportUserInsightDailyAuthenticationId) String() string {
	components := []string{
		fmt.Sprintf("Authentications Metric: %q", id.AuthenticationsMetricId),
	}
	return fmt.Sprintf("Report User Insight Daily Authentication (%s)", strings.Join(components, "\n"))
}
