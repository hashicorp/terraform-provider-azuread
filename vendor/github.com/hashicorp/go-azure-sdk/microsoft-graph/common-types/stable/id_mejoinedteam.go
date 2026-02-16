package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamId{}

// MeJoinedTeamId is a struct representing the Resource ID for a Me Joined Team
type MeJoinedTeamId struct {
	TeamId string
}

// NewMeJoinedTeamID returns a new MeJoinedTeamId struct
func NewMeJoinedTeamID(teamId string) MeJoinedTeamId {
	return MeJoinedTeamId{
		TeamId: teamId,
	}
}

// ParseMeJoinedTeamID parses 'input' into a MeJoinedTeamId
func ParseMeJoinedTeamID(input string) (*MeJoinedTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIDInsensitively(input string) (*MeJoinedTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	return nil
}

// ValidateMeJoinedTeamID checks that 'input' can be parsed as a Me Joined Team ID
func ValidateMeJoinedTeamID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team ID
func (id MeJoinedTeamId) ID() string {
	fmtString := "/me/joinedTeams/%s"
	return fmt.Sprintf(fmtString, id.TeamId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team ID
func (id MeJoinedTeamId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
	}
}

// String returns a human-readable description of this Me Joined Team ID
func (id MeJoinedTeamId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
	}
	return fmt.Sprintf("Me Joined Team (%s)", strings.Join(components, "\n"))
}
