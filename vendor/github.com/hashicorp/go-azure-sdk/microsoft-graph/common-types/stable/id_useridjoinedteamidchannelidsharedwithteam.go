package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdChannelIdSharedWithTeamId{}

// UserIdJoinedTeamIdChannelIdSharedWithTeamId is a struct representing the Resource ID for a User Id Joined Team Id Channel Id Shared With Team
type UserIdJoinedTeamIdChannelIdSharedWithTeamId struct {
	UserId                      string
	TeamId                      string
	ChannelId                   string
	SharedWithChannelTeamInfoId string
}

// NewUserIdJoinedTeamIdChannelIdSharedWithTeamID returns a new UserIdJoinedTeamIdChannelIdSharedWithTeamId struct
func NewUserIdJoinedTeamIdChannelIdSharedWithTeamID(userId string, teamId string, channelId string, sharedWithChannelTeamInfoId string) UserIdJoinedTeamIdChannelIdSharedWithTeamId {
	return UserIdJoinedTeamIdChannelIdSharedWithTeamId{
		UserId:                      userId,
		TeamId:                      teamId,
		ChannelId:                   channelId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
	}
}

// ParseUserIdJoinedTeamIdChannelIdSharedWithTeamID parses 'input' into a UserIdJoinedTeamIdChannelIdSharedWithTeamId
func ParseUserIdJoinedTeamIdChannelIdSharedWithTeamID(input string) (*UserIdJoinedTeamIdChannelIdSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdChannelIdSharedWithTeamId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdChannelIdSharedWithTeamId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdChannelIdSharedWithTeamIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdChannelIdSharedWithTeamId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdChannelIdSharedWithTeamIDInsensitively(input string) (*UserIdJoinedTeamIdChannelIdSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdChannelIdSharedWithTeamId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdChannelIdSharedWithTeamId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdChannelIdSharedWithTeamId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.SharedWithChannelTeamInfoId, ok = input.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdChannelIdSharedWithTeamID checks that 'input' can be parsed as a User Id Joined Team Id Channel Id Shared With Team ID
func ValidateUserIdJoinedTeamIdChannelIdSharedWithTeamID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdChannelIdSharedWithTeamID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Channel Id Shared With Team ID
func (id UserIdJoinedTeamIdChannelIdSharedWithTeamId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/channels/%s/sharedWithTeams/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChannelId, id.SharedWithChannelTeamInfoId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Channel Id Shared With Team ID
func (id UserIdJoinedTeamIdChannelIdSharedWithTeamId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("sharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Channel Id Shared With Team ID
func (id UserIdJoinedTeamIdChannelIdSharedWithTeamId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
	}
	return fmt.Sprintf("User Id Joined Team Id Channel Id Shared With Team (%s)", strings.Join(components, "\n"))
}
