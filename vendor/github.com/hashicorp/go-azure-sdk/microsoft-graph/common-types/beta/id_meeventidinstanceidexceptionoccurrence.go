package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeEventIdInstanceIdExceptionOccurrenceId{}

// MeEventIdInstanceIdExceptionOccurrenceId is a struct representing the Resource ID for a Me Event Id Instance Id Exception Occurrence
type MeEventIdInstanceIdExceptionOccurrenceId struct {
	EventId  string
	EventId1 string
	EventId2 string
}

// NewMeEventIdInstanceIdExceptionOccurrenceID returns a new MeEventIdInstanceIdExceptionOccurrenceId struct
func NewMeEventIdInstanceIdExceptionOccurrenceID(eventId string, eventId1 string, eventId2 string) MeEventIdInstanceIdExceptionOccurrenceId {
	return MeEventIdInstanceIdExceptionOccurrenceId{
		EventId:  eventId,
		EventId1: eventId1,
		EventId2: eventId2,
	}
}

// ParseMeEventIdInstanceIdExceptionOccurrenceID parses 'input' into a MeEventIdInstanceIdExceptionOccurrenceId
func ParseMeEventIdInstanceIdExceptionOccurrenceID(input string) (*MeEventIdInstanceIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEventIdInstanceIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEventIdInstanceIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeEventIdInstanceIdExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a MeEventIdInstanceIdExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseMeEventIdInstanceIdExceptionOccurrenceIDInsensitively(input string) (*MeEventIdInstanceIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEventIdInstanceIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEventIdInstanceIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeEventIdInstanceIdExceptionOccurrenceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	if id.EventId2, ok = input.Parsed["eventId2"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId2", input)
	}

	return nil
}

// ValidateMeEventIdInstanceIdExceptionOccurrenceID checks that 'input' can be parsed as a Me Event Id Instance Id Exception Occurrence ID
func ValidateMeEventIdInstanceIdExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeEventIdInstanceIdExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Event Id Instance Id Exception Occurrence ID
func (id MeEventIdInstanceIdExceptionOccurrenceId) ID() string {
	fmtString := "/me/events/%s/instances/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.EventId2)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Event Id Instance Id Exception Occurrence ID
func (id MeEventIdInstanceIdExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2"),
	}
}

// String returns a human-readable description of this Me Event Id Instance Id Exception Occurrence ID
func (id MeEventIdInstanceIdExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
	}
	return fmt.Sprintf("Me Event Id Instance Id Exception Occurrence (%s)", strings.Join(components, "\n"))
}
