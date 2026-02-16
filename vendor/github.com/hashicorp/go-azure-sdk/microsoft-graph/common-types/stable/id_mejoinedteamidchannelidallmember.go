package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdChannelIdAllMemberId{}

// MeJoinedTeamIdChannelIdAllMemberId is a struct representing the Resource ID for a Me Joined Team Id Channel Id All Member
type MeJoinedTeamIdChannelIdAllMemberId struct {
	TeamId               string
	ChannelId            string
	ConversationMemberId string
}

// NewMeJoinedTeamIdChannelIdAllMemberID returns a new MeJoinedTeamIdChannelIdAllMemberId struct
func NewMeJoinedTeamIdChannelIdAllMemberID(teamId string, channelId string, conversationMemberId string) MeJoinedTeamIdChannelIdAllMemberId {
	return MeJoinedTeamIdChannelIdAllMemberId{
		TeamId:               teamId,
		ChannelId:            channelId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseMeJoinedTeamIdChannelIdAllMemberID parses 'input' into a MeJoinedTeamIdChannelIdAllMemberId
func ParseMeJoinedTeamIdChannelIdAllMemberID(input string) (*MeJoinedTeamIdChannelIdAllMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdChannelIdAllMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdChannelIdAllMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdChannelIdAllMemberIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdChannelIdAllMemberId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdChannelIdAllMemberIDInsensitively(input string) (*MeJoinedTeamIdChannelIdAllMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdChannelIdAllMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdChannelIdAllMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdChannelIdAllMemberId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateMeJoinedTeamIdChannelIdAllMemberID checks that 'input' can be parsed as a Me Joined Team Id Channel Id All Member ID
func ValidateMeJoinedTeamIdChannelIdAllMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdChannelIdAllMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Channel Id All Member ID
func (id MeJoinedTeamIdChannelIdAllMemberId) ID() string {
	fmtString := "/me/joinedTeams/%s/channels/%s/allMembers/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChannelId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Channel Id All Member ID
func (id MeJoinedTeamIdChannelIdAllMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("allMembers", "allMembers", "allMembers"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Channel Id All Member ID
func (id MeJoinedTeamIdChannelIdAllMemberId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Me Joined Team Id Channel Id All Member (%s)", strings.Join(components, "\n"))
}
