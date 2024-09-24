package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdScheduleTimeOffRequestId{}

// UserIdJoinedTeamIdScheduleTimeOffRequestId is a struct representing the Resource ID for a User Id Joined Team Id Schedule Time Off Request
type UserIdJoinedTeamIdScheduleTimeOffRequestId struct {
	UserId           string
	TeamId           string
	TimeOffRequestId string
}

// NewUserIdJoinedTeamIdScheduleTimeOffRequestID returns a new UserIdJoinedTeamIdScheduleTimeOffRequestId struct
func NewUserIdJoinedTeamIdScheduleTimeOffRequestID(userId string, teamId string, timeOffRequestId string) UserIdJoinedTeamIdScheduleTimeOffRequestId {
	return UserIdJoinedTeamIdScheduleTimeOffRequestId{
		UserId:           userId,
		TeamId:           teamId,
		TimeOffRequestId: timeOffRequestId,
	}
}

// ParseUserIdJoinedTeamIdScheduleTimeOffRequestID parses 'input' into a UserIdJoinedTeamIdScheduleTimeOffRequestId
func ParseUserIdJoinedTeamIdScheduleTimeOffRequestID(input string) (*UserIdJoinedTeamIdScheduleTimeOffRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdScheduleTimeOffRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdScheduleTimeOffRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdScheduleTimeOffRequestIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdScheduleTimeOffRequestId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdScheduleTimeOffRequestIDInsensitively(input string) (*UserIdJoinedTeamIdScheduleTimeOffRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdScheduleTimeOffRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdScheduleTimeOffRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdScheduleTimeOffRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.TimeOffRequestId, ok = input.Parsed["timeOffRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "timeOffRequestId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdScheduleTimeOffRequestID checks that 'input' can be parsed as a User Id Joined Team Id Schedule Time Off Request ID
func ValidateUserIdJoinedTeamIdScheduleTimeOffRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdScheduleTimeOffRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Schedule Time Off Request ID
func (id UserIdJoinedTeamIdScheduleTimeOffRequestId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/schedule/timeOffRequests/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.TimeOffRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Schedule Time Off Request ID
func (id UserIdJoinedTeamIdScheduleTimeOffRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("timeOffRequests", "timeOffRequests", "timeOffRequests"),
		resourceids.UserSpecifiedSegment("timeOffRequestId", "timeOffRequestId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Schedule Time Off Request ID
func (id UserIdJoinedTeamIdScheduleTimeOffRequestId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Time Off Request: %q", id.TimeOffRequestId),
	}
	return fmt.Sprintf("User Id Joined Team Id Schedule Time Off Request (%s)", strings.Join(components, "\n"))
}
