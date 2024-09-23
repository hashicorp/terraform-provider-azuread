package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportDailyPrintUsageByUserId{}

// ReportDailyPrintUsageByUserId is a struct representing the Resource ID for a Report Daily Print Usage By User
type ReportDailyPrintUsageByUserId struct {
	PrintUsageByUserId string
}

// NewReportDailyPrintUsageByUserID returns a new ReportDailyPrintUsageByUserId struct
func NewReportDailyPrintUsageByUserID(printUsageByUserId string) ReportDailyPrintUsageByUserId {
	return ReportDailyPrintUsageByUserId{
		PrintUsageByUserId: printUsageByUserId,
	}
}

// ParseReportDailyPrintUsageByUserID parses 'input' into a ReportDailyPrintUsageByUserId
func ParseReportDailyPrintUsageByUserID(input string) (*ReportDailyPrintUsageByUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportDailyPrintUsageByUserId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportDailyPrintUsageByUserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportDailyPrintUsageByUserIDInsensitively parses 'input' case-insensitively into a ReportDailyPrintUsageByUserId
// note: this method should only be used for API response data and not user input
func ParseReportDailyPrintUsageByUserIDInsensitively(input string) (*ReportDailyPrintUsageByUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportDailyPrintUsageByUserId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportDailyPrintUsageByUserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportDailyPrintUsageByUserId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PrintUsageByUserId, ok = input.Parsed["printUsageByUserId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "printUsageByUserId", input)
	}

	return nil
}

// ValidateReportDailyPrintUsageByUserID checks that 'input' can be parsed as a Report Daily Print Usage By User ID
func ValidateReportDailyPrintUsageByUserID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportDailyPrintUsageByUserID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report Daily Print Usage By User ID
func (id ReportDailyPrintUsageByUserId) ID() string {
	fmtString := "/reports/dailyPrintUsageByUser/%s"
	return fmt.Sprintf(fmtString, id.PrintUsageByUserId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report Daily Print Usage By User ID
func (id ReportDailyPrintUsageByUserId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("dailyPrintUsageByUser", "dailyPrintUsageByUser", "dailyPrintUsageByUser"),
		resourceids.UserSpecifiedSegment("printUsageByUserId", "printUsageByUserId"),
	}
}

// String returns a human-readable description of this Report Daily Print Usage By User ID
func (id ReportDailyPrintUsageByUserId) String() string {
	components := []string{
		fmt.Sprintf("Print Usage By User: %q", id.PrintUsageByUserId),
	}
	return fmt.Sprintf("Report Daily Print Usage By User (%s)", strings.Join(components, "\n"))
}
