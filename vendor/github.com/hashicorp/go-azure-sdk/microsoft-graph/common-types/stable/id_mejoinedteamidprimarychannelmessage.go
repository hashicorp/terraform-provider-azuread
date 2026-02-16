package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdPrimaryChannelMessageId{}

// MeJoinedTeamIdPrimaryChannelMessageId is a struct representing the Resource ID for a Me Joined Team Id Primary Channel Message
type MeJoinedTeamIdPrimaryChannelMessageId struct {
	TeamId        string
	ChatMessageId string
}

// NewMeJoinedTeamIdPrimaryChannelMessageID returns a new MeJoinedTeamIdPrimaryChannelMessageId struct
func NewMeJoinedTeamIdPrimaryChannelMessageID(teamId string, chatMessageId string) MeJoinedTeamIdPrimaryChannelMessageId {
	return MeJoinedTeamIdPrimaryChannelMessageId{
		TeamId:        teamId,
		ChatMessageId: chatMessageId,
	}
}

// ParseMeJoinedTeamIdPrimaryChannelMessageID parses 'input' into a MeJoinedTeamIdPrimaryChannelMessageId
func ParseMeJoinedTeamIdPrimaryChannelMessageID(input string) (*MeJoinedTeamIdPrimaryChannelMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdPrimaryChannelMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdPrimaryChannelMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdPrimaryChannelMessageIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdPrimaryChannelMessageId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdPrimaryChannelMessageIDInsensitively(input string) (*MeJoinedTeamIdPrimaryChannelMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdPrimaryChannelMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdPrimaryChannelMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdPrimaryChannelMessageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.ChatMessageId, ok = input.Parsed["chatMessageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdPrimaryChannelMessageID checks that 'input' can be parsed as a Me Joined Team Id Primary Channel Message ID
func ValidateMeJoinedTeamIdPrimaryChannelMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdPrimaryChannelMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Primary Channel Message ID
func (id MeJoinedTeamIdPrimaryChannelMessageId) ID() string {
	fmtString := "/me/joinedTeams/%s/primaryChannel/messages/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChatMessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Primary Channel Message ID
func (id MeJoinedTeamIdPrimaryChannelMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Primary Channel Message ID
func (id MeJoinedTeamIdPrimaryChannelMessageId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
	}
	return fmt.Sprintf("Me Joined Team Id Primary Channel Message (%s)", strings.Join(components, "\n"))
}
