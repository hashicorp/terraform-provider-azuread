package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId{}

// UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId is a struct representing the Resource ID for a User Id Outlook Task Group Id Task Folder Id Task Id Attachment
type UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId struct {
	UserId              string
	OutlookTaskGroupId  string
	OutlookTaskFolderId string
	OutlookTaskId       string
	AttachmentId        string
}

// NewUserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID returns a new UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId struct
func NewUserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID(userId string, outlookTaskGroupId string, outlookTaskFolderId string, outlookTaskId string, attachmentId string) UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId {
	return UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId{
		UserId:              userId,
		OutlookTaskGroupId:  outlookTaskGroupId,
		OutlookTaskFolderId: outlookTaskFolderId,
		OutlookTaskId:       outlookTaskId,
		AttachmentId:        attachmentId,
	}
}

// ParseUserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID parses 'input' into a UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId
func ParseUserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID(input string) (*UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentIDInsensitively parses 'input' case-insensitively into a UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentIDInsensitively(input string) (*UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

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

// ValidateUserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID checks that 'input' can be parsed as a User Id Outlook Task Group Id Task Folder Id Task Id Attachment ID
func ValidateUserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Outlook Task Group Id Task Folder Id Task Id Attachment ID
func (id UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId) ID() string {
	fmtString := "/users/%s/outlook/taskGroups/%s/taskFolders/%s/tasks/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OutlookTaskGroupId, id.OutlookTaskFolderId, id.OutlookTaskId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Outlook Task Group Id Task Folder Id Task Id Attachment ID
func (id UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
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

// String returns a human-readable description of this User Id Outlook Task Group Id Task Folder Id Task Id Attachment ID
func (id UserIdOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Outlook Task Group: %q", id.OutlookTaskGroupId),
		fmt.Sprintf("Outlook Task Folder: %q", id.OutlookTaskFolderId),
		fmt.Sprintf("Outlook Task: %q", id.OutlookTaskId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Id Outlook Task Group Id Task Folder Id Task Id Attachment (%s)", strings.Join(components, "\n"))
}
