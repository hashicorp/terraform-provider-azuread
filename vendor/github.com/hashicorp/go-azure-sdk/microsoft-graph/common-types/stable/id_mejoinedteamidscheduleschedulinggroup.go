package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdScheduleSchedulingGroupId{}

// MeJoinedTeamIdScheduleSchedulingGroupId is a struct representing the Resource ID for a Me Joined Team Id Schedule Scheduling Group
type MeJoinedTeamIdScheduleSchedulingGroupId struct {
	TeamId            string
	SchedulingGroupId string
}

// NewMeJoinedTeamIdScheduleSchedulingGroupID returns a new MeJoinedTeamIdScheduleSchedulingGroupId struct
func NewMeJoinedTeamIdScheduleSchedulingGroupID(teamId string, schedulingGroupId string) MeJoinedTeamIdScheduleSchedulingGroupId {
	return MeJoinedTeamIdScheduleSchedulingGroupId{
		TeamId:            teamId,
		SchedulingGroupId: schedulingGroupId,
	}
}

// ParseMeJoinedTeamIdScheduleSchedulingGroupID parses 'input' into a MeJoinedTeamIdScheduleSchedulingGroupId
func ParseMeJoinedTeamIdScheduleSchedulingGroupID(input string) (*MeJoinedTeamIdScheduleSchedulingGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdScheduleSchedulingGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdScheduleSchedulingGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdScheduleSchedulingGroupIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdScheduleSchedulingGroupId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdScheduleSchedulingGroupIDInsensitively(input string) (*MeJoinedTeamIdScheduleSchedulingGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdScheduleSchedulingGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdScheduleSchedulingGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdScheduleSchedulingGroupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.SchedulingGroupId, ok = input.Parsed["schedulingGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "schedulingGroupId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdScheduleSchedulingGroupID checks that 'input' can be parsed as a Me Joined Team Id Schedule Scheduling Group ID
func ValidateMeJoinedTeamIdScheduleSchedulingGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdScheduleSchedulingGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Schedule Scheduling Group ID
func (id MeJoinedTeamIdScheduleSchedulingGroupId) ID() string {
	fmtString := "/me/joinedTeams/%s/schedule/schedulingGroups/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.SchedulingGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Schedule Scheduling Group ID
func (id MeJoinedTeamIdScheduleSchedulingGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("schedulingGroups", "schedulingGroups", "schedulingGroups"),
		resourceids.UserSpecifiedSegment("schedulingGroupId", "schedulingGroupId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Schedule Scheduling Group ID
func (id MeJoinedTeamIdScheduleSchedulingGroupId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Scheduling Group: %q", id.SchedulingGroupId),
	}
	return fmt.Sprintf("Me Joined Team Id Schedule Scheduling Group (%s)", strings.Join(components, "\n"))
}
