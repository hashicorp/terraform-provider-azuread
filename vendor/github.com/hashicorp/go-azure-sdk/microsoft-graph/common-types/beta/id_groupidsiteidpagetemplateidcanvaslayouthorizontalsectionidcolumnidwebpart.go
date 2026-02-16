package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId{}

// GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId is a struct representing the Resource ID for a Group Id Site Id Page Template Id Canvas Layout Horizontal Section Id Column Id Webpart
type GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId struct {
	GroupId                   string
	SiteId                    string
	PageTemplateId            string
	HorizontalSectionId       string
	HorizontalSectionColumnId string
	WebPartId                 string
}

// NewGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartID returns a new GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId struct
func NewGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartID(groupId string, siteId string, pageTemplateId string, horizontalSectionId string, horizontalSectionColumnId string, webPartId string) GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId {
	return GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId{
		GroupId:                   groupId,
		SiteId:                    siteId,
		PageTemplateId:            pageTemplateId,
		HorizontalSectionId:       horizontalSectionId,
		HorizontalSectionColumnId: horizontalSectionColumnId,
		WebPartId:                 webPartId,
	}
}

// ParseGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartID parses 'input' into a GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId
func ParseGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartID(input string) (*GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartIDInsensitively(input string) (*GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.PageTemplateId, ok = input.Parsed["pageTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "pageTemplateId", input)
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

// ValidateGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartID checks that 'input' can be parsed as a Group Id Site Id Page Template Id Canvas Layout Horizontal Section Id Column Id Webpart ID
func ValidateGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Page Template Id Canvas Layout Horizontal Section Id Column Id Webpart ID
func (id GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId) ID() string {
	fmtString := "/groups/%s/sites/%s/pageTemplates/%s/canvasLayout/horizontalSections/%s/columns/%s/webparts/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.PageTemplateId, id.HorizontalSectionId, id.HorizontalSectionColumnId, id.WebPartId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Page Template Id Canvas Layout Horizontal Section Id Column Id Webpart ID
func (id GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("pageTemplates", "pageTemplates", "pageTemplates"),
		resourceids.UserSpecifiedSegment("pageTemplateId", "pageTemplateId"),
		resourceids.StaticSegment("canvasLayout", "canvasLayout", "canvasLayout"),
		resourceids.StaticSegment("horizontalSections", "horizontalSections", "horizontalSections"),
		resourceids.UserSpecifiedSegment("horizontalSectionId", "horizontalSectionId"),
		resourceids.StaticSegment("columns", "columns", "columns"),
		resourceids.UserSpecifiedSegment("horizontalSectionColumnId", "horizontalSectionColumnId"),
		resourceids.StaticSegment("webparts", "webparts", "webparts"),
		resourceids.UserSpecifiedSegment("webPartId", "webPartId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Page Template Id Canvas Layout Horizontal Section Id Column Id Webpart ID
func (id GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Page Template: %q", id.PageTemplateId),
		fmt.Sprintf("Horizontal Section: %q", id.HorizontalSectionId),
		fmt.Sprintf("Horizontal Section Column: %q", id.HorizontalSectionColumnId),
		fmt.Sprintf("Web Part: %q", id.WebPartId),
	}
	return fmt.Sprintf("Group Id Site Id Page Template Id Canvas Layout Horizontal Section Id Column Id Webpart (%s)", strings.Join(components, "\n"))
}
