package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnId{}

// GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnId is a struct representing the Resource ID for a Group Id Site Id Page Template Id Canvas Layout Horizontal Section Id Column
type GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnId struct {
	GroupId                   string
	SiteId                    string
	PageTemplateId            string
	HorizontalSectionId       string
	HorizontalSectionColumnId string
}

// NewGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnID returns a new GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnId struct
func NewGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnID(groupId string, siteId string, pageTemplateId string, horizontalSectionId string, horizontalSectionColumnId string) GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnId {
	return GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnId{
		GroupId:                   groupId,
		SiteId:                    siteId,
		PageTemplateId:            pageTemplateId,
		HorizontalSectionId:       horizontalSectionId,
		HorizontalSectionColumnId: horizontalSectionColumnId,
	}
}

// ParseGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnID parses 'input' into a GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnId
func ParseGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnID(input string) (*GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIDInsensitively(input string) (*GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnID checks that 'input' can be parsed as a Group Id Site Id Page Template Id Canvas Layout Horizontal Section Id Column ID
func ValidateGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Page Template Id Canvas Layout Horizontal Section Id Column ID
func (id GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnId) ID() string {
	fmtString := "/groups/%s/sites/%s/pageTemplates/%s/canvasLayout/horizontalSections/%s/columns/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.PageTemplateId, id.HorizontalSectionId, id.HorizontalSectionColumnId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Page Template Id Canvas Layout Horizontal Section Id Column ID
func (id GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this Group Id Site Id Page Template Id Canvas Layout Horizontal Section Id Column ID
func (id GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Page Template: %q", id.PageTemplateId),
		fmt.Sprintf("Horizontal Section: %q", id.HorizontalSectionId),
		fmt.Sprintf("Horizontal Section Column: %q", id.HorizontalSectionColumnId),
	}
	return fmt.Sprintf("Group Id Site Id Page Template Id Canvas Layout Horizontal Section Id Column (%s)", strings.Join(components, "\n"))
}
