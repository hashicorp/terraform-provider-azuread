package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamScheduleTimeCardId{}

// GroupIdTeamScheduleTimeCardId is a struct representing the Resource ID for a Group Id Team Schedule Time Card
type GroupIdTeamScheduleTimeCardId struct {
	GroupId    string
	TimeCardId string
}

// NewGroupIdTeamScheduleTimeCardID returns a new GroupIdTeamScheduleTimeCardId struct
func NewGroupIdTeamScheduleTimeCardID(groupId string, timeCardId string) GroupIdTeamScheduleTimeCardId {
	return GroupIdTeamScheduleTimeCardId{
		GroupId:    groupId,
		TimeCardId: timeCardId,
	}
}

// ParseGroupIdTeamScheduleTimeCardID parses 'input' into a GroupIdTeamScheduleTimeCardId
func ParseGroupIdTeamScheduleTimeCardID(input string) (*GroupIdTeamScheduleTimeCardId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleTimeCardId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleTimeCardId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamScheduleTimeCardIDInsensitively parses 'input' case-insensitively into a GroupIdTeamScheduleTimeCardId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamScheduleTimeCardIDInsensitively(input string) (*GroupIdTeamScheduleTimeCardId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleTimeCardId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleTimeCardId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamScheduleTimeCardId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.TimeCardId, ok = input.Parsed["timeCardId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "timeCardId", input)
	}

	return nil
}

// ValidateGroupIdTeamScheduleTimeCardID checks that 'input' can be parsed as a Group Id Team Schedule Time Card ID
func ValidateGroupIdTeamScheduleTimeCardID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamScheduleTimeCardID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Schedule Time Card ID
func (id GroupIdTeamScheduleTimeCardId) ID() string {
	fmtString := "/groups/%s/team/schedule/timeCards/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.TimeCardId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Schedule Time Card ID
func (id GroupIdTeamScheduleTimeCardId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("timeCards", "timeCards", "timeCards"),
		resourceids.UserSpecifiedSegment("timeCardId", "timeCardId"),
	}
}

// String returns a human-readable description of this Group Id Team Schedule Time Card ID
func (id GroupIdTeamScheduleTimeCardId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Time Card: %q", id.TimeCardId),
	}
	return fmt.Sprintf("Group Id Team Schedule Time Card (%s)", strings.Join(components, "\n"))
}
