package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId{}

// GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId is a struct representing the Resource ID for a Group Id Conversation Id Thread Id Post Id In Reply To Attachment
type GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId struct {
	GroupId              string
	ConversationId       string
	ConversationThreadId string
	PostId               string
	AttachmentId         string
}

// NewGroupIdConversationIdThreadIdPostIdInReplyToAttachmentID returns a new GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId struct
func NewGroupIdConversationIdThreadIdPostIdInReplyToAttachmentID(groupId string, conversationId string, conversationThreadId string, postId string, attachmentId string) GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId {
	return GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId{
		GroupId:              groupId,
		ConversationId:       conversationId,
		ConversationThreadId: conversationThreadId,
		PostId:               postId,
		AttachmentId:         attachmentId,
	}
}

// ParseGroupIdConversationIdThreadIdPostIdInReplyToAttachmentID parses 'input' into a GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId
func ParseGroupIdConversationIdThreadIdPostIdInReplyToAttachmentID(input string) (*GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdConversationIdThreadIdPostIdInReplyToAttachmentIDInsensitively parses 'input' case-insensitively into a GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId
// note: this method should only be used for API response data and not user input
func ParseGroupIdConversationIdThreadIdPostIdInReplyToAttachmentIDInsensitively(input string) (*GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdConversationIdThreadIdPostIdInReplyToAttachmentID checks that 'input' can be parsed as a Group Id Conversation Id Thread Id Post Id In Reply To Attachment ID
func ValidateGroupIdConversationIdThreadIdPostIdInReplyToAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdConversationIdThreadIdPostIdInReplyToAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Conversation Id Thread Id Post Id In Reply To Attachment ID
func (id GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId) ID() string {
	fmtString := "/groups/%s/conversations/%s/threads/%s/posts/%s/inReplyTo/attachments/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationId, id.ConversationThreadId, id.PostId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Conversation Id Thread Id Post Id In Reply To Attachment ID
func (id GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("conversations", "conversations", "conversations"),
		resourceids.UserSpecifiedSegment("conversationId", "conversationId"),
		resourceids.StaticSegment("threads", "threads", "threads"),
		resourceids.UserSpecifiedSegment("conversationThreadId", "conversationThreadId"),
		resourceids.StaticSegment("posts", "posts", "posts"),
		resourceids.UserSpecifiedSegment("postId", "postId"),
		resourceids.StaticSegment("inReplyTo", "inReplyTo", "inReplyTo"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentId"),
	}
}

// String returns a human-readable description of this Group Id Conversation Id Thread Id Post Id In Reply To Attachment ID
func (id GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation: %q", id.ConversationId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
		fmt.Sprintf("Post: %q", id.PostId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Group Id Conversation Id Thread Id Post Id In Reply To Attachment (%s)", strings.Join(components, "\n"))
}
