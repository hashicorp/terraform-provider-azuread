package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOutlookTaskFolderIdTaskIdAttachmentId{}

// UserIdOutlookTaskFolderIdTaskIdAttachmentId is a struct representing the Resource ID for a User Id Outlook Task Folder Id Task Id Attachment
type UserIdOutlookTaskFolderIdTaskIdAttachmentId struct {
	UserId              string
	OutlookTaskFolderId string
	OutlookTaskId       string
	AttachmentId        string
}

// NewUserIdOutlookTaskFolderIdTaskIdAttachmentID returns a new UserIdOutlookTaskFolderIdTaskIdAttachmentId struct
func NewUserIdOutlookTaskFolderIdTaskIdAttachmentID(userId string, outlookTaskFolderId string, outlookTaskId string, attachmentId string) UserIdOutlookTaskFolderIdTaskIdAttachmentId {
	return UserIdOutlookTaskFolderIdTaskIdAttachmentId{
		UserId:              userId,
		OutlookTaskFolderId: outlookTaskFolderId,
		OutlookTaskId:       outlookTaskId,
		AttachmentId:        attachmentId,
	}
}

// ParseUserIdOutlookTaskFolderIdTaskIdAttachmentID parses 'input' into a UserIdOutlookTaskFolderIdTaskIdAttachmentId
func ParseUserIdOutlookTaskFolderIdTaskIdAttachmentID(input string) (*UserIdOutlookTaskFolderIdTaskIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOutlookTaskFolderIdTaskIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOutlookTaskFolderIdTaskIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOutlookTaskFolderIdTaskIdAttachmentIDInsensitively parses 'input' case-insensitively into a UserIdOutlookTaskFolderIdTaskIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserIdOutlookTaskFolderIdTaskIdAttachmentIDInsensitively(input string) (*UserIdOutlookTaskFolderIdTaskIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOutlookTaskFolderIdTaskIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOutlookTaskFolderIdTaskIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOutlookTaskFolderIdTaskIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
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

// ValidateUserIdOutlookTaskFolderIdTaskIdAttachmentID checks that 'input' can be parsed as a User Id Outlook Task Folder Id Task Id Attachment ID
func ValidateUserIdOutlookTaskFolderIdTaskIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOutlookTaskFolderIdTaskIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Outlook Task Folder Id Task Id Attachment ID
func (id UserIdOutlookTaskFolderIdTaskIdAttachmentId) ID() string {
	fmtString := "/users/%s/outlook/taskFolders/%s/tasks/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OutlookTaskFolderId, id.OutlookTaskId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Outlook Task Folder Id Task Id Attachment ID
func (id UserIdOutlookTaskFolderIdTaskIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("outlook", "outlook", "outlook"),
		resourceids.StaticSegment("taskFolders", "taskFolders", "taskFolders"),
		resourceids.UserSpecifiedSegment("outlookTaskFolderId", "outlookTaskFolderId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("outlookTaskId", "outlookTaskId"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentId"),
	}
}

// String returns a human-readable description of this User Id Outlook Task Folder Id Task Id Attachment ID
func (id UserIdOutlookTaskFolderIdTaskIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Outlook Task Folder: %q", id.OutlookTaskFolderId),
		fmt.Sprintf("Outlook Task: %q", id.OutlookTaskId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Id Outlook Task Folder Id Task Id Attachment (%s)", strings.Join(components, "\n"))
}
