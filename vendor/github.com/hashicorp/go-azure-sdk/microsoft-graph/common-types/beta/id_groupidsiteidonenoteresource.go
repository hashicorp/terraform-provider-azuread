package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdOnenoteResourceId{}

// GroupIdSiteIdOnenoteResourceId is a struct representing the Resource ID for a Group Id Site Id Onenote Resource
type GroupIdSiteIdOnenoteResourceId struct {
	GroupId           string
	SiteId            string
	OnenoteResourceId string
}

// NewGroupIdSiteIdOnenoteResourceID returns a new GroupIdSiteIdOnenoteResourceId struct
func NewGroupIdSiteIdOnenoteResourceID(groupId string, siteId string, onenoteResourceId string) GroupIdSiteIdOnenoteResourceId {
	return GroupIdSiteIdOnenoteResourceId{
		GroupId:           groupId,
		SiteId:            siteId,
		OnenoteResourceId: onenoteResourceId,
	}
}

// ParseGroupIdSiteIdOnenoteResourceID parses 'input' into a GroupIdSiteIdOnenoteResourceId
func ParseGroupIdSiteIdOnenoteResourceID(input string) (*GroupIdSiteIdOnenoteResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenoteResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenoteResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdOnenoteResourceIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdOnenoteResourceId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdOnenoteResourceIDInsensitively(input string) (*GroupIdSiteIdOnenoteResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdOnenoteResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdOnenoteResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdOnenoteResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.OnenoteResourceId, ok = input.Parsed["onenoteResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteResourceId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdOnenoteResourceID checks that 'input' can be parsed as a Group Id Site Id Onenote Resource ID
func ValidateGroupIdSiteIdOnenoteResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdOnenoteResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Onenote Resource ID
func (id GroupIdSiteIdOnenoteResourceId) ID() string {
	fmtString := "/groups/%s/sites/%s/onenote/resources/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.OnenoteResourceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Onenote Resource ID
func (id GroupIdSiteIdOnenoteResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("resources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("onenoteResourceId", "onenoteResourceId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Onenote Resource ID
func (id GroupIdSiteIdOnenoteResourceId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Onenote Resource: %q", id.OnenoteResourceId),
	}
	return fmt.Sprintf("Group Id Site Id Onenote Resource (%s)", strings.Join(components, "\n"))
}
