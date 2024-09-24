package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdThreadIdPostIdMentionId{}

// GroupIdThreadIdPostIdMentionId is a struct representing the Resource ID for a Group Id Thread Id Post Id Mention
type GroupIdThreadIdPostIdMentionId struct {
	GroupId              string
	ConversationThreadId string
	PostId               string
	MentionId            string
}

// NewGroupIdThreadIdPostIdMentionID returns a new GroupIdThreadIdPostIdMentionId struct
func NewGroupIdThreadIdPostIdMentionID(groupId string, conversationThreadId string, postId string, mentionId string) GroupIdThreadIdPostIdMentionId {
	return GroupIdThreadIdPostIdMentionId{
		GroupId:              groupId,
		ConversationThreadId: conversationThreadId,
		PostId:               postId,
		MentionId:            mentionId,
	}
}

// ParseGroupIdThreadIdPostIdMentionID parses 'input' into a GroupIdThreadIdPostIdMentionId
func ParseGroupIdThreadIdPostIdMentionID(input string) (*GroupIdThreadIdPostIdMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdThreadIdPostIdMentionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdThreadIdPostIdMentionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdThreadIdPostIdMentionIDInsensitively parses 'input' case-insensitively into a GroupIdThreadIdPostIdMentionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdThreadIdPostIdMentionIDInsensitively(input string) (*GroupIdThreadIdPostIdMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdThreadIdPostIdMentionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdThreadIdPostIdMentionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdThreadIdPostIdMentionId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdThreadIdPostIdMentionID checks that 'input' can be parsed as a Group Id Thread Id Post Id Mention ID
func ValidateGroupIdThreadIdPostIdMentionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdThreadIdPostIdMentionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Thread Id Post Id Mention ID
func (id GroupIdThreadIdPostIdMentionId) ID() string {
	fmtString := "/groups/%s/threads/%s/posts/%s/mentions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationThreadId, id.PostId, id.MentionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Thread Id Post Id Mention ID
func (id GroupIdThreadIdPostIdMentionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("threads", "threads", "threads"),
		resourceids.UserSpecifiedSegment("conversationThreadId", "conversationThreadId"),
		resourceids.StaticSegment("posts", "posts", "posts"),
		resourceids.UserSpecifiedSegment("postId", "postId"),
		resourceids.StaticSegment("mentions", "mentions", "mentions"),
		resourceids.UserSpecifiedSegment("mentionId", "mentionId"),
	}
}

// String returns a human-readable description of this Group Id Thread Id Post Id Mention ID
func (id GroupIdThreadIdPostIdMentionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
		fmt.Sprintf("Post: %q", id.PostId),
		fmt.Sprintf("Mention: %q", id.MentionId),
	}
	return fmt.Sprintf("Group Id Thread Id Post Id Mention (%s)", strings.Join(components, "\n"))
}
