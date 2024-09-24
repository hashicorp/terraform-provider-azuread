package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AuditLogDirectoryAuditId{}

// AuditLogDirectoryAuditId is a struct representing the Resource ID for a Audit Log Directory Audit
type AuditLogDirectoryAuditId struct {
	DirectoryAuditId string
}

// NewAuditLogDirectoryAuditID returns a new AuditLogDirectoryAuditId struct
func NewAuditLogDirectoryAuditID(directoryAuditId string) AuditLogDirectoryAuditId {
	return AuditLogDirectoryAuditId{
		DirectoryAuditId: directoryAuditId,
	}
}

// ParseAuditLogDirectoryAuditID parses 'input' into a AuditLogDirectoryAuditId
func ParseAuditLogDirectoryAuditID(input string) (*AuditLogDirectoryAuditId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AuditLogDirectoryAuditId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AuditLogDirectoryAuditId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAuditLogDirectoryAuditIDInsensitively parses 'input' case-insensitively into a AuditLogDirectoryAuditId
// note: this method should only be used for API response data and not user input
func ParseAuditLogDirectoryAuditIDInsensitively(input string) (*AuditLogDirectoryAuditId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AuditLogDirectoryAuditId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AuditLogDirectoryAuditId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AuditLogDirectoryAuditId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DirectoryAuditId, ok = input.Parsed["directoryAuditId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryAuditId", input)
	}

	return nil
}

// ValidateAuditLogDirectoryAuditID checks that 'input' can be parsed as a Audit Log Directory Audit ID
func ValidateAuditLogDirectoryAuditID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAuditLogDirectoryAuditID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Audit Log Directory Audit ID
func (id AuditLogDirectoryAuditId) ID() string {
	fmtString := "/auditLogs/directoryAudits/%s"
	return fmt.Sprintf(fmtString, id.DirectoryAuditId)
}

// Segments returns a slice of Resource ID Segments which comprise this Audit Log Directory Audit ID
func (id AuditLogDirectoryAuditId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("auditLogs", "auditLogs", "auditLogs"),
		resourceids.StaticSegment("directoryAudits", "directoryAudits", "directoryAudits"),
		resourceids.UserSpecifiedSegment("directoryAuditId", "directoryAuditId"),
	}
}

// String returns a human-readable description of this Audit Log Directory Audit ID
func (id AuditLogDirectoryAuditId) String() string {
	components := []string{
		fmt.Sprintf("Directory Audit: %q", id.DirectoryAuditId),
	}
	return fmt.Sprintf("Audit Log Directory Audit (%s)", strings.Join(components, "\n"))
}
