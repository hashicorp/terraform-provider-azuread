package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportServicePrincipalSignInActivityId{}

// ReportServicePrincipalSignInActivityId is a struct representing the Resource ID for a Report Service Principal Sign In Activity
type ReportServicePrincipalSignInActivityId struct {
	ServicePrincipalSignInActivityId string
}

// NewReportServicePrincipalSignInActivityID returns a new ReportServicePrincipalSignInActivityId struct
func NewReportServicePrincipalSignInActivityID(servicePrincipalSignInActivityId string) ReportServicePrincipalSignInActivityId {
	return ReportServicePrincipalSignInActivityId{
		ServicePrincipalSignInActivityId: servicePrincipalSignInActivityId,
	}
}

// ParseReportServicePrincipalSignInActivityID parses 'input' into a ReportServicePrincipalSignInActivityId
func ParseReportServicePrincipalSignInActivityID(input string) (*ReportServicePrincipalSignInActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportServicePrincipalSignInActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportServicePrincipalSignInActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportServicePrincipalSignInActivityIDInsensitively parses 'input' case-insensitively into a ReportServicePrincipalSignInActivityId
// note: this method should only be used for API response data and not user input
func ParseReportServicePrincipalSignInActivityIDInsensitively(input string) (*ReportServicePrincipalSignInActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportServicePrincipalSignInActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportServicePrincipalSignInActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportServicePrincipalSignInActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalSignInActivityId, ok = input.Parsed["servicePrincipalSignInActivityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalSignInActivityId", input)
	}

	return nil
}

// ValidateReportServicePrincipalSignInActivityID checks that 'input' can be parsed as a Report Service Principal Sign In Activity ID
func ValidateReportServicePrincipalSignInActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportServicePrincipalSignInActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report Service Principal Sign In Activity ID
func (id ReportServicePrincipalSignInActivityId) ID() string {
	fmtString := "/reports/servicePrincipalSignInActivities/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalSignInActivityId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report Service Principal Sign In Activity ID
func (id ReportServicePrincipalSignInActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("servicePrincipalSignInActivities", "servicePrincipalSignInActivities", "servicePrincipalSignInActivities"),
		resourceids.UserSpecifiedSegment("servicePrincipalSignInActivityId", "servicePrincipalSignInActivityId"),
	}
}

// String returns a human-readable description of this Report Service Principal Sign In Activity ID
func (id ReportServicePrincipalSignInActivityId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal Sign In Activity: %q", id.ServicePrincipalSignInActivityId),
	}
	return fmt.Sprintf("Report Service Principal Sign In Activity (%s)", strings.Join(components, "\n"))
}
