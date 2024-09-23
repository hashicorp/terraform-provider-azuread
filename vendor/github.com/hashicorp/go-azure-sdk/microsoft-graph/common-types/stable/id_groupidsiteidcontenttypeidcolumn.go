package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdContentTypeIdColumnId{}

// GroupIdSiteIdContentTypeIdColumnId is a struct representing the Resource ID for a Group Id Site Id Content Type Id Column
type GroupIdSiteIdContentTypeIdColumnId struct {
	GroupId            string
	SiteId             string
	ContentTypeId      string
	ColumnDefinitionId string
}

// NewGroupIdSiteIdContentTypeIdColumnID returns a new GroupIdSiteIdContentTypeIdColumnId struct
func NewGroupIdSiteIdContentTypeIdColumnID(groupId string, siteId string, contentTypeId string, columnDefinitionId string) GroupIdSiteIdContentTypeIdColumnId {
	return GroupIdSiteIdContentTypeIdColumnId{
		GroupId:            groupId,
		SiteId:             siteId,
		ContentTypeId:      contentTypeId,
		ColumnDefinitionId: columnDefinitionId,
	}
}

// ParseGroupIdSiteIdContentTypeIdColumnID parses 'input' into a GroupIdSiteIdContentTypeIdColumnId
func ParseGroupIdSiteIdContentTypeIdColumnID(input string) (*GroupIdSiteIdContentTypeIdColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdContentTypeIdColumnId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdContentTypeIdColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdContentTypeIdColumnIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdContentTypeIdColumnId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdContentTypeIdColumnIDInsensitively(input string) (*GroupIdSiteIdContentTypeIdColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdContentTypeIdColumnId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdContentTypeIdColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdContentTypeIdColumnId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.ContentTypeId, ok = input.Parsed["contentTypeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId", input)
	}

	if id.ColumnDefinitionId, ok = input.Parsed["columnDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "columnDefinitionId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdContentTypeIdColumnID checks that 'input' can be parsed as a Group Id Site Id Content Type Id Column ID
func ValidateGroupIdSiteIdContentTypeIdColumnID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdContentTypeIdColumnID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Content Type Id Column ID
func (id GroupIdSiteIdContentTypeIdColumnId) ID() string {
	fmtString := "/groups/%s/sites/%s/contentTypes/%s/columns/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ContentTypeId, id.ColumnDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Content Type Id Column ID
func (id GroupIdSiteIdContentTypeIdColumnId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("contentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeId"),
		resourceids.StaticSegment("columns", "columns", "columns"),
		resourceids.UserSpecifiedSegment("columnDefinitionId", "columnDefinitionId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Content Type Id Column ID
func (id GroupIdSiteIdContentTypeIdColumnId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
		fmt.Sprintf("Column Definition: %q", id.ColumnDefinitionId),
	}
	return fmt.Sprintf("Group Id Site Id Content Type Id Column (%s)", strings.Join(components, "\n"))
}
