package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeEventIdAttachmentId{}

// MeEventIdAttachmentId is a struct representing the Resource ID for a Me Event Id Attachment
type MeEventIdAttachmentId struct {
	EventId      string
	AttachmentId string
}

// NewMeEventIdAttachmentID returns a new MeEventIdAttachmentId struct
func NewMeEventIdAttachmentID(eventId string, attachmentId string) MeEventIdAttachmentId {
	return MeEventIdAttachmentId{
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseMeEventIdAttachmentID parses 'input' into a MeEventIdAttachmentId
func ParseMeEventIdAttachmentID(input string) (*MeEventIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEventIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEventIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeEventIdAttachmentIDInsensitively parses 'input' case-insensitively into a MeEventIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeEventIdAttachmentIDInsensitively(input string) (*MeEventIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeEventIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeEventIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeEventIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateMeEventIdAttachmentID checks that 'input' can be parsed as a Me Event Id Attachment ID
func ValidateMeEventIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeEventIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Event Id Attachment ID
func (id MeEventIdAttachmentId) ID() string {
	fmtString := "/me/events/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Event Id Attachment ID
func (id MeEventIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentId"),
	}
}

// String returns a human-readable description of this Me Event Id Attachment ID
func (id MeEventIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Event Id Attachment (%s)", strings.Join(components, "\n"))
}
