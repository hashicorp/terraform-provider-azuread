package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdChannelIdMessageIdHostedContentId{}

// UserIdJoinedTeamIdChannelIdMessageIdHostedContentId is a struct representing the Resource ID for a User Id Joined Team Id Channel Id Message Id Hosted Content
type UserIdJoinedTeamIdChannelIdMessageIdHostedContentId struct {
	UserId                     string
	TeamId                     string
	ChannelId                  string
	ChatMessageId              string
	ChatMessageHostedContentId string
}

// NewUserIdJoinedTeamIdChannelIdMessageIdHostedContentID returns a new UserIdJoinedTeamIdChannelIdMessageIdHostedContentId struct
func NewUserIdJoinedTeamIdChannelIdMessageIdHostedContentID(userId string, teamId string, channelId string, chatMessageId string, chatMessageHostedContentId string) UserIdJoinedTeamIdChannelIdMessageIdHostedContentId {
	return UserIdJoinedTeamIdChannelIdMessageIdHostedContentId{
		UserId:                     userId,
		TeamId:                     teamId,
		ChannelId:                  channelId,
		ChatMessageId:              chatMessageId,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseUserIdJoinedTeamIdChannelIdMessageIdHostedContentID parses 'input' into a UserIdJoinedTeamIdChannelIdMessageIdHostedContentId
func ParseUserIdJoinedTeamIdChannelIdMessageIdHostedContentID(input string) (*UserIdJoinedTeamIdChannelIdMessageIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdChannelIdMessageIdHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdChannelIdMessageIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdChannelIdMessageIdHostedContentIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdChannelIdMessageIdHostedContentId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdChannelIdMessageIdHostedContentIDInsensitively(input string) (*UserIdJoinedTeamIdChannelIdMessageIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdChannelIdMessageIdHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdChannelIdMessageIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdChannelIdMessageIdHostedContentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ChatMessageHostedContentId, ok = input.Parsed["chatMessageHostedContentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdChannelIdMessageIdHostedContentID checks that 'input' can be parsed as a User Id Joined Team Id Channel Id Message Id Hosted Content ID
func ValidateUserIdJoinedTeamIdChannelIdMessageIdHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdChannelIdMessageIdHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Channel Id Message Id Hosted Content ID
func (id UserIdJoinedTeamIdChannelIdMessageIdHostedContentId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/channels/%s/messages/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChannelId, id.ChatMessageId, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Channel Id Message Id Hosted Content ID
func (id UserIdJoinedTeamIdChannelIdMessageIdHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
		resourceids.StaticSegment("hostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Channel Id Message Id Hosted Content ID
func (id UserIdJoinedTeamIdChannelIdMessageIdHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("User Id Joined Team Id Channel Id Message Id Hosted Content (%s)", strings.Join(components, "\n"))
}
