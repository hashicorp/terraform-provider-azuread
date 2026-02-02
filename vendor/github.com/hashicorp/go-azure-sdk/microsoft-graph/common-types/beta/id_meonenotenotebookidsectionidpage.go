package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnenoteNotebookIdSectionIdPageId{}

// MeOnenoteNotebookIdSectionIdPageId is a struct representing the Resource ID for a Me Onenote Notebook Id Section Id Page
type MeOnenoteNotebookIdSectionIdPageId struct {
	NotebookId       string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewMeOnenoteNotebookIdSectionIdPageID returns a new MeOnenoteNotebookIdSectionIdPageId struct
func NewMeOnenoteNotebookIdSectionIdPageID(notebookId string, onenoteSectionId string, onenotePageId string) MeOnenoteNotebookIdSectionIdPageId {
	return MeOnenoteNotebookIdSectionIdPageId{
		NotebookId:       notebookId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseMeOnenoteNotebookIdSectionIdPageID parses 'input' into a MeOnenoteNotebookIdSectionIdPageId
func ParseMeOnenoteNotebookIdSectionIdPageID(input string) (*MeOnenoteNotebookIdSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteNotebookIdSectionIdPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteNotebookIdSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOnenoteNotebookIdSectionIdPageIDInsensitively parses 'input' case-insensitively into a MeOnenoteNotebookIdSectionIdPageId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteNotebookIdSectionIdPageIDInsensitively(input string) (*MeOnenoteNotebookIdSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteNotebookIdSectionIdPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteNotebookIdSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOnenoteNotebookIdSectionIdPageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.NotebookId, ok = input.Parsed["notebookId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "notebookId", input)
	}

	if id.OnenoteSectionId, ok = input.Parsed["onenoteSectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", input)
	}

	if id.OnenotePageId, ok = input.Parsed["onenotePageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", input)
	}

	return nil
}

// ValidateMeOnenoteNotebookIdSectionIdPageID checks that 'input' can be parsed as a Me Onenote Notebook Id Section Id Page ID
func ValidateMeOnenoteNotebookIdSectionIdPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteNotebookIdSectionIdPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Notebook Id Section Id Page ID
func (id MeOnenoteNotebookIdSectionIdPageId) ID() string {
	fmtString := "/me/onenote/notebooks/%s/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.NotebookId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Notebook Id Section Id Page ID
func (id MeOnenoteNotebookIdSectionIdPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("notebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookId"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
		resourceids.StaticSegment("pages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageId"),
	}
}

// String returns a human-readable description of this Me Onenote Notebook Id Section Id Page ID
func (id MeOnenoteNotebookIdSectionIdPageId) String() string {
	components := []string{
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Me Onenote Notebook Id Section Id Page (%s)", strings.Join(components, "\n"))
}
