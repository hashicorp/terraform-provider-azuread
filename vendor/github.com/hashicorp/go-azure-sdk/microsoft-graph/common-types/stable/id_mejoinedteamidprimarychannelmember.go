package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdPrimaryChannelMemberId{}

// MeJoinedTeamIdPrimaryChannelMemberId is a struct representing the Resource ID for a Me Joined Team Id Primary Channel Member
type MeJoinedTeamIdPrimaryChannelMemberId struct {
	TeamId               string
	ConversationMemberId string
}

// NewMeJoinedTeamIdPrimaryChannelMemberID returns a new MeJoinedTeamIdPrimaryChannelMemberId struct
func NewMeJoinedTeamIdPrimaryChannelMemberID(teamId string, conversationMemberId string) MeJoinedTeamIdPrimaryChannelMemberId {
	return MeJoinedTeamIdPrimaryChannelMemberId{
		TeamId:               teamId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseMeJoinedTeamIdPrimaryChannelMemberID parses 'input' into a MeJoinedTeamIdPrimaryChannelMemberId
func ParseMeJoinedTeamIdPrimaryChannelMemberID(input string) (*MeJoinedTeamIdPrimaryChannelMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdPrimaryChannelMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdPrimaryChannelMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdPrimaryChannelMemberIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdPrimaryChannelMemberId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdPrimaryChannelMemberIDInsensitively(input string) (*MeJoinedTeamIdPrimaryChannelMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdPrimaryChannelMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdPrimaryChannelMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdPrimaryChannelMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.ConversationMemberId, ok = input.Parsed["conversationMemberId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdPrimaryChannelMemberID checks that 'input' can be parsed as a Me Joined Team Id Primary Channel Member ID
func ValidateMeJoinedTeamIdPrimaryChannelMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdPrimaryChannelMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Primary Channel Member ID
func (id MeJoinedTeamIdPrimaryChannelMemberId) ID() string {
	fmtString := "/me/joinedTeams/%s/primaryChannel/members/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Primary Channel Member ID
func (id MeJoinedTeamIdPrimaryChannelMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("members", "members", "members"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Primary Channel Member ID
func (id MeJoinedTeamIdPrimaryChannelMemberId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Me Joined Team Id Primary Channel Member (%s)", strings.Join(components, "\n"))
}
