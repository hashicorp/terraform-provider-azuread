package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportMonthlyPrintUsageByUserId{}

// ReportMonthlyPrintUsageByUserId is a struct representing the Resource ID for a Report Monthly Print Usage By User
type ReportMonthlyPrintUsageByUserId struct {
	PrintUsageByUserId string
}

// NewReportMonthlyPrintUsageByUserID returns a new ReportMonthlyPrintUsageByUserId struct
func NewReportMonthlyPrintUsageByUserID(printUsageByUserId string) ReportMonthlyPrintUsageByUserId {
	return ReportMonthlyPrintUsageByUserId{
		PrintUsageByUserId: printUsageByUserId,
	}
}

// ParseReportMonthlyPrintUsageByUserID parses 'input' into a ReportMonthlyPrintUsageByUserId
func ParseReportMonthlyPrintUsageByUserID(input string) (*ReportMonthlyPrintUsageByUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportMonthlyPrintUsageByUserId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportMonthlyPrintUsageByUserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportMonthlyPrintUsageByUserIDInsensitively parses 'input' case-insensitively into a ReportMonthlyPrintUsageByUserId
// note: this method should only be used for API response data and not user input
func ParseReportMonthlyPrintUsageByUserIDInsensitively(input string) (*ReportMonthlyPrintUsageByUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportMonthlyPrintUsageByUserId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportMonthlyPrintUsageByUserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportMonthlyPrintUsageByUserId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PrintUsageByUserId, ok = input.Parsed["printUsageByUserId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "printUsageByUserId", input)
	}

	return nil
}

// ValidateReportMonthlyPrintUsageByUserID checks that 'input' can be parsed as a Report Monthly Print Usage By User ID
func ValidateReportMonthlyPrintUsageByUserID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportMonthlyPrintUsageByUserID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report Monthly Print Usage By User ID
func (id ReportMonthlyPrintUsageByUserId) ID() string {
	fmtString := "/reports/monthlyPrintUsageByUser/%s"
	return fmt.Sprintf(fmtString, id.PrintUsageByUserId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report Monthly Print Usage By User ID
func (id ReportMonthlyPrintUsageByUserId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("monthlyPrintUsageByUser", "monthlyPrintUsageByUser", "monthlyPrintUsageByUser"),
		resourceids.UserSpecifiedSegment("printUsageByUserId", "printUsageByUserId"),
	}
}

// String returns a human-readable description of this Report Monthly Print Usage By User ID
func (id ReportMonthlyPrintUsageByUserId) String() string {
	components := []string{
		fmt.Sprintf("Print Usage By User: %q", id.PrintUsageByUserId),
	}
	return fmt.Sprintf("Report Monthly Print Usage By User (%s)", strings.Join(components, "\n"))
}
