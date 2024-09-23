package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeContactIdExtensionId{}

// MeContactIdExtensionId is a struct representing the Resource ID for a Me Contact Id Extension
type MeContactIdExtensionId struct {
	ContactId   string
	ExtensionId string
}

// NewMeContactIdExtensionID returns a new MeContactIdExtensionId struct
func NewMeContactIdExtensionID(contactId string, extensionId string) MeContactIdExtensionId {
	return MeContactIdExtensionId{
		ContactId:   contactId,
		ExtensionId: extensionId,
	}
}

// ParseMeContactIdExtensionID parses 'input' into a MeContactIdExtensionId
func ParseMeContactIdExtensionID(input string) (*MeContactIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeContactIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeContactIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeContactIdExtensionIDInsensitively parses 'input' case-insensitively into a MeContactIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeContactIdExtensionIDInsensitively(input string) (*MeContactIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeContactIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeContactIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeContactIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ContactId, ok = input.Parsed["contactId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contactId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateMeContactIdExtensionID checks that 'input' can be parsed as a Me Contact Id Extension ID
func ValidateMeContactIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeContactIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Contact Id Extension ID
func (id MeContactIdExtensionId) ID() string {
	fmtString := "/me/contacts/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.ContactId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Contact Id Extension ID
func (id MeContactIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("contacts", "contacts", "contacts"),
		resourceids.UserSpecifiedSegment("contactId", "contactId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Me Contact Id Extension ID
func (id MeContactIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Contact: %q", id.ContactId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Contact Id Extension (%s)", strings.Join(components, "\n"))
}
