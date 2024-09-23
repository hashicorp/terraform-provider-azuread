package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamAllChannelId{}

// GroupIdTeamAllChannelId is a struct representing the Resource ID for a Group Id Team All Channel
type GroupIdTeamAllChannelId struct {
	GroupId   string
	ChannelId string
}

// NewGroupIdTeamAllChannelID returns a new GroupIdTeamAllChannelId struct
func NewGroupIdTeamAllChannelID(groupId string, channelId string) GroupIdTeamAllChannelId {
	return GroupIdTeamAllChannelId{
		GroupId:   groupId,
		ChannelId: channelId,
	}
}

// ParseGroupIdTeamAllChannelID parses 'input' into a GroupIdTeamAllChannelId
func ParseGroupIdTeamAllChannelID(input string) (*GroupIdTeamAllChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamAllChannelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamAllChannelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamAllChannelIDInsensitively parses 'input' case-insensitively into a GroupIdTeamAllChannelId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamAllChannelIDInsensitively(input string) (*GroupIdTeamAllChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamAllChannelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamAllChannelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamAllChannelId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ChannelId, ok = input.Parsed["channelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "channelId", input)
	}

	return nil
}

// ValidateGroupIdTeamAllChannelID checks that 'input' can be parsed as a Group Id Team All Channel ID
func ValidateGroupIdTeamAllChannelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamAllChannelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team All Channel ID
func (id GroupIdTeamAllChannelId) ID() string {
	fmtString := "/groups/%s/team/allChannels/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team All Channel ID
func (id GroupIdTeamAllChannelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("allChannels", "allChannels", "allChannels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
	}
}

// String returns a human-readable description of this Group Id Team All Channel ID
func (id GroupIdTeamAllChannelId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
	}
	return fmt.Sprintf("Group Id Team All Channel (%s)", strings.Join(components, "\n"))
}
