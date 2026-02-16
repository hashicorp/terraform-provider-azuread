package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnenoteSectionId{}

// MeOnenoteSectionId is a struct representing the Resource ID for a Me Onenote Section
type MeOnenoteSectionId struct {
	OnenoteSectionId string
}

// NewMeOnenoteSectionID returns a new MeOnenoteSectionId struct
func NewMeOnenoteSectionID(onenoteSectionId string) MeOnenoteSectionId {
	return MeOnenoteSectionId{
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseMeOnenoteSectionID parses 'input' into a MeOnenoteSectionId
func ParseMeOnenoteSectionID(input string) (*MeOnenoteSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOnenoteSectionIDInsensitively parses 'input' case-insensitively into a MeOnenoteSectionId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteSectionIDInsensitively(input string) (*MeOnenoteSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOnenoteSectionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OnenoteSectionId, ok = input.Parsed["onenoteSectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", input)
	}

	return nil
}

// ValidateMeOnenoteSectionID checks that 'input' can be parsed as a Me Onenote Section ID
func ValidateMeOnenoteSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Section ID
func (id MeOnenoteSectionId) ID() string {
	fmtString := "/me/onenote/sections/%s"
	return fmt.Sprintf(fmtString, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Section ID
func (id MeOnenoteSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
	}
}

// String returns a human-readable description of this Me Onenote Section ID
func (id MeOnenoteSectionId) String() string {
	components := []string{
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("Me Onenote Section (%s)", strings.Join(components, "\n"))
}
