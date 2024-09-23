package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportMonthlyPrintUsageSummariesByUserId{}

// ReportMonthlyPrintUsageSummariesByUserId is a struct representing the Resource ID for a Report Monthly Print Usage Summaries By User
type ReportMonthlyPrintUsageSummariesByUserId struct {
	PrintUsageByUserId string
}

// NewReportMonthlyPrintUsageSummariesByUserID returns a new ReportMonthlyPrintUsageSummariesByUserId struct
func NewReportMonthlyPrintUsageSummariesByUserID(printUsageByUserId string) ReportMonthlyPrintUsageSummariesByUserId {
	return ReportMonthlyPrintUsageSummariesByUserId{
		PrintUsageByUserId: printUsageByUserId,
	}
}

// ParseReportMonthlyPrintUsageSummariesByUserID parses 'input' into a ReportMonthlyPrintUsageSummariesByUserId
func ParseReportMonthlyPrintUsageSummariesByUserID(input string) (*ReportMonthlyPrintUsageSummariesByUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportMonthlyPrintUsageSummariesByUserId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportMonthlyPrintUsageSummariesByUserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportMonthlyPrintUsageSummariesByUserIDInsensitively parses 'input' case-insensitively into a ReportMonthlyPrintUsageSummariesByUserId
// note: this method should only be used for API response data and not user input
func ParseReportMonthlyPrintUsageSummariesByUserIDInsensitively(input string) (*ReportMonthlyPrintUsageSummariesByUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportMonthlyPrintUsageSummariesByUserId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportMonthlyPrintUsageSummariesByUserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportMonthlyPrintUsageSummariesByUserId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PrintUsageByUserId, ok = input.Parsed["printUsageByUserId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "printUsageByUserId", input)
	}

	return nil
}

// ValidateReportMonthlyPrintUsageSummariesByUserID checks that 'input' can be parsed as a Report Monthly Print Usage Summaries By User ID
func ValidateReportMonthlyPrintUsageSummariesByUserID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportMonthlyPrintUsageSummariesByUserID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report Monthly Print Usage Summaries By User ID
func (id ReportMonthlyPrintUsageSummariesByUserId) ID() string {
	fmtString := "/reports/monthlyPrintUsageSummariesByUser/%s"
	return fmt.Sprintf(fmtString, id.PrintUsageByUserId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report Monthly Print Usage Summaries By User ID
func (id ReportMonthlyPrintUsageSummariesByUserId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("monthlyPrintUsageSummariesByUser", "monthlyPrintUsageSummariesByUser", "monthlyPrintUsageSummariesByUser"),
		resourceids.UserSpecifiedSegment("printUsageByUserId", "printUsageByUserId"),
	}
}

// String returns a human-readable description of this Report Monthly Print Usage Summaries By User ID
func (id ReportMonthlyPrintUsageSummariesByUserId) String() string {
	components := []string{
		fmt.Sprintf("Print Usage By User: %q", id.PrintUsageByUserId),
	}
	return fmt.Sprintf("Report Monthly Print Usage Summaries By User (%s)", strings.Join(components, "\n"))
}
