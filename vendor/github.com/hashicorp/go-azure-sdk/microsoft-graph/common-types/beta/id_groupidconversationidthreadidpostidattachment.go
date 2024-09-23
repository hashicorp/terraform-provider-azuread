package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdConversationIdThreadIdPostIdAttachmentId{}

// GroupIdConversationIdThreadIdPostIdAttachmentId is a struct representing the Resource ID for a Group Id Conversation Id Thread Id Post Id Attachment
type GroupIdConversationIdThreadIdPostIdAttachmentId struct {
	GroupId              string
	ConversationId       string
	ConversationThreadId string
	PostId               string
	AttachmentId         string
}

// NewGroupIdConversationIdThreadIdPostIdAttachmentID returns a new GroupIdConversationIdThreadIdPostIdAttachmentId struct
func NewGroupIdConversationIdThreadIdPostIdAttachmentID(groupId string, conversationId string, conversationThreadId string, postId string, attachmentId string) GroupIdConversationIdThreadIdPostIdAttachmentId {
	return GroupIdConversationIdThreadIdPostIdAttachmentId{
		GroupId:              groupId,
		ConversationId:       conversationId,
		ConversationThreadId: conversationThreadId,
		PostId:               postId,
		AttachmentId:         attachmentId,
	}
}

// ParseGroupIdConversationIdThreadIdPostIdAttachmentID parses 'input' into a GroupIdConversationIdThreadIdPostIdAttachmentId
func ParseGroupIdConversationIdThreadIdPostIdAttachmentID(input string) (*GroupIdConversationIdThreadIdPostIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdConversationIdThreadIdPostIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdConversationIdThreadIdPostIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdConversationIdThreadIdPostIdAttachmentIDInsensitively parses 'input' case-insensitively into a GroupIdConversationIdThreadIdPostIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseGroupIdConversationIdThreadIdPostIdAttachmentIDInsensitively(input string) (*GroupIdConversationIdThreadIdPostIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdConversationIdThreadIdPostIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdConversationIdThreadIdPostIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdConversationIdThreadIdPostIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateGroupIdConversationIdThreadIdPostIdAttachmentID checks that 'input' can be parsed as a Group Id Conversation Id Thread Id Post Id Attachment ID
func ValidateGroupIdConversationIdThreadIdPostIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdConversationIdThreadIdPostIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Conversation Id Thread Id Post Id Attachment ID
func (id GroupIdConversationIdThreadIdPostIdAttachmentId) ID() string {
	fmtString := "/groups/%s/conversations/%s/threads/%s/posts/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationId, id.ConversationThreadId, id.PostId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Conversation Id Thread Id Post Id Attachment ID
func (id GroupIdConversationIdThreadIdPostIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("conversations", "conversations", "conversations"),
		resourceids.UserSpecifiedSegment("conversationId", "conversationId"),
		resourceids.StaticSegment("threads", "threads", "threads"),
		resourceids.UserSpecifiedSegment("conversationThreadId", "conversationThreadId"),
		resourceids.StaticSegment("posts", "posts", "posts"),
		resourceids.UserSpecifiedSegment("postId", "postId"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentId"),
	}
}

// String returns a human-readable description of this Group Id Conversation Id Thread Id Post Id Attachment ID
func (id GroupIdConversationIdThreadIdPostIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation: %q", id.ConversationId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
		fmt.Sprintf("Post: %q", id.PostId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Group Id Conversation Id Thread Id Post Id Attachment (%s)", strings.Join(components, "\n"))
}
