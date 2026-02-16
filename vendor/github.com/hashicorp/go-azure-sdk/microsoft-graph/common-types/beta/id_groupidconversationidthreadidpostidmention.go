package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdConversationIdThreadIdPostIdMentionId{}

// GroupIdConversationIdThreadIdPostIdMentionId is a struct representing the Resource ID for a Group Id Conversation Id Thread Id Post Id Mention
type GroupIdConversationIdThreadIdPostIdMentionId struct {
	GroupId              string
	ConversationId       string
	ConversationThreadId string
	PostId               string
	MentionId            string
}

// NewGroupIdConversationIdThreadIdPostIdMentionID returns a new GroupIdConversationIdThreadIdPostIdMentionId struct
func NewGroupIdConversationIdThreadIdPostIdMentionID(groupId string, conversationId string, conversationThreadId string, postId string, mentionId string) GroupIdConversationIdThreadIdPostIdMentionId {
	return GroupIdConversationIdThreadIdPostIdMentionId{
		GroupId:              groupId,
		ConversationId:       conversationId,
		ConversationThreadId: conversationThreadId,
		PostId:               postId,
		MentionId:            mentionId,
	}
}

// ParseGroupIdConversationIdThreadIdPostIdMentionID parses 'input' into a GroupIdConversationIdThreadIdPostIdMentionId
func ParseGroupIdConversationIdThreadIdPostIdMentionID(input string) (*GroupIdConversationIdThreadIdPostIdMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdConversationIdThreadIdPostIdMentionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdConversationIdThreadIdPostIdMentionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdConversationIdThreadIdPostIdMentionIDInsensitively parses 'input' case-insensitively into a GroupIdConversationIdThreadIdPostIdMentionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdConversationIdThreadIdPostIdMentionIDInsensitively(input string) (*GroupIdConversationIdThreadIdPostIdMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdConversationIdThreadIdPostIdMentionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdConversationIdThreadIdPostIdMentionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdConversationIdThreadIdPostIdMentionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.PostId, ok = input.Parsed["postId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "postId", input)
	}

	if id.MentionId, ok = input.Parsed["mentionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mentionId", input)
	}

	return nil
}

// ValidateGroupIdConversationIdThreadIdPostIdMentionID checks that 'input' can be parsed as a Group Id Conversation Id Thread Id Post Id Mention ID
func ValidateGroupIdConversationIdThreadIdPostIdMentionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdConversationIdThreadIdPostIdMentionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Conversation Id Thread Id Post Id Mention ID
func (id GroupIdConversationIdThreadIdPostIdMentionId) ID() string {
	fmtString := "/groups/%s/conversations/%s/threads/%s/posts/%s/mentions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationId, id.ConversationThreadId, id.PostId, id.MentionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Conversation Id Thread Id Post Id Mention ID
func (id GroupIdConversationIdThreadIdPostIdMentionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("conversations", "conversations", "conversations"),
		resourceids.UserSpecifiedSegment("conversationId", "conversationId"),
		resourceids.StaticSegment("threads", "threads", "threads"),
		resourceids.UserSpecifiedSegment("conversationThreadId", "conversationThreadId"),
		resourceids.StaticSegment("posts", "posts", "posts"),
		resourceids.UserSpecifiedSegment("postId", "postId"),
		resourceids.StaticSegment("mentions", "mentions", "mentions"),
		resourceids.UserSpecifiedSegment("mentionId", "mentionId"),
	}
}

// String returns a human-readable description of this Group Id Conversation Id Thread Id Post Id Mention ID
func (id GroupIdConversationIdThreadIdPostIdMentionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation: %q", id.ConversationId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
		fmt.Sprintf("Post: %q", id.PostId),
		fmt.Sprintf("Mention: %q", id.MentionId),
	}
	return fmt.Sprintf("Group Id Conversation Id Thread Id Post Id Mention (%s)", strings.Join(components, "\n"))
}
