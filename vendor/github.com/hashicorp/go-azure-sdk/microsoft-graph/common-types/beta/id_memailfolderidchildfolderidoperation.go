package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMailFolderIdChildFolderIdOperationId{}

// MeMailFolderIdChildFolderIdOperationId is a struct representing the Resource ID for a Me Mail Folder Id Child Folder Id Operation
type MeMailFolderIdChildFolderIdOperationId struct {
	MailFolderId          string
	MailFolderId1         string
	MailFolderOperationId string
}

// NewMeMailFolderIdChildFolderIdOperationID returns a new MeMailFolderIdChildFolderIdOperationId struct
func NewMeMailFolderIdChildFolderIdOperationID(mailFolderId string, mailFolderId1 string, mailFolderOperationId string) MeMailFolderIdChildFolderIdOperationId {
	return MeMailFolderIdChildFolderIdOperationId{
		MailFolderId:          mailFolderId,
		MailFolderId1:         mailFolderId1,
		MailFolderOperationId: mailFolderOperationId,
	}
}

// ParseMeMailFolderIdChildFolderIdOperationID parses 'input' into a MeMailFolderIdChildFolderIdOperationId
func ParseMeMailFolderIdChildFolderIdOperationID(input string) (*MeMailFolderIdChildFolderIdOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdChildFolderIdOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdChildFolderIdOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeMailFolderIdChildFolderIdOperationIDInsensitively parses 'input' case-insensitively into a MeMailFolderIdChildFolderIdOperationId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderIdChildFolderIdOperationIDInsensitively(input string) (*MeMailFolderIdChildFolderIdOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdChildFolderIdOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdChildFolderIdOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeMailFolderIdChildFolderIdOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MailFolderId, ok = input.Parsed["mailFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", input)
	}

	if id.MailFolderId1, ok = input.Parsed["mailFolderId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId1", input)
	}

	if id.MailFolderOperationId, ok = input.Parsed["mailFolderOperationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderOperationId", input)
	}

	return nil
}

// ValidateMeMailFolderIdChildFolderIdOperationID checks that 'input' can be parsed as a Me Mail Folder Id Child Folder Id Operation ID
func ValidateMeMailFolderIdChildFolderIdOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderIdChildFolderIdOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Id Child Folder Id Operation ID
func (id MeMailFolderIdChildFolderIdOperationId) ID() string {
	fmtString := "/me/mailFolders/%s/childFolders/%s/operations/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MailFolderId1, id.MailFolderOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Id Child Folder Id Operation ID
func (id MeMailFolderIdChildFolderIdOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("childFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1"),
		resourceids.StaticSegment("operations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("mailFolderOperationId", "mailFolderOperationId"),
	}
}

// String returns a human-readable description of this Me Mail Folder Id Child Folder Id Operation ID
func (id MeMailFolderIdChildFolderIdOperationId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
		fmt.Sprintf("Mail Folder Operation: %q", id.MailFolderOperationId),
	}
	return fmt.Sprintf("Me Mail Folder Id Child Folder Id Operation (%s)", strings.Join(components, "\n"))
}
