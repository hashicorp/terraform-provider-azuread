package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportDailyPrintUsageByPrinterId{}

// ReportDailyPrintUsageByPrinterId is a struct representing the Resource ID for a Report Daily Print Usage By Printer
type ReportDailyPrintUsageByPrinterId struct {
	PrintUsageByPrinterId string
}

// NewReportDailyPrintUsageByPrinterID returns a new ReportDailyPrintUsageByPrinterId struct
func NewReportDailyPrintUsageByPrinterID(printUsageByPrinterId string) ReportDailyPrintUsageByPrinterId {
	return ReportDailyPrintUsageByPrinterId{
		PrintUsageByPrinterId: printUsageByPrinterId,
	}
}

// ParseReportDailyPrintUsageByPrinterID parses 'input' into a ReportDailyPrintUsageByPrinterId
func ParseReportDailyPrintUsageByPrinterID(input string) (*ReportDailyPrintUsageByPrinterId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportDailyPrintUsageByPrinterId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportDailyPrintUsageByPrinterId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportDailyPrintUsageByPrinterIDInsensitively parses 'input' case-insensitively into a ReportDailyPrintUsageByPrinterId
// note: this method should only be used for API response data and not user input
func ParseReportDailyPrintUsageByPrinterIDInsensitively(input string) (*ReportDailyPrintUsageByPrinterId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportDailyPrintUsageByPrinterId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportDailyPrintUsageByPrinterId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportDailyPrintUsageByPrinterId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PrintUsageByPrinterId, ok = input.Parsed["printUsageByPrinterId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "printUsageByPrinterId", input)
	}

	return nil
}

// ValidateReportDailyPrintUsageByPrinterID checks that 'input' can be parsed as a Report Daily Print Usage By Printer ID
func ValidateReportDailyPrintUsageByPrinterID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportDailyPrintUsageByPrinterID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report Daily Print Usage By Printer ID
func (id ReportDailyPrintUsageByPrinterId) ID() string {
	fmtString := "/reports/dailyPrintUsageByPrinter/%s"
	return fmt.Sprintf(fmtString, id.PrintUsageByPrinterId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report Daily Print Usage By Printer ID
func (id ReportDailyPrintUsageByPrinterId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("dailyPrintUsageByPrinter", "dailyPrintUsageByPrinter", "dailyPrintUsageByPrinter"),
		resourceids.UserSpecifiedSegment("printUsageByPrinterId", "printUsageByPrinterId"),
	}
}

// String returns a human-readable description of this Report Daily Print Usage By Printer ID
func (id ReportDailyPrintUsageByPrinterId) String() string {
	components := []string{
		fmt.Sprintf("Print Usage By Printer: %q", id.PrintUsageByPrinterId),
	}
	return fmt.Sprintf("Report Daily Print Usage By Printer (%s)", strings.Join(components, "\n"))
}
