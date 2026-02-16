package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdScheduleTimeCardId{}

// MeJoinedTeamIdScheduleTimeCardId is a struct representing the Resource ID for a Me Joined Team Id Schedule Time Card
type MeJoinedTeamIdScheduleTimeCardId struct {
	TeamId     string
	TimeCardId string
}

// NewMeJoinedTeamIdScheduleTimeCardID returns a new MeJoinedTeamIdScheduleTimeCardId struct
func NewMeJoinedTeamIdScheduleTimeCardID(teamId string, timeCardId string) MeJoinedTeamIdScheduleTimeCardId {
	return MeJoinedTeamIdScheduleTimeCardId{
		TeamId:     teamId,
		TimeCardId: timeCardId,
	}
}

// ParseMeJoinedTeamIdScheduleTimeCardID parses 'input' into a MeJoinedTeamIdScheduleTimeCardId
func ParseMeJoinedTeamIdScheduleTimeCardID(input string) (*MeJoinedTeamIdScheduleTimeCardId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdScheduleTimeCardId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdScheduleTimeCardId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdScheduleTimeCardIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdScheduleTimeCardId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdScheduleTimeCardIDInsensitively(input string) (*MeJoinedTeamIdScheduleTimeCardId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdScheduleTimeCardId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdScheduleTimeCardId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdScheduleTimeCardId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.TimeCardId, ok = input.Parsed["timeCardId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "timeCardId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdScheduleTimeCardID checks that 'input' can be parsed as a Me Joined Team Id Schedule Time Card ID
func ValidateMeJoinedTeamIdScheduleTimeCardID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdScheduleTimeCardID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Schedule Time Card ID
func (id MeJoinedTeamIdScheduleTimeCardId) ID() string {
	fmtString := "/me/joinedTeams/%s/schedule/timeCards/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.TimeCardId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Schedule Time Card ID
func (id MeJoinedTeamIdScheduleTimeCardId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("timeCards", "timeCards", "timeCards"),
		resourceids.UserSpecifiedSegment("timeCardId", "timeCardId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Schedule Time Card ID
func (id MeJoinedTeamIdScheduleTimeCardId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Time Card: %q", id.TimeCardId),
	}
	return fmt.Sprintf("Me Joined Team Id Schedule Time Card (%s)", strings.Join(components, "\n"))
}
