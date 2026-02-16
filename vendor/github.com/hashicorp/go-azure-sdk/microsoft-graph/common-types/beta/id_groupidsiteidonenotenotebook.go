package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdOnenoteNotebookId{}

// GroupIdSiteIdOnenoteNotebookId is a struct representing the Resource ID for a Group Id Site Id Onenote Notebook
type GroupIdSiteIdOnenoteNotebookId struct {
	GroupId    string
	SiteId     string
	NotebookId string
}

// NewGroupIdSiteIdOnenoteNotebookID returns a new GroupIdSiteIdOnenoteNotebookId struct
func NewGroupIdSiteIdOnenoteNotebookID(groupId string, siteId string, notebookId string) GroupIdSiteIdOnenoteNotebookId {
	return GroupIdSiteIdOnenoteNotebookId{
		GroupId:    groupId,
		SiteId:     siteId,
		NotebookId: notebookId,
	}
}

// ParseGroupIdSiteIdOnenoteNotebookID parses 'input' into a GroupIdSiteIdOnenoteNotebookId
func ParseGroupIdSiteIdOnenoteNotebookID(input string) (*GroupIdSiteIdOnenoteNotebookId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenoteNotebookId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenoteNotebookId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdOnenoteNotebookIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdOnenoteNotebookId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdOnenoteNotebookIDInsensitively(input string) (*GroupIdSiteIdOnenoteNotebookId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenoteNotebookId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenoteNotebookId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdOnenoteNotebookId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateGroupIdSiteIdOnenoteNotebookID checks that 'input' can be parsed as a Group Id Site Id Onenote Notebook ID
func ValidateGroupIdSiteIdOnenoteNotebookID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdOnenoteNotebookID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Onenote Notebook ID
func (id GroupIdSiteIdOnenoteNotebookId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/notebooks/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.NotebookId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Onenote Notebook ID
func (id GroupIdSiteIdOnenoteNotebookId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("notebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Onenote Notebook ID
func (id GroupIdSiteIdOnenoteNotebookId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
	}
	return fmt.Sprintf("Group Id Site Id Onenote Notebook (%s)", strings.Join(components, "\n"))
}
