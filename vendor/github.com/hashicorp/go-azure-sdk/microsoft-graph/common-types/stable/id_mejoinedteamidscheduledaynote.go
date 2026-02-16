package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdScheduleDayNoteId{}

// MeJoinedTeamIdScheduleDayNoteId is a struct representing the Resource ID for a Me Joined Team Id Schedule Day Note
type MeJoinedTeamIdScheduleDayNoteId struct {
	TeamId    string
	DayNoteId string
}

// NewMeJoinedTeamIdScheduleDayNoteID returns a new MeJoinedTeamIdScheduleDayNoteId struct
func NewMeJoinedTeamIdScheduleDayNoteID(teamId string, dayNoteId string) MeJoinedTeamIdScheduleDayNoteId {
	return MeJoinedTeamIdScheduleDayNoteId{
		TeamId:    teamId,
		DayNoteId: dayNoteId,
	}
}

// ParseMeJoinedTeamIdScheduleDayNoteID parses 'input' into a MeJoinedTeamIdScheduleDayNoteId
func ParseMeJoinedTeamIdScheduleDayNoteID(input string) (*MeJoinedTeamIdScheduleDayNoteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdScheduleDayNoteId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdScheduleDayNoteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdScheduleDayNoteIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdScheduleDayNoteId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdScheduleDayNoteIDInsensitively(input string) (*MeJoinedTeamIdScheduleDayNoteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdScheduleDayNoteId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdScheduleDayNoteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdScheduleDayNoteId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.DayNoteId, ok = input.Parsed["dayNoteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dayNoteId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdScheduleDayNoteID checks that 'input' can be parsed as a Me Joined Team Id Schedule Day Note ID
func ValidateMeJoinedTeamIdScheduleDayNoteID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdScheduleDayNoteID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Schedule Day Note ID
func (id MeJoinedTeamIdScheduleDayNoteId) ID() string {
	fmtString := "/me/joinedTeams/%s/schedule/dayNotes/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.DayNoteId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Schedule Day Note ID
func (id MeJoinedTeamIdScheduleDayNoteId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("dayNotes", "dayNotes", "dayNotes"),
		resourceids.UserSpecifiedSegment("dayNoteId", "dayNoteId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Schedule Day Note ID
func (id MeJoinedTeamIdScheduleDayNoteId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Day Note: %q", id.DayNoteId),
	}
	return fmt.Sprintf("Me Joined Team Id Schedule Day Note (%s)", strings.Join(components, "\n"))
}
