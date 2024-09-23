package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportAppCredentialSignInActivityId{}

// ReportAppCredentialSignInActivityId is a struct representing the Resource ID for a Report App Credential Sign In Activity
type ReportAppCredentialSignInActivityId struct {
	AppCredentialSignInActivityId string
}

// NewReportAppCredentialSignInActivityID returns a new ReportAppCredentialSignInActivityId struct
func NewReportAppCredentialSignInActivityID(appCredentialSignInActivityId string) ReportAppCredentialSignInActivityId {
	return ReportAppCredentialSignInActivityId{
		AppCredentialSignInActivityId: appCredentialSignInActivityId,
	}
}

// ParseReportAppCredentialSignInActivityID parses 'input' into a ReportAppCredentialSignInActivityId
func ParseReportAppCredentialSignInActivityID(input string) (*ReportAppCredentialSignInActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportAppCredentialSignInActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportAppCredentialSignInActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportAppCredentialSignInActivityIDInsensitively parses 'input' case-insensitively into a ReportAppCredentialSignInActivityId
// note: this method should only be used for API response data and not user input
func ParseReportAppCredentialSignInActivityIDInsensitively(input string) (*ReportAppCredentialSignInActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportAppCredentialSignInActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportAppCredentialSignInActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportAppCredentialSignInActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AppCredentialSignInActivityId, ok = input.Parsed["appCredentialSignInActivityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appCredentialSignInActivityId", input)
	}

	return nil
}

// ValidateReportAppCredentialSignInActivityID checks that 'input' can be parsed as a Report App Credential Sign In Activity ID
func ValidateReportAppCredentialSignInActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportAppCredentialSignInActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report App Credential Sign In Activity ID
func (id ReportAppCredentialSignInActivityId) ID() string {
	fmtString := "/reports/appCredentialSignInActivities/%s"
	return fmt.Sprintf(fmtString, id.AppCredentialSignInActivityId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report App Credential Sign In Activity ID
func (id ReportAppCredentialSignInActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("appCredentialSignInActivities", "appCredentialSignInActivities", "appCredentialSignInActivities"),
		resourceids.UserSpecifiedSegment("appCredentialSignInActivityId", "appCredentialSignInActivityId"),
	}
}

// String returns a human-readable description of this Report App Credential Sign In Activity ID
func (id ReportAppCredentialSignInActivityId) String() string {
	components := []string{
		fmt.Sprintf("App Credential Sign In Activity: %q", id.AppCredentialSignInActivityId),
	}
	return fmt.Sprintf("Report App Credential Sign In Activity (%s)", strings.Join(components, "\n"))
}
