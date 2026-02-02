package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOutlookTaskFolderId{}

// UserIdOutlookTaskFolderId is a struct representing the Resource ID for a User Id Outlook Task Folder
type UserIdOutlookTaskFolderId struct {
	UserId              string
	OutlookTaskFolderId string
}

// NewUserIdOutlookTaskFolderID returns a new UserIdOutlookTaskFolderId struct
func NewUserIdOutlookTaskFolderID(userId string, outlookTaskFolderId string) UserIdOutlookTaskFolderId {
	return UserIdOutlookTaskFolderId{
		UserId:              userId,
		OutlookTaskFolderId: outlookTaskFolderId,
	}
}

// ParseUserIdOutlookTaskFolderID parses 'input' into a UserIdOutlookTaskFolderId
func ParseUserIdOutlookTaskFolderID(input string) (*UserIdOutlookTaskFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOutlookTaskFolderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOutlookTaskFolderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOutlookTaskFolderIDInsensitively parses 'input' case-insensitively into a UserIdOutlookTaskFolderId
// note: this method should only be used for API response data and not user input
func ParseUserIdOutlookTaskFolderIDInsensitively(input string) (*UserIdOutlookTaskFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOutlookTaskFolderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOutlookTaskFolderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOutlookTaskFolderId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.OutlookTaskFolderId, ok = input.Parsed["outlookTaskFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskFolderId", input)
	}

	return nil
}

// ValidateUserIdOutlookTaskFolderID checks that 'input' can be parsed as a User Id Outlook Task Folder ID
func ValidateUserIdOutlookTaskFolderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOutlookTaskFolderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Outlook Task Folder ID
func (id UserIdOutlookTaskFolderId) ID() string {
	fmtString := "/users/%s/outlook/taskFolders/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OutlookTaskFolderId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Outlook Task Folder ID
func (id UserIdOutlookTaskFolderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("outlook", "outlook", "outlook"),
		resourceids.StaticSegment("taskFolders", "taskFolders", "taskFolders"),
		resourceids.UserSpecifiedSegment("outlookTaskFolderId", "outlookTaskFolderId"),
	}
}

// String returns a human-readable description of this User Id Outlook Task Folder ID
func (id UserIdOutlookTaskFolderId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Outlook Task Folder: %q", id.OutlookTaskFolderId),
	}
	return fmt.Sprintf("User Id Outlook Task Folder (%s)", strings.Join(components, "\n"))
}
