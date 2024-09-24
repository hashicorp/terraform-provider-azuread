package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdListIdItemIdPermissionId{}

// GroupIdSiteIdListIdItemIdPermissionId is a struct representing the Resource ID for a Group Id Site Id List Id Item Id Permission
type GroupIdSiteIdListIdItemIdPermissionId struct {
	GroupId      string
	SiteId       string
	ListId       string
	ListItemId   string
	PermissionId string
}

// NewGroupIdSiteIdListIdItemIdPermissionID returns a new GroupIdSiteIdListIdItemIdPermissionId struct
func NewGroupIdSiteIdListIdItemIdPermissionID(groupId string, siteId string, listId string, listItemId string, permissionId string) GroupIdSiteIdListIdItemIdPermissionId {
	return GroupIdSiteIdListIdItemIdPermissionId{
		GroupId:      groupId,
		SiteId:       siteId,
		ListId:       listId,
		ListItemId:   listItemId,
		PermissionId: permissionId,
	}
}

// ParseGroupIdSiteIdListIdItemIdPermissionID parses 'input' into a GroupIdSiteIdListIdItemIdPermissionId
func ParseGroupIdSiteIdListIdItemIdPermissionID(input string) (*GroupIdSiteIdListIdItemIdPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdItemIdPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdItemIdPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdListIdItemIdPermissionIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdListIdItemIdPermissionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdListIdItemIdPermissionIDInsensitively(input string) (*GroupIdSiteIdListIdItemIdPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdItemIdPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdItemIdPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdListIdItemIdPermissionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.PermissionId, ok = input.Parsed["permissionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "permissionId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdListIdItemIdPermissionID checks that 'input' can be parsed as a Group Id Site Id List Id Item Id Permission ID
func ValidateGroupIdSiteIdListIdItemIdPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdListIdItemIdPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id List Id Item Id Permission ID
func (id GroupIdSiteIdListIdItemIdPermissionId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s/items/%s/permissions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId, id.ListItemId, id.PermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id List Id Item Id Permission ID
func (id GroupIdSiteIdListIdItemIdPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("listItemId", "listItemId"),
		resourceids.StaticSegment("permissions", "permissions", "permissions"),
		resourceids.UserSpecifiedSegment("permissionId", "permissionId"),
	}
}

// String returns a human-readable description of this Group Id Site Id List Id Item Id Permission ID
func (id GroupIdSiteIdListIdItemIdPermissionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
		fmt.Sprintf("List Item: %q", id.ListItemId),
		fmt.Sprintf("Permission: %q", id.PermissionId),
	}
	return fmt.Sprintf("Group Id Site Id List Id Item Id Permission (%s)", strings.Join(components, "\n"))
}
