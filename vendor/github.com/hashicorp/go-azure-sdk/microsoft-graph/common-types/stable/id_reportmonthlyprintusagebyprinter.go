package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportMonthlyPrintUsageByPrinterId{}

// ReportMonthlyPrintUsageByPrinterId is a struct representing the Resource ID for a Report Monthly Print Usage By Printer
type ReportMonthlyPrintUsageByPrinterId struct {
	PrintUsageByPrinterId string
}

// NewReportMonthlyPrintUsageByPrinterID returns a new ReportMonthlyPrintUsageByPrinterId struct
func NewReportMonthlyPrintUsageByPrinterID(printUsageByPrinterId string) ReportMonthlyPrintUsageByPrinterId {
	return ReportMonthlyPrintUsageByPrinterId{
		PrintUsageByPrinterId: printUsageByPrinterId,
	}
}

// ParseReportMonthlyPrintUsageByPrinterID parses 'input' into a ReportMonthlyPrintUsageByPrinterId
func ParseReportMonthlyPrintUsageByPrinterID(input string) (*ReportMonthlyPrintUsageByPrinterId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportMonthlyPrintUsageByPrinterId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportMonthlyPrintUsageByPrinterId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportMonthlyPrintUsageByPrinterIDInsensitively parses 'input' case-insensitively into a ReportMonthlyPrintUsageByPrinterId
// note: this method should only be used for API response data and not user input
func ParseReportMonthlyPrintUsageByPrinterIDInsensitively(input string) (*ReportMonthlyPrintUsageByPrinterId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportMonthlyPrintUsageByPrinterId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportMonthlyPrintUsageByPrinterId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportMonthlyPrintUsageByPrinterId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PrintUsageByPrinterId, ok = input.Parsed["printUsageByPrinterId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "printUsageByPrinterId", input)
	}

	return nil
}

// ValidateReportMonthlyPrintUsageByPrinterID checks that 'input' can be parsed as a Report Monthly Print Usage By Printer ID
func ValidateReportMonthlyPrintUsageByPrinterID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportMonthlyPrintUsageByPrinterID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report Monthly Print Usage By Printer ID
func (id ReportMonthlyPrintUsageByPrinterId) ID() string {
	fmtString := "/reports/monthlyPrintUsageByPrinter/%s"
	return fmt.Sprintf(fmtString, id.PrintUsageByPrinterId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report Monthly Print Usage By Printer ID
func (id ReportMonthlyPrintUsageByPrinterId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("monthlyPrintUsageByPrinter", "monthlyPrintUsageByPrinter", "monthlyPrintUsageByPrinter"),
		resourceids.UserSpecifiedSegment("printUsageByPrinterId", "printUsageByPrinterId"),
	}
}

// String returns a human-readable description of this Report Monthly Print Usage By Printer ID
func (id ReportMonthlyPrintUsageByPrinterId) String() string {
	components := []string{
		fmt.Sprintf("Print Usage By Printer: %q", id.PrintUsageByPrinterId),
	}
	return fmt.Sprintf("Report Monthly Print Usage By Printer (%s)", strings.Join(components, "\n"))
}
