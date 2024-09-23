package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdContentTypeIdColumnPositionId{}

// GroupIdSiteIdContentTypeIdColumnPositionId is a struct representing the Resource ID for a Group Id Site Id Content Type Id Column Position
type GroupIdSiteIdContentTypeIdColumnPositionId struct {
	GroupId            string
	SiteId             string
	ContentTypeId      string
	ColumnDefinitionId string
}

// NewGroupIdSiteIdContentTypeIdColumnPositionID returns a new GroupIdSiteIdContentTypeIdColumnPositionId struct
func NewGroupIdSiteIdContentTypeIdColumnPositionID(groupId string, siteId string, contentTypeId string, columnDefinitionId string) GroupIdSiteIdContentTypeIdColumnPositionId {
	return GroupIdSiteIdContentTypeIdColumnPositionId{
		GroupId:            groupId,
		SiteId:             siteId,
		ContentTypeId:      contentTypeId,
		ColumnDefinitionId: columnDefinitionId,
	}
}

// ParseGroupIdSiteIdContentTypeIdColumnPositionID parses 'input' into a GroupIdSiteIdContentTypeIdColumnPositionId
func ParseGroupIdSiteIdContentTypeIdColumnPositionID(input string) (*GroupIdSiteIdContentTypeIdColumnPositionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdContentTypeIdColumnPositionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdContentTypeIdColumnPositionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdContentTypeIdColumnPositionIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdContentTypeIdColumnPositionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdContentTypeIdColumnPositionIDInsensitively(input string) (*GroupIdSiteIdContentTypeIdColumnPositionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdContentTypeIdColumnPositionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdContentTypeIdColumnPositionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdContentTypeIdColumnPositionId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdSiteIdContentTypeIdColumnPositionID checks that 'input' can be parsed as a Group Id Site Id Content Type Id Column Position ID
func ValidateGroupIdSiteIdContentTypeIdColumnPositionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdContentTypeIdColumnPositionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Content Type Id Column Position ID
func (id GroupIdSiteIdContentTypeIdColumnPositionId) ID() string {
	fmtString := "/groups/%s/sites/%s/contentTypes/%s/columnPositions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ContentTypeId, id.ColumnDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Content Type Id Column Position ID
func (id GroupIdSiteIdContentTypeIdColumnPositionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("contentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeId"),
		resourceids.StaticSegment("columnPositions", "columnPositions", "columnPositions"),
		resourceids.UserSpecifiedSegment("columnDefinitionId", "columnDefinitionId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Content Type Id Column Position ID
func (id GroupIdSiteIdContentTypeIdColumnPositionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
		fmt.Sprintf("Column Definition: %q", id.ColumnDefinitionId),
	}
	return fmt.Sprintf("Group Id Site Id Content Type Id Column Position (%s)", strings.Join(components, "\n"))
}
