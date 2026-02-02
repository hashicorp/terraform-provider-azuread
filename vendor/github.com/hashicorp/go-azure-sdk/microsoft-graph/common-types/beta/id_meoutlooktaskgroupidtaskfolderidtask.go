package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOutlookTaskGroupIdTaskFolderIdTaskId{}

// MeOutlookTaskGroupIdTaskFolderIdTaskId is a struct representing the Resource ID for a Me Outlook Task Group Id Task Folder Id Task
type MeOutlookTaskGroupIdTaskFolderIdTaskId struct {
	OutlookTaskGroupId  string
	OutlookTaskFolderId string
	OutlookTaskId       string
}

// NewMeOutlookTaskGroupIdTaskFolderIdTaskID returns a new MeOutlookTaskGroupIdTaskFolderIdTaskId struct
func NewMeOutlookTaskGroupIdTaskFolderIdTaskID(outlookTaskGroupId string, outlookTaskFolderId string, outlookTaskId string) MeOutlookTaskGroupIdTaskFolderIdTaskId {
	return MeOutlookTaskGroupIdTaskFolderIdTaskId{
		OutlookTaskGroupId:  outlookTaskGroupId,
		OutlookTaskFolderId: outlookTaskFolderId,
		OutlookTaskId:       outlookTaskId,
	}
}

// ParseMeOutlookTaskGroupIdTaskFolderIdTaskID parses 'input' into a MeOutlookTaskGroupIdTaskFolderIdTaskId
func ParseMeOutlookTaskGroupIdTaskFolderIdTaskID(input string) (*MeOutlookTaskGroupIdTaskFolderIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOutlookTaskGroupIdTaskFolderIdTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOutlookTaskGroupIdTaskFolderIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOutlookTaskGroupIdTaskFolderIdTaskIDInsensitively parses 'input' case-insensitively into a MeOutlookTaskGroupIdTaskFolderIdTaskId
// note: this method should only be used for API response data and not user input
func ParseMeOutlookTaskGroupIdTaskFolderIdTaskIDInsensitively(input string) (*MeOutlookTaskGroupIdTaskFolderIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOutlookTaskGroupIdTaskFolderIdTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOutlookTaskGroupIdTaskFolderIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOutlookTaskGroupIdTaskFolderIdTaskId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateMeOutlookTaskGroupIdTaskFolderIdTaskID checks that 'input' can be parsed as a Me Outlook Task Group Id Task Folder Id Task ID
func ValidateMeOutlookTaskGroupIdTaskFolderIdTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOutlookTaskGroupIdTaskFolderIdTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Outlook Task Group Id Task Folder Id Task ID
func (id MeOutlookTaskGroupIdTaskFolderIdTaskId) ID() string {
	fmtString := "/me/outlook/taskGroups/%s/taskFolders/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.OutlookTaskGroupId, id.OutlookTaskFolderId, id.OutlookTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Outlook Task Group Id Task Folder Id Task ID
func (id MeOutlookTaskGroupIdTaskFolderIdTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("outlook", "outlook", "outlook"),
		resourceids.StaticSegment("taskGroups", "taskGroups", "taskGroups"),
		resourceids.UserSpecifiedSegment("outlookTaskGroupId", "outlookTaskGroupId"),
		resourceids.StaticSegment("taskFolders", "taskFolders", "taskFolders"),
		resourceids.UserSpecifiedSegment("outlookTaskFolderId", "outlookTaskFolderId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("outlookTaskId", "outlookTaskId"),
	}
}

// String returns a human-readable description of this Me Outlook Task Group Id Task Folder Id Task ID
func (id MeOutlookTaskGroupIdTaskFolderIdTaskId) String() string {
	components := []string{
		fmt.Sprintf("Outlook Task Group: %q", id.OutlookTaskGroupId),
		fmt.Sprintf("Outlook Task Folder: %q", id.OutlookTaskFolderId),
		fmt.Sprintf("Outlook Task: %q", id.OutlookTaskId),
	}
	return fmt.Sprintf("Me Outlook Task Group Id Task Folder Id Task (%s)", strings.Join(components, "\n"))
}
