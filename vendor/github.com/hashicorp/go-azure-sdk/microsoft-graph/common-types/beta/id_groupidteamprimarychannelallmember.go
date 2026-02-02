package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamPrimaryChannelAllMemberId{}

// GroupIdTeamPrimaryChannelAllMemberId is a struct representing the Resource ID for a Group Id Team Primary Channel All Member
type GroupIdTeamPrimaryChannelAllMemberId struct {
	GroupId              string
	ConversationMemberId string
}

// NewGroupIdTeamPrimaryChannelAllMemberID returns a new GroupIdTeamPrimaryChannelAllMemberId struct
func NewGroupIdTeamPrimaryChannelAllMemberID(groupId string, conversationMemberId string) GroupIdTeamPrimaryChannelAllMemberId {
	return GroupIdTeamPrimaryChannelAllMemberId{
		GroupId:              groupId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseGroupIdTeamPrimaryChannelAllMemberID parses 'input' into a GroupIdTeamPrimaryChannelAllMemberId
func ParseGroupIdTeamPrimaryChannelAllMemberID(input string) (*GroupIdTeamPrimaryChannelAllMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelAllMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelAllMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamPrimaryChannelAllMemberIDInsensitively parses 'input' case-insensitively into a GroupIdTeamPrimaryChannelAllMemberId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamPrimaryChannelAllMemberIDInsensitively(input string) (*GroupIdTeamPrimaryChannelAllMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelAllMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelAllMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamPrimaryChannelAllMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ConversationMemberId, ok = input.Parsed["conversationMemberId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", input)
	}

	return nil
}

// ValidateGroupIdTeamPrimaryChannelAllMemberID checks that 'input' can be parsed as a Group Id Team Primary Channel All Member ID
func ValidateGroupIdTeamPrimaryChannelAllMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamPrimaryChannelAllMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Primary Channel All Member ID
func (id GroupIdTeamPrimaryChannelAllMemberId) ID() string {
	fmtString := "/groups/%s/team/primaryChannel/allMembers/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Primary Channel All Member ID
func (id GroupIdTeamPrimaryChannelAllMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("allMembers", "allMembers", "allMembers"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberId"),
	}
}

// String returns a human-readable description of this Group Id Team Primary Channel All Member ID
func (id GroupIdTeamPrimaryChannelAllMemberId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Group Id Team Primary Channel All Member (%s)", strings.Join(components, "\n"))
}
