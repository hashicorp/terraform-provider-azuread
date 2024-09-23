package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdOnenoteSectionId{}

// GroupIdSiteIdOnenoteSectionId is a struct representing the Resource ID for a Group Id Site Id Onenote Section
type GroupIdSiteIdOnenoteSectionId struct {
	GroupId          string
	SiteId           string
	OnenoteSectionId string
}

// NewGroupIdSiteIdOnenoteSectionID returns a new GroupIdSiteIdOnenoteSectionId struct
func NewGroupIdSiteIdOnenoteSectionID(groupId string, siteId string, onenoteSectionId string) GroupIdSiteIdOnenoteSectionId {
	return GroupIdSiteIdOnenoteSectionId{
		GroupId:          groupId,
		SiteId:           siteId,
		OnenoteSectionId: onenoteSectionId,
	}
}

// ParseGroupIdSiteIdOnenoteSectionID parses 'input' into a GroupIdSiteIdOnenoteSectionId
func ParseGroupIdSiteIdOnenoteSectionID(input string) (*GroupIdSiteIdOnenoteSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenoteSectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenoteSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdOnenoteSectionIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdOnenoteSectionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdOnenoteSectionIDInsensitively(input string) (*GroupIdSiteIdOnenoteSectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenoteSectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenoteSectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdOnenoteSectionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.OnenoteSectionId, ok = input.Parsed["onenoteSectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdOnenoteSectionID checks that 'input' can be parsed as a Group Id Site Id Onenote Section ID
func ValidateGroupIdSiteIdOnenoteSectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdOnenoteSectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Onenote Section ID
func (id GroupIdSiteIdOnenoteSectionId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/sections/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.OnenoteSectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Onenote Section ID
func (id GroupIdSiteIdOnenoteSectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Onenote Section ID
func (id GroupIdSiteIdOnenoteSectionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
	}
	return fmt.Sprintf("Group Id Site Id Onenote Section (%s)", strings.Join(components, "\n"))
}
