package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnenoteSectionGroupId{}

// MeOnenoteSectionGroupId is a struct representing the Resource ID for a Me Onenote Section Group
type MeOnenoteSectionGroupId struct {
	SectionGroupId string
}

// NewMeOnenoteSectionGroupID returns a new MeOnenoteSectionGroupId struct
func NewMeOnenoteSectionGroupID(sectionGroupId string) MeOnenoteSectionGroupId {
	return MeOnenoteSectionGroupId{
		SectionGroupId: sectionGroupId,
	}
}

// ParseMeOnenoteSectionGroupID parses 'input' into a MeOnenoteSectionGroupId
func ParseMeOnenoteSectionGroupID(input string) (*MeOnenoteSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteSectionGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteSectionGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOnenoteSectionGroupIDInsensitively parses 'input' case-insensitively into a MeOnenoteSectionGroupId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteSectionGroupIDInsensitively(input string) (*MeOnenoteSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnenoteSectionGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnenoteSectionGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOnenoteSectionGroupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SectionGroupId, ok = input.Parsed["sectionGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", input)
	}

	return nil
}

// ValidateMeOnenoteSectionGroupID checks that 'input' can be parsed as a Me Onenote Section Group ID
func ValidateMeOnenoteSectionGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteSectionGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Section Group ID
func (id MeOnenoteSectionGroupId) ID() string {
	fmtString := "/me/onenote/sectionGroups/%s"
	return fmt.Sprintf(fmtString, id.SectionGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Section Group ID
func (id MeOnenoteSectionGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("sectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupId"),
	}
}

// String returns a human-readable description of this Me Onenote Section Group ID
func (id MeOnenoteSectionGroupId) String() string {
	components := []string{
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
	}
	return fmt.Sprintf("Me Onenote Section Group (%s)", strings.Join(components, "\n"))
}
