package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdThreadIdPostId{}

// GroupIdThreadIdPostId is a struct representing the Resource ID for a Group Id Thread Id Post
type GroupIdThreadIdPostId struct {
	GroupId              string
	ConversationThreadId string
	PostId               string
}

// NewGroupIdThreadIdPostID returns a new GroupIdThreadIdPostId struct
func NewGroupIdThreadIdPostID(groupId string, conversationThreadId string, postId string) GroupIdThreadIdPostId {
	return GroupIdThreadIdPostId{
		GroupId:              groupId,
		ConversationThreadId: conversationThreadId,
		PostId:               postId,
	}
}

// ParseGroupIdThreadIdPostID parses 'input' into a GroupIdThreadIdPostId
func ParseGroupIdThreadIdPostID(input string) (*GroupIdThreadIdPostId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdThreadIdPostId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdThreadIdPostId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdThreadIdPostIDInsensitively parses 'input' case-insensitively into a GroupIdThreadIdPostId
// note: this method should only be used for API response data and not user input
func ParseGroupIdThreadIdPostIDInsensitively(input string) (*GroupIdThreadIdPostId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdThreadIdPostId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdThreadIdPostId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdThreadIdPostId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateGroupIdThreadIdPostID checks that 'input' can be parsed as a Group Id Thread Id Post ID
func ValidateGroupIdThreadIdPostID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdThreadIdPostID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Thread Id Post ID
func (id GroupIdThreadIdPostId) ID() string {
	fmtString := "/groups/%s/threads/%s/posts/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationThreadId, id.PostId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Thread Id Post ID
func (id GroupIdThreadIdPostId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("threads", "threads", "threads"),
		resourceids.UserSpecifiedSegment("conversationThreadId", "conversationThreadId"),
		resourceids.StaticSegment("posts", "posts", "posts"),
		resourceids.UserSpecifiedSegment("postId", "postId"),
	}
}

// String returns a human-readable description of this Group Id Thread Id Post ID
func (id GroupIdThreadIdPostId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
		fmt.Sprintf("Post: %q", id.PostId),
	}
	return fmt.Sprintf("Group Id Thread Id Post (%s)", strings.Join(components, "\n"))
}
