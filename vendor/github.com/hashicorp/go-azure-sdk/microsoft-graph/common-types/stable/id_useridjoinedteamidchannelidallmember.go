package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdChannelIdAllMemberId{}

// UserIdJoinedTeamIdChannelIdAllMemberId is a struct representing the Resource ID for a User Id Joined Team Id Channel Id All Member
type UserIdJoinedTeamIdChannelIdAllMemberId struct {
	UserId               string
	TeamId               string
	ChannelId            string
	ConversationMemberId string
}

// NewUserIdJoinedTeamIdChannelIdAllMemberID returns a new UserIdJoinedTeamIdChannelIdAllMemberId struct
func NewUserIdJoinedTeamIdChannelIdAllMemberID(userId string, teamId string, channelId string, conversationMemberId string) UserIdJoinedTeamIdChannelIdAllMemberId {
	return UserIdJoinedTeamIdChannelIdAllMemberId{
		UserId:               userId,
		TeamId:               teamId,
		ChannelId:            channelId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseUserIdJoinedTeamIdChannelIdAllMemberID parses 'input' into a UserIdJoinedTeamIdChannelIdAllMemberId
func ParseUserIdJoinedTeamIdChannelIdAllMemberID(input string) (*UserIdJoinedTeamIdChannelIdAllMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdChannelIdAllMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdChannelIdAllMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdChannelIdAllMemberIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdChannelIdAllMemberId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdChannelIdAllMemberIDInsensitively(input string) (*UserIdJoinedTeamIdChannelIdAllMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdChannelIdAllMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdChannelIdAllMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdChannelIdAllMemberId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateUserIdJoinedTeamIdChannelIdAllMemberID checks that 'input' can be parsed as a User Id Joined Team Id Channel Id All Member ID
func ValidateUserIdJoinedTeamIdChannelIdAllMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdChannelIdAllMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Channel Id All Member ID
func (id UserIdJoinedTeamIdChannelIdAllMemberId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/channels/%s/allMembers/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChannelId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Channel Id All Member ID
func (id UserIdJoinedTeamIdChannelIdAllMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("allMembers", "allMembers", "allMembers"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Channel Id All Member ID
func (id UserIdJoinedTeamIdChannelIdAllMemberId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("User Id Joined Team Id Channel Id All Member (%s)", strings.Join(components, "\n"))
}
