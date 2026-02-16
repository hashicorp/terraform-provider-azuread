package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdChannelIdTabId{}

// UserIdJoinedTeamIdChannelIdTabId is a struct representing the Resource ID for a User Id Joined Team Id Channel Id Tab
type UserIdJoinedTeamIdChannelIdTabId struct {
	UserId     string
	TeamId     string
	ChannelId  string
	TeamsTabId string
}

// NewUserIdJoinedTeamIdChannelIdTabID returns a new UserIdJoinedTeamIdChannelIdTabId struct
func NewUserIdJoinedTeamIdChannelIdTabID(userId string, teamId string, channelId string, teamsTabId string) UserIdJoinedTeamIdChannelIdTabId {
	return UserIdJoinedTeamIdChannelIdTabId{
		UserId:     userId,
		TeamId:     teamId,
		ChannelId:  channelId,
		TeamsTabId: teamsTabId,
	}
}

// ParseUserIdJoinedTeamIdChannelIdTabID parses 'input' into a UserIdJoinedTeamIdChannelIdTabId
func ParseUserIdJoinedTeamIdChannelIdTabID(input string) (*UserIdJoinedTeamIdChannelIdTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdChannelIdTabId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdChannelIdTabId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdChannelIdTabIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdChannelIdTabId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdChannelIdTabIDInsensitively(input string) (*UserIdJoinedTeamIdChannelIdTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdChannelIdTabId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdChannelIdTabId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdChannelIdTabId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.ChannelId, ok = input.Parsed["channelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "channelId", input)
	}

	if id.TeamsTabId, ok = input.Parsed["teamsTabId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdChannelIdTabID checks that 'input' can be parsed as a User Id Joined Team Id Channel Id Tab ID
func ValidateUserIdJoinedTeamIdChannelIdTabID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdChannelIdTabID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Channel Id Tab ID
func (id UserIdJoinedTeamIdChannelIdTabId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/channels/%s/tabs/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChannelId, id.TeamsTabId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Channel Id Tab ID
func (id UserIdJoinedTeamIdChannelIdTabId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("tabs", "tabs", "tabs"),
		resourceids.UserSpecifiedSegment("teamsTabId", "teamsTabId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Channel Id Tab ID
func (id UserIdJoinedTeamIdChannelIdTabId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Teams Tab: %q", id.TeamsTabId),
	}
	return fmt.Sprintf("User Id Joined Team Id Channel Id Tab (%s)", strings.Join(components, "\n"))
}
