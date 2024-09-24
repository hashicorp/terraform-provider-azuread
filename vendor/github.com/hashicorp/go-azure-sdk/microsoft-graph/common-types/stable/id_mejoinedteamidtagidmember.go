package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdTagIdMemberId{}

// MeJoinedTeamIdTagIdMemberId is a struct representing the Resource ID for a Me Joined Team Id Tag Id Member
type MeJoinedTeamIdTagIdMemberId struct {
	TeamId              string
	TeamworkTagId       string
	TeamworkTagMemberId string
}

// NewMeJoinedTeamIdTagIdMemberID returns a new MeJoinedTeamIdTagIdMemberId struct
func NewMeJoinedTeamIdTagIdMemberID(teamId string, teamworkTagId string, teamworkTagMemberId string) MeJoinedTeamIdTagIdMemberId {
	return MeJoinedTeamIdTagIdMemberId{
		TeamId:              teamId,
		TeamworkTagId:       teamworkTagId,
		TeamworkTagMemberId: teamworkTagMemberId,
	}
}

// ParseMeJoinedTeamIdTagIdMemberID parses 'input' into a MeJoinedTeamIdTagIdMemberId
func ParseMeJoinedTeamIdTagIdMemberID(input string) (*MeJoinedTeamIdTagIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdTagIdMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdTagIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdTagIdMemberIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdTagIdMemberId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdTagIdMemberIDInsensitively(input string) (*MeJoinedTeamIdTagIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdTagIdMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdTagIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdTagIdMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.TeamworkTagId, ok = input.Parsed["teamworkTagId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagId", input)
	}

	if id.TeamworkTagMemberId, ok = input.Parsed["teamworkTagMemberId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagMemberId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdTagIdMemberID checks that 'input' can be parsed as a Me Joined Team Id Tag Id Member ID
func ValidateMeJoinedTeamIdTagIdMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdTagIdMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Tag Id Member ID
func (id MeJoinedTeamIdTagIdMemberId) ID() string {
	fmtString := "/me/joinedTeams/%s/tags/%s/members/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.TeamworkTagId, id.TeamworkTagMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Tag Id Member ID
func (id MeJoinedTeamIdTagIdMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("tags", "tags", "tags"),
		resourceids.UserSpecifiedSegment("teamworkTagId", "teamworkTagId"),
		resourceids.StaticSegment("members", "members", "members"),
		resourceids.UserSpecifiedSegment("teamworkTagMemberId", "teamworkTagMemberId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Tag Id Member ID
func (id MeJoinedTeamIdTagIdMemberId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Teamwork Tag: %q", id.TeamworkTagId),
		fmt.Sprintf("Teamwork Tag Member: %q", id.TeamworkTagMemberId),
	}
	return fmt.Sprintf("Me Joined Team Id Tag Id Member (%s)", strings.Join(components, "\n"))
}
