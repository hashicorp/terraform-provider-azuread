package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdPageId{}

// GroupIdSiteIdPageId is a struct representing the Resource ID for a Group Id Site Id Page
type GroupIdSiteIdPageId struct {
	GroupId        string
	SiteId         string
	BaseSitePageId string
}

// NewGroupIdSiteIdPageID returns a new GroupIdSiteIdPageId struct
func NewGroupIdSiteIdPageID(groupId string, siteId string, baseSitePageId string) GroupIdSiteIdPageId {
	return GroupIdSiteIdPageId{
		GroupId:        groupId,
		SiteId:         siteId,
		BaseSitePageId: baseSitePageId,
	}
}

// ParseGroupIdSiteIdPageID parses 'input' into a GroupIdSiteIdPageId
func ParseGroupIdSiteIdPageID(input string) (*GroupIdSiteIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdPageIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdPageId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdPageIDInsensitively(input string) (*GroupIdSiteIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdPageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.BaseSitePageId, ok = input.Parsed["baseSitePageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseSitePageId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdPageID checks that 'input' can be parsed as a Group Id Site Id Page ID
func ValidateGroupIdSiteIdPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Page ID
func (id GroupIdSiteIdPageId) ID() string {
	fmtString := "/groups/%s/sites/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.BaseSitePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Page ID
func (id GroupIdSiteIdPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("pages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("baseSitePageId", "baseSitePageId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Page ID
func (id GroupIdSiteIdPageId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Base Site Page: %q", id.BaseSitePageId),
	}
	return fmt.Sprintf("Group Id Site Id Page (%s)", strings.Join(components, "\n"))
}
