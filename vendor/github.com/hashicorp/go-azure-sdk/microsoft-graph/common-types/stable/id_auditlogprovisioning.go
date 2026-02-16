package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AuditLogProvisioningId{}

// AuditLogProvisioningId is a struct representing the Resource ID for a Audit Log Provisioning
type AuditLogProvisioningId struct {
	ProvisioningObjectSummaryId string
}

// NewAuditLogProvisioningID returns a new AuditLogProvisioningId struct
func NewAuditLogProvisioningID(provisioningObjectSummaryId string) AuditLogProvisioningId {
	return AuditLogProvisioningId{
		ProvisioningObjectSummaryId: provisioningObjectSummaryId,
	}
}

// ParseAuditLogProvisioningID parses 'input' into a AuditLogProvisioningId
func ParseAuditLogProvisioningID(input string) (*AuditLogProvisioningId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AuditLogProvisioningId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AuditLogProvisioningId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAuditLogProvisioningIDInsensitively parses 'input' case-insensitively into a AuditLogProvisioningId
// note: this method should only be used for API response data and not user input
func ParseAuditLogProvisioningIDInsensitively(input string) (*AuditLogProvisioningId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AuditLogProvisioningId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AuditLogProvisioningId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AuditLogProvisioningId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ProvisioningObjectSummaryId, ok = input.Parsed["provisioningObjectSummaryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "provisioningObjectSummaryId", input)
	}

	return nil
}

// ValidateAuditLogProvisioningID checks that 'input' can be parsed as a Audit Log Provisioning ID
func ValidateAuditLogProvisioningID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAuditLogProvisioningID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Audit Log Provisioning ID
func (id AuditLogProvisioningId) ID() string {
	fmtString := "/auditLogs/provisioning/%s"
	return fmt.Sprintf(fmtString, id.ProvisioningObjectSummaryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Audit Log Provisioning ID
func (id AuditLogProvisioningId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("auditLogs", "auditLogs", "auditLogs"),
		resourceids.StaticSegment("provisioning", "provisioning", "provisioning"),
		resourceids.UserSpecifiedSegment("provisioningObjectSummaryId", "provisioningObjectSummaryId"),
	}
}

// String returns a human-readable description of this Audit Log Provisioning ID
func (id AuditLogProvisioningId) String() string {
	components := []string{
		fmt.Sprintf("Provisioning Object Summary: %q", id.ProvisioningObjectSummaryId),
	}
	return fmt.Sprintf("Audit Log Provisioning (%s)", strings.Join(components, "\n"))
}
