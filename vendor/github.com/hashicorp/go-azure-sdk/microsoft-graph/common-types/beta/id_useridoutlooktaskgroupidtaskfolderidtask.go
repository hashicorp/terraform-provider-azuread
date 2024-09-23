package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOutlookTaskGroupIdTaskFolderIdTaskId{}

// UserIdOutlookTaskGroupIdTaskFolderIdTaskId is a struct representing the Resource ID for a User Id Outlook Task Group Id Task Folder Id Task
type UserIdOutlookTaskGroupIdTaskFolderIdTaskId struct {
	UserId              string
	OutlookTaskGroupId  string
	OutlookTaskFolderId string
	OutlookTaskId       string
}

// NewUserIdOutlookTaskGroupIdTaskFolderIdTaskID returns a new UserIdOutlookTaskGroupIdTaskFolderIdTaskId struct
func NewUserIdOutlookTaskGroupIdTaskFolderIdTaskID(userId string, outlookTaskGroupId string, outlookTaskFolderId string, outlookTaskId string) UserIdOutlookTaskGroupIdTaskFolderIdTaskId {
	return UserIdOutlookTaskGroupIdTaskFolderIdTaskId{
		UserId:              userId,
		OutlookTaskGroupId:  outlookTaskGroupId,
		OutlookTaskFolderId: outlookTaskFolderId,
		OutlookTaskId:       outlookTaskId,
	}
}

// ParseUserIdOutlookTaskGroupIdTaskFolderIdTaskID parses 'input' into a UserIdOutlookTaskGroupIdTaskFolderIdTaskId
func ParseUserIdOutlookTaskGroupIdTaskFolderIdTaskID(input string) (*UserIdOutlookTaskGroupIdTaskFolderIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOutlookTaskGroupIdTaskFolderIdTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOutlookTaskGroupIdTaskFolderIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOutlookTaskGroupIdTaskFolderIdTaskIDInsensitively parses 'input' case-insensitively into a UserIdOutlookTaskGroupIdTaskFolderIdTaskId
// note: this method should only be used for API response data and not user input
func ParseUserIdOutlookTaskGroupIdTaskFolderIdTaskIDInsensitively(input string) (*UserIdOutlookTaskGroupIdTaskFolderIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOutlookTaskGroupIdTaskFolderIdTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOutlookTaskGroupIdTaskFolderIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOutlookTaskGroupIdTaskFolderIdTaskId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateUserIdOutlookTaskGroupIdTaskFolderIdTaskID checks that 'input' can be parsed as a User Id Outlook Task Group Id Task Folder Id Task ID
func ValidateUserIdOutlookTaskGroupIdTaskFolderIdTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOutlookTaskGroupIdTaskFolderIdTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Outlook Task Group Id Task Folder Id Task ID
func (id UserIdOutlookTaskGroupIdTaskFolderIdTaskId) ID() string {
	fmtString := "/users/%s/outlook/taskGroups/%s/taskFolders/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OutlookTaskGroupId, id.OutlookTaskFolderId, id.OutlookTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Outlook Task Group Id Task Folder Id Task ID
func (id UserIdOutlookTaskGroupIdTaskFolderIdTaskId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this User Id Outlook Task Group Id Task Folder Id Task ID
func (id UserIdOutlookTaskGroupIdTaskFolderIdTaskId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Outlook Task Group: %q", id.OutlookTaskGroupId),
		fmt.Sprintf("Outlook Task Folder: %q", id.OutlookTaskFolderId),
		fmt.Sprintf("Outlook Task: %q", id.OutlookTaskId),
	}
	return fmt.Sprintf("User Id Outlook Task Group Id Task Folder Id Task (%s)", strings.Join(components, "\n"))
}
