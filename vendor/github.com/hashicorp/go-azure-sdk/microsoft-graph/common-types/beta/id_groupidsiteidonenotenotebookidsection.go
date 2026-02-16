package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdOnenoteNotebookIdSectionId{}

// GroupIdSiteIdOnenoteNotebookIdSectionId is a struct representing the Resource ID for a Group Id Site Id Onenote Notebook Id Section
type GroupIdSiteIdOnenoteNotebookIdSectionId struct {
	GroupId          string
	SiteId           string
	NotebookId       string
	OnenoteSectionId string
}

// NewGroupIdSiteIdOnenoteNotebookIdSectionID returns a new GroupIdSiteIdOnenoteNotebookIdSectionId struct
func NewGroupIdSiteIdOnenoteNotebookIdSectionID(groupId string, siteId string, notebookId string, onenoteSectionId string) GroupIdSiteIdOnenoteNotebookIdSectionId {
	return GroupIdSiteIdOnenoteNotebookIdSectionId{
		GroupId:          groupId,
		SiteId:           siteId,
		NotebookId:       notebookId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseGroupIdSiteIdOnenoteNotebookIdSectionID parses 'input' into a GroupIdSiteIdOnenoteNotebookIdSectionId
func ParseGroupIdSiteIdOnenoteNotebookIdSectionID(input string) (*GroupIdSiteIdOnenoteNotebookIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenoteNotebookIdSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenoteNotebookIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdOnenoteNotebookIdSectionIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdOnenoteNotebookIdSectionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdOnenoteNotebookIdSectionIDInsensitively(input string) (*GroupIdSiteIdOnenoteNotebookIdSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenoteNotebookIdSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenoteNotebookIdSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdOnenoteNotebookIdSectionId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateGroupIdSiteIdOnenoteNotebookIdSectionID checks that 'input' can be parsed as a Group Id Site Id Onenote Notebook Id Section ID
func ValidateGroupIdSiteIdOnenoteNotebookIdSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdOnenoteNotebookIdSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Onenote Notebook Id Section ID
func (id GroupIdSiteIdOnenoteNotebookIdSectionId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/notebooks/%s/sections/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.NotebookId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Onenote Notebook Id Section ID
func (id GroupIdSiteIdOnenoteNotebookIdSectionId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this Group Id Site Id Onenote Notebook Id Section ID
func (id GroupIdSiteIdOnenoteNotebookIdSectionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("Group Id Site Id Onenote Notebook Id Section (%s)", strings.Join(components, "\n"))
}
