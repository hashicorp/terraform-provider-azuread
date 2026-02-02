package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdIncomingChannelId{}

// MeJoinedTeamIdIncomingChannelId is a struct representing the Resource ID for a Me Joined Team Id Incoming Channel
type MeJoinedTeamIdIncomingChannelId struct {
	TeamId    string
	ChannelId string
}

// NewMeJoinedTeamIdIncomingChannelID returns a new MeJoinedTeamIdIncomingChannelId struct
func NewMeJoinedTeamIdIncomingChannelID(teamId string, channelId string) MeJoinedTeamIdIncomingChannelId {
	return MeJoinedTeamIdIncomingChannelId{
		TeamId:    teamId,
		ChannelId: channelId,
	}
}

// ParseMeJoinedTeamIdIncomingChannelID parses 'input' into a MeJoinedTeamIdIncomingChannelId
func ParseMeJoinedTeamIdIncomingChannelID(input string) (*MeJoinedTeamIdIncomingChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdIncomingChannelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdIncomingChannelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdIncomingChannelIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdIncomingChannelId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdIncomingChannelIDInsensitively(input string) (*MeJoinedTeamIdIncomingChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdIncomingChannelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdIncomingChannelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdIncomingChannelId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.ChannelId, ok = input.Parsed["channelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "channelId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdIncomingChannelID checks that 'input' can be parsed as a Me Joined Team Id Incoming Channel ID
func ValidateMeJoinedTeamIdIncomingChannelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdIncomingChannelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Incoming Channel ID
func (id MeJoinedTeamIdIncomingChannelId) ID() string {
	fmtString := "/me/joinedTeams/%s/incomingChannels/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChannelId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Incoming Channel ID
func (id MeJoinedTeamIdIncomingChannelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("incomingChannels", "incomingChannels", "incomingChannels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Incoming Channel ID
func (id MeJoinedTeamIdIncomingChannelId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
	}
	return fmt.Sprintf("Me Joined Team Id Incoming Channel (%s)", strings.Join(components, "\n"))
}
