package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdListIdContentTypeId{}

// GroupIdSiteIdListIdContentTypeId is a struct representing the Resource ID for a Group Id Site Id List Id Content Type
type GroupIdSiteIdListIdContentTypeId struct {
	GroupId       string
	SiteId        string
	ListId        string
	ContentTypeId string
}

// NewGroupIdSiteIdListIdContentTypeID returns a new GroupIdSiteIdListIdContentTypeId struct
func NewGroupIdSiteIdListIdContentTypeID(groupId string, siteId string, listId string, contentTypeId string) GroupIdSiteIdListIdContentTypeId {
	return GroupIdSiteIdListIdContentTypeId{
		GroupId:       groupId,
		SiteId:        siteId,
		ListId:        listId,
		ContentTypeId: contentTypeId,
	}
}

// ParseGroupIdSiteIdListIdContentTypeID parses 'input' into a GroupIdSiteIdListIdContentTypeId
func ParseGroupIdSiteIdListIdContentTypeID(input string) (*GroupIdSiteIdListIdContentTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdContentTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdContentTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdListIdContentTypeIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdListIdContentTypeId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdListIdContentTypeIDInsensitively(input string) (*GroupIdSiteIdListIdContentTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdContentTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdContentTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdListIdContentTypeId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateGroupIdSiteIdListIdContentTypeID checks that 'input' can be parsed as a Group Id Site Id List Id Content Type ID
func ValidateGroupIdSiteIdListIdContentTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdListIdContentTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id List Id Content Type ID
func (id GroupIdSiteIdListIdContentTypeId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s/contentTypes/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId, id.ContentTypeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id List Id Content Type ID
func (id GroupIdSiteIdListIdContentTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listId"),
		resourceids.StaticSegment("contentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeId"),
	}
}

// String returns a human-readable description of this Group Id Site Id List Id Content Type ID
func (id GroupIdSiteIdListIdContentTypeId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
	}
	return fmt.Sprintf("Group Id Site Id List Id Content Type (%s)", strings.Join(components, "\n"))
}
