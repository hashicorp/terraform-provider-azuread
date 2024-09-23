package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdOnenoteSectionGroupIdSectionId{}

// GroupIdSiteIdOnenoteSectionGroupIdSectionId is a struct representing the Resource ID for a Group Id Site Id Onenote Section Group Id Section
type GroupIdSiteIdOnenoteSectionGroupIdSectionId struct {
	GroupId          string
	SiteId           string
	SectionGroupId   string
	OnenoteSectionId string
}

// NewGroupIdSiteIdOnenoteSectionGroupIdSectionID returns a new GroupIdSiteIdOnenoteSectionGroupIdSectionId struct
func NewGroupIdSiteIdOnenoteSectionGroupIdSectionID(groupId string, siteId string, sectionGroupId string, onenoteSectionId string) GroupIdSiteIdOnenoteSectionGroupIdSectionId {
	return GroupIdSiteIdOnenoteSectionGroupIdSectionId{
		GroupId:          groupId,
		SiteId:           siteId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseGroupIdSiteIdOnenoteSectionGroupIdSectionID parses 'input' into a GroupIdSiteIdOnenoteSectionGroupIdSectionId
func ParseGroupIdSiteIdOnenoteSectionGroupIdSectionID(input string) (*GroupIdSiteIdOnenoteSectionGroupIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenoteSectionGroupIdSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenoteSectionGroupIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdOnenoteSectionGroupIdSectionIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdOnenoteSectionGroupIdSectionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdOnenoteSectionGroupIdSectionIDInsensitively(input string) (*GroupIdSiteIdOnenoteSectionGroupIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenoteSectionGroupIdSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenoteSectionGroupIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdOnenoteSectionGroupIdSectionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.SectionGroupId, ok = input.Parsed["sectionGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", input)
	}

	if id.OnenoteSectionId, ok = input.Parsed["onenoteSectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdOnenoteSectionGroupIdSectionID checks that 'input' can be parsed as a Group Id Site Id Onenote Section Group Id Section ID
func ValidateGroupIdSiteIdOnenoteSectionGroupIdSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdOnenoteSectionGroupIdSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Onenote Section Group Id Section ID
func (id GroupIdSiteIdOnenoteSectionGroupIdSectionId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/sectionGroups/%s/sections/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SectionGroupId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Onenote Section Group Id Section ID
func (id GroupIdSiteIdOnenoteSectionGroupIdSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("sectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupId"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Onenote Section Group Id Section ID
func (id GroupIdSiteIdOnenoteSectionGroupIdSectionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("Group Id Site Id Onenote Section Group Id Section (%s)", strings.Join(components, "\n"))
}
