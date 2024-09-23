package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdListIdItemIdVersionId{}

// GroupIdSiteIdListIdItemIdVersionId is a struct representing the Resource ID for a Group Id Site Id List Id Item Id Version
type GroupIdSiteIdListIdItemIdVersionId struct {
	GroupId           string
	SiteId            string
	ListId            string
	ListItemId        string
	ListItemVersionId string
}

// NewGroupIdSiteIdListIdItemIdVersionID returns a new GroupIdSiteIdListIdItemIdVersionId struct
func NewGroupIdSiteIdListIdItemIdVersionID(groupId string, siteId string, listId string, listItemId string, listItemVersionId string) GroupIdSiteIdListIdItemIdVersionId {
	return GroupIdSiteIdListIdItemIdVersionId{
		GroupId:           groupId,
		SiteId:            siteId,
		ListId:            listId,
		ListItemId:        listItemId,
		ListItemVersionId: listItemVersionId,
	}
}

// ParseGroupIdSiteIdListIdItemIdVersionID parses 'input' into a GroupIdSiteIdListIdItemIdVersionId
func ParseGroupIdSiteIdListIdItemIdVersionID(input string) (*GroupIdSiteIdListIdItemIdVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdItemIdVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdItemIdVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdListIdItemIdVersionIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdListIdItemIdVersionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdListIdItemIdVersionIDInsensitively(input string) (*GroupIdSiteIdListIdItemIdVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdItemIdVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdItemIdVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdListIdItemIdVersionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ListItemId, ok = input.Parsed["listItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "listItemId", input)
	}

	if id.ListItemVersionId, ok = input.Parsed["listItemVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "listItemVersionId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdListIdItemIdVersionID checks that 'input' can be parsed as a Group Id Site Id List Id Item Id Version ID
func ValidateGroupIdSiteIdListIdItemIdVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdListIdItemIdVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id List Id Item Id Version ID
func (id GroupIdSiteIdListIdItemIdVersionId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s/items/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId, id.ListItemId, id.ListItemVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id List Id Item Id Version ID
func (id GroupIdSiteIdListIdItemIdVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("listItemId", "listItemId"),
		resourceids.StaticSegment("versions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("listItemVersionId", "listItemVersionId"),
	}
}

// String returns a human-readable description of this Group Id Site Id List Id Item Id Version ID
func (id GroupIdSiteIdListIdItemIdVersionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
		fmt.Sprintf("List Item: %q", id.ListItemId),
		fmt.Sprintf("List Item Version: %q", id.ListItemVersionId),
	}
	return fmt.Sprintf("Group Id Site Id List Id Item Id Version (%s)", strings.Join(components, "\n"))
}
