package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AuditLogDirectoryProvisioningId{}

// AuditLogDirectoryProvisioningId is a struct representing the Resource ID for a Audit Log Directory Provisioning
type AuditLogDirectoryProvisioningId struct {
	ProvisioningObjectSummaryId string
}

// NewAuditLogDirectoryProvisioningID returns a new AuditLogDirectoryProvisioningId struct
func NewAuditLogDirectoryProvisioningID(provisioningObjectSummaryId string) AuditLogDirectoryProvisioningId {
	return AuditLogDirectoryProvisioningId{
		ProvisioningObjectSummaryId: provisioningObjectSummaryId,
	}
}

// ParseAuditLogDirectoryProvisioningID parses 'input' into a AuditLogDirectoryProvisioningId
func ParseAuditLogDirectoryProvisioningID(input string) (*AuditLogDirectoryProvisioningId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AuditLogDirectoryProvisioningId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AuditLogDirectoryProvisioningId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAuditLogDirectoryProvisioningIDInsensitively parses 'input' case-insensitively into a AuditLogDirectoryProvisioningId
// note: this method should only be used for API response data and not user input
func ParseAuditLogDirectoryProvisioningIDInsensitively(input string) (*AuditLogDirectoryProvisioningId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AuditLogDirectoryProvisioningId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AuditLogDirectoryProvisioningId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AuditLogDirectoryProvisioningId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ProvisioningObjectSummaryId, ok = input.Parsed["provisioningObjectSummaryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "provisioningObjectSummaryId", input)
	}

	return nil
}

// ValidateAuditLogDirectoryProvisioningID checks that 'input' can be parsed as a Audit Log Directory Provisioning ID
func ValidateAuditLogDirectoryProvisioningID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAuditLogDirectoryProvisioningID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Audit Log Directory Provisioning ID
func (id AuditLogDirectoryProvisioningId) ID() string {
	fmtString := "/auditLogs/directoryProvisioning/%s"
	return fmt.Sprintf(fmtString, id.ProvisioningObjectSummaryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Audit Log Directory Provisioning ID
func (id AuditLogDirectoryProvisioningId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("auditLogs", "auditLogs", "auditLogs"),
		resourceids.StaticSegment("directoryProvisioning", "directoryProvisioning", "directoryProvisioning"),
		resourceids.UserSpecifiedSegment("provisioningObjectSummaryId", "provisioningObjectSummaryId"),
	}
}

// String returns a human-readable description of this Audit Log Directory Provisioning ID
func (id AuditLogDirectoryProvisioningId) String() string {
	components := []string{
		fmt.Sprintf("Provisioning Object Summary: %q", id.ProvisioningObjectSummaryId),
	}
	return fmt.Sprintf("Audit Log Directory Provisioning (%s)", strings.Join(components, "\n"))
}
