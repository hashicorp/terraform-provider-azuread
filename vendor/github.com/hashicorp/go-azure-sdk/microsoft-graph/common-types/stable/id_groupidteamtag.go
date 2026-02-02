package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamTagId{}

// GroupIdTeamTagId is a struct representing the Resource ID for a Group Id Team Tag
type GroupIdTeamTagId struct {
	GroupId       string
	TeamworkTagId string
}

// NewGroupIdTeamTagID returns a new GroupIdTeamTagId struct
func NewGroupIdTeamTagID(groupId string, teamworkTagId string) GroupIdTeamTagId {
	return GroupIdTeamTagId{
		GroupId:       groupId,
		TeamworkTagId: teamworkTagId,
	}
}

// ParseGroupIdTeamTagID parses 'input' into a GroupIdTeamTagId
func ParseGroupIdTeamTagID(input string) (*GroupIdTeamTagId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamTagId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamTagId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamTagIDInsensitively parses 'input' case-insensitively into a GroupIdTeamTagId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamTagIDInsensitively(input string) (*GroupIdTeamTagId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamTagId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamTagId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamTagId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.TeamworkTagId, ok = input.Parsed["teamworkTagId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagId", input)
	}

	return nil
}

// ValidateGroupIdTeamTagID checks that 'input' can be parsed as a Group Id Team Tag ID
func ValidateGroupIdTeamTagID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamTagID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Tag ID
func (id GroupIdTeamTagId) ID() string {
	fmtString := "/groups/%s/team/tags/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.TeamworkTagId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Tag ID
func (id GroupIdTeamTagId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("tags", "tags", "tags"),
		resourceids.UserSpecifiedSegment("teamworkTagId", "teamworkTagId"),
	}
}

// String returns a human-readable description of this Group Id Team Tag ID
func (id GroupIdTeamTagId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Teamwork Tag: %q", id.TeamworkTagId),
	}
	return fmt.Sprintf("Group Id Team Tag (%s)", strings.Join(components, "\n"))
}
