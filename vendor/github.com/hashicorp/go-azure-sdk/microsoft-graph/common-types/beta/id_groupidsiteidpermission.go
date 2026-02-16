package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdPermissionId{}

// GroupIdSiteIdPermissionId is a struct representing the Resource ID for a Group Id Site Id Permission
type GroupIdSiteIdPermissionId struct {
	GroupId      string
	SiteId       string
	PermissionId string
}

// NewGroupIdSiteIdPermissionID returns a new GroupIdSiteIdPermissionId struct
func NewGroupIdSiteIdPermissionID(groupId string, siteId string, permissionId string) GroupIdSiteIdPermissionId {
	return GroupIdSiteIdPermissionId{
		GroupId:      groupId,
		SiteId:       siteId,
		PermissionId: permissionId,
	}
}

// ParseGroupIdSiteIdPermissionID parses 'input' into a GroupIdSiteIdPermissionId
func ParseGroupIdSiteIdPermissionID(input string) (*GroupIdSiteIdPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdPermissionIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdPermissionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdPermissionIDInsensitively(input string) (*GroupIdSiteIdPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdPermissionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.PermissionId, ok = input.Parsed["permissionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "permissionId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdPermissionID checks that 'input' can be parsed as a Group Id Site Id Permission ID
func ValidateGroupIdSiteIdPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Permission ID
func (id GroupIdSiteIdPermissionId) ID() string {
	fmtString := "/groups/%s/sites/%s/permissions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.PermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Permission ID
func (id GroupIdSiteIdPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("permissions", "permissions", "permissions"),
		resourceids.UserSpecifiedSegment("permissionId", "permissionId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Permission ID
func (id GroupIdSiteIdPermissionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Permission: %q", id.PermissionId),
	}
	return fmt.Sprintf("Group Id Site Id Permission (%s)", strings.Join(components, "\n"))
}
