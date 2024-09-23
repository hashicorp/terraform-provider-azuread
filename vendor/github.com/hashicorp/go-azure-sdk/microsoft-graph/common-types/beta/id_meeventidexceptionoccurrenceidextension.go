package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeEventIdExceptionOccurrenceIdExtensionId{}

// MeEventIdExceptionOccurrenceIdExtensionId is a struct representing the Resource ID for a Me Event Id Exception Occurrence Id Extension
type MeEventIdExceptionOccurrenceIdExtensionId struct {
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewMeEventIdExceptionOccurrenceIdExtensionID returns a new MeEventIdExceptionOccurrenceIdExtensionId struct
func NewMeEventIdExceptionOccurrenceIdExtensionID(eventId string, eventId1 string, extensionId string) MeEventIdExceptionOccurrenceIdExtensionId {
	return MeEventIdExceptionOccurrenceIdExtensionId{
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseMeEventIdExceptionOccurrenceIdExtensionID parses 'input' into a MeEventIdExceptionOccurrenceIdExtensionId
func ParseMeEventIdExceptionOccurrenceIdExtensionID(input string) (*MeEventIdExceptionOccurrenceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEventIdExceptionOccurrenceIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEventIdExceptionOccurrenceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeEventIdExceptionOccurrenceIdExtensionIDInsensitively parses 'input' case-insensitively into a MeEventIdExceptionOccurrenceIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeEventIdExceptionOccurrenceIdExtensionIDInsensitively(input string) (*MeEventIdExceptionOccurrenceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEventIdExceptionOccurrenceIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEventIdExceptionOccurrenceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeEventIdExceptionOccurrenceIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateMeEventIdExceptionOccurrenceIdExtensionID checks that 'input' can be parsed as a Me Event Id Exception Occurrence Id Extension ID
func ValidateMeEventIdExceptionOccurrenceIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeEventIdExceptionOccurrenceIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Event Id Exception Occurrence Id Extension ID
func (id MeEventIdExceptionOccurrenceIdExtensionId) ID() string {
	fmtString := "/me/events/%s/exceptionOccurrences/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Event Id Exception Occurrence Id Extension ID
func (id MeEventIdExceptionOccurrenceIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Me Event Id Exception Occurrence Id Extension ID
func (id MeEventIdExceptionOccurrenceIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Event Id Exception Occurrence Id Extension (%s)", strings.Join(components, "\n"))
}
