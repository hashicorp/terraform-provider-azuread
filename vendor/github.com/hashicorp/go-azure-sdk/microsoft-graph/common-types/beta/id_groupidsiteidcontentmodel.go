package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdContentModelId{}

// GroupIdSiteIdContentModelId is a struct representing the Resource ID for a Group Id Site Id Content Model
type GroupIdSiteIdContentModelId struct {
	GroupId        string
	SiteId         string
	ContentModelId string
}

// NewGroupIdSiteIdContentModelID returns a new GroupIdSiteIdContentModelId struct
func NewGroupIdSiteIdContentModelID(groupId string, siteId string, contentModelId string) GroupIdSiteIdContentModelId {
	return GroupIdSiteIdContentModelId{
		GroupId:        groupId,
		SiteId:         siteId,
		ContentModelId: contentModelId,
	}
}

// ParseGroupIdSiteIdContentModelID parses 'input' into a GroupIdSiteIdContentModelId
func ParseGroupIdSiteIdContentModelID(input string) (*GroupIdSiteIdContentModelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdContentModelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdContentModelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdContentModelIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdContentModelId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdContentModelIDInsensitively(input string) (*GroupIdSiteIdContentModelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdContentModelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdContentModelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdContentModelId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.ContentModelId, ok = input.Parsed["contentModelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contentModelId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdContentModelID checks that 'input' can be parsed as a Group Id Site Id Content Model ID
func ValidateGroupIdSiteIdContentModelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdContentModelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Content Model ID
func (id GroupIdSiteIdContentModelId) ID() string {
	fmtString := "/groups/%s/sites/%s/contentModels/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ContentModelId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Content Model ID
func (id GroupIdSiteIdContentModelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("contentModels", "contentModels", "contentModels"),
		resourceids.UserSpecifiedSegment("contentModelId", "contentModelId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Content Model ID
func (id GroupIdSiteIdContentModelId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Content Model: %q", id.ContentModelId),
	}
	return fmt.Sprintf("Group Id Site Id Content Model (%s)", strings.Join(components, "\n"))
}
