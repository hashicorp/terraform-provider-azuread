package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdThreadIdPostIdExtensionId{}

// GroupIdThreadIdPostIdExtensionId is a struct representing the Resource ID for a Group Id Thread Id Post Id Extension
type GroupIdThreadIdPostIdExtensionId struct {
	GroupId              string
	ConversationThreadId string
	PostId               string
	ExtensionId          string
}

// NewGroupIdThreadIdPostIdExtensionID returns a new GroupIdThreadIdPostIdExtensionId struct
func NewGroupIdThreadIdPostIdExtensionID(groupId string, conversationThreadId string, postId string, extensionId string) GroupIdThreadIdPostIdExtensionId {
	return GroupIdThreadIdPostIdExtensionId{
		GroupId:              groupId,
		ConversationThreadId: conversationThreadId,
		PostId:               postId,
		ExtensionId:          extensionId,
	}
}

// ParseGroupIdThreadIdPostIdExtensionID parses 'input' into a GroupIdThreadIdPostIdExtensionId
func ParseGroupIdThreadIdPostIdExtensionID(input string) (*GroupIdThreadIdPostIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdThreadIdPostIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdThreadIdPostIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdThreadIdPostIdExtensionIDInsensitively parses 'input' case-insensitively into a GroupIdThreadIdPostIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdThreadIdPostIdExtensionIDInsensitively(input string) (*GroupIdThreadIdPostIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdThreadIdPostIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdThreadIdPostIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdThreadIdPostIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateGroupIdThreadIdPostIdExtensionID checks that 'input' can be parsed as a Group Id Thread Id Post Id Extension ID
func ValidateGroupIdThreadIdPostIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdThreadIdPostIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Thread Id Post Id Extension ID
func (id GroupIdThreadIdPostIdExtensionId) ID() string {
	fmtString := "/groups/%s/threads/%s/posts/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationThreadId, id.PostId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Thread Id Post Id Extension ID
func (id GroupIdThreadIdPostIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("threads", "threads", "threads"),
		resourceids.UserSpecifiedSegment("conversationThreadId", "conversationThreadId"),
		resourceids.StaticSegment("posts", "posts", "posts"),
		resourceids.UserSpecifiedSegment("postId", "postId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Group Id Thread Id Post Id Extension ID
func (id GroupIdThreadIdPostIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
		fmt.Sprintf("Post: %q", id.PostId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Id Thread Id Post Id Extension (%s)", strings.Join(components, "\n"))
}
