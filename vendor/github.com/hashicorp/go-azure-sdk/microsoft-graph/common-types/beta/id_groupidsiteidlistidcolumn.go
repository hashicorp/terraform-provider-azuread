package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdListIdColumnId{}

// GroupIdSiteIdListIdColumnId is a struct representing the Resource ID for a Group Id Site Id List Id Column
type GroupIdSiteIdListIdColumnId struct {
	GroupId            string
	SiteId             string
	ListId             string
	ColumnDefinitionId string
}

// NewGroupIdSiteIdListIdColumnID returns a new GroupIdSiteIdListIdColumnId struct
func NewGroupIdSiteIdListIdColumnID(groupId string, siteId string, listId string, columnDefinitionId string) GroupIdSiteIdListIdColumnId {
	return GroupIdSiteIdListIdColumnId{
		GroupId:            groupId,
		SiteId:             siteId,
		ListId:             listId,
		ColumnDefinitionId: columnDefinitionId,
	}
}

// ParseGroupIdSiteIdListIdColumnID parses 'input' into a GroupIdSiteIdListIdColumnId
func ParseGroupIdSiteIdListIdColumnID(input string) (*GroupIdSiteIdListIdColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdColumnId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdListIdColumnIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdListIdColumnId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdListIdColumnIDInsensitively(input string) (*GroupIdSiteIdListIdColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdColumnId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdListIdColumnId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.ListId, ok = input.Parsed["listId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "listId", input)
	}

	if id.ColumnDefinitionId, ok = input.Parsed["columnDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "columnDefinitionId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdListIdColumnID checks that 'input' can be parsed as a Group Id Site Id List Id Column ID
func ValidateGroupIdSiteIdListIdColumnID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdListIdColumnID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id List Id Column ID
func (id GroupIdSiteIdListIdColumnId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s/columns/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId, id.ColumnDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id List Id Column ID
func (id GroupIdSiteIdListIdColumnId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listId"),
		resourceids.StaticSegment("columns", "columns", "columns"),
		resourceids.UserSpecifiedSegment("columnDefinitionId", "columnDefinitionId"),
	}
}

// String returns a human-readable description of this Group Id Site Id List Id Column ID
func (id GroupIdSiteIdListIdColumnId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
		fmt.Sprintf("Column Definition: %q", id.ColumnDefinitionId),
	}
	return fmt.Sprintf("Group Id Site Id List Id Column (%s)", strings.Join(components, "\n"))
}
