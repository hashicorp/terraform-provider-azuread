package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeChatIdMessageIdHostedContentId{}

// MeChatIdMessageIdHostedContentId is a struct representing the Resource ID for a Me Chat Id Message Id Hosted Content
type MeChatIdMessageIdHostedContentId struct {
	ChatId                     string
	ChatMessageId              string
	ChatMessageHostedContentId string
}

// NewMeChatIdMessageIdHostedContentID returns a new MeChatIdMessageIdHostedContentId struct
func NewMeChatIdMessageIdHostedContentID(chatId string, chatMessageId string, chatMessageHostedContentId string) MeChatIdMessageIdHostedContentId {
	return MeChatIdMessageIdHostedContentId{
		ChatId:                     chatId,
		ChatMessageId:              chatMessageId,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseMeChatIdMessageIdHostedContentID parses 'input' into a MeChatIdMessageIdHostedContentId
func ParseMeChatIdMessageIdHostedContentID(input string) (*MeChatIdMessageIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeChatIdMessageIdHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeChatIdMessageIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeChatIdMessageIdHostedContentIDInsensitively parses 'input' case-insensitively into a MeChatIdMessageIdHostedContentId
// note: this method should only be used for API response data and not user input
func ParseMeChatIdMessageIdHostedContentIDInsensitively(input string) (*MeChatIdMessageIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeChatIdMessageIdHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeChatIdMessageIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeChatIdMessageIdHostedContentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ChatId, ok = input.Parsed["chatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatId", input)
	}

	if id.ChatMessageId, ok = input.Parsed["chatMessageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", input)
	}

	if id.ChatMessageHostedContentId, ok = input.Parsed["chatMessageHostedContentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", input)
	}

	return nil
}

// ValidateMeChatIdMessageIdHostedContentID checks that 'input' can be parsed as a Me Chat Id Message Id Hosted Content ID
func ValidateMeChatIdMessageIdHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeChatIdMessageIdHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Chat Id Message Id Hosted Content ID
func (id MeChatIdMessageIdHostedContentId) ID() string {
	fmtString := "/me/chats/%s/messages/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.ChatId, id.ChatMessageId, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Chat Id Message Id Hosted Content ID
func (id MeChatIdMessageIdHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("chats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
		resourceids.StaticSegment("hostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentId"),
	}
}

// String returns a human-readable description of this Me Chat Id Message Id Hosted Content ID
func (id MeChatIdMessageIdHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("Me Chat Id Message Id Hosted Content (%s)", strings.Join(components, "\n"))
}
