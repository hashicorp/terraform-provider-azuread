package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdScheduleSwapShiftsChangeRequestId{}

// MeJoinedTeamIdScheduleSwapShiftsChangeRequestId is a struct representing the Resource ID for a Me Joined Team Id Schedule Swap Shifts Change Request
type MeJoinedTeamIdScheduleSwapShiftsChangeRequestId struct {
	TeamId                    string
	SwapShiftsChangeRequestId string
}

// NewMeJoinedTeamIdScheduleSwapShiftsChangeRequestID returns a new MeJoinedTeamIdScheduleSwapShiftsChangeRequestId struct
func NewMeJoinedTeamIdScheduleSwapShiftsChangeRequestID(teamId string, swapShiftsChangeRequestId string) MeJoinedTeamIdScheduleSwapShiftsChangeRequestId {
	return MeJoinedTeamIdScheduleSwapShiftsChangeRequestId{
		TeamId:                    teamId,
		SwapShiftsChangeRequestId: swapShiftsChangeRequestId,
	}
}

// ParseMeJoinedTeamIdScheduleSwapShiftsChangeRequestID parses 'input' into a MeJoinedTeamIdScheduleSwapShiftsChangeRequestId
func ParseMeJoinedTeamIdScheduleSwapShiftsChangeRequestID(input string) (*MeJoinedTeamIdScheduleSwapShiftsChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdScheduleSwapShiftsChangeRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdScheduleSwapShiftsChangeRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdScheduleSwapShiftsChangeRequestIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdScheduleSwapShiftsChangeRequestId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdScheduleSwapShiftsChangeRequestIDInsensitively(input string) (*MeJoinedTeamIdScheduleSwapShiftsChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdScheduleSwapShiftsChangeRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdScheduleSwapShiftsChangeRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdScheduleSwapShiftsChangeRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.SwapShiftsChangeRequestId, ok = input.Parsed["swapShiftsChangeRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "swapShiftsChangeRequestId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdScheduleSwapShiftsChangeRequestID checks that 'input' can be parsed as a Me Joined Team Id Schedule Swap Shifts Change Request ID
func ValidateMeJoinedTeamIdScheduleSwapShiftsChangeRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdScheduleSwapShiftsChangeRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Schedule Swap Shifts Change Request ID
func (id MeJoinedTeamIdScheduleSwapShiftsChangeRequestId) ID() string {
	fmtString := "/me/joinedTeams/%s/schedule/swapShiftsChangeRequests/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.SwapShiftsChangeRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Schedule Swap Shifts Change Request ID
func (id MeJoinedTeamIdScheduleSwapShiftsChangeRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("swapShiftsChangeRequests", "swapShiftsChangeRequests", "swapShiftsChangeRequests"),
		resourceids.UserSpecifiedSegment("swapShiftsChangeRequestId", "swapShiftsChangeRequestId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Schedule Swap Shifts Change Request ID
func (id MeJoinedTeamIdScheduleSwapShiftsChangeRequestId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Swap Shifts Change Request: %q", id.SwapShiftsChangeRequestId),
	}
	return fmt.Sprintf("Me Joined Team Id Schedule Swap Shifts Change Request (%s)", strings.Join(components, "\n"))
}
