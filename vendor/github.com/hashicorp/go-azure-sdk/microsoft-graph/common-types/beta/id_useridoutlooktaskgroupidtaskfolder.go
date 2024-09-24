package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOutlookTaskGroupIdTaskFolderId{}

// UserIdOutlookTaskGroupIdTaskFolderId is a struct representing the Resource ID for a User Id Outlook Task Group Id Task Folder
type UserIdOutlookTaskGroupIdTaskFolderId struct {
	UserId              string
	OutlookTaskGroupId  string
	OutlookTaskFolderId string
}

// NewUserIdOutlookTaskGroupIdTaskFolderID returns a new UserIdOutlookTaskGroupIdTaskFolderId struct
func NewUserIdOutlookTaskGroupIdTaskFolderID(userId string, outlookTaskGroupId string, outlookTaskFolderId string) UserIdOutlookTaskGroupIdTaskFolderId {
	return UserIdOutlookTaskGroupIdTaskFolderId{
		UserId:              userId,
		OutlookTaskGroupId:  outlookTaskGroupId,
		OutlookTaskFolderId: outlookTaskFolderId,
	}
}

// ParseUserIdOutlookTaskGroupIdTaskFolderID parses 'input' into a UserIdOutlookTaskGroupIdTaskFolderId
func ParseUserIdOutlookTaskGroupIdTaskFolderID(input string) (*UserIdOutlookTaskGroupIdTaskFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOutlookTaskGroupIdTaskFolderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOutlookTaskGroupIdTaskFolderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOutlookTaskGroupIdTaskFolderIDInsensitively parses 'input' case-insensitively into a UserIdOutlookTaskGroupIdTaskFolderId
// note: this method should only be used for API response data and not user input
func ParseUserIdOutlookTaskGroupIdTaskFolderIDInsensitively(input string) (*UserIdOutlookTaskGroupIdTaskFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOutlookTaskGroupIdTaskFolderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOutlookTaskGroupIdTaskFolderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOutlookTaskGroupIdTaskFolderId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateUserIdOutlookTaskGroupIdTaskFolderID checks that 'input' can be parsed as a User Id Outlook Task Group Id Task Folder ID
func ValidateUserIdOutlookTaskGroupIdTaskFolderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOutlookTaskGroupIdTaskFolderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Outlook Task Group Id Task Folder ID
func (id UserIdOutlookTaskGroupIdTaskFolderId) ID() string {
	fmtString := "/users/%s/outlook/taskGroups/%s/taskFolders/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OutlookTaskGroupId, id.OutlookTaskFolderId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Outlook Task Group Id Task Folder ID
func (id UserIdOutlookTaskGroupIdTaskFolderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("outlook", "outlook", "outlook"),
		resourceids.StaticSegment("taskGroups", "taskGroups", "taskGroups"),
		resourceids.UserSpecifiedSegment("outlookTaskGroupId", "outlookTaskGroupId"),
		resourceids.StaticSegment("taskFolders", "taskFolders", "taskFolders"),
		resourceids.UserSpecifiedSegment("outlookTaskFolderId", "outlookTaskFolderId"),
	}
}

// String returns a human-readable description of this User Id Outlook Task Group Id Task Folder ID
func (id UserIdOutlookTaskGroupIdTaskFolderId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Outlook Task Group: %q", id.OutlookTaskGroupId),
		fmt.Sprintf("Outlook Task Folder: %q", id.OutlookTaskFolderId),
	}
	return fmt.Sprintf("User Id Outlook Task Group Id Task Folder (%s)", strings.Join(components, "\n"))
}
