package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdPrimaryChannelAllMemberId{}

// MeJoinedTeamIdPrimaryChannelAllMemberId is a struct representing the Resource ID for a Me Joined Team Id Primary Channel All Member
type MeJoinedTeamIdPrimaryChannelAllMemberId struct {
	TeamId               string
	ConversationMemberId string
}

// NewMeJoinedTeamIdPrimaryChannelAllMemberID returns a new MeJoinedTeamIdPrimaryChannelAllMemberId struct
func NewMeJoinedTeamIdPrimaryChannelAllMemberID(teamId string, conversationMemberId string) MeJoinedTeamIdPrimaryChannelAllMemberId {
	return MeJoinedTeamIdPrimaryChannelAllMemberId{
		TeamId:               teamId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseMeJoinedTeamIdPrimaryChannelAllMemberID parses 'input' into a MeJoinedTeamIdPrimaryChannelAllMemberId
func ParseMeJoinedTeamIdPrimaryChannelAllMemberID(input string) (*MeJoinedTeamIdPrimaryChannelAllMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdPrimaryChannelAllMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdPrimaryChannelAllMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdPrimaryChannelAllMemberIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdPrimaryChannelAllMemberId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdPrimaryChannelAllMemberIDInsensitively(input string) (*MeJoinedTeamIdPrimaryChannelAllMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdPrimaryChannelAllMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdPrimaryChannelAllMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdPrimaryChannelAllMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.ConversationMemberId, ok = input.Parsed["conversationMemberId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdPrimaryChannelAllMemberID checks that 'input' can be parsed as a Me Joined Team Id Primary Channel All Member ID
func ValidateMeJoinedTeamIdPrimaryChannelAllMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdPrimaryChannelAllMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Primary Channel All Member ID
func (id MeJoinedTeamIdPrimaryChannelAllMemberId) ID() string {
	fmtString := "/me/joinedTeams/%s/primaryChannel/allMembers/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Primary Channel All Member ID
func (id MeJoinedTeamIdPrimaryChannelAllMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("allMembers", "allMembers", "allMembers"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Primary Channel All Member ID
func (id MeJoinedTeamIdPrimaryChannelAllMemberId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Me Joined Team Id Primary Channel All Member (%s)", strings.Join(components, "\n"))
}
