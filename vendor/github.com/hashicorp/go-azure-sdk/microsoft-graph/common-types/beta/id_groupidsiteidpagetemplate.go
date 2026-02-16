package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdPageTemplateId{}

// GroupIdSiteIdPageTemplateId is a struct representing the Resource ID for a Group Id Site Id Page Template
type GroupIdSiteIdPageTemplateId struct {
	GroupId        string
	SiteId         string
	PageTemplateId string
}

// NewGroupIdSiteIdPageTemplateID returns a new GroupIdSiteIdPageTemplateId struct
func NewGroupIdSiteIdPageTemplateID(groupId string, siteId string, pageTemplateId string) GroupIdSiteIdPageTemplateId {
	return GroupIdSiteIdPageTemplateId{
		GroupId:        groupId,
		SiteId:         siteId,
		PageTemplateId: pageTemplateId,
	}
}

// ParseGroupIdSiteIdPageTemplateID parses 'input' into a GroupIdSiteIdPageTemplateId
func ParseGroupIdSiteIdPageTemplateID(input string) (*GroupIdSiteIdPageTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdPageTemplateIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdPageTemplateId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdPageTemplateIDInsensitively(input string) (*GroupIdSiteIdPageTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageTemplateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdPageTemplateId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateGroupIdSiteIdPageTemplateID checks that 'input' can be parsed as a Group Id Site Id Page Template ID
func ValidateGroupIdSiteIdPageTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdPageTemplateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Page Template ID
func (id GroupIdSiteIdPageTemplateId) ID() string {
	fmtString := "/groups/%s/sites/%s/pageTemplates/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.PageTemplateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Page Template ID
func (id GroupIdSiteIdPageTemplateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("pageTemplates", "pageTemplates", "pageTemplates"),
		resourceids.UserSpecifiedSegment("pageTemplateId", "pageTemplateId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Page Template ID
func (id GroupIdSiteIdPageTemplateId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Page Template: %q", id.PageTemplateId),
	}
	return fmt.Sprintf("Group Id Site Id Page Template (%s)", strings.Join(components, "\n"))
}
