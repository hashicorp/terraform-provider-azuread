package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamScheduleTimeOffRequestId{}

// GroupIdTeamScheduleTimeOffRequestId is a struct representing the Resource ID for a Group Id Team Schedule Time Off Request
type GroupIdTeamScheduleTimeOffRequestId struct {
	GroupId          string
	TimeOffRequestId string
}

// NewGroupIdTeamScheduleTimeOffRequestID returns a new GroupIdTeamScheduleTimeOffRequestId struct
func NewGroupIdTeamScheduleTimeOffRequestID(groupId string, timeOffRequestId string) GroupIdTeamScheduleTimeOffRequestId {
	return GroupIdTeamScheduleTimeOffRequestId{
		GroupId:          groupId,
		TimeOffRequestId: timeOffRequestId,
	}
}

// ParseGroupIdTeamScheduleTimeOffRequestID parses 'input' into a GroupIdTeamScheduleTimeOffRequestId
func ParseGroupIdTeamScheduleTimeOffRequestID(input string) (*GroupIdTeamScheduleTimeOffRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleTimeOffRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleTimeOffRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamScheduleTimeOffRequestIDInsensitively parses 'input' case-insensitively into a GroupIdTeamScheduleTimeOffRequestId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamScheduleTimeOffRequestIDInsensitively(input string) (*GroupIdTeamScheduleTimeOffRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleTimeOffRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleTimeOffRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamScheduleTimeOffRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.TimeOffRequestId, ok = input.Parsed["timeOffRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "timeOffRequestId", input)
	}

	return nil
}

// ValidateGroupIdTeamScheduleTimeOffRequestID checks that 'input' can be parsed as a Group Id Team Schedule Time Off Request ID
func ValidateGroupIdTeamScheduleTimeOffRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamScheduleTimeOffRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Schedule Time Off Request ID
func (id GroupIdTeamScheduleTimeOffRequestId) ID() string {
	fmtString := "/groups/%s/team/schedule/timeOffRequests/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.TimeOffRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Schedule Time Off Request ID
func (id GroupIdTeamScheduleTimeOffRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("timeOffRequests", "timeOffRequests", "timeOffRequests"),
		resourceids.UserSpecifiedSegment("timeOffRequestId", "timeOffRequestId"),
	}
}

// String returns a human-readable description of this Group Id Team Schedule Time Off Request ID
func (id GroupIdTeamScheduleTimeOffRequestId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Time Off Request: %q", id.TimeOffRequestId),
	}
	return fmt.Sprintf("Group Id Team Schedule Time Off Request (%s)", strings.Join(components, "\n"))
}
