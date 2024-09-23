package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOutlookTaskGroupIdTaskFolderId{}

// MeOutlookTaskGroupIdTaskFolderId is a struct representing the Resource ID for a Me Outlook Task Group Id Task Folder
type MeOutlookTaskGroupIdTaskFolderId struct {
	OutlookTaskGroupId  string
	OutlookTaskFolderId string
}

// NewMeOutlookTaskGroupIdTaskFolderID returns a new MeOutlookTaskGroupIdTaskFolderId struct
func NewMeOutlookTaskGroupIdTaskFolderID(outlookTaskGroupId string, outlookTaskFolderId string) MeOutlookTaskGroupIdTaskFolderId {
	return MeOutlookTaskGroupIdTaskFolderId{
		OutlookTaskGroupId:  outlookTaskGroupId,
		OutlookTaskFolderId: outlookTaskFolderId,
	}
}

// ParseMeOutlookTaskGroupIdTaskFolderID parses 'input' into a MeOutlookTaskGroupIdTaskFolderId
func ParseMeOutlookTaskGroupIdTaskFolderID(input string) (*MeOutlookTaskGroupIdTaskFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOutlookTaskGroupIdTaskFolderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOutlookTaskGroupIdTaskFolderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOutlookTaskGroupIdTaskFolderIDInsensitively parses 'input' case-insensitively into a MeOutlookTaskGroupIdTaskFolderId
// note: this method should only be used for API response data and not user input
func ParseMeOutlookTaskGroupIdTaskFolderIDInsensitively(input string) (*MeOutlookTaskGroupIdTaskFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOutlookTaskGroupIdTaskFolderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOutlookTaskGroupIdTaskFolderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOutlookTaskGroupIdTaskFolderId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OutlookTaskGroupId, ok = input.Parsed["outlookTaskGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskGroupId", input)
	}

	if id.OutlookTaskFolderId, ok = input.Parsed["outlookTaskFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskFolderId", input)
	}

	return nil
}

// ValidateMeOutlookTaskGroupIdTaskFolderID checks that 'input' can be parsed as a Me Outlook Task Group Id Task Folder ID
func ValidateMeOutlookTaskGroupIdTaskFolderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOutlookTaskGroupIdTaskFolderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Outlook Task Group Id Task Folder ID
func (id MeOutlookTaskGroupIdTaskFolderId) ID() string {
	fmtString := "/me/outlook/taskGroups/%s/taskFolders/%s"
	return fmt.Sprintf(fmtString, id.OutlookTaskGroupId, id.OutlookTaskFolderId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Outlook Task Group Id Task Folder ID
func (id MeOutlookTaskGroupIdTaskFolderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("outlook", "outlook", "outlook"),
		resourceids.StaticSegment("taskGroups", "taskGroups", "taskGroups"),
		resourceids.UserSpecifiedSegment("outlookTaskGroupId", "outlookTaskGroupId"),
		resourceids.StaticSegment("taskFolders", "taskFolders", "taskFolders"),
		resourceids.UserSpecifiedSegment("outlookTaskFolderId", "outlookTaskFolderId"),
	}
}

// String returns a human-readable description of this Me Outlook Task Group Id Task Folder ID
func (id MeOutlookTaskGroupIdTaskFolderId) String() string {
	components := []string{
		fmt.Sprintf("Outlook Task Group: %q", id.OutlookTaskGroupId),
		fmt.Sprintf("Outlook Task Folder: %q", id.OutlookTaskFolderId),
	}
	return fmt.Sprintf("Me Outlook Task Group Id Task Folder (%s)", strings.Join(components, "\n"))
}
