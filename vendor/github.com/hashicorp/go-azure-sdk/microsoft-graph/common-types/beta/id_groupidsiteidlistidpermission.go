package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdListIdPermissionId{}

// GroupIdSiteIdListIdPermissionId is a struct representing the Resource ID for a Group Id Site Id List Id Permission
type GroupIdSiteIdListIdPermissionId struct {
	GroupId      string
	SiteId       string
	ListId       string
	PermissionId string
}

// NewGroupIdSiteIdListIdPermissionID returns a new GroupIdSiteIdListIdPermissionId struct
func NewGroupIdSiteIdListIdPermissionID(groupId string, siteId string, listId string, permissionId string) GroupIdSiteIdListIdPermissionId {
	return GroupIdSiteIdListIdPermissionId{
		GroupId:      groupId,
		SiteId:       siteId,
		ListId:       listId,
		PermissionId: permissionId,
	}
}

// ParseGroupIdSiteIdListIdPermissionID parses 'input' into a GroupIdSiteIdListIdPermissionId
func ParseGroupIdSiteIdListIdPermissionID(input string) (*GroupIdSiteIdListIdPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdListIdPermissionIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdListIdPermissionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdListIdPermissionIDInsensitively(input string) (*GroupIdSiteIdListIdPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListIdPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListIdPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdListIdPermissionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.PermissionId, ok = input.Parsed["permissionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "permissionId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdListIdPermissionID checks that 'input' can be parsed as a Group Id Site Id List Id Permission ID
func ValidateGroupIdSiteIdListIdPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdListIdPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id List Id Permission ID
func (id GroupIdSiteIdListIdPermissionId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s/permissions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId, id.PermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id List Id Permission ID
func (id GroupIdSiteIdListIdPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listId"),
		resourceids.StaticSegment("permissions", "permissions", "permissions"),
		resourceids.UserSpecifiedSegment("permissionId", "permissionId"),
	}
}

// String returns a human-readable description of this Group Id Site Id List Id Permission ID
func (id GroupIdSiteIdListIdPermissionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
		fmt.Sprintf("Permission: %q", id.PermissionId),
	}
	return fmt.Sprintf("Group Id Site Id List Id Permission (%s)", strings.Join(components, "\n"))
}
