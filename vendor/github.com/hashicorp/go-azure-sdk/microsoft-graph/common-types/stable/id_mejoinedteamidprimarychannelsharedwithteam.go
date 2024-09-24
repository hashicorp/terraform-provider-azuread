package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdPrimaryChannelSharedWithTeamId{}

// MeJoinedTeamIdPrimaryChannelSharedWithTeamId is a struct representing the Resource ID for a Me Joined Team Id Primary Channel Shared With Team
type MeJoinedTeamIdPrimaryChannelSharedWithTeamId struct {
	TeamId                      string
	SharedWithChannelTeamInfoId string
}

// NewMeJoinedTeamIdPrimaryChannelSharedWithTeamID returns a new MeJoinedTeamIdPrimaryChannelSharedWithTeamId struct
func NewMeJoinedTeamIdPrimaryChannelSharedWithTeamID(teamId string, sharedWithChannelTeamInfoId string) MeJoinedTeamIdPrimaryChannelSharedWithTeamId {
	return MeJoinedTeamIdPrimaryChannelSharedWithTeamId{
		TeamId:                      teamId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
	}
}

// ParseMeJoinedTeamIdPrimaryChannelSharedWithTeamID parses 'input' into a MeJoinedTeamIdPrimaryChannelSharedWithTeamId
func ParseMeJoinedTeamIdPrimaryChannelSharedWithTeamID(input string) (*MeJoinedTeamIdPrimaryChannelSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdPrimaryChannelSharedWithTeamId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdPrimaryChannelSharedWithTeamId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdPrimaryChannelSharedWithTeamIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdPrimaryChannelSharedWithTeamId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdPrimaryChannelSharedWithTeamIDInsensitively(input string) (*MeJoinedTeamIdPrimaryChannelSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdPrimaryChannelSharedWithTeamId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdPrimaryChannelSharedWithTeamId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdPrimaryChannelSharedWithTeamId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.SharedWithChannelTeamInfoId, ok = input.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdPrimaryChannelSharedWithTeamID checks that 'input' can be parsed as a Me Joined Team Id Primary Channel Shared With Team ID
func ValidateMeJoinedTeamIdPrimaryChannelSharedWithTeamID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdPrimaryChannelSharedWithTeamID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Primary Channel Shared With Team ID
func (id MeJoinedTeamIdPrimaryChannelSharedWithTeamId) ID() string {
	fmtString := "/me/joinedTeams/%s/primaryChannel/sharedWithTeams/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.SharedWithChannelTeamInfoId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Primary Channel Shared With Team ID
func (id MeJoinedTeamIdPrimaryChannelSharedWithTeamId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("sharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Primary Channel Shared With Team ID
func (id MeJoinedTeamIdPrimaryChannelSharedWithTeamId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
	}
	return fmt.Sprintf("Me Joined Team Id Primary Channel Shared With Team (%s)", strings.Join(components, "\n"))
}
