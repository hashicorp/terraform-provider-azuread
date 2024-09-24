package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdListId{}

// GroupIdSiteIdListId is a struct representing the Resource ID for a Group Id Site Id List
type GroupIdSiteIdListId struct {
	GroupId string
	SiteId  string
	ListId  string
}

// NewGroupIdSiteIdListID returns a new GroupIdSiteIdListId struct
func NewGroupIdSiteIdListID(groupId string, siteId string, listId string) GroupIdSiteIdListId {
	return GroupIdSiteIdListId{
		GroupId: groupId,
		SiteId:  siteId,
		ListId:  listId,
	}
}

// ParseGroupIdSiteIdListID parses 'input' into a GroupIdSiteIdListId
func ParseGroupIdSiteIdListID(input string) (*GroupIdSiteIdListId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdListIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdListId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdListIDInsensitively(input string) (*GroupIdSiteIdListId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdListId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdListId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdListId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateGroupIdSiteIdListID checks that 'input' can be parsed as a Group Id Site Id List ID
func ValidateGroupIdSiteIdListID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdListID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id List ID
func (id GroupIdSiteIdListId) ID() string {
	fmtString := "/groups/%s/sites/%s/lists/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ListId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id List ID
func (id GroupIdSiteIdListId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("lists", "lists", "lists"),
		resourceids.UserSpecifiedSegment("listId", "listId"),
	}
}

// String returns a human-readable description of this Group Id Site Id List ID
func (id GroupIdSiteIdListId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("List: %q", id.ListId),
	}
	return fmt.Sprintf("Group Id Site Id List (%s)", strings.Join(components, "\n"))
}
