package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamChannelIdTabId{}

// GroupIdTeamChannelIdTabId is a struct representing the Resource ID for a Group Id Team Channel Id Tab
type GroupIdTeamChannelIdTabId struct {
	GroupId    string
	ChannelId  string
	TeamsTabId string
}

// NewGroupIdTeamChannelIdTabID returns a new GroupIdTeamChannelIdTabId struct
func NewGroupIdTeamChannelIdTabID(groupId string, channelId string, teamsTabId string) GroupIdTeamChannelIdTabId {
	return GroupIdTeamChannelIdTabId{
		GroupId:    groupId,
		ChannelId:  channelId,
		TeamsTabId: teamsTabId,
	}
}

// ParseGroupIdTeamChannelIdTabID parses 'input' into a GroupIdTeamChannelIdTabId
func ParseGroupIdTeamChannelIdTabID(input string) (*GroupIdTeamChannelIdTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelIdTabId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelIdTabId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamChannelIdTabIDInsensitively parses 'input' case-insensitively into a GroupIdTeamChannelIdTabId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamChannelIdTabIDInsensitively(input string) (*GroupIdTeamChannelIdTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelIdTabId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelIdTabId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamChannelIdTabId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ChannelId, ok = input.Parsed["channelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "channelId", input)
	}

	if id.TeamsTabId, ok = input.Parsed["teamsTabId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", input)
	}

	return nil
}

// ValidateGroupIdTeamChannelIdTabID checks that 'input' can be parsed as a Group Id Team Channel Id Tab ID
func ValidateGroupIdTeamChannelIdTabID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamChannelIdTabID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Channel Id Tab ID
func (id GroupIdTeamChannelIdTabId) ID() string {
	fmtString := "/groups/%s/team/channels/%s/tabs/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId, id.TeamsTabId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Channel Id Tab ID
func (id GroupIdTeamChannelIdTabId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("tabs", "tabs", "tabs"),
		resourceids.UserSpecifiedSegment("teamsTabId", "teamsTabId"),
	}
}

// String returns a human-readable description of this Group Id Team Channel Id Tab ID
func (id GroupIdTeamChannelIdTabId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Teams Tab: %q", id.TeamsTabId),
	}
	return fmt.Sprintf("Group Id Team Channel Id Tab (%s)", strings.Join(components, "\n"))
}
