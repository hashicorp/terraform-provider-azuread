package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamPrimaryChannelMemberId{}

// GroupIdTeamPrimaryChannelMemberId is a struct representing the Resource ID for a Group Id Team Primary Channel Member
type GroupIdTeamPrimaryChannelMemberId struct {
	GroupId              string
	ConversationMemberId string
}

// NewGroupIdTeamPrimaryChannelMemberID returns a new GroupIdTeamPrimaryChannelMemberId struct
func NewGroupIdTeamPrimaryChannelMemberID(groupId string, conversationMemberId string) GroupIdTeamPrimaryChannelMemberId {
	return GroupIdTeamPrimaryChannelMemberId{
		GroupId:              groupId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseGroupIdTeamPrimaryChannelMemberID parses 'input' into a GroupIdTeamPrimaryChannelMemberId
func ParseGroupIdTeamPrimaryChannelMemberID(input string) (*GroupIdTeamPrimaryChannelMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamPrimaryChannelMemberIDInsensitively parses 'input' case-insensitively into a GroupIdTeamPrimaryChannelMemberId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamPrimaryChannelMemberIDInsensitively(input string) (*GroupIdTeamPrimaryChannelMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamPrimaryChannelMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ConversationMemberId, ok = input.Parsed["conversationMemberId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", input)
	}

	return nil
}

// ValidateGroupIdTeamPrimaryChannelMemberID checks that 'input' can be parsed as a Group Id Team Primary Channel Member ID
func ValidateGroupIdTeamPrimaryChannelMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamPrimaryChannelMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Primary Channel Member ID
func (id GroupIdTeamPrimaryChannelMemberId) ID() string {
	fmtString := "/groups/%s/team/primaryChannel/members/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Primary Channel Member ID
func (id GroupIdTeamPrimaryChannelMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("members", "members", "members"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberId"),
	}
}

// String returns a human-readable description of this Group Id Team Primary Channel Member ID
func (id GroupIdTeamPrimaryChannelMemberId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Group Id Team Primary Channel Member (%s)", strings.Join(components, "\n"))
}
