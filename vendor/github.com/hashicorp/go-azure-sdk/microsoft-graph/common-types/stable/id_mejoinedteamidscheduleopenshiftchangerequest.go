package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdScheduleOpenShiftChangeRequestId{}

// MeJoinedTeamIdScheduleOpenShiftChangeRequestId is a struct representing the Resource ID for a Me Joined Team Id Schedule Open Shift Change Request
type MeJoinedTeamIdScheduleOpenShiftChangeRequestId struct {
	TeamId                   string
	OpenShiftChangeRequestId string
}

// NewMeJoinedTeamIdScheduleOpenShiftChangeRequestID returns a new MeJoinedTeamIdScheduleOpenShiftChangeRequestId struct
func NewMeJoinedTeamIdScheduleOpenShiftChangeRequestID(teamId string, openShiftChangeRequestId string) MeJoinedTeamIdScheduleOpenShiftChangeRequestId {
	return MeJoinedTeamIdScheduleOpenShiftChangeRequestId{
		TeamId:                   teamId,
		OpenShiftChangeRequestId: openShiftChangeRequestId,
	}
}

// ParseMeJoinedTeamIdScheduleOpenShiftChangeRequestID parses 'input' into a MeJoinedTeamIdScheduleOpenShiftChangeRequestId
func ParseMeJoinedTeamIdScheduleOpenShiftChangeRequestID(input string) (*MeJoinedTeamIdScheduleOpenShiftChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdScheduleOpenShiftChangeRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdScheduleOpenShiftChangeRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdScheduleOpenShiftChangeRequestIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdScheduleOpenShiftChangeRequestId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdScheduleOpenShiftChangeRequestIDInsensitively(input string) (*MeJoinedTeamIdScheduleOpenShiftChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdScheduleOpenShiftChangeRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdScheduleOpenShiftChangeRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdScheduleOpenShiftChangeRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.OpenShiftChangeRequestId, ok = input.Parsed["openShiftChangeRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "openShiftChangeRequestId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdScheduleOpenShiftChangeRequestID checks that 'input' can be parsed as a Me Joined Team Id Schedule Open Shift Change Request ID
func ValidateMeJoinedTeamIdScheduleOpenShiftChangeRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdScheduleOpenShiftChangeRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Schedule Open Shift Change Request ID
func (id MeJoinedTeamIdScheduleOpenShiftChangeRequestId) ID() string {
	fmtString := "/me/joinedTeams/%s/schedule/openShiftChangeRequests/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.OpenShiftChangeRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Schedule Open Shift Change Request ID
func (id MeJoinedTeamIdScheduleOpenShiftChangeRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("openShiftChangeRequests", "openShiftChangeRequests", "openShiftChangeRequests"),
		resourceids.UserSpecifiedSegment("openShiftChangeRequestId", "openShiftChangeRequestId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Schedule Open Shift Change Request ID
func (id MeJoinedTeamIdScheduleOpenShiftChangeRequestId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Open Shift Change Request: %q", id.OpenShiftChangeRequestId),
	}
	return fmt.Sprintf("Me Joined Team Id Schedule Open Shift Change Request (%s)", strings.Join(components, "\n"))
}
