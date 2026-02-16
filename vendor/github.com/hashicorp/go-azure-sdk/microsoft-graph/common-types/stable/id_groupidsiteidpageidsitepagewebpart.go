package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdPageIdSitePageWebPartId{}

// GroupIdSiteIdPageIdSitePageWebPartId is a struct representing the Resource ID for a Group Id Site Id Page Id Site Page Web Part
type GroupIdSiteIdPageIdSitePageWebPartId struct {
	GroupId        string
	SiteId         string
	BaseSitePageId string
	WebPartId      string
}

// NewGroupIdSiteIdPageIdSitePageWebPartID returns a new GroupIdSiteIdPageIdSitePageWebPartId struct
func NewGroupIdSiteIdPageIdSitePageWebPartID(groupId string, siteId string, baseSitePageId string, webPartId string) GroupIdSiteIdPageIdSitePageWebPartId {
	return GroupIdSiteIdPageIdSitePageWebPartId{
		GroupId:        groupId,
		SiteId:         siteId,
		BaseSitePageId: baseSitePageId,
		WebPartId:      webPartId,
	}
}

// ParseGroupIdSiteIdPageIdSitePageWebPartID parses 'input' into a GroupIdSiteIdPageIdSitePageWebPartId
func ParseGroupIdSiteIdPageIdSitePageWebPartID(input string) (*GroupIdSiteIdPageIdSitePageWebPartId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageIdSitePageWebPartId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageIdSitePageWebPartId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdPageIdSitePageWebPartIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdPageIdSitePageWebPartId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdPageIdSitePageWebPartIDInsensitively(input string) (*GroupIdSiteIdPageIdSitePageWebPartId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdPageIdSitePageWebPartId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdPageIdSitePageWebPartId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdPageIdSitePageWebPartId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdSiteIdPageIdSitePageWebPartID checks that 'input' can be parsed as a Group Id Site Id Page Id Site Page Web Part ID
func ValidateGroupIdSiteIdPageIdSitePageWebPartID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdPageIdSitePageWebPartID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Page Id Site Page Web Part ID
func (id GroupIdSiteIdPageIdSitePageWebPartId) ID() string {
	fmtString := "/groups/%s/sites/%s/pages/%s/sitePage/webParts/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.BaseSitePageId, id.WebPartId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Page Id Site Page Web Part ID
func (id GroupIdSiteIdPageIdSitePageWebPartId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("pages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("baseSitePageId", "baseSitePageId"),
		resourceids.StaticSegment("sitePage", "sitePage", "sitePage"),
		resourceids.StaticSegment("webParts", "webParts", "webParts"),
		resourceids.UserSpecifiedSegment("webPartId", "webPartId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Page Id Site Page Web Part ID
func (id GroupIdSiteIdPageIdSitePageWebPartId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Base Site Page: %q", id.BaseSitePageId),
		fmt.Sprintf("Web Part: %q", id.WebPartId),
	}
	return fmt.Sprintf("Group Id Site Id Page Id Site Page Web Part (%s)", strings.Join(components, "\n"))
}
