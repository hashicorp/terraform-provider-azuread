package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId{}

// MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId is a struct representing the Resource ID for a Me Joined Team Id Channel Id Message Id Reply Id Hosted Content
type MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId struct {
	TeamId                     string
	ChannelId                  string
	ChatMessageId              string
	ChatMessageId1             string
	ChatMessageHostedContentId string
}

// NewMeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentID returns a new MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId struct
func NewMeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentID(teamId string, channelId string, chatMessageId string, chatMessageId1 string, chatMessageHostedContentId string) MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId {
	return MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId{
		TeamId:                     teamId,
		ChannelId:                  channelId,
		ChatMessageId:              chatMessageId,
		ChatMessageId1:             chatMessageId1,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseMeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentID parses 'input' into a MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId
func ParseMeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentID(input string) (*MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentIDInsensitively(input string) (*MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ChatMessageHostedContentId, ok = input.Parsed["chatMessageHostedContentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentID checks that 'input' can be parsed as a Me Joined Team Id Channel Id Message Id Reply Id Hosted Content ID
func ValidateMeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Channel Id Message Id Reply Id Hosted Content ID
func (id MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId) ID() string {
	fmtString := "/me/joinedTeams/%s/channels/%s/messages/%s/replies/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChannelId, id.ChatMessageId, id.ChatMessageId1, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Channel Id Message Id Reply Id Hosted Content ID
func (id MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("hostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Channel Id Message Id Reply Id Hosted Content ID
func (id MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("Me Joined Team Id Channel Id Message Id Reply Id Hosted Content (%s)", strings.Join(components, "\n"))
}
