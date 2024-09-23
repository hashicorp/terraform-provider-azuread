package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeEventIdInstanceId{}

// MeEventIdInstanceId is a struct representing the Resource ID for a Me Event Id Instance
type MeEventIdInstanceId struct {
	EventId  string
	EventId1 string
}

// NewMeEventIdInstanceID returns a new MeEventIdInstanceId struct
func NewMeEventIdInstanceID(eventId string, eventId1 string) MeEventIdInstanceId {
	return MeEventIdInstanceId{
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseMeEventIdInstanceID parses 'input' into a MeEventIdInstanceId
func ParseMeEventIdInstanceID(input string) (*MeEventIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEventIdInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEventIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeEventIdInstanceIDInsensitively parses 'input' case-insensitively into a MeEventIdInstanceId
// note: this method should only be used for API response data and not user input
func ParseMeEventIdInstanceIDInsensitively(input string) (*MeEventIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEventIdInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEventIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeEventIdInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	return nil
}

// ValidateMeEventIdInstanceID checks that 'input' can be parsed as a Me Event Id Instance ID
func ValidateMeEventIdInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeEventIdInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Event Id Instance ID
func (id MeEventIdInstanceId) ID() string {
	fmtString := "/me/events/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Event Id Instance ID
func (id MeEventIdInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1"),
	}
}

// String returns a human-readable description of this Me Event Id Instance ID
func (id MeEventIdInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Me Event Id Instance (%s)", strings.Join(components, "\n"))
}
