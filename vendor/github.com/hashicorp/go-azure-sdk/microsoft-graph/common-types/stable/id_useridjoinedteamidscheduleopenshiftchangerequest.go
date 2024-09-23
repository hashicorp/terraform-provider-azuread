package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId{}

// UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId is a struct representing the Resource ID for a User Id Joined Team Id Schedule Open Shift Change Request
type UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId struct {
	UserId                   string
	TeamId                   string
	OpenShiftChangeRequestId string
}

// NewUserIdJoinedTeamIdScheduleOpenShiftChangeRequestID returns a new UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId struct
func NewUserIdJoinedTeamIdScheduleOpenShiftChangeRequestID(userId string, teamId string, openShiftChangeRequestId string) UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId {
	return UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId{
		UserId:                   userId,
		TeamId:                   teamId,
		OpenShiftChangeRequestId: openShiftChangeRequestId,
	}
}

// ParseUserIdJoinedTeamIdScheduleOpenShiftChangeRequestID parses 'input' into a UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId
func ParseUserIdJoinedTeamIdScheduleOpenShiftChangeRequestID(input string) (*UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdScheduleOpenShiftChangeRequestIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdScheduleOpenShiftChangeRequestIDInsensitively(input string) (*UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.OpenShiftChangeRequestId, ok = input.Parsed["openShiftChangeRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "openShiftChangeRequestId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdScheduleOpenShiftChangeRequestID checks that 'input' can be parsed as a User Id Joined Team Id Schedule Open Shift Change Request ID
func ValidateUserIdJoinedTeamIdScheduleOpenShiftChangeRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdScheduleOpenShiftChangeRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Schedule Open Shift Change Request ID
func (id UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/schedule/openShiftChangeRequests/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.OpenShiftChangeRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Schedule Open Shift Change Request ID
func (id UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("openShiftChangeRequests", "openShiftChangeRequests", "openShiftChangeRequests"),
		resourceids.UserSpecifiedSegment("openShiftChangeRequestId", "openShiftChangeRequestId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Schedule Open Shift Change Request ID
func (id UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Open Shift Change Request: %q", id.OpenShiftChangeRequestId),
	}
	return fmt.Sprintf("User Id Joined Team Id Schedule Open Shift Change Request (%s)", strings.Join(components, "\n"))
}
