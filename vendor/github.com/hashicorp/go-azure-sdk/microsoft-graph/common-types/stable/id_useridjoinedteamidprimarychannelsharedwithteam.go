package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId{}

// UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId is a struct representing the Resource ID for a User Id Joined Team Id Primary Channel Shared With Team
type UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId struct {
	UserId                      string
	TeamId                      string
	SharedWithChannelTeamInfoId string
}

// NewUserIdJoinedTeamIdPrimaryChannelSharedWithTeamID returns a new UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId struct
func NewUserIdJoinedTeamIdPrimaryChannelSharedWithTeamID(userId string, teamId string, sharedWithChannelTeamInfoId string) UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId {
	return UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId{
		UserId:                      userId,
		TeamId:                      teamId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
	}
}

// ParseUserIdJoinedTeamIdPrimaryChannelSharedWithTeamID parses 'input' into a UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId
func ParseUserIdJoinedTeamIdPrimaryChannelSharedWithTeamID(input string) (*UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdPrimaryChannelSharedWithTeamIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdPrimaryChannelSharedWithTeamIDInsensitively(input string) (*UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.SharedWithChannelTeamInfoId, ok = input.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdPrimaryChannelSharedWithTeamID checks that 'input' can be parsed as a User Id Joined Team Id Primary Channel Shared With Team ID
func ValidateUserIdJoinedTeamIdPrimaryChannelSharedWithTeamID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdPrimaryChannelSharedWithTeamID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Primary Channel Shared With Team ID
func (id UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/primaryChannel/sharedWithTeams/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.SharedWithChannelTeamInfoId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Primary Channel Shared With Team ID
func (id UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("sharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Primary Channel Shared With Team ID
func (id UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
	}
	return fmt.Sprintf("User Id Joined Team Id Primary Channel Shared With Team (%s)", strings.Join(components, "\n"))
}
