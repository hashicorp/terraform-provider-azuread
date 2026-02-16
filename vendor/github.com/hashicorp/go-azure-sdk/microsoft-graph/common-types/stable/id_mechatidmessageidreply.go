package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeChatIdMessageIdReplyId{}

// MeChatIdMessageIdReplyId is a struct representing the Resource ID for a Me Chat Id Message Id Reply
type MeChatIdMessageIdReplyId struct {
	ChatId         string
	ChatMessageId  string
	ChatMessageId1 string
}

// NewMeChatIdMessageIdReplyID returns a new MeChatIdMessageIdReplyId struct
func NewMeChatIdMessageIdReplyID(chatId string, chatMessageId string, chatMessageId1 string) MeChatIdMessageIdReplyId {
	return MeChatIdMessageIdReplyId{
		ChatId:         chatId,
		ChatMessageId:  chatMessageId,
		ChatMessageId1: chatMessageId1,
	}
}

// ParseMeChatIdMessageIdReplyID parses 'input' into a MeChatIdMessageIdReplyId
func ParseMeChatIdMessageIdReplyID(input string) (*MeChatIdMessageIdReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeChatIdMessageIdReplyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeChatIdMessageIdReplyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeChatIdMessageIdReplyIDInsensitively parses 'input' case-insensitively into a MeChatIdMessageIdReplyId
// note: this method should only be used for API response data and not user input
func ParseMeChatIdMessageIdReplyIDInsensitively(input string) (*MeChatIdMessageIdReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeChatIdMessageIdReplyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeChatIdMessageIdReplyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeChatIdMessageIdReplyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

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

// ValidateMeChatIdMessageIdReplyID checks that 'input' can be parsed as a Me Chat Id Message Id Reply ID
func ValidateMeChatIdMessageIdReplyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeChatIdMessageIdReplyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Chat Id Message Id Reply ID
func (id MeChatIdMessageIdReplyId) ID() string {
	fmtString := "/me/chats/%s/messages/%s/replies/%s"
	return fmt.Sprintf(fmtString, id.ChatId, id.ChatMessageId, id.ChatMessageId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Chat Id Message Id Reply ID
func (id MeChatIdMessageIdReplyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("chats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
		resourceids.StaticSegment("replies", "replies", "replies"),
		resourceids.UserSpecifiedSegment("chatMessageId1", "chatMessageId1"),
	}
}

// String returns a human-readable description of this Me Chat Id Message Id Reply ID
func (id MeChatIdMessageIdReplyId) String() string {
	components := []string{
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
	}
	return fmt.Sprintf("Me Chat Id Message Id Reply (%s)", strings.Join(components, "\n"))
}
