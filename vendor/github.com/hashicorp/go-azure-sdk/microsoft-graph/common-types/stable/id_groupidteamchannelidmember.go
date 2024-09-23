package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamChannelIdMemberId{}

// GroupIdTeamChannelIdMemberId is a struct representing the Resource ID for a Group Id Team Channel Id Member
type GroupIdTeamChannelIdMemberId struct {
	GroupId              string
	ChannelId            string
	ConversationMemberId string
}

// NewGroupIdTeamChannelIdMemberID returns a new GroupIdTeamChannelIdMemberId struct
func NewGroupIdTeamChannelIdMemberID(groupId string, channelId string, conversationMemberId string) GroupIdTeamChannelIdMemberId {
	return GroupIdTeamChannelIdMemberId{
		GroupId:              groupId,
		ChannelId:            channelId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseGroupIdTeamChannelIdMemberID parses 'input' into a GroupIdTeamChannelIdMemberId
func ParseGroupIdTeamChannelIdMemberID(input string) (*GroupIdTeamChannelIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelIdMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamChannelIdMemberIDInsensitively parses 'input' case-insensitively into a GroupIdTeamChannelIdMemberId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamChannelIdMemberIDInsensitively(input string) (*GroupIdTeamChannelIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelIdMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamChannelIdMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ChannelId, ok = input.Parsed["channelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "channelId", input)
	}

	if id.ConversationMemberId, ok = input.Parsed["conversationMemberId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", input)
	}

	return nil
}

// ValidateGroupIdTeamChannelIdMemberID checks that 'input' can be parsed as a Group Id Team Channel Id Member ID
func ValidateGroupIdTeamChannelIdMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamChannelIdMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Channel Id Member ID
func (id GroupIdTeamChannelIdMemberId) ID() string {
	fmtString := "/groups/%s/team/channels/%s/members/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Channel Id Member ID
func (id GroupIdTeamChannelIdMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("members", "members", "members"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberId"),
	}
}

// String returns a human-readable description of this Group Id Team Channel Id Member ID
func (id GroupIdTeamChannelIdMemberId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Group Id Team Channel Id Member (%s)", strings.Join(components, "\n"))
}
