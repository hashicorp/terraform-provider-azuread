package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdConversationIdThreadIdPostIdInReplyToExtensionId{}

// GroupIdConversationIdThreadIdPostIdInReplyToExtensionId is a struct representing the Resource ID for a Group Id Conversation Id Thread Id Post Id In Reply To Extension
type GroupIdConversationIdThreadIdPostIdInReplyToExtensionId struct {
	GroupId              string
	ConversationId       string
	ConversationThreadId string
	PostId               string
	ExtensionId          string
}

// NewGroupIdConversationIdThreadIdPostIdInReplyToExtensionID returns a new GroupIdConversationIdThreadIdPostIdInReplyToExtensionId struct
func NewGroupIdConversationIdThreadIdPostIdInReplyToExtensionID(groupId string, conversationId string, conversationThreadId string, postId string, extensionId string) GroupIdConversationIdThreadIdPostIdInReplyToExtensionId {
	return GroupIdConversationIdThreadIdPostIdInReplyToExtensionId{
		GroupId:              groupId,
		ConversationId:       conversationId,
		ConversationThreadId: conversationThreadId,
		PostId:               postId,
		ExtensionId:          extensionId,
	}
}

// ParseGroupIdConversationIdThreadIdPostIdInReplyToExtensionID parses 'input' into a GroupIdConversationIdThreadIdPostIdInReplyToExtensionId
func ParseGroupIdConversationIdThreadIdPostIdInReplyToExtensionID(input string) (*GroupIdConversationIdThreadIdPostIdInReplyToExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdConversationIdThreadIdPostIdInReplyToExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdConversationIdThreadIdPostIdInReplyToExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdConversationIdThreadIdPostIdInReplyToExtensionIDInsensitively parses 'input' case-insensitively into a GroupIdConversationIdThreadIdPostIdInReplyToExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdConversationIdThreadIdPostIdInReplyToExtensionIDInsensitively(input string) (*GroupIdConversationIdThreadIdPostIdInReplyToExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdConversationIdThreadIdPostIdInReplyToExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdConversationIdThreadIdPostIdInReplyToExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdConversationIdThreadIdPostIdInReplyToExtensionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateGroupIdConversationIdThreadIdPostIdInReplyToExtensionID checks that 'input' can be parsed as a Group Id Conversation Id Thread Id Post Id In Reply To Extension ID
func ValidateGroupIdConversationIdThreadIdPostIdInReplyToExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdConversationIdThreadIdPostIdInReplyToExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Conversation Id Thread Id Post Id In Reply To Extension ID
func (id GroupIdConversationIdThreadIdPostIdInReplyToExtensionId) ID() string {
	fmtString := "/groups/%s/conversations/%s/threads/%s/posts/%s/inReplyTo/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationId, id.ConversationThreadId, id.PostId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Conversation Id Thread Id Post Id In Reply To Extension ID
func (id GroupIdConversationIdThreadIdPostIdInReplyToExtensionId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Group Id Conversation Id Thread Id Post Id In Reply To Extension ID
func (id GroupIdConversationIdThreadIdPostIdInReplyToExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation: %q", id.ConversationId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
		fmt.Sprintf("Post: %q", id.PostId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Id Conversation Id Thread Id Post Id In Reply To Extension (%s)", strings.Join(components, "\n"))
}
