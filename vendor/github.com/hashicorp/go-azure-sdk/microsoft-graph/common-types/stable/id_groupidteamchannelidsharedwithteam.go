package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamChannelIdSharedWithTeamId{}

// GroupIdTeamChannelIdSharedWithTeamId is a struct representing the Resource ID for a Group Id Team Channel Id Shared With Team
type GroupIdTeamChannelIdSharedWithTeamId struct {
	GroupId                     string
	ChannelId                   string
	SharedWithChannelTeamInfoId string
}

// NewGroupIdTeamChannelIdSharedWithTeamID returns a new GroupIdTeamChannelIdSharedWithTeamId struct
func NewGroupIdTeamChannelIdSharedWithTeamID(groupId string, channelId string, sharedWithChannelTeamInfoId string) GroupIdTeamChannelIdSharedWithTeamId {
	return GroupIdTeamChannelIdSharedWithTeamId{
		GroupId:                     groupId,
		ChannelId:                   channelId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
	}
}

// ParseGroupIdTeamChannelIdSharedWithTeamID parses 'input' into a GroupIdTeamChannelIdSharedWithTeamId
func ParseGroupIdTeamChannelIdSharedWithTeamID(input string) (*GroupIdTeamChannelIdSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelIdSharedWithTeamId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelIdSharedWithTeamId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamChannelIdSharedWithTeamIDInsensitively parses 'input' case-insensitively into a GroupIdTeamChannelIdSharedWithTeamId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamChannelIdSharedWithTeamIDInsensitively(input string) (*GroupIdTeamChannelIdSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelIdSharedWithTeamId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelIdSharedWithTeamId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamChannelIdSharedWithTeamId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateGroupIdTeamChannelIdSharedWithTeamID checks that 'input' can be parsed as a Group Id Team Channel Id Shared With Team ID
func ValidateGroupIdTeamChannelIdSharedWithTeamID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamChannelIdSharedWithTeamID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Channel Id Shared With Team ID
func (id GroupIdTeamChannelIdSharedWithTeamId) ID() string {
	fmtString := "/groups/%s/team/channels/%s/sharedWithTeams/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId, id.SharedWithChannelTeamInfoId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Channel Id Shared With Team ID
func (id GroupIdTeamChannelIdSharedWithTeamId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("sharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoId"),
	}
}

// String returns a human-readable description of this Group Id Team Channel Id Shared With Team ID
func (id GroupIdTeamChannelIdSharedWithTeamId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
	}
	return fmt.Sprintf("Group Id Team Channel Id Shared With Team (%s)", strings.Join(components, "\n"))
}
