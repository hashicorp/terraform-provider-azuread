package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdListIdContentTypeIdBaseTypeId{}

// GroupIdSiteIdListIdContentTypeIdBaseTypeId is a struct representing the Resource ID for a Group Id Site Id List Id Content Type Id Base Type
type GroupIdSiteIdListIdContentTypeIdBaseTypeId struct {
	GroupId        string
	SiteId         string
	ListId         string
	ContentTypeId  string
	ContentTypeId1 string
}

// NewGroupIdSiteIdListIdContentTypeIdBaseTypeID returns a new GroupIdSiteIdListIdContentTypeIdBaseTypeId struct
func NewGroupIdSiteIdListIdContentTypeIdBaseTypeID(groupId string, siteId string, listId string, contentTypeId string, contentTypeId1 string) GroupIdSiteIdListIdContentTypeIdBaseTypeId {
	return GroupIdSiteIdListIdContentTypeIdBaseTypeId{
		GroupId:        groupId,
		SiteId:         siteId,
		ListId:         listId,
		ContentTypeId:  contentTypeId,
		ContentTypeId1: contentTypeId1,
	}
}

// ParseGroupIdSiteIdListIdContentTypeIdBaseTypeID parses 'input' into a GroupIdSiteIdListIdContentTypeIdBaseTypeId
func ParseGroupIdSiteIdListIdContentTypeIdBaseTypeID(input string) (*GroupIdSiteIdListIdContentTypeIdBaseTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdContentTypeIdBaseTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdContentTypeIdBaseTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdListIdContentTypeIdBaseTypeIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdListIdContentTypeIdBaseTypeId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdListIdContentTypeIdBaseTypeIDInsensitively(input string) (*GroupIdSiteIdListIdContentTypeIdBaseTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdContentTypeIdBaseTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdContentTypeIdBaseTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdListIdContentTypeIdBaseTypeId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ContentTypeId1, ok = input.Parsed["contentTypeId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId1", input)
	}

	return nil
}

// ValidateGroupIdSiteIdListIdContentTypeIdBaseTypeID checks that 'input' can be parsed as a Group Id Site Id List Id Content Type Id Base Type ID
func ValidateGroupIdSiteIdListIdContentTypeIdBaseTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdListIdContentTypeIdBaseTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id List Id Content Type Id Base Type ID
func (id GroupIdSiteIdListIdContentTypeIdBaseTypeId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s/contentTypes/%s/baseTypes/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId, id.ContentTypeId, id.ContentTypeId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id List Id Content Type Id Base Type ID
func (id GroupIdSiteIdListIdContentTypeIdBaseTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listId"),
		resourceids.StaticSegment("contentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeId"),
		resourceids.StaticSegment("baseTypes", "baseTypes", "baseTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId1", "contentTypeId1"),
	}
}

// String returns a human-readable description of this Group Id Site Id List Id Content Type Id Base Type ID
func (id GroupIdSiteIdListIdContentTypeIdBaseTypeId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
		fmt.Sprintf("Content Type Id 1: %q", id.ContentTypeId1),
	}
	return fmt.Sprintf("Group Id Site Id List Id Content Type Id Base Type (%s)", strings.Join(components, "\n"))
}
