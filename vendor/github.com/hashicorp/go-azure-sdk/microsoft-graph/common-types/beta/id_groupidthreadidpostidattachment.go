package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdThreadIdPostIdAttachmentId{}

// GroupIdThreadIdPostIdAttachmentId is a struct representing the Resource ID for a Group Id Thread Id Post Id Attachment
type GroupIdThreadIdPostIdAttachmentId struct {
	GroupId              string
	ConversationThreadId string
	PostId               string
	AttachmentId         string
}

// NewGroupIdThreadIdPostIdAttachmentID returns a new GroupIdThreadIdPostIdAttachmentId struct
func NewGroupIdThreadIdPostIdAttachmentID(groupId string, conversationThreadId string, postId string, attachmentId string) GroupIdThreadIdPostIdAttachmentId {
	return GroupIdThreadIdPostIdAttachmentId{
		GroupId:              groupId,
		ConversationThreadId: conversationThreadId,
		PostId:               postId,
		AttachmentId:         attachmentId,
	}
}

// ParseGroupIdThreadIdPostIdAttachmentID parses 'input' into a GroupIdThreadIdPostIdAttachmentId
func ParseGroupIdThreadIdPostIdAttachmentID(input string) (*GroupIdThreadIdPostIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdThreadIdPostIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdThreadIdPostIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdThreadIdPostIdAttachmentIDInsensitively parses 'input' case-insensitively into a GroupIdThreadIdPostIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseGroupIdThreadIdPostIdAttachmentIDInsensitively(input string) (*GroupIdThreadIdPostIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdThreadIdPostIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdThreadIdPostIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdThreadIdPostIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateGroupIdThreadIdPostIdAttachmentID checks that 'input' can be parsed as a Group Id Thread Id Post Id Attachment ID
func ValidateGroupIdThreadIdPostIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdThreadIdPostIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Thread Id Post Id Attachment ID
func (id GroupIdThreadIdPostIdAttachmentId) ID() string {
	fmtString := "/groups/%s/threads/%s/posts/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationThreadId, id.PostId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Thread Id Post Id Attachment ID
func (id GroupIdThreadIdPostIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("threads", "threads", "threads"),
		resourceids.UserSpecifiedSegment("conversationThreadId", "conversationThreadId"),
		resourceids.StaticSegment("posts", "posts", "posts"),
		resourceids.UserSpecifiedSegment("postId", "postId"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentId"),
	}
}

// String returns a human-readable description of this Group Id Thread Id Post Id Attachment ID
func (id GroupIdThreadIdPostIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
		fmt.Sprintf("Post: %q", id.PostId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Group Id Thread Id Post Id Attachment (%s)", strings.Join(components, "\n"))
}
