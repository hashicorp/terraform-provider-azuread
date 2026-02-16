package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeEventId{}

// MeEventId is a struct representing the Resource ID for a Me Event
type MeEventId struct {
	EventId string
}

// NewMeEventID returns a new MeEventId struct
func NewMeEventID(eventId string) MeEventId {
	return MeEventId{
		EventId: eventId,
	}
}

// ParseMeEventID parses 'input' into a MeEventId
func ParseMeEventID(input string) (*MeEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeEventIDInsensitively parses 'input' case-insensitively into a MeEventId
// note: this method should only be used for API response data and not user input
func ParseMeEventIDInsensitively(input string) (*MeEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeEventId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	return nil
}

// ValidateMeEventID checks that 'input' can be parsed as a Me Event ID
func ValidateMeEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Event ID
func (id MeEventId) ID() string {
	fmtString := "/me/events/%s"
	return fmt.Sprintf(fmtString, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Event ID
func (id MeEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
	}
}

// String returns a human-readable description of this Me Event ID
func (id MeEventId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("Me Event (%s)", strings.Join(components, "\n"))
}
