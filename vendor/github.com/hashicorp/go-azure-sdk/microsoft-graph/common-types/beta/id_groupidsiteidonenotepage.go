package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdOnenotePageId{}

// GroupIdSiteIdOnenotePageId is a struct representing the Resource ID for a Group Id Site Id Onenote Page
type GroupIdSiteIdOnenotePageId struct {
	GroupId       string
	SiteId        string
	OnenotePageId string
}

// NewGroupIdSiteIdOnenotePageID returns a new GroupIdSiteIdOnenotePageId struct
func NewGroupIdSiteIdOnenotePageID(groupId string, siteId string, onenotePageId string) GroupIdSiteIdOnenotePageId {
	return GroupIdSiteIdOnenotePageId{
		GroupId:       groupId,
		SiteId:        siteId,
		OnenotePageId: onenotePageId,
	}
}

// ParseGroupIdSiteIdOnenotePageID parses 'input' into a GroupIdSiteIdOnenotePageId
func ParseGroupIdSiteIdOnenotePageID(input string) (*GroupIdSiteIdOnenotePageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenotePageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenotePageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdOnenotePageIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdOnenotePageId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdOnenotePageIDInsensitively(input string) (*GroupIdSiteIdOnenotePageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenotePageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenotePageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdOnenotePageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.OnenotePageId, ok = input.Parsed["onenotePageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdOnenotePageID checks that 'input' can be parsed as a Group Id Site Id Onenote Page ID
func ValidateGroupIdSiteIdOnenotePageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdOnenotePageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Onenote Page ID
func (id GroupIdSiteIdOnenotePageId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/pages/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Onenote Page ID
func (id GroupIdSiteIdOnenotePageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("pages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Onenote Page ID
func (id GroupIdSiteIdOnenotePageId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Group Id Site Id Onenote Page (%s)", strings.Join(components, "\n"))
}
