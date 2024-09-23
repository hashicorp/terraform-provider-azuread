package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId{}

// MeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId is a struct representing the Resource ID for a Me Outlook Task Group Id Task Folder Id Task Id Attachment
type MeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId struct {
	OutlookTaskGroupId  string
	OutlookTaskFolderId string
	OutlookTaskId       string
	AttachmentId        string
}

// NewMeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID returns a new MeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId struct
func NewMeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID(outlookTaskGroupId string, outlookTaskFolderId string, outlookTaskId string, attachmentId string) MeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId {
	return MeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId{
		OutlookTaskGroupId:  outlookTaskGroupId,
		OutlookTaskFolderId: outlookTaskFolderId,
		OutlookTaskId:       outlookTaskId,
		AttachmentId:        attachmentId,
	}
}

// ParseMeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID parses 'input' into a MeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId
func ParseMeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID(input string) (*MeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentIDInsensitively parses 'input' case-insensitively into a MeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentIDInsensitively(input string) (*MeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OutlookTaskGroupId, ok = input.Parsed["outlookTaskGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskGroupId", input)
	}

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

// ValidateMeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID checks that 'input' can be parsed as a Me Outlook Task Group Id Task Folder Id Task Id Attachment ID
func ValidateMeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Outlook Task Group Id Task Folder Id Task Id Attachment ID
func (id MeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId) ID() string {
	fmtString := "/me/outlook/taskGroups/%s/taskFolders/%s/tasks/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.OutlookTaskGroupId, id.OutlookTaskFolderId, id.OutlookTaskId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Outlook Task Group Id Task Folder Id Task Id Attachment ID
func (id MeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("outlook", "outlook", "outlook"),
		resourceids.StaticSegment("taskGroups", "taskGroups", "taskGroups"),
		resourceids.UserSpecifiedSegment("outlookTaskGroupId", "outlookTaskGroupId"),
		resourceids.StaticSegment("taskFolders", "taskFolders", "taskFolders"),
		resourceids.UserSpecifiedSegment("outlookTaskFolderId", "outlookTaskFolderId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("outlookTaskId", "outlookTaskId"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentId"),
	}
}

// String returns a human-readable description of this Me Outlook Task Group Id Task Folder Id Task Id Attachment ID
func (id MeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Outlook Task Group: %q", id.OutlookTaskGroupId),
		fmt.Sprintf("Outlook Task Folder: %q", id.OutlookTaskFolderId),
		fmt.Sprintf("Outlook Task: %q", id.OutlookTaskId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Outlook Task Group Id Task Folder Id Task Id Attachment (%s)", strings.Join(components, "\n"))
}
