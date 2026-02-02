package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamScheduleOpenShiftChangeRequestId{}

// GroupIdTeamScheduleOpenShiftChangeRequestId is a struct representing the Resource ID for a Group Id Team Schedule Open Shift Change Request
type GroupIdTeamScheduleOpenShiftChangeRequestId struct {
	GroupId                  string
	OpenShiftChangeRequestId string
}

// NewGroupIdTeamScheduleOpenShiftChangeRequestID returns a new GroupIdTeamScheduleOpenShiftChangeRequestId struct
func NewGroupIdTeamScheduleOpenShiftChangeRequestID(groupId string, openShiftChangeRequestId string) GroupIdTeamScheduleOpenShiftChangeRequestId {
	return GroupIdTeamScheduleOpenShiftChangeRequestId{
		GroupId:                  groupId,
		OpenShiftChangeRequestId: openShiftChangeRequestId,
	}
}

// ParseGroupIdTeamScheduleOpenShiftChangeRequestID parses 'input' into a GroupIdTeamScheduleOpenShiftChangeRequestId
func ParseGroupIdTeamScheduleOpenShiftChangeRequestID(input string) (*GroupIdTeamScheduleOpenShiftChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleOpenShiftChangeRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleOpenShiftChangeRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamScheduleOpenShiftChangeRequestIDInsensitively parses 'input' case-insensitively into a GroupIdTeamScheduleOpenShiftChangeRequestId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamScheduleOpenShiftChangeRequestIDInsensitively(input string) (*GroupIdTeamScheduleOpenShiftChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleOpenShiftChangeRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleOpenShiftChangeRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamScheduleOpenShiftChangeRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.OpenShiftChangeRequestId, ok = input.Parsed["openShiftChangeRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "openShiftChangeRequestId", input)
	}

	return nil
}

// ValidateGroupIdTeamScheduleOpenShiftChangeRequestID checks that 'input' can be parsed as a Group Id Team Schedule Open Shift Change Request ID
func ValidateGroupIdTeamScheduleOpenShiftChangeRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamScheduleOpenShiftChangeRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Schedule Open Shift Change Request ID
func (id GroupIdTeamScheduleOpenShiftChangeRequestId) ID() string {
	fmtString := "/groups/%s/team/schedule/openShiftChangeRequests/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.OpenShiftChangeRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Schedule Open Shift Change Request ID
func (id GroupIdTeamScheduleOpenShiftChangeRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("openShiftChangeRequests", "openShiftChangeRequests", "openShiftChangeRequests"),
		resourceids.UserSpecifiedSegment("openShiftChangeRequestId", "openShiftChangeRequestId"),
	}
}

// String returns a human-readable description of this Group Id Team Schedule Open Shift Change Request ID
func (id GroupIdTeamScheduleOpenShiftChangeRequestId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Open Shift Change Request: %q", id.OpenShiftChangeRequestId),
	}
	return fmt.Sprintf("Group Id Team Schedule Open Shift Change Request (%s)", strings.Join(components, "\n"))
}
