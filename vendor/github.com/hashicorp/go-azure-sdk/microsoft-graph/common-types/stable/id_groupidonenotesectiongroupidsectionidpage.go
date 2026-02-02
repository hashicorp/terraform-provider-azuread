package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdOnenoteSectionGroupIdSectionIdPageId{}

// GroupIdOnenoteSectionGroupIdSectionIdPageId is a struct representing the Resource ID for a Group Id Onenote Section Group Id Section Id Page
type GroupIdOnenoteSectionGroupIdSectionIdPageId struct {
	GroupId          string
	SectionGroupId   string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewGroupIdOnenoteSectionGroupIdSectionIdPageID returns a new GroupIdOnenoteSectionGroupIdSectionIdPageId struct
func NewGroupIdOnenoteSectionGroupIdSectionIdPageID(groupId string, sectionGroupId string, onenoteSectionId string, onenotePageId string) GroupIdOnenoteSectionGroupIdSectionIdPageId {
	return GroupIdOnenoteSectionGroupIdSectionIdPageId{
		GroupId:          groupId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseGroupIdOnenoteSectionGroupIdSectionIdPageID parses 'input' into a GroupIdOnenoteSectionGroupIdSectionIdPageId
func ParseGroupIdOnenoteSectionGroupIdSectionIdPageID(input string) (*GroupIdOnenoteSectionGroupIdSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteSectionGroupIdSectionIdPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteSectionGroupIdSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdOnenoteSectionGroupIdSectionIdPageIDInsensitively parses 'input' case-insensitively into a GroupIdOnenoteSectionGroupIdSectionIdPageId
// note: this method should only be used for API response data and not user input
func ParseGroupIdOnenoteSectionGroupIdSectionIdPageIDInsensitively(input string) (*GroupIdOnenoteSectionGroupIdSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteSectionGroupIdSectionIdPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteSectionGroupIdSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdOnenoteSectionGroupIdSectionIdPageId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.OnenotePageId, ok = input.Parsed["onenotePageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", input)
	}

	return nil
}

// ValidateGroupIdOnenoteSectionGroupIdSectionIdPageID checks that 'input' can be parsed as a Group Id Onenote Section Group Id Section Id Page ID
func ValidateGroupIdOnenoteSectionGroupIdSectionIdPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdOnenoteSectionGroupIdSectionIdPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Onenote Section Group Id Section Id Page ID
func (id GroupIdOnenoteSectionGroupIdSectionIdPageId) ID() string {
	fmtString := "/groups/%s/onenote/sectionGroups/%s/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SectionGroupId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Onenote Section Group Id Section Id Page ID
func (id GroupIdOnenoteSectionGroupIdSectionIdPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("sectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupId"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
		resourceids.StaticSegment("pages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageId"),
	}
}

// String returns a human-readable description of this Group Id Onenote Section Group Id Section Id Page ID
func (id GroupIdOnenoteSectionGroupIdSectionIdPageId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Group Id Onenote Section Group Id Section Id Page (%s)", strings.Join(components, "\n"))
}
