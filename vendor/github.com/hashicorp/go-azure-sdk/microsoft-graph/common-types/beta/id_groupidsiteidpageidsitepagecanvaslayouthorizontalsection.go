package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId{}

// GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId is a struct representing the Resource ID for a Group Id Site Id Page Id Site Page Canvas Layout Horizontal Section
type GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId struct {
	GroupId             string
	SiteId              string
	BaseSitePageId      string
	HorizontalSectionId string
}

// NewGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionID returns a new GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId struct
func NewGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionID(groupId string, siteId string, baseSitePageId string, horizontalSectionId string) GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId {
	return GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId{
		GroupId:             groupId,
		SiteId:              siteId,
		BaseSitePageId:      baseSitePageId,
		HorizontalSectionId: horizontalSectionId,
	}
}

// ParseGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionID parses 'input' into a GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId
func ParseGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionID(input string) (*GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIDInsensitively(input string) (*GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.HorizontalSectionId, ok = input.Parsed["horizontalSectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "horizontalSectionId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionID checks that 'input' can be parsed as a Group Id Site Id Page Id Site Page Canvas Layout Horizontal Section ID
func ValidateGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Page Id Site Page Canvas Layout Horizontal Section ID
func (id GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId) ID() string {
	fmtString := "/groups/%s/sites/%s/pages/%s/sitePage/canvasLayout/horizontalSections/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.BaseSitePageId, id.HorizontalSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Page Id Site Page Canvas Layout Horizontal Section ID
func (id GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("pages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("baseSitePageId", "baseSitePageId"),
		resourceids.StaticSegment("sitePage", "sitePage", "sitePage"),
		resourceids.StaticSegment("canvasLayout", "canvasLayout", "canvasLayout"),
		resourceids.StaticSegment("horizontalSections", "horizontalSections", "horizontalSections"),
		resourceids.UserSpecifiedSegment("horizontalSectionId", "horizontalSectionId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Page Id Site Page Canvas Layout Horizontal Section ID
func (id GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Base Site Page: %q", id.BaseSitePageId),
		fmt.Sprintf("Horizontal Section: %q", id.HorizontalSectionId),
	}
	return fmt.Sprintf("Group Id Site Id Page Id Site Page Canvas Layout Horizontal Section (%s)", strings.Join(components, "\n"))
}
