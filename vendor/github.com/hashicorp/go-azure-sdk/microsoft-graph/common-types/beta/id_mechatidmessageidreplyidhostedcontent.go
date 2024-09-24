package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeChatIdMessageIdReplyIdHostedContentId{}

// MeChatIdMessageIdReplyIdHostedContentId is a struct representing the Resource ID for a Me Chat Id Message Id Reply Id Hosted Content
type MeChatIdMessageIdReplyIdHostedContentId struct {
	ChatId                     string
	ChatMessageId              string
	ChatMessageId1             string
	ChatMessageHostedContentId string
}

// NewMeChatIdMessageIdReplyIdHostedContentID returns a new MeChatIdMessageIdReplyIdHostedContentId struct
func NewMeChatIdMessageIdReplyIdHostedContentID(chatId string, chatMessageId string, chatMessageId1 string, chatMessageHostedContentId string) MeChatIdMessageIdReplyIdHostedContentId {
	return MeChatIdMessageIdReplyIdHostedContentId{
		ChatId:                     chatId,
		ChatMessageId:              chatMessageId,
		ChatMessageId1:             chatMessageId1,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseMeChatIdMessageIdReplyIdHostedContentID parses 'input' into a MeChatIdMessageIdReplyIdHostedContentId
func ParseMeChatIdMessageIdReplyIdHostedContentID(input string) (*MeChatIdMessageIdReplyIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeChatIdMessageIdReplyIdHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeChatIdMessageIdReplyIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeChatIdMessageIdReplyIdHostedContentIDInsensitively parses 'input' case-insensitively into a MeChatIdMessageIdReplyIdHostedContentId
// note: this method should only be used for API response data and not user input
func ParseMeChatIdMessageIdReplyIdHostedContentIDInsensitively(input string) (*MeChatIdMessageIdReplyIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeChatIdMessageIdReplyIdHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeChatIdMessageIdReplyIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeChatIdMessageIdReplyIdHostedContentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ChatMessageHostedContentId, ok = input.Parsed["chatMessageHostedContentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", input)
	}

	return nil
}

// ValidateMeChatIdMessageIdReplyIdHostedContentID checks that 'input' can be parsed as a Me Chat Id Message Id Reply Id Hosted Content ID
func ValidateMeChatIdMessageIdReplyIdHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeChatIdMessageIdReplyIdHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Chat Id Message Id Reply Id Hosted Content ID
func (id MeChatIdMessageIdReplyIdHostedContentId) ID() string {
	fmtString := "/me/chats/%s/messages/%s/replies/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.ChatId, id.ChatMessageId, id.ChatMessageId1, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Chat Id Message Id Reply Id Hosted Content ID
func (id MeChatIdMessageIdReplyIdHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("chats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
		resourceids.StaticSegment("replies", "replies", "replies"),
		resourceids.UserSpecifiedSegment("chatMessageId1", "chatMessageId1"),
		resourceids.StaticSegment("hostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentId"),
	}
}

// String returns a human-readable description of this Me Chat Id Message Id Reply Id Hosted Content ID
func (id MeChatIdMessageIdReplyIdHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("Me Chat Id Message Id Reply Id Hosted Content (%s)", strings.Join(components, "\n"))
}
