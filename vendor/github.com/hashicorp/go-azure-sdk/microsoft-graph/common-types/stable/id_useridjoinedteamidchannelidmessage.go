package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdChannelIdMessageId{}

// UserIdJoinedTeamIdChannelIdMessageId is a struct representing the Resource ID for a User Id Joined Team Id Channel Id Message
type UserIdJoinedTeamIdChannelIdMessageId struct {
	UserId        string
	TeamId        string
	ChannelId     string
	ChatMessageId string
}

// NewUserIdJoinedTeamIdChannelIdMessageID returns a new UserIdJoinedTeamIdChannelIdMessageId struct
func NewUserIdJoinedTeamIdChannelIdMessageID(userId string, teamId string, channelId string, chatMessageId string) UserIdJoinedTeamIdChannelIdMessageId {
	return UserIdJoinedTeamIdChannelIdMessageId{
		UserId:        userId,
		TeamId:        teamId,
		ChannelId:     channelId,
		ChatMessageId: chatMessageId,
	}
}

// ParseUserIdJoinedTeamIdChannelIdMessageID parses 'input' into a UserIdJoinedTeamIdChannelIdMessageId
func ParseUserIdJoinedTeamIdChannelIdMessageID(input string) (*UserIdJoinedTeamIdChannelIdMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdChannelIdMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdChannelIdMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdChannelIdMessageIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdChannelIdMessageId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdChannelIdMessageIDInsensitively(input string) (*UserIdJoinedTeamIdChannelIdMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdChannelIdMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdChannelIdMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdChannelIdMessageId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ChatMessageId, ok = input.Parsed["chatMessageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdChannelIdMessageID checks that 'input' can be parsed as a User Id Joined Team Id Channel Id Message ID
func ValidateUserIdJoinedTeamIdChannelIdMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdChannelIdMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Channel Id Message ID
func (id UserIdJoinedTeamIdChannelIdMessageId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/channels/%s/messages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChannelId, id.ChatMessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Channel Id Message ID
func (id UserIdJoinedTeamIdChannelIdMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Channel Id Message ID
func (id UserIdJoinedTeamIdChannelIdMessageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
	}
	return fmt.Sprintf("User Id Joined Team Id Channel Id Message (%s)", strings.Join(components, "\n"))
}
