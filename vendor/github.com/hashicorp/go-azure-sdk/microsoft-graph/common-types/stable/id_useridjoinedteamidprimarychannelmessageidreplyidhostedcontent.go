package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId{}

// UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId is a struct representing the Resource ID for a User Id Joined Team Id Primary Channel Message Id Reply Id Hosted Content
type UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId struct {
	UserId                     string
	TeamId                     string
	ChatMessageId              string
	ChatMessageId1             string
	ChatMessageHostedContentId string
}

// NewUserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentID returns a new UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId struct
func NewUserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentID(userId string, teamId string, chatMessageId string, chatMessageId1 string, chatMessageHostedContentId string) UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId {
	return UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId{
		UserId:                     userId,
		TeamId:                     teamId,
		ChatMessageId:              chatMessageId,
		ChatMessageId1:             chatMessageId1,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseUserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentID parses 'input' into a UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId
func ParseUserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentID(input string) (*UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentIDInsensitively(input string) (*UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ChatMessageId1, ok = input.Parsed["chatMessageId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId1", input)
	}

	if id.ChatMessageHostedContentId, ok = input.Parsed["chatMessageHostedContentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentID checks that 'input' can be parsed as a User Id Joined Team Id Primary Channel Message Id Reply Id Hosted Content ID
func ValidateUserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Primary Channel Message Id Reply Id Hosted Content ID
func (id UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/primaryChannel/messages/%s/replies/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChatMessageId, id.ChatMessageId1, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Primary Channel Message Id Reply Id Hosted Content ID
func (id UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
		resourceids.StaticSegment("replies", "replies", "replies"),
		resourceids.UserSpecifiedSegment("chatMessageId1", "chatMessageId1"),
		resourceids.StaticSegment("hostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Primary Channel Message Id Reply Id Hosted Content ID
func (id UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("User Id Joined Team Id Primary Channel Message Id Reply Id Hosted Content (%s)", strings.Join(components, "\n"))
}
