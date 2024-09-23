package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdScheduleOpenShiftId{}

// UserIdJoinedTeamIdScheduleOpenShiftId is a struct representing the Resource ID for a User Id Joined Team Id Schedule Open Shift
type UserIdJoinedTeamIdScheduleOpenShiftId struct {
	UserId      string
	TeamId      string
	OpenShiftId string
}

// NewUserIdJoinedTeamIdScheduleOpenShiftID returns a new UserIdJoinedTeamIdScheduleOpenShiftId struct
func NewUserIdJoinedTeamIdScheduleOpenShiftID(userId string, teamId string, openShiftId string) UserIdJoinedTeamIdScheduleOpenShiftId {
	return UserIdJoinedTeamIdScheduleOpenShiftId{
		UserId:      userId,
		TeamId:      teamId,
		OpenShiftId: openShiftId,
	}
}

// ParseUserIdJoinedTeamIdScheduleOpenShiftID parses 'input' into a UserIdJoinedTeamIdScheduleOpenShiftId
func ParseUserIdJoinedTeamIdScheduleOpenShiftID(input string) (*UserIdJoinedTeamIdScheduleOpenShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdScheduleOpenShiftId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdScheduleOpenShiftId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdScheduleOpenShiftIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdScheduleOpenShiftId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdScheduleOpenShiftIDInsensitively(input string) (*UserIdJoinedTeamIdScheduleOpenShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdScheduleOpenShiftId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdScheduleOpenShiftId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdScheduleOpenShiftId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.OpenShiftId, ok = input.Parsed["openShiftId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "openShiftId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdScheduleOpenShiftID checks that 'input' can be parsed as a User Id Joined Team Id Schedule Open Shift ID
func ValidateUserIdJoinedTeamIdScheduleOpenShiftID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdScheduleOpenShiftID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Schedule Open Shift ID
func (id UserIdJoinedTeamIdScheduleOpenShiftId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/schedule/openShifts/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.OpenShiftId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Schedule Open Shift ID
func (id UserIdJoinedTeamIdScheduleOpenShiftId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("openShifts", "openShifts", "openShifts"),
		resourceids.UserSpecifiedSegment("openShiftId", "openShiftId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Schedule Open Shift ID
func (id UserIdJoinedTeamIdScheduleOpenShiftId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Open Shift: %q", id.OpenShiftId),
	}
	return fmt.Sprintf("User Id Joined Team Id Schedule Open Shift (%s)", strings.Join(components, "\n"))
}
