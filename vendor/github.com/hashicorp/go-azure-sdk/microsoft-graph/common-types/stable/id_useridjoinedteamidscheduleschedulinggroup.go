package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdScheduleSchedulingGroupId{}

// UserIdJoinedTeamIdScheduleSchedulingGroupId is a struct representing the Resource ID for a User Id Joined Team Id Schedule Scheduling Group
type UserIdJoinedTeamIdScheduleSchedulingGroupId struct {
	UserId            string
	TeamId            string
	SchedulingGroupId string
}

// NewUserIdJoinedTeamIdScheduleSchedulingGroupID returns a new UserIdJoinedTeamIdScheduleSchedulingGroupId struct
func NewUserIdJoinedTeamIdScheduleSchedulingGroupID(userId string, teamId string, schedulingGroupId string) UserIdJoinedTeamIdScheduleSchedulingGroupId {
	return UserIdJoinedTeamIdScheduleSchedulingGroupId{
		UserId:            userId,
		TeamId:            teamId,
		SchedulingGroupId: schedulingGroupId,
	}
}

// ParseUserIdJoinedTeamIdScheduleSchedulingGroupID parses 'input' into a UserIdJoinedTeamIdScheduleSchedulingGroupId
func ParseUserIdJoinedTeamIdScheduleSchedulingGroupID(input string) (*UserIdJoinedTeamIdScheduleSchedulingGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdScheduleSchedulingGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdScheduleSchedulingGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdScheduleSchedulingGroupIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdScheduleSchedulingGroupId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdScheduleSchedulingGroupIDInsensitively(input string) (*UserIdJoinedTeamIdScheduleSchedulingGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdScheduleSchedulingGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdScheduleSchedulingGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdScheduleSchedulingGroupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.SchedulingGroupId, ok = input.Parsed["schedulingGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "schedulingGroupId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdScheduleSchedulingGroupID checks that 'input' can be parsed as a User Id Joined Team Id Schedule Scheduling Group ID
func ValidateUserIdJoinedTeamIdScheduleSchedulingGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdScheduleSchedulingGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Schedule Scheduling Group ID
func (id UserIdJoinedTeamIdScheduleSchedulingGroupId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/schedule/schedulingGroups/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.SchedulingGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Schedule Scheduling Group ID
func (id UserIdJoinedTeamIdScheduleSchedulingGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("schedulingGroups", "schedulingGroups", "schedulingGroups"),
		resourceids.UserSpecifiedSegment("schedulingGroupId", "schedulingGroupId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Schedule Scheduling Group ID
func (id UserIdJoinedTeamIdScheduleSchedulingGroupId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Scheduling Group: %q", id.SchedulingGroupId),
	}
	return fmt.Sprintf("User Id Joined Team Id Schedule Scheduling Group (%s)", strings.Join(components, "\n"))
}
