package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdScheduleTimesOffId{}

// UserIdJoinedTeamIdScheduleTimesOffId is a struct representing the Resource ID for a User Id Joined Team Id Schedule Times Off
type UserIdJoinedTeamIdScheduleTimesOffId struct {
	UserId    string
	TeamId    string
	TimeOffId string
}

// NewUserIdJoinedTeamIdScheduleTimesOffID returns a new UserIdJoinedTeamIdScheduleTimesOffId struct
func NewUserIdJoinedTeamIdScheduleTimesOffID(userId string, teamId string, timeOffId string) UserIdJoinedTeamIdScheduleTimesOffId {
	return UserIdJoinedTeamIdScheduleTimesOffId{
		UserId:    userId,
		TeamId:    teamId,
		TimeOffId: timeOffId,
	}
}

// ParseUserIdJoinedTeamIdScheduleTimesOffID parses 'input' into a UserIdJoinedTeamIdScheduleTimesOffId
func ParseUserIdJoinedTeamIdScheduleTimesOffID(input string) (*UserIdJoinedTeamIdScheduleTimesOffId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdScheduleTimesOffId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdScheduleTimesOffId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdScheduleTimesOffIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdScheduleTimesOffId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdScheduleTimesOffIDInsensitively(input string) (*UserIdJoinedTeamIdScheduleTimesOffId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdScheduleTimesOffId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdScheduleTimesOffId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdScheduleTimesOffId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.TimeOffId, ok = input.Parsed["timeOffId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "timeOffId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdScheduleTimesOffID checks that 'input' can be parsed as a User Id Joined Team Id Schedule Times Off ID
func ValidateUserIdJoinedTeamIdScheduleTimesOffID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdScheduleTimesOffID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Schedule Times Off ID
func (id UserIdJoinedTeamIdScheduleTimesOffId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/schedule/timesOff/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.TimeOffId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Schedule Times Off ID
func (id UserIdJoinedTeamIdScheduleTimesOffId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("timesOff", "timesOff", "timesOff"),
		resourceids.UserSpecifiedSegment("timeOffId", "timeOffId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Schedule Times Off ID
func (id UserIdJoinedTeamIdScheduleTimesOffId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Time Off: %q", id.TimeOffId),
	}
	return fmt.Sprintf("User Id Joined Team Id Schedule Times Off (%s)", strings.Join(components, "\n"))
}
