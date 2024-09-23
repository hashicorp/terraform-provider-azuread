package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdChannelIdMessageIdReplyId{}

// MeJoinedTeamIdChannelIdMessageIdReplyId is a struct representing the Resource ID for a Me Joined Team Id Channel Id Message Id Reply
type MeJoinedTeamIdChannelIdMessageIdReplyId struct {
	TeamId         string
	ChannelId      string
	ChatMessageId  string
	ChatMessageId1 string
}

// NewMeJoinedTeamIdChannelIdMessageIdReplyID returns a new MeJoinedTeamIdChannelIdMessageIdReplyId struct
func NewMeJoinedTeamIdChannelIdMessageIdReplyID(teamId string, channelId string, chatMessageId string, chatMessageId1 string) MeJoinedTeamIdChannelIdMessageIdReplyId {
	return MeJoinedTeamIdChannelIdMessageIdReplyId{
		TeamId:         teamId,
		ChannelId:      channelId,
		ChatMessageId:  chatMessageId,
		ChatMessageId1: chatMessageId1,
	}
}

// ParseMeJoinedTeamIdChannelIdMessageIdReplyID parses 'input' into a MeJoinedTeamIdChannelIdMessageIdReplyId
func ParseMeJoinedTeamIdChannelIdMessageIdReplyID(input string) (*MeJoinedTeamIdChannelIdMessageIdReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdChannelIdMessageIdReplyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdChannelIdMessageIdReplyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdChannelIdMessageIdReplyIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdChannelIdMessageIdReplyId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdChannelIdMessageIdReplyIDInsensitively(input string) (*MeJoinedTeamIdChannelIdMessageIdReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdChannelIdMessageIdReplyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdChannelIdMessageIdReplyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdChannelIdMessageIdReplyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
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

// ValidateMeJoinedTeamIdChannelIdMessageIdReplyID checks that 'input' can be parsed as a Me Joined Team Id Channel Id Message Id Reply ID
func ValidateMeJoinedTeamIdChannelIdMessageIdReplyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdChannelIdMessageIdReplyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Channel Id Message Id Reply ID
func (id MeJoinedTeamIdChannelIdMessageIdReplyId) ID() string {
	fmtString := "/me/joinedTeams/%s/channels/%s/messages/%s/replies/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChannelId, id.ChatMessageId, id.ChatMessageId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Channel Id Message Id Reply ID
func (id MeJoinedTeamIdChannelIdMessageIdReplyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
		resourceids.StaticSegment("replies", "replies", "replies"),
		resourceids.UserSpecifiedSegment("chatMessageId1", "chatMessageId1"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Channel Id Message Id Reply ID
func (id MeJoinedTeamIdChannelIdMessageIdReplyId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
	}
	return fmt.Sprintf("Me Joined Team Id Channel Id Message Id Reply (%s)", strings.Join(components, "\n"))
}
