package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOutlookTaskFolderIdTaskId{}

// UserIdOutlookTaskFolderIdTaskId is a struct representing the Resource ID for a User Id Outlook Task Folder Id Task
type UserIdOutlookTaskFolderIdTaskId struct {
	UserId              string
	OutlookTaskFolderId string
	OutlookTaskId       string
}

// NewUserIdOutlookTaskFolderIdTaskID returns a new UserIdOutlookTaskFolderIdTaskId struct
func NewUserIdOutlookTaskFolderIdTaskID(userId string, outlookTaskFolderId string, outlookTaskId string) UserIdOutlookTaskFolderIdTaskId {
	return UserIdOutlookTaskFolderIdTaskId{
		UserId:              userId,
		OutlookTaskFolderId: outlookTaskFolderId,
		OutlookTaskId:       outlookTaskId,
	}
}

// ParseUserIdOutlookTaskFolderIdTaskID parses 'input' into a UserIdOutlookTaskFolderIdTaskId
func ParseUserIdOutlookTaskFolderIdTaskID(input string) (*UserIdOutlookTaskFolderIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOutlookTaskFolderIdTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOutlookTaskFolderIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOutlookTaskFolderIdTaskIDInsensitively parses 'input' case-insensitively into a UserIdOutlookTaskFolderIdTaskId
// note: this method should only be used for API response data and not user input
func ParseUserIdOutlookTaskFolderIdTaskIDInsensitively(input string) (*UserIdOutlookTaskFolderIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOutlookTaskFolderIdTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOutlookTaskFolderIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOutlookTaskFolderIdTaskId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateUserIdOutlookTaskFolderIdTaskID checks that 'input' can be parsed as a User Id Outlook Task Folder Id Task ID
func ValidateUserIdOutlookTaskFolderIdTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOutlookTaskFolderIdTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Outlook Task Folder Id Task ID
func (id UserIdOutlookTaskFolderIdTaskId) ID() string {
	fmtString := "/users/%s/outlook/taskFolders/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OutlookTaskFolderId, id.OutlookTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Outlook Task Folder Id Task ID
func (id UserIdOutlookTaskFolderIdTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("outlook", "outlook", "outlook"),
		resourceids.StaticSegment("taskFolders", "taskFolders", "taskFolders"),
		resourceids.UserSpecifiedSegment("outlookTaskFolderId", "outlookTaskFolderId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("outlookTaskId", "outlookTaskId"),
	}
}

// String returns a human-readable description of this User Id Outlook Task Folder Id Task ID
func (id UserIdOutlookTaskFolderIdTaskId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Outlook Task Folder: %q", id.OutlookTaskFolderId),
		fmt.Sprintf("Outlook Task: %q", id.OutlookTaskId),
	}
	return fmt.Sprintf("User Id Outlook Task Folder Id Task (%s)", strings.Join(components, "\n"))
}
