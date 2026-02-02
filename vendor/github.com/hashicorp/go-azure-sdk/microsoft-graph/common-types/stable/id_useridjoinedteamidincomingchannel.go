package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdIncomingChannelId{}

// UserIdJoinedTeamIdIncomingChannelId is a struct representing the Resource ID for a User Id Joined Team Id Incoming Channel
type UserIdJoinedTeamIdIncomingChannelId struct {
	UserId    string
	TeamId    string
	ChannelId string
}

// NewUserIdJoinedTeamIdIncomingChannelID returns a new UserIdJoinedTeamIdIncomingChannelId struct
func NewUserIdJoinedTeamIdIncomingChannelID(userId string, teamId string, channelId string) UserIdJoinedTeamIdIncomingChannelId {
	return UserIdJoinedTeamIdIncomingChannelId{
		UserId:    userId,
		TeamId:    teamId,
		ChannelId: channelId,
	}
}

// ParseUserIdJoinedTeamIdIncomingChannelID parses 'input' into a UserIdJoinedTeamIdIncomingChannelId
func ParseUserIdJoinedTeamIdIncomingChannelID(input string) (*UserIdJoinedTeamIdIncomingChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdIncomingChannelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdIncomingChannelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdIncomingChannelIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdIncomingChannelId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdIncomingChannelIDInsensitively(input string) (*UserIdJoinedTeamIdIncomingChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdIncomingChannelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdIncomingChannelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdIncomingChannelId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateUserIdJoinedTeamIdIncomingChannelID checks that 'input' can be parsed as a User Id Joined Team Id Incoming Channel ID
func ValidateUserIdJoinedTeamIdIncomingChannelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdIncomingChannelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Incoming Channel ID
func (id UserIdJoinedTeamIdIncomingChannelId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/incomingChannels/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChannelId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Incoming Channel ID
func (id UserIdJoinedTeamIdIncomingChannelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("incomingChannels", "incomingChannels", "incomingChannels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Incoming Channel ID
func (id UserIdJoinedTeamIdIncomingChannelId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
	}
	return fmt.Sprintf("User Id Joined Team Id Incoming Channel (%s)", strings.Join(components, "\n"))
}
