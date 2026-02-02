package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdChannelIdMessageId{}

// MeJoinedTeamIdChannelIdMessageId is a struct representing the Resource ID for a Me Joined Team Id Channel Id Message
type MeJoinedTeamIdChannelIdMessageId struct {
	TeamId        string
	ChannelId     string
	ChatMessageId string
}

// NewMeJoinedTeamIdChannelIdMessageID returns a new MeJoinedTeamIdChannelIdMessageId struct
func NewMeJoinedTeamIdChannelIdMessageID(teamId string, channelId string, chatMessageId string) MeJoinedTeamIdChannelIdMessageId {
	return MeJoinedTeamIdChannelIdMessageId{
		TeamId:        teamId,
		ChannelId:     channelId,
		ChatMessageId: chatMessageId,
	}
}

// ParseMeJoinedTeamIdChannelIdMessageID parses 'input' into a MeJoinedTeamIdChannelIdMessageId
func ParseMeJoinedTeamIdChannelIdMessageID(input string) (*MeJoinedTeamIdChannelIdMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdChannelIdMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdChannelIdMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdChannelIdMessageIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdChannelIdMessageId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdChannelIdMessageIDInsensitively(input string) (*MeJoinedTeamIdChannelIdMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdChannelIdMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdChannelIdMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdChannelIdMessageId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateMeJoinedTeamIdChannelIdMessageID checks that 'input' can be parsed as a Me Joined Team Id Channel Id Message ID
func ValidateMeJoinedTeamIdChannelIdMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdChannelIdMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Channel Id Message ID
func (id MeJoinedTeamIdChannelIdMessageId) ID() string {
	fmtString := "/me/joinedTeams/%s/channels/%s/messages/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChannelId, id.ChatMessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Channel Id Message ID
func (id MeJoinedTeamIdChannelIdMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Channel Id Message ID
func (id MeJoinedTeamIdChannelIdMessageId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
	}
	return fmt.Sprintf("Me Joined Team Id Channel Id Message (%s)", strings.Join(components, "\n"))
}
