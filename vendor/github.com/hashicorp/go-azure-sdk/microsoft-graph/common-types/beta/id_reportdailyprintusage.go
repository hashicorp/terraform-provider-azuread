package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportDailyPrintUsageId{}

// ReportDailyPrintUsageId is a struct representing the Resource ID for a Report Daily Print Usage
type ReportDailyPrintUsageId struct {
	PrintUsageId string
}

// NewReportDailyPrintUsageID returns a new ReportDailyPrintUsageId struct
func NewReportDailyPrintUsageID(printUsageId string) ReportDailyPrintUsageId {
	return ReportDailyPrintUsageId{
		PrintUsageId: printUsageId,
	}
}

// ParseReportDailyPrintUsageID parses 'input' into a ReportDailyPrintUsageId
func ParseReportDailyPrintUsageID(input string) (*ReportDailyPrintUsageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportDailyPrintUsageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportDailyPrintUsageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportDailyPrintUsageIDInsensitively parses 'input' case-insensitively into a ReportDailyPrintUsageId
// note: this method should only be used for API response data and not user input
func ParseReportDailyPrintUsageIDInsensitively(input string) (*ReportDailyPrintUsageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportDailyPrintUsageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportDailyPrintUsageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportDailyPrintUsageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PrintUsageId, ok = input.Parsed["printUsageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "printUsageId", input)
	}

	return nil
}

// ValidateReportDailyPrintUsageID checks that 'input' can be parsed as a Report Daily Print Usage ID
func ValidateReportDailyPrintUsageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportDailyPrintUsageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report Daily Print Usage ID
func (id ReportDailyPrintUsageId) ID() string {
	fmtString := "/reports/dailyPrintUsage/%s"
	return fmt.Sprintf(fmtString, id.PrintUsageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report Daily Print Usage ID
func (id ReportDailyPrintUsageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("dailyPrintUsage", "dailyPrintUsage", "dailyPrintUsage"),
		resourceids.UserSpecifiedSegment("printUsageId", "printUsageId"),
	}
}

// String returns a human-readable description of this Report Daily Print Usage ID
func (id ReportDailyPrintUsageId) String() string {
	components := []string{
		fmt.Sprintf("Print Usage: %q", id.PrintUsageId),
	}
	return fmt.Sprintf("Report Daily Print Usage (%s)", strings.Join(components, "\n"))
}
