package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMailFolderId{}

// MeMailFolderId is a struct representing the Resource ID for a Me Mail Folder
type MeMailFolderId struct {
	MailFolderId string
}

// NewMeMailFolderID returns a new MeMailFolderId struct
func NewMeMailFolderID(mailFolderId string) MeMailFolderId {
	return MeMailFolderId{
		MailFolderId: mailFolderId,
	}
}

// ParseMeMailFolderID parses 'input' into a MeMailFolderId
func ParseMeMailFolderID(input string) (*MeMailFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeMailFolderIDInsensitively parses 'input' case-insensitively into a MeMailFolderId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderIDInsensitively(input string) (*MeMailFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeMailFolderId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MailFolderId, ok = input.Parsed["mailFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", input)
	}

	return nil
}

// ValidateMeMailFolderID checks that 'input' can be parsed as a Me Mail Folder ID
func ValidateMeMailFolderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder ID
func (id MeMailFolderId) ID() string {
	fmtString := "/me/mailFolders/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder ID
func (id MeMailFolderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
	}
}

// String returns a human-readable description of this Me Mail Folder ID
func (id MeMailFolderId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
	}
	return fmt.Sprintf("Me Mail Folder (%s)", strings.Join(components, "\n"))
}
