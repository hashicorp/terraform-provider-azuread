package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportDailyPrintUsageSummariesByUserId{}

// ReportDailyPrintUsageSummariesByUserId is a struct representing the Resource ID for a Report Daily Print Usage Summaries By User
type ReportDailyPrintUsageSummariesByUserId struct {
	PrintUsageByUserId string
}

// NewReportDailyPrintUsageSummariesByUserID returns a new ReportDailyPrintUsageSummariesByUserId struct
func NewReportDailyPrintUsageSummariesByUserID(printUsageByUserId string) ReportDailyPrintUsageSummariesByUserId {
	return ReportDailyPrintUsageSummariesByUserId{
		PrintUsageByUserId: printUsageByUserId,
	}
}

// ParseReportDailyPrintUsageSummariesByUserID parses 'input' into a ReportDailyPrintUsageSummariesByUserId
func ParseReportDailyPrintUsageSummariesByUserID(input string) (*ReportDailyPrintUsageSummariesByUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportDailyPrintUsageSummariesByUserId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportDailyPrintUsageSummariesByUserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportDailyPrintUsageSummariesByUserIDInsensitively parses 'input' case-insensitively into a ReportDailyPrintUsageSummariesByUserId
// note: this method should only be used for API response data and not user input
func ParseReportDailyPrintUsageSummariesByUserIDInsensitively(input string) (*ReportDailyPrintUsageSummariesByUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportDailyPrintUsageSummariesByUserId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportDailyPrintUsageSummariesByUserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportDailyPrintUsageSummariesByUserId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PrintUsageByUserId, ok = input.Parsed["printUsageByUserId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "printUsageByUserId", input)
	}

	return nil
}

// ValidateReportDailyPrintUsageSummariesByUserID checks that 'input' can be parsed as a Report Daily Print Usage Summaries By User ID
func ValidateReportDailyPrintUsageSummariesByUserID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportDailyPrintUsageSummariesByUserID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report Daily Print Usage Summaries By User ID
func (id ReportDailyPrintUsageSummariesByUserId) ID() string {
	fmtString := "/reports/dailyPrintUsageSummariesByUser/%s"
	return fmt.Sprintf(fmtString, id.PrintUsageByUserId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report Daily Print Usage Summaries By User ID
func (id ReportDailyPrintUsageSummariesByUserId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("dailyPrintUsageSummariesByUser", "dailyPrintUsageSummariesByUser", "dailyPrintUsageSummariesByUser"),
		resourceids.UserSpecifiedSegment("printUsageByUserId", "printUsageByUserId"),
	}
}

// String returns a human-readable description of this Report Daily Print Usage Summaries By User ID
func (id ReportDailyPrintUsageSummariesByUserId) String() string {
	components := []string{
		fmt.Sprintf("Print Usage By User: %q", id.PrintUsageByUserId),
	}
	return fmt.Sprintf("Report Daily Print Usage Summaries By User (%s)", strings.Join(components, "\n"))
}
