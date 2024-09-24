package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdChannelIdMemberId{}

// MeJoinedTeamIdChannelIdMemberId is a struct representing the Resource ID for a Me Joined Team Id Channel Id Member
type MeJoinedTeamIdChannelIdMemberId struct {
	TeamId               string
	ChannelId            string
	ConversationMemberId string
}

// NewMeJoinedTeamIdChannelIdMemberID returns a new MeJoinedTeamIdChannelIdMemberId struct
func NewMeJoinedTeamIdChannelIdMemberID(teamId string, channelId string, conversationMemberId string) MeJoinedTeamIdChannelIdMemberId {
	return MeJoinedTeamIdChannelIdMemberId{
		TeamId:               teamId,
		ChannelId:            channelId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseMeJoinedTeamIdChannelIdMemberID parses 'input' into a MeJoinedTeamIdChannelIdMemberId
func ParseMeJoinedTeamIdChannelIdMemberID(input string) (*MeJoinedTeamIdChannelIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdChannelIdMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdChannelIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdChannelIdMemberIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdChannelIdMemberId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdChannelIdMemberIDInsensitively(input string) (*MeJoinedTeamIdChannelIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdChannelIdMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdChannelIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdChannelIdMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.ChannelId, ok = input.Parsed["channelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "channelId", input)
	}

	if id.ConversationMemberId, ok = input.Parsed["conversationMemberId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdChannelIdMemberID checks that 'input' can be parsed as a Me Joined Team Id Channel Id Member ID
func ValidateMeJoinedTeamIdChannelIdMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdChannelIdMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Channel Id Member ID
func (id MeJoinedTeamIdChannelIdMemberId) ID() string {
	fmtString := "/me/joinedTeams/%s/channels/%s/members/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChannelId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Channel Id Member ID
func (id MeJoinedTeamIdChannelIdMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("members", "members", "members"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Channel Id Member ID
func (id MeJoinedTeamIdChannelIdMemberId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Me Joined Team Id Channel Id Member (%s)", strings.Join(components, "\n"))
}
