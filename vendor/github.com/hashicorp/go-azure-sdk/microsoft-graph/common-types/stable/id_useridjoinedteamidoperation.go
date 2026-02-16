package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdOperationId{}

// UserIdJoinedTeamIdOperationId is a struct representing the Resource ID for a User Id Joined Team Id Operation
type UserIdJoinedTeamIdOperationId struct {
	UserId                string
	TeamId                string
	TeamsAsyncOperationId string
}

// NewUserIdJoinedTeamIdOperationID returns a new UserIdJoinedTeamIdOperationId struct
func NewUserIdJoinedTeamIdOperationID(userId string, teamId string, teamsAsyncOperationId string) UserIdJoinedTeamIdOperationId {
	return UserIdJoinedTeamIdOperationId{
		UserId:                userId,
		TeamId:                teamId,
		TeamsAsyncOperationId: teamsAsyncOperationId,
	}
}

// ParseUserIdJoinedTeamIdOperationID parses 'input' into a UserIdJoinedTeamIdOperationId
func ParseUserIdJoinedTeamIdOperationID(input string) (*UserIdJoinedTeamIdOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdOperationIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdOperationId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdOperationIDInsensitively(input string) (*UserIdJoinedTeamIdOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.TeamsAsyncOperationId, ok = input.Parsed["teamsAsyncOperationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamsAsyncOperationId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdOperationID checks that 'input' can be parsed as a User Id Joined Team Id Operation ID
func ValidateUserIdJoinedTeamIdOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Operation ID
func (id UserIdJoinedTeamIdOperationId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/operations/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.TeamsAsyncOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Operation ID
func (id UserIdJoinedTeamIdOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("operations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("teamsAsyncOperationId", "teamsAsyncOperationId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Operation ID
func (id UserIdJoinedTeamIdOperationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Teams Async Operation: %q", id.TeamsAsyncOperationId),
	}
	return fmt.Sprintf("User Id Joined Team Id Operation (%s)", strings.Join(components, "\n"))
}
