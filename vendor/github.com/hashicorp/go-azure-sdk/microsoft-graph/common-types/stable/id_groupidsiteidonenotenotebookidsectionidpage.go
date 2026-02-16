package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdOnenoteNotebookIdSectionIdPageId{}

// GroupIdSiteIdOnenoteNotebookIdSectionIdPageId is a struct representing the Resource ID for a Group Id Site Id Onenote Notebook Id Section Id Page
type GroupIdSiteIdOnenoteNotebookIdSectionIdPageId struct {
	GroupId          string
	SiteId           string
	NotebookId       string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewGroupIdSiteIdOnenoteNotebookIdSectionIdPageID returns a new GroupIdSiteIdOnenoteNotebookIdSectionIdPageId struct
func NewGroupIdSiteIdOnenoteNotebookIdSectionIdPageID(groupId string, siteId string, notebookId string, onenoteSectionId string, onenotePageId string) GroupIdSiteIdOnenoteNotebookIdSectionIdPageId {
	return GroupIdSiteIdOnenoteNotebookIdSectionIdPageId{
		GroupId:          groupId,
		SiteId:           siteId,
		NotebookId:       notebookId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseGroupIdSiteIdOnenoteNotebookIdSectionIdPageID parses 'input' into a GroupIdSiteIdOnenoteNotebookIdSectionIdPageId
func ParseGroupIdSiteIdOnenoteNotebookIdSectionIdPageID(input string) (*GroupIdSiteIdOnenoteNotebookIdSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenoteNotebookIdSectionIdPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenoteNotebookIdSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdOnenoteNotebookIdSectionIdPageIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdOnenoteNotebookIdSectionIdPageId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdOnenoteNotebookIdSectionIdPageIDInsensitively(input string) (*GroupIdSiteIdOnenoteNotebookIdSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenoteNotebookIdSectionIdPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenoteNotebookIdSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdOnenoteNotebookIdSectionIdPageId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.OnenoteSectionId, ok = input.Parsed["onenoteSectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", input)
	}

	if id.OnenotePageId, ok = input.Parsed["onenotePageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdOnenoteNotebookIdSectionIdPageID checks that 'input' can be parsed as a Group Id Site Id Onenote Notebook Id Section Id Page ID
func ValidateGroupIdSiteIdOnenoteNotebookIdSectionIdPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdOnenoteNotebookIdSectionIdPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Onenote Notebook Id Section Id Page ID
func (id GroupIdSiteIdOnenoteNotebookIdSectionIdPageId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/notebooks/%s/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.NotebookId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Onenote Notebook Id Section Id Page ID
func (id GroupIdSiteIdOnenoteNotebookIdSectionIdPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("notebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookId"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
		resourceids.StaticSegment("pages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Onenote Notebook Id Section Id Page ID
func (id GroupIdSiteIdOnenoteNotebookIdSectionIdPageId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Group Id Site Id Onenote Notebook Id Section Id Page (%s)", strings.Join(components, "\n"))
}
