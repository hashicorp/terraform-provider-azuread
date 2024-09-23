package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamScheduleTimesOffId{}

// GroupIdTeamScheduleTimesOffId is a struct representing the Resource ID for a Group Id Team Schedule Times Off
type GroupIdTeamScheduleTimesOffId struct {
	GroupId   string
	TimeOffId string
}

// NewGroupIdTeamScheduleTimesOffID returns a new GroupIdTeamScheduleTimesOffId struct
func NewGroupIdTeamScheduleTimesOffID(groupId string, timeOffId string) GroupIdTeamScheduleTimesOffId {
	return GroupIdTeamScheduleTimesOffId{
		GroupId:   groupId,
		TimeOffId: timeOffId,
	}
}

// ParseGroupIdTeamScheduleTimesOffID parses 'input' into a GroupIdTeamScheduleTimesOffId
func ParseGroupIdTeamScheduleTimesOffID(input string) (*GroupIdTeamScheduleTimesOffId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleTimesOffId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleTimesOffId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamScheduleTimesOffIDInsensitively parses 'input' case-insensitively into a GroupIdTeamScheduleTimesOffId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamScheduleTimesOffIDInsensitively(input string) (*GroupIdTeamScheduleTimesOffId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleTimesOffId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleTimesOffId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamScheduleTimesOffId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.TimeOffId, ok = input.Parsed["timeOffId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "timeOffId", input)
	}

	return nil
}

// ValidateGroupIdTeamScheduleTimesOffID checks that 'input' can be parsed as a Group Id Team Schedule Times Off ID
func ValidateGroupIdTeamScheduleTimesOffID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamScheduleTimesOffID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Schedule Times Off ID
func (id GroupIdTeamScheduleTimesOffId) ID() string {
	fmtString := "/groups/%s/team/schedule/timesOff/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.TimeOffId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Schedule Times Off ID
func (id GroupIdTeamScheduleTimesOffId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("timesOff", "timesOff", "timesOff"),
		resourceids.UserSpecifiedSegment("timeOffId", "timeOffId"),
	}
}

// String returns a human-readable description of this Group Id Team Schedule Times Off ID
func (id GroupIdTeamScheduleTimesOffId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Time Off: %q", id.TimeOffId),
	}
	return fmt.Sprintf("Group Id Team Schedule Times Off (%s)", strings.Join(components, "\n"))
}
