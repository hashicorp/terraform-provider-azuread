package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeChatIdOperationId{}

// MeChatIdOperationId is a struct representing the Resource ID for a Me Chat Id Operation
type MeChatIdOperationId struct {
	ChatId                string
	TeamsAsyncOperationId string
}

// NewMeChatIdOperationID returns a new MeChatIdOperationId struct
func NewMeChatIdOperationID(chatId string, teamsAsyncOperationId string) MeChatIdOperationId {
	return MeChatIdOperationId{
		ChatId:                chatId,
		TeamsAsyncOperationId: teamsAsyncOperationId,
	}
}

// ParseMeChatIdOperationID parses 'input' into a MeChatIdOperationId
func ParseMeChatIdOperationID(input string) (*MeChatIdOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeChatIdOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeChatIdOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeChatIdOperationIDInsensitively parses 'input' case-insensitively into a MeChatIdOperationId
// note: this method should only be used for API response data and not user input
func ParseMeChatIdOperationIDInsensitively(input string) (*MeChatIdOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeChatIdOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeChatIdOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeChatIdOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ChatId, ok = input.Parsed["chatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatId", input)
	}

	if id.TeamsAsyncOperationId, ok = input.Parsed["teamsAsyncOperationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamsAsyncOperationId", input)
	}

	return nil
}

// ValidateMeChatIdOperationID checks that 'input' can be parsed as a Me Chat Id Operation ID
func ValidateMeChatIdOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeChatIdOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Chat Id Operation ID
func (id MeChatIdOperationId) ID() string {
	fmtString := "/me/chats/%s/operations/%s"
	return fmt.Sprintf(fmtString, id.ChatId, id.TeamsAsyncOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Chat Id Operation ID
func (id MeChatIdOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("chats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatId"),
		resourceids.StaticSegment("operations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("teamsAsyncOperationId", "teamsAsyncOperationId"),
	}
}

// String returns a human-readable description of this Me Chat Id Operation ID
func (id MeChatIdOperationId) String() string {
	components := []string{
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Teams Async Operation: %q", id.TeamsAsyncOperationId),
	}
	return fmt.Sprintf("Me Chat Id Operation (%s)", strings.Join(components, "\n"))
}
