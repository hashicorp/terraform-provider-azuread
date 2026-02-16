package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdMemberId{}

// MeJoinedTeamIdMemberId is a struct representing the Resource ID for a Me Joined Team Id Member
type MeJoinedTeamIdMemberId struct {
	TeamId               string
	ConversationMemberId string
}

// NewMeJoinedTeamIdMemberID returns a new MeJoinedTeamIdMemberId struct
func NewMeJoinedTeamIdMemberID(teamId string, conversationMemberId string) MeJoinedTeamIdMemberId {
	return MeJoinedTeamIdMemberId{
		TeamId:               teamId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseMeJoinedTeamIdMemberID parses 'input' into a MeJoinedTeamIdMemberId
func ParseMeJoinedTeamIdMemberID(input string) (*MeJoinedTeamIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdMemberIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdMemberId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdMemberIDInsensitively(input string) (*MeJoinedTeamIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.ConversationMemberId, ok = input.Parsed["conversationMemberId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdMemberID checks that 'input' can be parsed as a Me Joined Team Id Member ID
func ValidateMeJoinedTeamIdMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Member ID
func (id MeJoinedTeamIdMemberId) ID() string {
	fmtString := "/me/joinedTeams/%s/members/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Member ID
func (id MeJoinedTeamIdMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("members", "members", "members"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Member ID
func (id MeJoinedTeamIdMemberId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Me Joined Team Id Member (%s)", strings.Join(components, "\n"))
}
