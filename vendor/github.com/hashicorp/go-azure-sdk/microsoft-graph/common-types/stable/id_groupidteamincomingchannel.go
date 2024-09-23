package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamIncomingChannelId{}

// GroupIdTeamIncomingChannelId is a struct representing the Resource ID for a Group Id Team Incoming Channel
type GroupIdTeamIncomingChannelId struct {
	GroupId   string
	ChannelId string
}

// NewGroupIdTeamIncomingChannelID returns a new GroupIdTeamIncomingChannelId struct
func NewGroupIdTeamIncomingChannelID(groupId string, channelId string) GroupIdTeamIncomingChannelId {
	return GroupIdTeamIncomingChannelId{
		GroupId:   groupId,
		ChannelId: channelId,
	}
}

// ParseGroupIdTeamIncomingChannelID parses 'input' into a GroupIdTeamIncomingChannelId
func ParseGroupIdTeamIncomingChannelID(input string) (*GroupIdTeamIncomingChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamIncomingChannelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamIncomingChannelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamIncomingChannelIDInsensitively parses 'input' case-insensitively into a GroupIdTeamIncomingChannelId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamIncomingChannelIDInsensitively(input string) (*GroupIdTeamIncomingChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamIncomingChannelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamIncomingChannelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamIncomingChannelId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ChannelId, ok = input.Parsed["channelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "channelId", input)
	}

	return nil
}

// ValidateGroupIdTeamIncomingChannelID checks that 'input' can be parsed as a Group Id Team Incoming Channel ID
func ValidateGroupIdTeamIncomingChannelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamIncomingChannelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Incoming Channel ID
func (id GroupIdTeamIncomingChannelId) ID() string {
	fmtString := "/groups/%s/team/incomingChannels/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Incoming Channel ID
func (id GroupIdTeamIncomingChannelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("incomingChannels", "incomingChannels", "incomingChannels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
	}
}

// String returns a human-readable description of this Group Id Team Incoming Channel ID
func (id GroupIdTeamIncomingChannelId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
	}
	return fmt.Sprintf("Group Id Team Incoming Channel (%s)", strings.Join(components, "\n"))
}
