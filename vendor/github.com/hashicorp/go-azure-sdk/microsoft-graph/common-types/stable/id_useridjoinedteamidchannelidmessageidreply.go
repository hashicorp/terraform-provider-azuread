package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdChannelIdMessageIdReplyId{}

// UserIdJoinedTeamIdChannelIdMessageIdReplyId is a struct representing the Resource ID for a User Id Joined Team Id Channel Id Message Id Reply
type UserIdJoinedTeamIdChannelIdMessageIdReplyId struct {
	UserId         string
	TeamId         string
	ChannelId      string
	ChatMessageId  string
	ChatMessageId1 string
}

// NewUserIdJoinedTeamIdChannelIdMessageIdReplyID returns a new UserIdJoinedTeamIdChannelIdMessageIdReplyId struct
func NewUserIdJoinedTeamIdChannelIdMessageIdReplyID(userId string, teamId string, channelId string, chatMessageId string, chatMessageId1 string) UserIdJoinedTeamIdChannelIdMessageIdReplyId {
	return UserIdJoinedTeamIdChannelIdMessageIdReplyId{
		UserId:         userId,
		TeamId:         teamId,
		ChannelId:      channelId,
		ChatMessageId:  chatMessageId,
		ChatMessageId1: chatMessageId1,
	}
}

// ParseUserIdJoinedTeamIdChannelIdMessageIdReplyID parses 'input' into a UserIdJoinedTeamIdChannelIdMessageIdReplyId
func ParseUserIdJoinedTeamIdChannelIdMessageIdReplyID(input string) (*UserIdJoinedTeamIdChannelIdMessageIdReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdChannelIdMessageIdReplyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdChannelIdMessageIdReplyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdChannelIdMessageIdReplyIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdChannelIdMessageIdReplyId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdChannelIdMessageIdReplyIDInsensitively(input string) (*UserIdJoinedTeamIdChannelIdMessageIdReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdChannelIdMessageIdReplyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdChannelIdMessageIdReplyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdChannelIdMessageIdReplyId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ChatMessageId1, ok = input.Parsed["chatMessageId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId1", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdChannelIdMessageIdReplyID checks that 'input' can be parsed as a User Id Joined Team Id Channel Id Message Id Reply ID
func ValidateUserIdJoinedTeamIdChannelIdMessageIdReplyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdChannelIdMessageIdReplyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Channel Id Message Id Reply ID
func (id UserIdJoinedTeamIdChannelIdMessageIdReplyId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/channels/%s/messages/%s/replies/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChannelId, id.ChatMessageId, id.ChatMessageId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Channel Id Message Id Reply ID
func (id UserIdJoinedTeamIdChannelIdMessageIdReplyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
		resourceids.StaticSegment("replies", "replies", "replies"),
		resourceids.UserSpecifiedSegment("chatMessageId1", "chatMessageId1"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Channel Id Message Id Reply ID
func (id UserIdJoinedTeamIdChannelIdMessageIdReplyId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
	}
	return fmt.Sprintf("User Id Joined Team Id Channel Id Message Id Reply (%s)", strings.Join(components, "\n"))
}
