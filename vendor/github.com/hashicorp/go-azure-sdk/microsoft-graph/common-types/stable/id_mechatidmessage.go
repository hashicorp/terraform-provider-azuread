package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeChatIdMessageId{}

// MeChatIdMessageId is a struct representing the Resource ID for a Me Chat Id Message
type MeChatIdMessageId struct {
	ChatId        string
	ChatMessageId string
}

// NewMeChatIdMessageID returns a new MeChatIdMessageId struct
func NewMeChatIdMessageID(chatId string, chatMessageId string) MeChatIdMessageId {
	return MeChatIdMessageId{
		ChatId:        chatId,
		ChatMessageId: chatMessageId,
	}
}

// ParseMeChatIdMessageID parses 'input' into a MeChatIdMessageId
func ParseMeChatIdMessageID(input string) (*MeChatIdMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeChatIdMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeChatIdMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeChatIdMessageIDInsensitively parses 'input' case-insensitively into a MeChatIdMessageId
// note: this method should only be used for API response data and not user input
func ParseMeChatIdMessageIDInsensitively(input string) (*MeChatIdMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeChatIdMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeChatIdMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeChatIdMessageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ChatId, ok = input.Parsed["chatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatId", input)
	}

	if id.ChatMessageId, ok = input.Parsed["chatMessageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", input)
	}

	return nil
}

// ValidateMeChatIdMessageID checks that 'input' can be parsed as a Me Chat Id Message ID
func ValidateMeChatIdMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeChatIdMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Chat Id Message ID
func (id MeChatIdMessageId) ID() string {
	fmtString := "/me/chats/%s/messages/%s"
	return fmt.Sprintf(fmtString, id.ChatId, id.ChatMessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Chat Id Message ID
func (id MeChatIdMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("chats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
	}
}

// String returns a human-readable description of this Me Chat Id Message ID
func (id MeChatIdMessageId) String() string {
	components := []string{
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
	}
	return fmt.Sprintf("Me Chat Id Message (%s)", strings.Join(components, "\n"))
}
