package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMailFolderIdChildFolderIdMessageIdAttachmentId{}

// UserIdMailFolderIdChildFolderIdMessageIdAttachmentId is a struct representing the Resource ID for a User Id Mail Folder Id Child Folder Id Message Id Attachment
type UserIdMailFolderIdChildFolderIdMessageIdAttachmentId struct {
	UserId        string
	MailFolderId  string
	MailFolderId1 string
	MessageId     string
	AttachmentId  string
}

// NewUserIdMailFolderIdChildFolderIdMessageIdAttachmentID returns a new UserIdMailFolderIdChildFolderIdMessageIdAttachmentId struct
func NewUserIdMailFolderIdChildFolderIdMessageIdAttachmentID(userId string, mailFolderId string, mailFolderId1 string, messageId string, attachmentId string) UserIdMailFolderIdChildFolderIdMessageIdAttachmentId {
	return UserIdMailFolderIdChildFolderIdMessageIdAttachmentId{
		UserId:        userId,
		MailFolderId:  mailFolderId,
		MailFolderId1: mailFolderId1,
		MessageId:     messageId,
		AttachmentId:  attachmentId,
	}
}

// ParseUserIdMailFolderIdChildFolderIdMessageIdAttachmentID parses 'input' into a UserIdMailFolderIdChildFolderIdMessageIdAttachmentId
func ParseUserIdMailFolderIdChildFolderIdMessageIdAttachmentID(input string) (*UserIdMailFolderIdChildFolderIdMessageIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdChildFolderIdMessageIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdChildFolderIdMessageIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdMailFolderIdChildFolderIdMessageIdAttachmentIDInsensitively parses 'input' case-insensitively into a UserIdMailFolderIdChildFolderIdMessageIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserIdMailFolderIdChildFolderIdMessageIdAttachmentIDInsensitively(input string) (*UserIdMailFolderIdChildFolderIdMessageIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdChildFolderIdMessageIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdChildFolderIdMessageIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdMailFolderIdChildFolderIdMessageIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

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

// ValidateUserIdMailFolderIdChildFolderIdMessageIdAttachmentID checks that 'input' can be parsed as a User Id Mail Folder Id Child Folder Id Message Id Attachment ID
func ValidateUserIdMailFolderIdChildFolderIdMessageIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdMailFolderIdChildFolderIdMessageIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Mail Folder Id Child Folder Id Message Id Attachment ID
func (id UserIdMailFolderIdChildFolderIdMessageIdAttachmentId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/childFolders/%s/messages/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.MailFolderId1, id.MessageId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Mail Folder Id Child Folder Id Message Id Attachment ID
func (id UserIdMailFolderIdChildFolderIdMessageIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
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

// String returns a human-readable description of this User Id Mail Folder Id Child Folder Id Message Id Attachment ID
func (id UserIdMailFolderIdChildFolderIdMessageIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Id Mail Folder Id Child Folder Id Message Id Attachment (%s)", strings.Join(components, "\n"))
}
