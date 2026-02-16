package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdRecycleBinItemId{}

// GroupIdSiteIdRecycleBinItemId is a struct representing the Resource ID for a Group Id Site Id Recycle Bin Item
type GroupIdSiteIdRecycleBinItemId struct {
	GroupId          string
	SiteId           string
	RecycleBinItemId string
}

// NewGroupIdSiteIdRecycleBinItemID returns a new GroupIdSiteIdRecycleBinItemId struct
func NewGroupIdSiteIdRecycleBinItemID(groupId string, siteId string, recycleBinItemId string) GroupIdSiteIdRecycleBinItemId {
	return GroupIdSiteIdRecycleBinItemId{
		GroupId:          groupId,
		SiteId:           siteId,
		RecycleBinItemId: recycleBinItemId,
	}
}

// ParseGroupIdSiteIdRecycleBinItemID parses 'input' into a GroupIdSiteIdRecycleBinItemId
func ParseGroupIdSiteIdRecycleBinItemID(input string) (*GroupIdSiteIdRecycleBinItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdRecycleBinItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdRecycleBinItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdRecycleBinItemIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdRecycleBinItemId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdRecycleBinItemIDInsensitively(input string) (*GroupIdSiteIdRecycleBinItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdRecycleBinItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdRecycleBinItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdRecycleBinItemId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.RecycleBinItemId, ok = input.Parsed["recycleBinItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "recycleBinItemId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdRecycleBinItemID checks that 'input' can be parsed as a Group Id Site Id Recycle Bin Item ID
func ValidateGroupIdSiteIdRecycleBinItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdRecycleBinItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Recycle Bin Item ID
func (id GroupIdSiteIdRecycleBinItemId) ID() string {
	fmtString := "/groups/%s/sites/%s/recycleBin/items/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.RecycleBinItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Recycle Bin Item ID
func (id GroupIdSiteIdRecycleBinItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("recycleBin", "recycleBin", "recycleBin"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("recycleBinItemId", "recycleBinItemId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Recycle Bin Item ID
func (id GroupIdSiteIdRecycleBinItemId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Recycle Bin Item: %q", id.RecycleBinItemId),
	}
	return fmt.Sprintf("Group Id Site Id Recycle Bin Item (%s)", strings.Join(components, "\n"))
}
