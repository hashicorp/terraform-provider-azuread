package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdChannelId{}

// UserIdJoinedTeamIdChannelId is a struct representing the Resource ID for a User Id Joined Team Id Channel
type UserIdJoinedTeamIdChannelId struct {
	UserId    string
	TeamId    string
	ChannelId string
}

// NewUserIdJoinedTeamIdChannelID returns a new UserIdJoinedTeamIdChannelId struct
func NewUserIdJoinedTeamIdChannelID(userId string, teamId string, channelId string) UserIdJoinedTeamIdChannelId {
	return UserIdJoinedTeamIdChannelId{
		UserId:    userId,
		TeamId:    teamId,
		ChannelId: channelId,
	}
}

// ParseUserIdJoinedTeamIdChannelID parses 'input' into a UserIdJoinedTeamIdChannelId
func ParseUserIdJoinedTeamIdChannelID(input string) (*UserIdJoinedTeamIdChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdChannelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdChannelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdChannelIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdChannelId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdChannelIDInsensitively(input string) (*UserIdJoinedTeamIdChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdChannelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdChannelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdChannelId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateUserIdJoinedTeamIdChannelID checks that 'input' can be parsed as a User Id Joined Team Id Channel ID
func ValidateUserIdJoinedTeamIdChannelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdChannelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Channel ID
func (id UserIdJoinedTeamIdChannelId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/channels/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChannelId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Channel ID
func (id UserIdJoinedTeamIdChannelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Channel ID
func (id UserIdJoinedTeamIdChannelId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
	}
	return fmt.Sprintf("User Id Joined Team Id Channel (%s)", strings.Join(components, "\n"))
}
