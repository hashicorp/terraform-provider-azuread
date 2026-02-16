package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnenoteNotebookIdSectionGroupId{}

// MeOnenoteNotebookIdSectionGroupId is a struct representing the Resource ID for a Me Onenote Notebook Id Section Group
type MeOnenoteNotebookIdSectionGroupId struct {
	NotebookId     string
	SectionGroupId string
}

// NewMeOnenoteNotebookIdSectionGroupID returns a new MeOnenoteNotebookIdSectionGroupId struct
func NewMeOnenoteNotebookIdSectionGroupID(notebookId string, sectionGroupId string) MeOnenoteNotebookIdSectionGroupId {
	return MeOnenoteNotebookIdSectionGroupId{
		NotebookId:     notebookId,
		SectionGroupId: sectionGroupId,
	}
}

// ParseMeOnenoteNotebookIdSectionGroupID parses 'input' into a MeOnenoteNotebookIdSectionGroupId
func ParseMeOnenoteNotebookIdSectionGroupID(input string) (*MeOnenoteNotebookIdSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteNotebookIdSectionGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteNotebookIdSectionGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOnenoteNotebookIdSectionGroupIDInsensitively parses 'input' case-insensitively into a MeOnenoteNotebookIdSectionGroupId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteNotebookIdSectionGroupIDInsensitively(input string) (*MeOnenoteNotebookIdSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteNotebookIdSectionGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteNotebookIdSectionGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOnenoteNotebookIdSectionGroupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.NotebookId, ok = input.Parsed["notebookId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "notebookId", input)
	}

	if id.SectionGroupId, ok = input.Parsed["sectionGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", input)
	}

	return nil
}

// ValidateMeOnenoteNotebookIdSectionGroupID checks that 'input' can be parsed as a Me Onenote Notebook Id Section Group ID
func ValidateMeOnenoteNotebookIdSectionGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteNotebookIdSectionGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Notebook Id Section Group ID
func (id MeOnenoteNotebookIdSectionGroupId) ID() string {
	fmtString := "/me/onenote/notebooks/%s/sectionGroups/%s"
	return fmt.Sprintf(fmtString, id.NotebookId, id.SectionGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Notebook Id Section Group ID
func (id MeOnenoteNotebookIdSectionGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("notebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookId"),
		resourceids.StaticSegment("sectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupId"),
	}
}

// String returns a human-readable description of this Me Onenote Notebook Id Section Group ID
func (id MeOnenoteNotebookIdSectionGroupId) String() string {
	components := []string{
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
	}
	return fmt.Sprintf("Me Onenote Notebook Id Section Group (%s)", strings.Join(components, "\n"))
}
