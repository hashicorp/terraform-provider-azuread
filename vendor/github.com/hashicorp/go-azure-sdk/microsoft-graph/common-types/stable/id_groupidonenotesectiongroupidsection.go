package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdOnenoteSectionGroupIdSectionId{}

// GroupIdOnenoteSectionGroupIdSectionId is a struct representing the Resource ID for a Group Id Onenote Section Group Id Section
type GroupIdOnenoteSectionGroupIdSectionId struct {
	GroupId          string
	SectionGroupId   string
	OnenoteSectionId string
}

// NewGroupIdOnenoteSectionGroupIdSectionID returns a new GroupIdOnenoteSectionGroupIdSectionId struct
func NewGroupIdOnenoteSectionGroupIdSectionID(groupId string, sectionGroupId string, onenoteSectionId string) GroupIdOnenoteSectionGroupIdSectionId {
	return GroupIdOnenoteSectionGroupIdSectionId{
		GroupId:          groupId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseGroupIdOnenoteSectionGroupIdSectionID parses 'input' into a GroupIdOnenoteSectionGroupIdSectionId
func ParseGroupIdOnenoteSectionGroupIdSectionID(input string) (*GroupIdOnenoteSectionGroupIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteSectionGroupIdSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteSectionGroupIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdOnenoteSectionGroupIdSectionIDInsensitively parses 'input' case-insensitively into a GroupIdOnenoteSectionGroupIdSectionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdOnenoteSectionGroupIdSectionIDInsensitively(input string) (*GroupIdOnenoteSectionGroupIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteSectionGroupIdSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteSectionGroupIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdOnenoteSectionGroupIdSectionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SectionGroupId, ok = input.Parsed["sectionGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", input)
	}

	if id.OnenoteSectionId, ok = input.Parsed["onenoteSectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", input)
	}

	return nil
}

// ValidateGroupIdOnenoteSectionGroupIdSectionID checks that 'input' can be parsed as a Group Id Onenote Section Group Id Section ID
func ValidateGroupIdOnenoteSectionGroupIdSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdOnenoteSectionGroupIdSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Onenote Section Group Id Section ID
func (id GroupIdOnenoteSectionGroupIdSectionId) ID() string {
	fmtString := "/groups/%s/onenote/sectionGroups/%s/sections/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SectionGroupId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Onenote Section Group Id Section ID
func (id GroupIdOnenoteSectionGroupIdSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("sectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupId"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
	}
}

// String returns a human-readable description of this Group Id Onenote Section Group Id Section ID
func (id GroupIdOnenoteSectionGroupIdSectionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("Group Id Onenote Section Group Id Section (%s)", strings.Join(components, "\n"))
}
