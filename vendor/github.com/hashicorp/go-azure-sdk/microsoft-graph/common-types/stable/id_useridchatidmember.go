package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdChatIdMemberId{}

// UserIdChatIdMemberId is a struct representing the Resource ID for a User Id Chat Id Member
type UserIdChatIdMemberId struct {
	UserId               string
	ChatId               string
	ConversationMemberId string
}

// NewUserIdChatIdMemberID returns a new UserIdChatIdMemberId struct
func NewUserIdChatIdMemberID(userId string, chatId string, conversationMemberId string) UserIdChatIdMemberId {
	return UserIdChatIdMemberId{
		UserId:               userId,
		ChatId:               chatId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseUserIdChatIdMemberID parses 'input' into a UserIdChatIdMemberId
func ParseUserIdChatIdMemberID(input string) (*UserIdChatIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdChatIdMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdChatIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdChatIdMemberIDInsensitively parses 'input' case-insensitively into a UserIdChatIdMemberId
// note: this method should only be used for API response data and not user input
func ParseUserIdChatIdMemberIDInsensitively(input string) (*UserIdChatIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdChatIdMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdChatIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdChatIdMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ChatId, ok = input.Parsed["chatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatId", input)
	}

	if id.ConversationMemberId, ok = input.Parsed["conversationMemberId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", input)
	}

	return nil
}

// ValidateUserIdChatIdMemberID checks that 'input' can be parsed as a User Id Chat Id Member ID
func ValidateUserIdChatIdMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdChatIdMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Chat Id Member ID
func (id UserIdChatIdMemberId) ID() string {
	fmtString := "/users/%s/chats/%s/members/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ChatId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Chat Id Member ID
func (id UserIdChatIdMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("chats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatId"),
		resourceids.StaticSegment("members", "members", "members"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberId"),
	}
}

// String returns a human-readable description of this User Id Chat Id Member ID
func (id UserIdChatIdMemberId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("User Id Chat Id Member (%s)", strings.Join(components, "\n"))
}
