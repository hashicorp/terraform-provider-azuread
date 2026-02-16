package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdThreadIdPostIdInReplyToMentionId{}

// GroupIdThreadIdPostIdInReplyToMentionId is a struct representing the Resource ID for a Group Id Thread Id Post Id In Reply To Mention
type GroupIdThreadIdPostIdInReplyToMentionId struct {
	GroupId              string
	ConversationThreadId string
	PostId               string
	MentionId            string
}

// NewGroupIdThreadIdPostIdInReplyToMentionID returns a new GroupIdThreadIdPostIdInReplyToMentionId struct
func NewGroupIdThreadIdPostIdInReplyToMentionID(groupId string, conversationThreadId string, postId string, mentionId string) GroupIdThreadIdPostIdInReplyToMentionId {
	return GroupIdThreadIdPostIdInReplyToMentionId{
		GroupId:              groupId,
		ConversationThreadId: conversationThreadId,
		PostId:               postId,
		MentionId:            mentionId,
	}
}

// ParseGroupIdThreadIdPostIdInReplyToMentionID parses 'input' into a GroupIdThreadIdPostIdInReplyToMentionId
func ParseGroupIdThreadIdPostIdInReplyToMentionID(input string) (*GroupIdThreadIdPostIdInReplyToMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdThreadIdPostIdInReplyToMentionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdThreadIdPostIdInReplyToMentionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdThreadIdPostIdInReplyToMentionIDInsensitively parses 'input' case-insensitively into a GroupIdThreadIdPostIdInReplyToMentionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdThreadIdPostIdInReplyToMentionIDInsensitively(input string) (*GroupIdThreadIdPostIdInReplyToMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdThreadIdPostIdInReplyToMentionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdThreadIdPostIdInReplyToMentionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdThreadIdPostIdInReplyToMentionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
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

// ValidateGroupIdThreadIdPostIdInReplyToMentionID checks that 'input' can be parsed as a Group Id Thread Id Post Id In Reply To Mention ID
func ValidateGroupIdThreadIdPostIdInReplyToMentionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdThreadIdPostIdInReplyToMentionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Thread Id Post Id In Reply To Mention ID
func (id GroupIdThreadIdPostIdInReplyToMentionId) ID() string {
	fmtString := "/groups/%s/threads/%s/posts/%s/inReplyTo/mentions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationThreadId, id.PostId, id.MentionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Thread Id Post Id In Reply To Mention ID
func (id GroupIdThreadIdPostIdInReplyToMentionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("threads", "threads", "threads"),
		resourceids.UserSpecifiedSegment("conversationThreadId", "conversationThreadId"),
		resourceids.StaticSegment("posts", "posts", "posts"),
		resourceids.UserSpecifiedSegment("postId", "postId"),
		resourceids.StaticSegment("inReplyTo", "inReplyTo", "inReplyTo"),
		resourceids.StaticSegment("mentions", "mentions", "mentions"),
		resourceids.UserSpecifiedSegment("mentionId", "mentionId"),
	}
}

// String returns a human-readable description of this Group Id Thread Id Post Id In Reply To Mention ID
func (id GroupIdThreadIdPostIdInReplyToMentionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
		fmt.Sprintf("Post: %q", id.PostId),
		fmt.Sprintf("Mention: %q", id.MentionId),
	}
	return fmt.Sprintf("Group Id Thread Id Post Id In Reply To Mention (%s)", strings.Join(components, "\n"))
}
