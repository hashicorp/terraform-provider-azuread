package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdContentTypeId{}

// GroupIdSiteIdContentTypeId is a struct representing the Resource ID for a Group Id Site Id Content Type
type GroupIdSiteIdContentTypeId struct {
	GroupId       string
	SiteId        string
	ContentTypeId string
}

// NewGroupIdSiteIdContentTypeID returns a new GroupIdSiteIdContentTypeId struct
func NewGroupIdSiteIdContentTypeID(groupId string, siteId string, contentTypeId string) GroupIdSiteIdContentTypeId {
	return GroupIdSiteIdContentTypeId{
		GroupId:       groupId,
		SiteId:        siteId,
		ContentTypeId: contentTypeId,
	}
}

// ParseGroupIdSiteIdContentTypeID parses 'input' into a GroupIdSiteIdContentTypeId
func ParseGroupIdSiteIdContentTypeID(input string) (*GroupIdSiteIdContentTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdContentTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdContentTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdContentTypeIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdContentTypeId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdContentTypeIDInsensitively(input string) (*GroupIdSiteIdContentTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdContentTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdContentTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdContentTypeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.ContentTypeId, ok = input.Parsed["contentTypeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdContentTypeID checks that 'input' can be parsed as a Group Id Site Id Content Type ID
func ValidateGroupIdSiteIdContentTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdContentTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Content Type ID
func (id GroupIdSiteIdContentTypeId) ID() string {
	fmtString := "/groups/%s/sites/%s/contentTypes/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ContentTypeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Content Type ID
func (id GroupIdSiteIdContentTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("contentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Content Type ID
func (id GroupIdSiteIdContentTypeId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
	}
	return fmt.Sprintf("Group Id Site Id Content Type (%s)", strings.Join(components, "\n"))
}
