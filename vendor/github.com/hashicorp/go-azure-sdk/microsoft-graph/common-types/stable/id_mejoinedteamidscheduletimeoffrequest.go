package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdScheduleTimeOffRequestId{}

// MeJoinedTeamIdScheduleTimeOffRequestId is a struct representing the Resource ID for a Me Joined Team Id Schedule Time Off Request
type MeJoinedTeamIdScheduleTimeOffRequestId struct {
	TeamId           string
	TimeOffRequestId string
}

// NewMeJoinedTeamIdScheduleTimeOffRequestID returns a new MeJoinedTeamIdScheduleTimeOffRequestId struct
func NewMeJoinedTeamIdScheduleTimeOffRequestID(teamId string, timeOffRequestId string) MeJoinedTeamIdScheduleTimeOffRequestId {
	return MeJoinedTeamIdScheduleTimeOffRequestId{
		TeamId:           teamId,
		TimeOffRequestId: timeOffRequestId,
	}
}

// ParseMeJoinedTeamIdScheduleTimeOffRequestID parses 'input' into a MeJoinedTeamIdScheduleTimeOffRequestId
func ParseMeJoinedTeamIdScheduleTimeOffRequestID(input string) (*MeJoinedTeamIdScheduleTimeOffRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdScheduleTimeOffRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdScheduleTimeOffRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdScheduleTimeOffRequestIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdScheduleTimeOffRequestId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdScheduleTimeOffRequestIDInsensitively(input string) (*MeJoinedTeamIdScheduleTimeOffRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdScheduleTimeOffRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdScheduleTimeOffRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdScheduleTimeOffRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.TimeOffRequestId, ok = input.Parsed["timeOffRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "timeOffRequestId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdScheduleTimeOffRequestID checks that 'input' can be parsed as a Me Joined Team Id Schedule Time Off Request ID
func ValidateMeJoinedTeamIdScheduleTimeOffRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdScheduleTimeOffRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Schedule Time Off Request ID
func (id MeJoinedTeamIdScheduleTimeOffRequestId) ID() string {
	fmtString := "/me/joinedTeams/%s/schedule/timeOffRequests/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.TimeOffRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Schedule Time Off Request ID
func (id MeJoinedTeamIdScheduleTimeOffRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("timeOffRequests", "timeOffRequests", "timeOffRequests"),
		resourceids.UserSpecifiedSegment("timeOffRequestId", "timeOffRequestId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Schedule Time Off Request ID
func (id MeJoinedTeamIdScheduleTimeOffRequestId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Time Off Request: %q", id.TimeOffRequestId),
	}
	return fmt.Sprintf("Me Joined Team Id Schedule Time Off Request (%s)", strings.Join(components, "\n"))
}
