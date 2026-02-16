package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMailFolderIdOperationId{}

// MeMailFolderIdOperationId is a struct representing the Resource ID for a Me Mail Folder Id Operation
type MeMailFolderIdOperationId struct {
	MailFolderId          string
	MailFolderOperationId string
}

// NewMeMailFolderIdOperationID returns a new MeMailFolderIdOperationId struct
func NewMeMailFolderIdOperationID(mailFolderId string, mailFolderOperationId string) MeMailFolderIdOperationId {
	return MeMailFolderIdOperationId{
		MailFolderId:          mailFolderId,
		MailFolderOperationId: mailFolderOperationId,
	}
}

// ParseMeMailFolderIdOperationID parses 'input' into a MeMailFolderIdOperationId
func ParseMeMailFolderIdOperationID(input string) (*MeMailFolderIdOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeMailFolderIdOperationIDInsensitively parses 'input' case-insensitively into a MeMailFolderIdOperationId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderIdOperationIDInsensitively(input string) (*MeMailFolderIdOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeMailFolderIdOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MailFolderId, ok = input.Parsed["mailFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", input)
	}

	if id.MailFolderOperationId, ok = input.Parsed["mailFolderOperationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderOperationId", input)
	}

	return nil
}

// ValidateMeMailFolderIdOperationID checks that 'input' can be parsed as a Me Mail Folder Id Operation ID
func ValidateMeMailFolderIdOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderIdOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Id Operation ID
func (id MeMailFolderIdOperationId) ID() string {
	fmtString := "/me/mailFolders/%s/operations/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MailFolderOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Id Operation ID
func (id MeMailFolderIdOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("operations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("mailFolderOperationId", "mailFolderOperationId"),
	}
}

// String returns a human-readable description of this Me Mail Folder Id Operation ID
func (id MeMailFolderIdOperationId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Operation: %q", id.MailFolderOperationId),
	}
	return fmt.Sprintf("Me Mail Folder Id Operation (%s)", strings.Join(components, "\n"))
}
