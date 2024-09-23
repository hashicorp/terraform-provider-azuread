package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId{}

// UserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId is a struct representing the Resource ID for a User Id Joined Team Id Channel Id Shared With Team Id Allowed Member
type UserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId struct {
	UserId                      string
	TeamId                      string
	ChannelId                   string
	SharedWithChannelTeamInfoId string
	ConversationMemberId        string
}

// NewUserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberID returns a new UserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId struct
func NewUserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberID(userId string, teamId string, channelId string, sharedWithChannelTeamInfoId string, conversationMemberId string) UserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId {
	return UserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId{
		UserId:                      userId,
		TeamId:                      teamId,
		ChannelId:                   channelId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
		ConversationMemberId:        conversationMemberId,
	}
}

// ParseUserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberID parses 'input' into a UserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId
func ParseUserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberID(input string) (*UserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberIDInsensitively(input string) (*UserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ConversationMemberId, ok = input.Parsed["conversationMemberId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberID checks that 'input' can be parsed as a User Id Joined Team Id Channel Id Shared With Team Id Allowed Member ID
func ValidateUserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Channel Id Shared With Team Id Allowed Member ID
func (id UserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/channels/%s/sharedWithTeams/%s/allowedMembers/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChannelId, id.SharedWithChannelTeamInfoId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Channel Id Shared With Team Id Allowed Member ID
func (id UserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("sharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoId"),
		resourceids.StaticSegment("allowedMembers", "allowedMembers", "allowedMembers"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Channel Id Shared With Team Id Allowed Member ID
func (id UserIdJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("User Id Joined Team Id Channel Id Shared With Team Id Allowed Member (%s)", strings.Join(components, "\n"))
}
