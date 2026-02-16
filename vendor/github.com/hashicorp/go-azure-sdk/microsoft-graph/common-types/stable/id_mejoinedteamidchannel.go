package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdChannelId{}

// MeJoinedTeamIdChannelId is a struct representing the Resource ID for a Me Joined Team Id Channel
type MeJoinedTeamIdChannelId struct {
	TeamId    string
	ChannelId string
}

// NewMeJoinedTeamIdChannelID returns a new MeJoinedTeamIdChannelId struct
func NewMeJoinedTeamIdChannelID(teamId string, channelId string) MeJoinedTeamIdChannelId {
	return MeJoinedTeamIdChannelId{
		TeamId:    teamId,
		ChannelId: channelId,
	}
}

// ParseMeJoinedTeamIdChannelID parses 'input' into a MeJoinedTeamIdChannelId
func ParseMeJoinedTeamIdChannelID(input string) (*MeJoinedTeamIdChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdChannelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdChannelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdChannelIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdChannelId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdChannelIDInsensitively(input string) (*MeJoinedTeamIdChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdChannelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdChannelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdChannelId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.ChannelId, ok = input.Parsed["channelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "channelId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdChannelID checks that 'input' can be parsed as a Me Joined Team Id Channel ID
func ValidateMeJoinedTeamIdChannelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdChannelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Channel ID
func (id MeJoinedTeamIdChannelId) ID() string {
	fmtString := "/me/joinedTeams/%s/channels/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChannelId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Channel ID
func (id MeJoinedTeamIdChannelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Channel ID
func (id MeJoinedTeamIdChannelId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
	}
	return fmt.Sprintf("Me Joined Team Id Channel (%s)", strings.Join(components, "\n"))
}
