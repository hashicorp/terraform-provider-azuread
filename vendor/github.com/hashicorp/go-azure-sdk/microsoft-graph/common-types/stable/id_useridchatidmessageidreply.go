package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdChatIdMessageIdReplyId{}

// UserIdChatIdMessageIdReplyId is a struct representing the Resource ID for a User Id Chat Id Message Id Reply
type UserIdChatIdMessageIdReplyId struct {
	UserId         string
	ChatId         string
	ChatMessageId  string
	ChatMessageId1 string
}

// NewUserIdChatIdMessageIdReplyID returns a new UserIdChatIdMessageIdReplyId struct
func NewUserIdChatIdMessageIdReplyID(userId string, chatId string, chatMessageId string, chatMessageId1 string) UserIdChatIdMessageIdReplyId {
	return UserIdChatIdMessageIdReplyId{
		UserId:         userId,
		ChatId:         chatId,
		ChatMessageId:  chatMessageId,
		ChatMessageId1: chatMessageId1,
	}
}

// ParseUserIdChatIdMessageIdReplyID parses 'input' into a UserIdChatIdMessageIdReplyId
func ParseUserIdChatIdMessageIdReplyID(input string) (*UserIdChatIdMessageIdReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdChatIdMessageIdReplyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdChatIdMessageIdReplyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdChatIdMessageIdReplyIDInsensitively parses 'input' case-insensitively into a UserIdChatIdMessageIdReplyId
// note: this method should only be used for API response data and not user input
func ParseUserIdChatIdMessageIdReplyIDInsensitively(input string) (*UserIdChatIdMessageIdReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdChatIdMessageIdReplyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdChatIdMessageIdReplyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdChatIdMessageIdReplyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ChatId, ok = input.Parsed["chatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatId", input)
	}

	if id.ChatMessageId, ok = input.Parsed["chatMessageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", input)
	}

	if id.ChatMessageId1, ok = input.Parsed["chatMessageId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId1", input)
	}

	return nil
}

// ValidateUserIdChatIdMessageIdReplyID checks that 'input' can be parsed as a User Id Chat Id Message Id Reply ID
func ValidateUserIdChatIdMessageIdReplyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdChatIdMessageIdReplyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Chat Id Message Id Reply ID
func (id UserIdChatIdMessageIdReplyId) ID() string {
	fmtString := "/users/%s/chats/%s/messages/%s/replies/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ChatId, id.ChatMessageId, id.ChatMessageId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Chat Id Message Id Reply ID
func (id UserIdChatIdMessageIdReplyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("chats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
		resourceids.StaticSegment("replies", "replies", "replies"),
		resourceids.UserSpecifiedSegment("chatMessageId1", "chatMessageId1"),
	}
}

// String returns a human-readable description of this User Id Chat Id Message Id Reply ID
func (id UserIdChatIdMessageIdReplyId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
	}
	return fmt.Sprintf("User Id Chat Id Message Id Reply (%s)", strings.Join(components, "\n"))
}
