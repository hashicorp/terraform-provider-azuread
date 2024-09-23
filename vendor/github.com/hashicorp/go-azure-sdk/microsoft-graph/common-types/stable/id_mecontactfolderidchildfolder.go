package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeContactFolderIdChildFolderId{}

// MeContactFolderIdChildFolderId is a struct representing the Resource ID for a Me Contact Folder Id Child Folder
type MeContactFolderIdChildFolderId struct {
	ContactFolderId  string
	ContactFolderId1 string
}

// NewMeContactFolderIdChildFolderID returns a new MeContactFolderIdChildFolderId struct
func NewMeContactFolderIdChildFolderID(contactFolderId string, contactFolderId1 string) MeContactFolderIdChildFolderId {
	return MeContactFolderIdChildFolderId{
		ContactFolderId:  contactFolderId,
		ContactFolderId1: contactFolderId1,
	}
}

// ParseMeContactFolderIdChildFolderID parses 'input' into a MeContactFolderIdChildFolderId
func ParseMeContactFolderIdChildFolderID(input string) (*MeContactFolderIdChildFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeContactFolderIdChildFolderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeContactFolderIdChildFolderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeContactFolderIdChildFolderIDInsensitively parses 'input' case-insensitively into a MeContactFolderIdChildFolderId
// note: this method should only be used for API response data and not user input
func ParseMeContactFolderIdChildFolderIDInsensitively(input string) (*MeContactFolderIdChildFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeContactFolderIdChildFolderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeContactFolderIdChildFolderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeContactFolderIdChildFolderId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ContactFolderId, ok = input.Parsed["contactFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId", input)
	}

	if id.ContactFolderId1, ok = input.Parsed["contactFolderId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId1", input)
	}

	return nil
}

// ValidateMeContactFolderIdChildFolderID checks that 'input' can be parsed as a Me Contact Folder Id Child Folder ID
func ValidateMeContactFolderIdChildFolderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeContactFolderIdChildFolderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Contact Folder Id Child Folder ID
func (id MeContactFolderIdChildFolderId) ID() string {
	fmtString := "/me/contactFolders/%s/childFolders/%s"
	return fmt.Sprintf(fmtString, id.ContactFolderId, id.ContactFolderId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Contact Folder Id Child Folder ID
func (id MeContactFolderIdChildFolderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("contactFolders", "contactFolders", "contactFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId", "contactFolderId"),
		resourceids.StaticSegment("childFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId1", "contactFolderId1"),
	}
}

// String returns a human-readable description of this Me Contact Folder Id Child Folder ID
func (id MeContactFolderIdChildFolderId) String() string {
	components := []string{
		fmt.Sprintf("Contact Folder: %q", id.ContactFolderId),
		fmt.Sprintf("Contact Folder Id 1: %q", id.ContactFolderId1),
	}
	return fmt.Sprintf("Me Contact Folder Id Child Folder (%s)", strings.Join(components, "\n"))
}
