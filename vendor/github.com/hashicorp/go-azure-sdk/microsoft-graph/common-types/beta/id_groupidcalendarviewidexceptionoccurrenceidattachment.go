package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdCalendarViewIdExceptionOccurrenceIdAttachmentId{}

// GroupIdCalendarViewIdExceptionOccurrenceIdAttachmentId is a struct representing the Resource ID for a Group Id Calendar View Id Exception Occurrence Id Attachment
type GroupIdCalendarViewIdExceptionOccurrenceIdAttachmentId struct {
	GroupId      string
	EventId      string
	EventId1     string
	AttachmentId string
}

// NewGroupIdCalendarViewIdExceptionOccurrenceIdAttachmentID returns a new GroupIdCalendarViewIdExceptionOccurrenceIdAttachmentId struct
func NewGroupIdCalendarViewIdExceptionOccurrenceIdAttachmentID(groupId string, eventId string, eventId1 string, attachmentId string) GroupIdCalendarViewIdExceptionOccurrenceIdAttachmentId {
	return GroupIdCalendarViewIdExceptionOccurrenceIdAttachmentId{
		GroupId:      groupId,
		EventId:      eventId,
		EventId1:     eventId1,
		AttachmentId: attachmentId,
	}
}

// ParseGroupIdCalendarViewIdExceptionOccurrenceIdAttachmentID parses 'input' into a GroupIdCalendarViewIdExceptionOccurrenceIdAttachmentId
func ParseGroupIdCalendarViewIdExceptionOccurrenceIdAttachmentID(input string) (*GroupIdCalendarViewIdExceptionOccurrenceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarViewIdExceptionOccurrenceIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarViewIdExceptionOccurrenceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdCalendarViewIdExceptionOccurrenceIdAttachmentIDInsensitively parses 'input' case-insensitively into a GroupIdCalendarViewIdExceptionOccurrenceIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseGroupIdCalendarViewIdExceptionOccurrenceIdAttachmentIDInsensitively(input string) (*GroupIdCalendarViewIdExceptionOccurrenceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarViewIdExceptionOccurrenceIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarViewIdExceptionOccurrenceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdCalendarViewIdExceptionOccurrenceIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateGroupIdCalendarViewIdExceptionOccurrenceIdAttachmentID checks that 'input' can be parsed as a Group Id Calendar View Id Exception Occurrence Id Attachment ID
func ValidateGroupIdCalendarViewIdExceptionOccurrenceIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdCalendarViewIdExceptionOccurrenceIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Calendar View Id Exception Occurrence Id Attachment ID
func (id GroupIdCalendarViewIdExceptionOccurrenceIdAttachmentId) ID() string {
	fmtString := "/groups/%s/calendarView/%s/exceptionOccurrences/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Calendar View Id Exception Occurrence Id Attachment ID
func (id GroupIdCalendarViewIdExceptionOccurrenceIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentId"),
	}
}

// String returns a human-readable description of this Group Id Calendar View Id Exception Occurrence Id Attachment ID
func (id GroupIdCalendarViewIdExceptionOccurrenceIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Group Id Calendar View Id Exception Occurrence Id Attachment (%s)", strings.Join(components, "\n"))
}
