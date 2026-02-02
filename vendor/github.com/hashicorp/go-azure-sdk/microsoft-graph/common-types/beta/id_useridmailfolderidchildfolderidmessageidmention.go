package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMailFolderIdChildFolderIdMessageIdMentionId{}

// UserIdMailFolderIdChildFolderIdMessageIdMentionId is a struct representing the Resource ID for a User Id Mail Folder Id Child Folder Id Message Id Mention
type UserIdMailFolderIdChildFolderIdMessageIdMentionId struct {
	UserId        string
	MailFolderId  string
	MailFolderId1 string
	MessageId     string
	MentionId     string
}

// NewUserIdMailFolderIdChildFolderIdMessageIdMentionID returns a new UserIdMailFolderIdChildFolderIdMessageIdMentionId struct
func NewUserIdMailFolderIdChildFolderIdMessageIdMentionID(userId string, mailFolderId string, mailFolderId1 string, messageId string, mentionId string) UserIdMailFolderIdChildFolderIdMessageIdMentionId {
	return UserIdMailFolderIdChildFolderIdMessageIdMentionId{
		UserId:        userId,
		MailFolderId:  mailFolderId,
		MailFolderId1: mailFolderId1,
		MessageId:     messageId,
		MentionId:     mentionId,
	}
}

// ParseUserIdMailFolderIdChildFolderIdMessageIdMentionID parses 'input' into a UserIdMailFolderIdChildFolderIdMessageIdMentionId
func ParseUserIdMailFolderIdChildFolderIdMessageIdMentionID(input string) (*UserIdMailFolderIdChildFolderIdMessageIdMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdChildFolderIdMessageIdMentionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdChildFolderIdMessageIdMentionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdMailFolderIdChildFolderIdMessageIdMentionIDInsensitively parses 'input' case-insensitively into a UserIdMailFolderIdChildFolderIdMessageIdMentionId
// note: this method should only be used for API response data and not user input
func ParseUserIdMailFolderIdChildFolderIdMessageIdMentionIDInsensitively(input string) (*UserIdMailFolderIdChildFolderIdMessageIdMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdChildFolderIdMessageIdMentionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdChildFolderIdMessageIdMentionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdMailFolderIdChildFolderIdMessageIdMentionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.MentionId, ok = input.Parsed["mentionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mentionId", input)
	}

	return nil
}

// ValidateUserIdMailFolderIdChildFolderIdMessageIdMentionID checks that 'input' can be parsed as a User Id Mail Folder Id Child Folder Id Message Id Mention ID
func ValidateUserIdMailFolderIdChildFolderIdMessageIdMentionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdMailFolderIdChildFolderIdMessageIdMentionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Mail Folder Id Child Folder Id Message Id Mention ID
func (id UserIdMailFolderIdChildFolderIdMessageIdMentionId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/childFolders/%s/messages/%s/mentions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.MailFolderId1, id.MessageId, id.MentionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Mail Folder Id Child Folder Id Message Id Mention ID
func (id UserIdMailFolderIdChildFolderIdMessageIdMentionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("childFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageId"),
		resourceids.StaticSegment("mentions", "mentions", "mentions"),
		resourceids.UserSpecifiedSegment("mentionId", "mentionId"),
	}
}

// String returns a human-readable description of this User Id Mail Folder Id Child Folder Id Message Id Mention ID
func (id UserIdMailFolderIdChildFolderIdMessageIdMentionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Mention: %q", id.MentionId),
	}
	return fmt.Sprintf("User Id Mail Folder Id Child Folder Id Message Id Mention (%s)", strings.Join(components, "\n"))
}
