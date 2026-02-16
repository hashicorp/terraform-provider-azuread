package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartId{}

// GroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartId is a struct representing the Resource ID for a Group Id Site Id Page Id Site Page Canvas Layout Vertical Section Webpart
type GroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartId struct {
	GroupId        string
	SiteId         string
	BaseSitePageId string
	WebPartId      string
}

// NewGroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartID returns a new GroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartId struct
func NewGroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartID(groupId string, siteId string, baseSitePageId string, webPartId string) GroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartId {
	return GroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartId{
		GroupId:        groupId,
		SiteId:         siteId,
		BaseSitePageId: baseSitePageId,
		WebPartId:      webPartId,
	}
}

// ParseGroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartID parses 'input' into a GroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartId
func ParseGroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartID(input string) (*GroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartIDInsensitively(input string) (*GroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.WebPartId, ok = input.Parsed["webPartId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "webPartId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartID checks that 'input' can be parsed as a Group Id Site Id Page Id Site Page Canvas Layout Vertical Section Webpart ID
func ValidateGroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Page Id Site Page Canvas Layout Vertical Section Webpart ID
func (id GroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartId) ID() string {
	fmtString := "/groups/%s/sites/%s/pages/%s/sitePage/canvasLayout/verticalSection/webparts/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.BaseSitePageId, id.WebPartId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Page Id Site Page Canvas Layout Vertical Section Webpart ID
func (id GroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("pages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("baseSitePageId", "baseSitePageId"),
		resourceids.StaticSegment("sitePage", "sitePage", "sitePage"),
		resourceids.StaticSegment("canvasLayout", "canvasLayout", "canvasLayout"),
		resourceids.StaticSegment("verticalSection", "verticalSection", "verticalSection"),
		resourceids.StaticSegment("webparts", "webparts", "webparts"),
		resourceids.UserSpecifiedSegment("webPartId", "webPartId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Page Id Site Page Canvas Layout Vertical Section Webpart ID
func (id GroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Base Site Page: %q", id.BaseSitePageId),
		fmt.Sprintf("Web Part: %q", id.WebPartId),
	}
	return fmt.Sprintf("Group Id Site Id Page Id Site Page Canvas Layout Vertical Section Webpart (%s)", strings.Join(components, "\n"))
}
