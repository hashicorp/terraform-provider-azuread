package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeChatIdPinnedMessageId{}

// MeChatIdPinnedMessageId is a struct representing the Resource ID for a Me Chat Id Pinned Message
type MeChatIdPinnedMessageId struct {
	ChatId                  string
	PinnedChatMessageInfoId string
}

// NewMeChatIdPinnedMessageID returns a new MeChatIdPinnedMessageId struct
func NewMeChatIdPinnedMessageID(chatId string, pinnedChatMessageInfoId string) MeChatIdPinnedMessageId {
	return MeChatIdPinnedMessageId{
		ChatId:                  chatId,
		PinnedChatMessageInfoId: pinnedChatMessageInfoId,
	}
}

// ParseMeChatIdPinnedMessageID parses 'input' into a MeChatIdPinnedMessageId
func ParseMeChatIdPinnedMessageID(input string) (*MeChatIdPinnedMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeChatIdPinnedMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeChatIdPinnedMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeChatIdPinnedMessageIDInsensitively parses 'input' case-insensitively into a MeChatIdPinnedMessageId
// note: this method should only be used for API response data and not user input
func ParseMeChatIdPinnedMessageIDInsensitively(input string) (*MeChatIdPinnedMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeChatIdPinnedMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeChatIdPinnedMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeChatIdPinnedMessageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ChatId, ok = input.Parsed["chatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatId", input)
	}

	if id.PinnedChatMessageInfoId, ok = input.Parsed["pinnedChatMessageInfoId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "pinnedChatMessageInfoId", input)
	}

	return nil
}

// ValidateMeChatIdPinnedMessageID checks that 'input' can be parsed as a Me Chat Id Pinned Message ID
func ValidateMeChatIdPinnedMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeChatIdPinnedMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Chat Id Pinned Message ID
func (id MeChatIdPinnedMessageId) ID() string {
	fmtString := "/me/chats/%s/pinnedMessages/%s"
	return fmt.Sprintf(fmtString, id.ChatId, id.PinnedChatMessageInfoId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Chat Id Pinned Message ID
func (id MeChatIdPinnedMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("chats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatId"),
		resourceids.StaticSegment("pinnedMessages", "pinnedMessages", "pinnedMessages"),
		resourceids.UserSpecifiedSegment("pinnedChatMessageInfoId", "pinnedChatMessageInfoId"),
	}
}

// String returns a human-readable description of this Me Chat Id Pinned Message ID
func (id MeChatIdPinnedMessageId) String() string {
	components := []string{
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Pinned Chat Message Info: %q", id.PinnedChatMessageInfoId),
	}
	return fmt.Sprintf("Me Chat Id Pinned Message (%s)", strings.Join(components, "\n"))
}
