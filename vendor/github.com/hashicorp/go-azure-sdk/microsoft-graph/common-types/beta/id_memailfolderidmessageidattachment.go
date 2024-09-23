package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMailFolderIdMessageIdAttachmentId{}

// MeMailFolderIdMessageIdAttachmentId is a struct representing the Resource ID for a Me Mail Folder Id Message Id Attachment
type MeMailFolderIdMessageIdAttachmentId struct {
	MailFolderId string
	MessageId    string
	AttachmentId string
}

// NewMeMailFolderIdMessageIdAttachmentID returns a new MeMailFolderIdMessageIdAttachmentId struct
func NewMeMailFolderIdMessageIdAttachmentID(mailFolderId string, messageId string, attachmentId string) MeMailFolderIdMessageIdAttachmentId {
	return MeMailFolderIdMessageIdAttachmentId{
		MailFolderId: mailFolderId,
		MessageId:    messageId,
		AttachmentId: attachmentId,
	}
}

// ParseMeMailFolderIdMessageIdAttachmentID parses 'input' into a MeMailFolderIdMessageIdAttachmentId
func ParseMeMailFolderIdMessageIdAttachmentID(input string) (*MeMailFolderIdMessageIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdMessageIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdMessageIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeMailFolderIdMessageIdAttachmentIDInsensitively parses 'input' case-insensitively into a MeMailFolderIdMessageIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderIdMessageIdAttachmentIDInsensitively(input string) (*MeMailFolderIdMessageIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdMessageIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdMessageIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeMailFolderIdMessageIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MailFolderId, ok = input.Parsed["mailFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", input)
	}

	if id.MessageId, ok = input.Parsed["messageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "messageId", input)
	}

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateMeMailFolderIdMessageIdAttachmentID checks that 'input' can be parsed as a Me Mail Folder Id Message Id Attachment ID
func ValidateMeMailFolderIdMessageIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderIdMessageIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Id Message Id Attachment ID
func (id MeMailFolderIdMessageIdAttachmentId) ID() string {
	fmtString := "/me/mailFolders/%s/messages/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MessageId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Id Message Id Attachment ID
func (id MeMailFolderIdMessageIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageId"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentId"),
	}
}

// String returns a human-readable description of this Me Mail Folder Id Message Id Attachment ID
func (id MeMailFolderIdMessageIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Mail Folder Id Message Id Attachment (%s)", strings.Join(components, "\n"))
}
