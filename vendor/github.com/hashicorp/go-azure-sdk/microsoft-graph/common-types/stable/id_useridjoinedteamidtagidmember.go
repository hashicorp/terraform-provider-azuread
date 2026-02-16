package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdTagIdMemberId{}

// UserIdJoinedTeamIdTagIdMemberId is a struct representing the Resource ID for a User Id Joined Team Id Tag Id Member
type UserIdJoinedTeamIdTagIdMemberId struct {
	UserId              string
	TeamId              string
	TeamworkTagId       string
	TeamworkTagMemberId string
}

// NewUserIdJoinedTeamIdTagIdMemberID returns a new UserIdJoinedTeamIdTagIdMemberId struct
func NewUserIdJoinedTeamIdTagIdMemberID(userId string, teamId string, teamworkTagId string, teamworkTagMemberId string) UserIdJoinedTeamIdTagIdMemberId {
	return UserIdJoinedTeamIdTagIdMemberId{
		UserId:              userId,
		TeamId:              teamId,
		TeamworkTagId:       teamworkTagId,
		TeamworkTagMemberId: teamworkTagMemberId,
	}
}

// ParseUserIdJoinedTeamIdTagIdMemberID parses 'input' into a UserIdJoinedTeamIdTagIdMemberId
func ParseUserIdJoinedTeamIdTagIdMemberID(input string) (*UserIdJoinedTeamIdTagIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdTagIdMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdTagIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdTagIdMemberIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdTagIdMemberId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdTagIdMemberIDInsensitively(input string) (*UserIdJoinedTeamIdTagIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdTagIdMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdTagIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdTagIdMemberId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.TeamworkTagMemberId, ok = input.Parsed["teamworkTagMemberId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagMemberId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdTagIdMemberID checks that 'input' can be parsed as a User Id Joined Team Id Tag Id Member ID
func ValidateUserIdJoinedTeamIdTagIdMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdTagIdMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Tag Id Member ID
func (id UserIdJoinedTeamIdTagIdMemberId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/tags/%s/members/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.TeamworkTagId, id.TeamworkTagMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Tag Id Member ID
func (id UserIdJoinedTeamIdTagIdMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("tags", "tags", "tags"),
		resourceids.UserSpecifiedSegment("teamworkTagId", "teamworkTagId"),
		resourceids.StaticSegment("members", "members", "members"),
		resourceids.UserSpecifiedSegment("teamworkTagMemberId", "teamworkTagMemberId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Tag Id Member ID
func (id UserIdJoinedTeamIdTagIdMemberId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Teamwork Tag: %q", id.TeamworkTagId),
		fmt.Sprintf("Teamwork Tag Member: %q", id.TeamworkTagMemberId),
	}
	return fmt.Sprintf("User Id Joined Team Id Tag Id Member (%s)", strings.Join(components, "\n"))
}
