package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeContactFolderIdChildFolderIdContactIdExtensionId{}

// MeContactFolderIdChildFolderIdContactIdExtensionId is a struct representing the Resource ID for a Me Contact Folder Id Child Folder Id Contact Id Extension
type MeContactFolderIdChildFolderIdContactIdExtensionId struct {
	ContactFolderId  string
	ContactFolderId1 string
	ContactId        string
	ExtensionId      string
}

// NewMeContactFolderIdChildFolderIdContactIdExtensionID returns a new MeContactFolderIdChildFolderIdContactIdExtensionId struct
func NewMeContactFolderIdChildFolderIdContactIdExtensionID(contactFolderId string, contactFolderId1 string, contactId string, extensionId string) MeContactFolderIdChildFolderIdContactIdExtensionId {
	return MeContactFolderIdChildFolderIdContactIdExtensionId{
		ContactFolderId:  contactFolderId,
		ContactFolderId1: contactFolderId1,
		ContactId:        contactId,
		ExtensionId:      extensionId,
	}
}

// ParseMeContactFolderIdChildFolderIdContactIdExtensionID parses 'input' into a MeContactFolderIdChildFolderIdContactIdExtensionId
func ParseMeContactFolderIdChildFolderIdContactIdExtensionID(input string) (*MeContactFolderIdChildFolderIdContactIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeContactFolderIdChildFolderIdContactIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeContactFolderIdChildFolderIdContactIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeContactFolderIdChildFolderIdContactIdExtensionIDInsensitively parses 'input' case-insensitively into a MeContactFolderIdChildFolderIdContactIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeContactFolderIdChildFolderIdContactIdExtensionIDInsensitively(input string) (*MeContactFolderIdChildFolderIdContactIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeContactFolderIdChildFolderIdContactIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeContactFolderIdChildFolderIdContactIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeContactFolderIdChildFolderIdContactIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateMeContactFolderIdChildFolderIdContactIdExtensionID checks that 'input' can be parsed as a Me Contact Folder Id Child Folder Id Contact Id Extension ID
func ValidateMeContactFolderIdChildFolderIdContactIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeContactFolderIdChildFolderIdContactIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Contact Folder Id Child Folder Id Contact Id Extension ID
func (id MeContactFolderIdChildFolderIdContactIdExtensionId) ID() string {
	fmtString := "/me/contactFolders/%s/childFolders/%s/contacts/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.ContactFolderId, id.ContactFolderId1, id.ContactId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Contact Folder Id Child Folder Id Contact Id Extension ID
func (id MeContactFolderIdChildFolderIdContactIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("contactFolders", "contactFolders", "contactFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId", "contactFolderId"),
		resourceids.StaticSegment("childFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId1", "contactFolderId1"),
		resourceids.StaticSegment("contacts", "contacts", "contacts"),
		resourceids.UserSpecifiedSegment("contactId", "contactId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Me Contact Folder Id Child Folder Id Contact Id Extension ID
func (id MeContactFolderIdChildFolderIdContactIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Contact Folder: %q", id.ContactFolderId),
		fmt.Sprintf("Contact Folder Id 1: %q", id.ContactFolderId1),
		fmt.Sprintf("Contact: %q", id.ContactId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Contact Folder Id Child Folder Id Contact Id Extension (%s)", strings.Join(components, "\n"))
}
