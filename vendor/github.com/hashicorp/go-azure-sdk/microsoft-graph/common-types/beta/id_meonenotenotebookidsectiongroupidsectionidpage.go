package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnenoteNotebookIdSectionGroupIdSectionIdPageId{}

// MeOnenoteNotebookIdSectionGroupIdSectionIdPageId is a struct representing the Resource ID for a Me Onenote Notebook Id Section Group Id Section Id Page
type MeOnenoteNotebookIdSectionGroupIdSectionIdPageId struct {
	NotebookId       string
	SectionGroupId   string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewMeOnenoteNotebookIdSectionGroupIdSectionIdPageID returns a new MeOnenoteNotebookIdSectionGroupIdSectionIdPageId struct
func NewMeOnenoteNotebookIdSectionGroupIdSectionIdPageID(notebookId string, sectionGroupId string, onenoteSectionId string, onenotePageId string) MeOnenoteNotebookIdSectionGroupIdSectionIdPageId {
	return MeOnenoteNotebookIdSectionGroupIdSectionIdPageId{
		NotebookId:       notebookId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseMeOnenoteNotebookIdSectionGroupIdSectionIdPageID parses 'input' into a MeOnenoteNotebookIdSectionGroupIdSectionIdPageId
func ParseMeOnenoteNotebookIdSectionGroupIdSectionIdPageID(input string) (*MeOnenoteNotebookIdSectionGroupIdSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteNotebookIdSectionGroupIdSectionIdPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteNotebookIdSectionGroupIdSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOnenoteNotebookIdSectionGroupIdSectionIdPageIDInsensitively parses 'input' case-insensitively into a MeOnenoteNotebookIdSectionGroupIdSectionIdPageId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteNotebookIdSectionGroupIdSectionIdPageIDInsensitively(input string) (*MeOnenoteNotebookIdSectionGroupIdSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteNotebookIdSectionGroupIdSectionIdPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteNotebookIdSectionGroupIdSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOnenoteNotebookIdSectionGroupIdSectionIdPageId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.OnenotePageId, ok = input.Parsed["onenotePageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", input)
	}

	return nil
}

// ValidateMeOnenoteNotebookIdSectionGroupIdSectionIdPageID checks that 'input' can be parsed as a Me Onenote Notebook Id Section Group Id Section Id Page ID
func ValidateMeOnenoteNotebookIdSectionGroupIdSectionIdPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteNotebookIdSectionGroupIdSectionIdPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Notebook Id Section Group Id Section Id Page ID
func (id MeOnenoteNotebookIdSectionGroupIdSectionIdPageId) ID() string {
	fmtString := "/me/onenote/notebooks/%s/sectionGroups/%s/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.NotebookId, id.SectionGroupId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Notebook Id Section Group Id Section Id Page ID
func (id MeOnenoteNotebookIdSectionGroupIdSectionIdPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("notebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookId"),
		resourceids.StaticSegment("sectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupId"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
		resourceids.StaticSegment("pages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageId"),
	}
}

// String returns a human-readable description of this Me Onenote Notebook Id Section Group Id Section Id Page ID
func (id MeOnenoteNotebookIdSectionGroupIdSectionIdPageId) String() string {
	components := []string{
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Me Onenote Notebook Id Section Group Id Section Id Page (%s)", strings.Join(components, "\n"))
}
