package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdChannelIdMemberId{}

// UserIdJoinedTeamIdChannelIdMemberId is a struct representing the Resource ID for a User Id Joined Team Id Channel Id Member
type UserIdJoinedTeamIdChannelIdMemberId struct {
	UserId               string
	TeamId               string
	ChannelId            string
	ConversationMemberId string
}

// NewUserIdJoinedTeamIdChannelIdMemberID returns a new UserIdJoinedTeamIdChannelIdMemberId struct
func NewUserIdJoinedTeamIdChannelIdMemberID(userId string, teamId string, channelId string, conversationMemberId string) UserIdJoinedTeamIdChannelIdMemberId {
	return UserIdJoinedTeamIdChannelIdMemberId{
		UserId:               userId,
		TeamId:               teamId,
		ChannelId:            channelId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseUserIdJoinedTeamIdChannelIdMemberID parses 'input' into a UserIdJoinedTeamIdChannelIdMemberId
func ParseUserIdJoinedTeamIdChannelIdMemberID(input string) (*UserIdJoinedTeamIdChannelIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdChannelIdMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdChannelIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdChannelIdMemberIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdChannelIdMemberId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdChannelIdMemberIDInsensitively(input string) (*UserIdJoinedTeamIdChannelIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdChannelIdMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdChannelIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdChannelIdMemberId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ConversationMemberId, ok = input.Parsed["conversationMemberId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdChannelIdMemberID checks that 'input' can be parsed as a User Id Joined Team Id Channel Id Member ID
func ValidateUserIdJoinedTeamIdChannelIdMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdChannelIdMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Channel Id Member ID
func (id UserIdJoinedTeamIdChannelIdMemberId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/channels/%s/members/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChannelId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Channel Id Member ID
func (id UserIdJoinedTeamIdChannelIdMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("members", "members", "members"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Channel Id Member ID
func (id UserIdJoinedTeamIdChannelIdMemberId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("User Id Joined Team Id Channel Id Member (%s)", strings.Join(components, "\n"))
}
