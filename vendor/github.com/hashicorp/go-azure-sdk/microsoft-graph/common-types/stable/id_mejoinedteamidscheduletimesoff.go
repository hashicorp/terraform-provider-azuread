package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdScheduleTimesOffId{}

// MeJoinedTeamIdScheduleTimesOffId is a struct representing the Resource ID for a Me Joined Team Id Schedule Times Off
type MeJoinedTeamIdScheduleTimesOffId struct {
	TeamId    string
	TimeOffId string
}

// NewMeJoinedTeamIdScheduleTimesOffID returns a new MeJoinedTeamIdScheduleTimesOffId struct
func NewMeJoinedTeamIdScheduleTimesOffID(teamId string, timeOffId string) MeJoinedTeamIdScheduleTimesOffId {
	return MeJoinedTeamIdScheduleTimesOffId{
		TeamId:    teamId,
		TimeOffId: timeOffId,
	}
}

// ParseMeJoinedTeamIdScheduleTimesOffID parses 'input' into a MeJoinedTeamIdScheduleTimesOffId
func ParseMeJoinedTeamIdScheduleTimesOffID(input string) (*MeJoinedTeamIdScheduleTimesOffId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdScheduleTimesOffId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdScheduleTimesOffId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdScheduleTimesOffIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdScheduleTimesOffId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdScheduleTimesOffIDInsensitively(input string) (*MeJoinedTeamIdScheduleTimesOffId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdScheduleTimesOffId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdScheduleTimesOffId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdScheduleTimesOffId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.TimeOffId, ok = input.Parsed["timeOffId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "timeOffId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdScheduleTimesOffID checks that 'input' can be parsed as a Me Joined Team Id Schedule Times Off ID
func ValidateMeJoinedTeamIdScheduleTimesOffID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdScheduleTimesOffID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Schedule Times Off ID
func (id MeJoinedTeamIdScheduleTimesOffId) ID() string {
	fmtString := "/me/joinedTeams/%s/schedule/timesOff/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.TimeOffId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Schedule Times Off ID
func (id MeJoinedTeamIdScheduleTimesOffId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("timesOff", "timesOff", "timesOff"),
		resourceids.UserSpecifiedSegment("timeOffId", "timeOffId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Schedule Times Off ID
func (id MeJoinedTeamIdScheduleTimesOffId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Time Off: %q", id.TimeOffId),
	}
	return fmt.Sprintf("Me Joined Team Id Schedule Times Off (%s)", strings.Join(components, "\n"))
}
