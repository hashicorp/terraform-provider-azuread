package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeEventIdExtensionId{}

// MeEventIdExtensionId is a struct representing the Resource ID for a Me Event Id Extension
type MeEventIdExtensionId struct {
	EventId     string
	ExtensionId string
}

// NewMeEventIdExtensionID returns a new MeEventIdExtensionId struct
func NewMeEventIdExtensionID(eventId string, extensionId string) MeEventIdExtensionId {
	return MeEventIdExtensionId{
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseMeEventIdExtensionID parses 'input' into a MeEventIdExtensionId
func ParseMeEventIdExtensionID(input string) (*MeEventIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEventIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEventIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeEventIdExtensionIDInsensitively parses 'input' case-insensitively into a MeEventIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeEventIdExtensionIDInsensitively(input string) (*MeEventIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEventIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEventIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeEventIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateMeEventIdExtensionID checks that 'input' can be parsed as a Me Event Id Extension ID
func ValidateMeEventIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeEventIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Event Id Extension ID
func (id MeEventIdExtensionId) ID() string {
	fmtString := "/me/events/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Event Id Extension ID
func (id MeEventIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Me Event Id Extension ID
func (id MeEventIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Event Id Extension (%s)", strings.Join(components, "\n"))
}
