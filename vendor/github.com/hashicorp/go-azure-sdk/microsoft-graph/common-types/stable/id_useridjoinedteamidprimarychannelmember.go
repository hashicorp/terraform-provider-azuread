package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdPrimaryChannelMemberId{}

// UserIdJoinedTeamIdPrimaryChannelMemberId is a struct representing the Resource ID for a User Id Joined Team Id Primary Channel Member
type UserIdJoinedTeamIdPrimaryChannelMemberId struct {
	UserId               string
	TeamId               string
	ConversationMemberId string
}

// NewUserIdJoinedTeamIdPrimaryChannelMemberID returns a new UserIdJoinedTeamIdPrimaryChannelMemberId struct
func NewUserIdJoinedTeamIdPrimaryChannelMemberID(userId string, teamId string, conversationMemberId string) UserIdJoinedTeamIdPrimaryChannelMemberId {
	return UserIdJoinedTeamIdPrimaryChannelMemberId{
		UserId:               userId,
		TeamId:               teamId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseUserIdJoinedTeamIdPrimaryChannelMemberID parses 'input' into a UserIdJoinedTeamIdPrimaryChannelMemberId
func ParseUserIdJoinedTeamIdPrimaryChannelMemberID(input string) (*UserIdJoinedTeamIdPrimaryChannelMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdPrimaryChannelMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdPrimaryChannelMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdPrimaryChannelMemberIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdPrimaryChannelMemberId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdPrimaryChannelMemberIDInsensitively(input string) (*UserIdJoinedTeamIdPrimaryChannelMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdPrimaryChannelMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdPrimaryChannelMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdPrimaryChannelMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.ConversationMemberId, ok = input.Parsed["conversationMemberId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdPrimaryChannelMemberID checks that 'input' can be parsed as a User Id Joined Team Id Primary Channel Member ID
func ValidateUserIdJoinedTeamIdPrimaryChannelMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdPrimaryChannelMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Primary Channel Member ID
func (id UserIdJoinedTeamIdPrimaryChannelMemberId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/primaryChannel/members/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Primary Channel Member ID
func (id UserIdJoinedTeamIdPrimaryChannelMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("members", "members", "members"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Primary Channel Member ID
func (id UserIdJoinedTeamIdPrimaryChannelMemberId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("User Id Joined Team Id Primary Channel Member (%s)", strings.Join(components, "\n"))
}
