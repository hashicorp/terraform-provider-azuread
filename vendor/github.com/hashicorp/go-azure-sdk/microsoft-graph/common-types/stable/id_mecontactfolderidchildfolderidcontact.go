package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeContactFolderIdChildFolderIdContactId{}

// MeContactFolderIdChildFolderIdContactId is a struct representing the Resource ID for a Me Contact Folder Id Child Folder Id Contact
type MeContactFolderIdChildFolderIdContactId struct {
	ContactFolderId  string
	ContactFolderId1 string
	ContactId        string
}

// NewMeContactFolderIdChildFolderIdContactID returns a new MeContactFolderIdChildFolderIdContactId struct
func NewMeContactFolderIdChildFolderIdContactID(contactFolderId string, contactFolderId1 string, contactId string) MeContactFolderIdChildFolderIdContactId {
	return MeContactFolderIdChildFolderIdContactId{
		ContactFolderId:  contactFolderId,
		ContactFolderId1: contactFolderId1,
		ContactId:        contactId,
	}
}

// ParseMeContactFolderIdChildFolderIdContactID parses 'input' into a MeContactFolderIdChildFolderIdContactId
func ParseMeContactFolderIdChildFolderIdContactID(input string) (*MeContactFolderIdChildFolderIdContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeContactFolderIdChildFolderIdContactId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeContactFolderIdChildFolderIdContactId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeContactFolderIdChildFolderIdContactIDInsensitively parses 'input' case-insensitively into a MeContactFolderIdChildFolderIdContactId
// note: this method should only be used for API response data and not user input
func ParseMeContactFolderIdChildFolderIdContactIDInsensitively(input string) (*MeContactFolderIdChildFolderIdContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeContactFolderIdChildFolderIdContactId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeContactFolderIdChildFolderIdContactId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeContactFolderIdChildFolderIdContactId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ContactFolderId, ok = input.Parsed["contactFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId", input)
	}

	if id.ContactFolderId1, ok = input.Parsed["contactFolderId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId1", input)
	}

	if id.ContactId, ok = input.Parsed["contactId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contactId", input)
	}

	return nil
}

// ValidateMeContactFolderIdChildFolderIdContactID checks that 'input' can be parsed as a Me Contact Folder Id Child Folder Id Contact ID
func ValidateMeContactFolderIdChildFolderIdContactID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeContactFolderIdChildFolderIdContactID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Contact Folder Id Child Folder Id Contact ID
func (id MeContactFolderIdChildFolderIdContactId) ID() string {
	fmtString := "/me/contactFolders/%s/childFolders/%s/contacts/%s"
	return fmt.Sprintf(fmtString, id.ContactFolderId, id.ContactFolderId1, id.ContactId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Contact Folder Id Child Folder Id Contact ID
func (id MeContactFolderIdChildFolderIdContactId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("contactFolders", "contactFolders", "contactFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId", "contactFolderId"),
		resourceids.StaticSegment("childFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId1", "contactFolderId1"),
		resourceids.StaticSegment("contacts", "contacts", "contacts"),
		resourceids.UserSpecifiedSegment("contactId", "contactId"),
	}
}

// String returns a human-readable description of this Me Contact Folder Id Child Folder Id Contact ID
func (id MeContactFolderIdChildFolderIdContactId) String() string {
	components := []string{
		fmt.Sprintf("Contact Folder: %q", id.ContactFolderId),
		fmt.Sprintf("Contact Folder Id 1: %q", id.ContactFolderId1),
		fmt.Sprintf("Contact: %q", id.ContactId),
	}
	return fmt.Sprintf("Me Contact Folder Id Child Folder Id Contact (%s)", strings.Join(components, "\n"))
}
