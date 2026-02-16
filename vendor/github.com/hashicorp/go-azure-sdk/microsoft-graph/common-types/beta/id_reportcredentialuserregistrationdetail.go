package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportCredentialUserRegistrationDetailId{}

// ReportCredentialUserRegistrationDetailId is a struct representing the Resource ID for a Report Credential User Registration Detail
type ReportCredentialUserRegistrationDetailId struct {
	CredentialUserRegistrationDetailsId string
}

// NewReportCredentialUserRegistrationDetailID returns a new ReportCredentialUserRegistrationDetailId struct
func NewReportCredentialUserRegistrationDetailID(credentialUserRegistrationDetailsId string) ReportCredentialUserRegistrationDetailId {
	return ReportCredentialUserRegistrationDetailId{
		CredentialUserRegistrationDetailsId: credentialUserRegistrationDetailsId,
	}
}

// ParseReportCredentialUserRegistrationDetailID parses 'input' into a ReportCredentialUserRegistrationDetailId
func ParseReportCredentialUserRegistrationDetailID(input string) (*ReportCredentialUserRegistrationDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportCredentialUserRegistrationDetailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportCredentialUserRegistrationDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportCredentialUserRegistrationDetailIDInsensitively parses 'input' case-insensitively into a ReportCredentialUserRegistrationDetailId
// note: this method should only be used for API response data and not user input
func ParseReportCredentialUserRegistrationDetailIDInsensitively(input string) (*ReportCredentialUserRegistrationDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportCredentialUserRegistrationDetailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportCredentialUserRegistrationDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportCredentialUserRegistrationDetailId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CredentialUserRegistrationDetailsId, ok = input.Parsed["credentialUserRegistrationDetailsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "credentialUserRegistrationDetailsId", input)
	}

	return nil
}

// ValidateReportCredentialUserRegistrationDetailID checks that 'input' can be parsed as a Report Credential User Registration Detail ID
func ValidateReportCredentialUserRegistrationDetailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportCredentialUserRegistrationDetailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report Credential User Registration Detail ID
func (id ReportCredentialUserRegistrationDetailId) ID() string {
	fmtString := "/reports/credentialUserRegistrationDetails/%s"
	return fmt.Sprintf(fmtString, id.CredentialUserRegistrationDetailsId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report Credential User Registration Detail ID
func (id ReportCredentialUserRegistrationDetailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("credentialUserRegistrationDetails", "credentialUserRegistrationDetails", "credentialUserRegistrationDetails"),
		resourceids.UserSpecifiedSegment("credentialUserRegistrationDetailsId", "credentialUserRegistrationDetailsId"),
	}
}

// String returns a human-readable description of this Report Credential User Registration Detail ID
func (id ReportCredentialUserRegistrationDetailId) String() string {
	components := []string{
		fmt.Sprintf("Credential User Registration Details: %q", id.CredentialUserRegistrationDetailsId),
	}
	return fmt.Sprintf("Report Credential User Registration Detail (%s)", strings.Join(components, "\n"))
}
