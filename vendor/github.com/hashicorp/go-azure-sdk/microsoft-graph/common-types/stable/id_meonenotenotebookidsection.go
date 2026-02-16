package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnenoteNotebookIdSectionId{}

// MeOnenoteNotebookIdSectionId is a struct representing the Resource ID for a Me Onenote Notebook Id Section
type MeOnenoteNotebookIdSectionId struct {
	NotebookId       string
	OnenoteSectionId string
}

// NewMeOnenoteNotebookIdSectionID returns a new MeOnenoteNotebookIdSectionId struct
func NewMeOnenoteNotebookIdSectionID(notebookId string, onenoteSectionId string) MeOnenoteNotebookIdSectionId {
	return MeOnenoteNotebookIdSectionId{
		NotebookId:       notebookId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseMeOnenoteNotebookIdSectionID parses 'input' into a MeOnenoteNotebookIdSectionId
func ParseMeOnenoteNotebookIdSectionID(input string) (*MeOnenoteNotebookIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteNotebookIdSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteNotebookIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOnenoteNotebookIdSectionIDInsensitively parses 'input' case-insensitively into a MeOnenoteNotebookIdSectionId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteNotebookIdSectionIDInsensitively(input string) (*MeOnenoteNotebookIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteNotebookIdSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteNotebookIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOnenoteNotebookIdSectionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.NotebookId, ok = input.Parsed["notebookId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "notebookId", input)
	}

	if id.OnenoteSectionId, ok = input.Parsed["onenoteSectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", input)
	}

	return nil
}

// ValidateMeOnenoteNotebookIdSectionID checks that 'input' can be parsed as a Me Onenote Notebook Id Section ID
func ValidateMeOnenoteNotebookIdSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteNotebookIdSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Notebook Id Section ID
func (id MeOnenoteNotebookIdSectionId) ID() string {
	fmtString := "/me/onenote/notebooks/%s/sections/%s"
	return fmt.Sprintf(fmtString, id.NotebookId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Notebook Id Section ID
func (id MeOnenoteNotebookIdSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("notebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookId"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
	}
}

// String returns a human-readable description of this Me Onenote Notebook Id Section ID
func (id MeOnenoteNotebookIdSectionId) String() string {
	components := []string{
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("Me Onenote Notebook Id Section (%s)", strings.Join(components, "\n"))
}
