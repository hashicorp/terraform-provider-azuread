package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamMemberId{}

// GroupIdTeamMemberId is a struct representing the Resource ID for a Group Id Team Member
type GroupIdTeamMemberId struct {
	GroupId              string
	ConversationMemberId string
}

// NewGroupIdTeamMemberID returns a new GroupIdTeamMemberId struct
func NewGroupIdTeamMemberID(groupId string, conversationMemberId string) GroupIdTeamMemberId {
	return GroupIdTeamMemberId{
		GroupId:              groupId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseGroupIdTeamMemberID parses 'input' into a GroupIdTeamMemberId
func ParseGroupIdTeamMemberID(input string) (*GroupIdTeamMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamMemberIDInsensitively parses 'input' case-insensitively into a GroupIdTeamMemberId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamMemberIDInsensitively(input string) (*GroupIdTeamMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ConversationMemberId, ok = input.Parsed["conversationMemberId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", input)
	}

	return nil
}

// ValidateGroupIdTeamMemberID checks that 'input' can be parsed as a Group Id Team Member ID
func ValidateGroupIdTeamMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Member ID
func (id GroupIdTeamMemberId) ID() string {
	fmtString := "/groups/%s/team/members/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Member ID
func (id GroupIdTeamMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("members", "members", "members"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberId"),
	}
}

// String returns a human-readable description of this Group Id Team Member ID
func (id GroupIdTeamMemberId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Group Id Team Member (%s)", strings.Join(components, "\n"))
}
