package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportMonthlyPrintUsageSummariesByPrinterId{}

// ReportMonthlyPrintUsageSummariesByPrinterId is a struct representing the Resource ID for a Report Monthly Print Usage Summaries By Printer
type ReportMonthlyPrintUsageSummariesByPrinterId struct {
	PrintUsageByPrinterId string
}

// NewReportMonthlyPrintUsageSummariesByPrinterID returns a new ReportMonthlyPrintUsageSummariesByPrinterId struct
func NewReportMonthlyPrintUsageSummariesByPrinterID(printUsageByPrinterId string) ReportMonthlyPrintUsageSummariesByPrinterId {
	return ReportMonthlyPrintUsageSummariesByPrinterId{
		PrintUsageByPrinterId: printUsageByPrinterId,
	}
}

// ParseReportMonthlyPrintUsageSummariesByPrinterID parses 'input' into a ReportMonthlyPrintUsageSummariesByPrinterId
func ParseReportMonthlyPrintUsageSummariesByPrinterID(input string) (*ReportMonthlyPrintUsageSummariesByPrinterId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportMonthlyPrintUsageSummariesByPrinterId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportMonthlyPrintUsageSummariesByPrinterId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportMonthlyPrintUsageSummariesByPrinterIDInsensitively parses 'input' case-insensitively into a ReportMonthlyPrintUsageSummariesByPrinterId
// note: this method should only be used for API response data and not user input
func ParseReportMonthlyPrintUsageSummariesByPrinterIDInsensitively(input string) (*ReportMonthlyPrintUsageSummariesByPrinterId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportMonthlyPrintUsageSummariesByPrinterId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportMonthlyPrintUsageSummariesByPrinterId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportMonthlyPrintUsageSummariesByPrinterId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PrintUsageByPrinterId, ok = input.Parsed["printUsageByPrinterId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "printUsageByPrinterId", input)
	}

	return nil
}

// ValidateReportMonthlyPrintUsageSummariesByPrinterID checks that 'input' can be parsed as a Report Monthly Print Usage Summaries By Printer ID
func ValidateReportMonthlyPrintUsageSummariesByPrinterID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportMonthlyPrintUsageSummariesByPrinterID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report Monthly Print Usage Summaries By Printer ID
func (id ReportMonthlyPrintUsageSummariesByPrinterId) ID() string {
	fmtString := "/reports/monthlyPrintUsageSummariesByPrinter/%s"
	return fmt.Sprintf(fmtString, id.PrintUsageByPrinterId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report Monthly Print Usage Summaries By Printer ID
func (id ReportMonthlyPrintUsageSummariesByPrinterId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("monthlyPrintUsageSummariesByPrinter", "monthlyPrintUsageSummariesByPrinter", "monthlyPrintUsageSummariesByPrinter"),
		resourceids.UserSpecifiedSegment("printUsageByPrinterId", "printUsageByPrinterId"),
	}
}

// String returns a human-readable description of this Report Monthly Print Usage Summaries By Printer ID
func (id ReportMonthlyPrintUsageSummariesByPrinterId) String() string {
	components := []string{
		fmt.Sprintf("Print Usage By Printer: %q", id.PrintUsageByPrinterId),
	}
	return fmt.Sprintf("Report Monthly Print Usage Summaries By Printer (%s)", strings.Join(components, "\n"))
}
