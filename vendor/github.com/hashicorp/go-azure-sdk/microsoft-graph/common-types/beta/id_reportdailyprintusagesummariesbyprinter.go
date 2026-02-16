package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportDailyPrintUsageSummariesByPrinterId{}

// ReportDailyPrintUsageSummariesByPrinterId is a struct representing the Resource ID for a Report Daily Print Usage Summaries By Printer
type ReportDailyPrintUsageSummariesByPrinterId struct {
	PrintUsageByPrinterId string
}

// NewReportDailyPrintUsageSummariesByPrinterID returns a new ReportDailyPrintUsageSummariesByPrinterId struct
func NewReportDailyPrintUsageSummariesByPrinterID(printUsageByPrinterId string) ReportDailyPrintUsageSummariesByPrinterId {
	return ReportDailyPrintUsageSummariesByPrinterId{
		PrintUsageByPrinterId: printUsageByPrinterId,
	}
}

// ParseReportDailyPrintUsageSummariesByPrinterID parses 'input' into a ReportDailyPrintUsageSummariesByPrinterId
func ParseReportDailyPrintUsageSummariesByPrinterID(input string) (*ReportDailyPrintUsageSummariesByPrinterId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportDailyPrintUsageSummariesByPrinterId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportDailyPrintUsageSummariesByPrinterId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportDailyPrintUsageSummariesByPrinterIDInsensitively parses 'input' case-insensitively into a ReportDailyPrintUsageSummariesByPrinterId
// note: this method should only be used for API response data and not user input
func ParseReportDailyPrintUsageSummariesByPrinterIDInsensitively(input string) (*ReportDailyPrintUsageSummariesByPrinterId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportDailyPrintUsageSummariesByPrinterId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportDailyPrintUsageSummariesByPrinterId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportDailyPrintUsageSummariesByPrinterId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PrintUsageByPrinterId, ok = input.Parsed["printUsageByPrinterId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "printUsageByPrinterId", input)
	}

	return nil
}

// ValidateReportDailyPrintUsageSummariesByPrinterID checks that 'input' can be parsed as a Report Daily Print Usage Summaries By Printer ID
func ValidateReportDailyPrintUsageSummariesByPrinterID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportDailyPrintUsageSummariesByPrinterID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report Daily Print Usage Summaries By Printer ID
func (id ReportDailyPrintUsageSummariesByPrinterId) ID() string {
	fmtString := "/reports/dailyPrintUsageSummariesByPrinter/%s"
	return fmt.Sprintf(fmtString, id.PrintUsageByPrinterId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report Daily Print Usage Summaries By Printer ID
func (id ReportDailyPrintUsageSummariesByPrinterId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("dailyPrintUsageSummariesByPrinter", "dailyPrintUsageSummariesByPrinter", "dailyPrintUsageSummariesByPrinter"),
		resourceids.UserSpecifiedSegment("printUsageByPrinterId", "printUsageByPrinterId"),
	}
}

// String returns a human-readable description of this Report Daily Print Usage Summaries By Printer ID
func (id ReportDailyPrintUsageSummariesByPrinterId) String() string {
	components := []string{
		fmt.Sprintf("Print Usage By Printer: %q", id.PrintUsageByPrinterId),
	}
	return fmt.Sprintf("Report Daily Print Usage Summaries By Printer (%s)", strings.Join(components, "\n"))
}
