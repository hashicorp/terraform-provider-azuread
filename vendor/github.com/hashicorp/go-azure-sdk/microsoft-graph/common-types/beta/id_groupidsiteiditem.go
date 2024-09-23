package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdItemId{}

// GroupIdSiteIdItemId is a struct representing the Resource ID for a Group Id Site Id Item
type GroupIdSiteIdItemId struct {
	GroupId    string
	SiteId     string
	BaseItemId string
}

// NewGroupIdSiteIdItemID returns a new GroupIdSiteIdItemId struct
func NewGroupIdSiteIdItemID(groupId string, siteId string, baseItemId string) GroupIdSiteIdItemId {
	return GroupIdSiteIdItemId{
		GroupId:    groupId,
		SiteId:     siteId,
		BaseItemId: baseItemId,
	}
}

// ParseGroupIdSiteIdItemID parses 'input' into a GroupIdSiteIdItemId
func ParseGroupIdSiteIdItemID(input string) (*GroupIdSiteIdItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdItemIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdItemId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdItemIDInsensitively(input string) (*GroupIdSiteIdItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdItemId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.BaseItemId, ok = input.Parsed["baseItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseItemId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdItemID checks that 'input' can be parsed as a Group Id Site Id Item ID
func ValidateGroupIdSiteIdItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Item ID
func (id GroupIdSiteIdItemId) ID() string {
	fmtString := "/groups/%s/sites/%s/items/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.BaseItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Item ID
func (id GroupIdSiteIdItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("baseItemId", "baseItemId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Item ID
func (id GroupIdSiteIdItemId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Base Item: %q", id.BaseItemId),
	}
	return fmt.Sprintf("Group Id Site Id Item (%s)", strings.Join(components, "\n"))
}
