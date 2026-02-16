package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnenoteNotebookIdSectionGroupIdSectionId{}

// MeOnenoteNotebookIdSectionGroupIdSectionId is a struct representing the Resource ID for a Me Onenote Notebook Id Section Group Id Section
type MeOnenoteNotebookIdSectionGroupIdSectionId struct {
	NotebookId       string
	SectionGroupId   string
	OnenoteSectionId string
}

// NewMeOnenoteNotebookIdSectionGroupIdSectionID returns a new MeOnenoteNotebookIdSectionGroupIdSectionId struct
func NewMeOnenoteNotebookIdSectionGroupIdSectionID(notebookId string, sectionGroupId string, onenoteSectionId string) MeOnenoteNotebookIdSectionGroupIdSectionId {
	return MeOnenoteNotebookIdSectionGroupIdSectionId{
		NotebookId:       notebookId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseMeOnenoteNotebookIdSectionGroupIdSectionID parses 'input' into a MeOnenoteNotebookIdSectionGroupIdSectionId
func ParseMeOnenoteNotebookIdSectionGroupIdSectionID(input string) (*MeOnenoteNotebookIdSectionGroupIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteNotebookIdSectionGroupIdSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteNotebookIdSectionGroupIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOnenoteNotebookIdSectionGroupIdSectionIDInsensitively parses 'input' case-insensitively into a MeOnenoteNotebookIdSectionGroupIdSectionId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteNotebookIdSectionGroupIdSectionIDInsensitively(input string) (*MeOnenoteNotebookIdSectionGroupIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteNotebookIdSectionGroupIdSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteNotebookIdSectionGroupIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOnenoteNotebookIdSectionGroupIdSectionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

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

// ValidateMeOnenoteNotebookIdSectionGroupIdSectionID checks that 'input' can be parsed as a Me Onenote Notebook Id Section Group Id Section ID
func ValidateMeOnenoteNotebookIdSectionGroupIdSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteNotebookIdSectionGroupIdSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Notebook Id Section Group Id Section ID
func (id MeOnenoteNotebookIdSectionGroupIdSectionId) ID() string {
	fmtString := "/me/onenote/notebooks/%s/sectionGroups/%s/sections/%s"
	return fmt.Sprintf(fmtString, id.NotebookId, id.SectionGroupId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Notebook Id Section Group Id Section ID
func (id MeOnenoteNotebookIdSectionGroupIdSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("notebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookId"),
		resourceids.StaticSegment("sectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupId"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
	}
}

// String returns a human-readable description of this Me Onenote Notebook Id Section Group Id Section ID
func (id MeOnenoteNotebookIdSectionGroupIdSectionId) String() string {
	components := []string{
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("Me Onenote Notebook Id Section Group Id Section (%s)", strings.Join(components, "\n"))
}
