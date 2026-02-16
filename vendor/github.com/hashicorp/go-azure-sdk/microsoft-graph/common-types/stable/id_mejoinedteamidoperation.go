package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdOperationId{}

// MeJoinedTeamIdOperationId is a struct representing the Resource ID for a Me Joined Team Id Operation
type MeJoinedTeamIdOperationId struct {
	TeamId                string
	TeamsAsyncOperationId string
}

// NewMeJoinedTeamIdOperationID returns a new MeJoinedTeamIdOperationId struct
func NewMeJoinedTeamIdOperationID(teamId string, teamsAsyncOperationId string) MeJoinedTeamIdOperationId {
	return MeJoinedTeamIdOperationId{
		TeamId:                teamId,
		TeamsAsyncOperationId: teamsAsyncOperationId,
	}
}

// ParseMeJoinedTeamIdOperationID parses 'input' into a MeJoinedTeamIdOperationId
func ParseMeJoinedTeamIdOperationID(input string) (*MeJoinedTeamIdOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdOperationIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdOperationId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdOperationIDInsensitively(input string) (*MeJoinedTeamIdOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.TeamsAsyncOperationId, ok = input.Parsed["teamsAsyncOperationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamsAsyncOperationId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdOperationID checks that 'input' can be parsed as a Me Joined Team Id Operation ID
func ValidateMeJoinedTeamIdOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Operation ID
func (id MeJoinedTeamIdOperationId) ID() string {
	fmtString := "/me/joinedTeams/%s/operations/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.TeamsAsyncOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Operation ID
func (id MeJoinedTeamIdOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("operations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("teamsAsyncOperationId", "teamsAsyncOperationId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Operation ID
func (id MeJoinedTeamIdOperationId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Teams Async Operation: %q", id.TeamsAsyncOperationId),
	}
	return fmt.Sprintf("Me Joined Team Id Operation (%s)", strings.Join(components, "\n"))
}
