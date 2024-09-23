package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdEventIdInstanceId{}

// GroupIdEventIdInstanceId is a struct representing the Resource ID for a Group Id Event Id Instance
type GroupIdEventIdInstanceId struct {
	GroupId  string
	EventId  string
	EventId1 string
}

// NewGroupIdEventIdInstanceID returns a new GroupIdEventIdInstanceId struct
func NewGroupIdEventIdInstanceID(groupId string, eventId string, eventId1 string) GroupIdEventIdInstanceId {
	return GroupIdEventIdInstanceId{
		GroupId:  groupId,
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseGroupIdEventIdInstanceID parses 'input' into a GroupIdEventIdInstanceId
func ParseGroupIdEventIdInstanceID(input string) (*GroupIdEventIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdEventIdInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdEventIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdEventIdInstanceIDInsensitively parses 'input' case-insensitively into a GroupIdEventIdInstanceId
// note: this method should only be used for API response data and not user input
func ParseGroupIdEventIdInstanceIDInsensitively(input string) (*GroupIdEventIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdEventIdInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdEventIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdEventIdInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	return nil
}

// ValidateGroupIdEventIdInstanceID checks that 'input' can be parsed as a Group Id Event Id Instance ID
func ValidateGroupIdEventIdInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdEventIdInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Event Id Instance ID
func (id GroupIdEventIdInstanceId) ID() string {
	fmtString := "/groups/%s/events/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Event Id Instance ID
func (id GroupIdEventIdInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1"),
	}
}

// String returns a human-readable description of this Group Id Event Id Instance ID
func (id GroupIdEventIdInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Group Id Event Id Instance (%s)", strings.Join(components, "\n"))
}
