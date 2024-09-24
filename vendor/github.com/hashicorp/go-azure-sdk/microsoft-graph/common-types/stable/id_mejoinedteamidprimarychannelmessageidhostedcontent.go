package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId{}

// MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId is a struct representing the Resource ID for a Me Joined Team Id Primary Channel Message Id Hosted Content
type MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId struct {
	TeamId                     string
	ChatMessageId              string
	ChatMessageHostedContentId string
}

// NewMeJoinedTeamIdPrimaryChannelMessageIdHostedContentID returns a new MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId struct
func NewMeJoinedTeamIdPrimaryChannelMessageIdHostedContentID(teamId string, chatMessageId string, chatMessageHostedContentId string) MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId {
	return MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId{
		TeamId:                     teamId,
		ChatMessageId:              chatMessageId,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseMeJoinedTeamIdPrimaryChannelMessageIdHostedContentID parses 'input' into a MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId
func ParseMeJoinedTeamIdPrimaryChannelMessageIdHostedContentID(input string) (*MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdPrimaryChannelMessageIdHostedContentIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdPrimaryChannelMessageIdHostedContentIDInsensitively(input string) (*MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

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

// ValidateMeJoinedTeamIdPrimaryChannelMessageIdHostedContentID checks that 'input' can be parsed as a Me Joined Team Id Primary Channel Message Id Hosted Content ID
func ValidateMeJoinedTeamIdPrimaryChannelMessageIdHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdPrimaryChannelMessageIdHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Primary Channel Message Id Hosted Content ID
func (id MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId) ID() string {
	fmtString := "/me/joinedTeams/%s/primaryChannel/messages/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChatMessageId, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Primary Channel Message Id Hosted Content ID
func (id MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
		resourceids.StaticSegment("hostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Primary Channel Message Id Hosted Content ID
func (id MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("Me Joined Team Id Primary Channel Message Id Hosted Content (%s)", strings.Join(components, "\n"))
}
