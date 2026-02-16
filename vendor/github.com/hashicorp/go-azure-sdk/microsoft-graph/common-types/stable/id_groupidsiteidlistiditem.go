package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdListIdItemId{}

// GroupIdSiteIdListIdItemId is a struct representing the Resource ID for a Group Id Site Id List Id Item
type GroupIdSiteIdListIdItemId struct {
	GroupId    string
	SiteId     string
	ListId     string
	ListItemId string
}

// NewGroupIdSiteIdListIdItemID returns a new GroupIdSiteIdListIdItemId struct
func NewGroupIdSiteIdListIdItemID(groupId string, siteId string, listId string, listItemId string) GroupIdSiteIdListIdItemId {
	return GroupIdSiteIdListIdItemId{
		GroupId:    groupId,
		SiteId:     siteId,
		ListId:     listId,
		ListItemId: listItemId,
	}
}

// ParseGroupIdSiteIdListIdItemID parses 'input' into a GroupIdSiteIdListIdItemId
func ParseGroupIdSiteIdListIdItemID(input string) (*GroupIdSiteIdListIdItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdListIdItemIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdListIdItemId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdListIdItemIDInsensitively(input string) (*GroupIdSiteIdListIdItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdListIdItemId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateGroupIdSiteIdListIdItemID checks that 'input' can be parsed as a Group Id Site Id List Id Item ID
func ValidateGroupIdSiteIdListIdItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdListIdItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id List Id Item ID
func (id GroupIdSiteIdListIdItemId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s/items/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId, id.ListItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id List Id Item ID
func (id GroupIdSiteIdListIdItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("listItemId", "listItemId"),
	}
}

// String returns a human-readable description of this Group Id Site Id List Id Item ID
func (id GroupIdSiteIdListIdItemId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
		fmt.Sprintf("List Item: %q", id.ListItemId),
	}
	return fmt.Sprintf("Group Id Site Id List Id Item (%s)", strings.Join(components, "\n"))
}
