package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdEventIdAttachmentId{}

// GroupIdEventIdAttachmentId is a struct representing the Resource ID for a Group Id Event Id Attachment
type GroupIdEventIdAttachmentId struct {
	GroupId      string
	EventId      string
	AttachmentId string
}

// NewGroupIdEventIdAttachmentID returns a new GroupIdEventIdAttachmentId struct
func NewGroupIdEventIdAttachmentID(groupId string, eventId string, attachmentId string) GroupIdEventIdAttachmentId {
	return GroupIdEventIdAttachmentId{
		GroupId:      groupId,
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseGroupIdEventIdAttachmentID parses 'input' into a GroupIdEventIdAttachmentId
func ParseGroupIdEventIdAttachmentID(input string) (*GroupIdEventIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdEventIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdEventIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdEventIdAttachmentIDInsensitively parses 'input' case-insensitively into a GroupIdEventIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseGroupIdEventIdAttachmentIDInsensitively(input string) (*GroupIdEventIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdEventIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdEventIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdEventIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdEventIdAttachmentID checks that 'input' can be parsed as a Group Id Event Id Attachment ID
func ValidateGroupIdEventIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdEventIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Event Id Attachment ID
func (id GroupIdEventIdAttachmentId) ID() string {
	fmtString := "/groups/%s/events/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Event Id Attachment ID
func (id GroupIdEventIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentId"),
	}
}

// String returns a human-readable description of this Group Id Event Id Attachment ID
func (id GroupIdEventIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Group Id Event Id Attachment (%s)", strings.Join(components, "\n"))
}
