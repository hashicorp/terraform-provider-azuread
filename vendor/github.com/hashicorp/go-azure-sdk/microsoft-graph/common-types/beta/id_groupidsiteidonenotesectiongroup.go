package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdOnenoteSectionGroupId{}

// GroupIdSiteIdOnenoteSectionGroupId is a struct representing the Resource ID for a Group Id Site Id Onenote Section Group
type GroupIdSiteIdOnenoteSectionGroupId struct {
	GroupId        string
	SiteId         string
	SectionGroupId string
}

// NewGroupIdSiteIdOnenoteSectionGroupID returns a new GroupIdSiteIdOnenoteSectionGroupId struct
func NewGroupIdSiteIdOnenoteSectionGroupID(groupId string, siteId string, sectionGroupId string) GroupIdSiteIdOnenoteSectionGroupId {
	return GroupIdSiteIdOnenoteSectionGroupId{
		GroupId:        groupId,
		SiteId:         siteId,
		SectionGroupId: sectionGroupId,
	}
}

// ParseGroupIdSiteIdOnenoteSectionGroupID parses 'input' into a GroupIdSiteIdOnenoteSectionGroupId
func ParseGroupIdSiteIdOnenoteSectionGroupID(input string) (*GroupIdSiteIdOnenoteSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenoteSectionGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenoteSectionGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdOnenoteSectionGroupIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdOnenoteSectionGroupId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdOnenoteSectionGroupIDInsensitively(input string) (*GroupIdSiteIdOnenoteSectionGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenoteSectionGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenoteSectionGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdOnenoteSectionGroupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.SectionGroupId, ok = input.Parsed["sectionGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdOnenoteSectionGroupID checks that 'input' can be parsed as a Group Id Site Id Onenote Section Group ID
func ValidateGroupIdSiteIdOnenoteSectionGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdOnenoteSectionGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Onenote Section Group ID
func (id GroupIdSiteIdOnenoteSectionGroupId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/sectionGroups/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SectionGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Onenote Section Group ID
func (id GroupIdSiteIdOnenoteSectionGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("sectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Onenote Section Group ID
func (id GroupIdSiteIdOnenoteSectionGroupId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
	}
	return fmt.Sprintf("Group Id Site Id Onenote Section Group (%s)", strings.Join(components, "\n"))
}
