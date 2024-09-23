package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamChannelIdMessageIdReplyId{}

// GroupIdTeamChannelIdMessageIdReplyId is a struct representing the Resource ID for a Group Id Team Channel Id Message Id Reply
type GroupIdTeamChannelIdMessageIdReplyId struct {
	GroupId        string
	ChannelId      string
	ChatMessageId  string
	ChatMessageId1 string
}

// NewGroupIdTeamChannelIdMessageIdReplyID returns a new GroupIdTeamChannelIdMessageIdReplyId struct
func NewGroupIdTeamChannelIdMessageIdReplyID(groupId string, channelId string, chatMessageId string, chatMessageId1 string) GroupIdTeamChannelIdMessageIdReplyId {
	return GroupIdTeamChannelIdMessageIdReplyId{
		GroupId:        groupId,
		ChannelId:      channelId,
		ChatMessageId:  chatMessageId,
		ChatMessageId1: chatMessageId1,
	}
}

// ParseGroupIdTeamChannelIdMessageIdReplyID parses 'input' into a GroupIdTeamChannelIdMessageIdReplyId
func ParseGroupIdTeamChannelIdMessageIdReplyID(input string) (*GroupIdTeamChannelIdMessageIdReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelIdMessageIdReplyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelIdMessageIdReplyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamChannelIdMessageIdReplyIDInsensitively parses 'input' case-insensitively into a GroupIdTeamChannelIdMessageIdReplyId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamChannelIdMessageIdReplyIDInsensitively(input string) (*GroupIdTeamChannelIdMessageIdReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelIdMessageIdReplyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelIdMessageIdReplyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamChannelIdMessageIdReplyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ChannelId, ok = input.Parsed["channelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "channelId", input)
	}

	if id.ChatMessageId, ok = input.Parsed["chatMessageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", input)
	}

	if id.ChatMessageId1, ok = input.Parsed["chatMessageId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId1", input)
	}

	return nil
}

// ValidateGroupIdTeamChannelIdMessageIdReplyID checks that 'input' can be parsed as a Group Id Team Channel Id Message Id Reply ID
func ValidateGroupIdTeamChannelIdMessageIdReplyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamChannelIdMessageIdReplyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Channel Id Message Id Reply ID
func (id GroupIdTeamChannelIdMessageIdReplyId) ID() string {
	fmtString := "/groups/%s/team/channels/%s/messages/%s/replies/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId, id.ChatMessageId, id.ChatMessageId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Channel Id Message Id Reply ID
func (id GroupIdTeamChannelIdMessageIdReplyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
		resourceids.StaticSegment("replies", "replies", "replies"),
		resourceids.UserSpecifiedSegment("chatMessageId1", "chatMessageId1"),
	}
}

// String returns a human-readable description of this Group Id Team Channel Id Message Id Reply ID
func (id GroupIdTeamChannelIdMessageIdReplyId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
	}
	return fmt.Sprintf("Group Id Team Channel Id Message Id Reply (%s)", strings.Join(components, "\n"))
}
