package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMailFolderIdMessageIdAttachmentId{}

// UserIdMailFolderIdMessageIdAttachmentId is a struct representing the Resource ID for a User Id Mail Folder Id Message Id Attachment
type UserIdMailFolderIdMessageIdAttachmentId struct {
	UserId       string
	MailFolderId string
	MessageId    string
	AttachmentId string
}

// NewUserIdMailFolderIdMessageIdAttachmentID returns a new UserIdMailFolderIdMessageIdAttachmentId struct
func NewUserIdMailFolderIdMessageIdAttachmentID(userId string, mailFolderId string, messageId string, attachmentId string) UserIdMailFolderIdMessageIdAttachmentId {
	return UserIdMailFolderIdMessageIdAttachmentId{
		UserId:       userId,
		MailFolderId: mailFolderId,
		MessageId:    messageId,
		AttachmentId: attachmentId,
	}
}

// ParseUserIdMailFolderIdMessageIdAttachmentID parses 'input' into a UserIdMailFolderIdMessageIdAttachmentId
func ParseUserIdMailFolderIdMessageIdAttachmentID(input string) (*UserIdMailFolderIdMessageIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdMessageIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdMessageIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdMailFolderIdMessageIdAttachmentIDInsensitively parses 'input' case-insensitively into a UserIdMailFolderIdMessageIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserIdMailFolderIdMessageIdAttachmentIDInsensitively(input string) (*UserIdMailFolderIdMessageIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdMessageIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdMessageIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdMailFolderIdMessageIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

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

// ValidateUserIdMailFolderIdMessageIdAttachmentID checks that 'input' can be parsed as a User Id Mail Folder Id Message Id Attachment ID
func ValidateUserIdMailFolderIdMessageIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdMailFolderIdMessageIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Mail Folder Id Message Id Attachment ID
func (id UserIdMailFolderIdMessageIdAttachmentId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/messages/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.MessageId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Mail Folder Id Message Id Attachment ID
func (id UserIdMailFolderIdMessageIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageId"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentId"),
	}
}

// String returns a human-readable description of this User Id Mail Folder Id Message Id Attachment ID
func (id UserIdMailFolderIdMessageIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Id Mail Folder Id Message Id Attachment (%s)", strings.Join(components, "\n"))
}
