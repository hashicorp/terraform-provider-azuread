package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId{}

// MeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId is a struct representing the Resource ID for a Me Joined Team Id Primary Channel Shared With Team Id Allowed Member
type MeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId struct {
	TeamId                      string
	SharedWithChannelTeamInfoId string
	ConversationMemberId        string
}

// NewMeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberID returns a new MeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId struct
func NewMeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberID(teamId string, sharedWithChannelTeamInfoId string, conversationMemberId string) MeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId {
	return MeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId{
		TeamId:                      teamId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
		ConversationMemberId:        conversationMemberId,
	}
}

// ParseMeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberID parses 'input' into a MeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId
func ParseMeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberID(input string) (*MeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberIDInsensitively(input string) (*MeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.SharedWithChannelTeamInfoId, ok = input.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", input)
	}

	if id.ConversationMemberId, ok = input.Parsed["conversationMemberId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberID checks that 'input' can be parsed as a Me Joined Team Id Primary Channel Shared With Team Id Allowed Member ID
func ValidateMeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Primary Channel Shared With Team Id Allowed Member ID
func (id MeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId) ID() string {
	fmtString := "/me/joinedTeams/%s/primaryChannel/sharedWithTeams/%s/allowedMembers/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.SharedWithChannelTeamInfoId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Primary Channel Shared With Team Id Allowed Member ID
func (id MeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("sharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoId"),
		resourceids.StaticSegment("allowedMembers", "allowedMembers", "allowedMembers"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Primary Channel Shared With Team Id Allowed Member ID
func (id MeJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Me Joined Team Id Primary Channel Shared With Team Id Allowed Member (%s)", strings.Join(components, "\n"))
}
