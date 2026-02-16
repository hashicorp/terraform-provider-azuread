package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamPrimaryChannelTabId{}

// GroupIdTeamPrimaryChannelTabId is a struct representing the Resource ID for a Group Id Team Primary Channel Tab
type GroupIdTeamPrimaryChannelTabId struct {
	GroupId    string
	TeamsTabId string
}

// NewGroupIdTeamPrimaryChannelTabID returns a new GroupIdTeamPrimaryChannelTabId struct
func NewGroupIdTeamPrimaryChannelTabID(groupId string, teamsTabId string) GroupIdTeamPrimaryChannelTabId {
	return GroupIdTeamPrimaryChannelTabId{
		GroupId:    groupId,
		TeamsTabId: teamsTabId,
	}
}

// ParseGroupIdTeamPrimaryChannelTabID parses 'input' into a GroupIdTeamPrimaryChannelTabId
func ParseGroupIdTeamPrimaryChannelTabID(input string) (*GroupIdTeamPrimaryChannelTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelTabId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelTabId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamPrimaryChannelTabIDInsensitively parses 'input' case-insensitively into a GroupIdTeamPrimaryChannelTabId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamPrimaryChannelTabIDInsensitively(input string) (*GroupIdTeamPrimaryChannelTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelTabId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelTabId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamPrimaryChannelTabId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.TeamsTabId, ok = input.Parsed["teamsTabId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", input)
	}

	return nil
}

// ValidateGroupIdTeamPrimaryChannelTabID checks that 'input' can be parsed as a Group Id Team Primary Channel Tab ID
func ValidateGroupIdTeamPrimaryChannelTabID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamPrimaryChannelTabID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Primary Channel Tab ID
func (id GroupIdTeamPrimaryChannelTabId) ID() string {
	fmtString := "/groups/%s/team/primaryChannel/tabs/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.TeamsTabId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Primary Channel Tab ID
func (id GroupIdTeamPrimaryChannelTabId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("tabs", "tabs", "tabs"),
		resourceids.UserSpecifiedSegment("teamsTabId", "teamsTabId"),
	}
}

// String returns a human-readable description of this Group Id Team Primary Channel Tab ID
func (id GroupIdTeamPrimaryChannelTabId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Teams Tab: %q", id.TeamsTabId),
	}
	return fmt.Sprintf("Group Id Team Primary Channel Tab (%s)", strings.Join(components, "\n"))
}
