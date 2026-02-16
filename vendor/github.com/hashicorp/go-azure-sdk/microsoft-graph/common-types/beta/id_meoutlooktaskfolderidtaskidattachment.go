package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOutlookTaskFolderIdTaskIdAttachmentId{}

// MeOutlookTaskFolderIdTaskIdAttachmentId is a struct representing the Resource ID for a Me Outlook Task Folder Id Task Id Attachment
type MeOutlookTaskFolderIdTaskIdAttachmentId struct {
	OutlookTaskFolderId string
	OutlookTaskId       string
	AttachmentId        string
}

// NewMeOutlookTaskFolderIdTaskIdAttachmentID returns a new MeOutlookTaskFolderIdTaskIdAttachmentId struct
func NewMeOutlookTaskFolderIdTaskIdAttachmentID(outlookTaskFolderId string, outlookTaskId string, attachmentId string) MeOutlookTaskFolderIdTaskIdAttachmentId {
	return MeOutlookTaskFolderIdTaskIdAttachmentId{
		OutlookTaskFolderId: outlookTaskFolderId,
		OutlookTaskId:       outlookTaskId,
		AttachmentId:        attachmentId,
	}
}

// ParseMeOutlookTaskFolderIdTaskIdAttachmentID parses 'input' into a MeOutlookTaskFolderIdTaskIdAttachmentId
func ParseMeOutlookTaskFolderIdTaskIdAttachmentID(input string) (*MeOutlookTaskFolderIdTaskIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOutlookTaskFolderIdTaskIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOutlookTaskFolderIdTaskIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOutlookTaskFolderIdTaskIdAttachmentIDInsensitively parses 'input' case-insensitively into a MeOutlookTaskFolderIdTaskIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeOutlookTaskFolderIdTaskIdAttachmentIDInsensitively(input string) (*MeOutlookTaskFolderIdTaskIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOutlookTaskFolderIdTaskIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOutlookTaskFolderIdTaskIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOutlookTaskFolderIdTaskIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OutlookTaskFolderId, ok = input.Parsed["outlookTaskFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskFolderId", input)
	}

	if id.OutlookTaskId, ok = input.Parsed["outlookTaskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskId", input)
	}

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateMeOutlookTaskFolderIdTaskIdAttachmentID checks that 'input' can be parsed as a Me Outlook Task Folder Id Task Id Attachment ID
func ValidateMeOutlookTaskFolderIdTaskIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOutlookTaskFolderIdTaskIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Outlook Task Folder Id Task Id Attachment ID
func (id MeOutlookTaskFolderIdTaskIdAttachmentId) ID() string {
	fmtString := "/me/outlook/taskFolders/%s/tasks/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.OutlookTaskFolderId, id.OutlookTaskId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Outlook Task Folder Id Task Id Attachment ID
func (id MeOutlookTaskFolderIdTaskIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("outlook", "outlook", "outlook"),
		resourceids.StaticSegment("taskFolders", "taskFolders", "taskFolders"),
		resourceids.UserSpecifiedSegment("outlookTaskFolderId", "outlookTaskFolderId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("outlookTaskId", "outlookTaskId"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentId"),
	}
}

// String returns a human-readable description of this Me Outlook Task Folder Id Task Id Attachment ID
func (id MeOutlookTaskFolderIdTaskIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Outlook Task Folder: %q", id.OutlookTaskFolderId),
		fmt.Sprintf("Outlook Task: %q", id.OutlookTaskId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Outlook Task Folder Id Task Id Attachment (%s)", strings.Join(components, "\n"))
}
