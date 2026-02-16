package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdOnenoteNotebookIdSectionGroupId{}

// GroupIdSiteIdOnenoteNotebookIdSectionGroupId is a struct representing the Resource ID for a Group Id Site Id Onenote Notebook Id Section Group
type GroupIdSiteIdOnenoteNotebookIdSectionGroupId struct {
	GroupId        string
	SiteId         string
	NotebookId     string
	SectionGroupId string
}

// NewGroupIdSiteIdOnenoteNotebookIdSectionGroupID returns a new GroupIdSiteIdOnenoteNotebookIdSectionGroupId struct
func NewGroupIdSiteIdOnenoteNotebookIdSectionGroupID(groupId string, siteId string, notebookId string, sectionGroupId string) GroupIdSiteIdOnenoteNotebookIdSectionGroupId {
	return GroupIdSiteIdOnenoteNotebookIdSectionGroupId{
		GroupId:        groupId,
		SiteId:         siteId,
		NotebookId:     notebookId,
		SectionGroupId: sectionGroupId,
	}
}

// ParseGroupIdSiteIdOnenoteNotebookIdSectionGroupID parses 'input' into a GroupIdSiteIdOnenoteNotebookIdSectionGroupId
func ParseGroupIdSiteIdOnenoteNotebookIdSectionGroupID(input string) (*GroupIdSiteIdOnenoteNotebookIdSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenoteNotebookIdSectionGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenoteNotebookIdSectionGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdOnenoteNotebookIdSectionGroupIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdOnenoteNotebookIdSectionGroupId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdOnenoteNotebookIdSectionGroupIDInsensitively(input string) (*GroupIdSiteIdOnenoteNotebookIdSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenoteNotebookIdSectionGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenoteNotebookIdSectionGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdOnenoteNotebookIdSectionGroupId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateGroupIdSiteIdOnenoteNotebookIdSectionGroupID checks that 'input' can be parsed as a Group Id Site Id Onenote Notebook Id Section Group ID
func ValidateGroupIdSiteIdOnenoteNotebookIdSectionGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdOnenoteNotebookIdSectionGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Onenote Notebook Id Section Group ID
func (id GroupIdSiteIdOnenoteNotebookIdSectionGroupId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/notebooks/%s/sectionGroups/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.NotebookId, id.SectionGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Onenote Notebook Id Section Group ID
func (id GroupIdSiteIdOnenoteNotebookIdSectionGroupId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this Group Id Site Id Onenote Notebook Id Section Group ID
func (id GroupIdSiteIdOnenoteNotebookIdSectionGroupId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
	}
	return fmt.Sprintf("Group Id Site Id Onenote Notebook Id Section Group (%s)", strings.Join(components, "\n"))
}
