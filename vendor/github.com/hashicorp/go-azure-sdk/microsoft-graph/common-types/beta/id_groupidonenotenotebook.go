package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdOnenoteNotebookId{}

// GroupIdOnenoteNotebookId is a struct representing the Resource ID for a Group Id Onenote Notebook
type GroupIdOnenoteNotebookId struct {
	GroupId    string
	NotebookId string
}

// NewGroupIdOnenoteNotebookID returns a new GroupIdOnenoteNotebookId struct
func NewGroupIdOnenoteNotebookID(groupId string, notebookId string) GroupIdOnenoteNotebookId {
	return GroupIdOnenoteNotebookId{
		GroupId:    groupId,
		NotebookId: notebookId,
	}
}

// ParseGroupIdOnenoteNotebookID parses 'input' into a GroupIdOnenoteNotebookId
func ParseGroupIdOnenoteNotebookID(input string) (*GroupIdOnenoteNotebookId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteNotebookId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteNotebookId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdOnenoteNotebookIDInsensitively parses 'input' case-insensitively into a GroupIdOnenoteNotebookId
// note: this method should only be used for API response data and not user input
func ParseGroupIdOnenoteNotebookIDInsensitively(input string) (*GroupIdOnenoteNotebookId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteNotebookId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteNotebookId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdOnenoteNotebookId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.NotebookId, ok = input.Parsed["notebookId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "notebookId", input)
	}

	return nil
}

// ValidateGroupIdOnenoteNotebookID checks that 'input' can be parsed as a Group Id Onenote Notebook ID
func ValidateGroupIdOnenoteNotebookID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdOnenoteNotebookID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Onenote Notebook ID
func (id GroupIdOnenoteNotebookId) ID() string {
	fmtString := "/groups/%s/onenote/notebooks/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.NotebookId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Onenote Notebook ID
func (id GroupIdOnenoteNotebookId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("notebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookId"),
	}
}

// String returns a human-readable description of this Group Id Onenote Notebook ID
func (id GroupIdOnenoteNotebookId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
	}
	return fmt.Sprintf("Group Id Onenote Notebook (%s)", strings.Join(components, "\n"))
}
