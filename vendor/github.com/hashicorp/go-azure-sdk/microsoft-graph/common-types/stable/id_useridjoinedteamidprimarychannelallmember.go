package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdPrimaryChannelAllMemberId{}

// UserIdJoinedTeamIdPrimaryChannelAllMemberId is a struct representing the Resource ID for a User Id Joined Team Id Primary Channel All Member
type UserIdJoinedTeamIdPrimaryChannelAllMemberId struct {
	UserId               string
	TeamId               string
	ConversationMemberId string
}

// NewUserIdJoinedTeamIdPrimaryChannelAllMemberID returns a new UserIdJoinedTeamIdPrimaryChannelAllMemberId struct
func NewUserIdJoinedTeamIdPrimaryChannelAllMemberID(userId string, teamId string, conversationMemberId string) UserIdJoinedTeamIdPrimaryChannelAllMemberId {
	return UserIdJoinedTeamIdPrimaryChannelAllMemberId{
		UserId:               userId,
		TeamId:               teamId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseUserIdJoinedTeamIdPrimaryChannelAllMemberID parses 'input' into a UserIdJoinedTeamIdPrimaryChannelAllMemberId
func ParseUserIdJoinedTeamIdPrimaryChannelAllMemberID(input string) (*UserIdJoinedTeamIdPrimaryChannelAllMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdPrimaryChannelAllMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdPrimaryChannelAllMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdPrimaryChannelAllMemberIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdPrimaryChannelAllMemberId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdPrimaryChannelAllMemberIDInsensitively(input string) (*UserIdJoinedTeamIdPrimaryChannelAllMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdPrimaryChannelAllMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdPrimaryChannelAllMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdPrimaryChannelAllMemberId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateUserIdJoinedTeamIdPrimaryChannelAllMemberID checks that 'input' can be parsed as a User Id Joined Team Id Primary Channel All Member ID
func ValidateUserIdJoinedTeamIdPrimaryChannelAllMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdPrimaryChannelAllMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Primary Channel All Member ID
func (id UserIdJoinedTeamIdPrimaryChannelAllMemberId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/primaryChannel/allMembers/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Primary Channel All Member ID
func (id UserIdJoinedTeamIdPrimaryChannelAllMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("allMembers", "allMembers", "allMembers"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Primary Channel All Member ID
func (id UserIdJoinedTeamIdPrimaryChannelAllMemberId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("User Id Joined Team Id Primary Channel All Member (%s)", strings.Join(components, "\n"))
}
