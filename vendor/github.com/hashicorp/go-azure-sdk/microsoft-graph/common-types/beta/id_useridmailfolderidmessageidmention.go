package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMailFolderIdMessageIdMentionId{}

// UserIdMailFolderIdMessageIdMentionId is a struct representing the Resource ID for a User Id Mail Folder Id Message Id Mention
type UserIdMailFolderIdMessageIdMentionId struct {
	UserId       string
	MailFolderId string
	MessageId    string
	MentionId    string
}

// NewUserIdMailFolderIdMessageIdMentionID returns a new UserIdMailFolderIdMessageIdMentionId struct
func NewUserIdMailFolderIdMessageIdMentionID(userId string, mailFolderId string, messageId string, mentionId string) UserIdMailFolderIdMessageIdMentionId {
	return UserIdMailFolderIdMessageIdMentionId{
		UserId:       userId,
		MailFolderId: mailFolderId,
		MessageId:    messageId,
		MentionId:    mentionId,
	}
}

// ParseUserIdMailFolderIdMessageIdMentionID parses 'input' into a UserIdMailFolderIdMessageIdMentionId
func ParseUserIdMailFolderIdMessageIdMentionID(input string) (*UserIdMailFolderIdMessageIdMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdMessageIdMentionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdMessageIdMentionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdMailFolderIdMessageIdMentionIDInsensitively parses 'input' case-insensitively into a UserIdMailFolderIdMessageIdMentionId
// note: this method should only be used for API response data and not user input
func ParseUserIdMailFolderIdMessageIdMentionIDInsensitively(input string) (*UserIdMailFolderIdMessageIdMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdMessageIdMentionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdMessageIdMentionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdMailFolderIdMessageIdMentionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.MentionId, ok = input.Parsed["mentionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mentionId", input)
	}

	return nil
}

// ValidateUserIdMailFolderIdMessageIdMentionID checks that 'input' can be parsed as a User Id Mail Folder Id Message Id Mention ID
func ValidateUserIdMailFolderIdMessageIdMentionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdMailFolderIdMessageIdMentionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Mail Folder Id Message Id Mention ID
func (id UserIdMailFolderIdMessageIdMentionId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/messages/%s/mentions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.MessageId, id.MentionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Mail Folder Id Message Id Mention ID
func (id UserIdMailFolderIdMessageIdMentionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageId"),
		resourceids.StaticSegment("mentions", "mentions", "mentions"),
		resourceids.UserSpecifiedSegment("mentionId", "mentionId"),
	}
}

// String returns a human-readable description of this User Id Mail Folder Id Message Id Mention ID
func (id UserIdMailFolderIdMessageIdMentionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Mention: %q", id.MentionId),
	}
	return fmt.Sprintf("User Id Mail Folder Id Message Id Mention (%s)", strings.Join(components, "\n"))
}
