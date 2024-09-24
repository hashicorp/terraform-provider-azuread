package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamScheduleDayNoteId{}

// GroupIdTeamScheduleDayNoteId is a struct representing the Resource ID for a Group Id Team Schedule Day Note
type GroupIdTeamScheduleDayNoteId struct {
	GroupId   string
	DayNoteId string
}

// NewGroupIdTeamScheduleDayNoteID returns a new GroupIdTeamScheduleDayNoteId struct
func NewGroupIdTeamScheduleDayNoteID(groupId string, dayNoteId string) GroupIdTeamScheduleDayNoteId {
	return GroupIdTeamScheduleDayNoteId{
		GroupId:   groupId,
		DayNoteId: dayNoteId,
	}
}

// ParseGroupIdTeamScheduleDayNoteID parses 'input' into a GroupIdTeamScheduleDayNoteId
func ParseGroupIdTeamScheduleDayNoteID(input string) (*GroupIdTeamScheduleDayNoteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleDayNoteId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleDayNoteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamScheduleDayNoteIDInsensitively parses 'input' case-insensitively into a GroupIdTeamScheduleDayNoteId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamScheduleDayNoteIDInsensitively(input string) (*GroupIdTeamScheduleDayNoteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleDayNoteId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleDayNoteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamScheduleDayNoteId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DayNoteId, ok = input.Parsed["dayNoteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dayNoteId", input)
	}

	return nil
}

// ValidateGroupIdTeamScheduleDayNoteID checks that 'input' can be parsed as a Group Id Team Schedule Day Note ID
func ValidateGroupIdTeamScheduleDayNoteID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamScheduleDayNoteID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Schedule Day Note ID
func (id GroupIdTeamScheduleDayNoteId) ID() string {
	fmtString := "/groups/%s/team/schedule/dayNotes/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DayNoteId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Schedule Day Note ID
func (id GroupIdTeamScheduleDayNoteId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("dayNotes", "dayNotes", "dayNotes"),
		resourceids.UserSpecifiedSegment("dayNoteId", "dayNoteId"),
	}
}

// String returns a human-readable description of this Group Id Team Schedule Day Note ID
func (id GroupIdTeamScheduleDayNoteId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Day Note: %q", id.DayNoteId),
	}
	return fmt.Sprintf("Group Id Team Schedule Day Note (%s)", strings.Join(components, "\n"))
}
