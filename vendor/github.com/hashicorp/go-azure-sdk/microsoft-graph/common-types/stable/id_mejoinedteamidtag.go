package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdTagId{}

// MeJoinedTeamIdTagId is a struct representing the Resource ID for a Me Joined Team Id Tag
type MeJoinedTeamIdTagId struct {
	TeamId        string
	TeamworkTagId string
}

// NewMeJoinedTeamIdTagID returns a new MeJoinedTeamIdTagId struct
func NewMeJoinedTeamIdTagID(teamId string, teamworkTagId string) MeJoinedTeamIdTagId {
	return MeJoinedTeamIdTagId{
		TeamId:        teamId,
		TeamworkTagId: teamworkTagId,
	}
}

// ParseMeJoinedTeamIdTagID parses 'input' into a MeJoinedTeamIdTagId
func ParseMeJoinedTeamIdTagID(input string) (*MeJoinedTeamIdTagId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdTagId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdTagId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdTagIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdTagId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdTagIDInsensitively(input string) (*MeJoinedTeamIdTagId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdTagId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdTagId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdTagId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.TeamworkTagId, ok = input.Parsed["teamworkTagId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdTagID checks that 'input' can be parsed as a Me Joined Team Id Tag ID
func ValidateMeJoinedTeamIdTagID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdTagID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Tag ID
func (id MeJoinedTeamIdTagId) ID() string {
	fmtString := "/me/joinedTeams/%s/tags/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.TeamworkTagId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Tag ID
func (id MeJoinedTeamIdTagId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("tags", "tags", "tags"),
		resourceids.UserSpecifiedSegment("teamworkTagId", "teamworkTagId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Tag ID
func (id MeJoinedTeamIdTagId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Teamwork Tag: %q", id.TeamworkTagId),
	}
	return fmt.Sprintf("Me Joined Team Id Tag (%s)", strings.Join(components, "\n"))
}
