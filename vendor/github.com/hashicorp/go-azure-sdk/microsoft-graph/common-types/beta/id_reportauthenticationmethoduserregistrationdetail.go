package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportAuthenticationMethodUserRegistrationDetailId{}

// ReportAuthenticationMethodUserRegistrationDetailId is a struct representing the Resource ID for a Report Authentication Method User Registration Detail
type ReportAuthenticationMethodUserRegistrationDetailId struct {
	UserRegistrationDetailsId string
}

// NewReportAuthenticationMethodUserRegistrationDetailID returns a new ReportAuthenticationMethodUserRegistrationDetailId struct
func NewReportAuthenticationMethodUserRegistrationDetailID(userRegistrationDetailsId string) ReportAuthenticationMethodUserRegistrationDetailId {
	return ReportAuthenticationMethodUserRegistrationDetailId{
		UserRegistrationDetailsId: userRegistrationDetailsId,
	}
}

// ParseReportAuthenticationMethodUserRegistrationDetailID parses 'input' into a ReportAuthenticationMethodUserRegistrationDetailId
func ParseReportAuthenticationMethodUserRegistrationDetailID(input string) (*ReportAuthenticationMethodUserRegistrationDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportAuthenticationMethodUserRegistrationDetailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportAuthenticationMethodUserRegistrationDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportAuthenticationMethodUserRegistrationDetailIDInsensitively parses 'input' case-insensitively into a ReportAuthenticationMethodUserRegistrationDetailId
// note: this method should only be used for API response data and not user input
func ParseReportAuthenticationMethodUserRegistrationDetailIDInsensitively(input string) (*ReportAuthenticationMethodUserRegistrationDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportAuthenticationMethodUserRegistrationDetailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportAuthenticationMethodUserRegistrationDetailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportAuthenticationMethodUserRegistrationDetailId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserRegistrationDetailsId, ok = input.Parsed["userRegistrationDetailsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userRegistrationDetailsId", input)
	}

	return nil
}

// ValidateReportAuthenticationMethodUserRegistrationDetailID checks that 'input' can be parsed as a Report Authentication Method User Registration Detail ID
func ValidateReportAuthenticationMethodUserRegistrationDetailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportAuthenticationMethodUserRegistrationDetailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report Authentication Method User Registration Detail ID
func (id ReportAuthenticationMethodUserRegistrationDetailId) ID() string {
	fmtString := "/reports/authenticationMethods/userRegistrationDetails/%s"
	return fmt.Sprintf(fmtString, id.UserRegistrationDetailsId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report Authentication Method User Registration Detail ID
func (id ReportAuthenticationMethodUserRegistrationDetailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("authenticationMethods", "authenticationMethods", "authenticationMethods"),
		resourceids.StaticSegment("userRegistrationDetails", "userRegistrationDetails", "userRegistrationDetails"),
		resourceids.UserSpecifiedSegment("userRegistrationDetailsId", "userRegistrationDetailsId"),
	}
}

// String returns a human-readable description of this Report Authentication Method User Registration Detail ID
func (id ReportAuthenticationMethodUserRegistrationDetailId) String() string {
	components := []string{
		fmt.Sprintf("User Registration Details: %q", id.UserRegistrationDetailsId),
	}
	return fmt.Sprintf("Report Authentication Method User Registration Detail (%s)", strings.Join(components, "\n"))
}
