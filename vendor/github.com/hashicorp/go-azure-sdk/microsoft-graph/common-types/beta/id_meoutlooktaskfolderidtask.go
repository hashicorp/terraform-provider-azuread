package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOutlookTaskFolderIdTaskId{}

// MeOutlookTaskFolderIdTaskId is a struct representing the Resource ID for a Me Outlook Task Folder Id Task
type MeOutlookTaskFolderIdTaskId struct {
	OutlookTaskFolderId string
	OutlookTaskId       string
}

// NewMeOutlookTaskFolderIdTaskID returns a new MeOutlookTaskFolderIdTaskId struct
func NewMeOutlookTaskFolderIdTaskID(outlookTaskFolderId string, outlookTaskId string) MeOutlookTaskFolderIdTaskId {
	return MeOutlookTaskFolderIdTaskId{
		OutlookTaskFolderId: outlookTaskFolderId,
		OutlookTaskId:       outlookTaskId,
	}
}

// ParseMeOutlookTaskFolderIdTaskID parses 'input' into a MeOutlookTaskFolderIdTaskId
func ParseMeOutlookTaskFolderIdTaskID(input string) (*MeOutlookTaskFolderIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOutlookTaskFolderIdTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOutlookTaskFolderIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOutlookTaskFolderIdTaskIDInsensitively parses 'input' case-insensitively into a MeOutlookTaskFolderIdTaskId
// note: this method should only be used for API response data and not user input
func ParseMeOutlookTaskFolderIdTaskIDInsensitively(input string) (*MeOutlookTaskFolderIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOutlookTaskFolderIdTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOutlookTaskFolderIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOutlookTaskFolderIdTaskId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OutlookTaskFolderId, ok = input.Parsed["outlookTaskFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskFolderId", input)
	}

	if id.OutlookTaskId, ok = input.Parsed["outlookTaskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskId", input)
	}

	return nil
}

// ValidateMeOutlookTaskFolderIdTaskID checks that 'input' can be parsed as a Me Outlook Task Folder Id Task ID
func ValidateMeOutlookTaskFolderIdTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOutlookTaskFolderIdTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Outlook Task Folder Id Task ID
func (id MeOutlookTaskFolderIdTaskId) ID() string {
	fmtString := "/me/outlook/taskFolders/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.OutlookTaskFolderId, id.OutlookTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Outlook Task Folder Id Task ID
func (id MeOutlookTaskFolderIdTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("outlook", "outlook", "outlook"),
		resourceids.StaticSegment("taskFolders", "taskFolders", "taskFolders"),
		resourceids.UserSpecifiedSegment("outlookTaskFolderId", "outlookTaskFolderId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("outlookTaskId", "outlookTaskId"),
	}
}

// String returns a human-readable description of this Me Outlook Task Folder Id Task ID
func (id MeOutlookTaskFolderIdTaskId) String() string {
	components := []string{
		fmt.Sprintf("Outlook Task Folder: %q", id.OutlookTaskFolderId),
		fmt.Sprintf("Outlook Task: %q", id.OutlookTaskId),
	}
	return fmt.Sprintf("Me Outlook Task Folder Id Task (%s)", strings.Join(components, "\n"))
}
