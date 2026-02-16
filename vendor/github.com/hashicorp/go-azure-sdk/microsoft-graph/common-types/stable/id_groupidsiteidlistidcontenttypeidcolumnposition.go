package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdListIdContentTypeIdColumnPositionId{}

// GroupIdSiteIdListIdContentTypeIdColumnPositionId is a struct representing the Resource ID for a Group Id Site Id List Id Content Type Id Column Position
type GroupIdSiteIdListIdContentTypeIdColumnPositionId struct {
	GroupId            string
	SiteId             string
	ListId             string
	ContentTypeId      string
	ColumnDefinitionId string
}

// NewGroupIdSiteIdListIdContentTypeIdColumnPositionID returns a new GroupIdSiteIdListIdContentTypeIdColumnPositionId struct
func NewGroupIdSiteIdListIdContentTypeIdColumnPositionID(groupId string, siteId string, listId string, contentTypeId string, columnDefinitionId string) GroupIdSiteIdListIdContentTypeIdColumnPositionId {
	return GroupIdSiteIdListIdContentTypeIdColumnPositionId{
		GroupId:            groupId,
		SiteId:             siteId,
		ListId:             listId,
		ContentTypeId:      contentTypeId,
		ColumnDefinitionId: columnDefinitionId,
	}
}

// ParseGroupIdSiteIdListIdContentTypeIdColumnPositionID parses 'input' into a GroupIdSiteIdListIdContentTypeIdColumnPositionId
func ParseGroupIdSiteIdListIdContentTypeIdColumnPositionID(input string) (*GroupIdSiteIdListIdContentTypeIdColumnPositionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdContentTypeIdColumnPositionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdContentTypeIdColumnPositionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdListIdContentTypeIdColumnPositionIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdListIdContentTypeIdColumnPositionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdListIdContentTypeIdColumnPositionIDInsensitively(input string) (*GroupIdSiteIdListIdContentTypeIdColumnPositionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdContentTypeIdColumnPositionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdContentTypeIdColumnPositionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdListIdContentTypeIdColumnPositionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ColumnDefinitionId, ok = input.Parsed["columnDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "columnDefinitionId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdListIdContentTypeIdColumnPositionID checks that 'input' can be parsed as a Group Id Site Id List Id Content Type Id Column Position ID
func ValidateGroupIdSiteIdListIdContentTypeIdColumnPositionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdListIdContentTypeIdColumnPositionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id List Id Content Type Id Column Position ID
func (id GroupIdSiteIdListIdContentTypeIdColumnPositionId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s/contentTypes/%s/columnPositions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId, id.ContentTypeId, id.ColumnDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id List Id Content Type Id Column Position ID
func (id GroupIdSiteIdListIdContentTypeIdColumnPositionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listId"),
		resourceids.StaticSegment("contentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeId"),
		resourceids.StaticSegment("columnPositions", "columnPositions", "columnPositions"),
		resourceids.UserSpecifiedSegment("columnDefinitionId", "columnDefinitionId"),
	}
}

// String returns a human-readable description of this Group Id Site Id List Id Content Type Id Column Position ID
func (id GroupIdSiteIdListIdContentTypeIdColumnPositionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
		fmt.Sprintf("Column Definition: %q", id.ColumnDefinitionId),
	}
	return fmt.Sprintf("Group Id Site Id List Id Content Type Id Column Position (%s)", strings.Join(components, "\n"))
}
