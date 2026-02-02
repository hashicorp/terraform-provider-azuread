package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdPrimaryChannelTabId{}

// UserIdJoinedTeamIdPrimaryChannelTabId is a struct representing the Resource ID for a User Id Joined Team Id Primary Channel Tab
type UserIdJoinedTeamIdPrimaryChannelTabId struct {
	UserId     string
	TeamId     string
	TeamsTabId string
}

// NewUserIdJoinedTeamIdPrimaryChannelTabID returns a new UserIdJoinedTeamIdPrimaryChannelTabId struct
func NewUserIdJoinedTeamIdPrimaryChannelTabID(userId string, teamId string, teamsTabId string) UserIdJoinedTeamIdPrimaryChannelTabId {
	return UserIdJoinedTeamIdPrimaryChannelTabId{
		UserId:     userId,
		TeamId:     teamId,
		TeamsTabId: teamsTabId,
	}
}

// ParseUserIdJoinedTeamIdPrimaryChannelTabID parses 'input' into a UserIdJoinedTeamIdPrimaryChannelTabId
func ParseUserIdJoinedTeamIdPrimaryChannelTabID(input string) (*UserIdJoinedTeamIdPrimaryChannelTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdPrimaryChannelTabId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdPrimaryChannelTabId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdPrimaryChannelTabIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdPrimaryChannelTabId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdPrimaryChannelTabIDInsensitively(input string) (*UserIdJoinedTeamIdPrimaryChannelTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdPrimaryChannelTabId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdPrimaryChannelTabId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdPrimaryChannelTabId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.TeamsTabId, ok = input.Parsed["teamsTabId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdPrimaryChannelTabID checks that 'input' can be parsed as a User Id Joined Team Id Primary Channel Tab ID
func ValidateUserIdJoinedTeamIdPrimaryChannelTabID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdPrimaryChannelTabID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Primary Channel Tab ID
func (id UserIdJoinedTeamIdPrimaryChannelTabId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/primaryChannel/tabs/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.TeamsTabId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Primary Channel Tab ID
func (id UserIdJoinedTeamIdPrimaryChannelTabId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("tabs", "tabs", "tabs"),
		resourceids.UserSpecifiedSegment("teamsTabId", "teamsTabId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Primary Channel Tab ID
func (id UserIdJoinedTeamIdPrimaryChannelTabId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Teams Tab: %q", id.TeamsTabId),
	}
	return fmt.Sprintf("User Id Joined Team Id Primary Channel Tab (%s)", strings.Join(components, "\n"))
}
