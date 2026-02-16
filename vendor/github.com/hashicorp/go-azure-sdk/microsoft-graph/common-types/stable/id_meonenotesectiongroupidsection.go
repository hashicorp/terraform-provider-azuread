package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnenoteSectionGroupIdSectionId{}

// MeOnenoteSectionGroupIdSectionId is a struct representing the Resource ID for a Me Onenote Section Group Id Section
type MeOnenoteSectionGroupIdSectionId struct {
	SectionGroupId   string
	OnenoteSectionId string
}

// NewMeOnenoteSectionGroupIdSectionID returns a new MeOnenoteSectionGroupIdSectionId struct
func NewMeOnenoteSectionGroupIdSectionID(sectionGroupId string, onenoteSectionId string) MeOnenoteSectionGroupIdSectionId {
	return MeOnenoteSectionGroupIdSectionId{
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseMeOnenoteSectionGroupIdSectionID parses 'input' into a MeOnenoteSectionGroupIdSectionId
func ParseMeOnenoteSectionGroupIdSectionID(input string) (*MeOnenoteSectionGroupIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteSectionGroupIdSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteSectionGroupIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOnenoteSectionGroupIdSectionIDInsensitively parses 'input' case-insensitively into a MeOnenoteSectionGroupIdSectionId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteSectionGroupIdSectionIDInsensitively(input string) (*MeOnenoteSectionGroupIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteSectionGroupIdSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteSectionGroupIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOnenoteSectionGroupIdSectionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SectionGroupId, ok = input.Parsed["sectionGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", input)
	}

	if id.OnenoteSectionId, ok = input.Parsed["onenoteSectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", input)
	}

	return nil
}

// ValidateMeOnenoteSectionGroupIdSectionID checks that 'input' can be parsed as a Me Onenote Section Group Id Section ID
func ValidateMeOnenoteSectionGroupIdSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteSectionGroupIdSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Section Group Id Section ID
func (id MeOnenoteSectionGroupIdSectionId) ID() string {
	fmtString := "/me/onenote/sectionGroups/%s/sections/%s"
	return fmt.Sprintf(fmtString, id.SectionGroupId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Section Group Id Section ID
func (id MeOnenoteSectionGroupIdSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("sectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupId"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
	}
}

// String returns a human-readable description of this Me Onenote Section Group Id Section ID
func (id MeOnenoteSectionGroupIdSectionId) String() string {
	components := []string{
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("Me Onenote Section Group Id Section (%s)", strings.Join(components, "\n"))
}
