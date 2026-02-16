package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdOnenoteNotebookIdSectionGroupIdSectionId{}

// GroupIdOnenoteNotebookIdSectionGroupIdSectionId is a struct representing the Resource ID for a Group Id Onenote Notebook Id Section Group Id Section
type GroupIdOnenoteNotebookIdSectionGroupIdSectionId struct {
	GroupId          string
	NotebookId       string
	SectionGroupId   string
	OnenoteSectionId string
}

// NewGroupIdOnenoteNotebookIdSectionGroupIdSectionID returns a new GroupIdOnenoteNotebookIdSectionGroupIdSectionId struct
func NewGroupIdOnenoteNotebookIdSectionGroupIdSectionID(groupId string, notebookId string, sectionGroupId string, onenoteSectionId string) GroupIdOnenoteNotebookIdSectionGroupIdSectionId {
	return GroupIdOnenoteNotebookIdSectionGroupIdSectionId{
		GroupId:          groupId,
		NotebookId:       notebookId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseGroupIdOnenoteNotebookIdSectionGroupIdSectionID parses 'input' into a GroupIdOnenoteNotebookIdSectionGroupIdSectionId
func ParseGroupIdOnenoteNotebookIdSectionGroupIdSectionID(input string) (*GroupIdOnenoteNotebookIdSectionGroupIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteNotebookIdSectionGroupIdSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteNotebookIdSectionGroupIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdOnenoteNotebookIdSectionGroupIdSectionIDInsensitively parses 'input' case-insensitively into a GroupIdOnenoteNotebookIdSectionGroupIdSectionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdOnenoteNotebookIdSectionGroupIdSectionIDInsensitively(input string) (*GroupIdOnenoteNotebookIdSectionGroupIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteNotebookIdSectionGroupIdSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteNotebookIdSectionGroupIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdOnenoteNotebookIdSectionGroupIdSectionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.NotebookId, ok = input.Parsed["notebookId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "notebookId", input)
	}

	if id.SectionGroupId, ok = input.Parsed["sectionGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", input)
	}

	if id.OnenoteSectionId, ok = input.Parsed["onenoteSectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", input)
	}

	return nil
}

// ValidateGroupIdOnenoteNotebookIdSectionGroupIdSectionID checks that 'input' can be parsed as a Group Id Onenote Notebook Id Section Group Id Section ID
func ValidateGroupIdOnenoteNotebookIdSectionGroupIdSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdOnenoteNotebookIdSectionGroupIdSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Onenote Notebook Id Section Group Id Section ID
func (id GroupIdOnenoteNotebookIdSectionGroupIdSectionId) ID() string {
	fmtString := "/groups/%s/onenote/notebooks/%s/sectionGroups/%s/sections/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.NotebookId, id.SectionGroupId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Onenote Notebook Id Section Group Id Section ID
func (id GroupIdOnenoteNotebookIdSectionGroupIdSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("notebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookId"),
		resourceids.StaticSegment("sectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupId"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
	}
}

// String returns a human-readable description of this Group Id Onenote Notebook Id Section Group Id Section ID
func (id GroupIdOnenoteNotebookIdSectionGroupIdSectionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("Group Id Onenote Notebook Id Section Group Id Section (%s)", strings.Join(components, "\n"))
}
