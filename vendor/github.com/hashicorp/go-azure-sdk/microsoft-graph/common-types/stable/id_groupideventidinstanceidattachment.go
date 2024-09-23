package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdEventIdInstanceIdAttachmentId{}

// GroupIdEventIdInstanceIdAttachmentId is a struct representing the Resource ID for a Group Id Event Id Instance Id Attachment
type GroupIdEventIdInstanceIdAttachmentId struct {
	GroupId      string
	EventId      string
	EventId1     string
	AttachmentId string
}

// NewGroupIdEventIdInstanceIdAttachmentID returns a new GroupIdEventIdInstanceIdAttachmentId struct
func NewGroupIdEventIdInstanceIdAttachmentID(groupId string, eventId string, eventId1 string, attachmentId string) GroupIdEventIdInstanceIdAttachmentId {
	return GroupIdEventIdInstanceIdAttachmentId{
		GroupId:      groupId,
		EventId:      eventId,
		EventId1:     eventId1,
		AttachmentId: attachmentId,
	}
}

// ParseGroupIdEventIdInstanceIdAttachmentID parses 'input' into a GroupIdEventIdInstanceIdAttachmentId
func ParseGroupIdEventIdInstanceIdAttachmentID(input string) (*GroupIdEventIdInstanceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdEventIdInstanceIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdEventIdInstanceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdEventIdInstanceIdAttachmentIDInsensitively parses 'input' case-insensitively into a GroupIdEventIdInstanceIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseGroupIdEventIdInstanceIdAttachmentIDInsensitively(input string) (*GroupIdEventIdInstanceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdEventIdInstanceIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdEventIdInstanceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdEventIdInstanceIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdEventIdInstanceIdAttachmentID checks that 'input' can be parsed as a Group Id Event Id Instance Id Attachment ID
func ValidateGroupIdEventIdInstanceIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdEventIdInstanceIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Event Id Instance Id Attachment ID
func (id GroupIdEventIdInstanceIdAttachmentId) ID() string {
	fmtString := "/groups/%s/events/%s/instances/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Event Id Instance Id Attachment ID
func (id GroupIdEventIdInstanceIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentId"),
	}
}

// String returns a human-readable description of this Group Id Event Id Instance Id Attachment ID
func (id GroupIdEventIdInstanceIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Group Id Event Id Instance Id Attachment (%s)", strings.Join(components, "\n"))
}
