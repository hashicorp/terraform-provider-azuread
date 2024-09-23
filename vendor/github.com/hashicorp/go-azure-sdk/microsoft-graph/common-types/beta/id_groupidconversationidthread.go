package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdConversationIdThreadId{}

// GroupIdConversationIdThreadId is a struct representing the Resource ID for a Group Id Conversation Id Thread
type GroupIdConversationIdThreadId struct {
	GroupId              string
	ConversationId       string
	ConversationThreadId string
}

// NewGroupIdConversationIdThreadID returns a new GroupIdConversationIdThreadId struct
func NewGroupIdConversationIdThreadID(groupId string, conversationId string, conversationThreadId string) GroupIdConversationIdThreadId {
	return GroupIdConversationIdThreadId{
		GroupId:              groupId,
		ConversationId:       conversationId,
		ConversationThreadId: conversationThreadId,
	}
}

// ParseGroupIdConversationIdThreadID parses 'input' into a GroupIdConversationIdThreadId
func ParseGroupIdConversationIdThreadID(input string) (*GroupIdConversationIdThreadId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdConversationIdThreadId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdConversationIdThreadId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdConversationIdThreadIDInsensitively parses 'input' case-insensitively into a GroupIdConversationIdThreadId
// note: this method should only be used for API response data and not user input
func ParseGroupIdConversationIdThreadIDInsensitively(input string) (*GroupIdConversationIdThreadId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdConversationIdThreadId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdConversationIdThreadId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdConversationIdThreadId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ConversationId, ok = input.Parsed["conversationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationId", input)
	}

	if id.ConversationThreadId, ok = input.Parsed["conversationThreadId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationThreadId", input)
	}

	return nil
}

// ValidateGroupIdConversationIdThreadID checks that 'input' can be parsed as a Group Id Conversation Id Thread ID
func ValidateGroupIdConversationIdThreadID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdConversationIdThreadID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Conversation Id Thread ID
func (id GroupIdConversationIdThreadId) ID() string {
	fmtString := "/groups/%s/conversations/%s/threads/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationId, id.ConversationThreadId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Conversation Id Thread ID
func (id GroupIdConversationIdThreadId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("conversations", "conversations", "conversations"),
		resourceids.UserSpecifiedSegment("conversationId", "conversationId"),
		resourceids.StaticSegment("threads", "threads", "threads"),
		resourceids.UserSpecifiedSegment("conversationThreadId", "conversationThreadId"),
	}
}

// String returns a human-readable description of this Group Id Conversation Id Thread ID
func (id GroupIdConversationIdThreadId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation: %q", id.ConversationId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
	}
	return fmt.Sprintf("Group Id Conversation Id Thread (%s)", strings.Join(components, "\n"))
}
