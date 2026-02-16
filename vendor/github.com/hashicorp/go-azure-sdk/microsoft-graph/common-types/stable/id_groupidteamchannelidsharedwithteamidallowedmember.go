package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamChannelIdSharedWithTeamIdAllowedMemberId{}

// GroupIdTeamChannelIdSharedWithTeamIdAllowedMemberId is a struct representing the Resource ID for a Group Id Team Channel Id Shared With Team Id Allowed Member
type GroupIdTeamChannelIdSharedWithTeamIdAllowedMemberId struct {
	GroupId                     string
	ChannelId                   string
	SharedWithChannelTeamInfoId string
	ConversationMemberId        string
}

// NewGroupIdTeamChannelIdSharedWithTeamIdAllowedMemberID returns a new GroupIdTeamChannelIdSharedWithTeamIdAllowedMemberId struct
func NewGroupIdTeamChannelIdSharedWithTeamIdAllowedMemberID(groupId string, channelId string, sharedWithChannelTeamInfoId string, conversationMemberId string) GroupIdTeamChannelIdSharedWithTeamIdAllowedMemberId {
	return GroupIdTeamChannelIdSharedWithTeamIdAllowedMemberId{
		GroupId:                     groupId,
		ChannelId:                   channelId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
		ConversationMemberId:        conversationMemberId,
	}
}

// ParseGroupIdTeamChannelIdSharedWithTeamIdAllowedMemberID parses 'input' into a GroupIdTeamChannelIdSharedWithTeamIdAllowedMemberId
func ParseGroupIdTeamChannelIdSharedWithTeamIdAllowedMemberID(input string) (*GroupIdTeamChannelIdSharedWithTeamIdAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelIdSharedWithTeamIdAllowedMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelIdSharedWithTeamIdAllowedMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamChannelIdSharedWithTeamIdAllowedMemberIDInsensitively parses 'input' case-insensitively into a GroupIdTeamChannelIdSharedWithTeamIdAllowedMemberId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamChannelIdSharedWithTeamIdAllowedMemberIDInsensitively(input string) (*GroupIdTeamChannelIdSharedWithTeamIdAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelIdSharedWithTeamIdAllowedMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelIdSharedWithTeamIdAllowedMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamChannelIdSharedWithTeamIdAllowedMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ChannelId, ok = input.Parsed["channelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "channelId", input)
	}

	if id.SharedWithChannelTeamInfoId, ok = input.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", input)
	}

	if id.ConversationMemberId, ok = input.Parsed["conversationMemberId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", input)
	}

	return nil
}

// ValidateGroupIdTeamChannelIdSharedWithTeamIdAllowedMemberID checks that 'input' can be parsed as a Group Id Team Channel Id Shared With Team Id Allowed Member ID
func ValidateGroupIdTeamChannelIdSharedWithTeamIdAllowedMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamChannelIdSharedWithTeamIdAllowedMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Channel Id Shared With Team Id Allowed Member ID
func (id GroupIdTeamChannelIdSharedWithTeamIdAllowedMemberId) ID() string {
	fmtString := "/groups/%s/team/channels/%s/sharedWithTeams/%s/allowedMembers/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId, id.SharedWithChannelTeamInfoId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Channel Id Shared With Team Id Allowed Member ID
func (id GroupIdTeamChannelIdSharedWithTeamIdAllowedMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("sharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoId"),
		resourceids.StaticSegment("allowedMembers", "allowedMembers", "allowedMembers"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberId"),
	}
}

// String returns a human-readable description of this Group Id Team Channel Id Shared With Team Id Allowed Member ID
func (id GroupIdTeamChannelIdSharedWithTeamIdAllowedMemberId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Group Id Team Channel Id Shared With Team Id Allowed Member (%s)", strings.Join(components, "\n"))
}
