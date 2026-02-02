package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdEventIdExtensionId{}

// GroupIdEventIdExtensionId is a struct representing the Resource ID for a Group Id Event Id Extension
type GroupIdEventIdExtensionId struct {
	GroupId     string
	EventId     string
	ExtensionId string
}

// NewGroupIdEventIdExtensionID returns a new GroupIdEventIdExtensionId struct
func NewGroupIdEventIdExtensionID(groupId string, eventId string, extensionId string) GroupIdEventIdExtensionId {
	return GroupIdEventIdExtensionId{
		GroupId:     groupId,
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseGroupIdEventIdExtensionID parses 'input' into a GroupIdEventIdExtensionId
func ParseGroupIdEventIdExtensionID(input string) (*GroupIdEventIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdEventIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdEventIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdEventIdExtensionIDInsensitively parses 'input' case-insensitively into a GroupIdEventIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdEventIdExtensionIDInsensitively(input string) (*GroupIdEventIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdEventIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdEventIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdEventIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateGroupIdEventIdExtensionID checks that 'input' can be parsed as a Group Id Event Id Extension ID
func ValidateGroupIdEventIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdEventIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Event Id Extension ID
func (id GroupIdEventIdExtensionId) ID() string {
	fmtString := "/groups/%s/events/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Event Id Extension ID
func (id GroupIdEventIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Group Id Event Id Extension ID
func (id GroupIdEventIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Id Event Id Extension (%s)", strings.Join(components, "\n"))
}
