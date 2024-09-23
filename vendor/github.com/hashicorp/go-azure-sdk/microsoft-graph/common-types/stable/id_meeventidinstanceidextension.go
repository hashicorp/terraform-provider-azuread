package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeEventIdInstanceIdExtensionId{}

// MeEventIdInstanceIdExtensionId is a struct representing the Resource ID for a Me Event Id Instance Id Extension
type MeEventIdInstanceIdExtensionId struct {
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewMeEventIdInstanceIdExtensionID returns a new MeEventIdInstanceIdExtensionId struct
func NewMeEventIdInstanceIdExtensionID(eventId string, eventId1 string, extensionId string) MeEventIdInstanceIdExtensionId {
	return MeEventIdInstanceIdExtensionId{
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseMeEventIdInstanceIdExtensionID parses 'input' into a MeEventIdInstanceIdExtensionId
func ParseMeEventIdInstanceIdExtensionID(input string) (*MeEventIdInstanceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEventIdInstanceIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEventIdInstanceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeEventIdInstanceIdExtensionIDInsensitively parses 'input' case-insensitively into a MeEventIdInstanceIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeEventIdInstanceIdExtensionIDInsensitively(input string) (*MeEventIdInstanceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEventIdInstanceIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEventIdInstanceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeEventIdInstanceIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateMeEventIdInstanceIdExtensionID checks that 'input' can be parsed as a Me Event Id Instance Id Extension ID
func ValidateMeEventIdInstanceIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeEventIdInstanceIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Event Id Instance Id Extension ID
func (id MeEventIdInstanceIdExtensionId) ID() string {
	fmtString := "/me/events/%s/instances/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Event Id Instance Id Extension ID
func (id MeEventIdInstanceIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Me Event Id Instance Id Extension ID
func (id MeEventIdInstanceIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Event Id Instance Id Extension (%s)", strings.Join(components, "\n"))
}
