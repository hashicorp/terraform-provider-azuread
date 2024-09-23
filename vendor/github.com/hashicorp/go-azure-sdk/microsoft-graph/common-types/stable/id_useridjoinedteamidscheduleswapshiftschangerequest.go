package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId{}

// UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId is a struct representing the Resource ID for a User Id Joined Team Id Schedule Swap Shifts Change Request
type UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId struct {
	UserId                    string
	TeamId                    string
	SwapShiftsChangeRequestId string
}

// NewUserIdJoinedTeamIdScheduleSwapShiftsChangeRequestID returns a new UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId struct
func NewUserIdJoinedTeamIdScheduleSwapShiftsChangeRequestID(userId string, teamId string, swapShiftsChangeRequestId string) UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId {
	return UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId{
		UserId:                    userId,
		TeamId:                    teamId,
		SwapShiftsChangeRequestId: swapShiftsChangeRequestId,
	}
}

// ParseUserIdJoinedTeamIdScheduleSwapShiftsChangeRequestID parses 'input' into a UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId
func ParseUserIdJoinedTeamIdScheduleSwapShiftsChangeRequestID(input string) (*UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdScheduleSwapShiftsChangeRequestIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdScheduleSwapShiftsChangeRequestIDInsensitively(input string) (*UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.SwapShiftsChangeRequestId, ok = input.Parsed["swapShiftsChangeRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "swapShiftsChangeRequestId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdScheduleSwapShiftsChangeRequestID checks that 'input' can be parsed as a User Id Joined Team Id Schedule Swap Shifts Change Request ID
func ValidateUserIdJoinedTeamIdScheduleSwapShiftsChangeRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdScheduleSwapShiftsChangeRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Schedule Swap Shifts Change Request ID
func (id UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/schedule/swapShiftsChangeRequests/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.SwapShiftsChangeRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Schedule Swap Shifts Change Request ID
func (id UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("swapShiftsChangeRequests", "swapShiftsChangeRequests", "swapShiftsChangeRequests"),
		resourceids.UserSpecifiedSegment("swapShiftsChangeRequestId", "swapShiftsChangeRequestId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Schedule Swap Shifts Change Request ID
func (id UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Swap Shifts Change Request: %q", id.SwapShiftsChangeRequestId),
	}
	return fmt.Sprintf("User Id Joined Team Id Schedule Swap Shifts Change Request (%s)", strings.Join(components, "\n"))
}
