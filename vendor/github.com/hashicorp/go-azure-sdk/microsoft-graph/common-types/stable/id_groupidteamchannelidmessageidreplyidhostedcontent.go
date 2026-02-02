package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamChannelIdMessageIdReplyIdHostedContentId{}

// GroupIdTeamChannelIdMessageIdReplyIdHostedContentId is a struct representing the Resource ID for a Group Id Team Channel Id Message Id Reply Id Hosted Content
type GroupIdTeamChannelIdMessageIdReplyIdHostedContentId struct {
	GroupId                    string
	ChannelId                  string
	ChatMessageId              string
	ChatMessageId1             string
	ChatMessageHostedContentId string
}

// NewGroupIdTeamChannelIdMessageIdReplyIdHostedContentID returns a new GroupIdTeamChannelIdMessageIdReplyIdHostedContentId struct
func NewGroupIdTeamChannelIdMessageIdReplyIdHostedContentID(groupId string, channelId string, chatMessageId string, chatMessageId1 string, chatMessageHostedContentId string) GroupIdTeamChannelIdMessageIdReplyIdHostedContentId {
	return GroupIdTeamChannelIdMessageIdReplyIdHostedContentId{
		GroupId:                    groupId,
		ChannelId:                  channelId,
		ChatMessageId:              chatMessageId,
		ChatMessageId1:             chatMessageId1,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseGroupIdTeamChannelIdMessageIdReplyIdHostedContentID parses 'input' into a GroupIdTeamChannelIdMessageIdReplyIdHostedContentId
func ParseGroupIdTeamChannelIdMessageIdReplyIdHostedContentID(input string) (*GroupIdTeamChannelIdMessageIdReplyIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelIdMessageIdReplyIdHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelIdMessageIdReplyIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamChannelIdMessageIdReplyIdHostedContentIDInsensitively parses 'input' case-insensitively into a GroupIdTeamChannelIdMessageIdReplyIdHostedContentId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamChannelIdMessageIdReplyIdHostedContentIDInsensitively(input string) (*GroupIdTeamChannelIdMessageIdReplyIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelIdMessageIdReplyIdHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelIdMessageIdReplyIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamChannelIdMessageIdReplyIdHostedContentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ChatMessageHostedContentId, ok = input.Parsed["chatMessageHostedContentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", input)
	}

	return nil
}

// ValidateGroupIdTeamChannelIdMessageIdReplyIdHostedContentID checks that 'input' can be parsed as a Group Id Team Channel Id Message Id Reply Id Hosted Content ID
func ValidateGroupIdTeamChannelIdMessageIdReplyIdHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamChannelIdMessageIdReplyIdHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Channel Id Message Id Reply Id Hosted Content ID
func (id GroupIdTeamChannelIdMessageIdReplyIdHostedContentId) ID() string {
	fmtString := "/groups/%s/team/channels/%s/messages/%s/replies/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId, id.ChatMessageId, id.ChatMessageId1, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Channel Id Message Id Reply Id Hosted Content ID
func (id GroupIdTeamChannelIdMessageIdReplyIdHostedContentId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("hostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentId"),
	}
}

// String returns a human-readable description of this Group Id Team Channel Id Message Id Reply Id Hosted Content ID
func (id GroupIdTeamChannelIdMessageIdReplyIdHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("Group Id Team Channel Id Message Id Reply Id Hosted Content (%s)", strings.Join(components, "\n"))
}
