package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeEventIdInstanceIdAttachmentId{}

// MeEventIdInstanceIdAttachmentId is a struct representing the Resource ID for a Me Event Id Instance Id Attachment
type MeEventIdInstanceIdAttachmentId struct {
	EventId      string
	EventId1     string
	AttachmentId string
}

// NewMeEventIdInstanceIdAttachmentID returns a new MeEventIdInstanceIdAttachmentId struct
func NewMeEventIdInstanceIdAttachmentID(eventId string, eventId1 string, attachmentId string) MeEventIdInstanceIdAttachmentId {
	return MeEventIdInstanceIdAttachmentId{
		EventId:      eventId,
		EventId1:     eventId1,
		AttachmentId: attachmentId,
	}
}

// ParseMeEventIdInstanceIdAttachmentID parses 'input' into a MeEventIdInstanceIdAttachmentId
func ParseMeEventIdInstanceIdAttachmentID(input string) (*MeEventIdInstanceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEventIdInstanceIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEventIdInstanceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeEventIdInstanceIdAttachmentIDInsensitively parses 'input' case-insensitively into a MeEventIdInstanceIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeEventIdInstanceIdAttachmentIDInsensitively(input string) (*MeEventIdInstanceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEventIdInstanceIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEventIdInstanceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeEventIdInstanceIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

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

// ValidateMeEventIdInstanceIdAttachmentID checks that 'input' can be parsed as a Me Event Id Instance Id Attachment ID
func ValidateMeEventIdInstanceIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeEventIdInstanceIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Event Id Instance Id Attachment ID
func (id MeEventIdInstanceIdAttachmentId) ID() string {
	fmtString := "/me/events/%s/instances/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Event Id Instance Id Attachment ID
func (id MeEventIdInstanceIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentId"),
	}
}

// String returns a human-readable description of this Me Event Id Instance Id Attachment ID
func (id MeEventIdInstanceIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Event Id Instance Id Attachment (%s)", strings.Join(components, "\n"))
}
