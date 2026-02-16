package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdListIdItemIdActivityId{}

// GroupIdSiteIdListIdItemIdActivityId is a struct representing the Resource ID for a Group Id Site Id List Id Item Id Activity
type GroupIdSiteIdListIdItemIdActivityId struct {
	GroupId           string
	SiteId            string
	ListId            string
	ListItemId        string
	ItemActivityOLDId string
}

// NewGroupIdSiteIdListIdItemIdActivityID returns a new GroupIdSiteIdListIdItemIdActivityId struct
func NewGroupIdSiteIdListIdItemIdActivityID(groupId string, siteId string, listId string, listItemId string, itemActivityOLDId string) GroupIdSiteIdListIdItemIdActivityId {
	return GroupIdSiteIdListIdItemIdActivityId{
		GroupId:           groupId,
		SiteId:            siteId,
		ListId:            listId,
		ListItemId:        listItemId,
		ItemActivityOLDId: itemActivityOLDId,
	}
}

// ParseGroupIdSiteIdListIdItemIdActivityID parses 'input' into a GroupIdSiteIdListIdItemIdActivityId
func ParseGroupIdSiteIdListIdItemIdActivityID(input string) (*GroupIdSiteIdListIdItemIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdItemIdActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdItemIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdListIdItemIdActivityIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdListIdItemIdActivityId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdListIdItemIdActivityIDInsensitively(input string) (*GroupIdSiteIdListIdItemIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdItemIdActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdItemIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdListIdItemIdActivityId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ItemActivityOLDId, ok = input.Parsed["itemActivityOLDId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityOLDId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdListIdItemIdActivityID checks that 'input' can be parsed as a Group Id Site Id List Id Item Id Activity ID
func ValidateGroupIdSiteIdListIdItemIdActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdListIdItemIdActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id List Id Item Id Activity ID
func (id GroupIdSiteIdListIdItemIdActivityId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s/items/%s/activities/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId, id.ListItemId, id.ItemActivityOLDId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id List Id Item Id Activity ID
func (id GroupIdSiteIdListIdItemIdActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("listItemId", "listItemId"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityOLDId", "itemActivityOLDId"),
	}
}

// String returns a human-readable description of this Group Id Site Id List Id Item Id Activity ID
func (id GroupIdSiteIdListIdItemIdActivityId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
		fmt.Sprintf("List Item: %q", id.ListItemId),
		fmt.Sprintf("Item Activity OLD: %q", id.ItemActivityOLDId),
	}
	return fmt.Sprintf("Group Id Site Id List Id Item Id Activity (%s)", strings.Join(components, "\n"))
}
