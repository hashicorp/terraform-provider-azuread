package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeChatId{}

// MeChatId is a struct representing the Resource ID for a Me Chat
type MeChatId struct {
	ChatId string
}

// NewMeChatID returns a new MeChatId struct
func NewMeChatID(chatId string) MeChatId {
	return MeChatId{
		ChatId: chatId,
	}
}

// ParseMeChatID parses 'input' into a MeChatId
func ParseMeChatID(input string) (*MeChatId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeChatId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeChatId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeChatIDInsensitively parses 'input' case-insensitively into a MeChatId
// note: this method should only be used for API response data and not user input
func ParseMeChatIDInsensitively(input string) (*MeChatId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeChatId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeChatId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeChatId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ChatId, ok = input.Parsed["chatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatId", input)
	}

	return nil
}

// ValidateMeChatID checks that 'input' can be parsed as a Me Chat ID
func ValidateMeChatID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeChatID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Chat ID
func (id MeChatId) ID() string {
	fmtString := "/me/chats/%s"
	return fmt.Sprintf(fmtString, id.ChatId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Chat ID
func (id MeChatId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("chats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatId"),
	}
}

// String returns a human-readable description of this Me Chat ID
func (id MeChatId) String() string {
	components := []string{
		fmt.Sprintf("Chat: %q", id.ChatId),
	}
	return fmt.Sprintf("Me Chat (%s)", strings.Join(components, "\n"))
}
