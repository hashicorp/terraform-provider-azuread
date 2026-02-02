package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdChatIdPinnedMessageId{}

// UserIdChatIdPinnedMessageId is a struct representing the Resource ID for a User Id Chat Id Pinned Message
type UserIdChatIdPinnedMessageId struct {
	UserId                  string
	ChatId                  string
	PinnedChatMessageInfoId string
}

// NewUserIdChatIdPinnedMessageID returns a new UserIdChatIdPinnedMessageId struct
func NewUserIdChatIdPinnedMessageID(userId string, chatId string, pinnedChatMessageInfoId string) UserIdChatIdPinnedMessageId {
	return UserIdChatIdPinnedMessageId{
		UserId:                  userId,
		ChatId:                  chatId,
		PinnedChatMessageInfoId: pinnedChatMessageInfoId,
	}
}

// ParseUserIdChatIdPinnedMessageID parses 'input' into a UserIdChatIdPinnedMessageId
func ParseUserIdChatIdPinnedMessageID(input string) (*UserIdChatIdPinnedMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdChatIdPinnedMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdChatIdPinnedMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdChatIdPinnedMessageIDInsensitively parses 'input' case-insensitively into a UserIdChatIdPinnedMessageId
// note: this method should only be used for API response data and not user input
func ParseUserIdChatIdPinnedMessageIDInsensitively(input string) (*UserIdChatIdPinnedMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdChatIdPinnedMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdChatIdPinnedMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdChatIdPinnedMessageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ChatId, ok = input.Parsed["chatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatId", input)
	}

	if id.PinnedChatMessageInfoId, ok = input.Parsed["pinnedChatMessageInfoId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "pinnedChatMessageInfoId", input)
	}

	return nil
}

// ValidateUserIdChatIdPinnedMessageID checks that 'input' can be parsed as a User Id Chat Id Pinned Message ID
func ValidateUserIdChatIdPinnedMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdChatIdPinnedMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Chat Id Pinned Message ID
func (id UserIdChatIdPinnedMessageId) ID() string {
	fmtString := "/users/%s/chats/%s/pinnedMessages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ChatId, id.PinnedChatMessageInfoId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Chat Id Pinned Message ID
func (id UserIdChatIdPinnedMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("chats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatId"),
		resourceids.StaticSegment("pinnedMessages", "pinnedMessages", "pinnedMessages"),
		resourceids.UserSpecifiedSegment("pinnedChatMessageInfoId", "pinnedChatMessageInfoId"),
	}
}

// String returns a human-readable description of this User Id Chat Id Pinned Message ID
func (id UserIdChatIdPinnedMessageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Pinned Chat Message Info: %q", id.PinnedChatMessageInfoId),
	}
	return fmt.Sprintf("User Id Chat Id Pinned Message (%s)", strings.Join(components, "\n"))
}
