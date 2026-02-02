package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamPrimaryChannelMessageIdReplyId{}

// GroupIdTeamPrimaryChannelMessageIdReplyId is a struct representing the Resource ID for a Group Id Team Primary Channel Message Id Reply
type GroupIdTeamPrimaryChannelMessageIdReplyId struct {
	GroupId        string
	ChatMessageId  string
	ChatMessageId1 string
}

// NewGroupIdTeamPrimaryChannelMessageIdReplyID returns a new GroupIdTeamPrimaryChannelMessageIdReplyId struct
func NewGroupIdTeamPrimaryChannelMessageIdReplyID(groupId string, chatMessageId string, chatMessageId1 string) GroupIdTeamPrimaryChannelMessageIdReplyId {
	return GroupIdTeamPrimaryChannelMessageIdReplyId{
		GroupId:        groupId,
		ChatMessageId:  chatMessageId,
		ChatMessageId1: chatMessageId1,
	}
}

// ParseGroupIdTeamPrimaryChannelMessageIdReplyID parses 'input' into a GroupIdTeamPrimaryChannelMessageIdReplyId
func ParseGroupIdTeamPrimaryChannelMessageIdReplyID(input string) (*GroupIdTeamPrimaryChannelMessageIdReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelMessageIdReplyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelMessageIdReplyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamPrimaryChannelMessageIdReplyIDInsensitively parses 'input' case-insensitively into a GroupIdTeamPrimaryChannelMessageIdReplyId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamPrimaryChannelMessageIdReplyIDInsensitively(input string) (*GroupIdTeamPrimaryChannelMessageIdReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelMessageIdReplyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelMessageIdReplyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamPrimaryChannelMessageIdReplyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ChatMessageId, ok = input.Parsed["chatMessageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", input)
	}

	if id.ChatMessageId1, ok = input.Parsed["chatMessageId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId1", input)
	}

	return nil
}

// ValidateGroupIdTeamPrimaryChannelMessageIdReplyID checks that 'input' can be parsed as a Group Id Team Primary Channel Message Id Reply ID
func ValidateGroupIdTeamPrimaryChannelMessageIdReplyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamPrimaryChannelMessageIdReplyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Primary Channel Message Id Reply ID
func (id GroupIdTeamPrimaryChannelMessageIdReplyId) ID() string {
	fmtString := "/groups/%s/team/primaryChannel/messages/%s/replies/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChatMessageId, id.ChatMessageId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Primary Channel Message Id Reply ID
func (id GroupIdTeamPrimaryChannelMessageIdReplyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
		resourceids.StaticSegment("replies", "replies", "replies"),
		resourceids.UserSpecifiedSegment("chatMessageId1", "chatMessageId1"),
	}
}

// String returns a human-readable description of this Group Id Team Primary Channel Message Id Reply ID
func (id GroupIdTeamPrimaryChannelMessageIdReplyId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
	}
	return fmt.Sprintf("Group Id Team Primary Channel Message Id Reply (%s)", strings.Join(components, "\n"))
}
