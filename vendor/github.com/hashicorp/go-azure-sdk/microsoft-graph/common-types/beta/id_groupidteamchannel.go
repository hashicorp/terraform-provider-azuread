package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamChannelId{}

// GroupIdTeamChannelId is a struct representing the Resource ID for a Group Id Team Channel
type GroupIdTeamChannelId struct {
	GroupId   string
	ChannelId string
}

// NewGroupIdTeamChannelID returns a new GroupIdTeamChannelId struct
func NewGroupIdTeamChannelID(groupId string, channelId string) GroupIdTeamChannelId {
	return GroupIdTeamChannelId{
		GroupId:   groupId,
		ChannelId: channelId,
	}
}

// ParseGroupIdTeamChannelID parses 'input' into a GroupIdTeamChannelId
func ParseGroupIdTeamChannelID(input string) (*GroupIdTeamChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamChannelIDInsensitively parses 'input' case-insensitively into a GroupIdTeamChannelId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamChannelIDInsensitively(input string) (*GroupIdTeamChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamChannelId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ChannelId, ok = input.Parsed["channelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "channelId", input)
	}

	return nil
}

// ValidateGroupIdTeamChannelID checks that 'input' can be parsed as a Group Id Team Channel ID
func ValidateGroupIdTeamChannelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamChannelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Channel ID
func (id GroupIdTeamChannelId) ID() string {
	fmtString := "/groups/%s/team/channels/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Channel ID
func (id GroupIdTeamChannelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
	}
}

// String returns a human-readable description of this Group Id Team Channel ID
func (id GroupIdTeamChannelId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
	}
	return fmt.Sprintf("Group Id Team Channel (%s)", strings.Join(components, "\n"))
}
