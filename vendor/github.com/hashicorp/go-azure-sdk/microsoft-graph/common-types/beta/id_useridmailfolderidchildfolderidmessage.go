package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMailFolderIdChildFolderIdMessageId{}

// UserIdMailFolderIdChildFolderIdMessageId is a struct representing the Resource ID for a User Id Mail Folder Id Child Folder Id Message
type UserIdMailFolderIdChildFolderIdMessageId struct {
	UserId        string
	MailFolderId  string
	MailFolderId1 string
	MessageId     string
}

// NewUserIdMailFolderIdChildFolderIdMessageID returns a new UserIdMailFolderIdChildFolderIdMessageId struct
func NewUserIdMailFolderIdChildFolderIdMessageID(userId string, mailFolderId string, mailFolderId1 string, messageId string) UserIdMailFolderIdChildFolderIdMessageId {
	return UserIdMailFolderIdChildFolderIdMessageId{
		UserId:        userId,
		MailFolderId:  mailFolderId,
		MailFolderId1: mailFolderId1,
		MessageId:     messageId,
	}
}

// ParseUserIdMailFolderIdChildFolderIdMessageID parses 'input' into a UserIdMailFolderIdChildFolderIdMessageId
func ParseUserIdMailFolderIdChildFolderIdMessageID(input string) (*UserIdMailFolderIdChildFolderIdMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdChildFolderIdMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdChildFolderIdMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdMailFolderIdChildFolderIdMessageIDInsensitively parses 'input' case-insensitively into a UserIdMailFolderIdChildFolderIdMessageId
// note: this method should only be used for API response data and not user input
func ParseUserIdMailFolderIdChildFolderIdMessageIDInsensitively(input string) (*UserIdMailFolderIdChildFolderIdMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdChildFolderIdMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdChildFolderIdMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdMailFolderIdChildFolderIdMessageId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateUserIdMailFolderIdChildFolderIdMessageID checks that 'input' can be parsed as a User Id Mail Folder Id Child Folder Id Message ID
func ValidateUserIdMailFolderIdChildFolderIdMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdMailFolderIdChildFolderIdMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Mail Folder Id Child Folder Id Message ID
func (id UserIdMailFolderIdChildFolderIdMessageId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/childFolders/%s/messages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.MailFolderId1, id.MessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Mail Folder Id Child Folder Id Message ID
func (id UserIdMailFolderIdChildFolderIdMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("childFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageId"),
	}
}

// String returns a human-readable description of this User Id Mail Folder Id Child Folder Id Message ID
func (id UserIdMailFolderIdChildFolderIdMessageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
		fmt.Sprintf("Message: %q", id.MessageId),
	}
	return fmt.Sprintf("User Id Mail Folder Id Child Folder Id Message (%s)", strings.Join(components, "\n"))
}
