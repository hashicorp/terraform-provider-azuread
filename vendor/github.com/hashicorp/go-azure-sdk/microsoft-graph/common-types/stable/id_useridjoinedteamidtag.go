package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdTagId{}

// UserIdJoinedTeamIdTagId is a struct representing the Resource ID for a User Id Joined Team Id Tag
type UserIdJoinedTeamIdTagId struct {
	UserId        string
	TeamId        string
	TeamworkTagId string
}

// NewUserIdJoinedTeamIdTagID returns a new UserIdJoinedTeamIdTagId struct
func NewUserIdJoinedTeamIdTagID(userId string, teamId string, teamworkTagId string) UserIdJoinedTeamIdTagId {
	return UserIdJoinedTeamIdTagId{
		UserId:        userId,
		TeamId:        teamId,
		TeamworkTagId: teamworkTagId,
	}
}

// ParseUserIdJoinedTeamIdTagID parses 'input' into a UserIdJoinedTeamIdTagId
func ParseUserIdJoinedTeamIdTagID(input string) (*UserIdJoinedTeamIdTagId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdTagId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdTagId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdTagIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdTagId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdTagIDInsensitively(input string) (*UserIdJoinedTeamIdTagId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdTagId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdTagId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdTagId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.TeamworkTagId, ok = input.Parsed["teamworkTagId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdTagID checks that 'input' can be parsed as a User Id Joined Team Id Tag ID
func ValidateUserIdJoinedTeamIdTagID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdTagID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Tag ID
func (id UserIdJoinedTeamIdTagId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/tags/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.TeamworkTagId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Tag ID
func (id UserIdJoinedTeamIdTagId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("tags", "tags", "tags"),
		resourceids.UserSpecifiedSegment("teamworkTagId", "teamworkTagId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Tag ID
func (id UserIdJoinedTeamIdTagId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Teamwork Tag: %q", id.TeamworkTagId),
	}
	return fmt.Sprintf("User Id Joined Team Id Tag (%s)", strings.Join(components, "\n"))
}
