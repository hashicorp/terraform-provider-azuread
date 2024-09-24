package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportApplicationSignInDetailedSummaryId{}

// ReportApplicationSignInDetailedSummaryId is a struct representing the Resource ID for a Report Application Sign In Detailed Summary
type ReportApplicationSignInDetailedSummaryId struct {
	ApplicationSignInDetailedSummaryId string
}

// NewReportApplicationSignInDetailedSummaryID returns a new ReportApplicationSignInDetailedSummaryId struct
func NewReportApplicationSignInDetailedSummaryID(applicationSignInDetailedSummaryId string) ReportApplicationSignInDetailedSummaryId {
	return ReportApplicationSignInDetailedSummaryId{
		ApplicationSignInDetailedSummaryId: applicationSignInDetailedSummaryId,
	}
}

// ParseReportApplicationSignInDetailedSummaryID parses 'input' into a ReportApplicationSignInDetailedSummaryId
func ParseReportApplicationSignInDetailedSummaryID(input string) (*ReportApplicationSignInDetailedSummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportApplicationSignInDetailedSummaryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportApplicationSignInDetailedSummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportApplicationSignInDetailedSummaryIDInsensitively parses 'input' case-insensitively into a ReportApplicationSignInDetailedSummaryId
// note: this method should only be used for API response data and not user input
func ParseReportApplicationSignInDetailedSummaryIDInsensitively(input string) (*ReportApplicationSignInDetailedSummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportApplicationSignInDetailedSummaryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportApplicationSignInDetailedSummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportApplicationSignInDetailedSummaryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApplicationSignInDetailedSummaryId, ok = input.Parsed["applicationSignInDetailedSummaryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationSignInDetailedSummaryId", input)
	}

	return nil
}

// ValidateReportApplicationSignInDetailedSummaryID checks that 'input' can be parsed as a Report Application Sign In Detailed Summary ID
func ValidateReportApplicationSignInDetailedSummaryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportApplicationSignInDetailedSummaryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report Application Sign In Detailed Summary ID
func (id ReportApplicationSignInDetailedSummaryId) ID() string {
	fmtString := "/reports/applicationSignInDetailedSummary/%s"
	return fmt.Sprintf(fmtString, id.ApplicationSignInDetailedSummaryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report Application Sign In Detailed Summary ID
func (id ReportApplicationSignInDetailedSummaryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("applicationSignInDetailedSummary", "applicationSignInDetailedSummary", "applicationSignInDetailedSummary"),
		resourceids.UserSpecifiedSegment("applicationSignInDetailedSummaryId", "applicationSignInDetailedSummaryId"),
	}
}

// String returns a human-readable description of this Report Application Sign In Detailed Summary ID
func (id ReportApplicationSignInDetailedSummaryId) String() string {
	components := []string{
		fmt.Sprintf("Application Sign In Detailed Summary: %q", id.ApplicationSignInDetailedSummaryId),
	}
	return fmt.Sprintf("Report Application Sign In Detailed Summary (%s)", strings.Join(components, "\n"))
}
