package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AuditLogCustomSecurityAttributeAuditId{}

// AuditLogCustomSecurityAttributeAuditId is a struct representing the Resource ID for a Audit Log Custom Security Attribute Audit
type AuditLogCustomSecurityAttributeAuditId struct {
	CustomSecurityAttributeAuditId string
}

// NewAuditLogCustomSecurityAttributeAuditID returns a new AuditLogCustomSecurityAttributeAuditId struct
func NewAuditLogCustomSecurityAttributeAuditID(customSecurityAttributeAuditId string) AuditLogCustomSecurityAttributeAuditId {
	return AuditLogCustomSecurityAttributeAuditId{
		CustomSecurityAttributeAuditId: customSecurityAttributeAuditId,
	}
}

// ParseAuditLogCustomSecurityAttributeAuditID parses 'input' into a AuditLogCustomSecurityAttributeAuditId
func ParseAuditLogCustomSecurityAttributeAuditID(input string) (*AuditLogCustomSecurityAttributeAuditId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AuditLogCustomSecurityAttributeAuditId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AuditLogCustomSecurityAttributeAuditId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAuditLogCustomSecurityAttributeAuditIDInsensitively parses 'input' case-insensitively into a AuditLogCustomSecurityAttributeAuditId
// note: this method should only be used for API response data and not user input
func ParseAuditLogCustomSecurityAttributeAuditIDInsensitively(input string) (*AuditLogCustomSecurityAttributeAuditId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AuditLogCustomSecurityAttributeAuditId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AuditLogCustomSecurityAttributeAuditId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AuditLogCustomSecurityAttributeAuditId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CustomSecurityAttributeAuditId, ok = input.Parsed["customSecurityAttributeAuditId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "customSecurityAttributeAuditId", input)
	}

	return nil
}

// ValidateAuditLogCustomSecurityAttributeAuditID checks that 'input' can be parsed as a Audit Log Custom Security Attribute Audit ID
func ValidateAuditLogCustomSecurityAttributeAuditID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAuditLogCustomSecurityAttributeAuditID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Audit Log Custom Security Attribute Audit ID
func (id AuditLogCustomSecurityAttributeAuditId) ID() string {
	fmtString := "/auditLogs/customSecurityAttributeAudits/%s"
	return fmt.Sprintf(fmtString, id.CustomSecurityAttributeAuditId)
}

// Segments returns a slice of Resource ID Segments which comprise this Audit Log Custom Security Attribute Audit ID
func (id AuditLogCustomSecurityAttributeAuditId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("auditLogs", "auditLogs", "auditLogs"),
		resourceids.StaticSegment("customSecurityAttributeAudits", "customSecurityAttributeAudits", "customSecurityAttributeAudits"),
		resourceids.UserSpecifiedSegment("customSecurityAttributeAuditId", "customSecurityAttributeAuditId"),
	}
}

// String returns a human-readable description of this Audit Log Custom Security Attribute Audit ID
func (id AuditLogCustomSecurityAttributeAuditId) String() string {
	components := []string{
		fmt.Sprintf("Custom Security Attribute Audit: %q", id.CustomSecurityAttributeAuditId),
	}
	return fmt.Sprintf("Audit Log Custom Security Attribute Audit (%s)", strings.Join(components, "\n"))
}
