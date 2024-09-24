package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdScheduleTimeOffReasonId{}

// MeJoinedTeamIdScheduleTimeOffReasonId is a struct representing the Resource ID for a Me Joined Team Id Schedule Time Off Reason
type MeJoinedTeamIdScheduleTimeOffReasonId struct {
	TeamId          string
	TimeOffReasonId string
}

// NewMeJoinedTeamIdScheduleTimeOffReasonID returns a new MeJoinedTeamIdScheduleTimeOffReasonId struct
func NewMeJoinedTeamIdScheduleTimeOffReasonID(teamId string, timeOffReasonId string) MeJoinedTeamIdScheduleTimeOffReasonId {
	return MeJoinedTeamIdScheduleTimeOffReasonId{
		TeamId:          teamId,
		TimeOffReasonId: timeOffReasonId,
	}
}

// ParseMeJoinedTeamIdScheduleTimeOffReasonID parses 'input' into a MeJoinedTeamIdScheduleTimeOffReasonId
func ParseMeJoinedTeamIdScheduleTimeOffReasonID(input string) (*MeJoinedTeamIdScheduleTimeOffReasonId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdScheduleTimeOffReasonId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdScheduleTimeOffReasonId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdScheduleTimeOffReasonIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdScheduleTimeOffReasonId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdScheduleTimeOffReasonIDInsensitively(input string) (*MeJoinedTeamIdScheduleTimeOffReasonId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdScheduleTimeOffReasonId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdScheduleTimeOffReasonId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdScheduleTimeOffReasonId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.TimeOffReasonId, ok = input.Parsed["timeOffReasonId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "timeOffReasonId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdScheduleTimeOffReasonID checks that 'input' can be parsed as a Me Joined Team Id Schedule Time Off Reason ID
func ValidateMeJoinedTeamIdScheduleTimeOffReasonID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdScheduleTimeOffReasonID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Schedule Time Off Reason ID
func (id MeJoinedTeamIdScheduleTimeOffReasonId) ID() string {
	fmtString := "/me/joinedTeams/%s/schedule/timeOffReasons/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.TimeOffReasonId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Schedule Time Off Reason ID
func (id MeJoinedTeamIdScheduleTimeOffReasonId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("timeOffReasons", "timeOffReasons", "timeOffReasons"),
		resourceids.UserSpecifiedSegment("timeOffReasonId", "timeOffReasonId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Schedule Time Off Reason ID
func (id MeJoinedTeamIdScheduleTimeOffReasonId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Time Off Reason: %q", id.TimeOffReasonId),
	}
	return fmt.Sprintf("Me Joined Team Id Schedule Time Off Reason (%s)", strings.Join(components, "\n"))
}
