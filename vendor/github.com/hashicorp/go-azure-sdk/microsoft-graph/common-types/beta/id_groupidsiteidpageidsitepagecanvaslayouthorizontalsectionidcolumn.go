package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId{}

// GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId is a struct representing the Resource ID for a Group Id Site Id Page Id Site Page Canvas Layout Horizontal Section Id Column
type GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId struct {
	GroupId                   string
	SiteId                    string
	BaseSitePageId            string
	HorizontalSectionId       string
	HorizontalSectionColumnId string
}

// NewGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnID returns a new GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId struct
func NewGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnID(groupId string, siteId string, baseSitePageId string, horizontalSectionId string, horizontalSectionColumnId string) GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId {
	return GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId{
		GroupId:                   groupId,
		SiteId:                    siteId,
		BaseSitePageId:            baseSitePageId,
		HorizontalSectionId:       horizontalSectionId,
		HorizontalSectionColumnId: horizontalSectionColumnId,
	}
}

// ParseGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnID parses 'input' into a GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId
func ParseGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnID(input string) (*GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIDInsensitively(input string) (*GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.HorizontalSectionColumnId, ok = input.Parsed["horizontalSectionColumnId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "horizontalSectionColumnId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnID checks that 'input' can be parsed as a Group Id Site Id Page Id Site Page Canvas Layout Horizontal Section Id Column ID
func ValidateGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Page Id Site Page Canvas Layout Horizontal Section Id Column ID
func (id GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId) ID() string {
	fmtString := "/groups/%s/sites/%s/pages/%s/sitePage/canvasLayout/horizontalSections/%s/columns/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.BaseSitePageId, id.HorizontalSectionId, id.HorizontalSectionColumnId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Page Id Site Page Canvas Layout Horizontal Section Id Column ID
func (id GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("columns", "columns", "columns"),
		resourceids.UserSpecifiedSegment("horizontalSectionColumnId", "horizontalSectionColumnId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Page Id Site Page Canvas Layout Horizontal Section Id Column ID
func (id GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Base Site Page: %q", id.BaseSitePageId),
		fmt.Sprintf("Horizontal Section: %q", id.HorizontalSectionId),
		fmt.Sprintf("Horizontal Section Column: %q", id.HorizontalSectionColumnId),
	}
	return fmt.Sprintf("Group Id Site Id Page Id Site Page Canvas Layout Horizontal Section Id Column (%s)", strings.Join(components, "\n"))
}
