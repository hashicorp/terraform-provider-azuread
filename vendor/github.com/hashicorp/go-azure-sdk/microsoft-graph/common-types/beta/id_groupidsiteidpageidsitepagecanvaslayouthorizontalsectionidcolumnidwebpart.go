package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartId{}

// GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartId is a struct representing the Resource ID for a Group Id Site Id Page Id Site Page Canvas Layout Horizontal Section Id Column Id Webpart
type GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartId struct {
	GroupId                   string
	SiteId                    string
	BaseSitePageId            string
	HorizontalSectionId       string
	HorizontalSectionColumnId string
	WebPartId                 string
}

// NewGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartID returns a new GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartId struct
func NewGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartID(groupId string, siteId string, baseSitePageId string, horizontalSectionId string, horizontalSectionColumnId string, webPartId string) GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartId {
	return GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartId{
		GroupId:                   groupId,
		SiteId:                    siteId,
		BaseSitePageId:            baseSitePageId,
		HorizontalSectionId:       horizontalSectionId,
		HorizontalSectionColumnId: horizontalSectionColumnId,
		WebPartId:                 webPartId,
	}
}

// ParseGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartID parses 'input' into a GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartId
func ParseGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartID(input string) (*GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartIDInsensitively(input string) (*GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.WebPartId, ok = input.Parsed["webPartId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "webPartId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartID checks that 'input' can be parsed as a Group Id Site Id Page Id Site Page Canvas Layout Horizontal Section Id Column Id Webpart ID
func ValidateGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Page Id Site Page Canvas Layout Horizontal Section Id Column Id Webpart ID
func (id GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartId) ID() string {
	fmtString := "/groups/%s/sites/%s/pages/%s/sitePage/canvasLayout/horizontalSections/%s/columns/%s/webparts/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.BaseSitePageId, id.HorizontalSectionId, id.HorizontalSectionColumnId, id.WebPartId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Page Id Site Page Canvas Layout Horizontal Section Id Column Id Webpart ID
func (id GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("webparts", "webparts", "webparts"),
		resourceids.UserSpecifiedSegment("webPartId", "webPartId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Page Id Site Page Canvas Layout Horizontal Section Id Column Id Webpart ID
func (id GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Base Site Page: %q", id.BaseSitePageId),
		fmt.Sprintf("Horizontal Section: %q", id.HorizontalSectionId),
		fmt.Sprintf("Horizontal Section Column: %q", id.HorizontalSectionColumnId),
		fmt.Sprintf("Web Part: %q", id.WebPartId),
	}
	return fmt.Sprintf("Group Id Site Id Page Id Site Page Canvas Layout Horizontal Section Id Column Id Webpart (%s)", strings.Join(components, "\n"))
}
