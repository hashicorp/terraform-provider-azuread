package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdCalendarCalendarViewIdAttachmentId{}

// GroupIdCalendarCalendarViewIdAttachmentId is a struct representing the Resource ID for a Group Id Calendar Calendar View Id Attachment
type GroupIdCalendarCalendarViewIdAttachmentId struct {
	GroupId      string
	EventId      string
	AttachmentId string
}

// NewGroupIdCalendarCalendarViewIdAttachmentID returns a new GroupIdCalendarCalendarViewIdAttachmentId struct
func NewGroupIdCalendarCalendarViewIdAttachmentID(groupId string, eventId string, attachmentId string) GroupIdCalendarCalendarViewIdAttachmentId {
	return GroupIdCalendarCalendarViewIdAttachmentId{
		GroupId:      groupId,
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseGroupIdCalendarCalendarViewIdAttachmentID parses 'input' into a GroupIdCalendarCalendarViewIdAttachmentId
func ParseGroupIdCalendarCalendarViewIdAttachmentID(input string) (*GroupIdCalendarCalendarViewIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarCalendarViewIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarCalendarViewIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdCalendarCalendarViewIdAttachmentIDInsensitively parses 'input' case-insensitively into a GroupIdCalendarCalendarViewIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseGroupIdCalendarCalendarViewIdAttachmentIDInsensitively(input string) (*GroupIdCalendarCalendarViewIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarCalendarViewIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarCalendarViewIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdCalendarCalendarViewIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateGroupIdCalendarCalendarViewIdAttachmentID checks that 'input' can be parsed as a Group Id Calendar Calendar View Id Attachment ID
func ValidateGroupIdCalendarCalendarViewIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdCalendarCalendarViewIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Calendar Calendar View Id Attachment ID
func (id GroupIdCalendarCalendarViewIdAttachmentId) ID() string {
	fmtString := "/groups/%s/calendar/calendarView/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Calendar Calendar View Id Attachment ID
func (id GroupIdCalendarCalendarViewIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentId"),
	}
}

// String returns a human-readable description of this Group Id Calendar Calendar View Id Attachment ID
func (id GroupIdCalendarCalendarViewIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Group Id Calendar Calendar View Id Attachment (%s)", strings.Join(components, "\n"))
}
