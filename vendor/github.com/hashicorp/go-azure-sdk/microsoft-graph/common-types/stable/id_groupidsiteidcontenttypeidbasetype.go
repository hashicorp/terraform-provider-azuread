package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdContentTypeIdBaseTypeId{}

// GroupIdSiteIdContentTypeIdBaseTypeId is a struct representing the Resource ID for a Group Id Site Id Content Type Id Base Type
type GroupIdSiteIdContentTypeIdBaseTypeId struct {
	GroupId        string
	SiteId         string
	ContentTypeId  string
	ContentTypeId1 string
}

// NewGroupIdSiteIdContentTypeIdBaseTypeID returns a new GroupIdSiteIdContentTypeIdBaseTypeId struct
func NewGroupIdSiteIdContentTypeIdBaseTypeID(groupId string, siteId string, contentTypeId string, contentTypeId1 string) GroupIdSiteIdContentTypeIdBaseTypeId {
	return GroupIdSiteIdContentTypeIdBaseTypeId{
		GroupId:        groupId,
		SiteId:         siteId,
		ContentTypeId:  contentTypeId,
		ContentTypeId1: contentTypeId1,
	}
}

// ParseGroupIdSiteIdContentTypeIdBaseTypeID parses 'input' into a GroupIdSiteIdContentTypeIdBaseTypeId
func ParseGroupIdSiteIdContentTypeIdBaseTypeID(input string) (*GroupIdSiteIdContentTypeIdBaseTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdContentTypeIdBaseTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdContentTypeIdBaseTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdContentTypeIdBaseTypeIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdContentTypeIdBaseTypeId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdContentTypeIdBaseTypeIDInsensitively(input string) (*GroupIdSiteIdContentTypeIdBaseTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdContentTypeIdBaseTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdContentTypeIdBaseTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdContentTypeIdBaseTypeId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ContentTypeId1, ok = input.Parsed["contentTypeId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId1", input)
	}

	return nil
}

// ValidateGroupIdSiteIdContentTypeIdBaseTypeID checks that 'input' can be parsed as a Group Id Site Id Content Type Id Base Type ID
func ValidateGroupIdSiteIdContentTypeIdBaseTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdContentTypeIdBaseTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Content Type Id Base Type ID
func (id GroupIdSiteIdContentTypeIdBaseTypeId) ID() string {
	fmtString := "/groups/%s/sites/%s/contentTypes/%s/baseTypes/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ContentTypeId, id.ContentTypeId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Content Type Id Base Type ID
func (id GroupIdSiteIdContentTypeIdBaseTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("contentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeId"),
		resourceids.StaticSegment("baseTypes", "baseTypes", "baseTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId1", "contentTypeId1"),
	}
}

// String returns a human-readable description of this Group Id Site Id Content Type Id Base Type ID
func (id GroupIdSiteIdContentTypeIdBaseTypeId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
		fmt.Sprintf("Content Type Id 1: %q", id.ContentTypeId1),
	}
	return fmt.Sprintf("Group Id Site Id Content Type Id Base Type (%s)", strings.Join(components, "\n"))
}
