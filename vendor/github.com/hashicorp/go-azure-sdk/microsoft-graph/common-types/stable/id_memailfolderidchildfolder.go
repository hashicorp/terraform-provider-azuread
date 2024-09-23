package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMailFolderIdChildFolderId{}

// MeMailFolderIdChildFolderId is a struct representing the Resource ID for a Me Mail Folder Id Child Folder
type MeMailFolderIdChildFolderId struct {
	MailFolderId  string
	MailFolderId1 string
}

// NewMeMailFolderIdChildFolderID returns a new MeMailFolderIdChildFolderId struct
func NewMeMailFolderIdChildFolderID(mailFolderId string, mailFolderId1 string) MeMailFolderIdChildFolderId {
	return MeMailFolderIdChildFolderId{
		MailFolderId:  mailFolderId,
		MailFolderId1: mailFolderId1,
	}
}

// ParseMeMailFolderIdChildFolderID parses 'input' into a MeMailFolderIdChildFolderId
func ParseMeMailFolderIdChildFolderID(input string) (*MeMailFolderIdChildFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdChildFolderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdChildFolderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeMailFolderIdChildFolderIDInsensitively parses 'input' case-insensitively into a MeMailFolderIdChildFolderId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderIdChildFolderIDInsensitively(input string) (*MeMailFolderIdChildFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdChildFolderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdChildFolderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeMailFolderIdChildFolderId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MailFolderId, ok = input.Parsed["mailFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", input)
	}

	if id.MailFolderId1, ok = input.Parsed["mailFolderId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId1", input)
	}

	return nil
}

// ValidateMeMailFolderIdChildFolderID checks that 'input' can be parsed as a Me Mail Folder Id Child Folder ID
func ValidateMeMailFolderIdChildFolderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderIdChildFolderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Id Child Folder ID
func (id MeMailFolderIdChildFolderId) ID() string {
	fmtString := "/me/mailFolders/%s/childFolders/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MailFolderId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Id Child Folder ID
func (id MeMailFolderIdChildFolderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("childFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1"),
	}
}

// String returns a human-readable description of this Me Mail Folder Id Child Folder ID
func (id MeMailFolderIdChildFolderId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
	}
	return fmt.Sprintf("Me Mail Folder Id Child Folder (%s)", strings.Join(components, "\n"))
}
