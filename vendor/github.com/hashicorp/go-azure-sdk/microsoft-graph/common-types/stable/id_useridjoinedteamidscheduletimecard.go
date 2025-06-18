package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdScheduleTimeCardId{}

// UserIdJoinedTeamIdScheduleTimeCardId is a struct representing the Resource ID for a User Id Joined Team Id Schedule Time Card
type UserIdJoinedTeamIdScheduleTimeCardId struct {
	UserId     string
	TeamId     string
	TimeCardId string
}

// NewUserIdJoinedTeamIdScheduleTimeCardID returns a new UserIdJoinedTeamIdScheduleTimeCardId struct
func NewUserIdJoinedTeamIdScheduleTimeCardID(userId string, teamId string, timeCardId string) UserIdJoinedTeamIdScheduleTimeCardId {
	return UserIdJoinedTeamIdScheduleTimeCardId{
		UserId:     userId,
		TeamId:     teamId,
		TimeCardId: timeCardId,
	}
}

// ParseUserIdJoinedTeamIdScheduleTimeCardID parses 'input' into a UserIdJoinedTeamIdScheduleTimeCardId
func ParseUserIdJoinedTeamIdScheduleTimeCardID(input string) (*UserIdJoinedTeamIdScheduleTimeCardId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdScheduleTimeCardId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdScheduleTimeCardId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdScheduleTimeCardIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdScheduleTimeCardId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdScheduleTimeCardIDInsensitively(input string) (*UserIdJoinedTeamIdScheduleTimeCardId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdScheduleTimeCardId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdScheduleTimeCardId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdScheduleTimeCardId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.TimeCardId, ok = input.Parsed["timeCardId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "timeCardId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdScheduleTimeCardID checks that 'input' can be parsed as a User Id Joined Team Id Schedule Time Card ID
func ValidateUserIdJoinedTeamIdScheduleTimeCardID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdScheduleTimeCardID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Schedule Time Card ID
func (id UserIdJoinedTeamIdScheduleTimeCardId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/schedule/timeCards/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.TimeCardId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Schedule Time Card ID
func (id UserIdJoinedTeamIdScheduleTimeCardId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("timeCards", "timeCards", "timeCards"),
		resourceids.UserSpecifiedSegment("timeCardId", "timeCardId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Schedule Time Card ID
func (id UserIdJoinedTeamIdScheduleTimeCardId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Time Card: %q", id.TimeCardId),
	}
	return fmt.Sprintf("User Id Joined Team Id Schedule Time Card (%s)", strings.Join(components, "\n"))
}
