package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AuditLogSignInId{}

// AuditLogSignInId is a struct representing the Resource ID for a Audit Log Sign In
type AuditLogSignInId struct {
	SignInId string
}

// NewAuditLogSignInID returns a new AuditLogSignInId struct
func NewAuditLogSignInID(signInId string) AuditLogSignInId {
	return AuditLogSignInId{
		SignInId: signInId,
	}
}

// ParseAuditLogSignInID parses 'input' into a AuditLogSignInId
func ParseAuditLogSignInID(input string) (*AuditLogSignInId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AuditLogSignInId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AuditLogSignInId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAuditLogSignInIDInsensitively parses 'input' case-insensitively into a AuditLogSignInId
// note: this method should only be used for API response data and not user input
func ParseAuditLogSignInIDInsensitively(input string) (*AuditLogSignInId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AuditLogSignInId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AuditLogSignInId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AuditLogSignInId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SignInId, ok = input.Parsed["signInId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "signInId", input)
	}

	return nil
}

// ValidateAuditLogSignInID checks that 'input' can be parsed as a Audit Log Sign In ID
func ValidateAuditLogSignInID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAuditLogSignInID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Audit Log Sign In ID
func (id AuditLogSignInId) ID() string {
	fmtString := "/auditLogs/signIns/%s"
	return fmt.Sprintf(fmtString, id.SignInId)
}

// Segments returns a slice of Resource ID Segments which comprise this Audit Log Sign In ID
func (id AuditLogSignInId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("auditLogs", "auditLogs", "auditLogs"),
		resourceids.StaticSegment("signIns", "signIns", "signIns"),
		resourceids.UserSpecifiedSegment("signInId", "signInId"),
	}
}

// String returns a human-readable description of this Audit Log Sign In ID
func (id AuditLogSignInId) String() string {
	components := []string{
		fmt.Sprintf("Sign In: %q", id.SignInId),
	}
	return fmt.Sprintf("Audit Log Sign In (%s)", strings.Join(components, "\n"))
}
