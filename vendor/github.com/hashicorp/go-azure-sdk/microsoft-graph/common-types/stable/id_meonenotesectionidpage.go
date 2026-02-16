package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnenoteSectionIdPageId{}

// MeOnenoteSectionIdPageId is a struct representing the Resource ID for a Me Onenote Section Id Page
type MeOnenoteSectionIdPageId struct {
	OnenoteSectionId string
	OnenotePageId    string
}

// NewMeOnenoteSectionIdPageID returns a new MeOnenoteSectionIdPageId struct
func NewMeOnenoteSectionIdPageID(onenoteSectionId string, onenotePageId string) MeOnenoteSectionIdPageId {
	return MeOnenoteSectionIdPageId{
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseMeOnenoteSectionIdPageID parses 'input' into a MeOnenoteSectionIdPageId
func ParseMeOnenoteSectionIdPageID(input string) (*MeOnenoteSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteSectionIdPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOnenoteSectionIdPageIDInsensitively parses 'input' case-insensitively into a MeOnenoteSectionIdPageId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteSectionIdPageIDInsensitively(input string) (*MeOnenoteSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteSectionIdPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOnenoteSectionIdPageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OnenoteSectionId, ok = input.Parsed["onenoteSectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", input)
	}

	if id.OnenotePageId, ok = input.Parsed["onenotePageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", input)
	}

	return nil
}

// ValidateMeOnenoteSectionIdPageID checks that 'input' can be parsed as a Me Onenote Section Id Page ID
func ValidateMeOnenoteSectionIdPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteSectionIdPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Section Id Page ID
func (id MeOnenoteSectionIdPageId) ID() string {
	fmtString := "/me/onenote/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Section Id Page ID
func (id MeOnenoteSectionIdPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
		resourceids.StaticSegment("pages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageId"),
	}
}

// String returns a human-readable description of this Me Onenote Section Id Page ID
func (id MeOnenoteSectionIdPageId) String() string {
	components := []string{
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Me Onenote Section Id Page (%s)", strings.Join(components, "\n"))
}
