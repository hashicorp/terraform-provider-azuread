package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdOnenoteNotebookIdSectionId{}

// GroupIdOnenoteNotebookIdSectionId is a struct representing the Resource ID for a Group Id Onenote Notebook Id Section
type GroupIdOnenoteNotebookIdSectionId struct {
	GroupId          string
	NotebookId       string
	OnenoteSectionId string
}

// NewGroupIdOnenoteNotebookIdSectionID returns a new GroupIdOnenoteNotebookIdSectionId struct
func NewGroupIdOnenoteNotebookIdSectionID(groupId string, notebookId string, onenoteSectionId string) GroupIdOnenoteNotebookIdSectionId {
	return GroupIdOnenoteNotebookIdSectionId{
		GroupId:          groupId,
		NotebookId:       notebookId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseGroupIdOnenoteNotebookIdSectionID parses 'input' into a GroupIdOnenoteNotebookIdSectionId
func ParseGroupIdOnenoteNotebookIdSectionID(input string) (*GroupIdOnenoteNotebookIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteNotebookIdSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteNotebookIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdOnenoteNotebookIdSectionIDInsensitively parses 'input' case-insensitively into a GroupIdOnenoteNotebookIdSectionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdOnenoteNotebookIdSectionIDInsensitively(input string) (*GroupIdOnenoteNotebookIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteNotebookIdSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteNotebookIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdOnenoteNotebookIdSectionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.NotebookId, ok = input.Parsed["notebookId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "notebookId", input)
	}

	if id.OnenoteSectionId, ok = input.Parsed["onenoteSectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", input)
	}

	return nil
}

// ValidateGroupIdOnenoteNotebookIdSectionID checks that 'input' can be parsed as a Group Id Onenote Notebook Id Section ID
func ValidateGroupIdOnenoteNotebookIdSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdOnenoteNotebookIdSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Onenote Notebook Id Section ID
func (id GroupIdOnenoteNotebookIdSectionId) ID() string {
	fmtString := "/groups/%s/onenote/notebooks/%s/sections/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.NotebookId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Onenote Notebook Id Section ID
func (id GroupIdOnenoteNotebookIdSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("notebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookId"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
	}
}

// String returns a human-readable description of this Group Id Onenote Notebook Id Section ID
func (id GroupIdOnenoteNotebookIdSectionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("Group Id Onenote Notebook Id Section (%s)", strings.Join(components, "\n"))
}
