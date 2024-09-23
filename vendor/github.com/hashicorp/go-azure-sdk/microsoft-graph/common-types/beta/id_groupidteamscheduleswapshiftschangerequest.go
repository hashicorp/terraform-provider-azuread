package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamScheduleSwapShiftsChangeRequestId{}

// GroupIdTeamScheduleSwapShiftsChangeRequestId is a struct representing the Resource ID for a Group Id Team Schedule Swap Shifts Change Request
type GroupIdTeamScheduleSwapShiftsChangeRequestId struct {
	GroupId                   string
	SwapShiftsChangeRequestId string
}

// NewGroupIdTeamScheduleSwapShiftsChangeRequestID returns a new GroupIdTeamScheduleSwapShiftsChangeRequestId struct
func NewGroupIdTeamScheduleSwapShiftsChangeRequestID(groupId string, swapShiftsChangeRequestId string) GroupIdTeamScheduleSwapShiftsChangeRequestId {
	return GroupIdTeamScheduleSwapShiftsChangeRequestId{
		GroupId:                   groupId,
		SwapShiftsChangeRequestId: swapShiftsChangeRequestId,
	}
}

// ParseGroupIdTeamScheduleSwapShiftsChangeRequestID parses 'input' into a GroupIdTeamScheduleSwapShiftsChangeRequestId
func ParseGroupIdTeamScheduleSwapShiftsChangeRequestID(input string) (*GroupIdTeamScheduleSwapShiftsChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleSwapShiftsChangeRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleSwapShiftsChangeRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamScheduleSwapShiftsChangeRequestIDInsensitively parses 'input' case-insensitively into a GroupIdTeamScheduleSwapShiftsChangeRequestId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamScheduleSwapShiftsChangeRequestIDInsensitively(input string) (*GroupIdTeamScheduleSwapShiftsChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleSwapShiftsChangeRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleSwapShiftsChangeRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamScheduleSwapShiftsChangeRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SwapShiftsChangeRequestId, ok = input.Parsed["swapShiftsChangeRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "swapShiftsChangeRequestId", input)
	}

	return nil
}

// ValidateGroupIdTeamScheduleSwapShiftsChangeRequestID checks that 'input' can be parsed as a Group Id Team Schedule Swap Shifts Change Request ID
func ValidateGroupIdTeamScheduleSwapShiftsChangeRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamScheduleSwapShiftsChangeRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Schedule Swap Shifts Change Request ID
func (id GroupIdTeamScheduleSwapShiftsChangeRequestId) ID() string {
	fmtString := "/groups/%s/team/schedule/swapShiftsChangeRequests/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SwapShiftsChangeRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Schedule Swap Shifts Change Request ID
func (id GroupIdTeamScheduleSwapShiftsChangeRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("swapShiftsChangeRequests", "swapShiftsChangeRequests", "swapShiftsChangeRequests"),
		resourceids.UserSpecifiedSegment("swapShiftsChangeRequestId", "swapShiftsChangeRequestId"),
	}
}

// String returns a human-readable description of this Group Id Team Schedule Swap Shifts Change Request ID
func (id GroupIdTeamScheduleSwapShiftsChangeRequestId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Swap Shifts Change Request: %q", id.SwapShiftsChangeRequestId),
	}
	return fmt.Sprintf("Group Id Team Schedule Swap Shifts Change Request (%s)", strings.Join(components, "\n"))
}
