package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdChannelIdTabId{}

// MeJoinedTeamIdChannelIdTabId is a struct representing the Resource ID for a Me Joined Team Id Channel Id Tab
type MeJoinedTeamIdChannelIdTabId struct {
	TeamId     string
	ChannelId  string
	TeamsTabId string
}

// NewMeJoinedTeamIdChannelIdTabID returns a new MeJoinedTeamIdChannelIdTabId struct
func NewMeJoinedTeamIdChannelIdTabID(teamId string, channelId string, teamsTabId string) MeJoinedTeamIdChannelIdTabId {
	return MeJoinedTeamIdChannelIdTabId{
		TeamId:     teamId,
		ChannelId:  channelId,
		TeamsTabId: teamsTabId,
	}
}

// ParseMeJoinedTeamIdChannelIdTabID parses 'input' into a MeJoinedTeamIdChannelIdTabId
func ParseMeJoinedTeamIdChannelIdTabID(input string) (*MeJoinedTeamIdChannelIdTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdChannelIdTabId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdChannelIdTabId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdChannelIdTabIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdChannelIdTabId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdChannelIdTabIDInsensitively(input string) (*MeJoinedTeamIdChannelIdTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdChannelIdTabId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdChannelIdTabId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdChannelIdTabId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.ChannelId, ok = input.Parsed["channelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "channelId", input)
	}

	if id.TeamsTabId, ok = input.Parsed["teamsTabId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdChannelIdTabID checks that 'input' can be parsed as a Me Joined Team Id Channel Id Tab ID
func ValidateMeJoinedTeamIdChannelIdTabID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdChannelIdTabID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Channel Id Tab ID
func (id MeJoinedTeamIdChannelIdTabId) ID() string {
	fmtString := "/me/joinedTeams/%s/channels/%s/tabs/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChannelId, id.TeamsTabId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Channel Id Tab ID
func (id MeJoinedTeamIdChannelIdTabId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("tabs", "tabs", "tabs"),
		resourceids.UserSpecifiedSegment("teamsTabId", "teamsTabId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Channel Id Tab ID
func (id MeJoinedTeamIdChannelIdTabId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Teams Tab: %q", id.TeamsTabId),
	}
	return fmt.Sprintf("Me Joined Team Id Channel Id Tab (%s)", strings.Join(components, "\n"))
}
