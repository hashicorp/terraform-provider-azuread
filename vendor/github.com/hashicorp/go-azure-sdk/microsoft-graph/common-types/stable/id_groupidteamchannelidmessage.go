package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamChannelIdMessageId{}

// GroupIdTeamChannelIdMessageId is a struct representing the Resource ID for a Group Id Team Channel Id Message
type GroupIdTeamChannelIdMessageId struct {
	GroupId       string
	ChannelId     string
	ChatMessageId string
}

// NewGroupIdTeamChannelIdMessageID returns a new GroupIdTeamChannelIdMessageId struct
func NewGroupIdTeamChannelIdMessageID(groupId string, channelId string, chatMessageId string) GroupIdTeamChannelIdMessageId {
	return GroupIdTeamChannelIdMessageId{
		GroupId:       groupId,
		ChannelId:     channelId,
		ChatMessageId: chatMessageId,
	}
}

// ParseGroupIdTeamChannelIdMessageID parses 'input' into a GroupIdTeamChannelIdMessageId
func ParseGroupIdTeamChannelIdMessageID(input string) (*GroupIdTeamChannelIdMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelIdMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelIdMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamChannelIdMessageIDInsensitively parses 'input' case-insensitively into a GroupIdTeamChannelIdMessageId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamChannelIdMessageIDInsensitively(input string) (*GroupIdTeamChannelIdMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelIdMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelIdMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamChannelIdMessageId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateGroupIdTeamChannelIdMessageID checks that 'input' can be parsed as a Group Id Team Channel Id Message ID
func ValidateGroupIdTeamChannelIdMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamChannelIdMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Channel Id Message ID
func (id GroupIdTeamChannelIdMessageId) ID() string {
	fmtString := "/groups/%s/team/channels/%s/messages/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId, id.ChatMessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Channel Id Message ID
func (id GroupIdTeamChannelIdMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
	}
}

// String returns a human-readable description of this Group Id Team Channel Id Message ID
func (id GroupIdTeamChannelIdMessageId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
	}
	return fmt.Sprintf("Group Id Team Channel Id Message (%s)", strings.Join(components, "\n"))
}
