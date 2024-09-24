package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberId{}

// GroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberId is a struct representing the Resource ID for a Group Id Team Primary Channel Shared With Team Id Allowed Member
type GroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberId struct {
	GroupId                     string
	SharedWithChannelTeamInfoId string
	ConversationMemberId        string
}

// NewGroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberID returns a new GroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberId struct
func NewGroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberID(groupId string, sharedWithChannelTeamInfoId string, conversationMemberId string) GroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberId {
	return GroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberId{
		GroupId:                     groupId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
		ConversationMemberId:        conversationMemberId,
	}
}

// ParseGroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberID parses 'input' into a GroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberId
func ParseGroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberID(input string) (*GroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberIDInsensitively parses 'input' case-insensitively into a GroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberIDInsensitively(input string) (*GroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SharedWithChannelTeamInfoId, ok = input.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", input)
	}

	if id.ConversationMemberId, ok = input.Parsed["conversationMemberId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", input)
	}

	return nil
}

// ValidateGroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberID checks that 'input' can be parsed as a Group Id Team Primary Channel Shared With Team Id Allowed Member ID
func ValidateGroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Primary Channel Shared With Team Id Allowed Member ID
func (id GroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberId) ID() string {
	fmtString := "/groups/%s/team/primaryChannel/sharedWithTeams/%s/allowedMembers/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SharedWithChannelTeamInfoId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Primary Channel Shared With Team Id Allowed Member ID
func (id GroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("sharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoId"),
		resourceids.StaticSegment("allowedMembers", "allowedMembers", "allowedMembers"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberId"),
	}
}

// String returns a human-readable description of this Group Id Team Primary Channel Shared With Team Id Allowed Member ID
func (id GroupIdTeamPrimaryChannelSharedWithTeamIdAllowedMemberId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Group Id Team Primary Channel Shared With Team Id Allowed Member (%s)", strings.Join(components, "\n"))
}
