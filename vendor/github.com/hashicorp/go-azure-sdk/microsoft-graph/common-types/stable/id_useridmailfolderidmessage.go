package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMailFolderIdMessageId{}

// UserIdMailFolderIdMessageId is a struct representing the Resource ID for a User Id Mail Folder Id Message
type UserIdMailFolderIdMessageId struct {
	UserId       string
	MailFolderId string
	MessageId    string
}

// NewUserIdMailFolderIdMessageID returns a new UserIdMailFolderIdMessageId struct
func NewUserIdMailFolderIdMessageID(userId string, mailFolderId string, messageId string) UserIdMailFolderIdMessageId {
	return UserIdMailFolderIdMessageId{
		UserId:       userId,
		MailFolderId: mailFolderId,
		MessageId:    messageId,
	}
}

// ParseUserIdMailFolderIdMessageID parses 'input' into a UserIdMailFolderIdMessageId
func ParseUserIdMailFolderIdMessageID(input string) (*UserIdMailFolderIdMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdMailFolderIdMessageIDInsensitively parses 'input' case-insensitively into a UserIdMailFolderIdMessageId
// note: this method should only be used for API response data and not user input
func ParseUserIdMailFolderIdMessageIDInsensitively(input string) (*UserIdMailFolderIdMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdMailFolderIdMessageId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateUserIdMailFolderIdMessageID checks that 'input' can be parsed as a User Id Mail Folder Id Message ID
func ValidateUserIdMailFolderIdMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdMailFolderIdMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Mail Folder Id Message ID
func (id UserIdMailFolderIdMessageId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/messages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.MessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Mail Folder Id Message ID
func (id UserIdMailFolderIdMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageId"),
	}
}

// String returns a human-readable description of this User Id Mail Folder Id Message ID
func (id UserIdMailFolderIdMessageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Message: %q", id.MessageId),
	}
	return fmt.Sprintf("User Id Mail Folder Id Message (%s)", strings.Join(components, "\n"))
}
