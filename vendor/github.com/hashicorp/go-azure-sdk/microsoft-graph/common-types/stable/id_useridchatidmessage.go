package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdChatIdMessageId{}

// UserIdChatIdMessageId is a struct representing the Resource ID for a User Id Chat Id Message
type UserIdChatIdMessageId struct {
	UserId        string
	ChatId        string
	ChatMessageId string
}

// NewUserIdChatIdMessageID returns a new UserIdChatIdMessageId struct
func NewUserIdChatIdMessageID(userId string, chatId string, chatMessageId string) UserIdChatIdMessageId {
	return UserIdChatIdMessageId{
		UserId:        userId,
		ChatId:        chatId,
		ChatMessageId: chatMessageId,
	}
}

// ParseUserIdChatIdMessageID parses 'input' into a UserIdChatIdMessageId
func ParseUserIdChatIdMessageID(input string) (*UserIdChatIdMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdChatIdMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdChatIdMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdChatIdMessageIDInsensitively parses 'input' case-insensitively into a UserIdChatIdMessageId
// note: this method should only be used for API response data and not user input
func ParseUserIdChatIdMessageIDInsensitively(input string) (*UserIdChatIdMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdChatIdMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdChatIdMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdChatIdMessageId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateUserIdChatIdMessageID checks that 'input' can be parsed as a User Id Chat Id Message ID
func ValidateUserIdChatIdMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdChatIdMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Chat Id Message ID
func (id UserIdChatIdMessageId) ID() string {
	fmtString := "/users/%s/chats/%s/messages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ChatId, id.ChatMessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Chat Id Message ID
func (id UserIdChatIdMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("chats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
	}
}

// String returns a human-readable description of this User Id Chat Id Message ID
func (id UserIdChatIdMessageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
	}
	return fmt.Sprintf("User Id Chat Id Message (%s)", strings.Join(components, "\n"))
}
