package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdPageTemplateIdWebPartId{}

// GroupIdSiteIdPageTemplateIdWebPartId is a struct representing the Resource ID for a Group Id Site Id Page Template Id Web Part
type GroupIdSiteIdPageTemplateIdWebPartId struct {
	GroupId        string
	SiteId         string
	PageTemplateId string
	WebPartId      string
}

// NewGroupIdSiteIdPageTemplateIdWebPartID returns a new GroupIdSiteIdPageTemplateIdWebPartId struct
func NewGroupIdSiteIdPageTemplateIdWebPartID(groupId string, siteId string, pageTemplateId string, webPartId string) GroupIdSiteIdPageTemplateIdWebPartId {
	return GroupIdSiteIdPageTemplateIdWebPartId{
		GroupId:        groupId,
		SiteId:         siteId,
		PageTemplateId: pageTemplateId,
		WebPartId:      webPartId,
	}
}

// ParseGroupIdSiteIdPageTemplateIdWebPartID parses 'input' into a GroupIdSiteIdPageTemplateIdWebPartId
func ParseGroupIdSiteIdPageTemplateIdWebPartID(input string) (*GroupIdSiteIdPageTemplateIdWebPartId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageTemplateIdWebPartId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageTemplateIdWebPartId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdPageTemplateIdWebPartIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdPageTemplateIdWebPartId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdPageTemplateIdWebPartIDInsensitively(input string) (*GroupIdSiteIdPageTemplateIdWebPartId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageTemplateIdWebPartId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageTemplateIdWebPartId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdPageTemplateIdWebPartId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdSiteIdPageTemplateIdWebPartID checks that 'input' can be parsed as a Group Id Site Id Page Template Id Web Part ID
func ValidateGroupIdSiteIdPageTemplateIdWebPartID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdPageTemplateIdWebPartID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Page Template Id Web Part ID
func (id GroupIdSiteIdPageTemplateIdWebPartId) ID() string {
	fmtString := "/groups/%s/sites/%s/pageTemplates/%s/webParts/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.PageTemplateId, id.WebPartId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Page Template Id Web Part ID
func (id GroupIdSiteIdPageTemplateIdWebPartId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("pageTemplates", "pageTemplates", "pageTemplates"),
		resourceids.UserSpecifiedSegment("pageTemplateId", "pageTemplateId"),
		resourceids.StaticSegment("webParts", "webParts", "webParts"),
		resourceids.UserSpecifiedSegment("webPartId", "webPartId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Page Template Id Web Part ID
func (id GroupIdSiteIdPageTemplateIdWebPartId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Page Template: %q", id.PageTemplateId),
		fmt.Sprintf("Web Part: %q", id.WebPartId),
	}
	return fmt.Sprintf("Group Id Site Id Page Template Id Web Part (%s)", strings.Join(components, "\n"))
}
