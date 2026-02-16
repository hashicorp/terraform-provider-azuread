package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportPartnerBillingOperationId{}

// ReportPartnerBillingOperationId is a struct representing the Resource ID for a Report Partner Billing Operation
type ReportPartnerBillingOperationId struct {
	OperationId string
}

// NewReportPartnerBillingOperationID returns a new ReportPartnerBillingOperationId struct
func NewReportPartnerBillingOperationID(operationId string) ReportPartnerBillingOperationId {
	return ReportPartnerBillingOperationId{
		OperationId: operationId,
	}
}

// ParseReportPartnerBillingOperationID parses 'input' into a ReportPartnerBillingOperationId
func ParseReportPartnerBillingOperationID(input string) (*ReportPartnerBillingOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportPartnerBillingOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportPartnerBillingOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportPartnerBillingOperationIDInsensitively parses 'input' case-insensitively into a ReportPartnerBillingOperationId
// note: this method should only be used for API response data and not user input
func ParseReportPartnerBillingOperationIDInsensitively(input string) (*ReportPartnerBillingOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportPartnerBillingOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportPartnerBillingOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportPartnerBillingOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OperationId, ok = input.Parsed["operationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "operationId", input)
	}

	return nil
}

// ValidateReportPartnerBillingOperationID checks that 'input' can be parsed as a Report Partner Billing Operation ID
func ValidateReportPartnerBillingOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportPartnerBillingOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report Partner Billing Operation ID
func (id ReportPartnerBillingOperationId) ID() string {
	fmtString := "/reports/partners/billing/operations/%s"
	return fmt.Sprintf(fmtString, id.OperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report Partner Billing Operation ID
func (id ReportPartnerBillingOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("partners", "partners", "partners"),
		resourceids.StaticSegment("billing", "billing", "billing"),
		resourceids.StaticSegment("operations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("operationId", "operationId"),
	}
}

// String returns a human-readable description of this Report Partner Billing Operation ID
func (id ReportPartnerBillingOperationId) String() string {
	components := []string{
		fmt.Sprintf("Operation: %q", id.OperationId),
	}
	return fmt.Sprintf("Report Partner Billing Operation (%s)", strings.Join(components, "\n"))
}
