package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AuditLogSignUpId{}

// AuditLogSignUpId is a struct representing the Resource ID for a Audit Log Sign Up
type AuditLogSignUpId struct {
	SelfServiceSignUpId string
}

// NewAuditLogSignUpID returns a new AuditLogSignUpId struct
func NewAuditLogSignUpID(selfServiceSignUpId string) AuditLogSignUpId {
	return AuditLogSignUpId{
		SelfServiceSignUpId: selfServiceSignUpId,
	}
}

// ParseAuditLogSignUpID parses 'input' into a AuditLogSignUpId
func ParseAuditLogSignUpID(input string) (*AuditLogSignUpId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AuditLogSignUpId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AuditLogSignUpId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAuditLogSignUpIDInsensitively parses 'input' case-insensitively into a AuditLogSignUpId
// note: this method should only be used for API response data and not user input
func ParseAuditLogSignUpIDInsensitively(input string) (*AuditLogSignUpId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AuditLogSignUpId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AuditLogSignUpId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AuditLogSignUpId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SelfServiceSignUpId, ok = input.Parsed["selfServiceSignUpId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "selfServiceSignUpId", input)
	}

	return nil
}

// ValidateAuditLogSignUpID checks that 'input' can be parsed as a Audit Log Sign Up ID
func ValidateAuditLogSignUpID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAuditLogSignUpID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Audit Log Sign Up ID
func (id AuditLogSignUpId) ID() string {
	fmtString := "/auditLogs/signUps/%s"
	return fmt.Sprintf(fmtString, id.SelfServiceSignUpId)
}

// Segments returns a slice of Resource ID Segments which comprise this Audit Log Sign Up ID
func (id AuditLogSignUpId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("auditLogs", "auditLogs", "auditLogs"),
		resourceids.StaticSegment("signUps", "signUps", "signUps"),
		resourceids.UserSpecifiedSegment("selfServiceSignUpId", "selfServiceSignUpId"),
	}
}

// String returns a human-readable description of this Audit Log Sign Up ID
func (id AuditLogSignUpId) String() string {
	components := []string{
		fmt.Sprintf("Self Service Sign Up: %q", id.SelfServiceSignUpId),
	}
	return fmt.Sprintf("Audit Log Sign Up (%s)", strings.Join(components, "\n"))
}
