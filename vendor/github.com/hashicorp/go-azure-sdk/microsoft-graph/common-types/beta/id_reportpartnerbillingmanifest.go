package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportPartnerBillingManifestId{}

// ReportPartnerBillingManifestId is a struct representing the Resource ID for a Report Partner Billing Manifest
type ReportPartnerBillingManifestId struct {
	ManifestId string
}

// NewReportPartnerBillingManifestID returns a new ReportPartnerBillingManifestId struct
func NewReportPartnerBillingManifestID(manifestId string) ReportPartnerBillingManifestId {
	return ReportPartnerBillingManifestId{
		ManifestId: manifestId,
	}
}

// ParseReportPartnerBillingManifestID parses 'input' into a ReportPartnerBillingManifestId
func ParseReportPartnerBillingManifestID(input string) (*ReportPartnerBillingManifestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportPartnerBillingManifestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportPartnerBillingManifestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportPartnerBillingManifestIDInsensitively parses 'input' case-insensitively into a ReportPartnerBillingManifestId
// note: this method should only be used for API response data and not user input
func ParseReportPartnerBillingManifestIDInsensitively(input string) (*ReportPartnerBillingManifestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportPartnerBillingManifestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportPartnerBillingManifestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportPartnerBillingManifestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManifestId, ok = input.Parsed["manifestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "manifestId", input)
	}

	return nil
}

// ValidateReportPartnerBillingManifestID checks that 'input' can be parsed as a Report Partner Billing Manifest ID
func ValidateReportPartnerBillingManifestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportPartnerBillingManifestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report Partner Billing Manifest ID
func (id ReportPartnerBillingManifestId) ID() string {
	fmtString := "/reports/partners/billing/manifests/%s"
	return fmt.Sprintf(fmtString, id.ManifestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report Partner Billing Manifest ID
func (id ReportPartnerBillingManifestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("partners", "partners", "partners"),
		resourceids.StaticSegment("billing", "billing", "billing"),
		resourceids.StaticSegment("manifests", "manifests", "manifests"),
		resourceids.UserSpecifiedSegment("manifestId", "manifestId"),
	}
}

// String returns a human-readable description of this Report Partner Billing Manifest ID
func (id ReportPartnerBillingManifestId) String() string {
	components := []string{
		fmt.Sprintf("Manifest: %q", id.ManifestId),
	}
	return fmt.Sprintf("Report Partner Billing Manifest (%s)", strings.Join(components, "\n"))
}
