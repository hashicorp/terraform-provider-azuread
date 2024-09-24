package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeEventIdExceptionOccurrenceId{}

// MeEventIdExceptionOccurrenceId is a struct representing the Resource ID for a Me Event Id Exception Occurrence
type MeEventIdExceptionOccurrenceId struct {
	EventId  string
	EventId1 string
}

// NewMeEventIdExceptionOccurrenceID returns a new MeEventIdExceptionOccurrenceId struct
func NewMeEventIdExceptionOccurrenceID(eventId string, eventId1 string) MeEventIdExceptionOccurrenceId {
	return MeEventIdExceptionOccurrenceId{
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseMeEventIdExceptionOccurrenceID parses 'input' into a MeEventIdExceptionOccurrenceId
func ParseMeEventIdExceptionOccurrenceID(input string) (*MeEventIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEventIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEventIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeEventIdExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a MeEventIdExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseMeEventIdExceptionOccurrenceIDInsensitively(input string) (*MeEventIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEventIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEventIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeEventIdExceptionOccurrenceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	return nil
}

// ValidateMeEventIdExceptionOccurrenceID checks that 'input' can be parsed as a Me Event Id Exception Occurrence ID
func ValidateMeEventIdExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeEventIdExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Event Id Exception Occurrence ID
func (id MeEventIdExceptionOccurrenceId) ID() string {
	fmtString := "/me/events/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Event Id Exception Occurrence ID
func (id MeEventIdExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1"),
	}
}

// String returns a human-readable description of this Me Event Id Exception Occurrence ID
func (id MeEventIdExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Me Event Id Exception Occurrence (%s)", strings.Join(components, "\n"))
}
