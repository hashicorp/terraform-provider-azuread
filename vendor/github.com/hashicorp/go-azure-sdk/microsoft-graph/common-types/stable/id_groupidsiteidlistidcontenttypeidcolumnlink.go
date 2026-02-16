package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdListIdContentTypeIdColumnLinkId{}

// GroupIdSiteIdListIdContentTypeIdColumnLinkId is a struct representing the Resource ID for a Group Id Site Id List Id Content Type Id Column Link
type GroupIdSiteIdListIdContentTypeIdColumnLinkId struct {
	GroupId       string
	SiteId        string
	ListId        string
	ContentTypeId string
	ColumnLinkId  string
}

// NewGroupIdSiteIdListIdContentTypeIdColumnLinkID returns a new GroupIdSiteIdListIdContentTypeIdColumnLinkId struct
func NewGroupIdSiteIdListIdContentTypeIdColumnLinkID(groupId string, siteId string, listId string, contentTypeId string, columnLinkId string) GroupIdSiteIdListIdContentTypeIdColumnLinkId {
	return GroupIdSiteIdListIdContentTypeIdColumnLinkId{
		GroupId:       groupId,
		SiteId:        siteId,
		ListId:        listId,
		ContentTypeId: contentTypeId,
		ColumnLinkId:  columnLinkId,
	}
}

// ParseGroupIdSiteIdListIdContentTypeIdColumnLinkID parses 'input' into a GroupIdSiteIdListIdContentTypeIdColumnLinkId
func ParseGroupIdSiteIdListIdContentTypeIdColumnLinkID(input string) (*GroupIdSiteIdListIdContentTypeIdColumnLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdContentTypeIdColumnLinkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdContentTypeIdColumnLinkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdListIdContentTypeIdColumnLinkIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdListIdContentTypeIdColumnLinkId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdListIdContentTypeIdColumnLinkIDInsensitively(input string) (*GroupIdSiteIdListIdContentTypeIdColumnLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdContentTypeIdColumnLinkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdContentTypeIdColumnLinkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdListIdContentTypeIdColumnLinkId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ContentTypeId, ok = input.Parsed["contentTypeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId", input)
	}

	if id.ColumnLinkId, ok = input.Parsed["columnLinkId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "columnLinkId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdListIdContentTypeIdColumnLinkID checks that 'input' can be parsed as a Group Id Site Id List Id Content Type Id Column Link ID
func ValidateGroupIdSiteIdListIdContentTypeIdColumnLinkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdListIdContentTypeIdColumnLinkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id List Id Content Type Id Column Link ID
func (id GroupIdSiteIdListIdContentTypeIdColumnLinkId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s/contentTypes/%s/columnLinks/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId, id.ContentTypeId, id.ColumnLinkId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id List Id Content Type Id Column Link ID
func (id GroupIdSiteIdListIdContentTypeIdColumnLinkId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listId"),
		resourceids.StaticSegment("contentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeId"),
		resourceids.StaticSegment("columnLinks", "columnLinks", "columnLinks"),
		resourceids.UserSpecifiedSegment("columnLinkId", "columnLinkId"),
	}
}

// String returns a human-readable description of this Group Id Site Id List Id Content Type Id Column Link ID
func (id GroupIdSiteIdListIdContentTypeIdColumnLinkId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
		fmt.Sprintf("Column Link: %q", id.ColumnLinkId),
	}
	return fmt.Sprintf("Group Id Site Id List Id Content Type Id Column Link (%s)", strings.Join(components, "\n"))
}
