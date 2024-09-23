package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdChannelIdSharedWithTeamId{}

// MeJoinedTeamIdChannelIdSharedWithTeamId is a struct representing the Resource ID for a Me Joined Team Id Channel Id Shared With Team
type MeJoinedTeamIdChannelIdSharedWithTeamId struct {
	TeamId                      string
	ChannelId                   string
	SharedWithChannelTeamInfoId string
}

// NewMeJoinedTeamIdChannelIdSharedWithTeamID returns a new MeJoinedTeamIdChannelIdSharedWithTeamId struct
func NewMeJoinedTeamIdChannelIdSharedWithTeamID(teamId string, channelId string, sharedWithChannelTeamInfoId string) MeJoinedTeamIdChannelIdSharedWithTeamId {
	return MeJoinedTeamIdChannelIdSharedWithTeamId{
		TeamId:                      teamId,
		ChannelId:                   channelId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
	}
}

// ParseMeJoinedTeamIdChannelIdSharedWithTeamID parses 'input' into a MeJoinedTeamIdChannelIdSharedWithTeamId
func ParseMeJoinedTeamIdChannelIdSharedWithTeamID(input string) (*MeJoinedTeamIdChannelIdSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdChannelIdSharedWithTeamId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdChannelIdSharedWithTeamId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdChannelIdSharedWithTeamIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdChannelIdSharedWithTeamId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdChannelIdSharedWithTeamIDInsensitively(input string) (*MeJoinedTeamIdChannelIdSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdChannelIdSharedWithTeamId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdChannelIdSharedWithTeamId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdChannelIdSharedWithTeamId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.ChannelId, ok = input.Parsed["channelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "channelId", input)
	}

	if id.SharedWithChannelTeamInfoId, ok = input.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdChannelIdSharedWithTeamID checks that 'input' can be parsed as a Me Joined Team Id Channel Id Shared With Team ID
func ValidateMeJoinedTeamIdChannelIdSharedWithTeamID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdChannelIdSharedWithTeamID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Channel Id Shared With Team ID
func (id MeJoinedTeamIdChannelIdSharedWithTeamId) ID() string {
	fmtString := "/me/joinedTeams/%s/channels/%s/sharedWithTeams/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChannelId, id.SharedWithChannelTeamInfoId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Channel Id Shared With Team ID
func (id MeJoinedTeamIdChannelIdSharedWithTeamId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("sharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Channel Id Shared With Team ID
func (id MeJoinedTeamIdChannelIdSharedWithTeamId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
	}
	return fmt.Sprintf("Me Joined Team Id Channel Id Shared With Team (%s)", strings.Join(components, "\n"))
}
