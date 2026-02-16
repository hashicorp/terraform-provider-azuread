package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeContactFolderId{}

// MeContactFolderId is a struct representing the Resource ID for a Me Contact Folder
type MeContactFolderId struct {
	ContactFolderId string
}

// NewMeContactFolderID returns a new MeContactFolderId struct
func NewMeContactFolderID(contactFolderId string) MeContactFolderId {
	return MeContactFolderId{
		ContactFolderId: contactFolderId,
	}
}

// ParseMeContactFolderID parses 'input' into a MeContactFolderId
func ParseMeContactFolderID(input string) (*MeContactFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeContactFolderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeContactFolderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeContactFolderIDInsensitively parses 'input' case-insensitively into a MeContactFolderId
// note: this method should only be used for API response data and not user input
func ParseMeContactFolderIDInsensitively(input string) (*MeContactFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeContactFolderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeContactFolderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeContactFolderId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ContactFolderId, ok = input.Parsed["contactFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId", input)
	}

	return nil
}

// ValidateMeContactFolderID checks that 'input' can be parsed as a Me Contact Folder ID
func ValidateMeContactFolderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeContactFolderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Contact Folder ID
func (id MeContactFolderId) ID() string {
	fmtString := "/me/contactFolders/%s"
	return fmt.Sprintf(fmtString, id.ContactFolderId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Contact Folder ID
func (id MeContactFolderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("contactFolders", "contactFolders", "contactFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId", "contactFolderId"),
	}
}

// String returns a human-readable description of this Me Contact Folder ID
func (id MeContactFolderId) String() string {
	components := []string{
		fmt.Sprintf("Contact Folder: %q", id.ContactFolderId),
	}
	return fmt.Sprintf("Me Contact Folder (%s)", strings.Join(components, "\n"))
}
