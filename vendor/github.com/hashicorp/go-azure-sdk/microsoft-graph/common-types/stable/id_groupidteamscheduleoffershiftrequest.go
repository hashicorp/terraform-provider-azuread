package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamScheduleOfferShiftRequestId{}

// GroupIdTeamScheduleOfferShiftRequestId is a struct representing the Resource ID for a Group Id Team Schedule Offer Shift Request
type GroupIdTeamScheduleOfferShiftRequestId struct {
	GroupId             string
	OfferShiftRequestId string
}

// NewGroupIdTeamScheduleOfferShiftRequestID returns a new GroupIdTeamScheduleOfferShiftRequestId struct
func NewGroupIdTeamScheduleOfferShiftRequestID(groupId string, offerShiftRequestId string) GroupIdTeamScheduleOfferShiftRequestId {
	return GroupIdTeamScheduleOfferShiftRequestId{
		GroupId:             groupId,
		OfferShiftRequestId: offerShiftRequestId,
	}
}

// ParseGroupIdTeamScheduleOfferShiftRequestID parses 'input' into a GroupIdTeamScheduleOfferShiftRequestId
func ParseGroupIdTeamScheduleOfferShiftRequestID(input string) (*GroupIdTeamScheduleOfferShiftRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleOfferShiftRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleOfferShiftRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamScheduleOfferShiftRequestIDInsensitively parses 'input' case-insensitively into a GroupIdTeamScheduleOfferShiftRequestId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamScheduleOfferShiftRequestIDInsensitively(input string) (*GroupIdTeamScheduleOfferShiftRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleOfferShiftRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleOfferShiftRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamScheduleOfferShiftRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.OfferShiftRequestId, ok = input.Parsed["offerShiftRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "offerShiftRequestId", input)
	}

	return nil
}

// ValidateGroupIdTeamScheduleOfferShiftRequestID checks that 'input' can be parsed as a Group Id Team Schedule Offer Shift Request ID
func ValidateGroupIdTeamScheduleOfferShiftRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamScheduleOfferShiftRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Schedule Offer Shift Request ID
func (id GroupIdTeamScheduleOfferShiftRequestId) ID() string {
	fmtString := "/groups/%s/team/schedule/offerShiftRequests/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.OfferShiftRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Schedule Offer Shift Request ID
func (id GroupIdTeamScheduleOfferShiftRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("offerShiftRequests", "offerShiftRequests", "offerShiftRequests"),
		resourceids.UserSpecifiedSegment("offerShiftRequestId", "offerShiftRequestId"),
	}
}

// String returns a human-readable description of this Group Id Team Schedule Offer Shift Request ID
func (id GroupIdTeamScheduleOfferShiftRequestId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Offer Shift Request: %q", id.OfferShiftRequestId),
	}
	return fmt.Sprintf("Group Id Team Schedule Offer Shift Request (%s)", strings.Join(components, "\n"))
}
