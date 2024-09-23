package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdAllChannelId{}

// UserIdJoinedTeamIdAllChannelId is a struct representing the Resource ID for a User Id Joined Team Id All Channel
type UserIdJoinedTeamIdAllChannelId struct {
	UserId    string
	TeamId    string
	ChannelId string
}

// NewUserIdJoinedTeamIdAllChannelID returns a new UserIdJoinedTeamIdAllChannelId struct
func NewUserIdJoinedTeamIdAllChannelID(userId string, teamId string, channelId string) UserIdJoinedTeamIdAllChannelId {
	return UserIdJoinedTeamIdAllChannelId{
		UserId:    userId,
		TeamId:    teamId,
		ChannelId: channelId,
	}
}

// ParseUserIdJoinedTeamIdAllChannelID parses 'input' into a UserIdJoinedTeamIdAllChannelId
func ParseUserIdJoinedTeamIdAllChannelID(input string) (*UserIdJoinedTeamIdAllChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdAllChannelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdAllChannelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdAllChannelIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdAllChannelId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdAllChannelIDInsensitively(input string) (*UserIdJoinedTeamIdAllChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdAllChannelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdAllChannelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdAllChannelId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateUserIdJoinedTeamIdAllChannelID checks that 'input' can be parsed as a User Id Joined Team Id All Channel ID
func ValidateUserIdJoinedTeamIdAllChannelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdAllChannelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id All Channel ID
func (id UserIdJoinedTeamIdAllChannelId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/allChannels/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChannelId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id All Channel ID
func (id UserIdJoinedTeamIdAllChannelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("allChannels", "allChannels", "allChannels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id All Channel ID
func (id UserIdJoinedTeamIdAllChannelId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
	}
	return fmt.Sprintf("User Id Joined Team Id All Channel (%s)", strings.Join(components, "\n"))
}
