package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdPrimaryChannelMessageIdReplyId{}

// MeJoinedTeamIdPrimaryChannelMessageIdReplyId is a struct representing the Resource ID for a Me Joined Team Id Primary Channel Message Id Reply
type MeJoinedTeamIdPrimaryChannelMessageIdReplyId struct {
	TeamId         string
	ChatMessageId  string
	ChatMessageId1 string
}

// NewMeJoinedTeamIdPrimaryChannelMessageIdReplyID returns a new MeJoinedTeamIdPrimaryChannelMessageIdReplyId struct
func NewMeJoinedTeamIdPrimaryChannelMessageIdReplyID(teamId string, chatMessageId string, chatMessageId1 string) MeJoinedTeamIdPrimaryChannelMessageIdReplyId {
	return MeJoinedTeamIdPrimaryChannelMessageIdReplyId{
		TeamId:         teamId,
		ChatMessageId:  chatMessageId,
		ChatMessageId1: chatMessageId1,
	}
}

// ParseMeJoinedTeamIdPrimaryChannelMessageIdReplyID parses 'input' into a MeJoinedTeamIdPrimaryChannelMessageIdReplyId
func ParseMeJoinedTeamIdPrimaryChannelMessageIdReplyID(input string) (*MeJoinedTeamIdPrimaryChannelMessageIdReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdPrimaryChannelMessageIdReplyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdPrimaryChannelMessageIdReplyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdPrimaryChannelMessageIdReplyIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdPrimaryChannelMessageIdReplyId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdPrimaryChannelMessageIdReplyIDInsensitively(input string) (*MeJoinedTeamIdPrimaryChannelMessageIdReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdPrimaryChannelMessageIdReplyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdPrimaryChannelMessageIdReplyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdPrimaryChannelMessageIdReplyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.ChatMessageId, ok = input.Parsed["chatMessageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", input)
	}

	if id.ChatMessageId1, ok = input.Parsed["chatMessageId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId1", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdPrimaryChannelMessageIdReplyID checks that 'input' can be parsed as a Me Joined Team Id Primary Channel Message Id Reply ID
func ValidateMeJoinedTeamIdPrimaryChannelMessageIdReplyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdPrimaryChannelMessageIdReplyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Primary Channel Message Id Reply ID
func (id MeJoinedTeamIdPrimaryChannelMessageIdReplyId) ID() string {
	fmtString := "/me/joinedTeams/%s/primaryChannel/messages/%s/replies/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChatMessageId, id.ChatMessageId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Primary Channel Message Id Reply ID
func (id MeJoinedTeamIdPrimaryChannelMessageIdReplyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
		resourceids.StaticSegment("replies", "replies", "replies"),
		resourceids.UserSpecifiedSegment("chatMessageId1", "chatMessageId1"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Primary Channel Message Id Reply ID
func (id MeJoinedTeamIdPrimaryChannelMessageIdReplyId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
	}
	return fmt.Sprintf("Me Joined Team Id Primary Channel Message Id Reply (%s)", strings.Join(components, "\n"))
}
