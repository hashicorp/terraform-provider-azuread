package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeChatIdTabId{}

// MeChatIdTabId is a struct representing the Resource ID for a Me Chat Id Tab
type MeChatIdTabId struct {
	ChatId     string
	TeamsTabId string
}

// NewMeChatIdTabID returns a new MeChatIdTabId struct
func NewMeChatIdTabID(chatId string, teamsTabId string) MeChatIdTabId {
	return MeChatIdTabId{
		ChatId:     chatId,
		TeamsTabId: teamsTabId,
	}
}

// ParseMeChatIdTabID parses 'input' into a MeChatIdTabId
func ParseMeChatIdTabID(input string) (*MeChatIdTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeChatIdTabId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeChatIdTabId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeChatIdTabIDInsensitively parses 'input' case-insensitively into a MeChatIdTabId
// note: this method should only be used for API response data and not user input
func ParseMeChatIdTabIDInsensitively(input string) (*MeChatIdTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeChatIdTabId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeChatIdTabId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeChatIdTabId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ChatId, ok = input.Parsed["chatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatId", input)
	}

	if id.TeamsTabId, ok = input.Parsed["teamsTabId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", input)
	}

	return nil
}

// ValidateMeChatIdTabID checks that 'input' can be parsed as a Me Chat Id Tab ID
func ValidateMeChatIdTabID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeChatIdTabID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Chat Id Tab ID
func (id MeChatIdTabId) ID() string {
	fmtString := "/me/chats/%s/tabs/%s"
	return fmt.Sprintf(fmtString, id.ChatId, id.TeamsTabId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Chat Id Tab ID
func (id MeChatIdTabId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("chats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatId"),
		resourceids.StaticSegment("tabs", "tabs", "tabs"),
		resourceids.UserSpecifiedSegment("teamsTabId", "teamsTabId"),
	}
}

// String returns a human-readable description of this Me Chat Id Tab ID
func (id MeChatIdTabId) String() string {
	components := []string{
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Teams Tab: %q", id.TeamsTabId),
	}
	return fmt.Sprintf("Me Chat Id Tab (%s)", strings.Join(components, "\n"))
}
