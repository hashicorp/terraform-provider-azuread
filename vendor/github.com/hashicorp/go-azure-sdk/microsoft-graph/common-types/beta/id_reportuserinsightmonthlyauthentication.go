package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightMonthlyAuthenticationId{}

// ReportUserInsightMonthlyAuthenticationId is a struct representing the Resource ID for a Report User Insight Monthly Authentication
type ReportUserInsightMonthlyAuthenticationId struct {
	AuthenticationsMetricId string
}

// NewReportUserInsightMonthlyAuthenticationID returns a new ReportUserInsightMonthlyAuthenticationId struct
func NewReportUserInsightMonthlyAuthenticationID(authenticationsMetricId string) ReportUserInsightMonthlyAuthenticationId {
	return ReportUserInsightMonthlyAuthenticationId{
		AuthenticationsMetricId: authenticationsMetricId,
	}
}

// ParseReportUserInsightMonthlyAuthenticationID parses 'input' into a ReportUserInsightMonthlyAuthenticationId
func ParseReportUserInsightMonthlyAuthenticationID(input string) (*ReportUserInsightMonthlyAuthenticationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightMonthlyAuthenticationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightMonthlyAuthenticationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportUserInsightMonthlyAuthenticationIDInsensitively parses 'input' case-insensitively into a ReportUserInsightMonthlyAuthenticationId
// note: this method should only be used for API response data and not user input
func ParseReportUserInsightMonthlyAuthenticationIDInsensitively(input string) (*ReportUserInsightMonthlyAuthenticationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserInsightMonthlyAuthenticationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserInsightMonthlyAuthenticationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportUserInsightMonthlyAuthenticationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AuthenticationsMetricId, ok = input.Parsed["authenticationsMetricId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationsMetricId", input)
	}

	return nil
}

// ValidateReportUserInsightMonthlyAuthenticationID checks that 'input' can be parsed as a Report User Insight Monthly Authentication ID
func ValidateReportUserInsightMonthlyAuthenticationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportUserInsightMonthlyAuthenticationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report User Insight Monthly Authentication ID
func (id ReportUserInsightMonthlyAuthenticationId) ID() string {
	fmtString := "/reports/userInsights/monthly/authentications/%s"
	return fmt.Sprintf(fmtString, id.AuthenticationsMetricId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report User Insight Monthly Authentication ID
func (id ReportUserInsightMonthlyAuthenticationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("userInsights", "userInsights", "userInsights"),
		resourceids.StaticSegment("monthly", "monthly", "monthly"),
		resourceids.StaticSegment("authentications", "authentications", "authentications"),
		resourceids.UserSpecifiedSegment("authenticationsMetricId", "authenticationsMetricId"),
	}
}

// String returns a human-readable description of this Report User Insight Monthly Authentication ID
func (id ReportUserInsightMonthlyAuthenticationId) String() string {
	components := []string{
		fmt.Sprintf("Authentications Metric: %q", id.AuthenticationsMetricId),
	}
	return fmt.Sprintf("Report User Insight Monthly Authentication (%s)", strings.Join(components, "\n"))
}
