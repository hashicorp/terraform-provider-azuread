package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionId{}

// GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionId is a struct representing the Resource ID for a Group Id Site Id Onenote Notebook Id Section Group Id Section
type GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionId struct {
	GroupId          string
	SiteId           string
	NotebookId       string
	SectionGroupId   string
	OnenoteSectionId string
}

// NewGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionID returns a new GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionId struct
func NewGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionID(groupId string, siteId string, notebookId string, sectionGroupId string, onenoteSectionId string) GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionId {
	return GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionId{
		GroupId:          groupId,
		SiteId:           siteId,
		NotebookId:       notebookId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionID parses 'input' into a GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionId
func ParseGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionID(input string) (*GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIDInsensitively(input string) (*GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.NotebookId, ok = input.Parsed["notebookId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "notebookId", input)
	}

	if id.SectionGroupId, ok = input.Parsed["sectionGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", input)
	}

	if id.OnenoteSectionId, ok = input.Parsed["onenoteSectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionID checks that 'input' can be parsed as a Group Id Site Id Onenote Notebook Id Section Group Id Section ID
func ValidateGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Onenote Notebook Id Section Group Id Section ID
func (id GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/notebooks/%s/sectionGroups/%s/sections/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.NotebookId, id.SectionGroupId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Onenote Notebook Id Section Group Id Section ID
func (id GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("notebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookId"),
		resourceids.StaticSegment("sectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupId"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Onenote Notebook Id Section Group Id Section ID
func (id GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("Group Id Site Id Onenote Notebook Id Section Group Id Section (%s)", strings.Join(components, "\n"))
}
