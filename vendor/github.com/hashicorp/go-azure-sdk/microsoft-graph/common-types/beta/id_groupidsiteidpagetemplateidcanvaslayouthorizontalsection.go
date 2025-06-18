package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionId{}

// GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionId is a struct representing the Resource ID for a Group Id Site Id Page Template Id Canvas Layout Horizontal Section
type GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionId struct {
	GroupId             string
	SiteId              string
	PageTemplateId      string
	HorizontalSectionId string
}

// NewGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionID returns a new GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionId struct
func NewGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionID(groupId string, siteId string, pageTemplateId string, horizontalSectionId string) GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionId {
	return GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionId{
		GroupId:             groupId,
		SiteId:              siteId,
		PageTemplateId:      pageTemplateId,
		HorizontalSectionId: horizontalSectionId,
	}
}

// ParseGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionID parses 'input' into a GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionId
func ParseGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionID(input string) (*GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIDInsensitively(input string) (*GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionID checks that 'input' can be parsed as a Group Id Site Id Page Template Id Canvas Layout Horizontal Section ID
func ValidateGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Page Template Id Canvas Layout Horizontal Section ID
func (id GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionId) ID() string {
	fmtString := "/groups/%s/sites/%s/pageTemplates/%s/canvasLayout/horizontalSections/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.PageTemplateId, id.HorizontalSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Page Template Id Canvas Layout Horizontal Section ID
func (id GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this Group Id Site Id Page Template Id Canvas Layout Horizontal Section ID
func (id GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Page Template: %q", id.PageTemplateId),
		fmt.Sprintf("Horizontal Section: %q", id.HorizontalSectionId),
	}
	return fmt.Sprintf("Group Id Site Id Page Template Id Canvas Layout Horizontal Section (%s)", strings.Join(components, "\n"))
}
