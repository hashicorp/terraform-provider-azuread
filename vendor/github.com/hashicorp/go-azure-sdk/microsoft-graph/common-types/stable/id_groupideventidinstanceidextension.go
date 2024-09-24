package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdEventIdInstanceIdExtensionId{}

// GroupIdEventIdInstanceIdExtensionId is a struct representing the Resource ID for a Group Id Event Id Instance Id Extension
type GroupIdEventIdInstanceIdExtensionId struct {
	GroupId     string
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewGroupIdEventIdInstanceIdExtensionID returns a new GroupIdEventIdInstanceIdExtensionId struct
func NewGroupIdEventIdInstanceIdExtensionID(groupId string, eventId string, eventId1 string, extensionId string) GroupIdEventIdInstanceIdExtensionId {
	return GroupIdEventIdInstanceIdExtensionId{
		GroupId:     groupId,
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseGroupIdEventIdInstanceIdExtensionID parses 'input' into a GroupIdEventIdInstanceIdExtensionId
func ParseGroupIdEventIdInstanceIdExtensionID(input string) (*GroupIdEventIdInstanceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdEventIdInstanceIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdEventIdInstanceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdEventIdInstanceIdExtensionIDInsensitively parses 'input' case-insensitively into a GroupIdEventIdInstanceIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdEventIdInstanceIdExtensionIDInsensitively(input string) (*GroupIdEventIdInstanceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdEventIdInstanceIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdEventIdInstanceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdEventIdInstanceIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateGroupIdEventIdInstanceIdExtensionID checks that 'input' can be parsed as a Group Id Event Id Instance Id Extension ID
func ValidateGroupIdEventIdInstanceIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdEventIdInstanceIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Event Id Instance Id Extension ID
func (id GroupIdEventIdInstanceIdExtensionId) ID() string {
	fmtString := "/groups/%s/events/%s/instances/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Event Id Instance Id Extension ID
func (id GroupIdEventIdInstanceIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Group Id Event Id Instance Id Extension ID
func (id GroupIdEventIdInstanceIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Id Event Id Instance Id Extension (%s)", strings.Join(components, "\n"))
}
