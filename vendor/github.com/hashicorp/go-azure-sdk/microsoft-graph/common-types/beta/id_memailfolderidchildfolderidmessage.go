package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMailFolderIdChildFolderIdMessageId{}

// MeMailFolderIdChildFolderIdMessageId is a struct representing the Resource ID for a Me Mail Folder Id Child Folder Id Message
type MeMailFolderIdChildFolderIdMessageId struct {
	MailFolderId  string
	MailFolderId1 string
	MessageId     string
}

// NewMeMailFolderIdChildFolderIdMessageID returns a new MeMailFolderIdChildFolderIdMessageId struct
func NewMeMailFolderIdChildFolderIdMessageID(mailFolderId string, mailFolderId1 string, messageId string) MeMailFolderIdChildFolderIdMessageId {
	return MeMailFolderIdChildFolderIdMessageId{
		MailFolderId:  mailFolderId,
		MailFolderId1: mailFolderId1,
		MessageId:     messageId,
	}
}

// ParseMeMailFolderIdChildFolderIdMessageID parses 'input' into a MeMailFolderIdChildFolderIdMessageId
func ParseMeMailFolderIdChildFolderIdMessageID(input string) (*MeMailFolderIdChildFolderIdMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdChildFolderIdMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdChildFolderIdMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeMailFolderIdChildFolderIdMessageIDInsensitively parses 'input' case-insensitively into a MeMailFolderIdChildFolderIdMessageId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderIdChildFolderIdMessageIDInsensitively(input string) (*MeMailFolderIdChildFolderIdMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdChildFolderIdMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdChildFolderIdMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeMailFolderIdChildFolderIdMessageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MailFolderId, ok = input.Parsed["mailFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", input)
	}

	if id.MailFolderId1, ok = input.Parsed["mailFolderId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId1", input)
	}

	if id.MessageId, ok = input.Parsed["messageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "messageId", input)
	}

	return nil
}

// ValidateMeMailFolderIdChildFolderIdMessageID checks that 'input' can be parsed as a Me Mail Folder Id Child Folder Id Message ID
func ValidateMeMailFolderIdChildFolderIdMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderIdChildFolderIdMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Id Child Folder Id Message ID
func (id MeMailFolderIdChildFolderIdMessageId) ID() string {
	fmtString := "/me/mailFolders/%s/childFolders/%s/messages/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MailFolderId1, id.MessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Id Child Folder Id Message ID
func (id MeMailFolderIdChildFolderIdMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("childFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageId"),
	}
}

// String returns a human-readable description of this Me Mail Folder Id Child Folder Id Message ID
func (id MeMailFolderIdChildFolderIdMessageId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
		fmt.Sprintf("Message: %q", id.MessageId),
	}
	return fmt.Sprintf("Me Mail Folder Id Child Folder Id Message (%s)", strings.Join(components, "\n"))
}
