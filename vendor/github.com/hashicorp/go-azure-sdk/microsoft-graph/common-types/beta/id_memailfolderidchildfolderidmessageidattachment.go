package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMailFolderIdChildFolderIdMessageIdAttachmentId{}

// MeMailFolderIdChildFolderIdMessageIdAttachmentId is a struct representing the Resource ID for a Me Mail Folder Id Child Folder Id Message Id Attachment
type MeMailFolderIdChildFolderIdMessageIdAttachmentId struct {
	MailFolderId  string
	MailFolderId1 string
	MessageId     string
	AttachmentId  string
}

// NewMeMailFolderIdChildFolderIdMessageIdAttachmentID returns a new MeMailFolderIdChildFolderIdMessageIdAttachmentId struct
func NewMeMailFolderIdChildFolderIdMessageIdAttachmentID(mailFolderId string, mailFolderId1 string, messageId string, attachmentId string) MeMailFolderIdChildFolderIdMessageIdAttachmentId {
	return MeMailFolderIdChildFolderIdMessageIdAttachmentId{
		MailFolderId:  mailFolderId,
		MailFolderId1: mailFolderId1,
		MessageId:     messageId,
		AttachmentId:  attachmentId,
	}
}

// ParseMeMailFolderIdChildFolderIdMessageIdAttachmentID parses 'input' into a MeMailFolderIdChildFolderIdMessageIdAttachmentId
func ParseMeMailFolderIdChildFolderIdMessageIdAttachmentID(input string) (*MeMailFolderIdChildFolderIdMessageIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdChildFolderIdMessageIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdChildFolderIdMessageIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeMailFolderIdChildFolderIdMessageIdAttachmentIDInsensitively parses 'input' case-insensitively into a MeMailFolderIdChildFolderIdMessageIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderIdChildFolderIdMessageIdAttachmentIDInsensitively(input string) (*MeMailFolderIdChildFolderIdMessageIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdChildFolderIdMessageIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdChildFolderIdMessageIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeMailFolderIdChildFolderIdMessageIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateMeMailFolderIdChildFolderIdMessageIdAttachmentID checks that 'input' can be parsed as a Me Mail Folder Id Child Folder Id Message Id Attachment ID
func ValidateMeMailFolderIdChildFolderIdMessageIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderIdChildFolderIdMessageIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Id Child Folder Id Message Id Attachment ID
func (id MeMailFolderIdChildFolderIdMessageIdAttachmentId) ID() string {
	fmtString := "/me/mailFolders/%s/childFolders/%s/messages/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MailFolderId1, id.MessageId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Id Child Folder Id Message Id Attachment ID
func (id MeMailFolderIdChildFolderIdMessageIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("childFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageId"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentId"),
	}
}

// String returns a human-readable description of this Me Mail Folder Id Child Folder Id Message Id Attachment ID
func (id MeMailFolderIdChildFolderIdMessageIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Mail Folder Id Child Folder Id Message Id Attachment (%s)", strings.Join(components, "\n"))
}
