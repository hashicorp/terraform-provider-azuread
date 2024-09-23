package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdScheduleShiftId{}

// MeJoinedTeamIdScheduleShiftId is a struct representing the Resource ID for a Me Joined Team Id Schedule Shift
type MeJoinedTeamIdScheduleShiftId struct {
	TeamId  string
	ShiftId string
}

// NewMeJoinedTeamIdScheduleShiftID returns a new MeJoinedTeamIdScheduleShiftId struct
func NewMeJoinedTeamIdScheduleShiftID(teamId string, shiftId string) MeJoinedTeamIdScheduleShiftId {
	return MeJoinedTeamIdScheduleShiftId{
		TeamId:  teamId,
		ShiftId: shiftId,
	}
}

// ParseMeJoinedTeamIdScheduleShiftID parses 'input' into a MeJoinedTeamIdScheduleShiftId
func ParseMeJoinedTeamIdScheduleShiftID(input string) (*MeJoinedTeamIdScheduleShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdScheduleShiftId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdScheduleShiftId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdScheduleShiftIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdScheduleShiftId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdScheduleShiftIDInsensitively(input string) (*MeJoinedTeamIdScheduleShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdScheduleShiftId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdScheduleShiftId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdScheduleShiftId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.ShiftId, ok = input.Parsed["shiftId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "shiftId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdScheduleShiftID checks that 'input' can be parsed as a Me Joined Team Id Schedule Shift ID
func ValidateMeJoinedTeamIdScheduleShiftID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdScheduleShiftID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Schedule Shift ID
func (id MeJoinedTeamIdScheduleShiftId) ID() string {
	fmtString := "/me/joinedTeams/%s/schedule/shifts/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ShiftId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Schedule Shift ID
func (id MeJoinedTeamIdScheduleShiftId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("shifts", "shifts", "shifts"),
		resourceids.UserSpecifiedSegment("shiftId", "shiftId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Schedule Shift ID
func (id MeJoinedTeamIdScheduleShiftId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Shift: %q", id.ShiftId),
	}
	return fmt.Sprintf("Me Joined Team Id Schedule Shift (%s)", strings.Join(components, "\n"))
}
