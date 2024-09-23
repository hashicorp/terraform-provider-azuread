package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdChannelIdMessageIdHostedContentId{}

// MeJoinedTeamIdChannelIdMessageIdHostedContentId is a struct representing the Resource ID for a Me Joined Team Id Channel Id Message Id Hosted Content
type MeJoinedTeamIdChannelIdMessageIdHostedContentId struct {
	TeamId                     string
	ChannelId                  string
	ChatMessageId              string
	ChatMessageHostedContentId string
}

// NewMeJoinedTeamIdChannelIdMessageIdHostedContentID returns a new MeJoinedTeamIdChannelIdMessageIdHostedContentId struct
func NewMeJoinedTeamIdChannelIdMessageIdHostedContentID(teamId string, channelId string, chatMessageId string, chatMessageHostedContentId string) MeJoinedTeamIdChannelIdMessageIdHostedContentId {
	return MeJoinedTeamIdChannelIdMessageIdHostedContentId{
		TeamId:                     teamId,
		ChannelId:                  channelId,
		ChatMessageId:              chatMessageId,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseMeJoinedTeamIdChannelIdMessageIdHostedContentID parses 'input' into a MeJoinedTeamIdChannelIdMessageIdHostedContentId
func ParseMeJoinedTeamIdChannelIdMessageIdHostedContentID(input string) (*MeJoinedTeamIdChannelIdMessageIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdChannelIdMessageIdHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdChannelIdMessageIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdChannelIdMessageIdHostedContentIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdChannelIdMessageIdHostedContentId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdChannelIdMessageIdHostedContentIDInsensitively(input string) (*MeJoinedTeamIdChannelIdMessageIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdChannelIdMessageIdHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdChannelIdMessageIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdChannelIdMessageIdHostedContentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ChatMessageHostedContentId, ok = input.Parsed["chatMessageHostedContentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdChannelIdMessageIdHostedContentID checks that 'input' can be parsed as a Me Joined Team Id Channel Id Message Id Hosted Content ID
func ValidateMeJoinedTeamIdChannelIdMessageIdHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdChannelIdMessageIdHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Channel Id Message Id Hosted Content ID
func (id MeJoinedTeamIdChannelIdMessageIdHostedContentId) ID() string {
	fmtString := "/me/joinedTeams/%s/channels/%s/messages/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChannelId, id.ChatMessageId, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Channel Id Message Id Hosted Content ID
func (id MeJoinedTeamIdChannelIdMessageIdHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
		resourceids.StaticSegment("hostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Channel Id Message Id Hosted Content ID
func (id MeJoinedTeamIdChannelIdMessageIdHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("Me Joined Team Id Channel Id Message Id Hosted Content (%s)", strings.Join(components, "\n"))
}
