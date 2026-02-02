package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId{}

// UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId is a struct representing the Resource ID for a User Id Joined Team Id Primary Channel Shared With Team Id Allowed Member
type UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId struct {
	UserId                      string
	TeamId                      string
	SharedWithChannelTeamInfoId string
	ConversationMemberId        string
}

// NewUserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberID returns a new UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId struct
func NewUserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberID(userId string, teamId string, sharedWithChannelTeamInfoId string, conversationMemberId string) UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId {
	return UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId{
		UserId:                      userId,
		TeamId:                      teamId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
		ConversationMemberId:        conversationMemberId,
	}
}

// ParseUserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberID parses 'input' into a UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId
func ParseUserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberID(input string) (*UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberIDInsensitively(input string) (*UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.SharedWithChannelTeamInfoId, ok = input.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", input)
	}

	if id.ConversationMemberId, ok = input.Parsed["conversationMemberId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberID checks that 'input' can be parsed as a User Id Joined Team Id Primary Channel Shared With Team Id Allowed Member ID
func ValidateUserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Primary Channel Shared With Team Id Allowed Member ID
func (id UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/primaryChannel/sharedWithTeams/%s/allowedMembers/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.SharedWithChannelTeamInfoId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Primary Channel Shared With Team Id Allowed Member ID
func (id UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("sharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoId"),
		resourceids.StaticSegment("allowedMembers", "allowedMembers", "allowedMembers"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Primary Channel Shared With Team Id Allowed Member ID
func (id UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("User Id Joined Team Id Primary Channel Shared With Team Id Allowed Member (%s)", strings.Join(components, "\n"))
}
