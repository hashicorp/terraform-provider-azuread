package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeContactFolderIdContactIdExtensionId{}

// MeContactFolderIdContactIdExtensionId is a struct representing the Resource ID for a Me Contact Folder Id Contact Id Extension
type MeContactFolderIdContactIdExtensionId struct {
	ContactFolderId string
	ContactId       string
	ExtensionId     string
}

// NewMeContactFolderIdContactIdExtensionID returns a new MeContactFolderIdContactIdExtensionId struct
func NewMeContactFolderIdContactIdExtensionID(contactFolderId string, contactId string, extensionId string) MeContactFolderIdContactIdExtensionId {
	return MeContactFolderIdContactIdExtensionId{
		ContactFolderId: contactFolderId,
		ContactId:       contactId,
		ExtensionId:     extensionId,
	}
}

// ParseMeContactFolderIdContactIdExtensionID parses 'input' into a MeContactFolderIdContactIdExtensionId
func ParseMeContactFolderIdContactIdExtensionID(input string) (*MeContactFolderIdContactIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeContactFolderIdContactIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeContactFolderIdContactIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeContactFolderIdContactIdExtensionIDInsensitively parses 'input' case-insensitively into a MeContactFolderIdContactIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeContactFolderIdContactIdExtensionIDInsensitively(input string) (*MeContactFolderIdContactIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeContactFolderIdContactIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeContactFolderIdContactIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeContactFolderIdContactIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ContactFolderId, ok = input.Parsed["contactFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId", input)
	}

	if id.ContactId, ok = input.Parsed["contactId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contactId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateMeContactFolderIdContactIdExtensionID checks that 'input' can be parsed as a Me Contact Folder Id Contact Id Extension ID
func ValidateMeContactFolderIdContactIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeContactFolderIdContactIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Contact Folder Id Contact Id Extension ID
func (id MeContactFolderIdContactIdExtensionId) ID() string {
	fmtString := "/me/contactFolders/%s/contacts/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.ContactFolderId, id.ContactId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Contact Folder Id Contact Id Extension ID
func (id MeContactFolderIdContactIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("contactFolders", "contactFolders", "contactFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId", "contactFolderId"),
		resourceids.StaticSegment("contacts", "contacts", "contacts"),
		resourceids.UserSpecifiedSegment("contactId", "contactId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Me Contact Folder Id Contact Id Extension ID
func (id MeContactFolderIdContactIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Contact Folder: %q", id.ContactFolderId),
		fmt.Sprintf("Contact: %q", id.ContactId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Contact Folder Id Contact Id Extension (%s)", strings.Join(components, "\n"))
}
