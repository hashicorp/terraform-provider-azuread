package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdPrimaryChannelMessageId{}

// UserIdJoinedTeamIdPrimaryChannelMessageId is a struct representing the Resource ID for a User Id Joined Team Id Primary Channel Message
type UserIdJoinedTeamIdPrimaryChannelMessageId struct {
	UserId        string
	TeamId        string
	ChatMessageId string
}

// NewUserIdJoinedTeamIdPrimaryChannelMessageID returns a new UserIdJoinedTeamIdPrimaryChannelMessageId struct
func NewUserIdJoinedTeamIdPrimaryChannelMessageID(userId string, teamId string, chatMessageId string) UserIdJoinedTeamIdPrimaryChannelMessageId {
	return UserIdJoinedTeamIdPrimaryChannelMessageId{
		UserId:        userId,
		TeamId:        teamId,
		ChatMessageId: chatMessageId,
	}
}

// ParseUserIdJoinedTeamIdPrimaryChannelMessageID parses 'input' into a UserIdJoinedTeamIdPrimaryChannelMessageId
func ParseUserIdJoinedTeamIdPrimaryChannelMessageID(input string) (*UserIdJoinedTeamIdPrimaryChannelMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdPrimaryChannelMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdPrimaryChannelMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdPrimaryChannelMessageIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdPrimaryChannelMessageId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdPrimaryChannelMessageIDInsensitively(input string) (*UserIdJoinedTeamIdPrimaryChannelMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdPrimaryChannelMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdPrimaryChannelMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdPrimaryChannelMessageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.ChatMessageId, ok = input.Parsed["chatMessageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdPrimaryChannelMessageID checks that 'input' can be parsed as a User Id Joined Team Id Primary Channel Message ID
func ValidateUserIdJoinedTeamIdPrimaryChannelMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdPrimaryChannelMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Primary Channel Message ID
func (id UserIdJoinedTeamIdPrimaryChannelMessageId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/primaryChannel/messages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChatMessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Primary Channel Message ID
func (id UserIdJoinedTeamIdPrimaryChannelMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Primary Channel Message ID
func (id UserIdJoinedTeamIdPrimaryChannelMessageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
	}
	return fmt.Sprintf("User Id Joined Team Id Primary Channel Message (%s)", strings.Join(components, "\n"))
}
