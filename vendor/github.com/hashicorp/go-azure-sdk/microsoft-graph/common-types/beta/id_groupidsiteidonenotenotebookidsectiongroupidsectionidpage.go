package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageId{}

// GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageId is a struct representing the Resource ID for a Group Id Site Id Onenote Notebook Id Section Group Id Section Id Page
type GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageId struct {
	GroupId          string
	SiteId           string
	NotebookId       string
	SectionGroupId   string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageID returns a new GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageId struct
func NewGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageID(groupId string, siteId string, notebookId string, sectionGroupId string, onenoteSectionId string, onenotePageId string) GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageId {
	return GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageId{
		GroupId:          groupId,
		SiteId:           siteId,
		NotebookId:       notebookId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageID parses 'input' into a GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageId
func ParseGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageID(input string) (*GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageIDInsensitively(input string) (*GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.OnenotePageId, ok = input.Parsed["onenotePageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageID checks that 'input' can be parsed as a Group Id Site Id Onenote Notebook Id Section Group Id Section Id Page ID
func ValidateGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Onenote Notebook Id Section Group Id Section Id Page ID
func (id GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/notebooks/%s/sectionGroups/%s/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.NotebookId, id.SectionGroupId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Onenote Notebook Id Section Group Id Section Id Page ID
func (id GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("pages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Onenote Notebook Id Section Group Id Section Id Page ID
func (id GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Group Id Site Id Onenote Notebook Id Section Group Id Section Id Page (%s)", strings.Join(components, "\n"))
}
