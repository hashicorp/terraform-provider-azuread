package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeContactId{}

// MeContactId is a struct representing the Resource ID for a Me Contact
type MeContactId struct {
	ContactId string
}

// NewMeContactID returns a new MeContactId struct
func NewMeContactID(contactId string) MeContactId {
	return MeContactId{
		ContactId: contactId,
	}
}

// ParseMeContactID parses 'input' into a MeContactId
func ParseMeContactID(input string) (*MeContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeContactId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeContactId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeContactIDInsensitively parses 'input' case-insensitively into a MeContactId
// note: this method should only be used for API response data and not user input
func ParseMeContactIDInsensitively(input string) (*MeContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeContactId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeContactId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeContactId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ContactId, ok = input.Parsed["contactId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contactId", input)
	}

	return nil
}

// ValidateMeContactID checks that 'input' can be parsed as a Me Contact ID
func ValidateMeContactID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeContactID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Contact ID
func (id MeContactId) ID() string {
	fmtString := "/me/contacts/%s"
	return fmt.Sprintf(fmtString, id.ContactId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Contact ID
func (id MeContactId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("contacts", "contacts", "contacts"),
		resourceids.UserSpecifiedSegment("contactId", "contactId"),
	}
}

// String returns a human-readable description of this Me Contact ID
func (id MeContactId) String() string {
	components := []string{
		fmt.Sprintf("Contact: %q", id.ContactId),
	}
	return fmt.Sprintf("Me Contact (%s)", strings.Join(components, "\n"))
}
