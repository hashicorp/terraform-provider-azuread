package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOutlookTaskIdAttachmentId{}

// MeOutlookTaskIdAttachmentId is a struct representing the Resource ID for a Me Outlook Task Id Attachment
type MeOutlookTaskIdAttachmentId struct {
	OutlookTaskId string
	AttachmentId  string
}

// NewMeOutlookTaskIdAttachmentID returns a new MeOutlookTaskIdAttachmentId struct
func NewMeOutlookTaskIdAttachmentID(outlookTaskId string, attachmentId string) MeOutlookTaskIdAttachmentId {
	return MeOutlookTaskIdAttachmentId{
		OutlookTaskId: outlookTaskId,
		AttachmentId:  attachmentId,
	}
}

// ParseMeOutlookTaskIdAttachmentID parses 'input' into a MeOutlookTaskIdAttachmentId
func ParseMeOutlookTaskIdAttachmentID(input string) (*MeOutlookTaskIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOutlookTaskIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOutlookTaskIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOutlookTaskIdAttachmentIDInsensitively parses 'input' case-insensitively into a MeOutlookTaskIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeOutlookTaskIdAttachmentIDInsensitively(input string) (*MeOutlookTaskIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOutlookTaskIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOutlookTaskIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOutlookTaskIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OutlookTaskId, ok = input.Parsed["outlookTaskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskId", input)
	}

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateMeOutlookTaskIdAttachmentID checks that 'input' can be parsed as a Me Outlook Task Id Attachment ID
func ValidateMeOutlookTaskIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOutlookTaskIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Outlook Task Id Attachment ID
func (id MeOutlookTaskIdAttachmentId) ID() string {
	fmtString := "/me/outlook/tasks/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.OutlookTaskId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Outlook Task Id Attachment ID
func (id MeOutlookTaskIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("outlook", "outlook", "outlook"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("outlookTaskId", "outlookTaskId"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentId"),
	}
}

// String returns a human-readable description of this Me Outlook Task Id Attachment ID
func (id MeOutlookTaskIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Outlook Task: %q", id.OutlookTaskId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Outlook Task Id Attachment (%s)", strings.Join(components, "\n"))
}
