package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentId{}

// UserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentId is a struct representing the Resource ID for a User Id Joined Team Id Primary Channel Message Id Hosted Content
type UserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentId struct {
	UserId                     string
	TeamId                     string
	ChatMessageId              string
	ChatMessageHostedContentId string
}

// NewUserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentID returns a new UserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentId struct
func NewUserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentID(userId string, teamId string, chatMessageId string, chatMessageHostedContentId string) UserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentId {
	return UserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentId{
		UserId:                     userId,
		TeamId:                     teamId,
		ChatMessageId:              chatMessageId,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseUserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentID parses 'input' into a UserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentId
func ParseUserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentID(input string) (*UserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentIDInsensitively(input string) (*UserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ChatMessageHostedContentId, ok = input.Parsed["chatMessageHostedContentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentID checks that 'input' can be parsed as a User Id Joined Team Id Primary Channel Message Id Hosted Content ID
func ValidateUserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Primary Channel Message Id Hosted Content ID
func (id UserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/primaryChannel/messages/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChatMessageId, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Primary Channel Message Id Hosted Content ID
func (id UserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
		resourceids.StaticSegment("hostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Primary Channel Message Id Hosted Content ID
func (id UserIdJoinedTeamIdPrimaryChannelMessageIdHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("User Id Joined Team Id Primary Channel Message Id Hosted Content (%s)", strings.Join(components, "\n"))
}
