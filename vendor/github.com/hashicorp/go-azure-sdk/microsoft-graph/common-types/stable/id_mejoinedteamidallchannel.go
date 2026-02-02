package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdAllChannelId{}

// MeJoinedTeamIdAllChannelId is a struct representing the Resource ID for a Me Joined Team Id All Channel
type MeJoinedTeamIdAllChannelId struct {
	TeamId    string
	ChannelId string
}

// NewMeJoinedTeamIdAllChannelID returns a new MeJoinedTeamIdAllChannelId struct
func NewMeJoinedTeamIdAllChannelID(teamId string, channelId string) MeJoinedTeamIdAllChannelId {
	return MeJoinedTeamIdAllChannelId{
		TeamId:    teamId,
		ChannelId: channelId,
	}
}

// ParseMeJoinedTeamIdAllChannelID parses 'input' into a MeJoinedTeamIdAllChannelId
func ParseMeJoinedTeamIdAllChannelID(input string) (*MeJoinedTeamIdAllChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdAllChannelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdAllChannelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdAllChannelIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdAllChannelId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdAllChannelIDInsensitively(input string) (*MeJoinedTeamIdAllChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdAllChannelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdAllChannelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdAllChannelId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.ChannelId, ok = input.Parsed["channelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "channelId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdAllChannelID checks that 'input' can be parsed as a Me Joined Team Id All Channel ID
func ValidateMeJoinedTeamIdAllChannelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdAllChannelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id All Channel ID
func (id MeJoinedTeamIdAllChannelId) ID() string {
	fmtString := "/me/joinedTeams/%s/allChannels/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChannelId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id All Channel ID
func (id MeJoinedTeamIdAllChannelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("allChannels", "allChannels", "allChannels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id All Channel ID
func (id MeJoinedTeamIdAllChannelId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
	}
	return fmt.Sprintf("Me Joined Team Id All Channel (%s)", strings.Join(components, "\n"))
}
