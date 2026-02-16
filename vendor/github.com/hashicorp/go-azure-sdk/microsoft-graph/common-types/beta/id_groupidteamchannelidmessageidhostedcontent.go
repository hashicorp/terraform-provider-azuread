package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamChannelIdMessageIdHostedContentId{}

// GroupIdTeamChannelIdMessageIdHostedContentId is a struct representing the Resource ID for a Group Id Team Channel Id Message Id Hosted Content
type GroupIdTeamChannelIdMessageIdHostedContentId struct {
	GroupId                    string
	ChannelId                  string
	ChatMessageId              string
	ChatMessageHostedContentId string
}

// NewGroupIdTeamChannelIdMessageIdHostedContentID returns a new GroupIdTeamChannelIdMessageIdHostedContentId struct
func NewGroupIdTeamChannelIdMessageIdHostedContentID(groupId string, channelId string, chatMessageId string, chatMessageHostedContentId string) GroupIdTeamChannelIdMessageIdHostedContentId {
	return GroupIdTeamChannelIdMessageIdHostedContentId{
		GroupId:                    groupId,
		ChannelId:                  channelId,
		ChatMessageId:              chatMessageId,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseGroupIdTeamChannelIdMessageIdHostedContentID parses 'input' into a GroupIdTeamChannelIdMessageIdHostedContentId
func ParseGroupIdTeamChannelIdMessageIdHostedContentID(input string) (*GroupIdTeamChannelIdMessageIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelIdMessageIdHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelIdMessageIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamChannelIdMessageIdHostedContentIDInsensitively parses 'input' case-insensitively into a GroupIdTeamChannelIdMessageIdHostedContentId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamChannelIdMessageIdHostedContentIDInsensitively(input string) (*GroupIdTeamChannelIdMessageIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelIdMessageIdHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelIdMessageIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamChannelIdMessageIdHostedContentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ChatMessageHostedContentId, ok = input.Parsed["chatMessageHostedContentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", input)
	}

	return nil
}

// ValidateGroupIdTeamChannelIdMessageIdHostedContentID checks that 'input' can be parsed as a Group Id Team Channel Id Message Id Hosted Content ID
func ValidateGroupIdTeamChannelIdMessageIdHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamChannelIdMessageIdHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Channel Id Message Id Hosted Content ID
func (id GroupIdTeamChannelIdMessageIdHostedContentId) ID() string {
	fmtString := "/groups/%s/team/channels/%s/messages/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId, id.ChatMessageId, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Channel Id Message Id Hosted Content ID
func (id GroupIdTeamChannelIdMessageIdHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
		resourceids.StaticSegment("hostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentId"),
	}
}

// String returns a human-readable description of this Group Id Team Channel Id Message Id Hosted Content ID
func (id GroupIdTeamChannelIdMessageIdHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("Group Id Team Channel Id Message Id Hosted Content (%s)", strings.Join(components, "\n"))
}
