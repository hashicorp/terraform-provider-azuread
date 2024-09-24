package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamScheduleTimeOffReasonId{}

// GroupIdTeamScheduleTimeOffReasonId is a struct representing the Resource ID for a Group Id Team Schedule Time Off Reason
type GroupIdTeamScheduleTimeOffReasonId struct {
	GroupId         string
	TimeOffReasonId string
}

// NewGroupIdTeamScheduleTimeOffReasonID returns a new GroupIdTeamScheduleTimeOffReasonId struct
func NewGroupIdTeamScheduleTimeOffReasonID(groupId string, timeOffReasonId string) GroupIdTeamScheduleTimeOffReasonId {
	return GroupIdTeamScheduleTimeOffReasonId{
		GroupId:         groupId,
		TimeOffReasonId: timeOffReasonId,
	}
}

// ParseGroupIdTeamScheduleTimeOffReasonID parses 'input' into a GroupIdTeamScheduleTimeOffReasonId
func ParseGroupIdTeamScheduleTimeOffReasonID(input string) (*GroupIdTeamScheduleTimeOffReasonId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleTimeOffReasonId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleTimeOffReasonId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamScheduleTimeOffReasonIDInsensitively parses 'input' case-insensitively into a GroupIdTeamScheduleTimeOffReasonId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamScheduleTimeOffReasonIDInsensitively(input string) (*GroupIdTeamScheduleTimeOffReasonId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleTimeOffReasonId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleTimeOffReasonId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamScheduleTimeOffReasonId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.TimeOffReasonId, ok = input.Parsed["timeOffReasonId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "timeOffReasonId", input)
	}

	return nil
}

// ValidateGroupIdTeamScheduleTimeOffReasonID checks that 'input' can be parsed as a Group Id Team Schedule Time Off Reason ID
func ValidateGroupIdTeamScheduleTimeOffReasonID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamScheduleTimeOffReasonID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Schedule Time Off Reason ID
func (id GroupIdTeamScheduleTimeOffReasonId) ID() string {
	fmtString := "/groups/%s/team/schedule/timeOffReasons/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.TimeOffReasonId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Schedule Time Off Reason ID
func (id GroupIdTeamScheduleTimeOffReasonId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("timeOffReasons", "timeOffReasons", "timeOffReasons"),
		resourceids.UserSpecifiedSegment("timeOffReasonId", "timeOffReasonId"),
	}
}

// String returns a human-readable description of this Group Id Team Schedule Time Off Reason ID
func (id GroupIdTeamScheduleTimeOffReasonId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Time Off Reason: %q", id.TimeOffReasonId),
	}
	return fmt.Sprintf("Group Id Team Schedule Time Off Reason (%s)", strings.Join(components, "\n"))
}
