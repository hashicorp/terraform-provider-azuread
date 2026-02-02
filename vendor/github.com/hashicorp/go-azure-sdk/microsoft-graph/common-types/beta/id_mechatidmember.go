package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeChatIdMemberId{}

// MeChatIdMemberId is a struct representing the Resource ID for a Me Chat Id Member
type MeChatIdMemberId struct {
	ChatId               string
	ConversationMemberId string
}

// NewMeChatIdMemberID returns a new MeChatIdMemberId struct
func NewMeChatIdMemberID(chatId string, conversationMemberId string) MeChatIdMemberId {
	return MeChatIdMemberId{
		ChatId:               chatId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseMeChatIdMemberID parses 'input' into a MeChatIdMemberId
func ParseMeChatIdMemberID(input string) (*MeChatIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeChatIdMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeChatIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeChatIdMemberIDInsensitively parses 'input' case-insensitively into a MeChatIdMemberId
// note: this method should only be used for API response data and not user input
func ParseMeChatIdMemberIDInsensitively(input string) (*MeChatIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeChatIdMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeChatIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeChatIdMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ChatId, ok = input.Parsed["chatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatId", input)
	}

	if id.ConversationMemberId, ok = input.Parsed["conversationMemberId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", input)
	}

	return nil
}

// ValidateMeChatIdMemberID checks that 'input' can be parsed as a Me Chat Id Member ID
func ValidateMeChatIdMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeChatIdMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Chat Id Member ID
func (id MeChatIdMemberId) ID() string {
	fmtString := "/me/chats/%s/members/%s"
	return fmt.Sprintf(fmtString, id.ChatId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Chat Id Member ID
func (id MeChatIdMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("chats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatId"),
		resourceids.StaticSegment("members", "members", "members"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberId"),
	}
}

// String returns a human-readable description of this Me Chat Id Member ID
func (id MeChatIdMemberId) String() string {
	components := []string{
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Me Chat Id Member (%s)", strings.Join(components, "\n"))
}
