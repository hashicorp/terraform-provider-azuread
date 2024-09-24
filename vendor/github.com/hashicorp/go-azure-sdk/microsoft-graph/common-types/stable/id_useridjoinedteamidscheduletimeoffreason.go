package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdScheduleTimeOffReasonId{}

// UserIdJoinedTeamIdScheduleTimeOffReasonId is a struct representing the Resource ID for a User Id Joined Team Id Schedule Time Off Reason
type UserIdJoinedTeamIdScheduleTimeOffReasonId struct {
	UserId          string
	TeamId          string
	TimeOffReasonId string
}

// NewUserIdJoinedTeamIdScheduleTimeOffReasonID returns a new UserIdJoinedTeamIdScheduleTimeOffReasonId struct
func NewUserIdJoinedTeamIdScheduleTimeOffReasonID(userId string, teamId string, timeOffReasonId string) UserIdJoinedTeamIdScheduleTimeOffReasonId {
	return UserIdJoinedTeamIdScheduleTimeOffReasonId{
		UserId:          userId,
		TeamId:          teamId,
		TimeOffReasonId: timeOffReasonId,
	}
}

// ParseUserIdJoinedTeamIdScheduleTimeOffReasonID parses 'input' into a UserIdJoinedTeamIdScheduleTimeOffReasonId
func ParseUserIdJoinedTeamIdScheduleTimeOffReasonID(input string) (*UserIdJoinedTeamIdScheduleTimeOffReasonId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdScheduleTimeOffReasonId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdScheduleTimeOffReasonId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdScheduleTimeOffReasonIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdScheduleTimeOffReasonId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdScheduleTimeOffReasonIDInsensitively(input string) (*UserIdJoinedTeamIdScheduleTimeOffReasonId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdScheduleTimeOffReasonId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdScheduleTimeOffReasonId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdScheduleTimeOffReasonId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.TimeOffReasonId, ok = input.Parsed["timeOffReasonId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "timeOffReasonId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdScheduleTimeOffReasonID checks that 'input' can be parsed as a User Id Joined Team Id Schedule Time Off Reason ID
func ValidateUserIdJoinedTeamIdScheduleTimeOffReasonID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdScheduleTimeOffReasonID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Schedule Time Off Reason ID
func (id UserIdJoinedTeamIdScheduleTimeOffReasonId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/schedule/timeOffReasons/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.TimeOffReasonId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Schedule Time Off Reason ID
func (id UserIdJoinedTeamIdScheduleTimeOffReasonId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("timeOffReasons", "timeOffReasons", "timeOffReasons"),
		resourceids.UserSpecifiedSegment("timeOffReasonId", "timeOffReasonId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Schedule Time Off Reason ID
func (id UserIdJoinedTeamIdScheduleTimeOffReasonId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Time Off Reason: %q", id.TimeOffReasonId),
	}
	return fmt.Sprintf("User Id Joined Team Id Schedule Time Off Reason (%s)", strings.Join(components, "\n"))
}
