package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnenoteNotebookId{}

// MeOnenoteNotebookId is a struct representing the Resource ID for a Me Onenote Notebook
type MeOnenoteNotebookId struct {
	NotebookId string
}

// NewMeOnenoteNotebookID returns a new MeOnenoteNotebookId struct
func NewMeOnenoteNotebookID(notebookId string) MeOnenoteNotebookId {
	return MeOnenoteNotebookId{
		NotebookId: notebookId,
	}
}

// ParseMeOnenoteNotebookID parses 'input' into a MeOnenoteNotebookId
func ParseMeOnenoteNotebookID(input string) (*MeOnenoteNotebookId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteNotebookId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteNotebookId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOnenoteNotebookIDInsensitively parses 'input' case-insensitively into a MeOnenoteNotebookId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteNotebookIDInsensitively(input string) (*MeOnenoteNotebookId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteNotebookId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteNotebookId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOnenoteNotebookId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.NotebookId, ok = input.Parsed["notebookId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "notebookId", input)
	}

	return nil
}

// ValidateMeOnenoteNotebookID checks that 'input' can be parsed as a Me Onenote Notebook ID
func ValidateMeOnenoteNotebookID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteNotebookID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Notebook ID
func (id MeOnenoteNotebookId) ID() string {
	fmtString := "/me/onenote/notebooks/%s"
	return fmt.Sprintf(fmtString, id.NotebookId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Notebook ID
func (id MeOnenoteNotebookId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("notebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookId"),
	}
}

// String returns a human-readable description of this Me Onenote Notebook ID
func (id MeOnenoteNotebookId) String() string {
	components := []string{
		fmt.Sprintf("Notebook: %q", id.NotebookId),
	}
	return fmt.Sprintf("Me Onenote Notebook (%s)", strings.Join(components, "\n"))
}
