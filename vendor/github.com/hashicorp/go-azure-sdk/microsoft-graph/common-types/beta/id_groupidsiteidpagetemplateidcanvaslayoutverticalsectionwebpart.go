package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartId{}

// GroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartId is a struct representing the Resource ID for a Group Id Site Id Page Template Id Canvas Layout Vertical Section Webpart
type GroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartId struct {
	GroupId        string
	SiteId         string
	PageTemplateId string
	WebPartId      string
}

// NewGroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartID returns a new GroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartId struct
func NewGroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartID(groupId string, siteId string, pageTemplateId string, webPartId string) GroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartId {
	return GroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartId{
		GroupId:        groupId,
		SiteId:         siteId,
		PageTemplateId: pageTemplateId,
		WebPartId:      webPartId,
	}
}

// ParseGroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartID parses 'input' into a GroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartId
func ParseGroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartID(input string) (*GroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartIDInsensitively(input string) (*GroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.WebPartId, ok = input.Parsed["webPartId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "webPartId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartID checks that 'input' can be parsed as a Group Id Site Id Page Template Id Canvas Layout Vertical Section Webpart ID
func ValidateGroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Page Template Id Canvas Layout Vertical Section Webpart ID
func (id GroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartId) ID() string {
	fmtString := "/groups/%s/sites/%s/pageTemplates/%s/canvasLayout/verticalSection/webparts/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.PageTemplateId, id.WebPartId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Page Template Id Canvas Layout Vertical Section Webpart ID
func (id GroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("pageTemplates", "pageTemplates", "pageTemplates"),
		resourceids.UserSpecifiedSegment("pageTemplateId", "pageTemplateId"),
		resourceids.StaticSegment("canvasLayout", "canvasLayout", "canvasLayout"),
		resourceids.StaticSegment("verticalSection", "verticalSection", "verticalSection"),
		resourceids.StaticSegment("webparts", "webparts", "webparts"),
		resourceids.UserSpecifiedSegment("webPartId", "webPartId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Page Template Id Canvas Layout Vertical Section Webpart ID
func (id GroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Page Template: %q", id.PageTemplateId),
		fmt.Sprintf("Web Part: %q", id.WebPartId),
	}
	return fmt.Sprintf("Group Id Site Id Page Template Id Canvas Layout Vertical Section Webpart (%s)", strings.Join(components, "\n"))
}
