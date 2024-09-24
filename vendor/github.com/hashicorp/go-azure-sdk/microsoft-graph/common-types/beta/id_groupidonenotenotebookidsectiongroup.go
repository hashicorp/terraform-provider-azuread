package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdOnenoteNotebookIdSectionGroupId{}

// GroupIdOnenoteNotebookIdSectionGroupId is a struct representing the Resource ID for a Group Id Onenote Notebook Id Section Group
type GroupIdOnenoteNotebookIdSectionGroupId struct {
	GroupId        string
	NotebookId     string
	SectionGroupId string
}

// NewGroupIdOnenoteNotebookIdSectionGroupID returns a new GroupIdOnenoteNotebookIdSectionGroupId struct
func NewGroupIdOnenoteNotebookIdSectionGroupID(groupId string, notebookId string, sectionGroupId string) GroupIdOnenoteNotebookIdSectionGroupId {
	return GroupIdOnenoteNotebookIdSectionGroupId{
		GroupId:        groupId,
		NotebookId:     notebookId,
		SectionGroupId: sectionGroupId,
	}
}

// ParseGroupIdOnenoteNotebookIdSectionGroupID parses 'input' into a GroupIdOnenoteNotebookIdSectionGroupId
func ParseGroupIdOnenoteNotebookIdSectionGroupID(input string) (*GroupIdOnenoteNotebookIdSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteNotebookIdSectionGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteNotebookIdSectionGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdOnenoteNotebookIdSectionGroupIDInsensitively parses 'input' case-insensitively into a GroupIdOnenoteNotebookIdSectionGroupId
// note: this method should only be used for API response data and not user input
func ParseGroupIdOnenoteNotebookIdSectionGroupIDInsensitively(input string) (*GroupIdOnenoteNotebookIdSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteNotebookIdSectionGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteNotebookIdSectionGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdOnenoteNotebookIdSectionGroupId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateGroupIdOnenoteNotebookIdSectionGroupID checks that 'input' can be parsed as a Group Id Onenote Notebook Id Section Group ID
func ValidateGroupIdOnenoteNotebookIdSectionGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdOnenoteNotebookIdSectionGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Onenote Notebook Id Section Group ID
func (id GroupIdOnenoteNotebookIdSectionGroupId) ID() string {
	fmtString := "/groups/%s/onenote/notebooks/%s/sectionGroups/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.NotebookId, id.SectionGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Onenote Notebook Id Section Group ID
func (id GroupIdOnenoteNotebookIdSectionGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("notebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookId"),
		resourceids.StaticSegment("sectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupId"),
	}
}

// String returns a human-readable description of this Group Id Onenote Notebook Id Section Group ID
func (id GroupIdOnenoteNotebookIdSectionGroupId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
	}
	return fmt.Sprintf("Group Id Onenote Notebook Id Section Group (%s)", strings.Join(components, "\n"))
}
