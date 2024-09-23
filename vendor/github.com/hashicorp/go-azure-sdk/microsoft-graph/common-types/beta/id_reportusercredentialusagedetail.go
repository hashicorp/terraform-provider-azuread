package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserCredentialUsageDetailId{}

// ReportUserCredentialUsageDetailId is a struct representing the Resource ID for a Report User Credential Usage Detail
type ReportUserCredentialUsageDetailId struct {
	UserCredentialUsageDetailsId string
}

// NewReportUserCredentialUsageDetailID returns a new ReportUserCredentialUsageDetailId struct
func NewReportUserCredentialUsageDetailID(userCredentialUsageDetailsId string) ReportUserCredentialUsageDetailId {
	return ReportUserCredentialUsageDetailId{
		UserCredentialUsageDetailsId: userCredentialUsageDetailsId,
	}
}

// ParseReportUserCredentialUsageDetailID parses 'input' into a ReportUserCredentialUsageDetailId
func ParseReportUserCredentialUsageDetailID(input string) (*ReportUserCredentialUsageDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserCredentialUsageDetailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserCredentialUsageDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportUserCredentialUsageDetailIDInsensitively parses 'input' case-insensitively into a ReportUserCredentialUsageDetailId
// note: this method should only be used for API response data and not user input
func ParseReportUserCredentialUsageDetailIDInsensitively(input string) (*ReportUserCredentialUsageDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportUserCredentialUsageDetailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportUserCredentialUsageDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportUserCredentialUsageDetailId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserCredentialUsageDetailsId, ok = input.Parsed["userCredentialUsageDetailsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userCredentialUsageDetailsId", input)
	}

	return nil
}

// ValidateReportUserCredentialUsageDetailID checks that 'input' can be parsed as a Report User Credential Usage Detail ID
func ValidateReportUserCredentialUsageDetailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportUserCredentialUsageDetailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report User Credential Usage Detail ID
func (id ReportUserCredentialUsageDetailId) ID() string {
	fmtString := "/reports/userCredentialUsageDetails/%s"
	return fmt.Sprintf(fmtString, id.UserCredentialUsageDetailsId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report User Credential Usage Detail ID
func (id ReportUserCredentialUsageDetailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("userCredentialUsageDetails", "userCredentialUsageDetails", "userCredentialUsageDetails"),
		resourceids.UserSpecifiedSegment("userCredentialUsageDetailsId", "userCredentialUsageDetailsId"),
	}
}

// String returns a human-readable description of this Report User Credential Usage Detail ID
func (id ReportUserCredentialUsageDetailId) String() string {
	components := []string{
		fmt.Sprintf("User Credential Usage Details: %q", id.UserCredentialUsageDetailsId),
	}
	return fmt.Sprintf("Report User Credential Usage Detail (%s)", strings.Join(components, "\n"))
}
