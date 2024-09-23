package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdEventId{}

// GroupIdEventId is a struct representing the Resource ID for a Group Id Event
type GroupIdEventId struct {
	GroupId string
	EventId string
}

// NewGroupIdEventID returns a new GroupIdEventId struct
func NewGroupIdEventID(groupId string, eventId string) GroupIdEventId {
	return GroupIdEventId{
		GroupId: groupId,
		EventId: eventId,
	}
}

// ParseGroupIdEventID parses 'input' into a GroupIdEventId
func ParseGroupIdEventID(input string) (*GroupIdEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdEventIDInsensitively parses 'input' case-insensitively into a GroupIdEventId
// note: this method should only be used for API response data and not user input
func ParseGroupIdEventIDInsensitively(input string) (*GroupIdEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdEventId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	return nil
}

// ValidateGroupIdEventID checks that 'input' can be parsed as a Group Id Event ID
func ValidateGroupIdEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Event ID
func (id GroupIdEventId) ID() string {
	fmtString := "/groups/%s/events/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Event ID
func (id GroupIdEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
	}
}

// String returns a human-readable description of this Group Id Event ID
func (id GroupIdEventId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("Group Id Event (%s)", strings.Join(components, "\n"))
}
