package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdContentTypeIdColumnLinkId{}

// GroupIdSiteIdContentTypeIdColumnLinkId is a struct representing the Resource ID for a Group Id Site Id Content Type Id Column Link
type GroupIdSiteIdContentTypeIdColumnLinkId struct {
	GroupId       string
	SiteId        string
	ContentTypeId string
	ColumnLinkId  string
}

// NewGroupIdSiteIdContentTypeIdColumnLinkID returns a new GroupIdSiteIdContentTypeIdColumnLinkId struct
func NewGroupIdSiteIdContentTypeIdColumnLinkID(groupId string, siteId string, contentTypeId string, columnLinkId string) GroupIdSiteIdContentTypeIdColumnLinkId {
	return GroupIdSiteIdContentTypeIdColumnLinkId{
		GroupId:       groupId,
		SiteId:        siteId,
		ContentTypeId: contentTypeId,
		ColumnLinkId:  columnLinkId,
	}
}

// ParseGroupIdSiteIdContentTypeIdColumnLinkID parses 'input' into a GroupIdSiteIdContentTypeIdColumnLinkId
func ParseGroupIdSiteIdContentTypeIdColumnLinkID(input string) (*GroupIdSiteIdContentTypeIdColumnLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdContentTypeIdColumnLinkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdContentTypeIdColumnLinkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdContentTypeIdColumnLinkIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdContentTypeIdColumnLinkId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdContentTypeIdColumnLinkIDInsensitively(input string) (*GroupIdSiteIdContentTypeIdColumnLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdContentTypeIdColumnLinkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdContentTypeIdColumnLinkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdContentTypeIdColumnLinkId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ColumnLinkId, ok = input.Parsed["columnLinkId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "columnLinkId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdContentTypeIdColumnLinkID checks that 'input' can be parsed as a Group Id Site Id Content Type Id Column Link ID
func ValidateGroupIdSiteIdContentTypeIdColumnLinkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdContentTypeIdColumnLinkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Content Type Id Column Link ID
func (id GroupIdSiteIdContentTypeIdColumnLinkId) ID() string {
	fmtString := "/groups/%s/sites/%s/contentTypes/%s/columnLinks/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ContentTypeId, id.ColumnLinkId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Content Type Id Column Link ID
func (id GroupIdSiteIdContentTypeIdColumnLinkId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("contentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeId"),
		resourceids.StaticSegment("columnLinks", "columnLinks", "columnLinks"),
		resourceids.UserSpecifiedSegment("columnLinkId", "columnLinkId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Content Type Id Column Link ID
func (id GroupIdSiteIdContentTypeIdColumnLinkId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
		fmt.Sprintf("Column Link: %q", id.ColumnLinkId),
	}
	return fmt.Sprintf("Group Id Site Id Content Type Id Column Link (%s)", strings.Join(components, "\n"))
}
