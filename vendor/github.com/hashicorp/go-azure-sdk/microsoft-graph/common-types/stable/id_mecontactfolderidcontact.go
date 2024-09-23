package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeContactFolderIdContactId{}

// MeContactFolderIdContactId is a struct representing the Resource ID for a Me Contact Folder Id Contact
type MeContactFolderIdContactId struct {
	ContactFolderId string
	ContactId       string
}

// NewMeContactFolderIdContactID returns a new MeContactFolderIdContactId struct
func NewMeContactFolderIdContactID(contactFolderId string, contactId string) MeContactFolderIdContactId {
	return MeContactFolderIdContactId{
		ContactFolderId: contactFolderId,
		ContactId:       contactId,
	}
}

// ParseMeContactFolderIdContactID parses 'input' into a MeContactFolderIdContactId
func ParseMeContactFolderIdContactID(input string) (*MeContactFolderIdContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeContactFolderIdContactId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeContactFolderIdContactId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeContactFolderIdContactIDInsensitively parses 'input' case-insensitively into a MeContactFolderIdContactId
// note: this method should only be used for API response data and not user input
func ParseMeContactFolderIdContactIDInsensitively(input string) (*MeContactFolderIdContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeContactFolderIdContactId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeContactFolderIdContactId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeContactFolderIdContactId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ContactFolderId, ok = input.Parsed["contactFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId", input)
	}

	if id.ContactId, ok = input.Parsed["contactId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contactId", input)
	}

	return nil
}

// ValidateMeContactFolderIdContactID checks that 'input' can be parsed as a Me Contact Folder Id Contact ID
func ValidateMeContactFolderIdContactID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeContactFolderIdContactID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Contact Folder Id Contact ID
func (id MeContactFolderIdContactId) ID() string {
	fmtString := "/me/contactFolders/%s/contacts/%s"
	return fmt.Sprintf(fmtString, id.ContactFolderId, id.ContactId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Contact Folder Id Contact ID
func (id MeContactFolderIdContactId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("contactFolders", "contactFolders", "contactFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId", "contactFolderId"),
		resourceids.StaticSegment("contacts", "contacts", "contacts"),
		resourceids.UserSpecifiedSegment("contactId", "contactId"),
	}
}

// String returns a human-readable description of this Me Contact Folder Id Contact ID
func (id MeContactFolderIdContactId) String() string {
	components := []string{
		fmt.Sprintf("Contact Folder: %q", id.ContactFolderId),
		fmt.Sprintf("Contact: %q", id.ContactId),
	}
	return fmt.Sprintf("Me Contact Folder Id Contact (%s)", strings.Join(components, "\n"))
}
